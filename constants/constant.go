package main

import "fmt"

const (
	// untyped constants
	a = 42
	b = 39.9
	c = "Bia Grig"
	// a int = 42 //typed constant
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
