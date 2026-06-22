package main

import "fmt"

func main() {
	arr1 := [6]string{"This", "is", "a", "golang", "interview", "questions"}
	//slice from an existing array
	s1 := arr1[1:4]
	fmt.Println(s1, len(s1), cap(s1))

	// slice from slice
	s2 := arr1[1:2]
	fmt.Println(s2, len(s2), cap(s2))

	// slcie literals
	arr2 := []string{"Programming", "is", "fun"}
	fmt.Println(arr2, len(arr2), cap(arr2))

	// slice using make function with length
	arr3 := make([]int, 3)
	fmt.Println(arr3, len(arr3), cap(arr3))

	// empty or nil slice
	var arr4 []int
	fmt.Println(arr4)

	// append
	var arr5 []int
	arr5 = append(arr5, 3)
	arr5 = append(arr5, 4)
	arr5 = append(arr5, 9)
	fmt.Println(arr5, len(arr5), cap(arr5))
	y := arr5
	x := arr5

	arr5 = append(arr5, 7)
	fmt.Println(arr5, len(arr5), cap(arr5))
	fmt.Println(y, len(y), cap(y))

	y = append(y, 5)
	y[0] = 17
	fmt.Println(arr5, len(arr5), cap(arr5))
	fmt.Println(y, len(y), cap(y))
	x[0] = 20
	fmt.Println(arr5, len(arr5), cap(arr5))
	fmt.Println(x, len(x), cap(x))
	fmt.Println(y, len(y), cap(y))

}
