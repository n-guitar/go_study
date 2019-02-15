package main

import (
    "fmt"
    "os/exec"
)

func main() {

cmd := exec.Command("date")
out, _ := cmd.Output()
fmt.Printf("結果: %s", out)

}
