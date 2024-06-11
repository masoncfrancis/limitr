#!/bin/bash

# Check if gum is installed
if ! command -v gum &> /dev/null; then
    echo "This script uses Gum to make things look nice."
    echo "Gum is not installed. Please install Gum before running this script."
    echo "You can find installation instructions for Gum at https://github.com/charmbracelet/gum?tab=readme-ov-file#installation"
    exit 1
fi

# Check if the version argument is provided
if [ -z "$1" ]; then
    echo "Usage: $0 <version>"
    exit 1
fi

# Project and version details
PROGRAM_NAME="limitr"
VERSION="$1"
OUTPUT_DIR="./out"
SOURCE_DIR="./cmd/limitr"

# List of OS and Architecture combinations
platforms=(
    "linux amd64"
    "linux arm64"
    "linux arm"
    "darwin amd64"
    "darwin arm64"
    "windows amd64"
    "windows arm64"
)

platforms_string=$(printf "\n - %s" "${platforms[@]}")

gum confirm "We will now build limitr for the following platforms:$platforms_string" --affirmative="Continue" --negative="Exit" || { echo "Limitr build cancelled"; exit 1; }

# Create the output directory if it doesn't exist
mkdir -p "$OUTPUT_DIR"

# Build the project for each platform
for platform in "${platforms[@]}"; do
    os_arch=(${platform})
    os=${os_arch[0]}
    arch=${os_arch[1]}
    output_name="$OUTPUT_DIR/${PROGRAM_NAME}_${VERSION}_${os}_${arch}"

    # Add .exe extension for Windows
    if [ "$os" = "windows" ]; then
        output_name+='.exe'
    fi

    # Set the environment variables for the platform to build for
    GOOS="$os" GOARCH="$arch"
    # build the project
    gum spin --title "Building for $os $arch..." --spinner minidot -- go build -o "$output_name" "$SOURCE_DIR"

    if [ $? -ne 0 ]; then
        echo "An error occurred while building for $os $arch"
        exit 1
    fi

    # Make the output file executable if not on Windows
    if [ "$os" != "windows" ]; then
        chmod +x "$output_name"
    fi
done

echo "Build completed successfully!"
