package main

import (
	"fmt"
	"time"
)

// 1. main is also a go routine
// 2. use go keyword infront of a func/method to create and execute a goroutine
// 3. by default no go routine waits for other goroutines to complete their execution
// 4. by default the order of execution is not determined; becase go routines are async

func main() {

	go func() {
		fmt.Println("Hello World-1")
	}()

	go greet()
	time.Sleep(time.Millisecond * 1)
	fmt.Println("end of main")

}

func greet() {
	go fmt.Print("Hello")
	go fmt.Print(" World!")
	go fmt.Println("--")
	//fmt.Println("Hello World!")
}
