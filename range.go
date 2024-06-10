package main

import "fmt"

func main() {
	a := []int{3,2,1}
	m := map[string]int{"foo":1}
	sum := 0
	for _, x := range a {
		sum += x
	}
	fmt.Println(sum)
	sum = 0
	for _, x := range m {
		sum += x
	}
	fmt.Println(sum)
}