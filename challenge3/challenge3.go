package challenge3

import (
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"

	"github.com/vvc-git/LabSec-Challenge.git/functions"
)

func ServerCertificateGenetor (intermediateCAbytes []byte, keyToSign *rsa.PrivateKey) ([]byte, *rsa.PrivateKey) {

	// Generate private public key pair
	var privateKey = functions.CreateKey()
	var publicKey  = privateKey.PublicKey

	// Creates x509 certificate with parameters related to CA root
	serverCert := Server_Certifcate()

	//  Parses certificate from the given ASN.1 DER data
	intermediateCA, err := x509.ParseCertificate(intermediateCAbytes)
	if err != nil {
		panic("Failed to parse certificate:" + err.Error())
	}

	// Sign using intermediate private key
	intermediateCASigned := functions.SignCertificate(serverCert, intermediateCA, &publicKey, keyToSign)

	// Create a PEM file certificate (you can also print in terminal)
	_ = functions.CreatePEMfile("Server_Certifcate", intermediateCASigned, privateKey)

	return intermediateCASigned, privateKey

	
}

func Server_Certifcate() *x509.Certificate {
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(3),
		Issuer: pkix.Name{
			CommonName: "SecurityLab.com",
		},
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
