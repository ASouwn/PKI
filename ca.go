package main

import pkiserver "github.com/ASouwn/PKI/pki-server"

func main() {
	pkiserver.StartCAServer("8082", "localhost:3001")
}
