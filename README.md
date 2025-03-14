# 🔍 Grepher

A simple and efficient implementation of the Unix `grep` command-line utility written in Go.

## 📝 Overview

Grepher is a lightweight tool that mimics the core functionality of the Unix grep utility. It allows you to search for patterns in files using regular expressions and provides several useful flags to customize your search.

This project demonstrates Go's powerful standard libraries and concise syntax while providing a useful tool for developers.

## ✨ Features

- Search for patterns using regular expressions
- Support for reading from files or standard input (stdin)
- Recursive directory searching
- Whole word matching
- Case-insensitive searching
- Line number display
- Match counting
- Invert match functionality

## 🚀 Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/grepher.git

# Navigate to the project directory
cd gogrep

# Build the executable
go build -o gogrep

# Optional: Move to a directory in your PATH
mv gogrep /usr/local/bin/
```

## 🛠️ Usage

Basic syntax:
```
./gogrep [OPTIONS] PATTERN [FILE...]
```

If no file is specified, GoGrep reads from standard input.

### Command-line Flags

| Flag | Description |
|------|-------------|
| `-i` | Perform case-insensitive matching |
| `-v` | Select non-matching lines (invert match) |
| `-n` | Prefix each line of output with its line number |
| `-c` | Only print a count of matching lines |
| `-r` | Recursively search subdirectories |
| `-w` | Match whole words only |

## 📚 How it Works

Grepher uses Go's `regexp` package to perform pattern matching on input files. It reads files line by line using bufio.Scanner, applies the specified matching criteria, and outputs the results.

The implementation includes:
- Command-line flag parsing with the `flag` package
- Regular expression compilation and matching
- File and directory handling
- Binary file detection (for recursive searches)
- Custom handling for special cases like stdin

## 🤝 Contributing

Contributions are welcome! Feel free to submit issues or pull requests.

1. Fork the repository
2. Create your feature branch: `git checkout -b feature-name`
3. Commit your changes: `git commit -m 'Add some feature'`
4. Push to the branch: `git push origin feature-name`
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.
