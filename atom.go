package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var sum atomic.Int64
	var wg sync.WaitGroup
	wg.Add(2)

	thread := func() {
		defer wg.Done()
		for range 1000 {
			sum.Add(1)
		}
	}
	thread()
	thread()

	wg.Wait()
	fmt.Println(sum.Load())
}