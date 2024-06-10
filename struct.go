package main

import "fmt"

type Person struct {
	name string
	age int
}

func main() {
	p := Person{name: "leslie", age: 100}
	fmt.Println((&p).name, p.age)	
}