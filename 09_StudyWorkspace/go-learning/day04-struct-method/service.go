package main

type UserService struct {
	count int
}

func (s UserService) Add() {

	s.count++
}

func (s *UserService) AddPointer() {

	s.count++
}
