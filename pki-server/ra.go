package pkiserver

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	rpctypes "github.com/ASouwn/PKI/shared-rpc-types"
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

func (r *RA) HandleCSR(csrPem *pem.Block, reply *string) error {
	log.Printf("get and handle CSR\n")
	if csrPem.Type != "CERTIFICATE REQUEST" {
		return nil
	}
	csrRequest, err := x509.ParseCertificateRequest(csrPem.Bytes)
	if err != nil {
		log.Printf("Failed to parse CSR: %v", err)
		return err
	}

	// Validate the CSR
	if err := ValidateIdentity(csrRequest); err != nil {
		return err
	}

	// Submit the CSR to CA
	// cer, err := SubmitCSRToCA(csrRequest, "localhost:8081")
	// if err != nil {
	// 	return err
	// }

	*reply = "send CSR to CA successfully"
	return nil
}

var _ rpctypes.RAServer = (*RA)(nil)

func StarRaServer(port, registerAddr string) {

	log.Println("Trying to regist Handle CSR server to register server")
	// Connect to the register server
	client, err := rpc.DialHTTP("tcp", registerAddr)
	if err != nil {
		log.Fatalf("Failed to connect to register server: %v", err)
	}
	defer client.Close()
	// Register the server with the RpcRegister service
	server := &rpctypes.Server{
		ServerInfo: rpctypes.ServerInfo{
			Network:       "tcp",
			ServerAddress: "localhost",
			Port:          port,
		},
		ServerKey: rpctypes.RAServerKey,
	}
	var reply string
	client.Call(rpctypes.WriteServerMethod, server, &reply)
	log.Printf("Reply: %s\n", reply)

	// Start the RA server
	// Handle the CSR submission
	log.Printf("Starting RA server on port %s...", port)
	rpc.RegisterName(rpctypes.RAServerKey.ServerName, new(RA))
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error starting listener: %v", err)
	}

	http.Serve(l, nil)
	log.Printf("RA server is running on port %s...\n", port)

}
