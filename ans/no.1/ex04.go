package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {

	var cmds []string

	flag.Parse()
	for i := 0; i < flag.NArg(); i++ {
		cmds = append(cmds, flag.Arg(i))
	}

	out, err := exec.Command(cmds[0], cmds[1:]...).Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
