package main

import (
	"customers/messagebroker"
	"fmt"
	"runtime"
)

var (
	KAFKA_CONN []string = []string{"localhost:29092"}
)

func main() {
	fmt.Println("Started subscriber")
	var ch <-chan []byte
	ch = messagebroker.Subscribe(KAFKA_CONN, "customer.created")
	for val := range ch {
		fmt.Println(string(val))
	}
	runtime.Goexit()
}
