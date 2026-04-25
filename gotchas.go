package main

import "fmt"

func goPointerGotChas() {

	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	fmt.Println("mySlice:", mySlice) //mySlice: [Hi There How Are You]

	updateSlice(mySlice)
	fmt.Println("mySlice:", mySlice) //mySlice: [Bye There How Are You]
	// Even not passing a pointer it mutates the original slice~ 🤯

	name := "Bill"
	updateValue(name)
	fmt.Println(name)   // still Bill 🤯
	fmt.Println(&name)  // memory address of the name variable
	fmt.Println(*&name) // the value  Bill again 🤯

}

func updateSlice(s []string) {
	s[0] = "Bye"
}

func updateValue(n string) {
	n = "Alex"
}
