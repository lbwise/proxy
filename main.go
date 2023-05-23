package main

import (
	"net/http"
	"fmt"
	"io"
	"os"
)

func main() {
	PORT := os.Getenv("PORT")
	server := http.NewServeMux()
	server.HandleFunc("/", handleProxy)
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), server)
}

func handleProxy(wr http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	method := r.URL.Query().Get("method")

	req, _ := http.NewRequest(method, url, r.Body)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		wr.WriteHeader(500)
		wr.Write([]byte("An error occured making your request"))
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		wr.WriteHeader(500)
		wr.Write([]byte("An error occured reading your response"))
	}

    wr.Write(resBody)
}