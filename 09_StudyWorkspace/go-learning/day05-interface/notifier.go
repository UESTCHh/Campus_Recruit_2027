package main

import "fmt"

type EmailNotifier struct {
}

type SMSNotifier struct {
}

func (e EmailNotifier) Send(message string) {
	fmt.Println("Email: ", message)
}

func (s *SMSNotifier) Send(message string) {
	fmt.Println("SMS:", message)
}

func SendNotification(n Notifier, message string) {
	n.Send(message)
}
