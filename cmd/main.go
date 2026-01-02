package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	searchHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "search page here")
	}

	http.HandleFunc("/", searchHandler)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("fatal err")
	}
}
