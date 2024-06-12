package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 10)
	for i := range 5 {
		c <- i
	}
	close(c)
	for t := range time.Tick(time.Millisecond * 500) {
		fmt.Printf("%s %d done\n", t, <- c)
	}
}