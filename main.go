package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	//"io/ioutil"

	//"crypto/x509/pkix"

	"fmt"

	//"io/ioutil"
	//"log"
	"math/big"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	rng = rand.Reader
	// Emissor
	issuerPrivKey = CreateKey()
	issuer        *x509.Certificate
)

const (
	keySize = 1024
)

func main() {

	// Challenge 1
	SelfSignedCertificate()

}

func SelfSignedCertificate() error {
	// O que é essa chaves (??)
	caRoot := CreateCAroot()
	_, err := x509.CreateCertificate(rng, caRoot, caRoot, &issuerPrivKey.PublicKey, issuerPrivKey)
	if err != nil {
		logrus.Errorf("Unable to create certificate: %v", err)
		return err
	}

	pemX509File, err := os.Create("X509_CA.pem")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var pemX509block = &pem.Block{
		Type:  "CA ROOT CERTIFICATE",
		Bytes: x509.MarshalPKCS1PrivateKey(issuerPrivKey),
	}

	// Writes the PEM encoding of pemX509 to pemX509File.
	err = pem.Encode(pemX509File, pemX509block)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pemX509File.Close()



	/*err = ioutil.WriteFile("CA ROOT", cer, 0644)
	if err != nil {
		logrus.Errorf("Unable to create certificate: %v", err)
		return err
	}
	*/
	return nil
}

func CreateKey() *rsa.PrivateKey {
	rnd := rand.Reader
	priv, err := rsa.GenerateKey(rnd, keySize)
	if err != nil {
		fmt.Printf("Unable to create a key: %v", err.Error())
		os.Exit(1)
	}

	return priv
}

/*func CreateCAroot() *x509.Certificate {
	// set up our CA certificate
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(2022),
		IsCA:                  true,
	}

	return ca

}*/

func CreateCAroot() *x509.Certificate {
	// set up our CA certificate
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(2019),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}
	return ca

}
