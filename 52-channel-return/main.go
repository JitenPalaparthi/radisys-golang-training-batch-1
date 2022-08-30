package main

import (
	"fmt"
	"time"
)

func main() {

	var ch1 <-chan int
	done := make(chan struct{})
	ch1 = sender()
	//ch2 := sender()
	//ch3 := sender()
	go receiver(ch1, done)
	//go receiver(ch2, done)
	//go receiver(ch3, done)
	// for val := range ch {
	// 	fmt.Println(val)
	// }
	// runtime.Goexit() // eventhrough it returns an error; make it simple
	<-done
	//<-done
	//<-done
}

func sender() <-chan int {
	send := make(chan int)
	// go func() {
	// 	for i := 1; i <= 100; i++ {
	// 		send <- i
	// 	}
	// 	close(send)
	// }()
	go innerSender(send)
	return send
}

func receiver(receive <-chan int, done chan<- struct{}) {
	// range loop on channel iterates until the channel is closed
	// for val := range receive {
	// 	fmt.Println(val)
	// }

	for i := 1; i <= 110; i++ { // sender is sending only 100 values but receiver tries to receive 110;type inference works here
		val, ok := <-receive
		fmt.Println(val, ok)
	}
	done <- struct{}{}
}

func innerSender(send chan<- int) {
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
