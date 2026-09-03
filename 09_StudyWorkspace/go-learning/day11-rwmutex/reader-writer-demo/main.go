package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.RWMutex

	readersReady := make(chan int, 2)
	releaseReaders := make(chan struct{})

	writerTrying := make(chan struct{})
	writerAcquired := make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(3)

	reader := func(id int) {
		defer wg.Done()

		mu.RLock()

		fmt.Println("reader acquired:", id)

		readersReady <- id

		<-releaseReaders

		fmt.Println("reader releasing:", id)

		mu.RUnlock()
	}

	go reader(1)
	go reader(2)

	<-readersReady
	<-readersReady

	go func() {
		defer wg.Done()

		fmt.Println("writer: trying Lock")

		close(writerTrying)

		mu.Lock()

		fmt.Println("writer: acquired Lock")

		close(writerAcquired)

		mu.Unlock()
	}()

	<-writerTrying

	time.Sleep(100 * time.Millisecond)

	select {
	case <-writerAcquired:
		fmt.Println("writer acquired while readers were active")
	default:
		fmt.Println("writer is blocked while readers hold RLock")
	}

	fmt.Println("main: release readers")

	close(releaseReaders)

	wg.Wait()

	fmt.Println("main: end")
}
