package main

import (
	"errors"
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Len:", len(slice1), "cap:", cap(slice1), "slice1:", slice1)
	slice1, _ = insert(slice1, 5, 500)
	fmt.Println("Len:", len(slice1), "cap:", cap(slice1), "slice1:", slice1)
	slice1, _ = insert(slice1, 0, 100)
	fmt.Println("Len:", len(slice1), "cap:", cap(slice1), "slice1:", slice1)

	slice2 := []int{100}
	slice2, _ = insert(slice2, 0, 99)
	fmt.Println(slice2)

}

func insert(slice []int, i, v int) ([]int, error) {
	if slice == nil {
		return nil, fmt.Errorf("nil slice")
	}
	if i < 0 || i >= len(slice) {
		return nil, errors.New("invalid index")

	}
	if i != 0 {
		slice = append(slice[:i], slice[i-1:]...)
		fmt.Println(slice)
		slice[i] = v
		return slice, nil
	}
	if len(slice) != 1 {
		slice = append(slice[:1], slice[i:]...)
		//fmt.Println(slice)
		slice[0] = v
		return slice, nil
	}

	slice = append(slice, slice[0])
	slice[0] = v
	return slice, nil

}
