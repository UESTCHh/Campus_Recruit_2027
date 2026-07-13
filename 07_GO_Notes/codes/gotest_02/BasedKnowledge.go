//Go 标记
//Go 程序可以由多个标记组成，可以是关键字，标识符，常量，字符串，符号。如以下 GO 语句由 6 个标记组成：
//
// fmt.Println("Hello, World!")
// 6 个标记是(每行一个)：
//
// 1. fmt
// 2. .
// 3. Println
// 4. (
// 5. "Hello, World!"
// 6. )

// 行分隔符
// 在 Go 程序中，一行代表一个语句结束。每个语句不需要像 C 家族中的其它语言一样以分号 ; 结尾，因为这些工作都将由 Go 编译器自动完成。
//
// 如果你打算将多个语句写在同一行，它们则必须使用 ; 人为区分，但在实际开发中我们并不鼓励这种做法。
//
// 以下为两个语句：
//
// fmt.Println("Hello, World!")
// fmt.Println("菜鸟教程：runoob.com")

// 注释
// 注释不会被编译，每一个包应该有相关注释。
//
// 单行注释是最常见的注释形式，你可以在任何地方使用以 // 开头的单行注释。多行注释也叫块注释，均已以 /* 开头，并以 */ 结尾。如：
//
// // 单行注释
// /*
// Author by 菜鸟教程
// 我是多行注释
// */

// 标识符
// 标识符用来命名变量、类型等程序实体。一个标识符实际上就是一个或是多个字母(A~Z和a~z)数字(0~9)、下划线_组成的序列，但是第一个字符必须是字母或下划线而不能是数字。
//
// 以下是有效的标识符：
//
// mahesh   kumar   abc   move_name   a_123
// myname50   _temp   j   a23b9   retVal
// 以下是无效的标识符：
//
// 1ab（以数字开头）
// case（Go 语言的关键字）
// a+b（运算符是不允许的）

// 字符串连接
// Go 语言的字符串连接可以通过 + 实现：
//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("Google" + "Runoob")
//}

//以上实例输出结果为：
//
//GoogleRunoob

//关键字
//下面列举了 Go 代码中会使用到的 25 个关键字或保留字：
//
//break	default	func	interface	select
//case	defer	go	map	struct
//chan	else	goto	package	switch
//const	fallthrough	if	range	type
//continue	for	import	return	var
//除了以上介绍的这些关键字，Go 语言还有 36 个预定义标识符：
//
//append	bool	byte	cap	close	complex	complex64	complex128	uint16
//copy	false	float32	float64	imag	int	int8	int16	uint32
//int32	int64	iota	len	make	new	nil	panic	uint64
//print	println	real	recover	string	true	uint	uint8	uintptr
//程序一般由关键字、常量、变量、运算符、类型和函数组成。
//
//程序中可能会使用到这些分隔符：括号 ()，中括号 [] 和大括号 {}。
//
//程序中可能会使用到这些标点符号：.、,、;、: 和 …。

//Go 语言的空格
//在 Go 语言中，空格通常用于分隔标识符、关键字、运算符和表达式，以提高代码的可读性。
//
//Go 语言中变量的声明必须使用空格隔开，如：
//
//var x int
//const Pi float64 = 3.14159265358979323846
//在运算符和操作数之间要使用空格能让程序更易阅读：
//
//无空格：
//
//fruit=apples+oranges;
//在变量与运算符间加入空格，程序看起来更加美观，如：
//
//fruit = apples + oranges;
//在关键字和表达式之间要使用空格。
//
//例如：
//
//if x > 0 {
//    // do something
//}
//在函数调用时，函数名和左边等号之间要使用空格，参数之间也要使用空格。
//
//例如：
//
//result := add(2, 3)

// 格式化字符串
// Go 语言中使用 fmt.Sprintf 或 fmt.Printf 格式化字符串并赋值给新串：
//
// Sprintf 根据格式化参数生成格式化的字符串并返回该字符串。
// Printf 根据格式化参数生成格式化的字符串并写入标准输出。
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// %d 表示整型数字，%s 表示字符串
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	//Go fmt.Sprintf 格式化字符串
	//
	//Go 可以使用 fmt.Sprintf 来格式化字符串，格式如下：
	//
	//fmt.Sprintf(格式化样式, 参数列表…)
	//格式化样式：字符串形式，格式化符号以 % 开头， %s 字符串格式，%d 十进制的整数格式。
	//参数列表：多个参数以逗号分隔，个数必须与格式化样式中的个数一一对应，否则运行时会报错。
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)

	// 另外一个实例，%d 表示整型
	const name, age = "Kim", 22
	s := fmt.Sprintf("%s is %d years old.\n", name, age)
	io.WriteString(os.Stdout, s) // 简单起见，忽略一些错误
}

//Go 字符串格式化符号:
//
//格  式	描  述
//%v	按值的本来值输出
//%+v	在 %v 基础上，对结构体字段名和值进行展开
//%#v	输出 Go 语言语法格式的值
//%T	输出 Go 语言语法格式的类型和值
//%%	输出 % 本体
//%b	整型以二进制方式显示
//%o	整型以八进制方式显示
//%d	整型以十进制方式显示
//%x	整型以十六进制方式显示
//%X	整型以十六进制、字母大写方式显示
//%U	Unicode 字符
//%f	浮点数
//%p	指针，十六进制方式显示
//对齐方式
//通过在格式化字符串中使用宽度和对齐参数，可以控制生成的字符串的对齐方式。
//
//常用的对齐参数有：
//
//%s：字符串格式，可以使用以下对齐参数：
//%s：默认对齐方式，左对齐。
//%10s：指定宽度为 10 的右对齐。
//%-10s：指定宽度为 10 的左对齐。
//%d：整数格式，可以使用以下对齐参数：
//%d：默认对齐方式，右对齐。
//%10d：指定宽度为 10 的右对齐。
//%-10d：指定宽度为 10 的左对齐。
//%f：浮点数格式，可以使用以下对齐参数：
//%f：默认对齐方式，右对齐。
//%10f：指定宽度为 10 的右对齐。
//%-10f：指定宽度为 10 的左对齐。

//Go fmt.Printf 格式化字符串
//
//fmt.Printf 是 Go 语言中一个功能强大的输出格式化函数，主要用于格式化字符串并将结果输出到标准输出（通常是控制台）。
//
//fmt.Printf 按照指定的格式化字符串对后续变量进行格式化，语法如下：
//
//fmt.Printf(格式化样式, 参数列表…)
//格式化样式：字符串形式，格式化符号以 % 开头， %s 字符串格式，%d 十进制的整数格式。
//参数列表：多个参数以逗号分隔，个数必须与格式化样式中的个数一一对应，否则运行时会报错。

//Go 字符串格式化符号:
//
//格式化符号	描述	示例
//通用格式
//%v	以默认格式输出变量	fmt.Printf("%v", 42)
//%+v	对结构体加字段名的方式输出	fmt.Printf("%+v", struct{A int}{A: 42})
//%#v	以 Go 语法格式化输出	fmt.Printf("%#v", map[string]int{"a": 1})
//%T	输出值的类型	fmt.Printf("%T", 42)
//%%	输出百分号	fmt.Printf("%%")
//布尔值
//%t	输出 true 或 false	fmt.Printf("%t", true)
//整数
//%b	二进制表示	fmt.Printf("%b", 5)
//%c	Unicode 对应字符	fmt.Printf("%c", 65)
//%d	十进制表示	fmt.Printf("%d", 42)
//%o	八进制表示	fmt.Printf("%o", 10)
//%x	十六进制表示（小写字母）	fmt.Printf("%x", 15)
//%X	十六进制表示（大写字母）	fmt.Printf("%X", 15)
//%U	Unicode 格式输出	fmt.Printf("%U", 65)
//浮点数
//%f	十进制浮点数	fmt.Printf("%f", 3.14)
//%e	科学计数法（小写 e）	fmt.Printf("%e", 3.14)
//%E	科学计数法（大写 E）	fmt.Printf("%E", 3.14)
//%g	自动选择 %f 或 %e 的简洁表示	fmt.Printf("%g", 3.14)
//字符串与字节
//%s	普通字符串	fmt.Printf("%s", "Go")
//%q	带双引号的字符串或字符	fmt.Printf("%q", "Go")
//%x	每个字节用两字符十六进制表示	fmt.Printf("%x", "abc")
//%X	十六进制（大写）表示	fmt.Printf("%X", "abc")
//指针
//%p	指针地址	fmt.Printf("%p", &x)
//格式化字符串由常规文本和格式化占位符组成。
//
//格式化占位符以 % 开头，后接一个或多个字符，指明格式化类型。
//
//格式化占位符的结构为：
//
//%[flags][width][.precision]verb
//flags：用于控制格式化输出的标志（可选）。
//-：左对齐。
//+：始终显示数值的符号。
//0：用零填充。
//#：为二进制、八进制、十六进制等加上前缀。
//空格：正数前加空格，负数前加 -。
//width：输出宽度（可选）。
//.precision：浮点数小数点后的位数（可选）。
//verb：用于指定数据的格式化方式。
//1. 格式化整数
//
//fmt.Printf("十进制: %d, 二进制: %b, 八进制: %o, 十六进制: %x\n", 42, 42, 42, 42)
//// 输出：十进制: 42, 二进制: 101010, 八进制: 52, 十六进制: 2a
//2. 格式化浮点数
//
//fmt.Printf("普通: %f, 科学计数法: %e, 自动: %g\n", 123.456, 123.456, 123.456)
//// 输出：普通: 123.456000, 科学计数法: 1.234560e+02, 自动: 123.456
//3. 格式化字符串
//
//fmt.Printf("普通字符串: %s, 带引号: %q, 十六进制: %x\n", "Go", "Go", "Go")
//// 输出：普通字符串: Go, 带引号: "Go", 十六进制: 476f
//4. 使用宽度和对齐
//
//fmt.Printf("|%10s|%-10s|\n", "right", "left")
//// 输出：|     right|left      |
//5. 格式化结构体
//
//type User struct {
//    Name string
//    Age  int
//}
//user := User{Name: "Alice", Age: 30}
//fmt.Printf("默认: %v\n带字段名: %+v\nGo语法: %#v\n", user, user, user)
//// 输出：
//// 默认: {Alice 30}
//// 带字段名: {Name:Alice Age:30}
//// Go语法: main.User{Name:"Alice", Age:30}
//6. 格式化布尔值
//
//fmt.Printf("布尔值: %t, %t\n", true, false)
//// 输出：布尔值: true, false
