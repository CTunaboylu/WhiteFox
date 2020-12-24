package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/lizrice/secure-connections/utils"
)

func main() {
	server := getServer()
	http.HandleFunc("/", myHandler)
	log.Fatal(server.ListenAndServeTLS("", ""))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling request")
	w.Write([]byte("Hey Gophercon\n"))
}

func getServer() *http.Server {
	tls := &tls.Config{
		GetCertificate: utils.CertReqFunc("", "")
		VerifyPeerCertificate: utils.CertificateChains,
	}
	server := &http.Server{
		Addr:      ":8081",
		TLSConfig: tls,
	}
	return server
}
