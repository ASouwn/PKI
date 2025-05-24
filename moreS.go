package main

import (
	rpctypes "github.com/ASouwn/PKI/shared-rpc-types"
	"github.com/ASouwn/PKI/utils"
)

func main() {
	servers := []rpctypes.Server{
		{
			ServerInfo: rpctypes.ServerInfo{
				Network:       "tcp",
				ServerAddress: "localhost",
				Port:          "10086",
			},
			ServerKey: rpctypes.ServerKey{
				ServerName: "CA4pay",
			},
		},
		{
			ServerInfo: rpctypes.ServerInfo{
				Network:       "tcp",
				ServerAddress: "localhost",
				Port:          "10087",
			},
			ServerKey: rpctypes.ServerKey{
				ServerName: "CA4line",
			},
		},
		{
			ServerInfo: rpctypes.ServerInfo{
				Network:       "tcp",
				ServerAddress: "localhost",
				Port:          "10088",
			},
			ServerKey: rpctypes.ServerKey{
				ServerName: "CA4user",
			},
		},
	}
	for _, s := range servers {
		utils.WriteServer(&s, "localhost:3001")
	}
}
