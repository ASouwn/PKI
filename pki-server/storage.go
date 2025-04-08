package pkiserver

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
)

// 将证书和私钥保存到文件系统
func SaveCertAndKeyToFile(cert *x509.Certificate, key *pem.Block, certPath, keyPath string) error {
	// 保存证书
	certFile, err := os.Create(certPath)
	if err != nil {
		return err
	}
	defer certFile.Close()

	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw})
	if _, err := certFile.Write(certPEM); err != nil {
		return err
	}

	// 保存私钥
	keyFile, err := os.Create(keyPath)
	if err != nil {
		return err
	}
	defer keyFile.Close()

	keyPEM := pem.EncodeToMemory(key)
	if _, err := keyFile.Write(keyPEM); err != nil {
		return err
	}

	return nil
}

// 从文件系统加载证书和私钥
func LoadCertAndKeyFromFile(certPath, keyPath string) (*x509.Certificate, *pem.Block, error) {
	// 加载证书
	certPEM, err := ioutil.ReadFile(certPath)
	if err != nil {
		return nil, nil, err
	}

	certBlock, _ := pem.Decode(certPEM)
	if certBlock == nil || certBlock.Type != "CERTIFICATE" {
		return nil, nil, errors.New("failed to decode certificate PEM")
	}

	cert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		return nil, nil, err
	}

	// 加载私钥
	keyPEM, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, nil, err
	}

	keyBlock, _ := pem.Decode(keyPEM)
	if keyBlock == nil {
		return nil, nil, errors.New("failed to decode key PEM")
	}

	return cert, keyBlock, nil
}
