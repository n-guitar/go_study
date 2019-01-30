package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

/* 第2回目のex02 */

func main() {
	file, err := os.Open("./all.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		time.Sleep(time.Second)
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
