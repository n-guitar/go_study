package main

/*
注意！
Scannerは一行毎にデータを読み込む仕様であり、一行読み込む毎に
json.Unmarshallをする作りにすると複数行あるjsonに対応できない。
（本プログラムが該当する）
複数行のjson = 一行中にjsonの始まりと終わりの"{}"が含まれない　ので
　"unexpected end of JSON input"
エラーが出て異常終了してしまう。
*/

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type goengineer struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     string `json:"age"`
	Ability struct {
		Programming string `json:"programming"`
		Operation   string `json:"operation"`
	}
}

func main() {
	file, err := os.Open("./kon.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var goeng []goengineer
	for scanner.Scan() {
		if err := json.Unmarshal(scanner.Bytes(), &goeng); err != nil {
			log.Fatal(err)
		}
		for _, g := range goeng {
			fmt.Printf("ID:%d\nName:%s\nAge:%s\n", g.ID, g.Name, g.Age)
		}
	}
}
