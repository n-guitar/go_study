package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	// [{"id":0,"name":"kon","age":"37","ability":{"programming":"shell","operation":"argus"}}]
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Age     string  `json:"age"`
	Ability Ability `json:"ability"`
}

type Ability struct {
	Programming string `json:"programming"`
	Operation   string `json:"operation"`
}

func main() {
	file, err := ioutil.ReadFile("./go_study_member.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "file open err!: %v\n", err)
		os.Exit(1)
	}

	var person []Person

	if err := json.Unmarshal(file, &person); err != nil {
		fmt.Fprintf(os.Stderr, "json unmarshal err!: %v\n", err)
		os.Exit(1)
	}
	for _, p := range person {
		if p.Name == "kon" {
			fmt.Printf("%sさん、%s歳　得意なことは%s, 口癖はまじかよ！\n", p.Name, p.Age, p.Ability.Programming)
		} else {
			fmt.Printf("%sさん、%s歳　得意なことは%s\n", p.Name, p.Age, p.Ability.Programming)
		}
	}
}
