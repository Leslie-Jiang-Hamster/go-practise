package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup;

	for i := range 3 {
		wg.Add(1)
		go func() {
			time.Sleep(time.Second * time.Duration(i))
			fmt.Println(i)
			wg.Done()
		}()
	}

	wg.Wait()
}