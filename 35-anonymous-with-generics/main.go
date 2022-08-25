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

	xyz1 := calc(10, 20, xyz[int])
	fmt.Println("xyz1:", xyz1)

	sub2 := calc(10.45, 5.67, func(a, b float64) float64 {
		return a - b
	})

	fmt.Println("sub2:", sub2)

}

func xyz[T Types](a, b T) T {
	return a + b/a - b
}

func calc[T Types](a, b T, f func(T, T) T) T {
	return f(a, b)
}

type Types interface {
	int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | int | float32 | float64
}

// type Calc[T Types, S string] struct {
// 	A, B T
// 	C, D S
// }
