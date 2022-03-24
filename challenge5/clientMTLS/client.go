package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"

	//"crypto/x509"

	"github.com/sirupsen/logrus"
)

func main() {

	// Read Certificate and keys
	intCertPEM := ReadAndCheck("../../2.intermediateCert.pem")
	clientCert := ReadAndCheck("../../5.clientCert.pem")
	clientKey := ReadAndCheck("../../6.clientKey.pem")

	// Create a set of certificate (pool) of CA
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(intCertPEM)

	// Generate tls.certificate for client
	clientTLSCert, err := tls.X509KeyPair(clientCert, clientKey)

	// Set required fields for tls client
	client := TLSClientTemp(clientTLSCert, certPool)

	// Client get
	response, err := client.Get("https://127.0.0.1:8443/hello")
	if err != nil {
		logrus.Printf("Response error!")
	}
	defer response.Body.Close()

	// Response html  
	html, err := io.ReadAll(response.Body)
	if err != nil {
		logrus.Fatal(err)
	}
	// Check if the handshake has completed
	if response.TLS.HandshakeComplete != false {
		fmt.Println("The handshake has concluded!")
	}

	fmt.Printf(string(html))

}


func TLSClientTemp(clientTLSCert tls.Certificate, certPool *x509.CertPool) *http.Client {

	confTls := &tls.Config{}
	// Certificate authorities that clients trust
	confTls.RootCAs = certPool
	// Client certificate that will be show for server
	confTls.Certificates = []tls.Certificate{clientTLSCert}

	// Setting http Trasport
	confTransp := &http.Transport{}
	confTransp.TLSClientConfig = confTls

	// Setting client http
	client := &http.Client{}
	client.Transport = confTransp

	return client

}

func ReadAndCheck(name string) []byte {

	PEMFile, err := os.ReadFile(name)
	if err != nil {
		logrus.Printf("Erro ao ler arquivo no cliente")
	}

	return PEMFile
}
