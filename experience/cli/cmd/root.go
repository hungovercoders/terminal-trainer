package cmd

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "terminal-trainer",
    Short: "Interactive terminal command learning tool",
    Long: `Terminal Trainer helps you learn command line tools through 
interactive tutorials and quick reference cheat-sheets.

Supported tools:
- Linux commands (ls, cd, mkdir, etc.)
- Git version control  
- Docker containerization
- Kubectl Kubernetes management

Examples:
  terminal-trainer quiz --topic linux
  terminal-trainer cheat --topic git`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("🚀 Welcome to Terminal Trainer!")
        fmt.Println()
        fmt.Println("Use 'terminal-trainer --help' to see available commands")
        fmt.Println("Use 'terminal-trainer quiz' to start learning interactively")
        fmt.Println("Use 'terminal-trainer cheat' for quick reference guides")
    },
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}

func init() {
    // Global flags available to all subcommands
    rootCmd.PersistentFlags().StringP("topic", "t", "linux", "Topic to focus on (linux, git, docker, kubectl)")
    rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")
}