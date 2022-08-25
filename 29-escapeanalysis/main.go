package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

// does not escape : stored in stack
// escapes to heap: compiler has decided to create in heap memory
// move to heap : at runtime a varialbe will be moved to heap
var num1 int = 100
var sum *int

func main() {
	//var num1 int = 100
	const max int = 9999
	var num2 any = num1

	//var num3 int = num2.(int) // type assertion
	println("num1:", num1, "max:", max, "num2:", num2)
	//fmt.Println(time.Now())
	slice1 := make([]string, 1000000000)
	for i := 0; i < 1000000000; i++ {
		//slice1 = append(slice1, "Hello World!")
		slice1[i] = "Hello World!"
	}
	println(slice1)

	sum1 := add(10, 12, 13, 14, 15)
	println("sum1:", sum1)

	sum = &num1
	sum2 := addPtr(10, 12, 13, 14, 15)
	println("sum2:", *sum2)
	var sum3 *int
	var num4 int = 0
	sum3 = &num4

	addPtrArg(sum3, 10, 12, 13, 14, 15)
	println("sum3:", *sum3)

	slice2 := generate(100)
	println(slice2)

	t1 := new(T1)
	t1.ID = 100
	t1.slice = []int{10, 20, 30}
	println(t1)
	//fmt.Println(time.Now())

	bytes, _ := json.Marshal(t1)

	//t2 := new(T1)
	//var t2 T1
	t2 := &T1{}
	json.Unmarshal(bytes, t2)
	fmt.Printf("%d %v \n", t2.ID, t2.slice)
}

func add(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func addPtr(nums ...int) *int {
	sum := new(int) // dangling ptr
	for _, v := range nums {
		*sum += v
	}
	return sum
}

func addPtrArg(sumout *int, nums ...int) {
	for _, v := range nums {
		*sumout += v
	}
}

func generate(n int) []int {
	arr := make([]int, 0)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(10000))
	}
	return arr
}

type T1 struct {
	ID    int
	slice []int
}
