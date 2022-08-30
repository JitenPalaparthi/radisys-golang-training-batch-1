package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int)
	go sender(ch)
	go receiver(ch)
	runtime.Goexit() // Terminates the goroutine that calls it.. Here the caller is main
	// It terminates main after all other goroutines finish of their execution
	// terminating main means not properly return main. That means the end of main
	// cannot be reached , hence error
}

func sender(send chan<- int) {
	//i := 1
	for i := 1; i <= 10; i++ {
		// if i == 10 {
		// 	return
		// }
		//i++
		time.Sleep(time.Microsecond * 1)
		send <- i
	}
}

func receiver(receive <-chan int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(<-receive)
	}
}
