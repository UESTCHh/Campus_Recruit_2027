package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(300 * time.Millisecond)
		ch1 <- "message from ch1"
	}()

	go func() {
		time.Sleep(600 * time.Millisecond)
		ch2 <- "message from ch2"
	}()

	select {
	case message := <-ch1:
		fmt.Println("received:", message)

	case message := <-ch2:
		fmt.Println("received:", message)
	}

	fmt.Println("main: end")
}
