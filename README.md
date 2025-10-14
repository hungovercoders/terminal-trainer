# terminal-trainer

Tool to help people learn how to use the terminal for a range of tools

## Build

```bash
# Build the CLI
cd experience/cli
go build -o terminal-trainer main.go

# Or use the convenience script from project root (auto-builds if needed)
./terminal-trainer.sh --help
```

## See all options

```bash
# Using convenience script from project root
./terminal-trainer.sh --help

# Or from CLI directory
cd experience/cli && ./terminal-trainer --help

# Or run directly without building  
cd experience/cli && go run . --help
```

## Interactive quiz

```bash
# Using convenience script
./terminal-trainer.sh quiz --topic linux --count 5 --verbose

# Or from CLI directory  
cd experience/cli && ./terminal-trainer quiz --topic linux --count 5 --verbose
```

## Quick reference

```bash
# Using convenience script
./terminal-trainer.sh cheat --topic linux

# Or from CLI directory
cd experience/cli && ./terminal-trainer cheat --topic linux
```

## Basic welcome

```bash
# Using convenience script
./terminal-trainer.sh

# Or from CLI directory
cd experience/cli && ./terminal-trainer
```
