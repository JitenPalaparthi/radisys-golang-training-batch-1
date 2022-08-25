package main

import "fmt"

func main() {
	// var ishape IShape
	//ishape = &Rect{L: 100, B: 102.45}
	// Area(ishape)
	// Perimeter(ishape)

	// ishape = &square
	// Area(ishape)
	// Perimeter(ishape)

	s1 := New(&Rect{L: 100, B: 102.45})
	s1.Area()
	s1.Perimeter()

	var square Square = 25.25
	s2 := New(&square)
	s2.Area()
	s2.Perimeter()

}

// func Area(ishape IShape) {
// 	fmt.Println("Area:", ishape.Area())
// }

// func Perimeter(ishape IShape) {
// 	fmt.Println("Perimeter:", ishape.Perimeter())
// }

type IShape interface {
	IArea
	IPerimeter
	//Error
}

type IArea interface {
	Area() float32
}

type IPerimeter interface {
	Perimeter() float32
}

// type Error interface {
// 	OnError()
// }

type Rect struct {
	L, B float32
}

func (r *Rect) Area() float32 {
	return r.L * r.B
}

func (r *Rect) Perimeter() float32 {
	return 2 * (r.L + r.B)
}

type Square float32

func (s *Square) Area() float32 {
	return float32(*s) * float32(*s)
}

func (s *Square) Perimeter() float32 {
	return 4 * float32(*s)
}

// ioc
type Shape struct {
	iShape IShape
}

func New(ishape IShape) *Shape {
	return &Shape{iShape: ishape}
}

func (s *Shape) Area() {
	fmt.Println("Area:", s.iShape.Area())
}

func (s *Shape) Perimeter() {
	fmt.Println("Perimeter:", s.iShape.Perimeter())
}

// add two more shapes
// cuboid, triangle
// implement IShape for both

// create a slice or array
// slice must be as follows []IShape
// create two or three instances of each shape .. rect,square, triangle, cuboid
// use range loop and call Area and Perimeter methods
