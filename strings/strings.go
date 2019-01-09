package main

import "fmt"

func main() {
	s := "Hello, playground"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s) // Converting the string to a slice of byte
	fmt.Println(bs)
	// fmt.Println(bs[0]) // index position 0
	fmt.Printf("%T\n", bs)
}
