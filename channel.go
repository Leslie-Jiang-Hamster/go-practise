package main

import "fmt"

func main() {
	msg := make(chan string, 2)

	msg <- "foo"
	msg <- "bar"
	
	fmt.Println(<- msg)
}