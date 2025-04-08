package main

import pkiserver "github.com/ASouwn/PKI/pki-server"

func main() {
	pkiserver.StarRaServer("8081", "localhost:3001")
}
