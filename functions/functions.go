package functions

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

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

func SignCertificate(subject, issuer *x509.Certificate, publicKey *rsa.PublicKey, privateKey *rsa.PrivateKey) []byte {

	// Sign x509 file using private key
	cert, err := x509.CreateCertificate(rand.Reader, subject, issuer, publicKey, privateKey)
	if err != nil {
		logrus.Errorf("Unable to create certificate: %v", err)
		return nil
	}
	return cert
}

// Create PEM encoding of cert for print
// Create a PEM file

// NAO PRECISA DA KEY
func CreatePEMfile(name string, cert []byte, priv *rsa.PrivateKey) []byte {
	
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
	
	
	/*pemfile, _ := os.Create(name)
	pem.Encode(pemfile, &blockPEM)
	pemfile.Close()*/

	return certPEM
}


func CreateKeyPEM (name string, priv *rsa.PrivateKey) []byte {

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
		
		
		/*pemfile, _ := os.Create(name)
		pem.Encode(pemfile, &blockPEM)
		pemfile.Close()*/
	
		return keyPEM
}


// helper function to create a cert template with a serial number and other required fields
func CertTemplate() (*x509.Certificate, error) {
    // generate a random serial number (a real cert authority would have some logic behind this)
    serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
    serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
    if err != nil {
        return nil, errors.New("failed to generate serial number: " + err.Error())
    }

    tmpl := x509.Certificate{
        SerialNumber:          serialNumber,
        Subject:               pkix.Name{Organization: []string{"Yhat, Inc."}},
        SignatureAlgorithm:    x509.SHA256WithRSA,
        NotBefore:             time.Now(),
        NotAfter:              time.Now().Add(time.Hour), // valid for an hour
        BasicConstraintsValid: true,
    }
    return &tmpl, nil
}

func CreateCert(template, parent *x509.Certificate, pub interface{}, parentPriv interface{}) (
    cert *x509.Certificate, certPEM []byte, err error) {

    certDER, err := x509.CreateCertificate(rand.Reader, template, parent, pub, parentPriv)
    if err != nil {
        return
    }
    // parse the resulting certificate so we can use it again
    cert, err = x509.ParseCertificate(certDER)
    if err != nil {
        return
    }
    // PEM encode the certificate (this is a standard TLS encoding)
    b := pem.Block{Type: "CERTIFICATE", Bytes: certDER}
    certPEM = pem.EncodeToMemory(&b)
    return
}