package terminaltrainer

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

