package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	h, e := http.Get("http://example.com")
	if e != nil {
		panic(e)
	}
	defer h.Body.Close()

	s := bufio.NewScanner(h.Body)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}