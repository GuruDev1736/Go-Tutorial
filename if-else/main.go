package main

import "fmt"

func main() {
	fmt.Println("This is the if else ")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else {
		result = "Not a regular user"
	}

	fmt.Println(result)
}
