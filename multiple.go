package main

import "fmt"

func foo() (int, int) {
	return 4, 2
}

func main() {
	_, two := foo()
	fmt.Println(two)
}