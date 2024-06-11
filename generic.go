package main

import (
	"cmp"
	"fmt"
)

func lt[T cmp.Ordered](x, y T) bool {
	return x < y
}

func main() {
	fmt.Println(1<2)
	fmt.Println(1.0<2.0)
	fmt.Println("1"<"2")
}