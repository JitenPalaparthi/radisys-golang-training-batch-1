package main

import (
	"fmt"
	"shapes/shape/cube"
	r "shapes/shape/rect" // r is alias .. if not you can directly call with rect.funcnames
	"shapes/shape/square"
)

func main() {
	//fmt.Println("Hello World")
	fmt.Println(r.Greet())

	a1 := r.Area(10.34, 13.67)
	fmt.Println("Area of Rect:", a1)

	p1 := r.Perimeter(10.34, 13.67)
	fmt.Println("Perimeter of Rect:", p1)

	a2 := square.Area(25.25)
	fmt.Println("Area of Square:", a2)
	p2 := square.Perimeter(25.25)
	fmt.Println("Perimeter of Square:", p2)

	a3 := cube.Area(10.34, 13.67, 12.45)

	fmt.Println("Area of Cube:", a3)

	p3 := cube.Perimeter(10.34, 13.67, 12.45)

	fmt.Println("Perimeter of Cube:", p3)
}

// to create a package
// 		1- GOPATH
// 		< Go 1.11
// bin/ pkg/ src/ --> export GOPATH=C:/Users/Jiten/Workspace/MyProject
// godeps
// 1.11 GO111MODULE=on

// 1- standard
// 2- user defined package
// 3- third party packages

// shapes
// 	shape
// 		rect.go

// shapes
// 	shape
//		cube
//			cube.go
// 		rect
// 			area.go
//			perimeter.go
//			misc.go
//      square
//			square.go

// create a simple pacakge called calc
// add
// sub
// mul
// divide
// write all these functiosn with any/interface{} so that it should accept
// float or int
