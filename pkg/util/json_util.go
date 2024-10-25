package util

import (
	"encoding/json"
	"fmt"
	"os"
)

// CreateTempJSONFile creates a temporary JSON file with the provided data
func CreateTempJSONFile(data map[string]string) (string, error) {

	// Convert sample data to JSON format
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// Create a temporary file
	tempFile, err := os.CreateTemp("", "sample_*.json")
	if err != nil {
		return "", err
	}
	defer func(tempFile *os.File) {
		err := tempFile.Close()
		if err != nil {
			_ = fmt.Errorf("failed to close file: %v", err)
		}
	}(tempFile)

	// Write JSON data to the file
	if _, err = tempFile.Write(jsonData); err != nil {
		return "", err
	}

	return tempFile.Name(), nil
}

// VerifyFileExists checks if a file exists at the given path
func VerifyFileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err == nil, err
}
