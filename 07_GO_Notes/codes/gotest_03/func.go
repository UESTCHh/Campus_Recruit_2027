package main

// Go 语言函数
// 函数是基本的代码块，用于执行一个任务。
//
// Go 语言最少有个 main() 函数。
//
// 你可以通过函数来划分不同功能，逻辑上每个函数执行的是指定的任务。
//
// 函数声明告诉了编译器函数的名称，返回类型，和参数。
//
// Go 语言标准库提供了多种可动用的内置的函数。例如，len() 函数可以接受不同类型参数并返回该类型的长度。如果我们传入的是字符串则返回字符串的长度，如果传入的是数组，则返回数组中包含的元素个数。
//
// 函数定义
// Go 语言函数定义格式如下：
//
//	func function_name( [parameter list] ) [return_types] {
//	  函数体
//	}
//
// 函数定义解析：
//
// func：函数由 func 开始声明
// function_name：函数名称，参数列表和返回值类型构成了函数签名。
// parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
// return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
// 函数体：函数定义的代码集合。
// 实例
// 以下实例为 max() 函数的代码，该函数传入两个整型参数 num1 和 num2，并返回这两个参数的最大值：
/* 函数返回两个数的最大值 */
//func max(num1, num2 int) int {
//	/* 声明局部变量 */
//	var result int
//
//	if num1 >= num2 {
//		result = num1
//	} else {
//		result = num2
//	}
//
//	return result
//}
//
//func main() {
//	fmt.Println(max(10, 2))
//}

// 函数调用
// 当创建函数时，你定义了函数需要做什么，通过调用该函数来执行指定任务。
//
// 调用函数，向函数传递参数，并返回值，例如：
//func main() {
//	/* 定义局部变量 */
//	var a int = 100
//	var b int = 200
//	var ret int
//
//	/* 调用函数并返回最大值 */
//	ret = max(a, b)
//
//	fmt.Printf("最大值是 : %d\n", ret)
//}
//
///* 函数返回两个数的最大值 */
//func max(num1, num2 int) int {
//	/* 定义局部变量 */
//	var result int
//
//	if num1 >= num2 {
//		result = num1
//	} else {
//		result = num2
//	}
//	return result
//}

//以上实例在 main() 函数中调用 max（）函数，执行结果为：
//
//最大值是 : 200

// 函数返回多个值
// Go 函数可以返回多个值，例如：
//func swap(x, y string) (string, string) {
//	return y, x
//}
//
//func main() {
//	a, b := swap("hello", "world")
//	fmt.Println(a, b)
//}
//以上实例执行结果为：world hello

//函数参数
//函数如果使用参数，该变量可称为函数的形参。
//
//形参就像定义在函数体内的局部变量。
//
//调用函数，可以通过两种方式来传递参数：
//
//传递类型	描述
//值传递	值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
//引用传递	引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
//默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。

//Go 语言函数值传递值
//Go 函数Go 函数
//
//传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
//
//默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。
//
//以下定义了 swap() 函数：
//
///* 定义相互交换值的函数 */
//func swap(x, y int) int {
//   var temp int
//
//   temp = x /* 保存 x 的值 */
//   x = y    /* 将 y 值赋给 x */
//   y = temp /* 将 temp 值赋给 y*/
//
//   return temp;
//}
//接下来，让我们使用值传递来调用 swap() 函数：
//
//实例
//package main
//
//import "fmt"
//
//func main() {
//   /* 定义局部变量 */
//   var a int = 100
//   var b int = 200
//
//   fmt.Printf("交换前 a 的值为 : %d\n", a )
//   fmt.Printf("交换前 b 的值为 : %d\n", b )
//
//   /* 通过调用函数来交换值 */
//   swap(a, b)
//
//   fmt.Printf("交换后 a 的值 : %d\n", a )
//   fmt.Printf("交换后 b 的值 : %d\n", b )
//}
//
///* 定义相互交换值的函数 */
//func swap(x, y int) int {
//   var temp int
//
//   temp = x /* 保存 x 的值 */
//   x = y    /* 将 y 值赋给 x */
//   y = temp /* 将 temp 值赋给 y*/
//
//   return temp;
//}
//以下代码执行结果为：
//
//交换前 a 的值为 : 100
//交换前 b 的值为 : 200
//交换后 a 的值 : 100
//交换后 b 的值 : 200
//程序中使用的是值传递, 所以两个值并没有实现交互，我们可以使用 引用传递 来实现交换效果。

//Go 语言函数引用传递值
//Go 函数Go 函数
//
//引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
//
//引用传递指针参数传递到函数内，以下是交换函数 swap() 使用了引用传递：
//
///* 定义交换值函数*/
//func swap(x *int, y *int) {
//   var temp int
//   temp = *x    /* 保持 x 地址上的值 */
//   *x = *y      /* 将 y 值赋给 x */
//   *y = temp    /* 将 temp 值赋给 y */
//}
//以下我们通过使用引用传递来调用 swap() 函数：
//
//package main
//
//import "fmt"
//
//func main() {
//   /* 定义局部变量 */
//   var a int = 100
//   var b int= 200
//
//   fmt.Printf("交换前，a 的值 : %d\n", a )
//   fmt.Printf("交换前，b 的值 : %d\n", b )
//
//   /* 调用 swap() 函数
//   * &a 指向 a 指针，a 变量的地址
//   * &b 指向 b 指针，b 变量的地址
//   */
//   swap(&a, &b)
//
//   fmt.Printf("交换后，a 的值 : %d\n", a )
//   fmt.Printf("交换后，b 的值 : %d\n", b )
//}
//
//func swap(x *int, y *int) {
//   var temp int
//   temp = *x    /* 保存 x 地址上的值 */
//   *x = *y      /* 将 y 值赋给 x */
//   *y = temp    /* 将 temp 值赋给 y */
//}
//以上代码执行结果为：
//
//交换前，a 的值 : 100
//交换前，b 的值 : 200
//交换后，a 的值 : 200
//交换后，b 的值 : 100

//函数用法
//函数用法	描述
//函数作为另外一个函数的实参	函数定义后可作为另外一个函数的实参数传入
//闭包	闭包是匿名函数，可在动态编程中使用
//方法	方法就是一个包含了接受者的函数
/* 定义一个函数，接收一个函数类型的参数 op */
/* op 的类型为 func(float64, float64) float64，表示接受两个 float64 参数、返回 float64 的函数 */
//func apply(op func(float64, float64) float64, a, b float64) float64 {
//	return op(a, b)
//}
//
//func main() {
//	/* 定义加法函数 */
//	add := func(a, b float64) float64 {
//		return a + b
//	}
//
//	/* 定义乘法函数 */
//	multiply := func(a, b float64) float64 {
//		return a * b
//	}
//
//	/* 将 add 函数作为实参传入 apply */
//	fmt.Println(apply(add, 3, 4)) // 输出：7
//
//	/* 将 multiply 函数作为实参传入 apply */
//	fmt.Println(apply(multiply, 3, 4)) // 输出：12
//
//	/* 也可以直接传入匿名函数 */
//	fmt.Println(apply(func(a, b float64) float64 {
//		return math.Pow(a, b) // a 的 b 次方
//	}, 2, 10)) // 输出：1024
//}
//以上代码执行结果为：
//
//7
//12
//1024

// 实际应用：自定义排序规则
// 函数作为实参最常见的应用之一是自定义排序规则。Go 标准库的 sort.Slice 就接受一个比较函数作为参数：
//func main() {
//	nums := []int{5, 2, 8, 1, 9, 3}
//
//	/* 将比较函数作为实参传入 sort.Slice，实现从小到大排序 */
//	sort.Slice(nums, func(i, j int) bool {
//		return nums[i] < nums[j]
//	})
//	fmt.Println("升序：", nums) // 输出：升序：[1 2 3 5 8 9]
//
//	/* 传入不同的比较函数，实现从大到小排序 */
//	sort.Slice(nums, func(i, j int) bool {
//		return nums[i] > nums[j]
//	})
//	fmt.Println("降序：", nums) // 输出：降序：[9 8 5 3 2 1]
//}
//以上代码执行结果为：
//
//升序：[1 2 3 5 8 9]
//降序：[9 8 5 3 2 1]
//通过将不同的函数作为实参传入，同一个 sort.Slice 就可以实现完全不同的排序逻辑，这正是函数作为实参的核心价值——将行为参数化，让代码更灵活复用。

// Go 语言函数闭包（匿名函数）
// Go 函数Go 函数
//
// Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
//
// 匿名函数是一种没有函数名的函数，通常用于在函数内部定义函数，或者作为函数参数进行传递。
//
// 以下实例中，我们创建了函数 getSequence() ，返回另外一个函数。该函数的目的是在闭包中递增 i 变量，代码如下：
//func getSequence() func() int {
//	i := 0
//	return func() int {
//		i += 1
//		return i
//	}
//}
//
//func main() {
//	/* nextNumber 为一个函数，函数 i 为 0 */
//	nextNumber := getSequence()
//
//	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
//	fmt.Println(nextNumber())
//	fmt.Println(nextNumber())
//	fmt.Println(nextNumber())
//
//	/* 创建新的函数 nextNumber1，并查看结果 */
//	nextNumber1 := getSequence()
//	fmt.Println(nextNumber1())
//	fmt.Println(nextNumber1())
//}
//以上代码执行结果为：
//
//1
//2
//3
//1
//2
//这正是 Go 语言中闭包的核心特性。
//
//简单来说：i := 0 只会在调用 getSequence() 时执行一次，而不是在调用返回的那个函数时执行。
//
//我们来拆解这个过程：
//
//1. nextNumber := getSequence() 发生了什么？
//调用 getSequence()，函数体开始执行。
//
//执行 i := 0，创建一个新的局部变量 i，并初始化为 0。
//
//返回一个匿名函数 func() int { i+=1; return i }。
//
//关键点：这个匿名函数捕获了外部的变量 i。它记住的是那个特定的 i，而不是值 0。
//
//getSequence() 执行结束，但它的局部变量 i 并没有被销毁，因为返回的匿名函数还引用着它。i 会逃逸到堆上，与这个闭包函数绑定在一起。
//
//nextNumber 现在就是这个闭包，带着一个私有的 i，当前值为 0。
//
//2. 调用 nextNumber() 时
//执行闭包内的代码：i += 1，它找到自己绑定的那个 i（目前是 0），加 1 变成 1，返回 1。
//
//注意：不会再执行 i := 0，因为那行代码在 getSequence 里，而 getSequence 早就返回了。
//
//第二次调用，同一个 i 从 1 变成 2，返回 2。
//
//第三次调用，从 2 变成 3，返回 3。
//
//所以输出了 1, 2, 3。
//
//3. nextNumber1 := getSequence() 再次调用
//这次又重新调用了 getSequence()，所以会重新执行 i := 0，创建一个全新的、独立的变量 i，初始值为 0。
//
//返回一个新的闭包，绑定着这个新的 i。
//
//nextNumber1() 调用时，它操作的是这个新的 i，从 0 变成 1，所以输出 1。
//
//第二次 nextNumber1() 输出 2。
//
//为什么不是 1, 1, 1, 1, 1？
//如果每次调用闭包时都重置 i，那就意味着每次调用都会执行 i := 0。但 i := 0 不在闭包的定义里，它在 getSequence 函数体里。闭包捕获的是变量本身，不是变量的副本，也不是重新执行外层代码。 它维持了那个变量的“记忆”。
//
//可以类比为：getSequence 是一个工厂，每调用一次就生产一个计数器（闭包）。每个计数器内部有自己的计数数字（i），工厂只在新生产时把数字归零（i := 0）。你按下同一个计数器，它的数字会累加；你换一个新计数器，它又从 1 开始。
//
//这就是闭包实现“有状态函数”的经典方式

// 以下实例我们定义了多个匿名函数，并展示了如何将匿名函数赋值给变量、在函数内部使用匿名函数以及将匿名函数作为参数传递给其他函数。
//func main() {
//	// 定义一个匿名函数并将其赋值给变量add
//	add := func(a, b int) int {
//		return a + b
//	}
//
//	// 调用匿名函数
//	result := add(3, 5)
//	fmt.Println("3 + 5 =", result)
//
//	// 在函数内部使用匿名函数
//	multiply := func(x, y int) int {
//		return x * y
//	}
//
//	product := multiply(4, 6)
//	fmt.Println("4 * 6 =", product)
//
//	// 将匿名函数作为参数传递给其他函数
//	calculate := func(operation func(int, int) int, x, y int) int {
//		return operation(x, y)
//	}
//
//	sum := calculate(add, 2, 8)
//	fmt.Println("2 + 8 =", sum)
//
//	// 也可以直接在函数调用中定义匿名函数
//	difference := calculate(func(a, b int) int {
//		return a - b
//	}, 10, 4)
//	fmt.Println("10 - 4 =", difference)
//}

//以上代码执行结果为：
//
//3 + 5 = 8
//4 * 6 = 24
//2 + 8 = 10
//10 - 4 = 6
//匿名函数的使用在 Go 语言中非常灵活，可以帮助简化代码结构和提高代码的可读性。

//Go 语言函数方法
//Go 函数Go 函数
//
//Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。语法格式如下：
//
//func (variable_name variable_data_type) function_name() [return_type]{
//   /* 函数体*/
//}
//下面定义一个结构体类型和该类型的一个方法：
///* 定义结构体 */
//type Circle struct {
//	radius float64
//}
//
//func main() {
//	var c1 Circle
//	c1.radius = 10.00
//	fmt.Println("圆的面积 = ", c1.getArea())
//}
//
//// 该 method 属于 Circle 类型对象中的方法
//func (c Circle) getArea() float64 {
//	//c.radius 即为 Circle 类型对象中的属性
//	return 3.14 * c.radius * c.radius
//}
//以上代码执行结果为：
//
//圆的面积 =  314
