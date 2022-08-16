package main

import (
	"fmt"
	"reflect"
)

func main() {
	// float32, float64
	// the below both real and imaginary are float64 and float64
	// c1 is complex128
	c1 := complex(12.13, .5)
	fmt.Println("Value:", c1, "Type:", reflect.TypeOf(c1))
	var (
		r1, i1 float32 = 12.13, .5
		r2, i2 float64 = 12.13, .5
	)
	c2 := complex(r1, i1)
	fmt.Println("Value:", c2, "Type:", reflect.TypeOf(c2))
	c3 := complex(r2, i2)
	fmt.Println("Value:", c3, "Type:", reflect.TypeOf(c3))

	c4 := 12.13 + .5i // another way of declaring complex number
	fmt.Println("Value:", c4, "Type:", reflect.TypeOf(c4))

	// arithemetic operations on complex numbers

	fmt.Println("Addition of complex numbers:", c1+c3)
	fmt.Println("Substraction of complex numbers:", c1-c3)
	fmt.Println("Multiplication of complex numbers:", c1*c3)
	fmt.Println("Divide of complex numbers:", c1/c3)

}
