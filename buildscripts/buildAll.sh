#!/bin/bash

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
)

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

    echo "Building for $os $arch..."

    # Set the environment variables and build the project
    GOOS="$os" GOARCH="$arch" go build -o "$output_name" "$SOURCE_DIR"

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
