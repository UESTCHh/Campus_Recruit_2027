package main

//Go 语言条件语句
//条件语句需要开发者通过指定一个或多个条件，并通过测试条件是否为 true 来决定是否执行指定语句，并在条件为 false 的情况在执行另外的语句。

//Go 语言提供了以下几种条件判断语句：
//
//语句	描述
//if 语句	if 语句 由一个布尔表达式后紧跟一个或多个语句组成。
//if...else 语句	if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。
//if 嵌套语句	你可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。
//switch 语句	switch 语句用于基于不同条件执行不同动作。
//select 语句	select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。
//注意：Go 没有三目运算符，所以不支持 ?: 形式的条件判断。

// 语法
// Go 编程语言中 if 语句的语法如下：
//
//	if 布尔表达式 {
//	  /* 在布尔表达式为 true 时执行 */
//	}
//
// If 在布尔表达式为 true 时，其后紧跟的语句块执行，如果为 false 则不执行。
// 实例
// 使用 if 判断一个数变量的大小：
//func main() {
// /* 定义局部变量 */
//	var a int = 10
// /* 使用 if 语句判断布尔表达式 */
//	if a < 20 {
//      /* 如果条件为 true 则执行以下语句 */
//		fmt.Println("a < 20")
//	}
//	fmt.Printf("a = %d\n", a)
//}
//以上代码执行结果为：
//
//a 小于 20
//a 的值为 : 10

// Go 语言 if...else 语句
// 语法
// Go 编程语言中 if...else 语句的语法如下：
//
// if 布尔表达式 {
// /* 在布尔表达式为 true 时执行 */
// } else {
// /* 在布尔表达式为 false 时执行 */
// }
// If 在布尔表达式为 true 时，其后紧跟的语句块执行，如果为 false 则执行 else 语句块。
// 实例
// 使用 if else 判断一个数的大小：
//func main() {
//	/* 局部变量定义 */
//	var a int = 100
//
//	/* 判断布尔表达式 */
//	if a < 20 {
//		/* 如果条件为 true 则执行以下语句 */
//		fmt.Println("a < 20")
//	} else {
//		/* 如果条件为 false 则执行以下语句 */
//		fmt.Println("a >= 20")
//	}
//	fmt.Printf("a = %d\n", a)
//}
//以上代码执行结果为：
//
//a 不小于 20
//a 的值为 : 100

// Go 语言 if 语句嵌套
// 语法
// Go 编程语言中 if...else 语句的语法如下：
//
//	if 布尔表达式 1 {
//		/* 在布尔表达式 1 为 true 时执行 */
//		if 布尔表达式 2 {
//		/* 在布尔表达式 2 为 true 时执行 */
//		}
//	}
//
// 你可以以同样的方式在 if 语句中嵌套 else if...else 语句
//
// 实例
// 嵌套使用 if 语句：
//func main() {
//	/* 定义局部变量 */
//	var a int = 100
//	var b int = 200
//
//	/* 判断条件 */
//	if a == 100 {
//		/* if 条件语句为 true 执行 */
//		if b == 200 {
//			/* if 条件语句为 true 执行 */
//			fmt.Printf("a 的值为 100 ， b 的值为 200\n" );
//		}
//	}
//	fmt.Printf("a 值为 : %d\n", a );
//	fmt.Printf("b 值为 : %d\n", b );
//}
//
//以上代码执行结果为：
//
//a 的值为 100 ， b 的值为 200
//a 值为 : 100
//b 值为 : 200

// Go 语言 switch 语句
// switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。
//
// switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加 break。
//
// switch 默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case，如果我们需要执行后面的 case，可以使用 fallthrough 。
//
// 语法
// Go 编程语言中 switch 语句的语法如下：
//
//	switch var1 {
//		case val1:
//		...
//		case val2:
//		...
//		default:
//		...
//	}
//
// 变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式。
// 您可以同时测试多个可能符合条件的值，使用逗号分割它们，例如：case val1, val2, val3。
//func main() {
//	/* 定义局部变量 */
//	var grade string = "B"
//	var marks int = 90
//
//	switch marks {
//	case 90:
//		grade = "A"
//	case 80:
//		grade = "B"
//	case 50, 60, 70:
//		grade = "C"
//	default:
//		grade = "D"
//	}
//
//	switch {
//	case grade == "A":
//		fmt.Printf("优秀!\n")
//	case grade == "B", grade == "C":
//		fmt.Printf("良好\n")
//	case grade == "D":
//		fmt.Printf("及格\n")
//	case grade == "F":
//		fmt.Printf("不及格\n")
//	default:
//		fmt.Printf("差\n")
//	}
//	fmt.Printf("你的等级是 %s\n", grade)
//}
//以上代码执行结果为：
//
//优秀!
//你的等级是 A

// Type Switch
// switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
//
// Type Switch 语法格式如下：
//
//	switch x.(type){
//		case type:
//		statement(s);
//		case type:
//		statement(s);
//		/* 你可以定义任意个数的case */
//		default: /* 可选 */
//		statement(s);
//	}
//func main() {
//	var x interface{}
//	switch i := x.(type) {
//	case nil:
//		fmt.Printf(" x 的类型 :%T", i)
//	case int:
//		fmt.Printf("x 是 int 型")
//	case float64:
//		fmt.Printf("x 是 float64 型")
//	case func(int) float64:
//		fmt.Printf("x 是 func(int) 型")
//	case bool, string:
//		fmt.Printf("x 是 bool 或 string 型")
//	default:
//		fmt.Printf("未知型")
//	}
//}
//以上代码执行结果为：
//
//x 的类型 :<nil>

// fallthrough
// 使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
//func main() {
//
//	switch {
//	case false:
//		fmt.Println("1、case 条件语句为 false")
//		fallthrough
//	case true:
//		fmt.Println("2、case 条件语句为 true")
//		fallthrough
//	case false:
//		fmt.Println("3、case 条件语句为 false")
//		fallthrough
//	case true:
//		fmt.Println("4、case 条件语句为 true")
//	case false:
//		fmt.Println("5、case 条件语句为 false")
//		fallthrough
//	default:
//		fmt.Println("6、默认 case")
//	}
//}
//以上代码执行结果为：
//
//2、case 条件语句为 true
//3、case 条件语句为 false
//4、case 条件语句为 true
//从以上代码输出的结果可以看出：switch 从第一个判断表达式为 true 的 case 开始执行，如果 case 带有 fallthrough，程序会继续执行下一条 case，且它不会去判断下一个 case 的表达式是否为 true。

// Go 语言 select 语句
//
// select 是 Go 中的一个控制结构，类似于 switch 语句。
//
// select 语句只能用于通道操作，每个 case 必须是一个通道操作，要么是发送要么是接收。
//
// select 语句会监听所有指定的通道上的操作，一旦其中一个通道准备好就会执行相应的代码块。
//
// 如果多个通道都准备好，那么 select 语句会随机选择一个通道执行。如果所有通道都没有准备好，那么执行 default 块中的代码。
//
// 语法
// Go 编程语言中 select 语句的语法如下：
//
//	select {
//	 case <- channel1:
//	   // 执行的代码
//	 case value := <- channel2:
//	   // 执行的代码
//	 case channel3 <- value:
//	   // 执行的代码
//
//	   // 你可以定义任意数量的 case
//
//	 default:
//	   // 所有通道都没有准备好，执行的代码
//	}
//
// 以下描述了 select 语句的语法：
//
// 每个 case 都必须是一个通道
// 所有 channel 表达式都会被求值
// 所有被发送的表达式都会被求值
// 如果任意某个通道可以进行，它就执行，其他被忽略。
// 如果有多个 case 都可以运行，select 会随机公平地选出一个执行，其他不会执行。
// 否则：
// 如果有 default 子句，则执行该语句。
// 如果没有 default 子句，select 将阻塞，直到某个通道可以运行；Go 不会重新对 channel 或值进行求值。
// 实例
// select 语句应用演示：
// func main() {
//
//		c1 := make(chan string)
//		c2 := make(chan string)
//
//		go func() {
//			time.Sleep(1 * time.Second)
//			c1 <- "one"
//		}()
//		go func() {
//			time.Sleep(2 * time.Second)
//			c2 <- "two"
//		}()
//
//		for i := 0; i < 2; i++ {
//			select {
//			case msg1 := <-c1:
//				fmt.Println("received", msg1)
//			case msg2 := <-c2:
//				fmt.Println("received", msg2)
//			}
//		}
//	}
//
// 以上代码执行结果为：
//
// received one
// received two
// 以上实例中，我们创建了两个通道 c1 和 c2。
//
// select 语句等待两个通道的数据。如果接收到 c1 的数据，就会打印 "received one"；如果接收到 c2 的数据，就会打印 "received two"。
// 以下实例中，我们定义了两个通道，并启动了两个协程（Goroutine）从这两个通道中获取数据。在 main 函数中，我们使用 select 语句在这两个通道中进行非阻塞的选择，如果两个通道都没有可用的数据，就执行 default 子句中的语句。
//
// 以下实例执行后会不断地从两个通道中获取到的数据，当两个通道都没有可用的数据时，会输出 "no message received"。
//func main() {
//	// 定义两个通道
//	ch1 := make(chan string)
//	ch2 := make(chan string)
//
//	// 启动两个 goroutine，分别从两个通道中获取数据
//	go func() {
//		for {
//			ch1 <- "from 1"
//		}
//	}()
//	go func() {
//		for {
//			ch2 <- "from 2"
//		}
//	}()
//
//	// 使用 select 语句非阻塞地从两个通道中获取数据
//	for {
//		select {
//		case msg1 := <-ch1:
//			fmt.Println(msg1)
//		case msg2 := <-ch2:
//			fmt.Println(msg2)
//		default:
//			// 如果两个通道都没有可用的数据，则执行这里的语句
//			fmt.Println("no message received")
//		}
//	}
//}
