package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MyMux struct {

}

func (p *MyMux)ServeHTTP(res http.ResponseWriter, req *http.Request)  {
	fmt.Fprintf(res, "Hi, This is an example of https service in golang!\n")

	fmt.Fprintf(res,
		`[{"Name":"jason","Age":35,"Weight":60.3,"Speciality":"computer science","Hobby":["tennis","swimming","reading"],"Score":725.5,"Secret":"SRRMb3ZlFFlvdSE="}]`)
}

func main()  {
	pool := x509.NewCertPool()
	caCertPath := "./cert_demo/server/ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	s := &http.Server{
		Addr: ":8080",
		Handler: &MyMux{},
		TLSConfig:&tls.Config{
			ClientCAs:pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}


	err = s.ListenAndServeTLS("./cert_demo/server/server.crt", "./cert_demo/server/server.key")
	if err != nil {
		fmt.Println("ListenAndServeTLS err:", err)
	}
}