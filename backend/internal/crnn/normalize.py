from PIL import Image, ImageOps
import os


def process_images(input_folder, output_folder):
    """
    Processes images in the input folder by converting them to grayscale,
    inverting the colors, resizing, and saving them to the output folder.

    Parameters:
        input_folder (str): The path to the input folder containing images.
        output_folder (str): The path to the output folder where processed images will be saved.
    """
    # Ensure the output folder exists
    if not os.path.exists(output_folder):
        os.makedirs(output_folder)

    # Traverse through all image files in the input folder
    for filename in os.listdir(input_folder):
        if filename.endswith(".png") or filename.endswith(".jpg"):
            # Open image
            img_path = os.path.join(input_folder, filename)
            img = Image.open(img_path)

            # Convert to grayscale and set resolution
            img = img.convert("L")

            # Invert colors: black to white, white to black
            img = ImageOps.invert(img)

            # Resize image
            width = 32
            height = int(width * img.height / img.width)
            img = img.resize((width, height), Image.LANCZOS)

            # Create a new white image as a background
            new_img = Image.new("L", (width, height), color="white")

            # Paste the resized image into the center of the new image
            offset = ((width - img.width) // 2, (height - img.height) // 2)
            new_img.paste(img, offset)

            # Save the processed image
            output_path = os.path.join(output_folder, filename)
            new_img.save(output_path, dpi=(300, 300))

            print(f"Processed: {filename}")


# def main():
#     # Define the input and output folders
#     input_folder = "./word_image"
#     output_folder = "./result"
#
#     # Call the function to process the images
#     process_images(input_folder, output_folder)
#
#
# # Run the main function if the script is executed directly
# if __name__ == "__main__":
#     main()
#
# print("All images processed successfully!")
