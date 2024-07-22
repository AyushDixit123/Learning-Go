package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"apple", "tomato"}
	fmt.Printf("type of fruitlist is %T", fruitList)
	fmt.Println(len(fruitList))
	fmt.Println(fruitList)
	// the mermory in a slice is dynamic

	//adding data in slice
	fruitList = append(fruitList, "mango", "banana")
	fmt.Println(fruitList)
	// : syntax in slice is used for slicing up the collection i.e.
	// it will show data of collections after the index mentioned in collection
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:3]) //starts mfrom position 1 to 3

	highScores := make([]int, 2) // we are creating a space in memory for this variable that is a slices with size 4
	highScores[0] = 123
	highScores[1] = 345
	//highScores[2] = 312 // this will create an error due to out of scope
	highScores = append(highScores, 321)
	// this will not cause an error
	//append reallocates the memory and data is allocated, increasing performance
	sort.Ints(highScores) //sorts int slices in increasing order
}
