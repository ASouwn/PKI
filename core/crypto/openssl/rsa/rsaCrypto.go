package rsa

/*
#cgo LDFLAGS: -lssl -lcrypto
#include <openssl/rsa.h>
#include <openssl/pem.h>
#include <openssl/err.h>
#include <stdlib.h>
*/
import "C"
import (
	"encoding/pem"
	"errors"
	"unsafe"

	"github.com/ASouwn/PKI/core/crypto"
)

type RSACrypto struct{}

// Decrypt implements crypto.CryptoInterface.
func (r *RSACrypto) Decrypt(ciphertext []byte, priBlock *pem.Block) ([]byte, error) {
	pemData := pem.EncodeToMemory(priBlock)
	if len(pemData) == 0 {
		return nil, errors.New("failed to encode PEM block to memory")
	}
	// 从 PEM 块中加载私钥
	privateKeyBio := C.BIO_new(C.BIO_s_mem())
	if privateKeyBio == nil {
		return nil, errors.New("failed to create BIO for private key")
	}
	defer C.BIO_free(privateKeyBio)

	if C.BIO_write(privateKeyBio, unsafe.Pointer(&pemData[0]), C.int(len(pemData))) <= 0 {
		return nil, errors.New("failed to write private key to BIO")
	}

	rsaKey := C.PEM_read_bio_RSAPrivateKey(privateKeyBio, nil, nil, nil)
	if rsaKey == nil {
		return nil, errors.New("failed to read RSA private key")
	}
	defer C.RSA_free(rsaKey)

	// 解密数据
	plaintext := make([]byte, C.RSA_size(rsaKey))
	plaintextLen := C.RSA_private_decrypt(C.int(len(ciphertext)), (*C.uchar)(unsafe.Pointer(&ciphertext[0])),
		(*C.uchar)(unsafe.Pointer(&plaintext[0])), rsaKey, C.RSA_PKCS1_OAEP_PADDING)
	if plaintextLen == -1 {
		return nil, errors.New("failed to decrypt data")
	}

	return plaintext[:plaintextLen], nil
}

// Encrypt implements crypto.CryptoInterface.
func (r *RSACrypto) Encrypt(origin []byte, pubBlock *pem.Block) ([]byte, error) {
	pemData := pem.EncodeToMemory(pubBlock)
	if len(pemData) == 0 {
		return nil, errors.New("failed to encode PEM block to memory")
	}
	// 从 PEM 块中加载公钥
	publicKeyBio := C.BIO_new(C.BIO_s_mem())
	if publicKeyBio == nil {
		return nil, errors.New("failed to create BIO for public key")
	}
	defer C.BIO_free(publicKeyBio)

	if C.BIO_write(publicKeyBio, unsafe.Pointer(&pemData[0]), C.int(len(pemData))) <= 0 {
		return nil, errors.New("failed to write public key to BIO")
	}

	rsaKey := C.PEM_read_bio_RSA_PUBKEY(publicKeyBio, nil, nil, nil)
	if rsaKey == nil {
		return nil, errors.New("failed to read RSA public key")
	}
	defer C.RSA_free(rsaKey)

	// 加密数据
	ciphertext := make([]byte, C.RSA_size(rsaKey))
	ciphertextLen := C.RSA_public_encrypt(C.int(len(origin)), (*C.uchar)(unsafe.Pointer(&origin[0])),
		(*C.uchar)(unsafe.Pointer(&ciphertext[0])), rsaKey, C.RSA_PKCS1_OAEP_PADDING)
	if ciphertextLen == -1 {
		return nil, errors.New("failed to encrypt data")
	}

	return ciphertext[:ciphertextLen], nil
}

// GenerateKeyPair implements crypto.CryptoInterface.
func (r *RSACrypto) GenerateKeyPair() (private *pem.Block, public *pem.Block, err error) {
	// 生成 RSA 密钥对
	rsaKey := C.RSA_new()
	if rsaKey == nil {
		return nil, nil, errors.New("failed to create RSA key")
	}
	defer C.RSA_free(rsaKey)

	// 设置公钥指数
	e := C.BN_new()
	if e == nil {
		return nil, nil, errors.New("failed to create BIGNUM")
	}
	defer C.BN_free(e)

	if C.BN_set_word(e, C.RSA_F4) != 1 {
		return nil, nil, errors.New("failed to set RSA exponent")
	}

	// 生成密钥对
	if C.RSA_generate_key_ex(rsaKey, 2048, e, nil) != 1 {
		return nil, nil, errors.New("failed to generate RSA key pair")
	}

	// 将私钥转换为 PEM 格式
	privateKeyBio := C.BIO_new(C.BIO_s_mem())
	if privateKeyBio == nil {
		return nil, nil, errors.New("failed to create BIO for private key")
	}
	defer C.BIO_free(privateKeyBio)

	if C.PEM_write_bio_RSAPrivateKey(privateKeyBio, rsaKey, nil, nil, 0, nil, nil) != 1 {
		return nil, nil, errors.New("failed to write private key to PEM")
	}

	privateKeyData := make([]byte, 4096)
	privateKeyLen := C.BIO_read(privateKeyBio, unsafe.Pointer(&privateKeyData[0]), C.int(len(privateKeyData)))
	if privateKeyLen <= 0 {
		return nil, nil, errors.New("failed to read private key from BIO")
	}
	privateKeyData = privateKeyData[:privateKeyLen]

	// 将公钥转换为 PEM 格式
	publicKeyBio := C.BIO_new(C.BIO_s_mem())
	if publicKeyBio == nil {
		return nil, nil, errors.New("failed to create BIO for public key")
	}
	defer C.BIO_free(publicKeyBio)

	if C.PEM_write_bio_RSA_PUBKEY(publicKeyBio, rsaKey) != 1 {
		return nil, nil, errors.New("failed to write public key to PEM")
	}

	publicKeyData := make([]byte, 4096)
	publicKeyLen := C.BIO_read(publicKeyBio, unsafe.Pointer(&publicKeyData[0]), C.int(len(publicKeyData)))
	if publicKeyLen <= 0 {
		return nil, nil, errors.New("failed to read public key from BIO")
	}
	publicKeyData = publicKeyData[:publicKeyLen]

	// 解析 PEM 块
	privateBlock, _ := pem.Decode(privateKeyData)
	publicBlock, _ := pem.Decode(publicKeyData)

	return privateBlock, publicBlock, nil
}

// Sign implements crypto.CryptoInterface.
func (r *RSACrypto) Sign(origin []byte, priBlock *pem.Block) ([]byte, error) {
	pemData := pem.EncodeToMemory(priBlock)
	if len(pemData) == 0 {
		return nil, errors.New("failed to encode PEM block to memory")
	}
	// 从 PEM 块中加载私钥
	privateKeyBio := C.BIO_new(C.BIO_s_mem())
	if privateKeyBio == nil {
		return nil, errors.New("failed to create BIO for private key")
	}
	defer C.BIO_free(privateKeyBio)

	if C.BIO_write(privateKeyBio, unsafe.Pointer(&pemData[0]), C.int(len(pemData))) <= 0 {
		return nil, errors.New("failed to write private key to BIO")
	}

	rsaKey := C.PEM_read_bio_RSAPrivateKey(privateKeyBio, nil, nil, nil)
	if rsaKey == nil {
		return nil, errors.New("failed to read RSA private key")
	}
	defer C.RSA_free(rsaKey)

	// 签名数据
	signature := make([]byte, C.RSA_size(rsaKey))
	var signatureLen C.uint

	// 处理空数据的情况
	var originPtr *C.uchar
	if len(origin) > 0 {
		originPtr = (*C.uchar)(unsafe.Pointer(&origin[0]))
	} else {
		// 对于空数据，传递一个非空指针
		originPtr = (*C.uchar)(unsafe.Pointer(&[]byte("null")[0]))
	}

	// sigh算法并不是直接根据origin的指针来确定，而是通过origin的长度来确定是否为空
	if C.RSA_sign(C.NID_sha256, originPtr, C.uint(len(origin)),
		(*C.uchar)(unsafe.Pointer(&signature[0])), &signatureLen, rsaKey) != 1 {
		return nil, errors.New("failed to sign data")
	}

	return signature[:signatureLen], nil
}

// Verify implements crypto.CryptoInterface.
func (r *RSACrypto) Verify(origin []byte, signature []byte, pubBlock *pem.Block) (bool, error) {
	pemData := pem.EncodeToMemory(pubBlock)
	if len(pemData) == 0 {
		return false, errors.New("failed to encode PEM block to memory")
	}
	// 从 PEM 块中加载公钥
	publicKeyBio := C.BIO_new(C.BIO_s_mem())
	if publicKeyBio == nil {
		return false, errors.New("failed to create BIO for public key")
	}
	defer C.BIO_free(publicKeyBio)

	if C.BIO_write(publicKeyBio, unsafe.Pointer(&pemData[0]), C.int(len(pemData))) <= 0 {
		return false, errors.New("failed to write public key to BIO")
	}

	rsaKey := C.PEM_read_bio_RSA_PUBKEY(publicKeyBio, nil, nil, nil)
	if rsaKey == nil {
		return false, errors.New("failed to read RSA public key")
	}
	defer C.RSA_free(rsaKey)

	// 验证签名
	if C.RSA_verify(C.NID_sha256, (*C.uchar)(unsafe.Pointer(&origin[0])), C.uint(len(origin)),
		(*C.uchar)(unsafe.Pointer(&signature[0])), C.uint(len(signature)), rsaKey) != 1 {
		return false, nil
	}

	return true, nil
}

var _ crypto.CryptoInterface = (*RSACrypto)(nil)
