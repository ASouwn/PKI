package pkiserver

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type SayHello struct{}

func (s *SayHello) Hello(args *struct{}, re *string) error {
	log.Printf("trying to call Hello")
	*re = "Hello, this is PKI!"
	return nil
}

// port like "1234" or "8080", else you like
func StartPKIServer(port string) {
	// Register the SayHello struct as an RPC service
	err := rpc.Register(new(SayHello))
	if err != nil {
		log.Fatal("Error registering SayHello:", err)
	}

	// Register the HTTP handler for RPC
	rpc.HandleHTTP()

	// Start the RPC server
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("Error starting listener:", err)
	}
	defer listener.Close()

	log.Printf("RPC server is running on port %s...\n", port)
	http.Serve(listener, nil)
	log.Println("RPC server stopped.")
}
