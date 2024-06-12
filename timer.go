package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTimer(time.Second)
	go func() {
		<- t.C
		fmt.Println("fire")
	}()
	t.Stop()
	fmt.Println("stop")
}