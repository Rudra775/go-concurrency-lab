package main

import (
	"fmt"
	"runtime"
)

func main() {
	ch := make(chan int)

	fmt.Println("Starting goroutine:", runtime.NumGoroutine())

	go func() {
		for v := range ch {
			fmt.Println("Received:", v)
		}
		fmt.Println("Goroutine exiting safely")
	}()

	ch <- 1
	ch <- 24

	fmt.Println("Ending Goroutine:", runtime.NumGoroutine())
}
