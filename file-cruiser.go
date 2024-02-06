package main // Declares this as the main package (executable program)

import (
	"fmt"  // Provides formatted I/O functions
	"os"   // Provides operating system functions, like file operations
	"path" // Provides path manipulation functions
	"sort" // Provides sorting functions
)

type Counter struct {
	dirs  int // Counts directories
	files int // Counts files
}

// Increments directory or file count based on path type
func (counter *Counter) index(path string) {
	stat, _ := os.Stat(path)
	if stat.IsDir() {
		counter.dirs += 1
	} else {
		counter.files += 1
	}
}

// Formats and returns a string with directory and file counts
func (counter *Counter) output() string {
	return fmt.Sprintf("\n%d directories, %d files", counter.dirs, counter.files)
}

// Returns a sorted list of subdirectory names within a base directory
func dirnamesFrom(base string) []string {
	file, err := os.Open(base) // Open the base directory
	if err != nil {
		fmt.Println(err) // Handle errors more gracefully
	}
	names, _ := file.Readdirnames(0) // Read subdirectory names
	file.Close()

	sort.Strings(names) // Sort names alphabetically
	return names
}

// Todo: skip executable file, and directories that starts with .
func jsonBuilder(base string, prefix string) string {
	names := dirnamesFrom(base)

	// Determine the type (file or directory)
	fileInfo, err := os.Stat(base)
	if err != nil {
		fmt.Println(err) // Handle errors appropriately
		return ""
	}
	nodeType := "file"
	if fileInfo.IsDir() {
		nodeType = "directory"
	}

	json := "{\n" + prefix + "\"name\": \"" + path.Base(base) + "\",\n" + prefix + "\"type\": \"" + nodeType + "\",\n" + prefix + "\"children\": ["

	for index, name := range names {
		if name[0] == '.' {
			continue
		}
		subpath := path.Join(base, name)

		if index == len(names)-1 {
			json += "\n" + prefix + "    " + jsonBuilder(subpath, prefix+"    ")
		} else {
			json += "\n" + prefix + "    " + jsonBuilder(subpath, prefix+"    ") + ","
		}
	}

	return json + "\n" + prefix + "]}\n"
}

func writeJSONToFile(json string) error {
	filename := "self.json"
	file, err := os.Create(filename)
	if err != nil {
		return err // Handle errors appropriately
	}
	defer file.Close()

	_, err = file.WriteString(json)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var directory string
	if len(os.Args) > 1 {
		directory = os.Args[1] // Use provided directory argument
	} else {
		directory = "." // Default to current directory
	}

	json := jsonBuilder(directory, "")

	err := writeJSONToFile(json)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("JSON written to self.json")
	}
}
