package crypto

import "encoding/pem"

// Crypto interface
type CryptoInterface interface {
	// 生成密钥对，并返回 PEM 格式的私钥和公钥块
	GenerateKeyPair() (private *pem.Block, public *pem.Block, err error)
	// 使用公钥对原始数据进行加密，并返回加密后的字节切片
	Encrypt(origin []byte, pubBlock *pem.Block) ([]byte, error)
	// 使用私钥对加密数据进行解密，并返回解密后的字节切片
	Decrypt(ciphertext []byte, priBlock *pem.Block) ([]byte, error)
	// 使用私钥对原始数据进行签名，并返回签名后的字节切片
	Sign(origin []byte, priBlock *pem.Block) ([]byte, error)
	// 使用公钥验证签名，并返回验证结果和可能的错误
	Verify(origin, signature []byte, pubBlock *pem.Block) (bool, error)
}
