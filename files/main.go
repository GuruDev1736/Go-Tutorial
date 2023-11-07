package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("This is the file program")

	content := "Hello ! my name is guruprasad"

	file, err := os.Create("./Guruprasad.txt")
	checkerror(err)

	length, err := io.WriteString(file, content)
	checkerror(err)

	fmt.Println("The length of the file is ", length)
	defer file.Close()

	readfile("./Guruprasad.txt")

}

func readfile(filename string) {
	result, err := ioutil.ReadFile(filename)
	checkerror(err)

	fmt.Println("Inside the file is : \n", string(result))
}

func checkerror(err error) {
	if err != nil {
		panic(err)
	}
}
