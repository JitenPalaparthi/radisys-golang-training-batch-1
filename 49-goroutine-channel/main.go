package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan struct{})
	go sender(ch)
	go receiver(ch, done)
	<-done
}

func sender(send chan<- int) {
	for i := 1; i <= 100; i++ {
		time.Sleep(time.Microsecond * 1)
		send <- i
	}
	close(send) // fundametally only sender must close the channel
	// close a channel does not mean the channel is nil.
	// close and nil channels are two different things.
	// when the channel is closed, still receiver is trying to receive then,
	// type inference work and channel emits 0 for int, false for bool etc..
}

func receiver(receive <-chan int, done chan<- struct{}) {
	for {
		val, ok := <-receive // only while reading values from channel can check whether the channel is closed or not
		// if ok is true then channel is not closed; Else channel is closed
		if !ok {
			done <- struct{}{}
			return
		}

		fmt.Println(val)
	}
}
