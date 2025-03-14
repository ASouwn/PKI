package rsa

import (
	"encoding/pem"

	"github.com/ASouwn/PKI/core/crypto"
)

type RSACrypto struct{}

// Decrypt implements crypto.CryptoInterface.
func (r *RSACrypto) Decrypt(ciphertext []byte, priBlock *pem.Block) ([]byte, error) {
	panic("unimplemented")
}

// Encrypt implements crypto.CryptoInterface.
func (r *RSACrypto) Encrypt(origin []byte, pubBlock *pem.Block) ([]byte, error) {
	panic("unimplemented")
}

// GenerateKeyPair implements crypto.CryptoInterface.
func (r *RSACrypto) GenerateKeyPair() (private *pem.Block, public *pem.Block, err error) {
	panic("unimplemented")
}

// Sign implements crypto.CryptoInterface.
func (r *RSACrypto) Sign(origin []byte, priBlock *pem.Block) ([]byte, error) {
	panic("unimplemented")
}

// Verify implements crypto.CryptoInterface.
func (r *RSACrypto) Verify(origin []byte, signature []byte, pubBlock *pem.Block) (bool, error) {
	panic("unimplemented")
}

var _ crypto.CryptoInterface = (*RSACrypto)(nil)
