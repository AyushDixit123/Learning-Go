package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("please enter the pizza rating between 1 to 5")
	numrating := bufio.NewReader(os.Stdin)

	input, _ := numrating.ReadString('\n')

	totalratee, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println("new ratings", totalratee+1)
	}

}
