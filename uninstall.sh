#!/bin/bash
# Uninstallation script for terminal-trainer CLI

echo "🗑️  Uninstalling Terminal Trainer CLI..."

# Get Go bin directory
GOBIN=$(go env GOBIN 2>/dev/null)
if [ -z "$GOBIN" ]; then
    GOBIN=$(go env GOPATH 2>/dev/null)/bin
fi

# Remove binaries
echo "📦 Removing binaries..."
rm -f "$GOBIN/terminal-trainer-cli" 2>/dev/null
rm -f "$GOBIN/terminaltrainer" 2>/dev/null  
rm -f "$GOBIN/tt" 2>/dev/null

# Remove shell aliases
echo "🐚 Removing shell aliases..."
for shell_rc in ~/.bashrc ~/.zshrc; do
    if [ -f "$shell_rc" ]; then
        sed -i '/alias terminaltrainer=/d' "$shell_rc" 2>/dev/null || true
        sed -i '/alias tt=/d' "$shell_rc" 2>/dev/null || true
        echo "   ✅ Cleaned $shell_rc"
    fi
done

echo ""
echo "🎉 Uninstallation complete!"
echo "💡 Restart your terminal or run 'source ~/.bashrc' to apply changes"