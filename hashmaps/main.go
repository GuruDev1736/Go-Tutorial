package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the hashmaps")

	language := make(map[string]string)
	language["Java"] = "java"
	language["Js"] = "javascript"
	language["react"] = "react"
	language["go"] = "Go"

	fmt.Println("List of all languages", language)
	fmt.Println("List of all languages", language["Js"])

	delete(language, "Js")
	fmt.Println("List of all languages", language)

	// loops

	for key, value := range language {
		fmt.Printf("for key %v , value is %v\n", key, value)
	}

}
