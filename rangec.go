package main

import "fmt"

func main() {
	c := make(chan string, 2)
	c <- "foo"
	c <- "bar"
	close(c)
	for x := range c {
		fmt.Println(x)
	}
}