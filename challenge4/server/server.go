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

func main() {

	// Regastar os certificados que o cliente reconhece
	intCertPEM, err := os.ReadFile("../../2.intermediateCert.pem")
	if err != nil {
		logrus.Printf("Erro ao ler arquivo no cliente")
	}

	// fazer o pool com o intCert.pem
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(intCertPEM)

	// criar um certificado tls para as chaves fornecida
	servCert, err := tls.LoadX509KeyPair("../../3.servCert.pem", "../../4.servKey.pem")
	if err != nil {
		logrus.Printf("nao conseguiu criar as chaves %v", err)
	}

	// configura o server
	config := &tls.Config{}
	// Certificado do servidor para mostrar para o cliente
	config.Certificates = []tls.Certificate{servCert}
	// Recquiring clients to show his certificate
	config.ClientAuth = tls.RequireAndVerifyClientCert 
	// Certificate authorities that server trust
	config.ClientCAs = certPool

	// Inicia o server
	// One can use generate_cert.go in crypto/tls to generate cert.pem and key.pem.
	http.HandleFunc("/hello", HelloServer)

	logrus.Printf("About to listen on 8443. Go to https://127.0.0.1:8443/hello")
	err = http.ListenAndServeTLS(":8443", "../../3.servCert.pem", "../../4.servKey.pem", nil)
	if err != nil {
		logrus.Printf("nao conseguiu iniciar o servr %v", err)
	}

    //http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

	logrus.Printf("Client %v connected ", r.RemoteAddr)
    fmt.Fprintf(w, "LabSEC challenge!")
}