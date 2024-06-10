package main

import "fmt"

type Job struct {
	salary int
}

type Person struct {
	name string
	Job
}

func main() {
	p := Person{"leslie", Job{130_000}}
	fmt.Println(p)
}