package main

import "fmt"

func main() {
	fmt.Println("sumOf:", sumOf(10, 10.11, 12, 14.56, 45.65))     // any/interface{} implementation
	fmt.Println("sumOf:", sumOf(10, 10.11, true, "Hello", 45.65)) // any/interface{} implementation

	fmt.Println("sum:", sum[float64](10.45, 34.56, 34.56, 67.54, 10))
	fmt.Println("sum:", sum[int](10, 3, 34, 67, 10))

	// s1 := sumOf("Hello", 10, false)
	// s2 := sumN[int](10, 20, 3, 40)

	var a, b, c, d float = 10.13, 12.45, 13.65, 34.98
	cs := sumC[float](a, b, c, d)
	fmt.Println("sumC:", cs.ToString())

}

func sumInts(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func sumOf(vals ...any) any {
	sum := 0.0 // float64
	for _, v := range vals {
		switch v.(type) {
		case float64:
			sum = sum + v.(float64)
		case float32:
			sum = sum + float64(v.(float32))
		case int:
			sum = sum + float64(v.(int))
		case int64:
			sum = sum + float64(v.(int64))
		case int32:
			sum = sum + float64(v.(int32))
		case int16:
			sum = sum + float64(v.(int16))
		case int8:
			sum = sum + float64(v.(int8))
		case uint64:
			sum = sum + float64(v.(uint64))
		case uint32:
			sum = sum + float64(v.(uint32))
		case uint16:
			sum = sum + float64(v.(uint16))
		case uint8:
			sum = sum + float64(v.(uint8))
		}
	}
	return sum
}

func sumN[T int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | int | float32 | float64](vals ...T) T {
	var s T
	for _, v := range vals {
		s += v
	}
	return s
}

func sum[T DT](vals ...T) T {
	var s T
	for _, v := range vals {
		s += v
	}
	return s
}

func sumC[T IFloat](vals ...T) T {
	var s T
	for _, v := range vals {
		s += v
	}
	return s
}

type DT interface {
	int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | int | float32 | float64
}

type float float64

func (f float) ToString() string {
	return fmt.Sprint(f)
}

type IFloat interface {
	float
	ToString() string
}

// Calc
// Add, Sub, Mul, Divide methods in clac
// Calc is a struct
// Calc should have a and b as fields
// the type a and type of b must be taken as generics
// all methods must work on those types.
// func Add[T]()(T) // T can be int,float of all variations
