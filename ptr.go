package main

import "fmt"

func main() {
	foo := 42
	ptr := &foo
	*ptr = 0
	fmt.Println(ptr)
	fmt.Println(foo)
}