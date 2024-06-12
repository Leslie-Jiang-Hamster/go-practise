package main

import (
	"fmt"
	"time"
)

func bar(c chan<- string) {
	time.Sleep(time.Second / 2)
	c <- "done"
}

func main() {
	c := make(chan string, 1)
	go bar(c)
	select {
		case msg := <- c:
			fmt.Println(msg)
		case <- time.After(time.Second):
			fmt.Println("timeout")
	}
}