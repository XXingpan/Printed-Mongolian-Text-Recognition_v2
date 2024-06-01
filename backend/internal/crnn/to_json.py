import json

def combine_lists(positions, results, output_path):
    """
    Combines the positions list with the results list into a new list and writes it to a file.

    Args:
        positions (list): A list of lists containing tuples that represent the positions of the word images.
        results (list): A list of tuples containing filenames and their associated text.
        output_path (str): Path to the file where the combined data will be saved.
    """
    # Flatten results into a dictionary for easy access
    result_dict = {filename: text for filename, text in results}

    # Remove empty lists from positions
    positions = [line for line in positions if line]

    combined_list = []

    for line in positions:
        for idx, pos in enumerate(line):
            filename = f"{idx + 1}.jpg"
            text = result_dict.get(filename, "")
            pos_list = list(pos)

            # Add position and text directly to the combined list
            combined_list.append([pos_list, text])

    # Convert to JSON string
    result_json = json.dumps(combined_list, indent=2, ensure_ascii=False)
    # print(combined_list)
    # Write JSON string to file
    try:
        with open(output_path, 'w', encoding="utf-8") as file:
            file.write(result_json)
            print(f"Successfully written to '{output_path}'")
    except IOError as e:
        print(f"Error: {e}")
    except Exception as e:
        print(f"Error: An unexpected error occurred: {e}")

# def main():
#     # Lists of positions and results (mocked data for illustration)
#     positions = [[(1, 5, 33, 56), (1, 90, 33, 128), (1, 231, 33, 87)], []]
#     results = [('1.jpg', 'ᠪᠠᠷᠢ'), ('2.jpg', 'ᠵᠤᠸᠵᠢᠠ'), ('3.jpg', 'ᠪᠢᠨᠰᠬ')]
#
#     # Output file path
#     output_path = './result_text/ocr_result.json'
#
#     # Combine lists and write to file
#     combine_lists(positions, results, output_path)
#
# if __name__ == "__main__":
#     main()
