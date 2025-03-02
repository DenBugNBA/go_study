package main

func main() {
	var ch chan int

	<-ch // deadlock - chan receive (nil chan)
}
