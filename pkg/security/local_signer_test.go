package security

import (
	"attestations/pkg/util"
	_ "embed"
	"testing"
)

//go:embed /keys/keypair-1/private-key.pem
var privateKey []byte

//go:embed /keys/keypair-1/public-key.pem
var publicKey []byte

func TestLocalSigner_Sign(t *testing.T) {

	//Arrange
	sampleData := map[string]string{
		"key": "value",
	}

	err, filePath := util.CreateJsonTempFile(sampleData, "LocalSigner_Sign*.json")
	if err != nil {
		t.Errorf("Failed to create temporary file: %v", err)
		return
	}
	defer func(name string) { util.RemoveFile(name) }(filePath)

	signer := LocalSigner{}
	err = signer.LoadKeyPairBytes(privateKey, publicKey)
	if err != nil {
		t.Errorf("Failed to Load Key Pair: %v", err)
		return
	}
	//Act

	//Assert

}

func TestCreateAndVerifyTempJSONFile(t *testing.T) {

	sampleData := map[string]string{
		"key": "value",
	}

	err, filePath := util.CreateJsonTempFile(sampleData)

	defer func(name string) { util.RemoveFile(name) }(filePath)

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
