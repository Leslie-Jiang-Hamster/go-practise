package main

import "fmt"

func main() {
	c := make(chan string)
	select {
	case c <- "foo":
		fmt.Println("done")
	default:
		fmt.Println("fail")
	}
}