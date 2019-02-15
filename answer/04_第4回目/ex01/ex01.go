package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/", http.StripPrefix("/", fs))

	http.ListenAndServe(":9999", nil)
}
