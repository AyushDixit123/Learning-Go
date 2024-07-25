package main

import "fmt"

func main() {
	result := adder(2, 3)
	fmt.Println(result)

	nums := pro(2, 3, 4, 5, 67, 8)
	fmt.Print("nums", nums)
	// handling multiple returns from functions like comma ok syntax
	anothernum, mess := prores(2, 3, 4, 5)
	println(anothernum)
	println(mess)
}

// no function can be declared in a function
func adder(input1 int, input2 int) int {
	return input1 + input2
}

// when no of arguments is unkwon
func pro(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}

// returning multiple things from function is similar to comma ok syntax
func prores(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "ayjj"
}
