package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
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
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	count := 0
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		count++
	})
	fmt.Printf("<a>タグの数は：%d個\n", count)
}
