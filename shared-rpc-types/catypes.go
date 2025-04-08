package sharedrpctypes

import "crypto/x509"

type CAServer interface {
	HandleCSR(csr *x509.CertificateRequest, reply *x509.Certificate) error
}
