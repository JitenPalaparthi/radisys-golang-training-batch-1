package main

import (
	"fmt"
)

// channel is a pipe that is used to communicate go routines
// chan is used to create a channel; make builtin func is used to instantiate a channel
// there are two kinds of channels. Buffered and unbuffered.
// unbuffered means chan with no size given

// 1- For unbuffered channel, sender is blocked until receiver receives the value
// 	- sender and reveicer are suppose to be go routines
// 2- The receiver is blocked untile the sender sends the value
// 3- To send a value to channel ch <- 100 ; To channel
// 4- To receive a value from channel <- ch ; From channel
// 5- To instantiate a int unbuffered channel use make(chan int)
// 6- as long as sender and receiver are separate goroutines no issue with the order;
// 	There is nothing like sender should be executed first.

func main() {
	var ch chan int     // ch can send and receive int type values through it.
	ch = make(chan int) // the size of the channel is zero
	// ch <- 100           // channel is receiving value; we are sending value to channel
	// num := <-ch         // value is being sent from the channel. That means value is being received from the channel
	go func() {
		// time.Sleep(time.Second * 2)
		ch <- 100
	}()

	// go func() {
	// 	ch <- 101
	// }()

	num := <-ch
	fmt.Println(num)

}
