package functions

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// Create a private key
func CreateKey() *rsa.PrivateKey {

	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		logrus.Errorf("Unable to create a key: %v", err)
		return nil
	}

	return priv
}

// Create a certificate template with required fields
func CertTemplate() (*x509.Certificate, error) {

	tmpl := x509.Certificate{
		SignatureAlgorithm:    x509.SHA256WithRSA,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0), // valid for an hour
		IsCA:                  true,
		BasicConstraintsValid: true,
	}
	return &tmpl, nil
}

// Sign x509 file using Issuer private key
func SignCertificate(sub, iss *x509.Certificate, pub *rsa.PublicKey, issuerPriv *rsa.PrivateKey) []byte {

	cert, err := x509.CreateCertificate(rand.Reader, sub, iss, pub, issuerPriv)
	if err != nil {
		logrus.Errorf("Unable to create certificate: %v", err)
		return nil
	}

	return cert
}

// Create key and certificates PEM file
func CreatePEMfile(name string, cert []byte, priv *rsa.PrivateKey) []byte {

	if priv == nil {
		return CreateCertPEM(name, cert)

	} else {
		return CreateKeyPEM(name, priv)
	}

}

// Create certificates in PEM file
func CreateCertPEM(name string, cert []byte) []byte {

	// block to be encoded
	blockPEM := pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert}

	certPEM := pem.EncodeToMemory(&blockPEM)

	// Create plain text PEM file.
	err := os.WriteFile(name, certPEM, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return certPEM
}

// Create private keys PEM file
func CreateKeyPEM(name string, priv *rsa.PrivateKey) []byte {

	// block to be encoded
	blockPEM := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(priv)}

	keyPEM := pem.EncodeToMemory(&blockPEM)

	// Create plain text PEM file.
	err := os.WriteFile(name, keyPEM, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return keyPEM
}
