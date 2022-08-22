package main

import "fmt"

func main() {

	var p1 Person
	p1.ID = 100
	p1.Name = "Jiten"
	p1.Email = "JitenP@Outlook.Com"
	p1.Contact = "9618558500"
	p1.Line1 = "1st Lane"
	p1.City = "Trivandrum"
	p1.Street = "Kesavadasa puram"
	p1.State = "Kerala"
	p1.Country = "India"
	p1.PIN = "590094"
	p1.Address.ID = 111
	fmt.Println(p1)

}

type Person struct {
	ID                   int
	Name, Email, Contact string
	//Addr                 Address
	Address // promoted field
	// if a filed is a promoted filed, members in that struct can directly be called.
	// in this exmaple p1.City can be accessed prvided by p1 is a variable of Person
}

type Address struct {
	ID                                       int
	Line1, City, Country, Street, State, PIN string
}
