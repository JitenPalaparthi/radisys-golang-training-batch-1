package main

import "fmt"

// range loop is used in three ways
// 1- arrays , slices --> index,value
// 2- maps --> key,value
// 3- channles --> value
func main() {
	str := "Hello World!"
	for i, v := range str {
		println("index: ", i, " Value: ", string(v))
	}

	for _, v := range str {
		println(" Value: ", string(v))
	}

	for i, _ := range str {
		println("index: ", i)
	}

	count := 1
	for range str {
		fmt.Println(str+"->", count)
		count++
	}
}

// Do string reverse using range loop
