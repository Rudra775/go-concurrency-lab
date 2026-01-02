package main

import (
	"os"
	"runtime/pprof"
)

func main() {
	ch := make(chan int)

	go func() {
		// wait a bit and dump
		for i := 0; i < 1e9; i++ {
		} // naive delay
		pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
	}()

	ch <- 1
	println("done")
}
