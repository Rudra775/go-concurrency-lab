package main

func main() {
	ch := make(chan int)

	<-ch //  blocks forever, waiting for a value

	println("done")
}
