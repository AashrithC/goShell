# goShell

An enhanced Unix shell wrapper built with **Go** featuring an interactive UI and intelligent autocomplete system that provides a modern, user-friendly command-line experience.

## Features

- **Interactive UI**: Clean, colorized interface with customizable prompt styling
- **Smart Autocomplete**: Context-aware tab completion with 32+ common Unix commands
- **Advanced File Completion**: Intelligent file/directory suggestions with type detection
- **Path Navigation**: Full support for absolute paths, relative paths, and `~` expansion
- **File Type Recognition**: Rich descriptions for 15+ file types (Go, Python, JavaScript, etc.)
- **Hidden Files Support**: Automatically show hidden files when typing `.`
- **Built-in Commands**: Enhanced implementations of `cd`, `pwd`, `clear`, and `exit`
- **Command History**: Navigate previous commands with arrow keys
- **Error Handling**: Graceful error recovery with helpful messages

## Prerequisites

- **Go**: Version 1.19 or higher
- **Operating System**: Unix-like system (Linux, macOS, WSL on Windows)
- **Terminal**: Color-capable terminal for optimal experience

## Installation

### Quick Start
```bash
git clone https://github.com/AashrithC/goShell
cd goShell
go run main.go
```

### Build and Install
```bash
git clone https://github.com/AashrithC/goShell
cd goShell
go build -o goShell main.go
./goShell
```

### System-wide Installation
```bash
go build -o goShell main.go
sudo mv goShell /usr/local/bin/
goShell
```

## Usage

### Getting Started
When you first run goShell, you'll see a colorized prompt with the `>` indicator. The shell supports all standard Unix commands plus enhanced features.

### Key Features
- **Tab Completion**: Press Tab for intelligent autocomplete suggestions
- **Command History**: Use ↑/↓ arrow keys to navigate command history
- **Exit**: Type `exit` or press Ctrl+C to quit

### Autocomplete Examples

#### Command Completion
```bash
gi<TAB>          # Suggests: git, gzip
ls -<TAB>        # Shows available flags and options
```

#### Smart File Completion
```bash
ls R<TAB>        # Completes to README.md (shows "Markdown")
cat main.<TAB>   # Completes to main.go (shows "Go source")
./goS<TAB>       # Completes to ./goShell (shows "Executable")
```

#### Path Navigation
```bash
cd ~/D<TAB>      # Expands to ~/Documents/ (shows "Directory")
cat /etc/pas<TAB> # Completes to /etc/passwd (shows "File")
ls src/<TAB>     # Shows contents of src/ directory
```

#### Hidden Files
```bash
.g<TAB>          # Shows .git/, .gitignore, etc.
ls .<TAB>        # Lists all hidden files in current directory
```

## Built-in Commands

goShell provides enhanced versions of common shell built-ins:

- **`cd <path>`** - Change directory with smart path expansion
  - Supports `~` (home directory), `.` (current), `..` (parent)
  - Handles both absolute and relative paths
  - Example: `cd ~/Documents/projects`

- **`pwd`** - Print current working directory
  - Shows full absolute path
  - Example: `/home/user/Documents/projects`

- **`clear`** - Clear terminal screen
  - Equivalent to Ctrl+L in most terminals

- **`exit`** - Exit goShell gracefully
  - Shows goodbye message before terminating

## Advanced Features

### File Type Detection
goShell recognizes and provides descriptions for various file types:
- **Source Code**: `.go`, `.py`, `.js`, `.ts`, `.sh`
- **Documents**: `.md`, `.txt`, `.json`, `.yaml`, `.xml`
- **Web Files**: `.html`, `.css`
- **Archives**: `.zip`, `.tar`, `.gz`
- **Images**: `.jpg`, `.png`, `.gif`
- **Executables**: Files with execute permissions

### Command Database
Autocomplete includes 32 commonly used Unix commands:
`ls`, `cd`, `pwd`, `mkdir`, `rm`, `cp`, `mv`, `cat`, `grep`, `find`, `git`, `vim`, `curl`, `ssh`, and more.

## Troubleshooting

### Common Issues

**"command not found" errors**
- Ensure Go is properly installed and in your PATH
- Verify you're in the correct directory when running `go run main.go`

**Autocomplete not working**
- Check that your terminal supports ANSI color codes
- Ensure you have read permissions for directories you're navigating

**Build failures**
- Update to Go 1.19 or higher
- Run `go mod tidy` to ensure dependencies are properly downloaded

### Getting Help
- Check that all dependencies are installed: `go mod download`
- Verify Go version: `go version`
- For issues, please open an issue on the GitHub repository

## Dependencies

- **[go-prompt](https://github.com/c-bata/go-prompt)** v0.2.6 - Interactive prompt library
- **Standard Library**: Uses Go's `os`, `filepath`, `strings`, and `os/exec` packages

## Contributing

Contributions are welcome! Please feel free to submit issues, feature requests, or pull requests.

## License

This project is open source. Please check the repository for license details.

#hello
