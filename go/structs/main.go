package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// type person struct {
// 	firstName string
// 	lastName  string
// 	contactInfo   contactInfo
// }

type person struct {
	firstName   string
	lastName    string
	contactInfo //same name variable as type name
}

func main() {

	//alex := person{firstName: "Alex", lastName: "Anderson"}
	//fmt.Println(alex)

	// var noor person
	// fmt.Println(noor)
	// fmt.Printf("%v", noor)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail",
			zipCode: 94000,
		},
	}

	jim.updateName("Jimmy")
	jim.print()

	pointerToJim := &jim
	pointerToJim.updateName("Jimmyyy")
	jim.print() //both ways work, go has this shortcut

}

// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName //won't change the name, cause pass by value.
// }

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName // will change the name, cause pass by reference.
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
