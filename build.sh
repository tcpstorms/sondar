#!/bin/bash

# Set the output directory for the compiled binary
OUTPUT_DIR="./bin"

# Set the name of the compiled binary
BINARY_NAME="nrat"

# Automatically detect and set the target operating system and target architecture
GOOS=$(go env GOOS)
GOARCH=$(go env GOARCH)

# Set the build type to "development" by default
BUILD_TYPE="development"

# Check the value of the first argument
case "$1" in
  "development")
    # If the first argument is "development", add debug flags and symbols to the compiled binary
    GO_FLAGS="-gcflags='all=-N -l'"
    GOFLAGS="-ldflags='-N -l'"
    ;;
  "production")
    # If the first argument is "production", remove symbol and debugging information from the compiled binary to reduce the file size
    GO_FLAGS="-ldflags=-s -ldflags=-w"
    # Enable aggressive optimization, optimize the code for the host machine's microarchitecture,
    # and use pipes instead of temporary files when invoking the assembler and linker
    GOFLAGS="-gcflags=-O3 -march=native -pipe"
    # Set the build type to "production"
    BUILD_TYPE="production"
    ;;
esac

mkdir -p "$OUTPUT_DIR"
printf "%s\n" "Building $BINARY_NAME for $GOOS/$GOARCH ($BUILD_TYPE build)"
GOOS="$GOOS" GOARCH="$GOARCH" go build "$GO_FLAGS" "$GOFLAGS" -o "$OUTPUT_DIR/$BINARY_NAME"
