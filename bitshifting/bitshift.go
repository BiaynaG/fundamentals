package main

import "fmt"

//func main() {
//	x := 4
//	fmt.Printf("%d\t\t%b\n", x, x)
//
//	y := x << 1 //shifiting x by 1 bit
//	fmt.Printf("%d\t\t%b", y, y)
//}

// Version 1
//func main() {
//	kb := 1024 //or// kb = 1 << 10
//	mb := 1024 * kb
//	gb := 1024 * mb
//	fmt.Printf("%d\t\t\t%b\n", kb, kb)
//	fmt.Printf("%d\t\t\t%b\n", mb, mb)
//	fmt.Printf("%d\t\t%b\n", gb, gb)
//}

// Version 2 using iota
const (
	_ = iota // throw the first iota which is 0, away
	//kb = 1024
	kb = 1 << (iota * 10) // the second iota is now 1
	mb = 1 << (iota * 10) // the third iota is 2, so will shift mb by 20
	gb = 1 << (iota * 10) // shifting by 30
)

func main() {
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
