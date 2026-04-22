package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Smith",
		contact: contactInfo{
			email:   "jim@example.com",
			zipCode: 12345,
		},
	}
	fmt.Println(jim)       // {Jim Smith {jim@example.com 12345}}
	fmt.Printf("%+v", jim) // {firstName:Jim lastName:Smith contact:{email:jim@example.com zipCode:12345}}
}
