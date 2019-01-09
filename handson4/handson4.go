// Write a program that
// 1. Assigns an int to a variable
// 2. Prints that int in decimal, binary, and hex
// 3. Shifts the bits of that int over 1 position to the left, and assigns that to a
// variable
// 4. Prints that variable in decimal, binary, and hex

package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)

	b := a << 1
	fmt.Printf("%d\t%b\t%#x", b, b, b)
}
