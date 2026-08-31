package main

import "fmt"

func calculate(result chan int) {
	result <- 42
}

func main() {
	result := make(chan int)

	go calculate(result)

	value := <-result

	fmt.Println("result:", value)
}
