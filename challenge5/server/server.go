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
		// Server certificate that will be use with client
		Certificates: []tls.Certificate{servTLSCert},
		// Recquiring clients to show his certificate
		//ClientAuth: tls.RequireAndVerifyClientCert,
		// Certificate authorities that server trust
		// ClientCAs:  certPool,
	}

	return s

}
