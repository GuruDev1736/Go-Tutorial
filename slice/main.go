package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is the slice program")

	// var fruit = []string{"apple", "banana"}
	// fruit = append(fruit, "fig", "strawberry")
	// fruit = append(fruit[1:2])

	// fmt.Println("The list of the fruit is ", fruit)
	// fmt.Printf("The type of the slice is %T\n ", fruit)
	// fmt.Println("The length of the fruit is ", len(fruit))

	// var highscore = make([]int, 4)

	// highscore[0] = 111
	// highscore[1] = 222
	// highscore[2] = 333
	// highscore[3] = 444

	// highscore = append(highscore, 555, 666, 10)

	// fmt.Println("The list is : ", highscore)

	// sort.Ints(highscore)
	// fmt.Println("The sorted slice is : ", highscore)
	// fmt.Println("The sorted slice is : ", sort.IntsAreSorted(highscore))

	// how to remove a value from slices based on index

	var course = []string{"reactjs , javascript , kotlin , flutter , java"}
	fmt.Println(course)
	var index int = 2
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course)

}
