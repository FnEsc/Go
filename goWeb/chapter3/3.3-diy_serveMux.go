package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", _indexHandler)
	mux.HandleFunc("/hi", _hiHandler)
	mux.HandleFunc("/hi/web", _hi_webHandler)
	// 以 / 结尾允许向上匹配， 非 / 结尾则只能精准匹配
	// 那么 /hi123 会匹配到 _indexHandler
	// 那么 /hi/web123 会匹配到 _indexHandler
	// 那么 /hi/web/123 还是会匹配到 _indexHandler

	server := &http.Server{
		Addr:         ":8086",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func _indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi _indexHandler")
}

func _hiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi _hiHandler")
}
func _hi_webHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi _hi_webHandler")
}
