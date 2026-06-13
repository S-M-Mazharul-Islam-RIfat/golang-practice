package main

import "fmt"

type user struct {
	name     string
	age      int32
	isActive bool
}

func main() {
	// Declare a variable of type user set to its zero value.
	var u1 user
	fmt.Printf("%+v \n", u1)

	// Declare a variable of type user and init using a struct literal.
	u2 := user{
		name:     "Rifat",
		age:      26,
		isActive: true,
	}

	fmt.Println("Name:", u2.name)
	fmt.Println("Age:", u2.age)
	fmt.Println("isActive:", u2.isActive)

	// Declare a varialbe using an anonymous struct.
	u3 := struct {
		name     string
		age      int32
		dollar   int
		isActive bool
	}{
		name:     "mj_riffu",
		age:      27,
		dollar:   100,
		isActive: true,
	}

	fmt.Printf("%+v\n", u3)
}
