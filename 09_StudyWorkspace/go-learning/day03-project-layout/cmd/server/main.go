package main

import (
	"fmt"

	"example.com/go-learning/day03-project-layout/internal/service"
)

func main() {

	name := service.GetUserName()

	fmt.Println(name)

}
