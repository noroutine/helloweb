package main

import (
	"fmt"
	"log"
	"net/http"
	"html"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q!!!", html.EscapeString(r.URL.Path))
}

func main( ) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}