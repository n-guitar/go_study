package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	get()
}

func get() {
	resp, err := http.Get("https://github.com/trending")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	count := 1
	doc.Find("html > body > div > div > div > div > div > ol > li").Each(func(_ int, s *goquery.Selection) {
		repoid, exist := s.Attr("id")
		reponame := strings.Trim(repoid, "pa-")
		if exist == true && count <= 10 {
			fmt.Printf("%d位\n   リポジトリ名 : %s\n", count, reponame)
			s.Find("a").Each(func(_ int, f *goquery.Selection) {
				hrefall, existhref := f.Attr("href")
				if existhref == true && strings.HasSuffix(hrefall, "/stargazers") == true {
					stars := strings.TrimSpace(f.Text())
					fmt.Printf("   スター数     : %s\n", stars)
				}
			})
			count++
		}
	})

}
