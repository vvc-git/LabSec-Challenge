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

// Make a request for a Server
/*func ClientMTLS(intCertPEM []byte, s *httptest.Server, clientTLSCert tls.Certificate) *http.Client {

	// create a set of trusted certs
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(intCertPEM)

	authedClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				// Certificate authorities that clients trust
				RootCAs:      certPool,
				// Client certificate which will be use with server
				Certificates: []tls.Certificate{clientTLSCert},
			},
		},
	}

	return authedClient

}*/

func clientDIAL() {

	// create a set of trusted certs
	//certPool := x509.NewCertPool()
	///certPool.AppendCertsFromPEM(intCertPEM)


	conf := &tls.Config{
		//InsecureSkipVerify: true,
		//RootCAs:      certPool,
   }

   conn, err := tls.Dial("tcp", "127.0.0.1:8443", conf)
   if err != nil {
	   logrus.Println(err)
	   return
   }
   conn.Handshake()
}

func main() {

	// ler o intCer.pem
	intCertPEM, err := os.ReadFile("../../2.intermediateCert.pem")
	if err != nil {
		logrus.Printf("Erro ao ler arquivo no cliente")
	}

	// fazer o pool com o intCert.pem
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(intCertPEM)


	// Ler o certicado do cliente do formato pem
	clientCert, err := os.ReadFile("../../5.clientCert.pem")
	if err != nil {
		logrus.Printf("Erro ao ler arquivo no cliente")
	}

	// Ler a key do cliente do formato pem
	clientKey, err := os.ReadFile("../../6.clientKey.pem")
	if err != nil {
		logrus.Printf("Erro ao ler arquivo no cliente")
	}

	// ler a chave e o certificado do cliente para criar um tls.certificate
	clientTLSCert, err := tls.X509KeyPair(clientCert, clientKey)

	// configuarar tls
	confTls := &tls.Config{}
	// Certificado que o cliente confia de uma AC
	confTls.RootCAs = certPool
	// Certificado que vai apresentar para o servidor
	confTls.Certificates = []tls.Certificate{clientTLSCert}

	// configurar http Trasport
	confTransp := &http.Transport{}
	confTransp.TLSClientConfig = confTls


	// configurar o client http
	client := &http.Client{}
	client.Transport = confTransp

	// client get
	response, _ := client.Get("https://127.0.0.1:8443/hello")
	defer response.Body.Close()

	// tratando erros e impriminda na tela
	html, err := io.ReadAll(response.Body)
	if err != nil {
	  logrus.Fatal(err)
	}
	fmt.Printf("%v\n", response.Status)
	fmt.Printf(string(html))

}