package main

import "fmt"

func main() {
	name := "hui"
	age := 20
	isLearningGo := true

	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("learning Go:", isLearningGo)

	message := greet(name)
	fmt.Println(message)

	score := 85
	result := classifyScore(score)
	fmt.Println("score:", score, "result:", result)

	printNumbers(3)

	scores := []int{80, 90, 70}
	total := sumScores(scores)

	fmt.Println("scores:", scores)
	fmt.Println("total:", total)
}
