package main

import (
	"fmt"
	"time"
)

// buffered channels
// sender and receiver is blocked only if the buffer is full
func main() {
	// ch := make(chan int, 2)
	// fmt.Println(cap(ch))
	// ch <- 100
	// fmt.Println(len(ch)) // length of the buffer
	// ch <- 101
	// fmt.Println(len(ch)) // length of the buffer
	// fmt.Println(<-ch)
	// ch <- 102
	// close(ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	done := make(chan struct{})
	out1, out2, out3 := sender()
	t1 := time.Now()
	go receiver(out1, "receiver-1", done)
	go receiver(out2, "receiver-2", done)
	go receiver(out3, "receiver-3", done)

	<-done
	<-done
	<-done
	fmt.Println(t1, "----", time.Now())
}

func sender() (<-chan int, <-chan int, <-chan int) {
	send1 := make(chan int, 3) // buffered channel
	send2 := make(chan int, 3)
	send3 := make(chan int, 3)
	go func() {
		for i := 1; i <= 10000; i++ {
			send1 <- i
			send2 <- i
			send3 <- i
		}
		close(send1)
		close(send2)
		close(send3)
	}()
	return send1, send2, send3
}

func receiver(receive <-chan int, str string, done chan<- struct{}) {
	for val := range receive {
		fmt.Println(str + "-->" + fmt.Sprint(val))
	}
	done <- struct{}{}
}
