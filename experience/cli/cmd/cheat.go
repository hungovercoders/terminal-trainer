package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"terminal-trainer/behaviour/terminaltrainer"

	"github.com/spf13/cobra"
)

var cheatCmd = &cobra.Command{
	Use:   "cheat",
	Short: "Display a quick reference cheat-sheet",
	Long: `Display a quick reference guide for command line tools.

This shows all available commands with their descriptions for quick lookup.

Example:
  terminal-trainer cheat --topic linux
  terminal-trainer cheat --topic git --compact`,
	Run: runCheat,
}

func init() {
	rootCmd.AddCommand(cheatCmd)
	cheatCmd.Flags().BoolP("compact", "c", false, "Use compact display format")
}

func runCheat(cmd *cobra.Command, args []string) {
	// Get flags
	topic, _ := cmd.Flags().GetString("topic")
	compact, _ := cmd.Flags().GetBool("compact")
	verbose, _ := cmd.Flags().GetBool("verbose")

	if verbose {
		fmt.Printf("🔧 Topic: %s\n", topic)
		fmt.Printf("📋 Compact mode: %t\n", compact)
		fmt.Println()
	}

	// Create the question engine
	engine := terminaltrainer.NewQuestionEngine()

	// Load knowledge base
	knowledgeFile := filepath.Join("/workspaces/terminal-trainer/knowledge", topic+".json")
	if err := loadKnowledgeBase(engine, knowledgeFile); err != nil {
		fmt.Printf("❌ Error loading %s knowledge: %v\n", topic, err)
		fmt.Printf("💡 Make sure %s exists\n", knowledgeFile)
		os.Exit(1)
	}

	// Get questions (we'll use them as reference items)
	questions := engine.GetAllQuestions()
	if len(questions) == 0 {
		fmt.Printf("❌ No reference items found for topic: %s\n", topic)
		os.Exit(1)
	}

	// Display cheat-sheet
	fmt.Printf("📚 %s Command Reference\n", strings.ToUpper(topic))
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println()

	for _, question := range questions {
		if compact {
			fmt.Printf("%-10s - %s\n", question.ExpectedAnswer, question.Description)
		} else {
			fmt.Printf("🔧 %s\n", question.ExpectedAnswer)
			fmt.Printf("   %s\n", question.Description)
			if verbose {
				fmt.Printf("   Category: %s\n", question.Category)
			}
			fmt.Println()
		}
	}

	fmt.Printf("💡 Total commands: %d\n", len(questions))
	fmt.Printf("🎯 Try 'terminal-trainer quiz -t %s' to test your knowledge!\n", topic)
}
