package main

import (
	"fmt"
	"shapes/shape/cube"
	"shapes/shape/rect"
)

func main() {
	l1 := new(float32)
	b1 := new(float32)
	*l1 = 12.34
	*b1 = 13.47
	r1 := &rect.Rect{L: l1, B: b1}

	fmt.Println("Area of Rect:", r1.Area())
	fmt.Println("Perimeter of Rect:", r1.Perimeter())

	fmt.Println(r1.A)
	fmt.Println(r1.P)

	c1 := new(cube.Cube)
	(*c1).L = 10.25
	c1.B = 12.34
	c1.H = 14.56

	fmt.Println("Area of Cube:", c1.Area())
	fmt.Println("Perimeter of Cube:", c1.Perimeter())

	*l1, *b1 = 10.23, 14.45
	if r2, err := rect.New(l1, b1); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Area of Rect:", r2.Area())
		fmt.Println("Perimeter of Rect:", r2.Perimeter())
	}
}
