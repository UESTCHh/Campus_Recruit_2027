package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	counter := 0

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 100000; i++ {
			counter++
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 100000; i++ {
			counter++
		}
	}()

	wg.Wait()

	fmt.Println("counter:", counter)
}
