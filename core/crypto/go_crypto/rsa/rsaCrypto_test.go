package rsa

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"reflect"
	"testing"
)

func TestRSACrypto_GenerateKeyPair(t *testing.T) {
	tests := []struct {
		name        string
		r           *RSACrypto
		wantPrivate *pem.Block
		wantPublic  *pem.Block
		wantErr     bool
	}{
		// TODO: Add test cases.
		{
			name:    "generateKeyPair",
			r:       &RSACrypto{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.r
			gotPrivate, gotPublic, err := r.GenerateKeyPair()

			if (err != nil) != tt.wantErr {
				t.Errorf("RSACrypto.GenerateKeyPair() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// public key
				if gotPrivate == nil || gotPrivate.Type != "RSA PRIVATE KEY" {
					t.Error("Generated private key is invalid")
				}

				// private key
				if gotPublic == nil || gotPublic.Type != "PUBLIC KEY" {
					t.Error("Generated public key is invalid")
				}

				// if rsa private key is valid
				_, err := x509.ParsePKCS1PrivateKey(gotPrivate.Bytes)
				if err != nil {
					t.Errorf("Failed to parse private key: %v", err)
				}

				// if rsa public key is valid
				_, err = x509.ParsePKIXPublicKey(gotPublic.Bytes)
				if err != nil {
					t.Errorf("Failed to parse public key: %v", err)
				}
			}
		})
	}
}

// func TestRSACrypto_Encrypt(t *testing.T) {
// 	type args struct {
// 		origin   []byte
// 		pubBlock *pem.Block
// 	}
// 	tests := []struct {
// 		name    string
// 		r       *RSACrypto
// 		args    args
// 		want    []byte
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := tt.r
// 			got, err := r.Encrypt(tt.args.origin, tt.args.pubBlock)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("RSACrypto.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("RSACrypto.Encrypt() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestRSACrypto_Decrypt(t *testing.T) {
	msg := "Long live the great People's Republic of China"

	tests := []struct {
		name    string
		r       *RSACrypto
		args    []byte
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "decrypt",
			r:       &RSACrypto{},
			args:    []byte(msg),
			want:    []byte(msg),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.r
			pri, pub, err := r.GenerateKeyPair()

			if (err != nil) != tt.wantErr {
				t.Errorf("RSACrypto.GenerateKeyPair() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			ciphertest, err := r.Encrypt(tt.args, pub)

			if (err != nil) != tt.wantErr {
				t.Errorf("RSACrypto.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := r.Decrypt(ciphertest, pri)

			if (err != nil) != tt.wantErr {
				t.Errorf("RSACrypto.Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RSACrypto.Decrypt() = %v, want %v", got, tt.want)
				return
			}

			fmt.Printf("the msg is: %v\nthe ciphertext is: %v\nthe got is: %s", msg, ciphertest, got)

		})
	}
}

func TestRSACrypto_Sign(t *testing.T) {
	// 生成测试密钥对
	r := &RSACrypto{}
	priBlock, pubBlock, err := r.GenerateKeyPair()
	if err != nil {
		t.Fatalf("生成测试密钥失败: %v", err)
	}

	// 生成错误私钥（无效PEM）
	invalidPriBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: []byte("invalid_private_key"),
	}

	tests := []struct {
		name        string
		args        func() ([]byte, *pem.Block) // 动态生成测试数据
		wantErr     bool
		expectValid bool
	}{
		{
			name: "正常签名-短文本",
			args: func() ([]byte, *pem.Block) {
				return []byte("Hello PKI"), priBlock
			},
			wantErr:     false,
			expectValid: true,
		},
		{
			name: "空数据签名",
			args: func() ([]byte, *pem.Block) {
				return []byte(""), priBlock
			},
			wantErr:     false,
			expectValid: true,
		},
		{
			name: "无效私钥签名",
			args: func() ([]byte, *pem.Block) {
				return []byte("test data"), invalidPriBlock
			},
			wantErr:     true,
			expectValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, key := tt.args()
			signature, err := r.Sign(data, key)

			if (err != nil) != tt.wantErr {
				t.Errorf("Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// 当期待有效时验证签名
			if tt.expectValid {
				valid, verifyErr := r.Verify(data, signature, pubBlock)
				if verifyErr != nil {
					t.Errorf("Verify() unexpected error: %v", verifyErr)
				}
				if !valid {
					t.Error("生成的签名无法通过验证")
				}
			}
		})
	}
}

func TestRSACrypto_Verify(t *testing.T) {
	r := &RSACrypto{}
	priBlock, pubBlock, _ := r.GenerateKeyPair()
	wrongPubBlock, _, _ := r.GenerateKeyPair() // 生成另一个密钥对的公钥

	// 生成有效签名
	data := []byte("Verification Test")
	validSig, _ := r.Sign(data, priBlock)

	tests := []struct {
		name    string
		args    func() ([]byte, []byte, *pem.Block)
		want    bool
		wantErr bool
	}{
		{
			name: "有效签名验证",
			args: func() ([]byte, []byte, *pem.Block) {
				return data, validSig, pubBlock
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "错误公钥验证",
			args: func() ([]byte, []byte, *pem.Block) {
				return data, validSig, wrongPubBlock
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "篡改原始数据",
			args: func() ([]byte, []byte, *pem.Block) {
				modifiedData := append([]byte(nil), data...)
				modifiedData[0] ^= 0xFF // 翻转第一个字节
				return modifiedData, validSig, pubBlock
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "篡改签名数据",
			args: func() ([]byte, []byte, *pem.Block) {
				modifiedSig := append([]byte(nil), validSig...)
				modifiedSig[0] ^= 0xFF
				return data, modifiedSig, pubBlock
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "无效公钥格式",
			args: func() ([]byte, []byte, *pem.Block) {
				invalidPub := &pem.Block{
					Type:  "PUBLIC KEY",
					Bytes: []byte("invalid_public_key"),
				}
				return data, validSig, invalidPub
			},
			want:    false,
			wantErr: true, // 预期解析公钥出错
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, sig, pub := tt.args()
			got, err := r.Verify(data, sig, pub)

			if (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}
