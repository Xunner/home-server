package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	if err != nil {
		fmt.Printf("helloHandler err: %v", err)
		return
	}
}

func main() {
	fmt.Println("Server start")

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", helloHandler)

	log.Fatal(http.ListenAndServe(":13579", nil))
}
