package main

import (
	"fmt"
	"math/rand"
	"time"
)

// main is the start point of execution
// main does not return
// main does not contain any arguments
func main() {
	println("Hello", "World!")
	fmt.Println("Hello Radisys minds!")
	var hello string = "Hello World!"
	fmt.Printf("%s", hello)
	fmt.Println("\nCurrent date and time", time.Now())
	fmt.Println("Random number", rand.Intn(100000))
}
