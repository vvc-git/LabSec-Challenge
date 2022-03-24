package main

import (
	"github.com/vvc-git/LabSec-Challenge.git/challenge1"
	"github.com/vvc-git/LabSec-Challenge.git/challenge2"
	"github.com/vvc-git/LabSec-Challenge.git/challenge3"
	"github.com/vvc-git/LabSec-Challenge.git/challenge5/client"
	"github.com/vvc-git/LabSec-Challenge.git/challenge5/server"
)

func main() {

	// Challenge 1
	rootDER, keyToSignRoot := challenge1.SelfSignedCert()

	// Challenge 2
	intDER, intPEM, keyToSignInt := challenge2.CreateIntCert(rootDER, keyToSignRoot)

	// Challenge 3
	servTLSCert := challenge3.ServerCertificateGenetor(intDER, intPEM, keyToSignInt)

	// Challenge 4

	// challenge 5
	s := server.ServerMTLS(intPEM, servTLSCert)

	clientTLSCert := client.ClientCertificateGenetor(intDER, intPEM, keyToSignInt)
	client.StartClient(intPEM, s, clientTLSCert)

}
