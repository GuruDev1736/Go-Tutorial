package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=dkjskgkjgkg"

func main() {
	fmt.Println("This is the handling url program")

	// fmt.Println(myurl)

	// result, err := url.Parse(myurl)

	// if err != nil {
	// 	panic(err)

	// }

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.RawQuery)
	// fmt.Println(result.Port())
	// fmt.Println(result.Query())

	// param := result.Query()
	// fmt.Println(param["coursename"])

	// for _, val := range param {

	// 	fmt.Println(val)
	// }

	parts := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=guruprasad",
	}

	anotherURL := parts.String()
	fmt.Println(anotherURL)
}
