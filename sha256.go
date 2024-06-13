package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "hello, world"
	fmt.Println(s)

	c := sha256.New()
	c.Write([]byte(s))
	fmt.Printf("%x\n", c.Sum(nil))
}