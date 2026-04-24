package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

//	func (p person) updateFirstName(firstName string) {
//		p.firstName = firstName
//	}

func (pointerToPerson *person) updateFirstName(firstName string) {
	(*pointerToPerson).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Smith",
		contactInfo: contactInfo{
			email:   "jim@example.com",
			zipCode: 12345,
		},
	}
	jim.print() // name: Jim
	// jim.updateFirstName("Jimmy")

	JimmyPointer := &jim
	JimmyPointer.updateFirstName("Jimmy")
	// jim.updateFirstName("Jimmy")
	jim.print() // name: Jimmy
}
