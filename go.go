package main

import (
	"fmt"
	"time"
)

func _foo() {
	for i := range 1000 {
		if i % 100 != 0 {
			continue
		}
		fmt.Println(i)
	}
}

func main() {
	go _foo()
	go _foo()
	time.Sleep(time.Second)
}