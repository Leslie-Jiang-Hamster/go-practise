package main

import "fmt"

func gen() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	next := gen()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}