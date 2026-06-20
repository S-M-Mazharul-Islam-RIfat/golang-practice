package main

import "fmt"

// standard or named function
func add(a int, b int) int {
	return a + b
}

// HOF
func HOF(x int, y int, f func(p int, q int) int) func(p int, q int) int {
	fmt.Println("Inside HOF", x+y+f(10, 2))
	return add
}

func main() {
	fmt.Println("Functional Programming in Golang")
	fmt.Println(add(5, 10))

	// anonymous and IIFE
	func(x int, y int) {
		fmt.Println(x + y)
	}(5, 7)

	// function expression
	sumOfThreeValues := func(x int, y int, z int) {
		fmt.Println(x + y + z)
	}
	sumOfThreeValues(10, 20, 30)

	// HOF
	anotherFunc := HOF(1, 2, add)
	fmt.Println(anotherFunc(5, 7))
}

func init() {
	fmt.Println("I am the init function 1")
}
func init() {
	fmt.Println("I am the init function 2")
}
