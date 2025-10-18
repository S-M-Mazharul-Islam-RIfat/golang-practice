package main

import "fmt"

func variableAndDataTypes() {
	// fmt.Println("Hello Gooo")

	/* go data types
	int - int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64
	float - float32,float64
	bool
	string
	*/

	var a int = 10
	var b = 20
	c := 30
	b = 52

	const value = 59
	fmt.Println(a, b, c, value)
}

func conditionalStatemnt() {
	CGPA := "A+"

	if CGPA == "A+" {
		fmt.Println("Pass")
	} else if CGPA == "B+" {
		fmt.Println("Ordhek Pass")
	} else {
		fmt.Println("Fail")
	}

	// a := 3
	// switch a {
	// case 1:
	// case 2, 3:
	// default:
	// }
}

func add(a int, b int) int {
	return a + b
}

func subAndMulti(a int, b int) (int, int) {
	return a - b, a * b
}

func getUserName() string {
	var name string
	fmt.Scanln(&name)
	return "Hello " + name
}

func main() {
	variableAndDataTypes()
	conditionalStatemnt()
	println(add(10, 20))
	sub, multi := subAndMulti(20, 10)
	fmt.Println(sub, multi)
	fmt.Println(getUserName())
}
