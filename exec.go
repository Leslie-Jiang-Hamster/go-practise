package main

import (
	"fmt"
	"os/exec"
)

func main() {
	s, e := exec.Command("ls").Output()
	if e != nil {
		panic(e)
	}
	fmt.Println(string(s))
}