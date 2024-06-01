
import os
import numpy as np
import cv2
import torch
import argparse
from crnn import CRNN
from normalize import process_images
from seg import process_image, clear_cache
from to_json import combine_lists

line_size = 32
vocab = list(range(0x1800, 0x180F)) + list(range(0x1810, 0x181A)) + list(range(0x1820, 0x1879)) + list(
    range(0x1880, 0x18AB)) + [0x202F]
vocab = "B " + "".join([chr(v) for v in vocab])
idx2char = {idx: char for idx, char in enumerate(vocab)}


def load_model_from_checkpoint(checkpoint_file_name, use_gpu=False):
    """Load a pretrained CRNN model."""
    model = CRNN(line_size, 1, len(vocab), 256)
    checkpoint = torch.load(checkpoint_file_name, map_location='cpu' if not use_gpu else None)
    model.load_state_dict(checkpoint['state_dict'])
    model.float()
    model.eval()
    model = model.cuda() if use_gpu else model.cpu()
    return model


def to_text(tensor, max_length=None, remove_repetitions=False):
    """Convert a tensor to text."""
    sentence = ''
    sequence = tensor.cpu().detach().numpy()
    for i in range(len(sequence)):
        if max_length is not None and i >= max_length:
            continue
        char = idx2char[sequence[i]]
        if char != 'B':
            if remove_repetitions and i != 0 and char == idx2char[sequence[i - 1]]:
                pass
            else:
                sentence += char
    return sentence


def ocr(image, model):
    """OCR for a single word image."""
    torch.set_grad_enabled(False)

    resized_img = np.array(np.rot90(image))
    resized_img = cv2.resize(resized_img, (line_size, resized_img.shape[0]))

    inputs = torch.from_numpy(resized_img / 255).float().unsqueeze(0).unsqueeze(0)
    outputs = model(inputs)
    prediction = outputs.softmax(2).max(2)[1]

    return to_text(prediction[:, 0], remove_repetitions=True)


def ocr_from_directory(directory, checkpoint_file_name, use_gpu=False):
    """Perform OCR on all images in the specified directory."""
    model = load_model_from_checkpoint(checkpoint_file_name, use_gpu)
    results = []

    for filename in os.listdir(directory):
        if filename.endswith('.png') or filename.endswith('.jpg'):
            image_path = os.path.join(directory, filename)
            image = cv2.imread(image_path, 0)
            if image is not None:
                text = ocr(image, model)
                results.append((filename, text))

    return results


if __name__ == '__main__':
    IMAGE_UPLOAD_PATH="./internal/crnn/upload/"
    upload_image_path=os.listdir(IMAGE_UPLOAD_PATH)
    # print(upload_image_path[0])
    all_positions = process_image(IMAGE_UPLOAD_PATH+upload_image_path[0])

    # Define the input and output folders
    input_folder = "./internal/crnn/preprocess/word_image"

    output_folder = "./internal/crnn/preprocess/result"

    # print(output_folder)
    # Call the function to process the images
    process_images(input_folder, output_folder)
    # 检查primary_folder目录是否为空
    if not os.listdir("./internal/crnn/preprocess/result"):
        output_folder = "./internal/crnn/preprocess/label_image"

    else:
        output_folder = "./internal/crnn/preprocess/result"
    results = ocr_from_directory(output_folder,"./internal/crnn/crnn_model.pth", use_gpu=False)
    combine_lists(all_positions,results,'./internal/crnn/result_text/ocr_result.json')

