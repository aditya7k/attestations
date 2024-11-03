package util

import (
	"encoding/json"
	"fmt"
)

func CreateJsonTempFile(sampleData map[string]string, pattern string) (error, string) {

	jsonData, err := json.Marshal(sampleData)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %v, %w", sampleData, err), ""
	}

	filePath, err := CreateTempFile(jsonData, pattern)
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %v, %w", jsonData, err), ""
	}

	return nil, filePath
}
