package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/p?k=v#f"
	u, _ := url.Parse(s)
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(net.SplitHostPort(u.Host))
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(url.ParseQuery(u.RawQuery))
}