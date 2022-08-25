package main

import "fmt"

func main() {
	// simple function
	// no arguments no retunns
	func() {
		fmt.Println("Hello World!-1")
	}()

	// with one argument no return type
	func(msg string) {
		fmt.Println(msg)
	}("Hello World!-2")

	f := func(msg string) string {
		return msg
	}

	// delcared a function
	sum := func(a, b int) int {
		return a + b
	}

	// declared and called func
	sum1 := func(a, b int) int {
		return a + b
	}(10, 20)

	str := f("Hello World!-2")
	fmt.Println(str)
	fmt.Println("sum:", sum(10, 20))
	fmt.Println("sum1:", sum1)

	add1 := calc(10, 20, sum) // http handlers, middlewares
	fmt.Println("add1:", add1)

	sub1 := calc(10, 20, func(a, b int) int {
		return a - b
	})

	fmt.Println("sub1:", sub1)

	mul1 := calc(10, 20, func(a, b int) int {
		return a * b
	})

	fmt.Println("mul1:", mul1)

	div1 := calc(10, 20, func(a, b int) int {
		return a / b
	})

	fmt.Println("div1:", div1)

	xyz1 := calc(10, 20, xyz)
	fmt.Println("xyz1:", xyz1)

}

func xyz(a, b int) int {
	return a + b/a - b
}

func calc(a, b int, f func(int, int) int) int {
	return f(a, b)
}
