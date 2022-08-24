package file

import (
	"fmt"
	"os"
)

// GetNamesInDir returns the names of files or folders in the given directory.
func GetNamesInDir(dir string) ([]string, error) {
	var fileNames []string

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("The given path was NOT FOUND: %s", dir)
		return nil, err
	}

	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}

	return fileNames, nil
}
