package main

import "fmt"

func main() {
	loginCount := 23

	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else {
		result = "name"
	}
	fmt.Println(result)
	// other format
	if num := 3; num < 10 {
		fmt.Println("ok")
	} else {
		fmt.Println("nx")
	}

}
