package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "os"
)

type Takita struct {
    Id          int        `json:"id,string"`
    Name        string     `json:"name"`
    Age         int        `json:"age"`
    Abillity    struct {
        Programming string `json:"shell"`
        Operation   string `json:"argus"`
    } `json:"abillity"`
}

func main() {

    raw, err := ioutil.ReadFile("./read.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var fc Item
    json.Unmarshal(raw, %fc)

    for _, ft := range raw
    fmt.Printf("%sさん、%s才¥n", p.Id, p.Age)
}
