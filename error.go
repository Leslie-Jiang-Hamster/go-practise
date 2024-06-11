package main

import (
	"errors"
	"fmt"
)

func main() {
	e1 := fmt.Errorf("I'm an error")
	e2 := fmt.Errorf("I'm an error, too. %w", e1)
	fmt.Println(errors.Is(e1, e2))
	fmt.Println(errors.Is(e2, e1))
	fmt.Println(e1)
}