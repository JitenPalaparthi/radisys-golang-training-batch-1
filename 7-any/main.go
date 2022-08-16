package main

import (
	"fmt"
	"reflect"
)

// any is nothing but interface{}
// rune is nothing but int32
// byte is nothing but uint8

// empty interface
// type any = interface{}
func main() {

	//var val interface{} // this is called as empty interface. This is capable of holding any kind of value
	var val any // interface{}

	val = 100
	fmt.Println("Value:", val, "Type:", reflect.TypeOf(val))
	// type casting --> v = T(v)
	// type assertion-> v = v.(T)
	var num1 int = val.(int) // here type casting does not work. There is a concept i.e type assertion
	fmt.Println("Value of num1:", num1)

	val = false
	fmt.Println("Value:", val, "Type:", reflect.TypeOf(val))
	var bool1 bool = val.(bool)
	fmt.Println("Value of bool1:", bool1)

	val = 100.100
	fmt.Println("Value:", val, "Type:", reflect.TypeOf(val))

	val = "Hello World"
	fmt.Println("Value:", val, "Type:", reflect.TypeOf(val))

	val = 'A'
	fmt.Println("Value:", val, "Type:", reflect.TypeOf(val))

	val = byte(65) // byte is nothing but uint8
	fmt.Println("Value:", val, "Type:", reflect.TypeOf(val))

	//fmt.Println(val)

	var a any = 100 // int
	var b int = 200
	var c float32 = 300.003

	d := a.(int) + b + int(c)
	fmt.Println("a+b+c:", d, "type:", reflect.TypeOf(d))
	e := float32(a.(int)) + float32(b) + c
	fmt.Println("a+b+c:", e, "type:", reflect.TypeOf(e))

}

// a,b,c
// a must be any
// b int
// c float32
//print (a+b+c)
