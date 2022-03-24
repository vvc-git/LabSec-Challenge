package clientSeUp

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"

	//"fmt"
	"log"
	"math/big"
	"net/http"
	"net/http/httptest"
	//"net/http/httputil"

	"github.com/vvc-git/LabSec-Challenge.git/functions"
)

// Fill the required fields for Client certificate
func Client_Certifcate(basicTmpl *x509.Certificate) *x509.Certificate {
	basicTmpl.SerialNumber = big.NewInt(4)
	basicTmpl.Subject.Organization = []string{"Client"}
	basicTmpl.KeyUsage = x509.KeyUsageDigitalSignature
	basicTmpl.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}

	return basicTmpl
}

// Generate a x509 Client Certificate
func ClientCertificateGenetor(intDER, intPEM []byte, keyToSign *rsa.PrivateKey) tls.Certificate {

	// Generate private public key pair
	var privateKey = functions.CreateKey()
	var publicKey = privateKey.PublicKey

	// Creates x509 certificate with parameters related to CA Intermediate
	basicTmpl, _ := functions.CertTemplate()
	clientCert := Client_Certifcate(basicTmpl)

	//  PEM to x509.certificate
	intCert, err := x509.ParseCertificate(intDER)
	if err != nil {
		panic("Failed to parse certificate:" + err.Error())
	}

	// Sign using intermediate private key
	clientCASigned := functions.SignCertificate(clientCert, intCert, &publicKey, keyToSign)

	// Create a PEM file certificate
	clientCertPEM := functions.CreatePEMfile("5.clientCert.pem", clientCASigned, nil)
	clientKeyPEM := functions.CreatePEMfile("6.clientKey.pem", nil, privateKey)

	// Create a TLS cert using the private key and certificate
	clientTLSCert, err := tls.X509KeyPair(clientCertPEM, clientKeyPEM)
	if err != nil {
		log.Fatalf("invalid key pair: %v", err)
	}

	return clientTLSCert

}

// Make a request for a Server
func Client(intCertPEM []byte, s *httptest.Server) *http.Client {

	client := &http.Client{}
	
	return client

}
