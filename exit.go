package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println(42)

	os.Exit(0)
}