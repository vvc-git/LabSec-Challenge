package serverSetUp

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	//"net/http/httptest"
)

// Set up server MTLS
func serverSetUp(intCertPEM []byte, servTLSCert tls.Certificate) *http.Server {

	//lab := func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("LabSEC Challenge! 2")) }

	// create a set of trusted certs
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(intCertPEM)

	// Set the server with certificate generated
	s := &http.Server{}

	//s.TLSConfig = &tls.Config{}
	s.TLSConfig.Certificates = []tls.Certificate{servTLSCert}
	s.TLSConfig.ClientAuth = tls.RequireAndVerifyClientCert

	return s

}
