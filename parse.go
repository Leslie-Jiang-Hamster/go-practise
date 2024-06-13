package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.ParseInt("0xd", 0, 64))
	fmt.Println(strconv.Atoi("9"))
}