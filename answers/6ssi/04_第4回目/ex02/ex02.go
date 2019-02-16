package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/", http.StripPrefix("/", fs))
	go http.ListenAndServe(":9999", nil)
	go get()
	time.Sleep(time.Second)
}

func get() {
	resp, err := http.Get("http://localhost:9999")
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
