package main

import (
	"errors"
	"fmt"
)

// add 函数：两个参数同为 int 类型，返回一个 int
func add(a, b int) int {
	return a + b
}

// divide 函数：返回两个值——结果和错误
// Go 惯用 (result, error) 模式进行错误处理
func divide(a, b float64) (float64, error) {
	if b == 0 {
		// 返回零值和自定义错误
		return 0, errors.New("除数不能为零")
	}
	return a / b, nil // nil 表示没有错误
}

// swap 函数：多返回值用于返回多个计算结果
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// 调用普通函数
	sum := add(10, 20)
	fmt.Println("10 + 20 =", sum)

	// 调用多返回值函数并检查错误
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	// 测试除零情况
	_, err = divide(10, 0) // _ 用于忽略不需要的返回值
	if err != nil {
		fmt.Println("除零错误:", err) // 预期输出此错误
	}

	// 多返回值交换
	first, second := swap("RUNOOB", "Hello")
	fmt.Println("交换后:", first, second)
}
