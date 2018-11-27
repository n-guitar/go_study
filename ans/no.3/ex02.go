package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hello" {
		fmt.Fprintf(w, "hello web server!")
	} else {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	}
}
