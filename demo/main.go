package main

import (
	"crypto/x509/pkix"
	"encoding/pem"
	"log"

	rpctypes "github.com/ASouwn/PKI/shared-rpc-types"
	"github.com/ASouwn/PKI/utils"
)

// func main() {
// 	// Generate a new key pair
// 	log.Println("Generating a new key pair...")
// 	pri, pub, _ := utils.GenerateKeyPair()

// 	// verify the key pair
// 	originMsg := "hello world"
// 	encryptedMsg, _ := utils.Encrypt([]byte(originMsg), pub)
// 	decryptedMsg, _ := utils.Decrypt(encryptedMsg, pri)
// 	log.Printf("Verify key pair: %v\n", string(decryptedMsg) == originMsg)

// 	// send the public key to the RA to verify the identity and submit the CSR
// 	// create a new CSR
// 	csr, _ := utils.CreateCSR(pri, pkix.Name{
// 		CommonName:   "ahsown",
// 		Organization: []string{"asouwn.xyz"},
// 		Country:      []string{"CN"},
// 	})
// 	log.Printf("CSR: \n%+s\n", pem.EncodeToMemory(csr))

// 	// submit the CSR to RA and get the x509 certificate from ca
// 	cer, _ := utils.SubmitCSRToRA(csr, "localhost:3001")
// 	log.Printf("Submit CSR to CA: \n%+v\n", cer)
// }

func main() {
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
	cert, ok := reply.(*pem.Block)
	if !ok {
		log.Printf("failed to assert type *pem.Block\n")
	}
	log.Printf("got reply: \n%+s\n", pem.EncodeToMemory(cert))
}
