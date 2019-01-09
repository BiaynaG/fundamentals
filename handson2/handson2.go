// Using the following operators, write expressions and assign their values to variables:
// ==		<=		>=		!=		<		>

package main

import "fmt"

func main() {
	x := 10
	y := 20
	a := (x == y)
	b := (x <= y)
	c := (x >= y)
	d := (x != y)
	e := (x < y)
	f := (x > y)
	fmt.Println(a, b, c, d, e, f)
}
