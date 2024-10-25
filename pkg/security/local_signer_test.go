package security

import (
	"attestations/pkg/util"
	"os"
	"testing"
)

func TestCreateAndVerifyTempJSONFile(t *testing.T) {

	// Define a sample JSON structure
	sampleData := map[string]string{
		"key": "value",
	}

	// Call the function to create the temporary JSON file
	filePath, err := util.CreateTempJSONFile(sampleData)
	if err != nil {
		t.Errorf("Failed to create temporary file: %v", err)
		return
	}

	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Errorf("Failed to remove temporary file: %v", err)
		}
	}(filePath) // Clean up the file afterward

	// Verify the file exists
	exists, err := util.VerifyFileExists(filePath)
	if err != nil {
		t.Errorf("Failed to verify file existence: %v", err)
		return
	}
	if !exists {
		t.Errorf("File does not exist at: %s", filePath)
		return
	}
}
