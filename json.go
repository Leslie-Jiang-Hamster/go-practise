package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s, _ := json.Marshal([]int{1, 2, 4})
	fmt.Printf("\"%s\"\n", s)
	x := []int{}
	json.Unmarshal(s, &x)
	fmt.Println(x)
}