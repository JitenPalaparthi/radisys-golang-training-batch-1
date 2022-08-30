package main

import "fmt"

// defer defers the execution to the end of the caller
func main() {
	// defer func() {
	// 	defer fmt.Println("End of main")
	// }()

	// defer fmt.Println("call-1")
	// defer fmt.Println("call-2")
	// defer fmt.Println("call-3")
	// fmt.Println("Beginning of main")

	i := 100

	defer func(i int) {
		i = i + 100
		fmt.Println("in defer", i)
	}(i)

	i = i / 100

	fmt.Println("in main", i)
	// reverse using defer
	str := "Hello World"
	for _, v := range str {
		defer fmt.Print(string(v))
	}

}
