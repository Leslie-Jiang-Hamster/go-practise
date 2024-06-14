package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.WriteFile("foo", []byte("hello"), 0644)
	if err != nil {
		panic(err)
	}
	text, e := os.ReadFile("foo")
	if e != nil {
		panic(e)
	}
	fmt.Println(string(text))
}