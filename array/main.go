package main

import "fmt"

func main() {
	fmt.Println("This is the array program")

	var names [4]string

	names[0] = "Guruprasad"
	names[1] = "meow"
	names[2] = "ayush"
	names[3] = "daivat"

	fmt.Println("The names are : ", names)
	fmt.Println("The names are : ", len(names))

	var fruit = [3]string{"apple", "banana", "fig"}
	fmt.Println("The fruit are : ", fruit)
	fmt.Println("The fruit are : ", len(fruit))
}
