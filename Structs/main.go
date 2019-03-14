package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	alex := person{
		firstName: "Oladiran",
		lastName:  "Segun",
		contact: contactInfo{
			email:   "sheghunoladiran9@gmail.com",
			zipCode: 900211,
		},
	}

	alex.updateName("Suraju")
	// alex.print()
}

func (p *person) updateName(firstName string) {

	p.firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
