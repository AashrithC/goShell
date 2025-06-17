# goShell

An enhanced Unix shell wrapper built with **Go** featuring an interactive UI and intelligent autocomplete.

## Features

- **Interactive UI**: Clean, colorized interface powered by go-prompt
- **Smart Autocomplete**: zsh-like tab completion for commands and files
- **Path Navigation**: Support for absolute paths, relative paths, and `~` expansion
- **File Type Detection**: Rich descriptions for different file types
- **Hidden Files**: Show hidden files when typing `.`
- **Built-in Commands**: Enhanced `cd`, `pwd`, `clear`, and `exit`

## Installation

```bash
git clone https://github.com/Aashrithc/goShell
cd goShell
go run main.go
```

## Usage

- Use **Tab** for autocomplete suggestions
- Arrow keys for command history navigation
- Type `exit` to quit the shell

### Autocomplete Examples

- `ls R<TAB>` → completes to `README.md`
- `cd ~/D<TAB>` → completes to `~/Documents/`
- `cat src/<TAB>` → shows files in `src/` directory
- `.g<TAB>` → shows hidden files starting with `.g`

## Built-in Commands

- `cd <path>` - Change directory (supports `~` and relative paths)
- `pwd` - Print working directory
- `clear` - Clear screen
- `exit` - Exit shell

## Dependencies

- [go-prompt](https://github.com/c-bata/go-prompt) - Interactive prompt library
