package main

import "fmt"

const login string = "Guruprasad" // global

func main() {

	// String
	var username string = "Guruprasad"
	fmt.Printf(username)
	fmt.Printf("\nThe type of variable is %T \n", username)

	// boolean
	var isloggedin bool = true
	fmt.Println(isloggedin)
	fmt.Printf("The type of variable is %T \n", isloggedin)

	// unsignedInteger
	var smallvalue uint = 10000
	fmt.Println(smallvalue)
	fmt.Printf("The type of variable is %T \n", smallvalue)

	//float
	var float float64 = 1256.052454
	fmt.Println(float)
	fmt.Printf("The type of variable is %T \n", float)

	// default values
	var integer int
	fmt.Println(integer)
	fmt.Printf("The type of variable is %T \n", integer)

	// implicit type
	var website = "This is the code "
	fmt.Println(website)

	// novar style
	name := "Guru"
	fmt.Println(name)

	//Global variable
	fmt.Println(login)

}
