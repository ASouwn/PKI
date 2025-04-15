package pkiserver

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"reflect"

	rpctypes "github.com/ASouwn/PKI/shared-rpc-types"
	"github.com/ASouwn/PKI/utils"
)

type RA struct{}

// 验证证书申请者的身份
func ValidateIdentity(csr *x509.CertificateRequest) error {
	if csr == nil {
		return fmt.Errorf("CSR is nil")
	}
	return nil
}

func (r *RA) HandleCSR(csrPem *pem.Block, reply *pem.Block) error {
	log.Printf("get and handle CSR\n")
	if csrPem.Type != "CERTIFICATE REQUEST" {
		log.Printf("got wrong csr type\n")
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
	log.Printf("trying to call GetRedServer with method(%s) and args(%s)\n", rpctypes.CAHandleCSRMethod, reflect.TypeOf(csrRequest))
	cert, err := utils.GetRedServer(rpctypes.CAHandleCSRMethod, csrRequest, "localhost:3001")
	if err != nil {
		return fmt.Errorf("got wrong when submit csr to ca: %v", err)
	}
	log.Printf("after submit csr to ca\n")
	certParsed, ok := cert.(*pem.Block)
	if !ok {
		return fmt.Errorf("failed to assert type *pem.Block")
	}

	*reply = *certParsed
	return nil
}

var _ rpctypes.RAServer = (*RA)(nil)

func StarRaServer(port, registerAddr string) {

	log.Println("Trying to regist Handle CSR server to register server")

	utils.WriteServer(&rpctypes.Server{
		ServerKey: rpctypes.RAServerKey,
		ServerInfo: rpctypes.ServerInfo{
			Network:       "tcp",
			ServerAddress: "localhost",
			Port:          port,
		},
	}, registerAddr)
	// Connect to the register server
	// client, err := rpc.DialHTTP("tcp", registerAddr)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to register server: %v", err)
	// }
	// defer client.Close()
	// Register the server with the RpcRegister service
	// server := &rpctypes.Server{
	// 	ServerInfo: rpctypes.ServerInfo{
	// 		Network:       "tcp",
	// 		ServerAddress: "localhost",
	// 		Port:          port,
	// 	},
	// 	ServerKey: rpctypes.RAServerKey,
	// }
	// var reply string
	// client.Call(rpctypes.WriteServerMethod, server, &reply)
	// log.Printf("Reply: %s\n", reply)

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
