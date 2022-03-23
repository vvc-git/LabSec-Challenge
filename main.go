package main

import (
	//"fmt"

	"fmt"

	"github.com/vvc-git/LabSec-Challenge.git/challenge1"
	"github.com/vvc-git/LabSec-Challenge.git/challenge2"
	"github.com/vvc-git/LabSec-Challenge.git/challenge3"
	//"github.com/vvc-git/LabSec-Challenge.git/challenge4"
	//"github.com/vvc-git/LabSec-Challenge.git/challenge5"
)



func main() {

	// Challenge 1
	rootPEM, keyToSignRoot := challenge1.SelfSignedCACertificate()
	//fmt.Printf(string(rootPEM), "\n")
	// Challenge 2
	IntermediatePEM, keyToSignIntermediate := challenge2.CreateIntermediateCACertificate(rootPEM, keyToSignRoot)
	//fmt.Printf(string(IntermediatePEM))
	// Challenge 3
	serverPEM := challenge3.ServerCertificateGenetor(IntermediatePEM, keyToSignIntermediate)
	fmt.Sprintln(serverPEM)
	// Challenge 5
	// Client TLS certificate
	//clientPEM := challenge5.ClientCertificateGenetor(IntermediatePEM, keyToSignIntermediate)
	//fmt.Sprintln(clientPEM)

	// Challenge 4
	//challenge4.TLSserver(serverPEM)
	// Set up Client 
	// challenge5.TLSClient(clientPEM)
}

