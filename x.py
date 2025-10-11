import os

onedrive_path = r"D:\GoogleDrive"  # your local OneDrive folder


def is_empty_dir(path):
    # Return True if directory has no files (including hidden) in any depth
    for _, dirs, files in os.walk(path):
        if files:  # has any file at any depth
            return False
    return True


empty_dirs = []

for root, dirs, files in os.walk(onedrive_path):
    # ignore if this dir contains files
    if not files:
        # check if no subdirectory contains files either
        if is_empty_dir(root):
            empty_dirs.append(root)

# Print results
if empty_dirs:
    print("Empty folders (would NOT sync in Linux onedrive):\n")
    for d in empty_dirs:
        print(d)
    print(f"\nTotal empty folders: {len(empty_dirs)}")
else:
    print("No empty folders found.")
