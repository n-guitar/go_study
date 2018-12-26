package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		fmt.Printf("url_name:%s\n", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "http get err:%v\n", err)
			os.Exit(1)
		}
		status := resp.StatusCode
		fmt.Printf("ステータス:%d\n", status)
	}
}
