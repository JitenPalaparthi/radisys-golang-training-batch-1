package main

import "fmt"

func main() {

	p1 := Person{ID: 101, Name: "Jiten", Contact: "9618558500"}
	//DisplayPerson(p1)
	p1.Display()
	str := p1.ToString()
	fmt.Println(str)

	p1.T1.Display()

	p1.T2.Display()

	var t1 T1
	t1.Display()

	var t2 T2
	t2.Display()

	Display()
}

type Person struct {
	ID                   int
	Name, Email, Contact string
	T1
	T2
}

func DisplayPerson(p Person) {
	fmt.Println("ID:", p.ID)
	fmt.Println("Name:", p.Name)
	fmt.Println("Contact Number:", p.Contact)
}

func (p Person) Display() {
	fmt.Println("ID:", p.ID)
	fmt.Println("Name:", p.Name)
	fmt.Println("Contact Number:", p.Contact)
}

func (p Person) ToString() string {
	return fmt.Sprint(p)
}

type Address struct {
	City, State string
}

func (a Address) Display() {
	fmt.Println("City:", a.City)
	fmt.Println("State:", a.State)
}

func Display() {
	fmt.Println("Hello World--from Func")
}

type T1 struct{} // empty struct .A struct does not containg any fielcs

func (t1 T1) Display() {
	fmt.Println("Hello World-- from T1")
}

type T2 struct{} // empty struct

func (t2 T2) Display() {
	fmt.Println("Hello World-- from T2")
}
