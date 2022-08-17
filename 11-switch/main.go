package main

import "fmt"

func main() {
	num := 80
	switch {
	case num%8 == 0:
		println(num, "is divisible by 8")
		fallthrough

	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough

	case num%2 == 0:
		println(num, "is divisible by 2")

	default:
		println(num, "not divisible by 2,4,8")
	}

	var ival interface{}
	var age int8 = 100
	ival = age
	switch ival.(type) {
	case int:
		fmt.Println(ival, "is of type int")
	case int8:
		fmt.Println(ival, "is of type int8")
	case int16:
		fmt.Println(ival, "is of type int16")
	case int32:
		fmt.Println(ival, "is of type int32")
	case float64:
		fmt.Println(ival, "if of type float64")
	default:
		fmt.Println(ival, "is another type")
	}
}
