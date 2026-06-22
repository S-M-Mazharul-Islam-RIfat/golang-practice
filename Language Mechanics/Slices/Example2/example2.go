package main

import "fmt"

func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i := range slice {
		fmt.Printf("[%d] %p %s\n", i, &slice[i], slice[i])
	}
}

func main() {
	slice := make([]string, 5, 8)

	slice[0] = "Apple"
	slice[1] = "Orange"
	slice[2] = "Banana"
	slice[3] = "Grape"
	slice[4] = "Plum"

	inspectSlice(slice)

	var data []int

	for record := 1; record <= 10; record++ {
		data = append(data, record)
	}

	fmt.Println(data, len(data), cap(data))

	slice1 := []string{"A", "B", "C", "D", "E"}
	slice1 = append(slice1, "G")
	slice2 := slice1[2:4]

	fmt.Println(slice1)
	fmt.Println(slice2, len(slice2), cap(slice2))

	// slice2[0] = "K"

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice3 := slice1[2:4:4]
	fmt.Println(slice3, len(slice3), cap(slice3))

	// slice3[0] = "K"

	// fmt.Println(slice1)
	// fmt.Println(slice2)
	// fmt.Println(slice3)

	slice3 = append(slice3, "P")

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

}
