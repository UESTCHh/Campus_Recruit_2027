package main

import "fmt"

func main() {
	// 方式一：标准声明（显式类型）
	var age int = 25

	// 方式二：声明后赋值（类型在声明时指定）
	var score float64
	score = 98.5

	// 方式三：短变量声明（最常用，仅限函数内部）
	name := "runoob"

	// 方式四：多变量同时声明
	x, y := 10, 20

	// 常量声明（编译时确定，不可修改）
	const Pi = 3.14159
	const AppName = "RUNOOB"

	fmt.Println(name, age, score, x, y, Pi, AppName)
}
