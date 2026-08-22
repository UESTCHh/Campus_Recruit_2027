package main

import (
	"fmt"
	"math"
)

// Shape 接口：定义几何形状的行为规范
type Shape interface {
	Area() float64      // 计算面积
	Perimeter() float64 // 计算周长
}

// Circle 圆形结构体
type Circle struct {
	Radius float64
}

// Circle 实现 Shape 接口（隐式实现，无需显式声明）
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle 矩形结构体
type Rectangle struct {
	Width, Height float64
}

// Rectangle 也隐式实现了 Shape 接口
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// PrintShape 接受任何实现了 Shape 接口的类型
func PrintShape(s Shape) {
	fmt.Printf("面积: %.2f, 周长: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	c := Circle{Radius: 5}              // 半径为 5 的圆
	r := Rectangle{Width: 4, Height: 6} // 宽 4 高 6 的矩形

	fmt.Println("圆形 (RUNOOB 示例):")
	PrintShape(c)

	fmt.Println("矩形 (RUNOOB 示例):")
	PrintShape(r)
}
