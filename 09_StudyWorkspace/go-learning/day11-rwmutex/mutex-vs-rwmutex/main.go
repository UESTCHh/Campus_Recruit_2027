package main

import (
	"fmt"
	"sync"
	"time"
)

type MutexStore struct {
	mu sync.Mutex
}

func (s *MutexStore) Read(id int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	fmt.Println("Mutex reader start:", id)

	time.Sleep(500 * time.Millisecond)

	fmt.Println("Mutex reader end:", id)
}

type RWMutexStore struct {
	mu sync.RWMutex
}

func (s *RWMutexStore) Read(id int) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	fmt.Println("RWMutex reader start:", id)

	time.Sleep(500 * time.Millisecond)

	fmt.Println("RWMutex reader end:", id)
}

func runMutexReaders() {
	var store MutexStore
	var wg sync.WaitGroup

	start := time.Now()

	wg.Add(3)

	for i := 1; i <= 3; i++ {
		go func(id int) {
			defer wg.Done()
			store.Read(id)
		}(i)
	}

	wg.Wait()

	fmt.Println(
		"Mutex elapsed:",
		time.Since(start),
	)
}

func runRWMutexReaders() {
	var store RWMutexStore
	var wg sync.WaitGroup

	start := time.Now()

	wg.Add(3)

	for i := 1; i <= 3; i++ {
		go func(id int) {
			defer wg.Done()
			store.Read(id)
		}(i)
	}

	wg.Wait()

	fmt.Println(
		"RWMutex elapsed:",
		time.Since(start),
	)
}

func main() {
	fmt.Println("=== Mutex ===")
	runMutexReaders()

	fmt.Println()

	fmt.Println("=== RWMutex ===")
	runRWMutexReaders()
}
