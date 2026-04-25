package main

import "fmt"

func goPointerGotChas() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	fmt.Println("mySlice:", mySlice) //mySlice: [Hi There How Are You]

	updateSlice(mySlice)
	fmt.Println("mySlice:", mySlice) //mySlice: [Bye There How Are You]
	// Even not passing a pointer it mutates the original slice~ 🤯

}

func updateSlice(s []string) {
	s[0] = "Bye"
}
