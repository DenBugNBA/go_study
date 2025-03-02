package main

func main() {
	var ch chan int // nil-канал

	// panic: close of nil channel
	close(ch)
}
