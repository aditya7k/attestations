package security

import (
	"attestations/pkg/statement/provenance"
	"attestations/pkg/util"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed keypair-1/private-key.pem
var privateKey []byte

//go:embed keypair-1/public-key.pem
var publicKey []byte

func TestLocalSigner_Sign(t *testing.T) {

	//Arrange
	attestationBytes, err := getAttestation()
	if err != nil {
		t.Error(err)
	}

	signer, err := getSigner()
	if err != nil {
		t.Error(err)
	}

	//Act
	signature, err := signer.Sign(attestationBytes)
	if err != nil {
		t.Errorf("error signing message: %v\n", err)
	}
	assert.NotNil(t, signature)

	//Assert
	verified, err := signer.VerifySignature(attestationBytes, signature)
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

func getSigner() (LocalSigner, error) {

	signer := LocalSigner{}
	err := signer.LoadKeyPairBytes(privateKey, publicKey)
	if err != nil {
		return signer, fmt.Errorf("failed to Load Key Pair: %v", err)
	}
	return signer, err
}

func getAttestation() ([]byte, error) {

	predicateStatementDTO := provenance.PredicateStatementDTO{
		ConfigURI:          "https://github.com/example/repo",
		ConfigDigest:       map[string]string{"sha1": "abc123"},
		ConfigEntryPoint:   "build script",
		PredicateBuilderId: "example.com/builder",
		PredicateBuildType: "https://example.com/build/type",
		SubjectName:        "example.com/my-artifact",
		SubjectDigest:      map[string]string{"sha256": "abcd1234"},
	}

	// Act
	statement := provenance.BuildProvenanceStatement(predicateStatementDTO)

	return json.Marshal(statement)
}

func getJsonBytes() ([]byte, error) {

	sampleData := map[string]string{
		"key": "value",
	}

	sampleDataBytes, err := json.Marshal(sampleData)
	if err != nil {
		return nil, fmt.Errorf("error marshaling statement: %v", err)
	}
	return sampleDataBytes, err
}
