package server

import (
	//"crypto/rsa"
	//"crypto/rand"
	//"crypto/rsa"
	//"crypto/tls"
	"crypto/tls"
	"crypto/x509"
	//"encoding/pem"
	//"fmt"
	//"log"
	//"net"
	"net/http"
	"net/http/httptest"
	//"net/http/httputil"
	//"github.com/vvc-git/LabSec-Challenge.git/functions"
)


func Server(intCertPEM []byte ,servTLSCert tls.Certificate) (*httptest.Server) {

	/*// ----------- CHALLENGE 3: CREATE A CERTIFICATE FOR SERVER ------------------------
	// create a key-pair for the server
	servKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("generating random key: %v", err)
	}

	// create a template for the server
	servCertTmpl, err := functions.CertTemplate()
	if err != nil {
		log.Fatalf("creating cert template: %v", err)
	}
	servCertTmpl.KeyUsage = x509.KeyUsageDigitalSignature
	servCertTmpl.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}
	servCertTmpl.IPAddresses = []net.IP{net.ParseIP("127.0.0.1")}

	rootCert, err := x509.ParseCertificate(rootCertDER)
	if err != nil {
		panic("Failed to parse certificate:" + err.Error())
	}
	


	
	// create a certificate which wraps the server's public key, sign it with the root private key
	_, servCertPEM, err := functions.CreateCert(servCertTmpl, rootCert, &servKey.PublicKey, rootKey)
	if err != nil {
		log.Fatalf("error creating cert: %v", err)
		}	
	


	// provide the private key and the cert
	servKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(servKey),
	})

	// Create a TLS cert using the private key and certificate
	servTLSCert, err := tls.X509KeyPair(servCertPEM, servKeyPEM)
	if err != nil {
		log.Fatalf("invalid key pair: %v", err)
	}*/

	// ----------- CHALLENGE 3: CREATE A CERTIFICATE FOR SERVER ------------------------
	ok := func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("HI!")) }
	//s := httptest.NewUnstartedServer(http.HandlerFunc(ok))

	
	// create a pool of trusted certs
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(intCertPEM)


	// Configure the server to present the certficate we created
	// create another test server and use the certificate
	s := httptest.NewUnstartedServer(http.HandlerFunc(ok))
	s.TLS = &tls.Config{
		Certificates: []tls.Certificate{servTLSCert},
		// Getting the Server to Trust the Client
		// (Client have to show his certificate)
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs: certPool,
	}

	// configure a client to use trust those certificates


	s.StartTLS()
	return s
	//s.Close()
	
}






