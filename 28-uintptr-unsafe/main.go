package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var num1 int = 100

	var ptr1 *int = &num1
	fmt.Println(ptr1)

	// ptr1+= 8 // this is pointer arthemetic. Very common operation in C lang
	// 0xc00001e098 0xc00001e106

	arr1 := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("Add of 0th element", &arr1[0])
	size := unsafe.Sizeof(arr1[0])
	fmt.Println("Size of int", size)

	arrptr := uintptr(unsafe.Pointer(&arr1[0]))
	arrptr += size // performing arthemetic operation on uintptr

	val := (*int)(unsafe.Pointer(arrptr))
	fmt.Println(*val)
	arr1ptr := uintptr(unsafe.Pointer(&arr1[0]))
	fmt.Println(*(*int)(unsafe.Pointer(arr1ptr)))
	for i := 1; i < len(arr1); i++ {
		arr1ptr += size
		fmt.Println(*(*int)(unsafe.Pointer(arr1ptr)))
	}
	uintptr1 := uintptr(0xc000014370)
	val1 := (*int)(unsafe.Pointer(uintptr1))
	fmt.Println(*val1)
}
