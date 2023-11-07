package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")

	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDatee := time.Date(2004, time.November, 11, 23, 23, 0, 0, time.UTC)

	fmt.Println(createdDatee.Format("01-02-2006 Monday "))

}
