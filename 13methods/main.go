package main

// structs are same as classes
import "fmt"

func main() {
	// there is no inheritence in golang
	// using structure created
	hitesh := User{"ayush", "ayush.go.dev", true, 20}
	fmt.Println(hitesh)
	hitesh.GetStatus()

}

// u in capitaal bcz as it is like a template it is needed to be exported out
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// passing struct in arg of function makes it method
func (u User) GetStatus() {
	fmt.Println("is user active: ", u.Status)
}
