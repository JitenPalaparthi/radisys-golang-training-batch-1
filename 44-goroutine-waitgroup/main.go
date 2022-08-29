package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	wg.Add(1)
	i := 1
	go func() {
		wg.Add(100)
		for {
			go func() {
				mu.Lock()
				if i > 100 {
					wg.Done()
					mu.Unlock()
					return
				} else {
					//mu.Lock()
					fmt.Println("i-->", i)
					i++
					wg.Done()
				}
				mu.Unlock()
			}()
		}
	}()
	wg.Wait()
}
