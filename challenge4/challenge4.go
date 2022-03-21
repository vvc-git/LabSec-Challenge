package challenge4

/*import (
	"crypto/tls"
	"log"
)


func TLServer (certPem []byte, keyPem []byte) {

	// 
	TLScert, err := tls.X509KeyPair(certPem, keyPem)
	if err != nil {
		log.Fatal(err)
	}

	// Set up config for tls server
	cfg := SetUpServer(TLScert)

	ser := tls.Server( , cfg))
}


func SetUpServer (cert tls.Certificate) *tls.Config{

	serv := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	return serv
}*/