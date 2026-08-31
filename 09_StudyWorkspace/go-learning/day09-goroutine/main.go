package main

import (
	"fmt"
	"sync"
	"time"
)

func doWork(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		fmt.Println(name, "step:", i)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	fmt.Println("main: start")

	var wg sync.WaitGroup

	wg.Add(2)

	go doWork("worker-A", &wg)
	go doWork("worker-B", &wg)

	wg.Wait()

	fmt.Println("main: end")
}
