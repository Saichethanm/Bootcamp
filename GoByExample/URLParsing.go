package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	fmt.Println(u.User.Password())
	fmt.Println(u.Host)
	fmt.Println(net.SplitHostPort(u.Host))
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery)
	fmt.Println(url.ParseQuery(u.RawQuery))
}
