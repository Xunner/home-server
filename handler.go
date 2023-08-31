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

func eventJsHandler(w http.ResponseWriter, r *http.Request) {
	// log
	queries := r.URL.Query()
	log.Printf("RequestURI=%s, Header=%v, queries=%v", r.URL.RequestURI(), r.Header, queries)

	baseUrl := "https://analytics.tiktok.com/i18n/pixel/events.js"
	params := url.Values{}
	params.Add("sdkid", queries.Get("sdkid"))
	params.Add("lib", queries.Get("lib"))

	// 创建一个新的URL
	fullUrl := fmt.Sprintf("%s?%s", baseUrl, params.Encode())

	// 发送GET请求
	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Printf("http.Get err: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Body.Close() err: %v", err)
		}
	}(resp.Body)

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("io.ReadAll err: %v", err)
	}

	// Copy body
	_, err = fmt.Fprintf(w, string(body))
	if err != nil {
		log.Printf("return body err: %v", err)
		return
	}
	// Copy all headers
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	// 设置响应的状态码为200 OK
	w.WriteHeader(http.StatusOK)
}
