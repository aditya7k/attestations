package security

type Signer interface {
	// Sign signs the given data and returns the signature
	Sign(data []byte) ([]byte, error)
	// LoadKeyPair loads the private and public key pair from the given paths
	LoadKeyPair(privateKeyPath string, publicKeyPath string) error
	// VerifySignature verifies the signature of the given data
	VerifySignature(data []byte, signature []byte) (bool, error)
}
