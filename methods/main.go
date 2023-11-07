package main

import "fmt"

func main() {
	fmt.Println("This is the method program")

	guru := User{"Guruprasad Bhagat", 19, "Bekarai nagar", 1002, 2}

	guru.getexp()
	guru.setemail()

}

type User struct {
	Name    string
	age     int
	address string
	emp_id  int
	exp     int
}

func (u User) getexp() {

	fmt.Println("The experiencee of the user is : ", u.exp)
}

func (u User) setemail() {

	u.age = 155

	fmt.Println("The new age of the user is ", u.age)

}
