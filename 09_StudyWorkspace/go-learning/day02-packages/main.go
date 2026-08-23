package main

import (
	"fmt"

	"example.com/go-learning/day02-packages/calculator"
)

func main() {
	sum := calculator.Add(10, 5)
	difference := calculator.Subtract(10, 5)

	fmt.Println("10 + 5 =", sum)
	fmt.Println("10 - 5 =", difference)

	fmt.Println(calculator.Version)
}
