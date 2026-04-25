package main

import "fmt"

type customSlice []string

func goPointerGotChas() {

	mySlice := customSlice{"Hi", "There", "How", "Are", "You"}
	fmt.Println("mySlice:", mySlice) //mySlice: [Hi There How Are You]

	updateSlice(mySlice)
	fmt.Println("mySlice:", mySlice) //mySlice: [Bye There How Are You]
	// Even not passing a pointer it mutates the original slice~ 🤯
	mySlice.updateByValueReceiver()
	fmt.Println("mySlice:", mySlice) //mySlice: [By Value Receiver There How Are You]
	// Mutate the original slice too using value receiver
	// The slice has [length,cap,ptr to head] in memory
	//Even passing a copy of the slice, both slices have the same [ptr to head]

	name := "Bill"
	updateString(name)
	fmt.Println(name)   // still Bill
	fmt.Println(&name)  // memory address of the name variable
	fmt.Println(*&name) // the value  Bill again 🤯
	updateStringWithPointer(&name)
	fmt.Println(name) // Alex with pointer

	namePointer := &name
	fmt.Println(&namePointer)
	printPointer(namePointer)
	// even passing the pointer as a parameter, it  prints differently memory address
	// the pointers are cloned, their head is in different memory locations but the reference is the same

}

func updateSlice(s customSlice) {
	s[0] = "Bye"
}

func (s customSlice) updateByValueReceiver() {
	s[0] = "By Value Receiver"
}

func updateString(n string) {
	n = "Alex"
}
func updateStringWithPointer(n *string) {
	*n = "Alex with pointer"
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}
