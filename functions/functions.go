package functions

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"github.com/sirupsen/logrus"
)


func CreateKey() *rsa.PrivateKey {

	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("Unable to create a key: %v", err.Error())
		os.Exit(1)
	}

	return priv
}

func SignCertificate(Subject, issuer *x509.Certificate, publicKey *rsa.PublicKey, privateKey *rsa.PrivateKey) []byte {

	// Sign x509 file using private key
	cert, err := x509.CreateCertificate(rand.Reader, issuer, issuer, publicKey, privateKey)
	if err != nil {
		logrus.Errorf("Unable to create certificate: %v", err)
		return nil
	}
	return cert
}

// Create PEM encoding of cert for print
// Create a PEM file
func CreatePEMfile(name string, cert []byte, priv *rsa.PrivateKey) []byte {

	// block to be encoded
	blockPEM := pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert}

	certPEM := pem.EncodeToMemory(&blockPEM)

	// Create plain text PEM file.
	pemfile, _ := os.Create(name)
	pem.Encode(pemfile, &blockPEM)
	pemfile.Close()

	return certPEM
}
