package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<html><body>hello web server!</body></html>")
	})

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {

		// func (r *Request) FormValue(key string) string
		val := r.FormValue("say")

		fmt.Fprintf(w, "<html><body>echo:%s</body></html>",
			// func EscapeString(s string) string
			html.EscapeString(val))
	})

	err := http.ListenAndServe("localhost:8000", nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
