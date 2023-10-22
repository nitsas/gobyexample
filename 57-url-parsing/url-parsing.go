package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	urlString := "postgres://usrname:passwd@host.com:5432/path?foo=first&foo=second&hello[]=world&bar=&baz#fragment"
	fmt.Println("urlString:", urlString)

	u, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}
	fmt.Println("u := url.Parse(urlString):", u)
	fmt.Println()

	fmt.Println("u.Scheme:", u.Scheme)
	fmt.Println("u.User:", u.User)
	fmt.Println("u.User.Username():", u.User.Username())
	password, hasPassword := u.User.Password()
	fmt.Printf("u.User.Password(): %s, %t\n", password, hasPassword)
	// hasPassword would be false if the url didn't contain a password
	fmt.Println("u.Host:", u.Host)
	hostname, port, err := net.SplitHostPort(u.Host)
	// port is a string
	fmt.Printf("net.SplitHostPort(u.Host): %s, %s, %e\n", hostname, port, err)
	fmt.Println("u.Path:", u.Path)
	fmt.Println("u.Fragment:", u.Fragment)
	fmt.Println("u.RawQuery:", u.RawQuery)
	q, err := url.ParseQuery(u.RawQuery)
	fmt.Printf("url.ParseQuery(u.RawQuery): %#v, %e\n", q, err)
	// q["hello"], q["foo"], q["bar"], q["baz"] are all non-empty []string
}
