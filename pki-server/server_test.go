package pkiserver

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"log"
	"testing"

	rpctypes "github.com/ASouwn/PKI/shared-rpc-types"
	"github.com/ASouwn/PKI/utils"
)

func TestCAServer(t *testing.T) {
	privPem, _, err := utils.GenerateKeyPair()
	if err != nil {
		log.Printf("failed to generate key pair: %v\n", err)
	}
	argPem, err := utils.CreateCSR(privPem, pkix.Name{
		CommonName:   "ahsown",
		Organization: []string{"asouwn.xyz"},
		Country:      []string{"CN"},
	})
	if err != nil {
		log.Printf("failed to create CSR: %v\n", err)
	}

	reply, err := utils.GetRedServer(rpctypes.CAHandleCSRMethod, argPem, "localhost:3001")
	if err != nil {
		log.Printf("failed to get server: %v\n", err)
	}
	cert, ok := reply.(*x509.Certificate)
	if !ok {
		log.Printf("failed to assert type *x509.Certificate\n")
	}
	log.Printf("got reply: %v\n", cert)
}
