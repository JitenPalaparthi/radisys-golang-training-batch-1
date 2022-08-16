package main

import (
	"fmt"
	"reflect"
)

func main() {
	// compiler must evaluate the value of a constant at compile time
	const PI float32 = 3.14
	fmt.Println("Value", PI, "Type:", reflect.TypeOf(PI))

	const (
		MAX, MIN = 10, 1.1
	)
	fmt.Println("Value", MAX, "Type:", reflect.TypeOf(MAX))
	fmt.Println("Value", MIN, "Type:", reflect.TypeOf(MIN))

	var num int8 = 10

	// constants must be evaluated at compile time
	const SQUARE1 = 10 * 10   // OK
	const SQUARE2 = MAX * MAX // OK

	//const SQUARE3 = num * num // NOT OK
	fmt.Println("Value", SQUARE1, "Type:", reflect.TypeOf(SQUARE1))
	fmt.Println("Value", SQUARE2, "Type:", reflect.TypeOf(SQUARE2))

	fmt.Println(num)
}
