package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 1; i <= 1000; i++ {
			func() {
				for j := 0; j < 10; j++ {
					go fmt.Println("I-->", i, "-->", j)
				}
			}()
		}
	}()
	time.Sleep(time.Second * 5)
	fmt.Println("end of main")

}
