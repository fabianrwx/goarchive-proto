#!/bin/bash
set -e # Exit immediately if a command exits with a non-zero status

SERVICE_NAME=$1
RELEASE_VERSION=$2

# Install necessary tools
sudo apt-get update
sudo apt-get install -y protobuf-compiler golang-goprotobuf-dev
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate protobuf and gRPC files
protoc --go_out=./golang --go_opt=paths=source_relative \
  --go-grpc_out=./golang --go-grpc_opt=paths=source_relative \
  ./${SERVICE_NAME}/*.proto

# Ensure the generated service directory exists before moving into it
if [ ! -d "./golang/${SERVICE_NAME}" ]; then
  echo "Directory ./golang/${SERVICE_NAME} does not exist. Generation might have failed."
  exit 1
fi

# Move into the service directory to generate the go mod file
cd golang/${SERVICE_NAME}

# Initialize Go module if not already initialized
if [ ! -f "go.mod" ]; then
  go mod init github.com/fabianrwx/myshopstatistics-proto/golang/${SERVICE_NAME} || true
fi

go mod tidy

cd ../../

# Configure git
git config --global user.email "example@gmail.com"
git config --global user.name "example_user"

# Add and commit changes
git commit -am "proto update" || true

echo "golang/${SERVICE_NAME}/${RELEASE_VERSION}"
# Force replace the tag if it already exists since this will be the latest execution of proto-gen script

# Create and push a tag
git tag -fa "golang/${SERVICE_NAME}/${RELEASE_VERSION}" -m "golang/${SERVICE_NAME}/${RELEASE_VERSION}"
git push -f origin refs/tags/golang/${SERVICE_NAME}/${RELEASE_VERSION} 