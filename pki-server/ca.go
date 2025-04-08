package pkiserver

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net"
	"net/http"
	"net/rpc"
	"time"

	rpctypes "github.com/ASouwn/PKI/shared-rpc-types"
	"github.com/ASouwn/PKI/utils"
)

var (
	caCertPath = "./relese/certificate/ca.cert"
	caKeyPath  = "./relese/certificate/caKey.pem"
)

// 生成自签名根证书
func CreateRootCA(subject pkix.Name, validity time.Duration) (*x509.Certificate, *pem.Block, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	template := &x509.Certificate{
		SerialNumber:          GenerateSerialNumber(),
		Subject:               subject,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(validity),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, template, template, &privKey.PublicKey, privKey)
	if err != nil {
		return nil, nil, err
	}

	cert, err := x509.ParseCertificate(certBytes)
	if err != nil {
		return nil, nil, err
	}

	privKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey),
	}

	return cert, privKeyPEM, nil
}

// 签发中间CA证书
func IssueIntermediateCA(parentCert *x509.Certificate, parentKey *pem.Block, subject pkix.Name, validity time.Duration) (*x509.Certificate, *pem.Block, error) {
	parentPrivKey, err := parsePrivateKey(parentKey)
	if err != nil {
		return nil, nil, err
	}

	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	template := &x509.Certificate{
		SerialNumber:          GenerateSerialNumber(),
		Subject:               subject,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(validity),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, template, parentCert, &privKey.PublicKey, parentPrivKey)
	if err != nil {
		return nil, nil, err
	}

	cert, err := x509.ParseCertificate(certBytes)
	if err != nil {
		return nil, nil, err
	}

	privKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey),
	}

	return cert, privKeyPEM, nil
}

// Helper function to generate a serial number
func GenerateSerialNumber() *big.Int {
	serial, _ := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	return serial
}

// Helper function to parse private key from PEM block
func parsePrivateKey(block *pem.Block) (*rsa.PrivateKey, error) {
	if block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("invalid private key type")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

type CA struct{}

var _ rpctypes.CAServer = (*CA)(nil)

func (c *CA) HandleCSR(csr *x509.CertificateRequest, reply *x509.Certificate) error {
	log.Printf("try to handle csr from RA")
	// 加载ca的证书与私钥
	caCert, caKey, err := LoadCertAndKeyFromFile(caCertPath, caKeyPath)
	// 创建证书模板
	template := &x509.Certificate{
		SerialNumber:          GenerateSerialNumber(),
		Subject:               csr.Subject,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}
	// 4. 签发证书
	certBytes, err := x509.CreateCertificate(
		rand.Reader,
		template,
		caCert,
		csr.PublicKey,
		caKey,
	)
	if err != nil {
		return err
	}
	cer, err := x509.ParseCertificate(certBytes)
	*reply = *cer
	return nil
}

func StartCAServer(port, registerAddr string) {
	// Generate CA cert and key
	cert, privPem, err := CreateRootCA(pkix.Name{
		CommonName:         "My Root CA",
		Organization:       []string{"NJUPT"},
		OrganizationalUnit: []string{"NJUPT"},
		Country:            []string{"CN"},
		Province:           []string{"JianSu"},
		Locality:           []string{"NJUPT"},
	}, 365*24*time.Hour)
	if err != nil {
		log.Fatalf("Error creating root CA: %v\n", err)
	}
	err = SaveCertAndKeyToFile(cert, privPem, caCertPath, caKeyPath)
	if err != nil {
		fmt.Printf("got wrong when save rootcert and keypair: %v\n", err)
	}

	// Submit CAServer to regist center
	utils.WriteServer(&rpctypes.Server{
		ServerKey: rpctypes.CAServerKey,
		ServerInfo: rpctypes.ServerInfo{
			Network:       "tcp",
			ServerAddress: "localhost",
			Port:          port,
		},
	}, registerAddr)

	// Start the CA server
	log.Printf("Starting CA server on port %s...", port)
	rpc.RegisterName(rpctypes.CAServerKey.ServerName, new(CA))
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error starting listener: %v", err)
	}

	log.Printf("CA server is running on port %s...\n", port)
	http.Serve(l, nil)
}
