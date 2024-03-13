package main
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)
// FileNode represents a file or directory in the file system.
type FileNode struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Content  string `json:"content,omitempty"`
	Children []*FileNode `json:"children,omitempty"`
}
// traverseDirectory recursively traverses the given directory and builds the file tree structure.
func traverseDirectory(dirPath string) (*FileNode, error) {
	fileInfo, err := os.Stat(dirPath)
	if err != nil {
		return nil, err
	}
	node := &FileNode{
		Name: filepath.Base(dirPath),
		Type: "file",
	}
	if fileInfo.IsDir() {
		node.Type = "directory"
		entries, err := os.ReadDir(dirPath)
		if err != nil {
			return nil, err
		}
		sort.Slice(entries, func(i, j int) bool {
			return entries[i].Name() < entries[j].Name()
		})
		for _, entry := range entries {
			if shouldSkipEntry(entry.Name()) {
				continue
			}
			entryPath := filepath.Join(dirPath, entry.Name())
			child, err := traverseDirectory(entryPath)
			if err != nil {
				return nil, err
			}
			node.Children = append(node.Children, child)
		}
	} else {
		// If it's a file, read the contents and store them in the FileNode
		content, err := ioutil.ReadFile(dirPath)
		if err != nil {
			return nil, err
		}
		// Skip the content if the file is named "file-cruiser"
		// which is an executable file
		if filepath.Base(dirPath) != "file-cruiser" {
			node.Content = string(content)
		}
	}
	return node, nil
}
// shouldSkipEntry returns true if the given entry name should be skipped
func shouldSkipEntry(name string) bool {
	return strings.HasPrefix(name, ".")
}
// writeJSONToFile writes the given file tree as JSON to the specified file path.
func writeJSONToFile(filePath string, root *FileNode) error {
	jsonData, err := json.MarshalIndent(root, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, jsonData, 0644)
}
func main() {
	dirPath := "."
	if len(os.Args) > 1 {
		dirPath = os.Args[1]
	}
	root, err := traverseDirectory(dirPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = writeJSONToFile("file-tree.json", root)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}