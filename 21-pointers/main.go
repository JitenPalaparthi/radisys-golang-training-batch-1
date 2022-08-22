package main

import "fmt"

func main() {

	var num1 int = 100
	var ptr1 *int
	var ptr2 **int

	ptr1 = &num1
	ptr2 = &ptr1

	fmt.Println("Address of num1", &num1)
	fmt.Println("Address of num1 through ptr", ptr1)

	fmt.Println("Value of num1", num1)
	fmt.Println("Value of num1 through ptr", *ptr1) // value at pointer is pointed to
	*ptr1 = 200
	fmt.Println("Value of num1", num1)

	changeVal(&num1, 201)
	fmt.Println("Value of num1", num1)
	fmt.Println("Value of num1 through ptr", *ptr1)

	changeVal(ptr1, 202)
	fmt.Println("Value of num1", num1)
	fmt.Println("Value of num1 through ptr", *ptr1)
	fmt.Println("Value of num1 through ptr of ptr", **ptr2)

	changeVal(*ptr2, 203)
	fmt.Println("Value of num1", num1)
	fmt.Println("Value of num1 through ptr", *ptr1)
	fmt.Println("Value of num1 through ptr of ptr", **ptr2)
	// pointer arthimetic is forbidden
	// can do but not directly

	ptr3 := &num1 // shorthand declaration
	fmt.Println("Value of num1", num1)
	fmt.Println("Value of num1 through ptr", *ptr3)

	str1 := new(string) // new allocates memory and returns the address
	*str1 = "Hello"
	str2 := new(string)
	*str2 = " World!"
	fmt.Println(concat(str1, str2))

	//var str3, str4 *string
	str3 := new(string) // type inference ""
	str4 := new(string) // type inference ""
	fmt.Println("Concat-->", concat(str3, str4))

	// str := "Hello World"
	// var str3 *string
	// str3 = &str

}

func changeVal(num *int, val int) {
	*num = val
}

func concat(str1, str2 *string) string {
	return *str1 + *str2
}
