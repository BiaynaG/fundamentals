package main

import "fmt"

const (
	a = iota // successive untyped integers
	b = iota // + 1
	c = iota
)

const (
	d = iota
	e // no need to use iota again
	f
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
