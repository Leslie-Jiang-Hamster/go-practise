package main

import "fmt"

func main() {
	for i := range 5 {
		if i % 2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}