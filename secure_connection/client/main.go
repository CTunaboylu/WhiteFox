package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var hostname string = "localhost"

func main() {
	client := getClient()
	resp, err := client.Get("http://" + hostname + ":8081")
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	body, what2 := ioutil.ReadAll(resp.Body)
	fmt.Printf("[what2]%v of type %T\n", what2, what2)
	fmt.Printf("[body]%v of type %T\n", body, body)

	fmt.Printf("Status: %s Body: %s\n ", resp.Status, string(body))
}

func getClient() *http.Client {
	config := &tls.Config{
		GetClientCertificate:  utils.ClientCertReqFunc("", ""),
		VerifyPeerCertificate: utils.CertificateChains,
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}
	return client
}
