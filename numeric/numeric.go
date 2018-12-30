package main

import "fmt"

var a int
var b float64 // if we want to speciafy the numeric (float) type, use var

func main() {
	a = 10 // use the simple assignment (=), rather than redeclaration (:=)
	b = 42.34534
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
}

