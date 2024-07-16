package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "dnjnfnfik"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin) //reader is a pointer
	fmt.Println("enter the rating for pizza")

	// whatever the reader reads is stored in comma ok syntax variable
	// comma ok || err err

	input, _ := reader.ReadString('\n') //it shows it will keep reading until next line ius hit
	fmt.Println("thanks for rating,", input)
}
