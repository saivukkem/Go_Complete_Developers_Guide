package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email    string
	postcode string
}

func main() {
	anusha := person{
		firstName: "Dakshayani",
		lastName:  "Chandu",
		contactInfo: contactInfo{
			email:    "d.chandu@email.com",
			postcode: "M1 1M",
		},
	}

	anusha.updateFirstName("Anu")
	anusha.updateLastname("Daks")
	anusha.print()
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson *person) updateLastname(newLastName string) {
	(*pointerToPerson).lastName = newLastName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
