package main

import "fmt"

func greet(name string) string {
	return "Hello, " + name + "!"
}

func classifyScore(score int) string {
	if score >= 60 {
		return "passed"
	}

	return "failed"
}

func printNumbers(limit int) {
	for number := 1; number <= limit; number++ {
		fmt.Println("number:", number)
	}
}

func sumScores(scores []int) int {
	total := 0

	for _, score := range scores {
		total += score
	}

	return total
}
