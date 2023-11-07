package main

import "fmt"

func main() {
	fmt.Println("This is the function")

	read("Hello user")

	result, mymessage := proadd(1, 2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Println(result)
	fmt.Println(mymessage)

}

func proadd(value ...int) (int, string) {

	total := 0

	for _, val := range value {

		total += val

	}
	return total, "This is the proadd function"

}

func read(text string) {
	fmt.Println(text)
}
