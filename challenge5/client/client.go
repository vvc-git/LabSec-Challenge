package client

import (
	//"crypto/rsa"
	//"crypto/rand"
	//"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"

	//"encoding/pem"
	"fmt"
	"log"

	//"net"
	"net/http"
	//"net/http/httptest"
	"net/http/httptest"
	"net/http/httputil"

	"github.com/vvc-git/LabSec-Challenge.git/functions"
)

//clientTLSCert tls.Certificate
/*func Client(rootCertDER []byte, rootKey *rsa.PrivateKey, intCertPEM []byte, s *httptest.Server) {

	// create a key-pair for the client
	clientKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("generating random key: %v", err)
	}

	// create a template for the client
	clientCertTmpl, err := functions.CertTemplate()
	if err != nil {
		log.Fatalf("creating cert template: %v", err)
	}
	clientCertTmpl.KeyUsage = x509.KeyUsageDigitalSignature
	clientCertTmpl.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}

	rootCert, err := x509.ParseCertificate(rootCertDER)
	if err != nil {
		panic("Failed to parse certificate:" + err.Error())
	}

	// the root cert signs the cert by again providing its private key
	_, clientCertPEM, err := functions.CreateCert(clientCertTmpl, rootCert, &clientKey.PublicKey, rootKey)
	if err != nil {
		log.Fatalf("error creating cert: %v", err)
	}

	// encode and load the cert and private key for the client
	clientKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(clientKey),
	})
	
	clientTLSCert, err := tls.X509KeyPair(clientCertPEM, clientKeyPEM)
	if err != nil {
		log.Fatalf("invalid key pair: %v", err)
	}

	// sem o cliente se autenticando
	/*client := &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: certPool},
	},
	}

	// create a pool of trusted certs
	certPool := x509.NewCertPool()
	// rootCertPEM == 
	certPool.AppendCertsFromPEM(intCertPEM)


	authedClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      certPool,
				Certificates: []tls.Certificate{clientTLSCert},
			},
		},
	}

	resp, err := authedClient.Get(s.URL)
	if err != nil {
		log.Fatalf("could not make GET request: %v", err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatalf("could not dump response: %v", err)
	}
	fmt.Printf("%s\n", dump)

}*/

func StartClient (intCertPEM []byte, s *httptest.Server,  clientTLSCert tls.Certificate) {


	// create a pool of trusted certs
	certPool := x509.NewCertPool()
	// rootCertPEM == 
	certPool.AppendCertsFromPEM(intCertPEM)

	authedClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      certPool,
				Certificates: []tls.Certificate{clientTLSCert},
			},
		},
	}

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


func ClientCertificateGenetor (intDER, intPEM []byte, keyToSign *rsa.PrivateKey) (tls.Certificate) {

	// Generate private public key pair
	var privateKey = functions.CreateKey()
	var publicKey  = privateKey.PublicKey

	// Creates x509 certificate with parameters related to CA Intermediate
	basicTmpl, _ := functions.CertTemplate()		
	clientCert      := Client_Certifcate(basicTmpl)

	//  PEM to x509.certificate
	intCert, err := x509.ParseCertificate(intDER)
	if err != nil {
		panic("Failed to parse certificate:" + err.Error())
	}

	// Sign using intermediate private key
	clientCASigned := functions.SignCertificate(clientCert, intCert, &publicKey, keyToSign)


	// Create a PEM file certificate (It's posbile to print in terminal)
	clientCertPEM := functions.CreatePEMfile("cert.pem", clientCASigned, privateKey)
	clientKeyPEM  := functions.CreateKeyPEM("key.pem", privateKey)


	// Create a TLS cert using the private key and certificate
	clientTLSCert, err := tls.X509KeyPair(clientCertPEM, clientKeyPEM)
	if err != nil {
		log.Fatalf("invalid key pair: %v", err)
	}

	return clientTLSCert

	
}

func Client_Certifcate(basicTmpl *x509.Certificate) *x509.Certificate {

	basicTmpl.KeyUsage = x509.KeyUsageDigitalSignature
	basicTmpl.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}


	return basicTmpl
}