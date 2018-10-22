package main

import (
	"net/url"
	"fmt"
	"net"
)

func URLParse() {

	s := "postgres://user:password@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)

	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	/*
	To get query params in a string of k=v format, user RawQuery.
	One can also parse query params into a map.
	The parsed query param maps are from strings to slices of strings,
	so index into [0] if you only want the first value.
	 */
	 fmt.Println("u.RawQuery:", u.RawQuery)
	 m, _ := url.ParseQuery(u.RawQuery)
	 fmt.Println("Map:", m)
	 fmt.Print(m["k"][0])

}