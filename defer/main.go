package main

import "fmt"

func main() {

	fmt.Println("This is the deefer function")

	defer fmt.Println("world")
	fmt.Println("Hello")
	getfefer()

}

func getfefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
