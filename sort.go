package main

import (
	"fmt"
	"slices"
)

func main() {
	arr := []int{2, 1, 3}
	fmt.Println(slices.IsSorted(arr))
	slices.Sort(arr)
	fmt.Println(arr)
}