# Cruising Linux File System with go

A simple Go program that recursively traverses a given directory and generates a JSON representation of the file system structure, including file names, types (file or directory), and contents (for non-binary files).

## What does it do?

- Recursive directory traversal to explore the entire file system hierarchy
- Differentiation between files and directories for proper handling
- Sorting of directory entries for organized output
- Skipping of hidden files and directories
- Option to exclude specific files
- JSON output

## Usage

```bash
1. Build the program:
go build file-cruiser.go
2. Run the program:
./file-cruise
```
