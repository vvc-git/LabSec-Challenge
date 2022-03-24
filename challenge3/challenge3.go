package challenge3

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"log"
	"math/big"
	"net"

	"github.com/vvc-git/LabSec-Challenge.git/functions"
)

// Generate a x509 Server Certificate
func ServCertGenetor(intDER, intPEM []byte, keyToSign *rsa.PrivateKey) tls.Certificate {

	// Generate private public key pair
	var privateKey = functions.CreateKey()
	var publicKey = privateKey.PublicKey

	// Creates x509 certificate with parameters related to CA Intermediate
	basicTmpl, _ := functions.CertTemplate()
	serverCert := Server_Certifcate(basicTmpl)

	//  PEM to x509.certificate
	intCert, err := x509.ParseCertificate(intDER)
	if err != nil {
		panic("Failed to parse certificate:" + err.Error())
	}

	// Sign using intermediate private key
	ServerCASigned := functions.SignCertificate(serverCert, intCert, &publicKey, keyToSign)

	// Create a PEM file certificate (It's posbile to print in terminal)
	servCertPEM := functions.CreatePEMfile("3.servCert.pem", ServerCASigned, nil)
	servKeyPEM := functions.CreatePEMfile("4.servKey.pem", nil, privateKey)

	// Create a TLS cert using the private key and certificate
	servTLSCert, err := tls.X509KeyPair(servCertPEM, servKeyPEM)
	if err != nil {
		log.Fatalf("invalid key pair: %v", err)
	}

	return servTLSCert

}

// Fill the required fields for Server certificate
func Server_Certifcate(basicTmpl *x509.Certificate) *x509.Certificate {
	basicTmpl.SerialNumber = big.NewInt(3)
	basicTmpl.Subject.Organization = []string{"Server"}
	basicTmpl.DNSNames = []string{"localhost"}
	basicTmpl.KeyUsage = x509.KeyUsageDigitalSignature
	basicTmpl.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}
	// Using localhost
	basicTmpl.IPAddresses = []net.IP{net.ParseIP("127.0.0.1")}

	return basicTmpl
}
