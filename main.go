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

// Value receiver
func (p person) tryUpdateFirstName(lastName string) {
	p.lastName = lastName
}

// Pointer receiver
// The * is a  type description - it means we are working with a pointer to a person
func (p *person) updateFirstName(firstName string) {
	// (*arg) This is an operator - it means we want to manipulate the value the pointer is referencing
	//var p *person
	(*p).firstName = firstName
	//var p *person
	fmt.Println("updated first name", p) //updated first name &{Jimmy Smith {jim@example.com 12345}}
	//even not using the * operator in front of the p, it prints with & symbol
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

	jim.print() //firstName:Jim lastName:Smith contactInfo:{email:jim@example.com zipCode:12345}}
	jim.tryUpdateFirstName("Jimmy")
	jim.print() //firstName:Jim lastName:Smith contactInfo:{email:jim@example.com zipCode:12345}}

	// the & operator is used to get the address of a variable
	JimmyPointer := &jim
	fmt.Println("var:", jim)              //var: {Jim Smith {jim@example.com 12345}}
	fmt.Println("pointer:", JimmyPointer) //pointer: &{Jim Smith {jim@example.com 12345}}
	//var JimmyPointer *person
	JimmyPointer.updateFirstName("Jimmy")
	//var jim person
	jim.print() //{firstName:Jimmy lastName:Smith contactInfo:{email:jim@example.com zipCode:12345}}

	//var jim person
	jim.updateFirstName("Bob")
	// even with miss match types, it works
	// if the method declare the receiver as a pointer, it will work even calling the method on a value jim

	jim.print() //firstName:Bob lastName:Smith contactInfo:{email:jim@example.com zipCode:12345}}

}
