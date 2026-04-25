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

func (p person) tryUpdateFirstName(lastName string) {
	p.lastName = lastName
}

// The * is a  type description - it means we are working with a pointer to a person
func (p *person) updateFirstName(firstName string) {
	// (*arg) This is an operator - it means we want to manipulate the value the pointer is referencing
	(*p).firstName = firstName
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
	JimmyPointer.updateFirstName("Jimmy")
	// jim.updateFirstName("Jimmy")
	jim.print() //{firstName:Jimmy lastName:Smith contactInfo:{email:jim@example.com zipCode:12345}}

}
