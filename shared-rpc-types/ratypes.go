package sharedrpctypes

import (
	"crypto/x509"
	"encoding/pem"
)

type RAServer interface {
	HandleCSR(csrPem *pem.Block, reply *x509.Certificate) error
}
