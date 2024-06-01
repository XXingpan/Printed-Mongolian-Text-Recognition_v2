import cv2
import numpy as np
import os
import shutil


def clear_cache(folder_path):
    """
    Clears all files and subdirectories in the specified folder.

    Args:
        folder_path (str): The path to the folder to clear.
    """
    # Ensure the folder exists
    if not os.path.exists(folder_path):
        print(f"Folder '{folder_path}' does not exist.")
        return

    if not os.path.isdir(folder_path):
        print(f"'{folder_path}' is not a folder path.")
        return

    files = os.listdir(folder_path)

    for file_name in files:
        file_path = os.path.join(folder_path, file_name)
        try:
            if os.path.isfile(file_path):
                os.unlink(file_path)
            elif os.path.isdir(file_path):
                shutil.rmtree(file_path)
        except Exception as e:
            print(f"Error deleting '{file_path}': {e}")


def extract_peek_ranges_from_array(array_vals, minimum_val=1000, minimum_range=2):
    """
    Extracts ranges of peak values from an array.

    Args:
        array_vals (np.array): Array to process.
        minimum_val (int): Minimum value to consider a peak.
        minimum_range (int): Minimum range length for a peak.

    Returns:
        list: List of tuples with start and end indices of peaks.
    """
    start_i = None
    end_i = None
    peek_ranges = []
    for i, val in enumerate(array_vals):
        if val > minimum_val and start_i is None:
            start_i = i
        elif val > minimum_val and start_i is not None:
            pass
        elif val < minimum_val and start_i is not None:
            end_i = i
            if end_i - start_i >= minimum_range:
                peek_ranges.append((start_i, end_i))
            start_i = None
            end_i = None
        elif val < minimum_val and start_i is None:
            pass
        else:
            raise ValueError("cannot parse this case...")
    return peek_ranges


def process_image(IMAGE_UPLOAD_PATH):
    """
    Processes an image to identify text lines and words, saving results to files and returning positions.

    Args:
        IMAGE_UPLOAD_PATH (str): Path to the image to process.

    Returns:
        list: List of positions of text lines and words.
    """
    # Clear cache directories
    clear_cache("./internal/crnn/preprocess/word_image")
    clear_cache("./internal/crnn/preprocess/label_image")
    clear_cache("./internal/crnn/preprocess/result")

    # Read and process the image
    image_color = cv2.imread(IMAGE_UPLOAD_PATH)
    new_shape = (image_color.shape[1], image_color.shape[0])
    image_color = cv2.resize(image_color, new_shape)
    image = cv2.cvtColor(image_color, cv2.COLOR_BGR2GRAY)
    adaptive_threshold = cv2.adaptiveThreshold(
        image,
        255,
        cv2.ADAPTIVE_THRESH_GAUSSIAN_C,
        cv2.THRESH_BINARY_INV, 11, 2)

    horizontal_sum = np.sum(adaptive_threshold, axis=0)

    peek_ranges = extract_peek_ranges_from_array(horizontal_sum)

    line_seg_adaptive_threshold = np.copy(adaptive_threshold)
    for i, peek_range in enumerate(peek_ranges):
        x = peek_range[0]
        y = 0
        w = peek_range[1]
        h = line_seg_adaptive_threshold.shape[0]
        pt1 = (x, y)
        pt2 = (x + w, y + h)
        cv2.rectangle(line_seg_adaptive_threshold, pt1, pt2, 255)

    vertical_peek_ranges2d = []
    for peek_range in peek_ranges:
        start_x = 0
        end_x = line_seg_adaptive_threshold.shape[0]
        line_img = adaptive_threshold[start_x:end_x, peek_range[0]:peek_range[1]]
        vertical_sum = np.sum(line_img, axis=1)
        vertical_peek_ranges = extract_peek_ranges_from_array(vertical_sum, minimum_val=2, minimum_range=1)
        vertical_peek_ranges2d.append(vertical_peek_ranges)

    cnt = 1
    color = (0, 0, 255)
    min_word_height = 15  # Minimum height threshold for words

    all_positions = []  # Store all image positions

    for i, peek_range in enumerate(peek_ranges):
        line_positions = []  # Store positions for each line
        for vertical_range in vertical_peek_ranges2d[i]:
            x = peek_range[0]
            y = vertical_range[0]
            w = peek_range[1] - x
            h = vertical_range[1] - y

            if h >= min_word_height:
                patch = adaptive_threshold[y:y + h, x:x + w]
                cv2.imwrite(f'.\\internal\\crnn\\preprocess\\word_image\\{cnt}.jpg', patch)
                cnt += 1

                pt1 = (x, y)
                pt2 = (x + w, y + h)
                cv2.rectangle(image_color, pt1, pt2, color)

                line_positions.append((x, y, w, h))

        all_positions.append(line_positions)

    # Save final labeled image
    cv2.imwrite('./internal/crnn/preprocess/label_image/label.jpg', image_color)

    return all_positions
