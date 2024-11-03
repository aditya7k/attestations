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
}

func (s *LocalSigner) Sign(data []byte) ([]byte, error) {

	signer, err := signature.LoadECDSASigner(s.privateKey, crypto.SHA256)
	if err != nil {
		return nil, fmt.Errorf("error loading ECDSA signer: %w\n", err)
	}

	sig, err := signer.SignMessage(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("error signing message: %w\n", err)
	}

	return sig, nil
}

func (s *LocalSigner) LoadKeyPair(privateKeyPath string, publicKeyPath string) error {

	privateKey, err := loadPrivateKey(privateKeyPath)
	if err != nil {
		return err
	}
	s.privateKey = privateKey

	publicKey, err := loadPublicKey(publicKeyPath)
	if err != nil {
		return err
	}
	s.publicKey = publicKey

	return nil
}

func (s *LocalSigner) LoadKeyPairBytes(privateKeyBytes []byte, publicKeyBytes []byte) error {

	privateKey, err := loadPrivateKeyBytes(privateKeyBytes)
	if err != nil {
		return err
	}
	s.privateKey = privateKey

	publicKey, err := loadPublicKeyBytes(publicKeyBytes)
	if err != nil {
		return err
	}
	s.publicKey = publicKey

	return nil
}

func (s *LocalSigner) VerifySignature(data []byte, signatureBytes []byte) (bool, error) {

	verifier, err := signature.LoadECDSAVerifier(s.publicKey, crypto.SHA256)
	if err != nil {
		return false, fmt.Errorf("error loading ECDSA verifier: %v\n", err)
	}

	err = verifier.VerifySignature(bytes.NewReader(signatureBytes), bytes.NewReader(data))
	if err != nil {
		return false, fmt.Errorf("error verifying signature: %v\n", err)
	}

	return true, nil
}

func loadPrivateKey(filePath string) (*ecdsa.PrivateKey, error) {

	privateKeyPEM, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read private key file: %s,  %w", filePath, err)
	}

	privateKey, err := loadPrivateKeyBytes(privateKeyPEM)
	if err != nil {
		return nil, fmt.Errorf("failed to load private key bytes from: %s, %w", filePath, err)
	}
	return privateKey, nil
}

func loadPrivateKeyBytes(privateKeyBytes []byte) (*ecdsa.PrivateKey, error) {

	// Decode the PEM block
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	// Parse the private key
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ECDSA private key: %w", err)
	}

	return privateKey, nil
}

func loadPublicKey(filePath string) (*ecdsa.PublicKey, error) {

	// Read the public key from the file
	publicKeyPEM, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read public key file: %s, %w", filePath, err)
	}

	// Type assert to *ecdsa.PublicKey
	ecdsaPubKey, err := loadPublicKeyBytes(publicKeyPEM)
	if err != nil {
		return nil, fmt.Errorf("failed to load public key bytes from %s", filePath)
	}

	return ecdsaPubKey, nil
}

func loadPublicKeyBytes(publicKeyBytes []byte) (*ecdsa.PublicKey, error) {

	// Decode the PEM block
	block, _ := pem.Decode(publicKeyBytes)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	// Parse the public key
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ECDSA public key: %w", err)
	}

	// Type assert to *ecdsa.PublicKey
	ecdsaPubKey, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an ECDSA public key")
	}

	return ecdsaPubKey, nil
}
