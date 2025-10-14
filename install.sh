#!/bin/bash
# Installation script for terminal-trainer CLI

set -e

echo "🚀 Installing Terminal Trainer CLI..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go first: https://golang.org/doc/install"
    exit 1
fi

# Get the project directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
CLI_DIR="$SCRIPT_DIR/experience/cli"

# Install the CLI
echo "📦 Building and installing CLI..."
cd "$CLI_DIR"
go install .

# Create aliases
GOBIN=$(go env GOBIN)
if [ -z "$GOBIN" ]; then
    GOBIN=$(go env GOPATH)/bin
fi

echo "🔗 Creating convenient aliases..."
ln -sf "$GOBIN/terminal-trainer-cli" "$GOBIN/terminaltrainer" 2>/dev/null || true
ln -sf "$GOBIN/terminal-trainer-cli" "$GOBIN/tt" 2>/dev/null || true

# Add shell aliases
echo "🐚 Adding shell aliases..."
for shell_rc in ~/.bashrc ~/.zshrc; do
    if [ -f "$shell_rc" ]; then
        # Remove existing aliases
        sed -i '/alias terminaltrainer=/d' "$shell_rc" 2>/dev/null || true
        sed -i '/alias tt=/d' "$shell_rc" 2>/dev/null || true
        
        # Add new aliases
        echo 'alias terminaltrainer="terminal-trainer-cli"' >> "$shell_rc"
        echo 'alias tt="terminal-trainer-cli"' >> "$shell_rc"
        echo "   ✅ Updated $shell_rc"
    fi
done

echo ""
echo "🎉 Installation complete!"
echo ""
echo "You can now use the CLI with:"
echo "  terminal-trainer-cli --help    # Full name"
echo "  terminaltrainer --help         # Alias"  
echo "  tt --help                      # Short alias"
echo ""
echo "💡 Restart your terminal or run 'source ~/.bashrc' to use aliases"
echo ""
echo "🚀 Try it out:"
echo "  tt quiz --topic linux"
echo "  terminaltrainer cheat --topic git"