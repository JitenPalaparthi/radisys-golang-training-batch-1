package shape

import "fmt"

// init does not have parameters
// init does not have return type
func init() {
	fmt.Println("Calling init")
	L = 1
	B = 1
}

func init() {
	fmt.Println("Calling init again")
}

func init() {
	fmt.Println("Calling init again and again")
}

var (
	L, B float32
)

func AreaOfRect() float32 {
	return L * B
}

func PerimeterOfRect() float32 {
	return 2 * (L + B)
}
