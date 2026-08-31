package main

import "fmt"

func produce(ch chan int) {
	for i := 1; i <= 3; i++ {
		ch <- i * 10
	}

	close(ch)
}

func main() {
	ch := make(chan int)

	go produce(ch)

	for value := range ch {
		fmt.Println("received:", value)
	}

	fmt.Println("channel closed")
}
