package security

import (
	"attestations/pkg/util"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

//go:embed keypair-1/private-key.pem
var privateKey []byte

//go:embed keypair-1/public-key.pem
var publicKey []byte

func TestLocalSigner_Sign(t *testing.T) {

	//Arrange
	sampleData := map[string]string{
		"key": "value",
	}

	sampleDataBytes, err := json.Marshal(sampleData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error marshaling statement: %v\n", err)
		return
	}

	signer := LocalSigner{}
	err = signer.LoadKeyPairBytes(privateKey, publicKey)
	if err != nil {
		t.Errorf("Failed to Load Key Pair: %v", err)
		return
	}

	//Act
	sign, err := signer.Sign(sampleDataBytes)
	if err != nil {
		t.Errorf("error signing message: %v\n", err)
	}
	assert.NotNil(t, sign)

	//Assert
	verified, err := signer.VerifySignature(sampleDataBytes, sign)
	if err != nil {
		t.Errorf("error verifying signature: %v\n", err)
	}

	assert.True(t, verified)
}

func TestCreateAndVerifyTempJSONFile(t *testing.T) {

	sampleData := map[string]string{
		"key": "value",
	}

	err, filePath := util.CreateJsonTempFile(sampleData, "CreateAndVerify*.json")

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
