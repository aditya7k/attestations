package util

import (
	"fmt"
	"os"
)

// CreateTempFile creates a temporary file with the given data and pattern
// The pattern should contain '*' to be replaced by a random string
func CreateTempFile(data []byte, pattern string) (string, error) {

	tempFile, err := os.CreateTemp("", pattern)
	if err != nil {
		return "", err
	}

	defer func(tempFile *os.File) { closeFile(tempFile) }(tempFile)

	if _, err = tempFile.Write(data); err != nil {
		return "", err
	}

	return tempFile.Name(), nil
}

func VerifyFileExists(filePath string) (bool, error) {

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err == nil, err
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		_ = fmt.Errorf("failed to close file: %s, %w", file.Name(), err)
	}
}

func RemoveFile(name string) {

	err := os.Remove(name)
	if err != nil {
		_ = fmt.Errorf("failed to remove temporary file: %s, %w", name, err)
	}
}
