#!/bin/bash

# Ensure VERSION is provided
if [ -z "$VERSION" ]; then
  echo "VERSION environment variable is required!"
  exit 1
fi

go mod tidy

# Define Go OS and architecture combinations
declare -a os_array=("linux" "darwin" "windows")
declare -a arch_array=("amd64" "arm64")

# Directory to store the output binaries
OUTPUT_DIR="/output"

# Loop through each OS and architecture to build the binaries
for os in "${os_array[@]}"; do
  for arch in "${arch_array[@]}"; do
    # Set the binary output name
    BINARY_NAME="goftree-${os}-${arch}"

    # Append .exe for Windows binaries
    if [ "$os" == "windows" ]; then
      BINARY_NAME="${BINARY_NAME}.exe"
    fi

    echo "Building for ${os}/${arch}..."

    # Set GOOS and GOARCH for cross-compiling and build the binary
    GOOS=$os GOARCH=$arch go build -o "${OUTPUT_DIR}/${BINARY_NAME}" -ldflags "-X github.com/EdgeLordKirito/goFTree/internal/goFTree/version.Version=${VERSION}"

    # Check if the go build command failed
    if [ $? -ne 0 ]; then
      echo "Error: Build failed for ${os}/${arch}"
      exit 1  # Exit the script if build fails
    fi

    echo "Built binary for ${os}/${arch}: ${OUTPUT_DIR}/${BINARY_NAME}"
  done
done

echo "Build process completed!"
