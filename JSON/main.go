package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("This is the JSON program")
	// encodeJSON()
	DecodeJson()
}

func encodeJSON() {
	lcoCourse := []Course{
		{"ReactJs", 299, "lco", "Guruprasad", []string{"WEB , DEV"}},
		{"Android", 99, "lco", "Guruprasad233", []string{"FullStack , DEV"}},
		{"Spring", 2299, "lco", "Guruprasad123", nil},
	}

	finalJSON, err := json.MarshalIndent(lcoCourse, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJSON)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
                "name": "ReactJs",
                "price": 299,
                "platform": "lco",
                "tags": ["WEB , DEV"]
    }

	`)

	var lcoCourse Course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v  and value is %v and Type is: %T\n", k, v, v)
	}

}
