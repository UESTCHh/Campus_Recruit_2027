package main

import "testing"

type FakeNotifier struct {
	messages []string
}

func (f *FakeNotifier) Send(message string) {
	f.messages = append(f.messages, message)
}

func TestUserServiceRegister(t *testing.T) {
	fake := FakeNotifier{}

	service := NewUserService(&fake)

	service.Register("UESTCHh")

	if len(fake.messages) != 1 {
		t.Fatalf("expected 1 message, got %d", len(fake.messages))
	}

	if fake.messages[0] != "welcome UESTCHh" {
		t.Fatalf("expected welcome UTESTCHh, got %s", fake.messages[0])
	}
}
