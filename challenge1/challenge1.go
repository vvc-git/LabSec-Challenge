package challenge1

import (
	"crypto/rsa"
	"crypto/x509"
	"math/big"

	"github.com/vvc-git/LabSec-Challenge.git/functions"
)

// Generate a new x509 Certificate for root CA
func SelfSignedCert() ([]byte, *rsa.PrivateKey) {
	// Generate private and public key
	var privateKey = functions.CreateKey()
	var publicKey = privateKey.PublicKey

	// Creates x509 certificate with parameters related to CA root
	basicTmpl, _ := functions.CertTemplate()
	caRoot := CreateRootCACert(basicTmpl)

	// Use private key to sign itself
	caSigned := functions.SignCertificate(caRoot, caRoot, &publicKey, privateKey)

	_ = functions.CreatePEMfile("1.rootCert.pem", caSigned, nil)

	return caSigned, privateKey
}

// Fill the required fields for Intermediate CA certificate
func CreateRootCACert(basicTmpl *x509.Certificate) *x509.Certificate {

	basicTmpl.SerialNumber = big.NewInt(1)
	basicTmpl.Subject.Organization = []string{"root CA"}
	basicTmpl.Subject.Country = []string{"BR"}

	return basicTmpl
}
