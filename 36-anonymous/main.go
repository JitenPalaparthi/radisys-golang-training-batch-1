package main

import (
	"fmt"
)

type calcFunc func(int, int) int

func main() {

	var funcMap map[string]calcFunc
	substract := func(a int, b int) int {
		return a - b
	}
	a, b := 20, 10
	funcMap = make(map[string]calcFunc)
	funcMap["add"] = addition
	funcMap["sub"] = substract
	funcMap["mul"] = func(i1, i2 int) int {
		return i1 * i2
	}
	funcMap["div"] = func(a1, b1 int) int {
		return (a1 / b1)
	}
	funcMap["div-with-constant"] = division(2) // just to make sure to modify a function to add additional parameter

	for k, f := range funcMap {
		if k != "div-with-constant" {
			fmt.Println(k, ":", f(a, b))
		} else {
			func1 := f(a, b)
			fmt.Println(k, ":", func1)
		}
	}

}

func addition(a, b int) int {
	return a + b
}

func division(c int) calcFunc {
	return func(a1, b1 int) int {
		return (a1 / b1) + c
	}
}

// calcFunc add it to slice
// call each item in the slice and execute each func
