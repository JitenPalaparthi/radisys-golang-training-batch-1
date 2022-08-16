package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num1 int

	var char rune // imp

	char = 19000
	char = 'A'

	fmt.Println("Value and type of num1", num1, reflect.TypeOf(num1))
	fmt.Println("Value and type of char", char, reflect.TypeOf(char))

	var (
		a, b, c int
	)

	fmt.Println(a, b, c)

	//var num2 int =100
	var (
		num2, num3 int = 100, 200
	)

	fmt.Println(num2, num3)

	var age1 = 100 // no need to mention the type here

	fmt.Println(age1, reflect.TypeOf(age1))

	// shorthand declaration of variables

	num4 := 100 //int
	fmt.Printf("Value : %v Type: %T", num4, num4)
	//var num5 float32 = 100.05
	num5 := 100.05 // float64
	fmt.Printf("\nValue : %v Type: %T", num5, num5)

	var1 := false // bool
	fmt.Printf("\nValue : %v Type: %T", var1, var1)

	str1 := "Hello World" // string
	fmt.Printf("\nValue : %v Type: %T", str1, str1)

	str2 := 'A' //rune
	fmt.Printf("\nValue : %v Type: %T", str2, str2)

	var (
		no, name, age, isMarried = 101, "Jiten", 38, true
	)

	fmt.Println(no, name, age, isMarried)

}

// Numbers
//	int8,int16,int32,int64
//	uint8,uint16,uint32,uint64
//  int --> eithr 4bytes or 8byTypeOf
//  float 32, float 64

// bool --> true/false
// byte --> 1 byte .
// rune - 4 bytes // used to store unicode chars
// string --> collection of chars. Strings are immutable

// complex , complex64, complex128 --> real and imaginary value

// builtin

// type inference
// numbers --> 0
// bool ->false
// string -> empty string
