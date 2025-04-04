package sharedrpctypes

var (
	ServerNmae        = "RpcRegister"             // used by server to regist
	GetServerMethod   = "RpcRegister.GetServer"   // used by client to get server
	WriteServerMethod = "RpcRegister.WriteServer" // used by client to write server
)

// This struct is used to register the server information
// will get the server information from the client
// network: tcp or udp, others the server provides
// to use the addr: serverAddress+":"+port
type ServerInfo struct {
	Network       string
	ServerAddress string
	Port          string
}

// This struct is used to register the server information
// function in regist server: func GetServer(args *ServerKey, reply *ServerInfo)
// the reply is the ServerInfo struct
type ServerKey struct {
	ServerName string
}

// This struct is used to register the server information
// function in regist server: func WriteServer(args *Server, reply *string)
// the reply is the string "Server registered successfully"
type Server struct {
	ServerInfo ServerInfo
	ServerKey  ServerKey
}

// 请通过RpcRegister.ServiceName来调用接口
type RpcRegister interface {
	WriteServer(args *Server, reply *string) error
	GetServer(args *ServerKey, reply *ServerInfo) error
}
