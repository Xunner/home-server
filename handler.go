package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	if err != nil {
		log.Printf("helloHandler err: %v", err)
		return
	}
}

func eventJsHandler(w http.ResponseWriter, r *http.Request) {
	// log
	reader, err := r.GetBody()
	if err != nil {
		log.Println(err)
	}
	var bs []byte
	n, err := reader.Read(bs)
	if err != nil {
		log.Println(err, n)
	}
	log.Printf("URI=%s, Header=%v, body=%s", r.URL.RequestURI(), r.Header, string(bs))

	// return
	_, err = fmt.Fprintf(w, "On developing...")
	if err != nil {
		log.Printf("eventJsHandler err: %v", err)
		return
	}
}
