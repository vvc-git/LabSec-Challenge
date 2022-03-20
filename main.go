package main

import (
	"fmt"

	"github.com/vvc-git/LabSec-Challenge.git/challenge1"
	"github.com/vvc-git/LabSec-Challenge.git/challenge2"
)



func main() {

	// Challenge 1
	root, keyToSignRoot := challenge1.SelfSignedCACertificate()
	//fmt.Printf(string(pem))
	// Challenge 2
	pemIntermediate, _ := challenge2.CreateIntermediateACCertificate(root, keyToSignRoot)
	fmt.Printf(string(pemIntermediate))
}

