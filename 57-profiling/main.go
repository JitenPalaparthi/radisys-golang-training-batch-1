package main

import (
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func main() {
	// code

	go func() {
		_ = http.ListenAndServe(":8848", nil)
	}()
	runtime.Goexit()
}
