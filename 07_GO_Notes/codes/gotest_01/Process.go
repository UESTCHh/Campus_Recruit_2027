package main

import "fmt"

func main() {
	// if 语句：条件表达式不需要括号，但大括号必须有
	score := 85
	if score >= 90 {
		fmt.Println("RUNOOB 评级: A")
	} else if score >= 60 {
		fmt.Println("RUNOOB 评级: B")
	} else {
		fmt.Println("RUNOOB 评级: C")
	}

	// for 循环一：标准三要素形式（类似 C 的 for）
	fmt.Print("计数: ")
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// for 循环二：仅条件（类似 while）
	sum := 0
	n := 1
	for n <= 10 {
		sum += n
		n++
	}
	fmt.Println("1 到 10 之和:", sum)

	// for 循环三：range 遍历（最常用）
	fruits := []string{"apple", "banana", "cherry"}
	for index, fruit := range fruits {
		fmt.Printf("索引 %d: %s\n", index, fruit)
	}

	// switch 语句：每个 case 自动 break，无需手动添加
	day := "周一"
	switch day {
	case "周一":
		fmt.Println("新的一周开始了")
	case "周五":
		fmt.Println("周末快到了")
	case "周六":
	case "周天":
		fmt.Println("放假了！")
	default:
		fmt.Println("普通的一天")
	}
}
