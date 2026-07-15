package main

//Go 语言数组
//Go 语言提供了数组类型的数据结构。
//
//数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整型、字符串或者自定义类型。
//
//相对于去声明 number0, number1, ..., number99 的变量，使用数组形式 numbers[0], numbers[1] ..., numbers[99] 更加方便且易于扩展。
//
//数组元素可以通过索引（位置）来读取（或者修改），索引从 0 开始，第一个元素索引为 0，第二个索引为 1，以此类推。
//
//声明数组
//Go 语言数组声明需要指定元素类型及元素个数，语法格式如下：
//
//var arrayName [size]dataType
//其中，arrayName 是数组的名称，size 是数组的大小，dataType 是数组中元素的数据类型。
//以下定义了数组 balance 长度为 10 类型为 float32：
//
//var balance [10]float32

//初始化数组
//以下演示了数组初始化：
//
//以下实例声明一个名为 numbers 的整数数组，其大小为 5，在声明时，数组中的每个元素都会根据其数据类型进行默认初始化，对于整数类型，初始值为 0。
//
//var numbers [5]int
//还可以使用初始化列表来初始化数组的元素：
//
//var numbers = [5]int{1, 2, 3, 4, 5}
//以上代码声明一个大小为 5 的整数数组，并将其中的元素分别初始化为 1、2、3、4 和 5。
//
//另外，还可以使用 := 简短声明语法来声明和初始化数组：
//
//numbers := [5]int{1, 2, 3, 4, 5}
//以上代码创建一个名为 numbers 的整数数组，并将其大小设置为 5，并初始化元素的值。
//
//注意：在 Go 语言中，数组的大小是类型的一部分，因此不同大小的数组是不兼容的，也就是说 [5]int 和 [10]int 是不同的类型。
//
//以下定义了数组 balance 长度为 5 类型为 float32，并初始化数组的元素：
//
//var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
//我们也可以通过字面量在声明数组的同时快速初始化数组：
//
//balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
//如果数组长度不确定，可以使用 ... 代替数组的长度，编译器会根据元素个数自行推断数组的长度：
//
//var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
//或
//balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
//如果设置了数组的长度，我们还可以通过指定下标来初始化元素：
//
////  将索引为 1 和 3 的元素初始化
//balance := [5]float32{1:2.0,3:7.0}
//初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
//
//如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：
//
// balance[4] = 50.0
//以上实例读取了第五个元素。数组元素可以通过索引（位置）来读取（或者修改），索引从 0 开始，第一个元素索引为 0，第二个索引为 1，以此类推。

// 访问数组元素
// 数组元素可以通过索引（位置）来读取。格式为数组名后加中括号，中括号中为索引的值。例如：
//
// var salary float32 = balance[9]
// 以上实例读取了数组 balance 第 10 个元素的值。
//
// 以下演示了数组完整操作（声明、赋值、访问）的实例：
//实例 1
//package main
//
//import "fmt"
//
//func main() {
//	var n [10]int /* n 是一个长度为 10 的数组 */
//	var i,j int
//
//	/* 为数组 n 初始化元素 */
//	for i = 0; i < 10; i++ {
//		n[i] = i + 100 /* 设置元素为 i + 100 */
//	}
//
//	/* 输出每个数组元素的值 */
//	for j = 0; j < 10; j++ {
//		fmt.Printf("Element[%d] = %d\n", j, n[j] )
//	}
//}
//以上实例执行结果如下：
//
//Element[0] = 100
//Element[1] = 101
//Element[2] = 102
//Element[3] = 103
//Element[4] = 104
//Element[5] = 105
//Element[6] = 106
//Element[7] = 107
//Element[8] = 108
//Element[9] = 109

//实例 2
//package main
//
//import "fmt"
//
//func main() {
//	var i,j,k int
//	// 声明数组的同时快速初始化数组
//	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
//
//	/* 输出数组元素 */         ...
//	for i = 0; i < 5; i++ {
//		fmt.Printf("balance[%d] = %f\n", i, balance[i] )
//	}
//
//	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
//	/* 输出每个数组元素的值 */
//	for j = 0; j < 5; j++ {
//		fmt.Printf("balance2[%d] = %f\n", j, balance2[j] )
//	}
//
//	//  将索引为 1 和 3 的元素初始化
//	balance3 := [5]float32{1:2.0,3:7.0}
//	for k = 0; k < 5; k++ {
//		fmt.Printf("balance3[%d] = %f\n", k, balance3[k] )
//	}
//}
//
//以上实例执行结果如下：
//
//balance[0] = 1000.000000
//balance[1] = 2.000000
//balance[2] = 3.400000
//balance[3] = 7.000000
//balance[4] = 50.000000
//balance2[0] = 1000.000000
//balance2[1] = 2.000000
//balance2[2] = 3.400000
//balance2[3] = 7.000000
//balance2[4] = 50.000000
//balance3[0] = 0.000000
//balance3[1] = 2.000000
//balance3[2] = 0.000000
//balance3[3] = 7.000000
//balance3[4] = 0.000000

//更多内容
//数组对 Go 语言来说是非常重要的，以下我们将介绍数组更多的内容：
//
//内容	描述
//多维数组	Go 语言支持多维数组，最简单的多维数组是二维数组
//向函数传递数组	你可以向函数传递数组参数

// Go 语言多维数组
// Go 语言数组Go 语言数组
//
// 多维数组是数组的数组，可以把它想象成一个数据表格或矩阵，其中每个元素都可以通过多个索引来访问。
//
// Go 语言中的多维数组可用于处理表格数据、矩阵运算或游戏棋盘等结构化的信息。
//
// 基本概念
// 一维数组：像一条直线上的点，只需要一个坐标（索引）就能找到特定元素
// 二维数组：像一个表格，需要行和列两个坐标
// 三维数组：像一个立方体，需要长、宽、高三个坐标
// 更高维度：理论上可以有多维，但实际编程中二维和三维最常用
// 可以从我们平时的停车位来比拟：
//
// 一维数组 = 一排停车位（只需要车位编号）
// 二维数组 = 多层停车场（需要楼层和车位编号）
// 三维数组 = 多个多层停车场（需要停车场编号、楼层和车位编号）
// Go 语言支持多维数组，以下为常用的多维数组声明方式：
//
// var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type
// 声明与初始化：
//
// // 声明一个二维数组 var 数组名 [行数][列数]元素类型 // 声明并初始化 var 数组名 [行数][列数]元素类型 = [行数][列数]元素类型{初始化值}
// 以下实例声明了三维的整型数组：
//
// var threedim [5][10][4]int
// 二维数组
// 二维数组是最简单的多维数组，二维数组本质上是由一维数组组成的。二维数组定义方式如下：
//
// var arrayName [ x ][ y ] variable_type
// variable_type 为 Go 语言的数据类型，arrayName 为数组名，二维数组可认为是一个表格，x 为行，y 为列
// 二维数组中的元素可通过 a[ i ][ j ] 来访问。
//
//	func main() {
//		// Step 1: 创建数组
//		values := [][]int{}
//
//		// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
//		row1 := []int{1, 2, 3}
//		row2 := []int{4, 5, 6}
//		values = append(values, row1)
//		values = append(values, row2)
//
//		// Step 3: 显示两行数据
//		fmt.Println("Row 1")
//		fmt.Println(values[0])
//		fmt.Println("Row 2")
//		fmt.Println(values[1])
//
//		// Step 4: 访问第一个元素
//		fmt.Println("第一个元素为：")
//		fmt.Println(values[0][0])
//	}
//
// 以上实例运行输出结果为：
//
// Row 1
// [1 2 3]
// Row 2
// [4 5 6]
// 第一个元素为：
// 1
// 初始化二维数组
// 多维数组可通过大括号来初始值。以下实例为一个 3 行 4 列的二维数组：
//
// a := [3][4]int{
// {0, 1, 2, 3} ,   /*  第一行索引为 0 */
// {4, 5, 6, 7} ,   /*  第二行索引为 1 */
// {8, 9, 10, 11},   /* 第三行索引为 2 */
// }
// 注意：以上代码中倒数第二行的 } 必须要有逗号，因为最后一行的 } 不能单独一行，也可以写成这样：
// a := [3][4]int{
// {0, 1, 2, 3} ,   /*  第一行索引为 0 */
// {4, 5, 6, 7} ,   /*  第二行索引为 1 */
// {8, 9, 10, 11}}   /* 第三行索引为 2 */
// 以下实例初始化一个 2 行 2 列 的二维数组：
//实例
//package main
//
//import "fmt"
//
//func main() {
//	// 创建二维数组
//	sites := [2][2]string{}
//
//	// 向二维数组添加元素
//	sites[0][0] = "Google"
//	sites[0][1] = "Runoob"
//	sites[1][0] = "Taobao"
//	sites[1][1] = "Weibo"
//
//	// 显示结果
//	fmt.Println(sites)
//}
//以上实例运行输出结果为：
//
//[[Google Runoob] [Taobao Weibo]]

// 访问二维数组
// 二维数组通过指定坐标来访问。如数组中的行索引与列索引，例如：
//
// val := a[2][3]
// 或
// var value int = a[2][3]
// 以上实例访问了二维数组 val 第三行的第四个元素。
//
// 二维数组可以使用循环嵌套来输出元素：
//func main() {
//	/* 数组 - 5 行 2 列*/
//	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
//	var i, j int
//
//	/* 输出数组元素 */
//	for i = 0; i < 5; i++ {
//		for j = 0; j < 2; j++ {
//			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
//		}
//	}
//}
//以上实例运行输出结果为：
//
//a[0][0] = 0
//a[0][1] = 0
//a[1][0] = 1
//a[1][1] = 2
//a[2][0] = 2
//a[2][1] = 4
//a[3][0] = 3
//a[3][1] = 6
//a[4][0] = 4
//a[4][1] = 8

// 以下实例创建各个维度元素数量不一致的多维数组：
//func main() {
//	// 创建空的二维数组
//	animals := [][]string{}
//
//	// 创建三一维数组，各数组长度不同
//	row1 := []string{"fish", "shark", "eel"}
//	row2 := []string{"bird"}
//	row3 := []string{"lizard", "salamander"}
//
//	// 使用 append() 函数将一维数组添加到二维数组中
//	animals = append(animals, row1)
//	animals = append(animals, row2)
//	animals = append(animals, row3)
//
//	// 循环输出
//	for i := range animals {
//		fmt.Printf("Row: %v\n", i)
//		fmt.Println(animals[i])
//	}
//}
//
//以上实例运行输出结果为：
//
//Row: 0
//[fish shark eel]
//Row: 1
//[bird]
//Row: 2
//[lizard salamander]

// 访问和修改数组元素
// 访问元素
//func main() {
//	// 创建一个3x3的矩阵
//	matrix := [3][3]int{
//		{1, 2, 3},
//		{4, 5, 6},
//		{7, 8, 9},
//	}
//
//	// 访问单个元素
//	fmt.Println("第一行第二列:", matrix[0][1]) // 输出: 2
//	fmt.Println("第三行第三列:", matrix[2][2]) // 输出: 9
//
//	// 访问整行
//	fmt.Println("第二行:", matrix[1]) // 输出: [4 5 6]
//
//	// 遍历所有元素
//	fmt.Println("\n遍历所有元素:")
//	for i := 0; i < 3; i++ {
//		for j := 0; j < 3; j++ {
//			fmt.Printf("matrix[%d][%d] = %d\n", i, j, matrix[i][j])
//		}
//	}
//}
//第一行第二列: 2
//第三行第三列: 9
//第二行: [4 5 6]
//
//遍历所有元素:
//matrix[0][0] = 1
//matrix[0][1] = 2
//matrix[0][2] = 3
//matrix[1][0] = 4
//matrix[1][1] = 5
//matrix[1][2] = 6
//matrix[2][0] = 7
//matrix[2][1] = 8
//matrix[2][2] = 9

// 修改元素
// 实例
// package main
//
// import "fmt"
//
//	func main() {
//	   // 创建一个2x2的零值矩阵
//	   var grid [2][2]int
//
//	   fmt.Println("修改前的矩阵:", grid)
//
//	   // 修改特定位置的元素
//	   grid[0][0] = 10
//	   grid[0][1] = 20
//	   grid[1][0] = 30
//	   grid[1][1] = 40
//
//	   fmt.Println("修改后的矩阵:", grid)
//
//	   // 批量修改一行
//	   grid[0] = [2]int{100, 200}
//	   fmt.Println("修改第一行后:", grid)
//	}

//实例
//func main() {
//	// 创建一个2x2的零值矩阵
//	var grid [2][2]int
//
//	fmt.Println("修改前的矩阵:", grid)
//
//	// 修改特定位置的元素
//	grid[0][0] = 10
//	grid[0][1] = 20
//	grid[1][0] = 30
//	grid[1][1] = 40
//
//	fmt.Println("修改后的矩阵:", grid)
//
//	// 批量修改一行
//	grid[0] = [2]int{100, 200}
//	fmt.Println("修改第一行后:", grid)
//}
//修改前的矩阵: [[0 0] [0 0]]
//修改后的矩阵: [[10 20] [30 40]]
//修改第一行后: [[100 200] [30 40]]

// 三维及更高维数组
// 三维数组示例
// 实例
//func main() {
//	// 声明一个2x3x4的三维数组
//	// 可以理解为：2个平面，每个平面有3行4列
//	var cube [2][3][4]int
//
//	// 初始化三维数组
//	cube = [2][3][4]int{
//		{ // 第一个平面
//			{1, 2, 3, 4},
//			{5, 6, 7, 8},
//			{9, 10, 11, 12},
//		},
//		{ // 第二个平面
//			{13, 14, 15, 16},
//			{17, 18, 19, 20},
//			{21, 22, 23, 24},
//		},
//	}
//
//	// 访问三维数组元素
//	fmt.Println("cube[0][1][2] =", cube[0][1][2]) // 输出: 7
//	fmt.Println("cube[1][2][3] =", cube[1][2][3]) // 输出: 24
//
//	// 遍历三维数组
//	fmt.Println("\n三维数组内容:")
//	for i := 0; i < 2; i++ {
//		fmt.Printf("平面 %d:\n", i)
//		for j := 0; j < 3; j++ {
//			for k := 0; k < 4; k++ {
//				fmt.Printf("%3d ", cube[i][j][k])
//			}
//			fmt.Println()
//		}
//		fmt.Println()
//	}
//}
//cube[0][1][2] = 7
//cube[1][2][3] = 24
//
//三维数组内容:
//平面 0:
//1   2   3   4
//5   6   7   8
//9  10  11  12
//
//平面 1:
//13  14  15  16
//17  18  19  20
//21  22  23  24

// 多维数组的常用操作
// 1. 使用 range 遍历
//func main() {
//	// 创建一个二维数组
//	scores := [3][4]int{
//		{85, 90, 78, 92},
//		{88, 76, 95, 89},
//		{92, 85, 88, 90},
//	}
//
//	fmt.Println("学生成绩表:")
//
//	// 使用 range 遍历二维数组
//	for i, row := range scores {
//		fmt.Printf("学生 %d 的成绩: ", i+1)
//		for j, score := range row {
//			fmt.Printf("%d ", score)
//			// 如果需要索引和值都使用
//			_ = j // 避免未使用变量警告
//		}
//		fmt.Println()
//	}
//
//	// 只关心值，不关心索引
//	total := 0
//	count := 0
//	for _, row := range scores {
//		for _, score := range row {
//			total += score
//			count++
//		}
//	}
//	fmt.Printf("\n平均分: %.2f\n", float64(total)/float64(count))
//}
//学生成绩表:
//学生 1 的成绩: 85 90 78 92
//学生 2 的成绩: 88 76 95 89
//学生 3 的成绩: 92 85 88 90
//
//平均分: 87.33

// 2. 数组长度获取
// 实例
//func main() {
//	// 创建一个不规则的多维数组
//	jagged := [3][3]int{
//		{1, 2, 3},
//		{4, 5},
//		{6, 7, 8, 9}, // 注意：这里会编译错误，因为每行必须长度一致
//	}
//
//	// 正确的示例：获取数组维度
//	matrix := [4][5]int{}
//
//	// 获取行数
//	rows := len(matrix)
//	fmt.Println("行数:", rows) // 输出: 4
//
//	// 获取第一行的列数（所有行长度相同）
//	cols := len(matrix[0])
//	fmt.Println("列数:", cols) // 输出: 5
//
//	// 获取总元素数
//	totalElements := rows * cols
//	fmt.Println("总元素数:", totalElements) // 输出: 20
//}

// 3. 数组比较
// 实例
//func main() {
//	// 创建两个相同的二维数组
//	a := [2][2]int{{1, 2}, {3, 4}}
//	b := [2][2]int{{1, 2}, {3, 4}}
//	c := [2][2]int{{1, 2}, {3, 5}}
//
//	// 数组可以直接比较（只有当维度完全相同时）
//	fmt.Println("a == b:", a == b) // 输出: true
//	fmt.Println("a == c:", a == c) // 输出: false
//
//	// 注意：不同维度的数组不能比较
//	//d := [2][3]int{{1, 2, 3}, {4, 5, 6}}
//	//fmt.Println(a == d) // 编译错误：类型不匹配
//}

// 实际应用场景
// 场景1：游戏棋盘（井字棋）
// 实例
//func main() {
//	// 初始化一个3x3的井字棋棋盘
//	var board [3][3]string
//
//	// 初始化为空
//	for i := 0; i < 3; i++ {
//		for j := 0; j < 3; j++ {
//			board[i][j] = " "
//		}
//	}
//
//	// 模拟下棋
//	board[0][0] = "X"
//	board[1][1] = "O"
//	board[2][2] = "X"
//
//	// 打印棋盘
//	fmt.Println("井字棋棋盘:")
//	for i := 0; i < 3; i++ {
//		for j := 0; j < 3; j++ {
//			fmt.Printf(" %s ", board[i][j])
//			if j < 2 {
//				fmt.Printf("|")
//			}
//		}
//		fmt.Println()
//		if i < 2 {
//			fmt.Println("---+---+---")
//		}
//	}
//}

// 场景2：学生成绩管理系统
// 实例
//func main() {
//	// 定义：3个学生，每个学生有4门课程
//	var grades [3][4]float64
//
//	// 输入学生成绩
//	grades = [3][4]float64{
//		{85.5, 90.0, 78.5, 92.0}, // 学生1的成绩
//		{88.0, 76.5, 95.0, 89.5}, // 学生2的成绩
//		{92.5, 85.0, 88.5, 90.0}, // 学生3的成绩
//	}
//
//	// 计算每个学生的平均分
//	fmt.Println("学生成绩统计:")
//	for i, studentGrades := range grades {
//		sum := 0.0
//		for _, grade := range studentGrades {
//			sum += grade
//		}
//		average := sum / float64(len(studentGrades))
//		fmt.Printf("学生 %d: 平均分 = %.2f\n", i+1, average)
//	}
//
//	// 计算每门课程的平均分
//	fmt.Println("\n课程平均分:")
//	for j := 0; j < 4; j++ {
//		sum := 0.0
//		for i := 0; i < 3; i++ {
//			sum += grades[i][j]
//		}
//		average := sum / 3.0
//		fmt.Printf("课程 %d: 平均分 = %.2f\n", j+1, average)
//	}
//}

// 场景3：图像像素处理
// 实例
//func main() {
//	// 模拟一个简单的3x3灰度图像
//	// 每个像素值范围：0(黑色) ~ 255(白色)
//	var image [3][3]int
//
//	// 初始化图像（一个简单的渐变）
//	for i := 0; i < 3; i++ {
//		for j := 0; j < 3; j++ {
//			image[i][j] = (i + j) * 50
//		}
//	}
//
//	// 显示原始图像
//	fmt.Println("原始图像:")
//	displayImage(image)
//
//	// 图像处理：增加亮度
//	fmt.Println("\n增加亮度后的图像:")
//	for i := 0; i < 3; i++ {
//		for j := 0; j < 3; j++ {
//			// 增加亮度，但不超过255
//			newValue := image[i][j] + 50
//			if newValue > 255 {
//				newValue = 255
//			}
//			image[i][j] = newValue
//		}
//	}
//	displayImage(image)
//}
//
//func displayImage(img [3][3]int) {
//	for i := 0; i < 3; i++ {
//		for j := 0; j < 3; j++ {
//			fmt.Printf("%3d ", img[i][j])
//		}
//		fmt.Println()
//	}
//}
//原始图像:
//0  50 100
//50 100 150
//100 150 200
//
//增加亮度后的图像:
//50 100 150
//100 150 200
//150 200 250

// 注意事项和最佳实践
// 1. 数组长度是类型的一部分
//
//	func main() {
//		// 这两个是不同的类型！
//		var a [2][3]int
//		var b [3][2]int
//
//		// 以下代码会编译错误
//		// a = b  // 错误：类型不匹配
//
//		fmt.Printf("a 的类型: %T\n", a) // [2][3]int
//		fmt.Printf("b 的类型: %T\n", b) // [3][2]int
//	}
//
// 2. 值类型 vs 引用类型
//func main() {
//	// 数组是值类型
//	original := [2][2]int{{1, 2}, {3, 4}}
//
//	// 赋值会创建副本
//	copy := original
//
//	// 修改副本不会影响原始数组
//	copy[0][0] = 100
//
//	fmt.Println("原始数组:", original) // [[1 2] [3 4]]
//	fmt.Println("副本数组:", copy)     // [[100 2] [3 4]]
//
//	// 如果需要引用语义，可以使用切片（后续文章会介绍）
//}
//3. 性能考虑
//内存连续：多维数组在内存中是连续存储的，访问速度快
//固定大小：数组长度在编译时确定，无法动态改变
//适合场景：当数据大小已知且固定时，数组是最佳选择

//常见问题解答
//Q1: 多维数组和嵌套切片有什么区别？
//特性	多维数组	嵌套切片
//大小	固定，编译时确定	动态，运行时可变
//内存	连续分配	可能不连续
//性能	访问速度快	稍慢，有额外开销
//使用场景	数据大小已知	数据大小可变

//Q2: 如何创建不规则的二维数组？
//Go 的数组要求每行长度相同。如果需要不规则结构，应该使用切片：
//
//实例
//// 使用切片创建不规则结构
//irregular := [][]int{
//    {1, 2, 3},
//    {4, 5},          // 这行只有2个元素
//    {6, 7, 8, 9},    // 这行有4个元素
//}

//Q3: 多维数组可以作为函数参数吗？
//可以，但要注意数组是值类型，传递大数组会有性能开销：
//
//实例
//func processMatrix(matrix [3][3]int) [3][3]int {
//    // 处理矩阵...
//    matrix[0][0] = 100
//    return matrix
//}
//
//// 更好的方式是使用指针或切片
//func processMatrixPtr(matrix *[3][3]int) {
//    matrix[0][0] = 100
//}

// Go 语言向函数传递数组
// Go 语言数组Go 语言数组
//
// Go 语言中的数组是值类型，因此在将数组传递给函数时，实际上是传递数组的副本。
//
// 如果你想向函数传递数组参数，你需要在函数定义时，声明形参为数组，我们可以通过以下两种方式来声明：
//
// 方式一
// 形参设定数组大小：
//
//	func myFunction(param [10]int) {
//	   ....
//	}
//
// 方式二
// 形参未设定数组大小：
//
//	func myFunction(param []int) {
//	   ....
//	}
//
// 如果你想要在函数内修改原始数组，可以通过传递数组的指针来实现。
//
// 实例
// 让我们看下以下实例，实例中函数接收整型数组参数，另一个参数指定了数组元素的个数，并返回平均值：
// 实例
// func getAverage(arr []int, size int) float32
//
//	{
//	  var i int
//	  var avg, sum float32
//
//	  for i = 0; i < size; ++i {
//	     sum += arr[i]
//	  }
//
//	  avg = sum / size
//
//	  return avg;
//	}
//
// 接下来我们来调用这个函数：
//
// 实例
// package main
//
// import "fmt"
//
//	func main() {
//	  /* 数组长度为 5 */
//	  var  balance = [5]int {1000, 2, 3, 17, 50}
//	  var avg float32
//
//	  /* 数组作为参数传递给函数 */
//	  avg = getAverage( balance, 5 ) ;
//
//	  /* 输出返回的平均值 */
//	  fmt.Printf( "平均值为: %f ", avg );
//	}
//
//	func getAverage(arr []int, size int) float32 {
//	  var i,sum int
//	  var avg float32
//
//	  for i = 0; i < size;i++ {
//	     sum += arr[i]
//	  }
//
//	  avg = float32(sum) / float32(size)
//
//	  return avg;
//	}
//
// 以上实例执行输出结果为：
//
// 平均值为: 214.399994
// 以上实例中我们使用的形参并未设定数组大小。
//
// 浮点数计算输出有一定的偏差，你也可以转整型来设置精度。
//实例
//package main
//import (
//    "fmt"
//)
//func main() {
//    a := 1.69
//    b := 1.7
//    c := a * b      // 结果应该是2.873
//    fmt.Println(c)  // 输出的是2.8729999999999998
//}
//设置固定精度：
//
//实例
//package main
//import (
//    "fmt"
//)
//func main() {
//    a := 1690           // 表示1.69
//    b := 1700           // 表示1.70
//    c := a * b          // 结果应该是2873000表示 2.873
//    fmt.Println(c)      // 内部编码
//    fmt.Println(float64(c) / 1000000) // 显示
//}

// 如果你想要在函数内修改原始数组，可以通过传递数组的指针来实现。
//
// 以下实例演示如何向函数传递数组，函数接受一个数组和数组的指针作为参数：
// 函数接受一个数组作为参数
//实例
//package main
//
//import "fmt"
//
//// 函数接受一个数组作为参数
//func modifyArray(arr [5]int) {
//    for i := 0; i < len(arr); i++ {
//        arr[i] = arr[i] * 2
//    }
//}
//
//// 函数接受一个数组的指针作为参数
//func modifyArrayWithPointer(arr *[5]int) {
//    for i := 0; i < len(*arr); i++ {
//        (*arr)[i] = (*arr)[i] * 2
//    }
//}
//
//func main() {
//    // 创建一个包含5个元素的整数数组
//    myArray := [5]int{1, 2, 3, 4, 5}
//
//    fmt.Println("Original Array:", myArray)
//
//    // 传递数组给函数，但不会修改原始数组的值
//    modifyArray(myArray)
//    fmt.Println("Array after modifyArray:", myArray)
//
//    // 传递数组的指针给函数，可以修改原始数组的值
//    modifyArrayWithPointer(&myArray)
//    fmt.Println("Array after modifyArrayWithPointer:", myArray)
//}
//在上面的例子中，modifyArray 函数接受一个数组，并尝试修改数组的值，但在主函数中调用后，原始数组并未被修改。相反，modifyArrayWithPointer 函数接受一个数组的指针，并通过指针修改了原始数组的值。
//
//以上实例执行输出结果为：
//
//Original Array: [1 2 3 4 5]
//Array after modifyArray: [1 2 3 4 5]
//Array after modifyArrayWithPointer: [2 4 6 8 10]
