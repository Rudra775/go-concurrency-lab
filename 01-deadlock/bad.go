package main

func main() {
	ch := make(chan int) // unbuffered channel

	ch <- 1 // cd blocks forever because nobody is receiving

	println("done") // this never executes
}
