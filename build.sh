#!/bin/bash

set -e

BINARY_NAME="go-desent"
REMOTE_PATH="/home/$VM_USERNAME/"

echo "Building..."
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o "$BINARY_NAME"

echo "Uploading..."
scp -i "$SSH_PATH" "$BINARY_NAME" "$VM_USERNAME@$VM_IP:$REMOTE_PATH"

echo "Done 🚀"