package os

import (
	"fmt"
	"os"
	"testing"
)

func isIncluded(t *testing.T, list []string, str string) bool {
	t.Helper()
	for _, s := range list {
		if s == str {
			return true
		}
	}
	return false
}

const FILE_NAME = "f.go"

func TestGetFileNames(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("I don't know what's happening.")
		return
	}
	fileNames, err := GetNamesInDir(pwd)
	if err != nil {
		fmt.Println("It seems the GetFileNames function doesn't work correctly")
		return
	}
	if !isIncluded(t, fileNames, FILE_NAME) {
		t.Errorf("%s is not included in the given list", FILE_NAME)
	}
}
