package ecc

import (
	"encoding/pem"

	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"errors"
	"math/big"

	"github.com/ASouwn/PKI/pki-server/crypto"
)

type ECCCrypto struct{}

// Decrypt implements crypto.CryptoInterface.
func (e *ECCCrypto) Decrypt(ciphertext []byte, priBlock *pem.Block) ([]byte, error) {
	// ECC is not typically used for direct encryption/decryption, but for key agreement.
	// Here, we return an error to indicate it's not supported.
	return nil, errors.New("ECC does not support direct decryption")
}

// Encrypt implements crypto.CryptoInterface.
func (e *ECCCrypto) Encrypt(origin []byte, pubBlock *pem.Block) ([]byte, error) {
	// ECC is not typically used for direct encryption/decryption, but for key agreement.
	// Here, we return an error to indicate it's not supported.
	return nil, errors.New("ECC does not support direct encryption")
}

// GenerateKeyPair implements crypto.CryptoInterface.
func (e *ECCCrypto) GenerateKeyPair() (private *pem.Block, public *pem.Block, err error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	privBytes, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		return nil, nil, err
	}
	pubBytes, err := x509.MarshalPKIXPublicKey(&priv.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	privBlock := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: privBytes,
	}
	pubBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	}
	return privBlock, pubBlock, nil
}

// Sign implements crypto.CryptoInterface.
func (e *ECCCrypto) Sign(origin []byte, priBlock *pem.Block) ([]byte, error) {
	privKey, err := x509.ParseECPrivateKey(priBlock.Bytes)
	if err != nil {
		return nil, err
	}
	hash := sha256.Sum256(origin)
	r, s, err := ecdsa.Sign(rand.Reader, privKey, hash[:])
	if err != nil {
		return nil, err
	}
	return asn1.Marshal(struct{ R, S *big.Int }{r, s})
}

// Verify implements crypto.CryptoInterface.
func (e *ECCCrypto) Verify(origin []byte, signature []byte, pubBlock *pem.Block) (bool, error) {
	pubKeyIfc, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		return false, err
	}
	pubKey, ok := pubKeyIfc.(*ecdsa.PublicKey)
	if !ok {
		return false, errors.New("not ECDSA public key")
	}
	var sig struct{ R, S *big.Int }
	_, err = asn1.Unmarshal(signature, &sig)
	if err != nil {
		return false, err
	}
	hash := sha256.Sum256(origin)
	return ecdsa.Verify(pubKey, hash[:], sig.R, sig.S), nil
}

var _ crypto.CryptoInterface = (*ECCCrypto)(nil)
