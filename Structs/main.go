package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // we could just use contactInfo, but this is more readable
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	var alex person // go assigns zero value to the fields
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Println(alex)
	fmt.Printf("%+v", alex)

	jim := person{firstName: "Jim", lastName: "Party", contact: contactInfo{email: "jim@xyz.com", zipCode: 12345}}
	fmt.Printf("%+v", jim)

	jim.print()

	jim.updateName("Jimmy")

	jim.print() // the name is not updated because the receiver is a copy of the original struct

	jimPointer := &jim

	jimPointer.updateName1("Jimmy")

	jim.print() // the name is updated because the receiver is a pointer to the original struct

	jim.updateName1("JimmyBoi") // this is equivalent to the above line, but there is a type mismatch. Go automatically converts the receiver to a pointer

	jim.print()

	/*
		We would need to use pointers if we are working with value types like int, string, float, bool, etc.
		We don't need to use pointers with slices, maps, channels, pointers, functions because they are already references
	*/

}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p *person) updateName1(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
