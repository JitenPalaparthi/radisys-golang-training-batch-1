package main

import (
	"fmt"
	"sync"
	"time"
)

// do not use global variables across multiple go routines
// use mutex
var num int = 0 // data race

func main() {

	mu := new(sync.Mutex)
	go func() {
		for i := 1; i <= 100; i++ {
			go increment(mu)
		}
	}()
	// go func() {
	// 	for i := 1; i <= 100; i++ {
	// 		go decrement(mu)
	// 	}
	// }()
	time.Sleep(time.Second * 1)
	fmt.Println("\n", num)

	// var slice1 []int

	// slice1 = append(slice1, 1000, 10001, 10002) // append instantiate slice as well
	// fmt.Println(slice1)

	// slice2 := new([]int)
}

func increment(mu *sync.Mutex) {
	mu.Lock()
	num++
	fmt.Println(num)
	mu.Unlock()
}

func decrement(mu *sync.Mutex) {
	mu.Lock()
	num--
	fmt.Println(num)
	mu.Unlock()
}
