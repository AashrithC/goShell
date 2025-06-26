package main

//hi
import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/c-bata/go-prompt"
)

// ErrNoPath is returned when cd command is called without a path argument
var ErrNoPath = errors.New("path required")

// commonCommands contains frequently used Unix commands for autocomplete
var commonCommands = []string{
	"ls", "cd", "pwd", "mkdir", "rmdir", "rm", "cp", "mv", "cat", "less", "more",
	"grep", "find", "which", "man", "ps", "kill", "top", "df", "du", "free",
	"git", "vim", "nano", "curl", "wget", "ssh", "scp", "tar", "gzip", "unzip",
	"chmod", "chown", "sudo", "history", "clear", "exit",
}

func main() {
	fmt.Println("goShell - Enhanced Unix Shell Wrapper")
	fmt.Println("Type 'exit' to quit, use Tab for autocomplete")
	fmt.Println("----------------------------------------")
	
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("> "),
		prompt.OptionTitle("goShell"),
		prompt.OptionHistory([]string{}),
		prompt.OptionPrefixTextColor(prompt.Blue),
		prompt.OptionPreviewSuggestionTextColor(prompt.Blue),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionSuggestionBGColor(prompt.DarkGray),
	)
	p.Run()
}

// executor handles command execution
func executor(input string) {
	input = strings.TrimSpace(input)
	if input == "" {
		return
	}
	
	if err := execInput(input); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

// completer provides autocomplete suggestions
func completer(d prompt.Document) []prompt.Suggest {
	text := d.TextBeforeCursor()
	words := strings.Fields(text)
	
	// If we're typing the first word, suggest commands
	if len(words) <= 1 {
		word := ""
		if len(words) == 1 {
			word = words[0]
		}
		return getCommandSuggestions(word)
	}
	
	// For subsequent words, suggest files/directories
	currentWord := ""
	if len(text) > 0 && text[len(text)-1] != ' ' {
		// We're in the middle of typing a word
		lastSpace := strings.LastIndex(text, " ")
		if lastSpace >= 0 {
			currentWord = text[lastSpace+1:]
		}
	}
	
	return getFileSuggestions(currentWord)
}

// getCommandSuggestions returns command suggestions
func getCommandSuggestions(word string) []prompt.Suggest {
	suggestions := []prompt.Suggest{}
	
	for _, cmd := range commonCommands {
		if strings.HasPrefix(cmd, word) {
			suggestions = append(suggestions, prompt.Suggest{
				Text:        cmd,
				Description: "Command",
			})
		}
	}
	
	return suggestions
}

// getFileSuggestions returns file and directory suggestions with zsh-like behavior
func getFileSuggestions(currentWord string) []prompt.Suggest {
	suggestions := []prompt.Suggest{}
	
	// Determine the directory to search and the prefix to match
	searchDir := "."
	prefix := currentWord
	
	// Handle paths with directories (e.g., "src/ma" -> search in "src/", match "ma*")
	if strings.Contains(currentWord, "/") {
		lastSlash := strings.LastIndex(currentWord, "/")
		searchDir = currentWord[:lastSlash+1]
		prefix = currentWord[lastSlash+1:]
		
		// Handle absolute paths
		if strings.HasPrefix(searchDir, "/") {
			// Absolute path
		} else if strings.HasPrefix(searchDir, "~/") {
			// Home directory expansion
			home, err := os.UserHomeDir()
			if err == nil {
				searchDir = filepath.Join(home, searchDir[2:])
			}
		} else {
			// Relative path
			pwd, err := os.Getwd()
			if err == nil {
				searchDir = filepath.Join(pwd, searchDir)
			}
		}
	}
	
	// Clean the search directory path
	searchDir = filepath.Clean(searchDir)
	
	// Read directory contents
	files, err := os.ReadDir(searchDir)
	if err != nil {
		return suggestions
	}
	
	for _, file := range files {
		name := file.Name()
		
		// Skip hidden files unless prefix starts with "."
		if strings.HasPrefix(name, ".") && !strings.HasPrefix(prefix, ".") {
			continue
		}
		
		// Check if the file matches the prefix
		if strings.HasPrefix(name, prefix) {
			var suggestion prompt.Suggest
			
			if file.IsDir() {
				// For directories, construct the full path and add trailing slash
				fullPath := constructFullPath(currentWord, name, true)
				suggestion = prompt.Suggest{
					Text:        fullPath,
					Description: "Directory",
				}
			} else {
				// For files, construct the full path
				fullPath := constructFullPath(currentWord, name, false)
				
				// Get file info for better descriptions
				desc := getFileDescription(filepath.Join(searchDir, name))
				suggestion = prompt.Suggest{
					Text:        fullPath,
					Description: desc,
				}
			}
			
			suggestions = append(suggestions, suggestion)
		}
	}
	
	return suggestions
}

// constructFullPath builds the complete path for autocomplete suggestions
func constructFullPath(currentWord, fileName string, isDir bool) string {
	if strings.Contains(currentWord, "/") {
		// Replace the last part after the last slash
		lastSlash := strings.LastIndex(currentWord, "/")
		basePath := currentWord[:lastSlash+1]
		result := basePath + fileName
		if isDir {
			result += "/"
		}
		return result
	}
	
	// Simple case: just the filename
	if isDir {
		return fileName + "/"
	}
	return fileName
}

// getFileDescription provides descriptive information about files
func getFileDescription(filePath string) string {
	info, err := os.Stat(filePath)
	if err != nil {
		return "File"
	}
	
	// Check if it's executable
	if info.Mode()&0111 != 0 {
		return "Executable"
	}
	
	// Check file extension for common types
	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".go":
		return "Go source"
	case ".js", ".ts":
		return "JavaScript/TypeScript"
	case ".py":
		return "Python script"
	case ".md":
		return "Markdown"
	case ".txt":
		return "Text file"
	case ".json":
		return "JSON file"
	case ".yaml", ".yml":
		return "YAML file"
	case ".xml":
		return "XML file"
	case ".html", ".htm":
		return "HTML file"
	case ".css":
		return "CSS file"
	case ".sh":
		return "Shell script"
	case ".zip", ".tar", ".gz":
		return "Archive"
	case ".jpg", ".jpeg", ".png", ".gif":
		return "Image"
	default:
		return "File"
	}
}

// execInput parses and executes a command line input
func execInput(line string) error {
	line = strings.TrimSpace(line)
	tokens := strings.Fields(line)

	// Handle empty input
	if len(tokens) == 0 {
		return nil
	}

	// Handle built-in commands
	switch tokens[0] {
	case "cd":
		if len(tokens) < 2 {
			return ErrNoPath
		}
		return changeDirectory(tokens[1])
	case "pwd":
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}
		fmt.Println(pwd)
		return nil
	case "clear":
		fmt.Print("\033[2J\033[H")
		return nil
	case "exit":
		fmt.Println("Goodbye!")
		os.Exit(0)
	}

	// Execute external commands
	proc := exec.Command(tokens[0], tokens[1:]...)
	proc.Stderr = os.Stderr
	proc.Stdout = os.Stdout
	proc.Stdin = os.Stdin
	return proc.Run()
}

// changeDirectory handles directory changes with path expansion
func changeDirectory(path string) error {
	// Handle ~ expansion
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		path = filepath.Join(home, path[1:])
	}
	
	// Handle relative paths
	if !filepath.IsAbs(path) {
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}
		path = filepath.Join(pwd, path)
	}
	
	return os.Chdir(path)
}