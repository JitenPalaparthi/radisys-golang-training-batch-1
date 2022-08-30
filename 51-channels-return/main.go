package main

import (
	"fmt"
	"time"
)

var isClosed bool

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan struct{})
	go sender(ch1)
	go sender(ch2)
	go receiver(ch1, done)
	go receiver(ch2, done)
	<-done
	<-done
}

func sender(send chan<- int) {
	for i := 1; i <= 100; i++ {
		time.Sleep(time.Microsecond * 1)
		//if !isClosed {
		send <- i
		//}
	}
	//isClosed = true
	close(send) // fundametally only sender must close the channel
	// close a channel does not mean the channel is nil.
	// close and nil channels are two different things.
	// when the channel is closed, still receiver is trying to receive then,
	// type inference work and channel emits 0 for int, false for bool etc..

}

func receiver(receive <-chan int, done chan<- struct{}) {
	// range loop on channel iterates until the channel is closed
	for val := range receive {
		fmt.Println(val)
	}
	done <- struct{}{}
}

// never use global variables to read or write data
// never try to implement go routine inside a loop
// if same goroutine is called multiple times, when it is a sender, ensure to use
// multiple channels. The reason behind that is one of the goroutine may send data
// to a closed channel; becase sender only closes the channel.Sender can never know
// if a channel is closed and hence it crashes the application
