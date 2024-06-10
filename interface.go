package main

import "fmt"

type Square struct {
	width int
}

func (s *Square) area() int {	
	return s.width * s.width
}

type Shape interface {
	area() int
}

func foo_(s Shape) int {
	return s.area() - 1
}

func main() {
	s := &Square{10}
	fmt.Println(foo_(s))
}