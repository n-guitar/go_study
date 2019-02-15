package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var lines int = 0

func set(url string) {
	doc, _ := goquery.NewDocument(url)
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		//url, _ := s.Attr("href")
		//fmt.Println(url)
		lines += 1
	})
}

func main() {
	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/", http.StripPrefix("/", fs))

	go http.ListenAndServe(":9999", nil)
	url := "http://localhost:9999"
	set(url)
	fmt.Printf("Number of [a] tags is %d\n", lines)
}
