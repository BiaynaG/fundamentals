// Program that prints a number in decimal, binary, and hex

package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("%d\n", a) // decimal, base 10
	fmt.Printf("%b\n", a) // binary, base 2
	fmt.Printf("%x\n", a) // hexadecimal, base 16 with lower letter

	// shorter
	// fmt.Printf("%d\t%b\t%#x", a, a, a)
}
