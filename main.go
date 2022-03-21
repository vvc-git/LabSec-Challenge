package main

import (
	"fmt"

	"github.com/vvc-git/LabSec-Challenge.git/challenge1"
	"github.com/vvc-git/LabSec-Challenge.git/challenge2"
	"github.com/vvc-git/LabSec-Challenge.git/challenge3"
)



func main() {

	// Challenge 1
	root, keyToSignRoot := challenge1.SelfSignedCACertificate()
	fmt.Printf(string(root), "\n")
	// Challenge 2
	Intermediate, keyToSignIntermediate := challenge2.CreateIntermediateCACertificate(root, keyToSignRoot)
	fmt.Printf(string(Intermediate))
	// Challenge 3
	challenge3.ServerCertificateGenetor(Intermediate, keyToSignIntermediate)
}

