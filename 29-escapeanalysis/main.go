package main

func main() {
	var num1 int = 100
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
	//fmt.Println(time.Now())
}

// does not escape : stored in stack
// escapes to heap: compiler has decided to create in heap memory
// moved to heap
