package main

import "fmt"

func main() {
	var fruits [5]string
	fruits[0] = "Apple"
	fruits[1] = "Orage"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}

	numbers := [5]int{10, 20, 30, 40, 50}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}
}
