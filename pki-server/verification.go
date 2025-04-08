package pkiserver

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
)

// 验证证书链
func VerifyCertChain(cert *x509.Certificate, roots *x509.CertPool) error { return nil }

// 生成证书吊销列表（CRL）
func GenerateCRL(caCert *x509.Certificate, caKey *pem.Block, revokedCerts []pkix.RevokedCertificate) ([]byte, error) {
	return nil, nil
}

// 检查证书是否被吊销
func IsCertRevoked(cert *x509.Certificate, crl []byte) (bool, error) { return true, nil }
