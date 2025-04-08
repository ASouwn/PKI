// 代理包，提供接口简化rpc服务的调用

package utils

import (
	"fmt"
	"log"
	"net/rpc"
	"reflect"
	"strings"

	rpctypes "github.com/ASouwn/PKI/shared-rpc-types"
)

// 对于服务端的使用函数，将服务写入注册中心
func WriteServer(server *rpctypes.Server, registerAddress string) error {
	client, err := rpc.DialHTTP("tcp", registerAddress)
	if err != nil {
		return fmt.Errorf("failed to connect to register server: %v", err)
	}
	defer client.Close()

	var reply string
	err = client.Call(rpctypes.WriteServerMethod, server, &reply)
	if reply != "If server registered successfully: true" {
		return nil
	}
	return fmt.Errorf("failed to register server: %v", err)
}

// 对于客户端的使用函数
// 只需要将将要调用的服务(serverMethod)和参数(args)还有注册中心的地址(registerAddress)传入即可
func GetRedServer(serverMethod string, args interface{}, registerAddress string) (interface{}, error) {
	log.Printf("the type of args is %s\n", reflect.TypeOf(args))
	serverKey := strings.Split(serverMethod, ".")[0]

	// 连接注册中心
	// 通过注册中心获取服务端的地址和端口号
	clientRpc, err := rpc.DialHTTP("tcp", registerAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to register server: %v", err)
	}
	defer clientRpc.Close()
	var serverInfo rpctypes.ServerInfo
	err = clientRpc.Call(rpctypes.GetServerMethod, &rpctypes.ServerKey{
		ServerName: serverKey,
	}, &serverInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to get server info: %v", err)
	}
	log.Printf("serched ServerMethod(%s) addr is %s", serverMethod, serverInfo.ServerAddress+":"+serverInfo.Port)
	// 连接服务端
	// 通过服务端的地址和端口号连接服务端
	clientServer, err := rpc.DialHTTP("tcp", serverInfo.ServerAddress+":"+serverInfo.Port)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %v", err)
	}
	defer clientServer.Close()

	// todo: 需要通过serverMethod来断言args的类型，同时断言reply的类型，然后连接服务并返回relpy
	var reply interface{}
	err = clientServer.Call(serverMethod, args, &reply)
	if err != nil {
		return nil, fmt.Errorf("failed to call server method: %v", err)
	}
	return &reply, nil
}
