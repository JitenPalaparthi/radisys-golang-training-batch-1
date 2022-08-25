package main

import "fmt"

func main() {

	c1 := Calc[int]{A: 100, B: 90}
	fmt.Println("c1 Add:", c1.Add())
	fmt.Println("c1 Sub:", c1.Sub())
	fmt.Println("c1 Multipy:", c1.Multipy())
	fmt.Println("c1 Divide:", c1.Divide())
	fmt.Println()
	c2 := Calc[float32]{A: 100.23, B: 90.34}
	fmt.Println("c2 Add:", c2.Add())
	fmt.Println("c2 Sub:", c2.Sub())
	fmt.Println("c2 Multipy:", c2.Multipy())
	fmt.Println("c2 Divide:", c2.Divide())
	fmt.Println()

	c3 := Calc[float64]{A: 100.23, B: 90.34}
	fmt.Println("c3 Add:", c3.Add())
	fmt.Println("c3 Sub:", c3.Sub())
	fmt.Println("c3 Multipy:", c3.Multipy())
	fmt.Println("c3 Divide:", c3.Divide())

}

type Types interface {
	int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | int | float32 | float64
}

type Calc[T Types] struct {
	A, B T
}

func (c Calc[T]) Add() T {
	return c.A + c.B
}

func (c Calc[T]) Sub() T {
	return c.A - c.B
}

func (c Calc[T]) Multipy() T {
	return c.A * c.B
}

func (c Calc[T]) Divide() T {
	return c.A / c.B
}
