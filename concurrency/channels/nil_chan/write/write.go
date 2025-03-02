package main

func main() {
	var ch chan int

	ch <- 1 // deadlock - chan send (nil chan)
}
