package challenge5

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"

	"github.com/vvc-git/LabSec-Challenge.git/functions"
)

func Client(intermediateCA []byte) {

	// Read the key pair to create certificate
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatal(err)
	}

	/*// Create a CA certificate pool and add cert.pem to it
	caCert, err := ioutil.ReadFile("cert2.pem")
	if err != nil {
		log.Fatal(err)
	}*/
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(intermediateCA)

	// Create a HTTPS client and supply the created CA pool and certificate
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	// Request /hello via the created HTTPS client over port 8443 via GET
	r, err := client.Get("https://localhost:8443/hello")
	if err != nil {
		log.Fatal(err)
	}

	// Read the response body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response body to stdout
	fmt.Printf("%s\n", body)
}

func ClientCertificateGenetor (intermediateCAbytes []byte, keyToSign *rsa.PrivateKey) ([]byte) {

	// Generate private public key pair
	var privateKey = functions.CreateKey()
	var publicKey  = privateKey.PublicKey

	// Creates x509 certificate with parameters related to CA root
	serverCert := Client_Certifcate()

	//  Parses certificate from the given ASN.1 DER data
	intermediateCA, err := x509.ParseCertificate(intermediateCAbytes)
	if err != nil {
		panic("Failed to parse certificate:" + err.Error())
	}

	// Sign using intermediate private key
	intermediateCASigned := functions.SignCertificate(serverCert, intermediateCA, &publicKey, keyToSign)

	// Create a PEM file certificate (It's posbile to print in terminal)
	_ = functions.CreatePEMfile("cert.pem", intermediateCASigned, privateKey)
	_ = functions.CreateKeyPEM("key.pem", privateKey)

	return intermediateCASigned

}

func Client_Certifcate() *x509.Certificate {
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(3),
		Issuer: pkix.Name{
			CommonName: "",
		},
		Subject: pkix.Name{
			Organization:  []string{""},
			Country:       []string{"BR"},
			Province:      []string{""},
			Locality:      []string{"São Paulo"},
			StreetAddress: []string{""},
			PostalCode:    []string{""},
			CommonName:    "localhost",

		},
		//NotBefore:             time.Now(),
		//NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA: true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}
	return ca
}