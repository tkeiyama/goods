package file

import (
	"fmt"
	"os"
)

// ReadNames returns the names of files and folders in the given directory.
func GetNames(dir string) ([]string, error) {
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
