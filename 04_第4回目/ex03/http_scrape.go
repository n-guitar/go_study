package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	Scrape()
}

func Scrape() {
	// Request the HTML page.
	res, err := http.Get("http://localhost:9999")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d &s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	title := doc.Find("title").Text()
	fmt.Println("title:", title)
	fmt.Println("------------")

	// var count = 1
	doc.Find("a").Each(func(i int, s *goquery.Selection) {

		linkText := s.Text()
		linkUrl, _ := s.Attr("href")
		fmt.Println("int:", i)
		fmt.Println("linkText:", linkText)
		fmt.Println("linkUrl:", linkUrl)
		fmt.Println("------------")
		// count++
	})
	// fmt.Println("aタグの合計数:", count)

}
