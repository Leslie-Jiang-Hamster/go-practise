package main

import (
	"fmt"
	"slices"
)

func main() {
	s1 := make([]int, 3)
	s2 := []int{1, 2, 3}
	fmt.Println(slices.Equal(s1, s2))
	s1 = append([]int{}, 1, 2, 3)
	fmt.Println(slices.Equal(s1, s2))
}