package challenge1

import (
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"github.com/vvc-git/LabSec-Challenge.git/functions"
)

func SelfSignedCACertificate() ([]byte, *rsa.PrivateKey) {
	// Generate private and public key
	var privateKey = functions.CreateKey()
	var publicKey = privateKey.PublicKey

	// Creates x509 certificate with parameters related to CA root
	caRoot := CreateRootCACert()

	// Use private key to sign itself
	caSigned := functions.SignCertificate(caRoot, caRoot, &publicKey, privateKey)

	_ = functions.CreatePEMfile("CA_Certificate", caSigned, privateKey)
	//CreateDERfile(cert)

	return caSigned, privateKey 
}

func CreateRootCACert() *x509.Certificate {
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

