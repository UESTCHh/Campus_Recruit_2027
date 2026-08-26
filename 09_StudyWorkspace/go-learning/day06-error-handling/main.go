package main

import (
	"errors"
	"fmt"
)

func main() {
	name, err := GetUserDisplayName(100)

	if err != nil {
		fmt.Println("full error:", err)

		if errors.Is(err, ErrUserNotFound) {
			fmt.Println("matched: ErrUserNotFound")
			return
		}
		if errors.Is(err, ErrInvalidUserID) {
			fmt.Println("matched: ErrInvalidUserID")
			return
		}

		fmt.Println("unknown error")
		return
	}

	fmt.Println("user:", name)
}
