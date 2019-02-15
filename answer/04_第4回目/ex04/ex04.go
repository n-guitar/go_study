package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// TrendRankingを出力させる。
func GetRank(url string) {
	trend := 1
	doc, _ := goquery.NewDocument(url)
	doc.Find("li").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("id")
		if trend <= 10 {
			//　空行を出力させないため、len(url)を使用する。
			if len(url) > 0 {
				// リポ名の頭に「pa-」がつくので、削除する。
				replace := strings.Replace(url, "pa-", "", 1)
				fmt.Printf("Trend%d位は	→	%s\n", trend, replace)
				trend++
			}
		}
	})
}

func format(r *regexp.Regexp, target string) string {
	result := strings.Replace(target, "\n", "", -1)
	result = strings.Replace(result, " ", "", -1)
	return result
}

// StargaxersのURLを取得し、starまで取得してくる。
func GetStars(url string) {
	stars := 1
	doc, _ := goquery.NewDocument(url)
	doc.Find("div.f6.text-gray.mt-2 > a").Each(func(_ int, s *goquery.Selection) {
		hikakin, _ := s.Attr("href")
		if stars <= 10 {
			if strings.Contains(hikakin, "stargazers") {
				fmt.Println(format(nil, s.Text()))
				stars++
			}
		}
	})
}

func main() {
	url := "https://github.com/trending"
	GetRank(url)
	GetStars(url)
}
