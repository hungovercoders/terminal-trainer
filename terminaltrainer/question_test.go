package terminaltrainer

import (
    "encoding/json"
    "os"
    "strings"
    "testing"
)

// ... keep all existing tests ...

func TestBasicStructs(t *testing.T) {
    cmd := Command{
        Command:     "ls",
        Description: "List files in current directory",
        Question:    "What command lists files in the current directory?",
    }

    if cmd.Command != "ls" {
        t.Errorf("Expected 'ls', got '%s'", cmd.Command)
    }

    t.Logf("Created command: %s", cmd.Command)

    category := CommandCategory{
        Category: "file_management",
        Commands: []Command{cmd},
    }

    if category.Category != "file_management" {
        t.Errorf("Expected 'file_management', got '%s'", category.Category)
    }

    if len(category.Commands) != 1 {
        t.Errorf("Expected 1 command, got %d", len(category.Commands))
    }

    t.Logf("Category: %s with %d commands", category.Category, len(category.Commands))
}

func TestJSONLoading(t *testing.T) {
    jsonString := `{
        "category": "file_management",
        "commands": [
            {
                "command": "ls",
                "description": "List files in current directory",
                "question": "What command lists files in the current directory?"
            }
        ]
    }`

    var category CommandCategory
    err := json.Unmarshal([]byte(jsonString), &category)
    
    if err != nil {
        t.Fatalf("Failed to parse JSON: %v", err)
    }

    if category.Category != "file_management" {
        t.Errorf("Expected 'file_management', got '%s'", category.Category)
    }

    firstCommand := category.Commands[0]
    if firstCommand.Command != "ls" {
        t.Errorf("Expected 'ls', got '%s'", firstCommand.Command)
    }

    t.Logf("Loaded from JSON - Category: %s", category.Category)
    t.Logf("First command: %s - %s", firstCommand.Command, firstCommand.Description)
}

func TestQuestionLogic(t *testing.T) {
    cmd := Command{
        Command:     "ls",
        Description: "List files in current directory",
        Question:    "What command lists files in the current directory?",
    }

    question := ConvertToQuestion(cmd, "file_management")

    if question.ExpectedAnswer != "ls" {
        t.Errorf("Expected answer 'ls', got '%s'", question.ExpectedAnswer)
    }

    if !ValidateAnswer(question, "ls") {
        t.Error("Expected 'ls' to be a correct answer")
    }

    if ValidateAnswer(question, "cd") {
        t.Error("Expected 'cd' to be incorrect")
    }

    if !ValidateAnswer(question, "LS") {
        t.Error("Expected 'LS' to be correct (case insensitive)")
    }

    t.Logf("Question: %s", question.Text)
    t.Logf("Expected answer: %s", question.ExpectedAnswer)
}

func TestQuestionEngine(t *testing.T) {
    // Create a question engine
    engine := NewQuestionEngine()

    // JSON that matches your linux.json structure
    jsonData := `{
        "category": "file_management",
        "commands": [
            {
                "command": "ls",
                "description": "List files in current directory",
                "question": "What command lists files in the current directory?"
            },
            {
                "command": "cd",
                "description": "Change directory",
                "question": "What command changes the current directory?"
            }
        ]
    }`

    // Load the JSON into the engine
    reader := strings.NewReader(jsonData)
    err := engine.LoadFromJSON(reader)
    if err != nil {
        t.Fatalf("Failed to load JSON: %v", err)
    }

    // Test that we can get all questions
    questions := engine.GetAllQuestions()
    if len(questions) != 2 {
        t.Errorf("Expected 2 questions, got %d", len(questions))
    }

    // Test question count
    if engine.GetQuestionCount() != 2 {
        t.Errorf("Expected question count 2, got %d", engine.GetQuestionCount())
    }

    // Test first question
    firstQ := questions[0]
    if firstQ.ExpectedAnswer != "ls" {
        t.Errorf("Expected first answer 'ls', got '%s'", firstQ.ExpectedAnswer)
    }

    if firstQ.Category != "file_management" {
        t.Errorf("Expected category 'file_management', got '%s'", firstQ.Category)
    }

    // Test answering questions
    if !ValidateAnswer(firstQ, "ls") {
        t.Error("Expected 'ls' to be correct")
    }

    t.Logf("Engine loaded %d questions from JSON!", engine.GetQuestionCount())
    t.Logf("First question: %s", firstQ.Text)
    t.Logf("Expected answer: %s", firstQ.ExpectedAnswer)
}

// NEW TEST: Load your actual linux.json file
func TestLoadLinuxJSON(t *testing.T) {
    // Create a question engine
    engine := NewQuestionEngine()

    // Open your actual linux.json file
    file, err := os.Open("/workspaces/terminal-trainer/knowledge/linux.json")
    if err != nil {
        t.Fatalf("Failed to open linux.json: %v", err)
    }
    defer file.Close() // Always close files when done

    // Load the file into the engine
    err = engine.LoadFromJSON(file)
    if err != nil {
        t.Fatalf("Failed to load linux.json: %v", err)
    }

    // Test that all questions from your file are loaded
    questions := engine.GetAllQuestions()
    if len(questions) == 0 {
        t.Error("Expected some questions from linux.json, got none")
    }

    t.Logf("Loaded %d questions from your linux.json file!", len(questions))

    // Show all the questions from your file
    for i, question := range questions {
        t.Logf("Question %d: %s", i+1, question.Text)
        t.Logf("  Expected answer: %s", question.ExpectedAnswer)
        t.Logf("  Description: %s", question.Description)
        t.Logf("  Category: %s", question.Category)
        t.Logf("")
    }

    // Test that we can answer one of the questions
    if len(questions) > 0 {
        firstQuestion := questions[0]
        if ValidateAnswer(firstQuestion, firstQuestion.ExpectedAnswer) {
            t.Logf("✅ Answer validation works for real questions!")
        } else {
            t.Error("❌ Answer validation failed for real question")
        }
    }
}