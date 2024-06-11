package main

import "fmt"

func main() {
	msg := make(chan string)

	go func() {
		msg <- "foo"
		msg <- "bar"
	}()
	
	fmt.Println(<- msg)
}