package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println(time.Now().Format(time.RFC3339), "Server start")

	initRouter()

	go func() {
		log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))
	}()

	log.Fatal(http.ListenAndServe(":80", nil))
}

func initRouter() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/event/js", eventJsHandler)
}
