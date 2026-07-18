package main

//Go 语言切片(Slice)
//Go 语言切片是对数组的抽象。
//
//Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
//
//定义切片
//你可以声明一个未指定大小的数组来定义切片：
//
//var identifier []type
//切片不需要说明长度。
//
//或使用 make() 函数来创建切片:
//
//var slice1 []type = make([]type, len)
//
//也可以简写为
//
//slice1 := make([]type, len)
//也可以指定容量，其中 capacity 为可选参数。
//
//make([]T, length, capacity)
//这里 len 是数组的长度并且也是切片的初始长度。
//
//切片初始化
//s :=[] int {1,2,3 }
//直接初始化切片，[] 表示是切片类型，{1,2,3} 初始化值依次是 1,2,3，其 cap=len=3。
//
//s := arr[:]
//初始化切片 s，是数组 arr 的引用。
//
//s := arr[startIndex:endIndex]
//将 arr 中从下标 startIndex 到 endIndex-1 下的元素创建为一个新的切片。
//
//s := arr[startIndex:]
//默认 endIndex 时将表示一直到arr的最后一个元素。
//
//s := arr[:endIndex]
//默认 startIndex 时将表示从 arr 的第一个元素开始。
//
//s1 := s[startIndex:endIndex]
//通过切片 s 初始化切片 s1。
//
//s :=make([]int,len,cap)
//通过内置函数 make() 初始化切片s，[]int 标识为其元素类型为 int 的切片。

//len() 和 cap() 函数
//切片是可索引的，并且可以由 len() 方法获取长度。
//
//切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。
//
//以下为具体实例：
//
//实例
//package main
//
//import "fmt"
//
//func main() {
//   var numbers = make([]int,3,5)
//
//   printSlice(numbers)
//}
//
//func printSlice(x []int){
//   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
//}
//以上实例运行输出结果为:
//
//len=3 cap=5 slice=[0 0 0]

//空(nil)切片
//一个切片在未初始化之前默认为 nil，长度为 0，实例如下：
//
//实例
//package main
//
//import "fmt"
//
//func main() {
//   var numbers []int
//
//   printSlice(numbers)
//
//   if(numbers == nil){
//      fmt.Printf("切片是空的")
//   }
//}
//
//func printSlice(x []int){
//   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
//}
//以上实例运行输出结果为:
//
//len=0 cap=0 slice=[]
//切片是空的

//切片截取
//可以通过设置下限及上限来设置截取切片 [lower-bound:upper-bound]，实例如下：
//
//实例
//package main
//
//import "fmt"
//
//func main() {
//   /* 创建切片 */
//   numbers := []int{0,1,2,3,4,5,6,7,8}
//   printSlice(numbers)
//
//   /* 打印原始切片 */
//   fmt.Println("numbers ==", numbers)
//
//   /* 打印子切片从索引1(包含) 到索引4(不包含)*/
//   fmt.Println("numbers[1:4] ==", numbers[1:4])
//
//   /* 默认下限为 0*/
//   fmt.Println("numbers[:3] ==", numbers[:3])
//
//   /* 默认上限为 len(s)*/
//   fmt.Println("numbers[4:] ==", numbers[4:])
//
//   numbers1 := make([]int,0,5)
//   printSlice(numbers1)
//
//   /* 打印子切片从索引  0(包含) 到索引 2(不包含) */
//   number2 := numbers[:2]
//   printSlice(number2)
//
//   /* 打印子切片从索引 2(包含) 到索引 5(不包含) */
//   number3 := numbers[2:5]
//   printSlice(number3)
//
//}
//
//func printSlice(x []int){
//   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
//}
//执行以上代码输出结果为：
//
//len=9 cap=9 slice=[0 1 2 3 4 5 6 7 8]
//numbers == [0 1 2 3 4 5 6 7 8]
//numbers[1:4] == [1 2 3]
//numbers[:3] == [0 1 2]
//numbers[4:] == [4 5 6 7 8]
//len=0 cap=5 slice=[]
//len=2 cap=9 slice=[0 1]
//len=3 cap=7 slice=[2 3 4]
//你困惑的点在于 **`cap`（容量）** 的计算，特别是为什么 `number2` 的容量是 9，而 `number3` 的容量是 7。
//
//根本原因在于：**切片截取后的新切片，与原切片共享同一个底层数组。新切片的容量（cap），指的是从新切片的起始索引（`low`）到原底层数组末尾的元素个数。**
//
//计算公式非常简单：
//**新切片的 `cap` = 原切片的 `cap` - `low`**（下限索引）
//**新切片的 `len` = `high` - `low`**（上限减下限）
//
//---
//
//我们来逐一拆解你的疑问：
//
//### 1. `len=0 cap=5 slice=[]` （这是 `numbers1`）
//这是你在代码中单独使用 `make` 创建的：
//`numbers1 := make([]int, 0, 5)`
//它分配了一个**全新的底层数组**，长度是 0，容量是 5。它和上面的 `numbers` 数组没有任何关系，所以输出就是 `len=0 cap=5`。
//
//---
//
//### 2. `len=2 cap=9 slice=[0 1]` （这是 `number2`）
//代码：`number2 := numbers[:2]`
//
//- **下限 `low`**：默认是 `0`
//- **上限 `high`**：`2`
//- **长度**：`high - low` = `2 - 0` = **2**（即元素 [0, 1]）✅
//- **容量**：因为新切片从底层数组的**索引 0** 开始看，一直到原数组末尾（索引 8），一共还有 **9** 个元素。所以 `cap = 原cap(9) - low(0)` = **9**。
//
//---
//
//### 3. `len=3 cap=7 slice=[2 3 4]` （这是 `number3`）
//代码：`number3 := numbers[2:5]`
//
//- **下限 `low`**：`2`
//- **上限 `high`**：`5`
//- **长度**：`high - low` = `5 - 2` = **3**（即元素 [2, 3, 4]）✅
//- **容量**：新切片从底层数组的**索引 2** 开始看，一直到原数组末尾（索引 8），元素个数是 `8 - 2 + 1 = 7` 个。所以 `cap = 原cap(9) - low(2)` = **7**。
//
//---
//
//### 用一张图来理解底层数组
//
//原数组 `numbers` 在内存中长这样（下标 0 到 8）：
//
//```
//下标:  0    1    2    3    4    5    6    7    8
//值:  [ 0  | 1  | 2  | 3  | 4  | 5  | 6  | 7  | 8 ]
//      ▲                        ▲
//      |                        |
//   number2 从这里看           number3 从这里看
//   (low=0)                   (low=2)
//   到末尾有 9 个元素          到末尾有 7 个元素
//```
//
//- `number2` 站在下标 `0` 的位置看后面，能看到 `0~8` 共 **9** 个位置。
//- `number3` 站在下标 `2` 的位置看后面，能看到 `2~8` 共 **7** 个位置。
//
//---
//
//### 总结关键点（避免误区）
//
//- **`len`（长度）**：是你**截取的那一段**有多长（`high - low`）。
//- **`cap`（容量）**：是截取起点 **`low` 到原数组末尾**还有多长（`原cap - low`）。
//- `cap` 并不是 `high - low`，除非你刚好截取到末尾（比如 `[4:]` 时，cap 才等于 len）。
//
//所以，`number2` 的 `cap=9` 意味着虽然你现在只能用前 2 个位置，但底层数组还给你预留了后面 7 个位置（下标 2~8），如果对 `number2` 进行 `append` 操作，只要不超过容量 9，就不会触发扩容，且会覆盖原 `numbers` 后续的元素。

//append() 和 copy() 函数
//如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
//
//下面的代码描述了从拷贝切片的 copy 方法和向切片追加新元素的 append 方法。
//
//实例
//package main
//
//import "fmt"
//
//func main() {
//   var numbers []int
//   printSlice(numbers)
//
//   /* 允许追加空切片 */
//   numbers = append(numbers, 0)
//   printSlice(numbers)
//
//   /* 向切片添加一个元素 */
//   numbers = append(numbers, 1)
//   printSlice(numbers)
//
//   /* 同时添加多个元素 */
//   numbers = append(numbers, 2,3,4)
//   printSlice(numbers)
//
//   /* 创建切片 numbers1 是之前切片的两倍容量*/
//   numbers1 := make([]int, len(numbers), (cap(numbers))*2)
//
//   /* 拷贝 numbers 的内容到 numbers1 */
//   copy(numbers1,numbers)
//   printSlice(numbers1)
//}
//
//func printSlice(x []int){
//   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
//}
//以上代码执行输出结果为：
//
//len=0 cap=0 slice=[]
//len=1 cap=1 slice=[0]
//len=2 cap=2 slice=[0 1]
//len=5 cap=6 slice=[0 1 2 3 4]
//len=5 cap=12 slice=[0 1 2 3 4]
