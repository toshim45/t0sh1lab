package main

import (
	"crypto/tls"
	"log"

	"github.com/go-resty/resty/v2"
)

func main() {
	log.Println("client")
	// Parsing public/private key pair from a pair of files. The files must contain PEM encoded data.
	cert, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		log.Fatalf("ERROR client certificate: %v", err)
	}

	// Create a resty client
	client := resty.New()
	client.SetRootCertificate("ca.crt") // self-signed
	client.SetCertificates(cert)
	resp, err := client.R().
		EnableTrace().
		Post("https://macko-ss.local/hello") // or get

	if err != nil {
		log.Fatalf("ERROR http request: %v", err)
	}

	log.Printf("request-trace: %+v", resp.Request.TraceInfo())

	log.Println("response-status: ", resp.Status())
	log.Println("response-body: ", resp)
}

/*
* at go-lab directory
go run ./cmd/tls-http/client/main.go
*/
