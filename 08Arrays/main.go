package main

import (
	"fmt"
)

func main() {

	var fruitList [4]string // in arrays it is compulsion to provide the number of how much data is going to come in  i.e. have to provide it with a number

	fruitList[0] = "apple"
	fmt.Println(fruitList)
	//here in array with you will find apple in zero index followed by blank spaces indicating no data is present on these indexes
	fmt.Println(len(fruitList)) //so it strctily gives four blocks of memory to array irrespective of data is entered or not

	//In Go, array lengths must be constant values known at compile time. You cannot define the length of an array using a variable. Instead, you can use a slice, which can be dynamically sized.
}
