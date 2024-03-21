import os

def search_word_in_files(directory, word):
    found_files = []

    # Iterate through all files and directories
    for root, _, files in os.walk(directory):
        for file in files:
            # Construct the file path
            file_path = os.path.join(root, file)
            try:
                # Check if the file is a binary file
                if file.endswith(('.tgz', '.tar', '.zip')):
                    # Skip binary files
                    continue

                # Open the file in read mode
                with open(file_path, 'r') as f:
                    # Read the file line by line
                    for line in f:
                        # Check if the word exists in the line
                        if word in line:
                            # Add the file path to the list of found files
                            found_files.append(file_path)
                            break  # Stop searching this file once the word is found
            except Exception as e:
                print(f"Error reading file {file_path}: {e}")

    return found_files

if __name__ == "__main__":
    word = input("Enter the word to search for: ")

    found_files = search_word_in_files("/home/mouhameden/Downloads/cells-main", word)

    if found_files:
        print("Word found in the following files:")
        for file_path in found_files:
            print(file_path)
    else:
        print("Word not found in any files.")
