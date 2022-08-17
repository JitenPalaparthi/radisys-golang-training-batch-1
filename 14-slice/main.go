package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// slice is an expandable array
// slice is a reference type

// make--> to create slice, map  or chan
// nil can be checked on pointers, slice, map, chan , interface
func main() {
	var slice1 []int // declare a slice

	if slice1 == nil {
		fmt.Println("Slice is nil because it is not instantiated")
	}
	slice1 = make([]int, 5) // instantiate a slice
	fmt.Println(slice1)

	if slice1 == nil {
		fmt.Println("Slice is nil because it is not instantiated")
	}
	for i := 0; i < len(slice1); i++ {
		slice1[i] = rand.Intn(1000)
	}
	fmt.Println(slice1)

	// shorthand declaration
	slice2 := []int{10, 15, 14, 13, 15} // slice with values
	fmt.Println(slice2)

	slice3 := make([]int, 1) // length = 5 , capacity = 10
	fmt.Println(slice3, "Len:", len(slice3), "Cap:", cap(slice3))
	// slice3[5] = 100
	// slice3[6] = 101
	fmt.Println(slice3, "Len:", len(slice3), "Cap:", cap(slice3))
	slice3[0] = 500
	//slice3 = append(slice3, 500) // to add a new element to the slice more than its length
	slice3 = append(slice3, 501) // to add a new element to the slice more than its length
	slice3 = append(slice3, 502) // to add a new element to the slice more than its length
	slice3 = append(slice3, 503) // to add a new element to the slice more than its length
	slice3 = append(slice3, 504) // to add a new element to the slice more than its length
	slice3 = append(slice3, 505) // to add a new element to the slice more than its length
	slice3 = append(slice3, 505, 506, 507, 508, 509, 510)
	fmt.Println(slice3, "Len:", len(slice3), "Cap:", cap(slice3))
	if err := changeVal(slice3, 6, 605); err != nil {
		fmt.Println(err)
	}
	fmt.Println(slice3, "Len:", len(slice3), "Cap:", cap(slice3))
}

func changeVal(arr []int, i uint, v int) error {
	if int(i) >= len(arr) {
		return errors.New("invalid index")
	}
	arr[i] = v
	return nil
}

// reverse a slice
// deep copy a slice
// by default slice does shallow copy because of reference type
