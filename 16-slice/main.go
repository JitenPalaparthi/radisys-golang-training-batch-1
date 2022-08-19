package main

import "fmt"

func main() {

	slice1 := []int{10, 12, 123, 2131, 54, 45435, 453, 6, 453, 35, 343, 453, 12, 13}

	slice2 := slice1 // reference copy

	slice2[0] = 100

	fmt.Println(slice1)

	slice3 := slice1[:6] // slice1[0:6]

	slice3[0] = 1000
	slice3 = append(slice3, 99999)

	fmt.Println("Slice3", len(slice3), cap(slice3), slice3)
	fmt.Println("Slice1", len(slice1), cap(slice1), slice1)
	fmt.Println("Slice2", len(slice2), cap(slice2), slice2)

	slice4 := slice1[5:10]
	fmt.Println("Slice4", len(slice4), cap(slice4), slice4)
	slice4[0] = 66666
	fmt.Println("Slice1", len(slice1), cap(slice1), slice1)
	slice5 := slice1[8:]
	fmt.Println("Slice5", len(slice5), cap(slice5), slice5)

	slice6 := make([]int, len(slice1))

	// var slice6 []int

	// for i, v := range slice1 {
	// 	slice6[i] = v
	// }

	copy(slice6, slice1) // deep copy

	slice6[0] = 111111

	fmt.Println("Slice1", len(slice1), cap(slice1), slice1)
	fmt.Println("Slice6", len(slice6), cap(slice6), slice6)

}
