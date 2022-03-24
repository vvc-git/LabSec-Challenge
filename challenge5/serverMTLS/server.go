package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	//"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

const (
	intCertPath  = "../../2.intermediateCert.pem"
	servCertPath = "../../3.servCert.pem"
	servKeyPath  = "../../4.servKey.pem"
	port         = ":8443"
	add          = "https://127.0.0.1:8443/hello"
)

const usersResp = `
<html>
<body>
<p> LabSEC challenge
<p> This is the v call to this
</body>
</html>
`

var counter int

func main() {

	// Read intermediate certificate
	intCertPEM, err := os.ReadFile(intCertPath)
	if err != nil {
		logrus.Printf("Erro ao ler arquivo no cliente")
	}

	// Create a set of certificate (pool) of CA
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(intCertPEM)

	// Generate tls.certificate for server
	servCert, err := tls.LoadX509KeyPair(servCertPath, servKeyPath)
	if err != nil {
		logrus.Printf("nao conseguiu criar as chaves %v", err)
	}

	// Server settings
	config := &tls.Config{}
	// Certificate used to prove server's authenticity
	config.Certificates = []tls.Certificate{servCert}
	// Recquiring clients to show his certificate
	config.ClientAuth = tls.RequireAndVerifyClientCert
	// Certificate authorities that server trust
	config.ClientCAs = certPool

	// Server handler function
	http.HandleFunc("/hello", HelloServer)

	// Instantiate server
	logrus.Printf("About to listen on %s. Go to %s ", port, add)
	err = http.ListenAndServeTLS(port, servCertPath, servKeyPath, nil)
	if err != nil {
		logrus.Printf("nao conseguiu iniciar o servr %v", err)
	}

}

// Handler function
func HelloServer(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	// Show for server
	logrus.Printf("Client %v connected ", r.RemoteAddr)
	// Show for client
	counter++
	fmt.Fprintf(w, usersResp)
}