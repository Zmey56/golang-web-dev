package main

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/crypto/acme"
	"log"
	"net/http"
	"rsc.io/letsencrypt"
)

func main() {
	http.HandleFunc("/", foo)
	letsencrypt.Client = &acme.Client{DirectoryURL: "https://acme-staging-v02.api.letsencrypt.org/directory"}

	var m letsencrypt.Manager
	if err := m.CacheFile("lets encrypt.cache"); err != nil {
		log.Fatalln(err)
	}

	go http.ListenAndServe(":8080", http.HandlerFunc(letsencrypt.RedirectHTTP))

	srv := &http.Server{
		Addr: ":10443",
		TLSConfig: &tls.Config{
			GetCertificate: m.GetCertificate,
		},
	}

	log.Fatalln(srv.ListenAndServeTLS("", ""))
}

func foo(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello TLS")
}
