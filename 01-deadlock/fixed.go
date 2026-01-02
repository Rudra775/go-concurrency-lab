package main

func main() {
	ch := make(chan int)

	go func() {
		v := <-ch
		println("received:", v)
	}()

	ch <- 1

	println("done")
}
