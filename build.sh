#!/bin/bash

# Ensure VERSION is provided
if [ -z "$VERSION" ]; then
  echo "VERSION environment variable is required!"
  exit 1
fi

go mod tidy

# Directory to store the output binaries
OUTPUT_DIR="/output"

# Declare an associative array to store binary names and their checksums
declare -A checksum_map

# Function to build Go binary for a specific OS and architecture
build_go_binary() {
  local os=$1
  local arch=$2

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

  # Calculate the checksum for the built binary and store it in the checksum_map
  CHECKSUM=$(sha256sum "${OUTPUT_DIR}/${BINARY_NAME}" | awk '{ print $1 }')
  checksum_map["${BINARY_NAME}"]=$CHECKSUM
}

# Build for all architectures except 386
for os in "linux" "darwin" "windows"; do
  for arch in "amd64" "arm64"; do
    build_go_binary $os $arch
  done
done

# Now handle the 386 architecture separately
declare -a os_386_array=("linux" "windows")
for os in "${os_386_array[@]}"; do
  build_go_binary $os "386"
done

# Create the checksums.txt file and write the binary names and their checksums
echo "Creating checksums file..."

CHECKSUM_FILE="${OUTPUT_DIR}/checksums.txt"
> "$CHECKSUM_FILE"  # Clear the file before writing

# Iterate over the checksum_map and write to the checksums.txt file
for binary_name in "${!checksum_map[@]}"; do
  echo "$binary_name" >> "$CHECKSUM_FILE"
  echo "${checksum_map[$binary_name]}" >> "$CHECKSUM_FILE"
  echo "" >> "$CHECKSUM_FILE"  # Add an empty line between entries
done

echo "Checksums file created: ${CHECKSUM_FILE}"

echo "Build process completed!"
