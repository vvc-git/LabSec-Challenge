package challenge2

import (
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"

	"github.com/vvc-git/LabSec-Challenge.git/functions"
)

func CreateIntermediateACCertificate(rootCAbytes []byte, keyToSign *rsa.PrivateKey)  ([]byte, *rsa.PrivateKey){

	// Generate private and public key
	var privateKey = functions.CreateKey()
	var publicKey = privateKey.PublicKey

	// Creates x509 certificate with parameters related to CA root
	intermediateCA := IntermediateAC()

	//  Parses certificate from the given ASN.1 DER data
	rootCA, err := x509.ParseCertificate(rootCAbytes)
	if err != nil {
		panic("Failed to parse certificate:" + err.Error())
	}


	// Use private key to sign itself
	intermediateCASigned := functions.SignCertificate(intermediateCA, rootCA, &publicKey, privateKey)

	// Create a PEM file certificate (you can also print in terminal)
	_ = functions.CreatePEMfile("Intermdiate_CA", intermediateCASigned, privateKey)

	return intermediateCASigned, privateKey

	
}

func IntermediateAC() *x509.Certificate {
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(2022),
		Subject: pkix.Name{
			Organization:  []string{""},
			Country:       []string{"BR"},
			Province:      []string{""},
			Locality:      []string{"São Paulo"},
			StreetAddress: []string{""},
			PostalCode:    []string{""},
		},
		//NotBefore:             time.Now(),
		//NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA: true,
		//ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		//KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		//BasicConstraintsValid: true,
	}
	return ca
}