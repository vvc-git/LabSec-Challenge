package main

import (
	"github.com/vvc-git/LabSec-Challenge.git/challenge1"
	"github.com/vvc-git/LabSec-Challenge.git/challenge2"
	"github.com/vvc-git/LabSec-Challenge.git/challenge3"
	"github.com/vvc-git/LabSec-Challenge.git/challenge5/clientCertGen"
)


func main() {

	// Challenge 1
	rootDER, keyToSignRoot := challenge1.SelfSignedCert()

	// Challenge 2
	intDER, intPEM, keyToSignInt := challenge2.CreateIntCert(rootDER, keyToSignRoot)

	// Challenge 3
	_ = challenge3.ServCertGenetor(intDER, intPEM, keyToSignInt)
	
	// Challenge 4

	//curl -Lv --cacert 3.servCert.pem  https://localhost:8443/hello

	// challenge 5
	_ = ClientCertGen.ClientCertGen(intDER, intPEM, keyToSignInt)

	//curl -Lv --cacert <path-to/3.servCert.pem> --cert <path-to/5.clientCert.pem> --key <path-to6.clientKey.pem>  https://localhost:8443/hello
	// Example - go to clientMTLS folder:
	//curl -Lv --cacert ../../3.servCert.pem --cert ../../5.clientCert.pem --key ../../6.clientKey.pem  https://localhost:8443/hello

}
