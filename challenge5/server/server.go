package server

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"net/http/httptest"
)

// Set up server MTLS
func ServerMTLS(intCertPEM []byte, servTLSCert tls.Certificate) *httptest.Server {

	lab := func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("LabSEC Challenge! ")) }

	// create a set of trusted certs
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(intCertPEM)

	// Set the server with certificate generated
	s := httptest.NewUnstartedServer(http.HandlerFunc(lab))

	s.TLS = &tls.Config{
		Certificates: []tls.Certificate{servTLSCert},
		// Client have to show his certificate
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  certPool,
	}

	return s

}
