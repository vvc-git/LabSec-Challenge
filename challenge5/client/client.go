package client

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"

	"fmt"
	"log"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"

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
func ClientMTLS(intCertPEM []byte, s *httptest.Server, clientTLSCert tls.Certificate) *http.Client {

	// create a set of trusted certs
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(intCertPEM)

	authedClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				// Certificate authorities that clients trust
				RootCAs: certPool,
				// Client certificate which will be use with server
				Certificates: []tls.Certificate{clientTLSCert},
			},
		},
	}

	return authedClient

}

func StartClientMTLS(authedClient *http.Client, s *httptest.Server) {

	resp, err := authedClient.Get(s.URL)
	if err != nil {
		log.Fatalf("could not make GET request: %v", err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatalf("could not dump response: %v", err)
	}
	fmt.Printf("%s\n", dump)

}
