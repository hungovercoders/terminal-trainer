# terminal-trainer

Tool to help people learn how to use the terminal for a range of tools

## Installation

### Quick Install (Recommended)

```bash
# Clone and install globally
git clone https://github.com/hungovercoders/terminal-trainer
cd terminal-trainer
./install.sh
```

After installation, you can use:

- `terminaltrainer` or `tt` (aliases)  
- `terminal-trainer-cli` (full name)

### Manual Install

```bash
# Install directly with go
cd experience/cli
go install .

# Create your own aliases
echo 'alias tt="terminal-trainer-cli"' >> ~/.bashrc
source ~/.bashrc
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
