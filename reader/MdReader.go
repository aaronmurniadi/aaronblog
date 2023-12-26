package reader

import (
	"fmt"
	"github.com/chasefleming/elem-go"
	"os"
	"path/filepath"
)

func GetPosts() *elem.Element {
	directoryPath := "./contents/posts"
	files, err := ListFiles(directoryPath)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	fmt.Println("Found posts in", directoryPath, ":")
	for _, file := range files {
		fmt.Println(file)
	}
	return nil
}

func ListFiles(dirPath string) ([]string, error) {
	var files []string

	// Read directory contents
	dirEntries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	// Iterate through directory entries
	for _, entry := range dirEntries {
		if entry.IsDir() {
			// Skip directories if needed
			continue
		}
		// Get full file path
		fullPath := filepath.Join(dirPath, entry.Name())
		files = append(files, fullPath)
	}

	return files, nil
}
