package main

import "fmt"

func main() {
	fmt.Println("This is the loop statement")

	//days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday "}

	// fmt.Println(days, "\n")

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// fmt.Println("\n")

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// fmt.Println("\n")

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n ", index, day)
	// }

	rough := 1

	for rough < 10 {

		if rough == 5 {
			rough++
			break
		}

		fmt.Println("Value is ", rough)
		rough++

	}

}
