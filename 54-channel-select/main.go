package main

import (
	"fmt"
	"time"
)

func main() {

	s1 := server1()
	s2 := server2()
	s3 := server3()
	t1 := time.After(time.Millisecond * 10)

loop:
	select {
	case str := <-s1:
		fmt.Println(str)
	case str := <-s2:
		fmt.Println(str)
	case str := <-s3:
		fmt.Println(str)
	case <-t1:
		fmt.Println("Timedout")
	default:
		//fmt.Println("no channel has returned")
		goto loop
	}

	//}

}

func server1() <-chan string {
	out := make(chan string)
	go func() {
		time.Sleep(time.Millisecond * 10)
		out <- "Sever-1"
	}()
	return out
}

func server2() <-chan string {
	out := make(chan string)
	go func() {
		time.Sleep(time.Millisecond * 10)
		out <- "Sever-2"
	}()
	return out
}

func server3() <-chan string {
	out := make(chan string)
	go func() {
		time.Sleep(time.Millisecond * 10)
		out <- "Sever-3"
	}()
	return out
}
