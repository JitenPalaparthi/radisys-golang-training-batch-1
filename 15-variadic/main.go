package main

import "fmt"

func main() {

	// fmt.Println("How", "Are", "You", "Doing")

	fmt.Println("Sum:", sumOf(10, 20, 30, 40, 50, 60, 160, 170, 180, 190))
	fmt.Println("Sum:", sumOf(10, 20))
	fmt.Println("Sum:", sumOf())

	slice := []int{10, 20, 30, 40, 50, 60, 160, 170, 180, 190}
	fmt.Println("Sum:", sumOf(slice...)) // ... must be put next to slice

	arr := [5]int{10, 11, 12, 13, 14}
	fmt.Println("Sum:", sumOf(arr[:]...)) // ... must be put next to slice

}

// Variadic parameter must be the last parameter
// Varaidic parameter , you can pass zero arguments or many arguments
// you can pass a slice as a argument to variadic parameter but you need to use ... eclipse symbol append to the slice
func sumOf(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// insert operation in slices
func insert(arr []int, i, v int) {
	if i != 0 {
		arr = append(arr[:i], arr[i-1:]...)
		arr[i] = v
	} else {
		arr = append(arr[:0], arr[i:]...)
		arr[i] = v
	}
}
