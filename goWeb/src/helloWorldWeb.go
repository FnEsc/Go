package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	server := &http.Server{
		Addr: "0.0.0.0:80",
	}
	http.HandleFunc("/", hello)
	server.ListenAndServe()
}
