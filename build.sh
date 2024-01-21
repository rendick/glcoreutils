#!/bin/bash

# Set the source and build directories
src_directory="./src"
build_directory="./build"

# Create the build directory if it doesn't exist
mkdir -p "$build_directory"

# Change to the source directory
cd "$src_directory" || exit

# Find all .go files and build them
for file in *.go; do
    if [ -f "$file" ]; then
        echo "Building: $file"
        go build -o "$build_directory/$(basename "$file" .go)" "$file"
        if [ $? -eq 0 ]; then
            echo "Build successful: $file"
        else
            echo "Build failed: $file"
        fi
    fi
done
