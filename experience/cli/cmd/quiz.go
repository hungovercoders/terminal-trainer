package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"terminal-trainer/behaviour/terminaltrainer"

	"github.com/spf13/cobra"
)

// quizCmd represents the quiz command
var quizCmd = &cobra.Command{
	Use:   "quiz",
	Short: "Start an interactive quiz session",
	Long: `Start an interactive question and answer session to test your knowledge.

You'll be asked questions about command line tools and need to provide the correct answer.
Type 'quit' at any time to exit the quiz.

Example:
  terminal-trainer quiz --topic linux
  terminal-trainer quiz --topic git --count 5`,
	Run: runQuiz,
}

func init() {
	rootCmd.AddCommand(quizCmd)

	// Quiz-specific flags
	quizCmd.Flags().IntP("count", "c", 0, "Number of questions to ask (0 = all)")
	quizCmd.Flags().BoolP("shuffle", "s", false, "Randomize question order")
}

func runQuiz(cmd *cobra.Command, args []string) {
	// Get flags from root command (topic, verbose)
	topic, _ := cmd.Flags().GetString("topic")
	verbose, _ := cmd.Flags().GetBool("verbose")

	// Get quiz-specific flags
	count, _ := cmd.Flags().GetInt("count")
	shuffle, _ := cmd.Flags().GetBool("shuffle")

	if verbose {
		fmt.Printf("🔧 Starting quiz for topic: %s\n", topic)
		fmt.Printf("🔢 Question limit: %d (0=all)\n", count)
		fmt.Printf("🔀 Shuffle questions: %t\n", shuffle)
		fmt.Println()
	}

	// Create question engine (using your behaviour logic)
	engine := terminaltrainer.NewQuestionEngine()

	// Load knowledge base
	knowledgeFile := filepath.Join("/workspaces/terminal-trainer/knowledge", topic+".json")
	if err := loadKnowledgeBase(engine, knowledgeFile); err != nil {
		fmt.Printf("❌ Error loading %s knowledge: %v\n", topic, err)
		fmt.Printf("💡 Make sure %s exists\n", knowledgeFile)
		os.Exit(1)
	}

	// Get questions using your engine
	questions := engine.GetAllQuestions()
	if len(questions) == 0 {
		fmt.Printf("❌ No questions found for topic: %s\n", topic)
		os.Exit(1)
	}

	// Limit questions if count specified
	if count > 0 && count < len(questions) {
		questions = questions[:count]
	}

	fmt.Printf("🚀 Starting %s quiz with %d questions!\n", topic, len(questions))
	fmt.Println("💡 Type 'quit' to exit anytime")
	fmt.Println()

	// Run interactive quiz
	runInteractiveQuiz(questions, verbose)
}

// loadKnowledgeBase loads JSON knowledge into the engine
func loadKnowledgeBase(engine *terminaltrainer.QuestionEngine, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", filePath, err)
	}
	defer file.Close()

	return engine.LoadFromJSON(file)
}

// runInteractiveQuiz handles the Q&A interaction
func runInteractiveQuiz(questions []terminaltrainer.Question, verbose bool) {
	scanner := bufio.NewScanner(os.Stdin)
	correctAnswers := 0
	totalQuestions := len(questions)

	for i, question := range questions {
		fmt.Printf("Question %d/%d: %s\n", i+1, totalQuestions, question.Text)
		fmt.Print("Your answer: ")

		// Get user input
		if !scanner.Scan() {
			break
		}
		userAnswer := strings.TrimSpace(scanner.Text())

		// Handle quit
		if strings.ToLower(userAnswer) == "quit" {
			fmt.Println("👋 Thanks for practicing!")
			break
		}

		// Validate answer using your behaviour logic
		if terminaltrainer.ValidateAnswer(question, userAnswer) {
			correctAnswers++
			fmt.Printf("✅ Correct! '%s' - %s\n", question.ExpectedAnswer, question.Description)
			if verbose {
				fmt.Printf("🏷️  Category: %s\n", question.Category)
			}
		} else {
			fmt.Printf("❌ Incorrect. The answer is '%s' - %s\n",
				question.ExpectedAnswer, question.Description)
			if verbose {
				fmt.Printf("🏷️  Category: %s\n", question.Category)
			}
		}
		fmt.Println()
	}

	// Show final score
	percentage := float64(correctAnswers) / float64(totalQuestions) * 100
	fmt.Printf("🏆 Quiz Complete! Score: %d/%d (%.1f%%)\n",
		correctAnswers, totalQuestions, percentage)

	if percentage >= 90 {
		fmt.Println("🌟 Excellent! You're a command line expert!")
	} else if percentage >= 70 {
		fmt.Println("👍 Good work! Keep practicing!")
	} else {
		fmt.Println("📚 Keep studying - practice makes perfect!")
	}
}
