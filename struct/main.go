package main

import "fmt"

func main() {
	fmt.Println("This is struct")

	// no inheritance in go lang , no super , no parent

	guru := User{"Guruprasad", 19, "guruprasad1736@gmail.com", true}
	fmt.Println("The all details is :", guru)
	fmt.Printf("The name of the user is : %v\n ", guru.Name)

}

type User struct {
	Name   string
	age    int
	email  string
	status bool
}
