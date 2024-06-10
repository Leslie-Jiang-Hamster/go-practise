package main

import "fmt"

func sum_(nums ...int) int {
	sum_ := 0
	for _, x := range nums {
		sum_ += x
	}
	return sum_
}

func main() {
	s := []int{1, 2, 3}
	fmt.Println(sum_(s...))
}