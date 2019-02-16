package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

/* 第3回目のex04 */

func main() {
	inputurl := flag.String("url", "http://localhost:8000", "Type your URL without http://")
	flag.Parse()
	resp, err := http.Get(*inputurl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	code := resp.StatusCode
	fmt.Printf("%d\n", code)
}
