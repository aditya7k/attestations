package security

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/sigstore/sigstore/pkg/signature"
)

// LocalSigner signs data locally using a private key
type LocalSigner struct {
	// privateKey is the private key used for signing
	privateKey *ecdsa.PrivateKey
	// PublicKey is the public key used for verification
	publicKey *ecdsa.PublicKey
	// PublicKeyPath is the path to the public key
	PublicKeyPath string
}

func (s *LocalSigner) Sign(data []byte) ([]byte, error) {
	// Sign the data using the private key

	// Step 5: Sign the attestation using the generated private key
	signer, err := signature.LoadECDSASigner(s.privateKey, crypto.SHA256)
	if err != nil {
		return nil, fmt.Errorf("error loading ECDSA signer: %v\n", err)
	}

	// Sign the attestation content
	sig, err := signer.SignMessage(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("error signing message: %v\n", err)
	}

	return sig, nil
}

func (s *LocalSigner) LoadKeyPair(privateKeyPath string, publicKeyPath string) error {

	privateKey, err := loadPrivateKey(privateKeyPath)
	if err != nil {
		return fmt.Errorf("failed to load private key: %w", err)
	}
	s.privateKey = privateKey

	publicKey, err := loadPublicKey(publicKeyPath)
	if err != nil {
		return fmt.Errorf("failed to load public key: %w", err)
	}
	s.publicKey = publicKey

	return nil
}

func (s *LocalSigner) VerifySignature(data []byte, signature []byte) (bool, error) {
	// Verify the signature
	return false, nil
}

func loadPrivateKey(filePath string) (*ecdsa.PrivateKey, error) {
	// Step 1: Read the private key from the file
	privKeyPEM, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read private key file: %w", err)
	}

	// Step 2: Decode the PEM block
	block, _ := pem.Decode(privKeyPEM)
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	// Step 3: Parse the private key
	privKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ECDSA private key: %w", err)
	}

	return privKey, nil
}

func loadPublicKey(filePath string) (*ecdsa.PublicKey, error) {
	// Read the public key from the file
	pubKeyPEM, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read public key file: %w", err)
	}

	// Decode the PEM block
	block, _ := pem.Decode(pubKeyPEM)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	// Parse the public key
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ECDSA public key: %w", err)
	}

	// Type assert to *ecdsa.PublicKey
	ecdsaPubKey, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an ECDSA public key")
	}

	return ecdsaPubKey, nil
}
