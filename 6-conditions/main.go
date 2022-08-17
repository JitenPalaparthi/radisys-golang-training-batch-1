package main

import (
	"fmt"
	"math/rand"
)

type char = rune // creating a type using underlining type

type charStr = char // creating a userdefined type from another userdefined type

var (
	a, b, c int
)

func main() { //{ func, if else, for, switch, select

	var age uint8 = 18

	if age >= 18 {
		fmt.Println("age is ", age, ".So eligible for vote")
	} else {
		fmt.Println("age is ", age, ".So not eligible for vote")
	}

	var gender charStr = 'F' // internally it is 70 // it is a rune but custom type created at the top

	if age >= 21 {
		fmt.Println("age is ", age, " and gender is ", string(gender), ".So eligible for marriage")
	} else if age >= 18 && gender == 'F' {
		fmt.Println("age is ", age, " and gender is ", string(gender), ".So eligible for marriage")
	} else {
		fmt.Println("Not eligible for marriage")
	}

	// conversion either Spring or strconv
	a, b, c = 49, 49, 49
	if a > b && a > c {
		fmt.Println(a, "a is bigger")
	} else if b > a && b > c {
		fmt.Println(b, "b is bigger")
	} else if c > a && c > b {
		fmt.Println(c, "c is bigger")
	}

	for i := 1; i <= 10; i++ {
		if num := rand.Intn(100); num%2 == 0 {
			fmt.Println(num, "is an even number")
		} else {
			fmt.Println(num, "is an odd number")
		}
	}
}

// if a,b,c are equal say equal
// if a ,b are equal but c is smaller then say a and b are equal and bigger
