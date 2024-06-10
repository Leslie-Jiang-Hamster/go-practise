package main

import "fmt"

func main() {
	m := map[string]int{}
	_, found := m["foo"]
	fmt.Println(found)
	m["foo"] = 42
	fmt.Println(m["foo"])
}