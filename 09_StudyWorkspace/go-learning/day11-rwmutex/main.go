package main

import (
	"fmt"
	"sync"
	"time"
)

type UserStore struct {
	mu    sync.RWMutex
	users map[int]string
}

func NewUserStore() *UserStore {
	return &UserStore{
		users: map[int]string{
			1: "Alice",
			2: "Bob",
		},
	}
}

func (s *UserStore) Get(id int) string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	fmt.Println("reader start:", id)

	time.Sleep(500 * time.Millisecond)

	value := s.users[id]

	fmt.Println("reader end:", id)

	return value
}

func (s *UserStore) Set(id int, name string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	fmt.Println("writer start:", id)

	time.Sleep(500 * time.Millisecond)

	s.users[id] = name

	fmt.Println("writer end:", id)
}

func main() {
	store := NewUserStore()

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()

		fmt.Println(
			"reader result:",
			store.Get(1),
		)
	}()

	go func() {
		defer wg.Done()

		fmt.Println(
			"reader result:",
			store.Get(2),
		)
	}()

	go func() {
		defer wg.Done()

		store.Set(3, "Charlie")
	}()

	wg.Wait()
}
