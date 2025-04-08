package sharedrpctypes

var (
	// rpc
	GetServerMethod   = "RpcRegister.GetServer"   // used by client to get server
	WriteServerMethod = "RpcRegister.WriteServer" // used by client to write server
	// ra
	RAHandleCSRMethod = "RA.HandleCSR" // used by client to handle CSR
	// ca
	CAHandleCSRMethod = "CA.HandleCSR" // used by client to CA to handle CSR

	RPCServerKey = ServerKey{ServerName: "RpcRegister"} // used by server to regist
	RAServerKey  = ServerKey{ServerName: "RA"}          // used by server to regist
	CAServerKey  = ServerKey{ServerName: "CA"}
)

type ArgsReply struct {
	Args  interface{}
	Reply interface{}
}
