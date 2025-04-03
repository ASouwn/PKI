package crypto

import "github.com/ASouwn/PKI/pki_server/crypto/go_crypto/rsa"

func getCryptoInstance() CryptoInterface {
	instance := &rsa.RSACrypto{}
	return instance
}
