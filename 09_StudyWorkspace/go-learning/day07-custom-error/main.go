package main

import (
	"errors"
	"fmt"
)

func main() {
	err := RegisterUser(10)

	if err == nil {
		fmt.Println("register success")
		return
	}

	fmt.Println("full error:", err)

	var validationErr *ValidationError

	if errors.As(err, &validationErr) {
		fmt.Println("matched ValidationError")
		fmt.Println("field:", validationErr.Field)
		fmt.Println("value:", validationErr.Value)
		fmt.Println("message:", validationErr.Message)
		return
	}

	fmt.Println("unknown error")
}
