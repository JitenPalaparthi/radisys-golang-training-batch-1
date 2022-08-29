package main

import "fmt"

func main() {

	msgCh := make(chan string)
	done := make(chan struct{})
	go sender(msgCh)
	go receiver(msgCh, done)
	<-done
}

func sender(send chan<- string) { // send only channel chan <-
	for i := 1; i <= 10; i++ {
		send <- "Hello-->" + fmt.Sprint(i)
	}

}

func receiver(receive <-chan string, done chan<- struct{}) { // receive only channel <-chan
	for i := 1; i <= 10; i++ {
		msg := <-receive
		fmt.Println(msg)
	}
	done <- struct{}{}
}

// modify receiver function to return done channel
