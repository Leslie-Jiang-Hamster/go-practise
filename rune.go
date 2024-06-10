package main

import "fmt"

const s = "สวัสดี"

func main() {
	for _, c := range s {
		fmt.Printf("%#U\n", c)
	}
	fmt.Println('a' == 97)
}