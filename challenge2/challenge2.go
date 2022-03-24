package challenge2

import (
	"crypto/rsa"
	"crypto/x509"

	//"crypto/x509/pkix"
	"math/big"
	//"time"

	"github.com/vvc-git/LabSec-Challenge.git/functions"
)

// Generate a new x509 Certificate for Intermediate CA
func CreateIntCert(rootDER []byte, keyToSign *rsa.PrivateKey) ([]byte, []byte, *rsa.PrivateKey) {

	// Generate private public key pair
	var privateKey = functions.CreateKey()
	var publicKey = privateKey.PublicKey

	// Creates x509 certificate with parameters related to CA root
	basicTmpl, _ := functions.CertTemplate()
	intCert := IntCertificate(basicTmpl)

	//  Parses certificate from the given ASN.1 DER data
	rootCA, err := x509.ParseCertificate(rootDER)
	if err != nil {
		panic("Failed to parse certificate:" + err.Error())
	}

	// Sign using root private key
	intermediateCASigned := functions.SignCertificate(intCert, rootCA, &publicKey, keyToSign)

	// Create a PEM file certificate (you can also print in terminal)
	intermediatePEM := functions.CreatePEMfile("2.intermediateCert.pem", intermediateCASigned, nil)

	return intermediateCASigned, intermediatePEM, privateKey

}

// Fill the required fields for Intermediate CA certificate
func IntCertificate(basicTmpl *x509.Certificate) *x509.Certificate {

	basicTmpl.SerialNumber = big.NewInt(2)
	basicTmpl.Subject.Organization = []string{"Intermdiate CA"}
	basicTmpl.Subject.Country = []string{"BR"}

	return basicTmpl
}
