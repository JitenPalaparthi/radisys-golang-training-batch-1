package main

import (
	"demo/greeting"

	"fmt"
	"sync"
)

var (
	wg  *sync.WaitGroup
	g1  *greeting.Greetings
	err error
)

// init is a special func
// it is called as soon as the package is called
// any number of init functons can be called
func init() {
	//wg = &sync.WaitGroup{}
	wg = new(sync.WaitGroup)
	g1, err = greeting.New(wg)
}

func main() {
	wg.Add(1)
	go greet()
	if err != nil {
		fmt.Println(err)
	} else {
		wg.Add(1)
		go g1.Greet("Hello World-2")
	}

	wg.Wait()
}

// can use global variable
func greet() {
	if wg != nil {
		fmt.Println("Hello World-1")
		wg.Done()
	}
}
