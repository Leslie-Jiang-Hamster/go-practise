package main

import "fmt"

type square struct {
	width int
}

func (s *square) area() int {
	return s.width * s.width 
}

func main() {
	s := square{10}
	fmt.Println(s.area())
}