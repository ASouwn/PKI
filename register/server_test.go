package register

import (
	"log"
	"net/rpc"
	"testing"

	rpctypes "github.com/ASouwn/PKI/shared-rpc-types"
)

func TestClient(t *testing.T) {
	client, err := rpc.DialHTTP("tcp", "localhost:3001")
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer client.Close()

	args := &Server{
		ServerInfo: ServerInfo{
			Network:       "tcp",
			ServerAddress: "localhost",
			Port:          "8080",
		},
		ServerKey: ServerKey{
			ServerName: "hello",
		},
	}
	var reply string
	err = client.Call(rpctypes.WriteServerMethod, args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	log.Printf("Reply: %s\n", reply)
}

func TestGetServer(t *testing.T) {
	client, err := rpc.DialHTTP("tcp", "localhost:3001")
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer client.Close()

	args := &ServerKey{
		ServerName: "hello",
	}
	var reply ServerInfo
	err = client.Call(rpctypes.GetServerMethod, args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	log.Printf("Reply: %s\n", reply)
	if reply.ServerAddress == "" && reply.Port == "" && reply.Network == "" {
		t.Fatalf("Failed to get server: %+v", reply)
	} else {
		log.Printf("Server info: %+v\n", reply)
	}
}
