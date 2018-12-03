package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type person struct {
	// [{"id":0,"name":"kon","age":"37","ability":{"programming":"shell","operation":"argus"}}]
	id      string    `json:"id"`
	name    string    `json:"name"`
	age     string    `json:"age"`
	ability []ability `json:"ability"`
}

type ability struct {
	programming string `json:"programming"`
	operation   string `json:"operation"`
}

func main() {
	file, err := ioutil.ReadFile(`./kon.json`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "file open err!: %v\n", err)
		os.Exit(1)
	}

	jsonBytes := []byte(file)
	var p person
	fmt.Println(string(jsonBytes))
	json.Unmarshal(jsonBytes, &p)
	fmt.Println(p)
	fmt.Printf("基本情報: %s\t%s\t%s\t%s\t", p.id, p.name, p.age, p.ability)
}
