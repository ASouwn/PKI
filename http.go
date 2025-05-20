package main

import (
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"net/http"

	"github.com/ASouwn/PKI/utils"
)

func main() {
	log.Println("Generating a new key pair...")
	pri, _, _ := utils.GenerateKeyPair()

	csr, _ := utils.CreateCSR(pri, pkix.Name{
		CommonName:   "ahsown",
		Organization: []string{"asouwn.xyz"},
		Country:      []string{"CN"},
	})

	cer, _ := utils.SubmitCSRToRA(csr, "localhost:3001")

	http.HandleFunc("/key", func(w http.ResponseWriter, r *http.Request) {
		// Add CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Private Key: \n%s\n", pem.EncodeToMemory(pri))
		// 客户端不用展示公钥
		// fmt.Fprintf(w, "Public Key: \n%s\n", pem.EncodeToMemory(pub))
	})

	http.HandleFunc("/csr", func(w http.ResponseWriter, r *http.Request) {
		// Add CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		// Handle preflight requests
		if r.Method == http.MethodOptions {
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "CSR: \n%s\n", pem.EncodeToMemory(csr))
	})
	http.HandleFunc("/cert", func(w http.ResponseWriter, r *http.Request) {
		// Add CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		// Handle preflight requests
		if r.Method == http.MethodOptions {
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Certificate: \n%s\n", pem.EncodeToMemory(cer))
	})
	http.ListenAndServe(":8080", nil)
}
