package main

import (
	"fmt"
	"time"
)

func Foo(x int) {
	time.Sleep(time.Second)
	fmt.Println(x)
	sync <- x
}

var sync = make(chan int)

func main() {
	go Foo(1)
	go Foo(2)
	<- sync
	<- sync
	fmt.Println("done")
}