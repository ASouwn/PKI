package main

import (
	"github.com/ASouwn/PKI/register"
)

func main() {
	// Start the PKI server on port 1234
	// pkiServer.StartPKIServer("1234")
	register.StartRegisterServer("3001")
}
