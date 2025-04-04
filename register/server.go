package register

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	rpctypes "github.com/ASouwn/PKI/shared-rpc-types"
)

type ServerKey = rpctypes.ServerKey
type ServerInfo = rpctypes.ServerInfo
type Server = rpctypes.Server

type Register struct {
	ServerMap map[ServerKey]ServerInfo
}

func NewRegister() *Register {
	log.Println("init a Register instance")
	return &Register{
		ServerMap: make(map[ServerKey]ServerInfo),
	}
}

// write server to the center
func (r *Register) WriteServer(args *Server, reply *string) error {
	log.Println("Call WriteServer")
	r.ServerMap[args.ServerKey] = args.ServerInfo
	*reply = "Server registered successfully"
	return nil
}

// get server from the center
func (r *Register) GetServer(args *ServerKey, reply *ServerInfo) error {
	log.Println("Call GetServer")
	if _, ok := r.ServerMap[*args]; !ok {
		*reply = ServerInfo{}
		return fmt.Errorf("Server not found")
	}
	*reply = r.ServerMap[*args]
	return nil
}

var _ rpctypes.RpcRegister = (*Register)(nil)

func StartRegisterServer(port string) error {
	err := rpc.RegisterName(rpctypes.ServerNmae, NewRegister())
	if err != nil {
		return fmt.Errorf("Error registering RPC server: %v", err)
	}
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("Error starting listener: %v", err)
	}
	defer listener.Close()
	fmt.Printf("Register server is running on port %s...\n", port)
	http.Serve(listener, nil)
	return nil
}
