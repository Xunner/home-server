package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	if err != nil {
		fmt.Printf("helloHandler err: %v", err)
		return
	}
}

func main() {
	fmt.Println(time.Now().Format(time.RFC3339), "Server start")

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", helloHandler)

	go func() {
		log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))
	}()

	log.Fatal(http.ListenAndServe(":80", nil))
}
