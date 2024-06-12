package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	s := []string{"fooo", "bar"}
	cmp_ := func (x, y string) int {
		return cmp.Compare(len(x), len(y))
	}
	slices.SortFunc(s, cmp_)
	fmt.Println(s)
}