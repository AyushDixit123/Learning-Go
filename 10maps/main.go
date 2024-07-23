package main

import "fmt"

func main() {
	//[key]value
	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["ja"] = "java"

	fmt.Println("list of all languages", languages)
	//not comma seperated values

	//delete values

	delete(languages, "js")

}
