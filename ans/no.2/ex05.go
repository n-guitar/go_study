package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name string
	Age  string
}

func main() {
	bytes, err := ioutil.ReadFile("./j.json")
	if err != nil {
		panic(err)
	}

	var ps []Person
	if err := json.Unmarshal(bytes, &ps); err != nil {
		panic(err)
	}

	for _, p := range ps {
		fmt.Println(p.Name, p.Age)
	}
}
