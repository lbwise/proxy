package main

import (
	"net/http"
	"fmt"
	"io"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/", handleProxy)
	http.ListenAndServe(":8080", server)
}

func handleProxy(wr http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	method := r.URL.Query().Get("method")
	fmt.Println(method, url)
	req, err := http.NewRequest(method, url, r.Body)
	if err != nil {
		fmt.Println(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		wr.WriteHeader(500)
		wr.Write([]byte("An error occured making your request"))
	}
	fmt.Println(res)
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		wr.WriteHeader(500)
		wr.Write([]byte("An error occured reading your response"))
	}
    wr.Write(resBody)
}