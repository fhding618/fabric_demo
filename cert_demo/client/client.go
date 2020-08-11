package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {
	pool := x509.NewCertPool()
	caCertPath := "./cert_demo/client/ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	cliCrt, err := tls.LoadX509KeyPair("./cert_demo/client/client.crt", "./cert_demo/client/client.key")
	if err != nil {
		fmt.Println("Loadx509KeyPair err: ", err)
		return
	}

	tr := &http.Transport{
		TLSClientConfig:&tls.Config{
			RootCAs: pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{Transport:tr}
	resp, err := client.Get("https://test.com:8080")
	if err != nil {
		fmt.Println("Get error: ", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
