package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("https://github.com/trending?since=daily")
	if err != nil {
		fmt.Print("url scarapping failed")
	}
	var rankCount = 0
	doc.Find("ol > li ").Each(func(_ int, s *goquery.Selection) {
		url := s.Children().Find("a")
		star := s.Children().Find("span.float-sm-right")
		rankName, _ := url.Attr("href")
		rankText := star.Text()
		var rankNumber []string
		rankNumber = strings.Fields(rankText)
		rankCount++
		if 10 >= rankCount {
			fmt.Printf("%d位\t スター数:%s\t リポジトリ名：https://github.com%s\n", rankCount, rankNumber[0], rankName)
		}
	})
}
