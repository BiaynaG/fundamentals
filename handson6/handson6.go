// Using iota, create 4 constants for the NEXT 4 years. Print the constant values.

package main

import "fmt"

const (
	_  = iota
	y  = 2019
	y1 = y + iota
	y2 = y + iota
	y3 = y + iota
	y4 = y + iota
)

func main() {
	fmt.Println(y1, y2, y3, y4)
}
