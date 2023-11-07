package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointer")

	// var one *int
	// fmt.Println("value of pointer is ", one)

	mynumber := 23

	var numpointer = &mynumber

	fmt.Println("The address of the number is ", numpointer)
	fmt.Println("The value of the number is ", *numpointer)

	*numpointer = *numpointer * 2
	fmt.Println("The new value is : ", mynumber)
	fmt.Println("The new value is : ", *numpointer)

}
