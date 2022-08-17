package main

import "fmt"

//
func main() {
	// var num int
	// const max int
	var arr1 [5]int
	fmt.Println(arr1) // type inference works here. So arr1 [0,0,0,0,0]
	arr1[0] = 100
	arr1[1] = 101
	arr1[2] = 102
	arr1[3] = 103
	arr1[4] = 104
	//arr[5] = 105
	//fmt.Println(arr1)
	sum := 0
	for _, v := range arr1 {
		sum = sum + v
	}

	println("Sum of arr1:", sum)
	// creating and assigning values to a array
	var arr2 [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println(arr2)

	arr3 := [5]int{101, 103, 105, 107, 109} // short hand declaration
	fmt.Println(arr3)

	// arr4 the size is evaluated by the compiler at compile time
	arr4 := [...]int{123, 43, 12, 5, 12, 76, 345, 68, 45, 787, 454, 9, 54, 245}
	fmt.Println(arr4)

	fmt.Println("Length of arr4", len(arr4))
	fmt.Println("Capacity of arr4", cap(arr4))

	max := arr4[0]
	for i := 1; i < len(arr4); i++ {
		if arr4[i] > max {
			max = arr4[i]
		}
		print(arr4[i], " ")
	}

	fmt.Println("Max value in arr4:", max)

}

// reverse the array
// find min number in an array
// find prime numbers in an array
