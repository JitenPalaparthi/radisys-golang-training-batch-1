package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr1 := [2][2][3]int{{{1, 2, 3}, {4, 5, 6}}, {{7, 8, 9}, {10, 11, 12}}}
	fmt.Println(arr1)

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			for k := 0; k < len(arr1[i][j]); k++ {
				print(arr1[i][j][k], " ")
			}
			println()
		}
		println()
	}
	// type of an array
	arr2 := [5]int{49, 34, 54, 67, 89}
	fmt.Println("Values of arr2", arr2, "Type of arr2", reflect.TypeOf(arr2))
	arr3 := [5]int{}
	arr3 = arr2
	fmt.Println("Values of arr3", arr3, "Type of arr2", reflect.TypeOf(arr3))

	err := changeVal(arr3, 6, 334) // call by value
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(arr3)
	}

	max, min, _ := MinAndMax(arr3)
	//max, min, _ := MinAndMax2(arr3)
	fmt.Println("Max:", max, "Min:", min)

	max1, _, _ := MinAndMax(arr3)
	fmt.Println("Max:", max1)

	// what are values of arr3 ?

}

// nil for interfaces, interface{}, pointers, slices, map , chan
func changeVal(arr [5]int, i uint, v int) (err error) {
	if int(i) >= len(arr) {
		//return errors.New("invalid index")
		return fmt.Errorf("invalid index")
	}
	arr[i] = v
	fmt.Println("inside function:", arr)
	return nil
}

func MinAndMax(arr [5]int) (max int, min int, err error) {

	if len(arr) == 0 {
		return 0, 0, fmt.Errorf("invalid array")
	}
	max = arr[0]
	min = arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	return max, min, nil
}

func MinAndMax2(arr [5]int) (int, int, error) {

	if len(arr) == 0 {
		return 0, 0, fmt.Errorf("invalid array")
	}
	max := arr[0]
	min := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	return max, min, nil
}

// stack or heap
// escape analysis

// go.dev

// Effective Go
