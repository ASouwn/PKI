package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

// GenerateKeyPair 生成 RSA 2048长度 密钥对，并返回 PEM 格式的私钥和公钥块
func GenerateKeyPair() (private *pem.Block, public *pem.Block, err error) {
	// 生成 RSA 密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048) // 2048 是密钥长度
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate RSA key pair: %v", err)
	}

	// 将私钥编码为 PEM 格式
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	// 将公钥编码为 PEM 格式
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal public key: %v", err)
	}
	publicKeyPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	return privateKeyPEM, publicKeyPEM, nil
}

// Encrypt 使用 RSA 公钥加密数据
func Encrypt(origin []byte, pubBlock *pem.Block) ([]byte, error) {
	// 解码 PEM 块中的公钥
	publicKey, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}

	// 类型断言为 RSA 公钥
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("invalid public key type, expected RSA public key")
	}

	// 使用 RSA 公钥加密数据
	ciphertext, err := rsa.EncryptOAEP(
		sha256.New(), // 使用 SHA-256 作为哈希函数
		rand.Reader,  // 随机数生成器
		rsaPublicKey, // RSA 公钥
		origin,       // 原始数据
		nil,          // 可选标签
	)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt data: %v", err)
	}

	return ciphertext, nil
}

// Decrypt 使用 RSA 私钥解密数据
func Decrypt(ciphertext []byte, priBlock *pem.Block) ([]byte, error) {
	// 解码 PEM 块中的私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(priBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	// 使用 RSA 私钥解密数据
	plaintext, err := rsa.DecryptOAEP(
		sha256.New(), // 使用 SHA-256 作为哈希函数
		rand.Reader,  // 随机数生成器
		privateKey,   // RSA 私钥
		ciphertext,   // 加密数据
		nil,          // 可选标签
	)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt data: %v", err)
	}

	return plaintext, nil
}
