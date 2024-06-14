package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	p := filepath.Join("foo", "bar")
	fmt.Println(p, filepath.Ext("baz.jpg"), filepath.Dir(p))
}