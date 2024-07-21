package main

import "fmt"

// pointers give yu an assurity that when you are passing in variables yu are not passing variables rather their memory adresses
// pointers are refrecncde to their memory adress
func main() {
	var ptr *int     //this datatype is pointer which is responsible for holding integers into that
	fmt.Println(ptr) //intial value in a empty ptr is nil

	myNumb := 23

	var ptr1 = &myNumb // this ptr is pointer which is pointing towards or refrencing the memory adrees of this variable

	fmt.Println(ptr1)  // it shows that pointer is pointing or stored the memory adress of variable
	fmt.Println(*ptr1) // the value inside this pointer is the actual value of this variable which is accessed by putting * int front of the pointer
}
