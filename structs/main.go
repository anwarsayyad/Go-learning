package main

import "fmt"

type contactinfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact contactinfo
}


func main() {
	alex := person{
		firstName: "Anwar",
		lastName: "Sayyad",
		contact: contactinfo{
			email: "anwar@gmail.com",
			zipCode: 59125,
		},
	}
	// we can also declare struct by person{firstName: "fitst Name", lastName: "Last name"}


	// var alex person
	// alex.firstName = "Anwar"
	// alex.lastName = "sayyad"
	// alex.contact.email = "anwarsayyad@gmail.com"
	// alex.contact.zipCode = 590065
	alex.update_firstname("Afrid")
	alex.print()
}

// reciver for the person struct
func (p person) print() {
	fmt.Printf("%+v",p)
}

func (p person) update_firstname(newFirstName string) {
	p.firstName = newFirstName
} 