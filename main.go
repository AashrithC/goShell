package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)


func main() {
    scanner := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("> ")
        line, err := scanner.ReadString('\n')
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }

        if err = execInput(line); err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
    }
}

// ErrNoPath is returned when cd command is called without a path argument
var ErrNoPath = errors.New("path required")

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
        return os.Chdir(tokens[1])
    case "exit":
        os.Exit(0)
    }

    // Execute external commands
    proc := exec.Command(tokens[0], tokens[1:]...)
    proc.Stderr = os.Stderr
    proc.Stdout = os.Stdout
    return proc.Run()
}