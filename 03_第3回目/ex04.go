package main

import (
	"fmt"
	"net/http"
)

func main() {
	req, _ := http.NewRequest("GET", "http://golang.jp/", nil)

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	code := resp.StatusCode
	fmt.Printf("%d\n", code)
}
