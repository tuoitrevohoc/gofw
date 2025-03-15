#!/bin/bash

# Prompt for package name
read -p "Please enter the package name (e.g. your-github-handle/your-project-name): " PACKAGE_NAME

# Check if the package name is empty
if [ -z "$PACKAGE_NAME" ]; then
    echo "Package name cannot be empty."
    exit 1
fi

# Check if the current folder is empty
if [ "$(ls -A .)" ]; then
    echo "Current folder is not empty"
    exit 1
fi

# Create a temporary directory
TEMP_DIR=$(mktemp -d)

# Checkout the repository into the temporary directory
git clone https://github.com/tuoitrevohoc/gofw-template "$TEMP_DIR"

echo "Contents of temporary directory:"
ls -la "$TEMP_DIR"

# Replace all occurrences of the string in the files
find "$TEMP_DIR" -type f -exec echo "Executing: sed -i '' 's|tuoitrevohoc/gofw-template|$PACKAGE_NAME|g' {}" \; -exec sed -i '' "s|github.com/tuoitrevohoc/gofw-template|$PACKAGE_NAME|g" {} \;


# Copy all files and folders except the .git folder to the current directory
rsync -av --exclude='.git' "$TEMP_DIR/" ./

# Print success message
echo "Congratulations! Open the project in VS Code"

# Clean up the temporary directory
rm -rf "$TEMP_DIR"
