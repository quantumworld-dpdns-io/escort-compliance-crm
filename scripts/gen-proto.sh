#!/bin/bash
set -euo pipefail

echo "=== Generating Protobuf Code ==="

PROTO_DIR="packages/proto"
OUT_DIR="packages/proto/gen"

mkdir -p "$OUT_DIR"

# Go
protoc --go_out="$OUT_DIR/go" --go_opt=paths=source_relative \
    --go-grpc_out="$OUT_DIR/go" --go-grpc_opt=paths=source_relative \
    -I "$PROTO_DIR" \
    $(find "$PROTO_DIR" -name "*.proto")

# Python
python -m grpc_tools.protoc \
    --python_out="$OUT_DIR/python" \
    --grpc_python_out="$OUT_DIR/python" \
    -I "$PROTO_DIR" \
    $(find "$PROTO_DIR" -name "*.proto")

# TypeScript
npx ts-proto \
    --protoPath="$PROTO_DIR" \
    --out="$OUT_DIR/typescript" \
    $(find "$PROTO_DIR" -name "*.proto" -exec basename {} \;)

echo "✓ Protobuf code generated"
