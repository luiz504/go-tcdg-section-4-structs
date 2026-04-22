package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	// alex := person{
	// 	"Alex",
	// 	"Smith",
	// }

	alex := person{
		firstName: "Alex",
		lastName:  "Smith",
	}

	fmt.Println(alex)         // {Alex Smith}
	fmt.Printf("%+v\n", alex) // {firstName:Alex lastName:Smith}

}
