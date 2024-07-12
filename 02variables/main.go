package main

import "fmt"

func main() {
	var username string = "ayush"
	//you always need to be use something if have declared it in your go program
	fmt.Print(username)
	//fmt.Printf("variables is of type %T", username)

	//instead of long, we have int8, int16,     int32,int64

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	//default values: go doesnt store an garbage value fot int,float- 0 string is empty

	//no var style: go dosent require var keyword to declare varaible and can be implicilty type
	// but this declartion is only allowed in a method(here main())
	numberofUsers := 3000
	fmt.Println(numberofUsers)

	// if firstletter of variable/method is capital it is Public

}
