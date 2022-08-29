package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	// add increases the counter
	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Hello World-1")
		wg.Done() // Done decreases the counter
	}(wg)

	go greet1(wg)

	//wg.Add(1)
	go func(wg *sync.WaitGroup) {
		greet2()
		wg.Done()
	}(wg)
	fmt.Println("Hello World-2")
	wg.Wait()
	// Wait waits until the counter is zero
}

func greet1(wg *sync.WaitGroup) {
	//wg.Add(1)
	fmt.Println("Hello World-3")
	wg.Done()
}

func greet2() {
	fmt.Println("Hello World-4")
}

// 0- sync.Mutex
// 1- WaitGroup
// 2- runtmie.Goexit
// 3- runtime.Gosched
