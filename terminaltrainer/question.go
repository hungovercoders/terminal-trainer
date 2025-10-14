package terminaltrainer

import (
    "encoding/json"
    "io"
    "strings"
)

// Command represents a single command from our JSON file
type Command struct {
    Command     string `json:"command"`
    Description string `json:"description"`
    Question    string `json:"question"`
}

// CommandCategory represents a category of commands (like "file_management")
type CommandCategory struct {
    Category string    `json:"category"`
    Commands []Command `json:"commands"`
}

// Question represents a question ready to be asked to a user
type Question struct {
    Text           string // The question text
    ExpectedAnswer string // What answer we expect
    Description    string // Description of the command
    Category       string // Which category this belongs to
}

// QuestionEngine manages loading knowledge and creating questions
type QuestionEngine struct {
    categories []CommandCategory // Store all loaded categories
}

// NewQuestionEngine creates a new question engine
func NewQuestionEngine() *QuestionEngine {
    return &QuestionEngine{
        categories: make([]CommandCategory, 0), // Start with empty list
    }
}

// LoadFromJSON loads a command category from a JSON reader
func (qe *QuestionEngine) LoadFromJSON(reader io.Reader) error {
    var category CommandCategory
    decoder := json.NewDecoder(reader)
    
    if err := decoder.Decode(&category); err != nil {
        return err
    }
    
    // Add this category to our collection
    qe.categories = append(qe.categories, category)
    return nil
}

// GetAllQuestions converts all loaded commands into questions
func (qe *QuestionEngine) GetAllQuestions() []Question {
    var questions []Question
    
    // Loop through each category
    for _, category := range qe.categories {
        // Loop through each command in the category
        for _, command := range category.Commands {
            question := ConvertToQuestion(command, category.Category)
            questions = append(questions, question)
        }
    }
    
    return questions
}

// GetQuestionCount returns how many questions are available
func (qe *QuestionEngine) GetQuestionCount() int {
    return len(qe.GetAllQuestions())
}

// ConvertToQuestion converts a Command into a Question
func ConvertToQuestion(cmd Command, category string) Question {
    return Question{
        Text:           cmd.Question,
        ExpectedAnswer: cmd.Command,
        Description:    cmd.Description,
        Category:       category,
    }
}

// ValidateAnswer checks if the user's answer matches the expected answer
func ValidateAnswer(question Question, userAnswer string) bool {
    // Make both answers lowercase and remove extra spaces
    expected := strings.TrimSpace(strings.ToLower(question.ExpectedAnswer))
    user := strings.TrimSpace(strings.ToLower(userAnswer))
    
    return expected == user
}

