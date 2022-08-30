package main

import "fmt"

func main() {

	msgCh := make(chan string)
	done := make(chan struct{})
	go sender(msgCh, "sender-1")
	go sender(msgCh, "sender-2")
	go receiver(msgCh, done, "receiver-1")
	go receiver(msgCh, done, "receiver-2")
	<-done
	<-done
}

func sender(send chan<- string, str string) { // send only channel chan <-
	for i := 1; i <= 10; i++ {
		send <- "Hello-->" + fmt.Sprint(i) + str
	}
}

func receiver(receive <-chan string, done chan<- struct{}, str string) { // receive only channel <-chan
	for i := 1; i <= 10; i++ {
		msg := <-receive
		fmt.Println(msg + str)
	}
	done <- struct{}{}
}

// modify receiver function to return done channel
