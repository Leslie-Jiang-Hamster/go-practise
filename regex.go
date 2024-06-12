package main

import (
	"fmt"
	"regexp"
)

func main() {
	r, _ := regexp.Compile("p[a-z]*h")
	fmt.Println(r.MatchString("peach porch"))
	fmt.Println(r.FindAllString("peach porch", -1))
}