package utils

import (
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"strings"
)

// 生成证书签名请求（CSR PEM格式）
// CreateCSR generates a Certificate Signing Request (CSR) using the provided private key and subject information.
// The privateKey should be a PEM-encoded private key, and the subject specifies the details of the entity requesting the certificate.
func CreateCSR(privateBlock *pem.Block, subject pkix.Name) (*pem.Block, error) {
	if !strings.Contains(privateBlock.Type, "PRIVATE KEY") {
		log.Printf("Invalid private key PEM block type: %s", privateBlock.Type)
		return nil, errors.New("invalid private key PEM block type")
	}
	// Decode the private key from the PEM block
	privKey, err := x509.ParsePKCS1PrivateKey(privateBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	// Create a CSR template with the provided subject
	csrTemplate := x509.CertificateRequest{
		Subject: subject,
	}

	// Generate the CSR
	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, privKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create CSR: %v", err)
	}

	// Encode the CSR to PEM format
	return &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes}, nil
}

// 将CSR提交到ra，由ra提交给ca并返回证书
func SubmitCSRToRA(csrPem *pem.Block, raAddr string) (*pem.Block, error) {
	// Verify the CSR
	// Parse the CSR from the PEM block
	_, err := x509.ParseCertificateRequest(csrPem.Bytes)
	if err != nil {
		log.Printf("Failed to parse CSR: %v", err)
		return nil, err
	}

	// Send the CSR to the RA for verification and submission to the CA
	// By RPC

	return nil, nil
}
