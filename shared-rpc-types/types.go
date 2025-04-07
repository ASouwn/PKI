package sharedrpctypes

var (
	// rpc
	GetServerMethod   = "RpcRegister.GetServer"   // used by client to get server
	WriteServerMethod = "RpcRegister.WriteServer" // used by client to write server
	// ra
	RAHandleCSRMethod = "RA.HandleCSR" // used by client to handle CSR

	RPCServerKey = ServerKey{ServerName: "RpcRegister"} // used by server to regist
	RAServerKey  = ServerKey{ServerName: "RA"}          // used by server to regist

)

type ArgsReply struct {
	Args  interface{}
	Reply interface{}
}
