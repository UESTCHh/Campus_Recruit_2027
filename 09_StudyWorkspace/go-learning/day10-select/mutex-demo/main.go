package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	counter := 0

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 100; i++ {
			mu.Lock()

			counter++

			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 100; i++ {
			mu.Lock()

			counter++

			mu.Unlock()
		}
	}()

	wg.Wait()

	fmt.Println("counter:", counter)
}
