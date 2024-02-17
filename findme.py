import os
import glob

def find_image(filename, directory):
    # Iterate through all files and directories within the given directory
    for root, dirs, files in os.walk(directory):
        # Check if the filename exists in the list of files
        if filename in files:
            # Return the full path of the image file if found
            return os.path.join(root, filename)
    # Return None if the image file is not found
    return None

# Directory to search within
search_directory = '/'

# Image filename to search for
image_filename = 'LoginBoxLogo.png'

# Call the function to find the image file
image_path = find_image(image_filename, search_directory)

if image_path:
    print(f"Image found at: {image_path}")
else:
    print("Image not found.")
