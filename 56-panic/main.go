package main

import "fmt"

func main() {
	defer fmt.Println("end of main")
	//v := 0
	//fmt.Println(100 / v)
	fn := new(string) // its a pointer
	*fn = "Jiten"
	var ln *string //nil pointer
	fullName(fn, ln)
	greet()
}

func greet() {
	fmt.Println("Hello World")
}
func fullName(fn, ln *string) {
	defer fmt.Println("full name fun ends")
	defer recoverMe() // when there is a panic, all defers are executed.
	// hence make recover logic in defer
	fmt.Println("calling fullname func")
	//func() {
	if fn == nil {
		panic("Firstname cannot be nil")
	}
	if ln == nil {
		panic("LastName cannot be nil")
	}
	//}()
	fmt.Println(fn, ln)
}

func recoverMe() {
	if r := recover(); r != nil {
		fmt.Println("recover me-->", r)
	}
}
