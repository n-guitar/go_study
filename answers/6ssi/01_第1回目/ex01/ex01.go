package main

/* ex01 */

import "fmt"
import "os"
import "os/exec"

func main() {
	fmt.Printf("Command \"ls\" returned :\n%s", comexec("ls"))
}

func comexec(comstr string) []byte {
	cmd := exec.Command(comstr)
	out, err := cmd.Output()
	/* Here's another option...And it's much more simple.
	out, err := exec.Command(comstr).Output()		*/
	if err != nil {
		fmt.Printf("Oops -> %s", err)
		os.Exit(1)
		/* Godoc says code "0" indicates success. */
	}
	return out
}
