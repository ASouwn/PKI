package crypto

import "github.com/ASouwn/PKI/core/crypto/go_crypto/rsa"

func getCryptoInstance() CryptoInterface {
	instance := &rsa.RSACrypto{}
	return instance
}
