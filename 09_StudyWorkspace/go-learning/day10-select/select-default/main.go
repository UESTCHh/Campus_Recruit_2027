package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(500 * time.Millisecond)
		ch <- "hello"
	}()

	select {
	case message := <-ch:
		fmt.Println("received:", message)

	default:
		fmt.Println("no message ready")
	}

	fmt.Println("main: end")
}
