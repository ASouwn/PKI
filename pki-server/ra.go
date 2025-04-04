package pkiserver

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
)

type RA struct{}

// 验证证书申请者的身份
func ValidateIdentity(csr *x509.CertificateRequest) error {
	if csr == nil {
		return fmt.Errorf("CSR is nil")
	}
	return nil
}

// 提交证书申请到CA
func SubmitCSRToCA(csr *x509.CertificateRequest, addr string) (*pem.Block, error) {
	return nil, nil
}

func (r *RA) HandleCSR(csrPem *pem.Block) (*pem.Block, error) {
	if csrPem.Type != "CERTIFICATE REQUEST" {
		return nil, nil
	}
	csrRequest, err := x509.ParseCertificateRequest(csrPem.Bytes)
	if err != nil {
		log.Printf("Failed to parse CSR: %v", err)
		return nil, err
	}

	// Validate the CSR
	if err := ValidateIdentity(csrRequest); err != nil {
		return nil, err
	}

	// Submit the CSR to CA
	cer, err := SubmitCSRToCA(csrRequest, "localhost:8081")
	if err != nil {
		return nil, err
	}
	return cer, nil
}

func (r *RA) SayHello(a struct{}, reply *string) error {
	*reply = "Hello, RA!"
	return nil
}

func StarRaServer(port int) {
	// Start the RA server
	// Handle the CSR submission
	log.Printf("Starting RA server on port %d...", port)

}
