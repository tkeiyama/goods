package goods

import (
	"fmt"
	"os"
)

// GetFileNames gets file names on a given path and returns them as a string array.
func GetFileNames(path string) ([]string, error) {
	var fileNames []string

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("The given path was NOT FOUND: %s", path)
		return nil, err
	}

	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}

	return fileNames, nil
}
