# Terminal Trainer Agents Hints

Tool to help people learn how to use the terminal for a range of tools.
This can take the form of cheat-sheets, interactive tutorials, or other resources.
Example command line tools that people will want to learn include:

- git
- docker
- kubectl
- linux commands

The tool will be built in go.
The tool will leverage cobra-cli for the command line interface.
The tool will leverage bubbletea for interactive command line experiences.
Test driven development should be leveraged.

The CLI will be built in go so that installation is simple and it can be used on any platform.
The logic and state will be kept in the behaviour directory separate to the CLI itself so that it can be used in other contexts, such as a web app or a desktop app.
The CLI is an experience and will be kept in the experience directory separate from the behaviour.

## Features

- Interactive tutorials for common command line tools
- Cheat-sheets for quick reference

This repository uses [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) for commit messages.
