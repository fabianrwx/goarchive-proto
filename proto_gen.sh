#!/bin/bash

set -e # Exit immediately if a command exits with a non-zero status

# Create base directories if they do not exist
mkdir -p golang

# Find all directories containing .proto files
PROTO_DIRS=$(find . -type f -name "*.proto" -exec dirname {} \; | sort -u)

# Generate protobuf and gRPC files for Go in the appropriate directories
for dir in $PROTO_DIRS; do

    # Remove leading ./ from directory name
    dir_clean=${dir#./}

    # Create output directory if it does not exist
    output_dir="${dir_clean}"

    mkdir -p "$output_dir"

    # Generate protobuf and gRPC files for Golang
    protoc --go_out="golang" --go_opt=paths=source_relative \
           --go-grpc_out="golang" --go-grpc_opt=paths=source_relative \
           "${dir_clean}"/*.proto

    # Run go mod tidy in the output directory
    pushd "golang/$output_dir" > /dev/null
    go mod init || true
    go mod tidy
    popd > /dev/null

done

echo "**** Successfully generated protobuf and gRPC files for Go ****"
