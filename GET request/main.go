package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("This is the getrequest")

	// performGET()
	performPost()

}

func performGET() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("status code is : ", response.StatusCode)
	fmt.Println("Content Lengeth of the  code is : ", response.ContentLength)

	var responseString strings.Builder
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)

	}
	bytecode, _ := responseString.Write(content)

	fmt.Println("the bytecount is ", bytecode)
	fmt.Println("the bytecount is ", responseString.String())
	// fmt.Println("Content : ", string(content))
}

func performPost() {
	const myurl = "http://localhost:8000/post"

	// fake Json payload

	requestbody := strings.NewReader(`
		{
			"name" : "Guruprasad"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestbody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))

}
