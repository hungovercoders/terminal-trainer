# terminal-trainer

Tool to help people learn how to use the terminal for a range of tools

## Installation

### Quick Install Script

**Linux/macOS:**

```bash
curl -sSL https://raw.githubusercontent.com/hungovercoders/terminal-trainer/main/install-release.sh | bash
```

**Windows (PowerShell):**

```powershell
irm https://raw.githubusercontent.com/hungovercoders/terminal-trainer/main/install-windows.ps1 | iex
```

**Windows (Git Bash/WSL):**

```bash
curl -sSL https://raw.githubusercontent.com/hungovercoders/terminal-trainer/main/install-release.sh | bash
```

### Download Binary (Manual)

1. Go to [Releases](https://github.com/hungovercoders/terminal-trainer/releases)
2. Download the binary for your OS/architecture:
   - `terminal-trainer_Linux_x86_64.tar.gz` (Linux)
   - `terminal-trainer_Darwin_x86_64.tar.gz` (macOS Intel)
   - `terminal-trainer_Darwin_arm64.tar.gz` (macOS Apple Silicon)
   - `terminal-trainer_Windows_x86_64.zip` (Windows)
3. Extract and add to your PATH:

```bash
# Linux/macOS
tar -xzf terminal-trainer_*.tar.gz
sudo mv terminal-trainer /usr/local/bin/

# Windows (PowerShell)
Expand-Archive terminal-trainer_Windows_x86_64.zip -DestinationPath C:\tools\
# Add C:\tools\ to your PATH environment variable
```

### Go Install (For Go Users)

```bash
go install github.com/hungovercoders/terminal-trainer/experience/cli@latest
```

### Build from Source

```bash
git clone https://github.com/hungovercoders/terminal-trainer
cd terminal-trainer/experience/cli
go build -o terminal-trainer main.go
sudo mv terminal-trainer /usr/local/bin/
```

### Uninstall

```bash
./uninstall.sh
```

## Build

```bash
# Build the CLI
cd experience/cli
go build -o terminal-trainer main.go

# Or use the convenience script from project root (auto-builds if needed)
./terminal-trainer.sh --help
```

## Usage Examples

After installation, use any of these commands:

```bash
# Short and sweet! 
tt --help
tt quiz --topic linux --count 5
tt cheat --topic git

# Full alias
terminaltrainer --help
terminaltrainer quiz --topic docker

# Full binary name  
terminal-trainer-cli --help
```

## Development Usage

If you're working on the code locally:

```bash
# Using convenience script
./terminal-trainer.sh --help

# Or from CLI directory
cd experience/cli && ./terminal-trainer --help

# Or run directly without building  
cd experience/cli && go run . --help
```

## Available Topics

- `linux` - Basic Linux commands (ls, cd, mkdir, etc.)
- `git` - Git version control commands  
- `docker` - Docker containerization commands
- `kubectl` - Kubernetes management commands
- `vscode` - VS Code shortcuts and commands

## Available Commands

- `quiz` - Interactive learning with questions
- `cheat` - Quick reference guide
- `help` - Show help information
