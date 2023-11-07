package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("This is the webrequest")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type is %T\n", response)

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("The response is ", string(data))

}
