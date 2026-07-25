package main

//Go 语言循环语句
//在不少实际问题中有许多具有规律性的重复操作，因此在程序中就需要重复执行某些语句。
//Go 语言提供了以下几种类型循环处理语句：
//
//循环类型	描述
//for 循环	重复执行语句块
//循环嵌套	在 for 循环中嵌套一个或多个 for 循环

//Go 语言 for 循环
//Go 语言循环语句Go 语言循环语句
//
//for 循环是一个循环控制结构，可以执行指定次数的循环。
//
//语法
//Go 语言的 For 循环有 3 种形式，只有其中的一种使用分号。
//
//和 C 语言的 for 一样：
//
//for init; condition; post { }
//和 C 的 while 一样：
//
//for condition { }
//和 C 的 for(;;) 一样：
//
//for { }
//init： 一般为赋值表达式，给控制变量赋初值；
//condition： 关系表达式或逻辑表达式，循环控制条件；
//post： 一般为赋值表达式，给控制变量增量或减量。
//for语句执行过程如下：
//
//1、先对表达式 1 赋初值；
//
//2、判别赋值表达式 init 是否满足给定条件，若其值为真，满足循环条件，则执行循环体内语句，然后执行 post，进入第二次循环，再判别 condition；否则判断 condition 的值为假，不满足条件，就终止for循环，执行循环体外语句。
//
//for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：
//
//for key, value := range oldMap {
//    newMap[key] = value
//}
//以上代码中的 key 和 value 是可以省略。
//
//如果只想读取 key，格式如下：
//
//for key := range oldMap
//或者这样：
//
//for key, _ := range oldMap
//
//如果只想读取 value，格式如下：
//
//for _, value := range oldMap

// 实例
// 计算 1 到 10 的数字之和：
//
//	func main() {
//		sum := 0
//		for i := 0; i <= 10; i++ {
//			sum += i
//		}
//		fmt.Println(sum)
//	}
//
// 输出结果为：
//
// 55
// init 和 post 参数是可选的，我们可以直接省略它，类似 While 语句。
//
// 以下实例在 sum 小于 10 的时候计算 sum 自相加后的值：
//
//	func main() {
//		sum := 1
//		for sum <= 10 {
//			sum += sum
//		}
//		fmt.Println(sum)
//
//		// 这样写也可以，更像 While 语句形式
//		for sum <= 10 {
//			sum += sum
//		}
//		fmt.Println(sum)
//	}
//
// 输出结果为：
//
// 16
// 16
// 无限循环:
//
//	func main() {
//		sum := 0
//		for {
//			sum++ // 无限循环下去
//		}
//		fmt.Println(sum) // 无法输出
//	}
//
// 要停止无限循环，可以在命令窗口按下ctrl-c 。
//
// # For-each range 循环
//
// 这种格式的循环可以对字符串、数组、切片等进行迭代输出元素。
//func main() {
//	strings := []string{"google", "runoob"}
//	for i, s := range strings {
//		fmt.Println(i, s)
//	}
//
//	numbers := [6]int{1, 2, 3, 5, 6} // 自动补0
//	for i, x := range numbers {
//		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
//	}
//}
//以上实例运行输出结果为:
//0 google
//1 runoob
//第 0 位 x 的值 = 1
//第 1 位 x 的值 = 2
//第 2 位 x 的值 = 3
//第 3 位 x 的值 = 5
//第 4 位 x 的值 = 6
//第 5 位 x 的值 = 0

// for 循环的 range 格式可以省略 key 和 value，如下实例：
//func main() {
//	map1 := make(map[int]float32)
//	map1[1] = 1.0
//	map1[2] = 2.0
//	map1[3] = 3.0
//	map1[4] = 4.0
//
//	// 读取 key 和 value
//	for key, value := range map1 {
//		fmt.Printf("key is: %d - value is: %f\n", key, value)
//	}
//
//	// 读取 key
//	for key := range map1 {
//		fmt.Printf("key is: %d\n", key)
//	}
//
//	// 读取 value
//	for _, value := range map1 {
//		fmt.Printf("value is: %f\n", value)
//	}
//}
//以上实例运行输出结果为:
//
//key is: 4 - value is: 4.000000
//key is: 1 - value is: 1.000000
//key is: 2 - value is: 2.000000
//key is: 3 - value is: 3.000000
//key is: 1
//key is: 2
//key is: 3
//key is: 4
//value is: 1.000000
//value is: 2.000000
//value is: 3.000000
//value is: 4.000000

// Go 语言循环嵌套
// Go 语言循环语句Go 语言循环语句
//
// Go 语言允许用户在循环内使用循环。接下来我们将为大家介绍嵌套循环的使用。
//
// 语法
// 以下为 Go 语言嵌套循环的格式：
//
// for [condition |  ( init; condition; increment ) | Range]
//
//	{
//	  for [condition |  ( init; condition; increment ) | Range]
//	  {
//	     statement(s);
//	  }
//	  statement(s);
//	}
//
// 实例
// 以下实例使用循环嵌套来输出 2 到 100 间的素数：
//func main() {
//	/* 定义局部变量 */
//	var i, j int
//
//	for i = 2; i < 100; i++ {
//		for j = 2; j <= (i / j); j++ {
//			if i%j == 0 {
//				break // 如果发现因子，则不是素数
//			}
//		}
//		if j > (i / j) {
//			fmt.Printf("%d  是素数\n", i)
//		}
//	}
//}
// 可读性更好的版本：
//for i := 2; i < 100; i++ {
//    isPrime := true
//    for j := 2; j*j <= i; j++ { // 直接使用 j*j <= i
//        if i%j == 0 {
//            isPrime = false
//            break
//        }
//    }
//    if isPrime {
//        fmt.Printf("%d  是素数\n", i)
//    }
//}
//以上实例运行输出结果为:
//
//2  是素数
//3  是素数
//5  是素数
//7  是素数
//11  是素数
//13  是素数
//17  是素数
//19  是素数
//23  是素数
//29  是素数
//31  是素数
//37  是素数
//41  是素数
//43  是素数
//47  是素数
//53  是素数
//59  是素数
//61  是素数
//67  是素数
//71  是素数
//73  是素数
//79  是素数
//83  是素数
//89  是素数
//97  是素数

//循环控制语句
//循环控制语句可以控制循环体内语句的执行过程。
//
//GO 语言支持以下几种循环控制语句：
//
//控制语句	描述
//break 语句	经常用于中断当前 for 循环或跳出 switch 语句
//continue 语句	跳过当前循环的剩余语句，然后继续进行下一轮循环。
//goto 语句	将控制转移到被标记的语句。

// Go 语言 break 语句
// Go 语言循环语句 Go语言循环语句
//
// 在 Go 语言中，break 语句用于终止当前循环或者 switch 语句的执行，并跳出该循环或者 switch 语句的代码块。
//
// break 语句可以用于以下几个方面：。
//
// 用于循环语句中跳出循环，并开始执行循环之后的语句。
// break 在 switch 语句中在执行一条 case 后跳出语句的作用。
// break 可应用在 select 语句中。
// 在多重循环中，可以用标号 label 标出想 break 的循环。
// 语法
// break 语法格式如下：
//
// break
// 实例
// 在 for 循环中使用 break：
//
//	func main() {
//		for i := 0; i < 10; i++ {
//			if i == 5 {
//				break // 当 i 等于 5 时跳出循环
//			}
//			fmt.Println(i)
//		}
//	}
//
// 输出结果：
//
// 0
// 1
// 2
// 3
// 4
// 在变量 a 大于 15 的时候跳出循环：
//
//	func main() {
//		/* 定义局部变量 */
//		var a int = 10
//
//		/* for 循环 */
//		for a < 20 {
//			fmt.Printf("a 的值为 : %d\n", a)
//			a++
//			if a > 15 {
//				/* a 大于 15 时使用 break 语句跳出循环 */
//				break
//			}
//		}
//	}
//
// 以上实例中当 a 的值大于 15 时，break 语句被执行，整个循环被终止。执行结果为：
//
// a 的值为 : 10
// a 的值为 : 11
// a 的值为 : 12
// a 的值为 : 13
// a 的值为 : 14
// a 的值为 : 15
// 以下实例有多重循环，演示了使用标记和不使用标记的区别：
// func main() {
//
//	// 不使用标记
//	fmt.Println("---- break ----")
//	for i := 1; i <= 3; i++ {
//		fmt.Printf("i: %d\n", i)
//		for i2 := 11; i2 <= 13; i2++ {
//			fmt.Printf("i2: %d\n", i2)
//			break
//		}
//	}
//
//	// 使用标记
//	fmt.Println("---- break label ----")
//
// re:
//
//		for i := 1; i <= 3; i++ {
//			fmt.Printf("i: %d\n", i)
//			for i2 := 11; i2 <= 13; i2++ {
//				fmt.Printf("i2: %d\n", i2)
//				break re
//			}
//		}
//	}
//
// 以上实例执行结果为：
//
// ---- break ----
// i: 1
// i2: 11
// i: 2
// i2: 11
// i: 3
// i2: 11
// ---- break label ----
// i: 1
// i2: 11
// 在 switch 语句中使用 break：
//
//	func main() {
//		day := "Tuesday"
//		switch day {
//		case "Monday":
//			fmt.Println("It's Monday.")
//		case "Tuesday":
//			fmt.Println("It's Tuesday.")
//			break // 跳出 switch 语句
//		case "Wednesday":
//			fmt.Println("It's Wednesday.")
//		}
//	}
//
// 输出结果：
//
// It's Tuesday.
// 在 select 语句中使用 break：
//
//	func main() {
//		ch1 := make(chan int)
//		ch2 := make(chan int)
//
//		go func() {
//			time.Sleep(2 * time.Second)
//			ch1 <- 1
//		}()
//
//		go func() {
//			time.Sleep(1 * time.Second)
//			ch2 <- 2
//		}()
//
//		select {
//		case <-ch1:
//			fmt.Println("Received from ch1.")
//		case <-ch2:
//			fmt.Println("Received from ch2.")
//			break // 跳出 select 语句
//		}
//	}
//
// 输出结果：
//
// Received from ch2.
// 在 Go 语言中，break 语句在 select 语句中的应用是相对特殊的。由于 select 语句的特性，break 语句并不能直接用于跳出 select 语句本身，因为 select 语句是非阻塞的，它会一直等待所有的通信操作都准备就绪。如果需要提前结束 select 语句的执行，可以使用 return 或者 goto 语句来达到相同的效果。
// 以下实例，展示了在 select 语句中使用 return 来提前结束执行的情况：
//func process(ch chan int) {
//	for {
//		select {
//		case val := <-ch:
//			fmt.Println("Received value:", val)
//			// 执行一些逻辑
//			if val == 5 {
//				return // 提前结束 select 语句的执行
//			}
//		default:
//			fmt.Println("No value received yet.")
//			time.Sleep(500 * time.Millisecond)
//		}
//	}
//}
//
//func main() {
//	ch := make(chan int)
//
//	go process(ch)
//
//	time.Sleep(2 * time.Second)
//	ch <- 1
//	time.Sleep(1 * time.Second)
//	ch <- 3
//	time.Sleep(1 * time.Second)
//	ch <- 5
//	time.Sleep(1 * time.Second)
//	ch <- 7
//
//	time.Sleep(2 * time.Second)
//}
//以上实例中，process 函数在一个无限循环中使用 select 语句等待通道 ch 上的数据。当接收到数据时，会执行一些逻辑。当接收到的值等于 5 时，使用 return 提前结束 select 语句的执行。
//
//输出结果：
//
//No value received yet.
//No value received yet.
//Received value: 1
//No value received yet.
//Received value: 3
//No value received yet.
//Received value: 5
//通过使用 return，我们可以在 select 语句中提前终止执行，并返回到调用者的代码中。
//
//需要注意的是，使用 return 语句会立即终止当前的函数执行，所以请根据实际需求来决定在 select 语句中使用何种方式来提前结束执行。

// 无限循环
// 如果循环中条件语句永远不为 false 则会进行无限循环，我们可以通过 for 循环语句中只设置一个条件表达式来执行无限循环：
//func main() {
//	for true { // 也可以啥也不写， for {}
//		fmt.Printf("这是无限循环。\n")
//	}
//}
