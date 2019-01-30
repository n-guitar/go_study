package main

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
			fmt.Printf("ID:%d\nName:%s\nAge:%s\n Ability(Programming):%s\n Ability(Operation):%s\n", g.ID, g.Name, g.Age, g.Ability.Programming, g.Ability.Operation)
		}
	}
}
