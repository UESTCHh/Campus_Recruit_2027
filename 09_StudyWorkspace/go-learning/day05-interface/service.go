package main

import "fmt"

type Notifier interface {
	Send(message string)
}

type UserService struct {
	notifier Notifier
}

func NewUserService(notifier Notifier) *UserService {
	return &UserService{
		notifier: notifier,
	}
}

func (s *UserService) Register(name string) {
	fmt.Println("register user:", name)
	s.notifier.Send("welcome " + name)
}
