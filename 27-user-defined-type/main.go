package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	var num1 myInt
	var num2 Integer = 100

	num1 = myInt(num2)
	fmt.Println("Value of num1:", num1, "Type of num1", reflect.TypeOf(num1))
	fmt.Println("Value of num2", num2, "Type of num2", reflect.TypeOf(num2))

	str1 := num1.ToString()
	fmt.Println("str1:", str1, "Type of str1:", reflect.TypeOf(str1))

	str2 := myInt(num2).ToString() // num2 is int .. to call ToString method num2 is casted to myInt
	fmt.Println("str2:", str2, "Type of str2:", reflect.TypeOf(str2))

	var mymap1 myMap
	mymap1 = make(myMap)
	mymap1["Bangalore-1"] = 560086
	mymap1["Trivandrum-1"] = 690031
	fmt.Println("Value of myap1", mymap1, "Type of mymap1", reflect.TypeOf(mymap1))

	var map1 map[string]object
	map1 = make(map[string]interface{})
	//map1 = mymap1 // its a reference copy
	fmt.Println("Value of map1", map1, "Type of map1", reflect.TypeOf(map1))
	map1["Bangalore-2"] = 560031
	fmt.Println("Value of myap1", mymap1, "Type of mymap1", reflect.TypeOf(mymap1))

	if ok, err := mymap1.Delete("Bangalore-2"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("deleted?", ok)
	}
	fmt.Println("Value of myap1", mymap1, "Type of mymap1", reflect.TypeOf(mymap1))

	if ok, err := mymap1.Delete("Bangalore-2"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("deleted?", ok)
	}

	var mymap2 myMap // only declared but not instantiated
	if ok, err := mymap2.Delete("Bangalore-2"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("deleted?", ok)
	}

	if ok, err := myMap(map1).Delete("Bangalore-1"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Deleted?", ok)
	}

	fmt.Println("Value of map1", map1, "Type of map1", reflect.TypeOf(map1))
	fmt.Println("value of mymap1", mymap1, "Type of mymap1", reflect.TypeOf(mymap1))
}

type myInt int

func (mi myInt) ToString() string {
	return fmt.Sprint(mi)
}

type myMap map[string]any

func (mm myMap) Delete(key string) (bool, error) {
	if mm == nil {
		return false, errors.New("nil map")
	}
	if _, ok := mm[key]; !ok {
		return false, errors.New("no key")
	}

	delete(mm, key)
	return true, nil
}

type Integer = int

type double = float64

type long = int64

type object = any

// read/reearch in builtin standard package

// create user defined map
// insert, update, get, delete methods
// each method should have error as last parameter
// if the specified operations can't be performed the error must retunr reason.

// go.dev
// google builtin pacakge in go
