package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	if err != nil {
		log.Printf("helloHandler err: %v", err)
		return
	}
}
