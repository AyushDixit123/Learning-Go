// handling time in golang
package main

//hello
import (
	"fmt"
	"time"
)

func main() {

	presentTime := time.Now() // gives current time
	fmt.Println(presentTime)

	//changing format of time
	fmt.Println(presentTime.Format("01-02-2006")) // this date will always be used as standard date for formatting
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	fmt.Println(presentTime.Format("15:04:05"))

	//creating time from valueswe are manually entering
	createdDate := time.Date(2017, time.August, 24, 14, 41, 0, 0, time.Local)
	fmt.Println(createdDate)
}
