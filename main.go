package main

import (
	//"fmt"

	"fmt"

	"github.com/vvc-git/LabSec-Challenge.git/challenge1"
	"github.com/vvc-git/LabSec-Challenge.git/challenge2"
	"github.com/vvc-git/LabSec-Challenge.git/challenge4"
	//"github.com/vvc-git/LabSec-Challenge.git/functions"
	//"github.com/vvc-git/LabSec-Challenge.git/challenge5"
)




func main() {

	// Challenge 1
	rootPEM, keyToSignRoot := challenge1.SelfSignedCACertificate()
	//fmt.Printf(string(rootPEM), "\n")
	// Challenge 2
	IntermediateDER, IntermediatePEM, keyToSignIntermediate := challenge2.CreateIntermediateCACertificate(rootPEM, keyToSignRoot)
	//fmt.Printf(string(IntermediatePEM))
	// Challenge 3
	//serverPEM := challenge3.ServerCertificateGenetor(IntermediatePEM, keyToSignIntermediate)
	//fmt.Sprintln(serverPEM)

	// challenge 4.1
	// Testando com chave privada da intermediario 
	s := challenge4.Server(IntermediateDER, IntermediatePEM, keyToSignIntermediate)
	fmt.Print(s)
	//challenge4.Client(IntermediateDER, IntermediatePEM, keyToSignIntermediate, s)


	// Challenge 5
	// Client TLS certificate
	//clientPEM := challenge5.ClientCertificateGenetor(IntermediatePEM, keyToSignIntermediate)
	//fmt.Sprintln(clientPEM)

	// Challenge 4
	//challenge4.TLSserver(serverPEM)
	// Set up Client 
	// challenge5.TLSClient(clientPEM)
}



/*func server(rootCertDER []byte, rootCertPEM []byte, rootKey *rsa.PrivateKey) {

	// create a key-pair for the server
	servKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("generating random key: %v", err)
	}

	// create a template for the server
	servCertTmpl, err := CertTemplate()
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
	_, servCertPEM, err := CreateCert(servCertTmpl, rootCert, &servKey.PublicKey, rootKey)
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
	}

	ok := func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("HI!")) }
	//s := httptest.NewUnstartedServer(http.HandlerFunc(ok))

	
	// create a pool of trusted certs
	certPool := x509.NewCertPool()
	// rootCertPEM == 
	certPool.AppendCertsFromPEM(rootCertPEM)


	// Configure the server to present the certficate we created
	// create another test server and use the certificate
	s := httptest.NewUnstartedServer(http.HandlerFunc(ok))
	s.TLS = &tls.Config{
		Certificates: []tls.Certificate{servTLSCert},
		// Getting the Server to Trust the Client
		// Client have to show his certificate
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs: certPool,
	}

	// configure a client to use trust those certificates


	// create a key-pair for the client
	clientKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("generating random key: %v", err)
	}

	// create a template for the client
	clientCertTmpl, err := CertTemplate()
	if err != nil {
		log.Fatalf("creating cert template: %v", err)
	}
	clientCertTmpl.KeyUsage = x509.KeyUsageDigitalSignature
	clientCertTmpl.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}

	// the root cert signs the cert by again providing its private key
	_, clientCertPEM, err := CreateCert(clientCertTmpl, rootCert, &clientKey.PublicKey, rootKey)
	if err != nil {
		log.Fatalf("error creating cert: %v", err)
	}

	// encode and load the cert and private key for the client
	clientKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(clientKey),
	})
	
	clientTLSCert, err := tls.X509KeyPair(clientCertPEM, clientKeyPEM)
	if err != nil {
		log.Fatalf("invalid key pair: %v", err)
	}


	/*client := &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: certPool},
	},
	}

	authedClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      certPool,
				Certificates: []tls.Certificate{clientTLSCert},
			},
		},
	}

	s.StartTLS()
	resp, err := authedClient.Get(s.URL)
	s.Close()
	if err != nil {
		log.Fatalf("could not make GET request: %v", err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatalf("could not dump response: %v", err)
	}
	fmt.Printf("%s\n", dump)



}*/

