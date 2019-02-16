package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/* 第3回目のex03 */

func main() {
	inputurl := flag.String("url", "http://localhost:8000", "Type your URL without http://")
	flag.Parse()
	resp, err := http.Get(*inputurl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q\n", body[:])
}
