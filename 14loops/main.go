package main

import "fmt"

func main() {
	days := []string{"sun", "mon", "tue", "fri"}

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }
	//range autmocially loops through the slice(in this case)
	for i := range days {
		fmt.Println(days[i])
	} //i is an index

	//for each loop returns both index and value
	for index, day := range days {
		fmt.Printf("index is %v and value is %v", index, day)
	}
	//goto
	rouge := 1

	for rouge < 10 {
		rouge++

		if rouge == 2 {
			goto lco
		}
	}
lco:
	fmt.Println("laeraning")

}
