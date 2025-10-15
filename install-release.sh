#!/bin/bash
# Installation script for terminal-trainer CLI from GitHub releases

set -e

echo "🚀 Installing Terminal Trainer CLI..."

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64) ARCH="x86_64";;
    arm64|aarch64) ARCH="arm64";;
    i386|i686) ARCH="i386";;
    *) 
        echo "❌ Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

case $OS in
    linux) OS="Linux";;
    darwin) OS="Darwin";;
    *)
        echo "❌ Unsupported OS: $OS"
        echo "Please download manually from: https://github.com/hungovercoders/terminal-trainer/releases"
        exit 1
        ;;
esac

# Get latest release
echo "🔍 Getting latest release..."
LATEST_RELEASE=$(curl -s https://api.github.com/repos/hungovercoders/terminal-trainer/releases/latest | grep -o '"tag_name": "[^"]*' | grep -o '[^"]*$')

if [ -z "$LATEST_RELEASE" ]; then
    echo "❌ Could not get latest release information"
    exit 1
fi

echo "📦 Latest version: $LATEST_RELEASE"

# Download URL
DOWNLOAD_URL="https://github.com/hungovercoders/terminal-trainer/releases/download/${LATEST_RELEASE}/terminal-trainer_${OS}_${ARCH}.tar.gz"

echo "⬇️  Downloading terminal-trainer ${LATEST_RELEASE}..."
TEMP_DIR=$(mktemp -d)
cd "$TEMP_DIR"

if ! curl -L -f "${DOWNLOAD_URL}" -o terminal-trainer.tar.gz; then
    echo "❌ Download failed. Available releases:"
    curl -s https://api.github.com/repos/hungovercoders/terminal-trainer/releases/latest | grep "browser_download_url"
    exit 1
fi

# Extract
echo "📦 Extracting..."
tar -xzf terminal-trainer.tar.gz

# Install to /usr/local/bin (with sudo) or ~/bin (without sudo)
INSTALL_DIR=""
if [ -w "/usr/local/bin" ]; then
    INSTALL_DIR="/usr/local/bin"
    echo "📍 Installing to /usr/local/bin..."
elif command -v sudo &> /dev/null; then
    INSTALL_DIR="/usr/local/bin"
    echo "📍 Installing to /usr/local/bin (requires sudo)..."
    sudo mkdir -p "$INSTALL_DIR"
    sudo mv terminal-trainer "$INSTALL_DIR/"
else
    INSTALL_DIR="$HOME/bin"
    echo "📍 Installing to ~/bin..."
    mkdir -p "$INSTALL_DIR"
    mv terminal-trainer "$INSTALL_DIR/"
    
    # Add ~/bin to PATH if not already there
    for shell_rc in ~/.bashrc ~/.zshrc ~/.profile; do
        if [ -f "$shell_rc" ] && ! grep -q "$HOME/bin" "$shell_rc"; then
            echo 'export PATH="$HOME/bin:$PATH"' >> "$shell_rc"
            echo "   ✅ Added ~/bin to PATH in $shell_rc"
        fi
    done
fi

if [ "$INSTALL_DIR" = "/usr/local/bin" ] && [ ! -f "/usr/local/bin/terminal-trainer" ]; then
    sudo mv terminal-trainer "$INSTALL_DIR/"
fi

# Make executable
if [ "$INSTALL_DIR" = "/usr/local/bin" ]; then
    sudo chmod +x "$INSTALL_DIR/terminal-trainer"
else
    chmod +x "$INSTALL_DIR/terminal-trainer"
fi

# Create aliases
echo "🔗 Creating convenient aliases..."
for shell_rc in ~/.bashrc ~/.zshrc; do
    if [ -f "$shell_rc" ]; then
        # Remove existing aliases
        sed -i '/alias terminaltrainer=/d' "$shell_rc" 2>/dev/null || true
        sed -i '/alias tt=/d' "$shell_rc" 2>/dev/null || true
        
        # Add new aliases
        echo 'alias terminaltrainer="terminal-trainer"' >> "$shell_rc"
        echo 'alias tt="terminal-trainer"' >> "$shell_rc"
        echo "   ✅ Updated $shell_rc"
    fi
done

# Cleanup
cd /
rm -rf "$TEMP_DIR"

echo ""
echo "🎉 Installation complete!"
echo ""
echo "You can now use the CLI with:"
echo "  terminal-trainer --help        # Full name"
echo "  terminaltrainer --help         # Alias"  
echo "  tt --help                      # Short alias"
echo ""
echo "💡 Restart your terminal or run 'source ~/.bashrc' to use aliases"
echo ""
echo "🚀 Try it out:"
echo "  tt quiz --topic linux"
echo "  terminaltrainer cheat --topic git"