package main

import (
	"fmt"
	"runtime"
)

func main() {
	ch := make(chan int)

	fmt.Println("Starting goroutine:", runtime.NumGoroutine())

	go func() {
		for range ch {

		}
	}()

	fmt.Println("Ending goroutine:", runtime.NumGoroutine())
}
