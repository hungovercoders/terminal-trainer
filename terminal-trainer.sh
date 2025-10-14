#!/bin/bash
# Convenience script to run terminal-trainer from anywhere in the project

# Get the directory where this script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
CLI_DIR="$SCRIPT_DIR/experience/cli"

# Check if binary exists, if not try to build it
if [ ! -f "$CLI_DIR/terminal-trainer" ]; then
    echo "Binary not found, building..."
    cd "$CLI_DIR" && go build -o terminal-trainer main.go
    if [ $? -ne 0 ]; then
        echo "Build failed!"
        exit 1
    fi
fi

# Run the CLI with all passed arguments
"$CLI_DIR/terminal-trainer" "$@"