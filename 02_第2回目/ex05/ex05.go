package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	file, err := ioutil.ReadFile("./go_study_member.json")
	if err != nil {
		log.Fatal(err)
	}
	var goeng []goengineer
	if err := json.Unmarshal(file, &goeng); err != nil {
		log.Fatal(err)
	}
	for _, g := range goeng {
		fmt.Printf("ID:%d\nName:%s\nAge:%s\n", g.ID, g.Name, g.Age)
	}
}
