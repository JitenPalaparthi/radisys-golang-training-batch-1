package main

import (
	"fmt"
	"unsafe"
)

func main() {

	t1 := T1{id: 101, name: "Jiten", isMarried: true, address: "Bangalore, Rajaji nagar"}
	fmt.Println("t1:", t1, "Size of t1:", unsafe.Sizeof(t1))

	t2 := T2{id: 101, name: "Jiten", isMarried: true, address: "Bangalore, Rajaji nagar"}
	fmt.Println("t2:", t2, "Size of t1:", unsafe.Sizeof(t2))

	// arr1 := [4]string{"How", "are", "you", "doing?"}

	// //size := unsafe.Sizeof(arr1[0])

	// arrptr := uintptr(unsafe.Pointer(&arr1[0]))

	// fmt.Println(*(*string)(unsafe.Pointer(arrptr)))

	// for i := 1; i < 4; i++ {
	// 	//arrptr := uintptr(unsafe.Pointer(&arr1[0]))
	// 	arrptr += 16 //size
	// 	fmt.Println(*(*string)(unsafe.Pointer(arrptr)))
	// }

}

type T1 struct {
	id        uint16 // 8 ---> 2 --6
	name      string // 16 -->16 --0
	isMarried bool   // 8  -->1  --7
	address   string // 16 -->16 --0
	ref       uint32 //     --> 4 -- 1
	canVote   bool   //     --> 1 -- 1
} //

type T2 struct {
	name      string // 16
	address   string // 16
	id        uint16 // 8   --> 2 -- 6
	isMarried bool   //     --> 1 -- 5
	canVote   bool   //     --> 1 -- 1
	ref       uint32 //     --> 4 -- 1
}
