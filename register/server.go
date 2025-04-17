package register

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"sync"

	rpctypes "github.com/ASouwn/PKI/shared-rpc-types"
)

type ServerKey = rpctypes.ServerKey
type ServerInfo = rpctypes.ServerInfo
type Server = rpctypes.Server

type Register struct {
	mu        sync.RWMutex
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
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ServerMap[args.ServerKey] = args.ServerInfo
	log.Printf("%s is registed in this addr: %s", args.ServerKey.ServerName, args.ServerInfo.ServerAddress+":"+args.ServerInfo.Port)
	*reply = "If server registered successfully: true"
	return nil
}

// get server from the center
func (r *Register) GetServer(args *ServerKey, reply *ServerInfo) error {
	log.Printf("Call GetServer: trying to call %s\n", args.ServerName)
	r.mu.RLock()
	defer r.mu.RUnlock()
	info, ok := r.ServerMap[*args]
	if !ok {
		return fmt.Errorf("Server not found")
	}
	*reply = info
	return nil
}

var _ rpctypes.RpcRegister = (*Register)(nil)

func StartRegisterServer(port string) error {
	register := NewRegister()
	http.HandleFunc("/servers", func(w http.ResponseWriter, r *http.Request) {
		register.mu.RLock()
		defer register.mu.RUnlock()

		w.Header().Set("Content-Type", "application/json")
		serializableMap := make(map[string]ServerInfo)
		for key, value := range register.ServerMap {
			serializableMap[key.ServerName] = value
		}
		if err := json.NewEncoder(w).Encode(serializableMap); err != nil {
			http.Error(w, "Failed to encode server map", http.StatusInternalServerError)
		}
	})
	err := rpc.RegisterName(rpctypes.RPCServerKey.ServerName, register)
	if err != nil {
		return fmt.Errorf("error registering RPC server: %v", err)
	}
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("error starting listener: %v", err)
	}
	defer listener.Close()
	fmt.Printf("Register server is running on port %s...\n", port)
	http.Serve(listener, nil)
	return nil
}
