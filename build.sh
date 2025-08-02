#!/bin/bash

# Define the output binary name
APP_NAME="workbuddy-api"

# Define the main Go file
MAIN_FILE="main.go"

# Define the target OS and architecture for cross-compilation
# Use 'linux' for deployment on most servers
GOOS_TARGET="linux"
GOARCH_TARGET="amd64"

echo "Building Go Fiber application for $GOOS_TARGET/$GOARCH_TARGET..."

# Clean up any old builds
rm -f $APP_NAME

# Build the application with optimized flags for production
# -o: specifies the output file name
# -ldflags: removes debug symbols to reduce binary size
# GOOS and GOARCH: environment variables for cross-compilation
GOOS=$GOOS_TARGET GOARCH=$GOARCH_TARGET go build -o $APP_NAME -ldflags="-s -w" $MAIN_FILE

# Check if the build was successful
if [ $? -eq 0 ]; then
  echo "✅ Build successful! Executable created: ./$APP_NAME"
  echo "To run your application, use: ./$APP_NAME"
else
  echo "❌ Build failed!"
  exit 1
fi