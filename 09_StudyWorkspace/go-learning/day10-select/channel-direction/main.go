package main

import "fmt"

func produce(ch chan<- int) {
	ch <- 42
}

func consume(ch <-chan int) {
	value := <-ch
	fmt.Println("received:", value)
}

func main() {
	ch := make(chan int)

	go produce(ch)

	consume(ch)
}
