package main

import "fmt"

func main() {
	c := make(chan string, 1)
	c <- "bar"
	close(c)
	msg, more := <- c;
	fmt.Println(msg, more)
	msg, more = <- c;
	fmt.Println(msg, more)
}