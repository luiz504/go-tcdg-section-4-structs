package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// very easy to fuck up if need to change the struct fields order, find some lint config to help with that.
	// alex := person{
	// 	"Alex",
	// 	"Smith",
	// }

	// alex := person{
	// 	firstName: "Alex",
	// 	lastName:  "Smith",
	// }

	var alex person
	fmt.Println(alex)       // { } // Zero Value
	fmt.Printf("%+v", alex) // {firstName:"" lastName:""}

	alex.firstName = "Alex"
	alex.lastName = "Smith"
	fmt.Println(alex)       // {Alex Smith}
	fmt.Printf("%+v", alex) // {firstName:"Alex" lastName:"Smith"}

}
