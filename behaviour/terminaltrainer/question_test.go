package terminaltrainer

import (
    "encoding/json"
    "testing"
)

func TestCommand(t *testing.T) {
    // Create a Command instance
    cmd := Command{
        Command:     "ls",
        Description: "List files",
        Question:    "What lists files?",
    }

    // Test that our fields work
    if cmd.Command != "ls" {
        t.Errorf("Expected 'ls', got '%s'", cmd.Command)
    }

    t.Logf("Command: %s", cmd.Command)
    t.Logf("Description: %s", cmd.Description)
    t.Logf("Question: %s", cmd.Question)
}

func TestCommandCategory(t *testing.T) {
    jsonData := `{
        "category": "file_management",
        "commands": [
            {
                "command": "ls",
                "description": "List files in the current directory",
                "question": "What command lists files in the current directory?"
            },
            {
                "command": "cd",
                "description": "Change the current directory",
                "question": "What command changes the current directory?"
            }
        ]
    }`

    var category CommandCategory
    err := json.Unmarshal([]byte(jsonData), &category)
    if err != nil {
        t.Fatalf("Failed to parse JSON: %v", err)
    }

    // Test the category
    if category.Category != "file_management" {
        t.Errorf("Expected 'file_management', got '%s'", category.Category)
    }

    // Test the commands
    if len(category.Commands) != 2 {
        t.Errorf("Expected 2 commands, got %d", len(category.Commands))
    }

    // Test the first command
    cmd1 := category.Commands[0]
    if cmd1.Command != "ls" {
        t.Errorf("Expected 'ls', got '%s'", cmd1.Command)
    }
    if cmd1.Description != "List files in the current directory" {
        t.Errorf("Expected 'List files in the current directory', got '%s'", cmd1.Description)
    }
    if cmd1.Question != "What command lists files in the current directory?" {
        t.Errorf("Expected 'What command lists files in the current directory?', got '%s'", cmd1.Question)
    }

    // Test the second command
    cmd2 := category.Commands[1]
    if cmd2.Command != "cd" {
        t.Errorf("Expected 'cd', got '%s'", cmd2.Command)
    }
    if cmd2.Description != "Change the current directory" {
        t.Errorf("Expected 'Change the current directory', got '%s'", cmd2.Description)
    }
    if cmd2.Question != "What command changes the current directory?" {
        t.Errorf("Expected 'What command changes the current directory?', got '%s'", cmd2.Question)
    }
}