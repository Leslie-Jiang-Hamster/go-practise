package main

import (
	"fmt"
	"sync"
)

func main() {
	mu := sync.Mutex{}
	sum := 0

	thread := func() {
		for range 100 {
			mu.Lock()
			sum++
			mu.Unlock()
		}
	}
	thread()
	thread()

	fmt.Println(sum)
}