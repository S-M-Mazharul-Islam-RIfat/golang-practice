package main

import "fmt"

func main() {
	// Declare variable that are set to their zero value.
	var a int
	var b string
	var c float32
	var d bool

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n\n", d, d)

	// Declare variable and initialize
	// Using the short variable declaration operatort.

	aa := 10
	bb := "rifat"
	cc := 3.14
	dd := false

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n\n", dd, dd)

	aaa := int32(64)

	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)

}
