package main

import (
	"fmt"
	"os"
)

func main() {
	text, err := os.ReadFile("foo")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(text))
}