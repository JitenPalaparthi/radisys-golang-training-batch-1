package main

import "fmt"

func main() {

	var p1 Person
	var a1 Address

	a1.Line1 = "1St Lane"
	a1.Street = "Kesavdasapuram"
	a1.City = "trivandrum"
	a1.State = "Kerala"
	a1.Country = "India"
	a1.PIN = "590031"

	p1.ID = 100
	p1.Name = "Jiten"
	p1.Address = a1
	p1.Email = "Jitenp@outlook.com"
	p1.Contact = "9618558500"

	fmt.Println(p1)

	p2 := Person{ID: 101, Name: "Ravi", Address: Address{Line1: "Kamat Street", City: "Bangalore"}, Email: "Ravi.K@Gmail.Com", Contact: "99991100111"}
	fmt.Println(p2)
}

type Person struct {
	ID      int
	Name    string
	Address Address // can give same name of type as a filed
	Email   string
	Contact string
}

type Address struct {
	Line1, Street, City, State, Country, PIN string
}

// 1- struct
// 2- user defined type from any existing type
// Methods are called on types. Methods can be created on types
