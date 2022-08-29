package main

import (
	"demo/shape"
	"fmt"
)

func main() {
	a1 := shape.AreaOfRect()
	p1 := shape.PerimeterOfRect()
	fmt.Println(a1, p1)
}
