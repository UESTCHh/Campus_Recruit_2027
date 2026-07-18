package main

import "fmt"

// Go 语言Map(集合)
// Map 是一种无序的键值对的集合。
//
// Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
//
// Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，遍历 Map 时返回的键值对的顺序是不确定的。
//
// 在获取 Map 的值时，如果键不存在，返回该类型的零值，例如 int 类型的零值是 0，string 类型的零值是 ""。
//
// Map 是引用类型，如果将一个 Map 传递给一个函数或赋值给另一个变量，它们都指向同一个底层数据结构，因此对 Map 的修改会影响到所有引用它的变量。
//
// 定义 Map
// 可以使用内建函数 make 或使用 map 关键字来定义 Map:
//
// /* 使用 make 函数 */
// map_variable := make(map[KeyType]ValueType, initialCapacity)
// 其中 KeyType 是键的类型，ValueType 是值的类型，initialCapacity 是可选的参数，用于指定 Map 的初始容量。Map 的容量是指 Map 中可以保存的键值对的数量，当 Map 中的键值对数量达到容量时，Map 会自动扩容。如果不指定 initialCapacity，Go 语言会根据实际情况选择一个合适的值。
//
// 实例
// // 创建一个空的 Map
// m := make(map[string]int)
//
// // 创建一个初始容量为 10 的 Map
// m := make(map[string]int, 10)
// 也可以使用字面量创建 Map：
//
// // 使用字面量创建 Map
//
//	m := map[string]int{
//	   "apple": 1,
//	   "banana": 2,
//	   "orange": 3,
//	}
//
// 获取元素：
//
// // 获取键值对
// v1 := m["apple"]
// v2, ok := m["pear"]  // 如果键不存在，ok 的值为 false，v2 的值为该类型的零值
// 修改元素：
//
// // 修改键值对
// m["apple"] = 5
// 获取 Map 的长度：
//
// // 获取 Map 的长度
// len := len(m)
// 遍历 Map：
//
// // 遍历 Map
//
//	for k, v := range m {
//	   fmt.Printf("key=%s, value=%d\n", k, v)
//	}
//
// 删除元素：
//
// // 删除键值对
// delete(m, "banana")

// 实例
// 下面实例演示了创建和使用map:
//
//	func main() {
//		//var siteMap map[string]string /*创建集合 */
//		siteMap := make(map[string]string)
//
//		/* map 插入 key - value 对,各个国家对应的首都 */
//		siteMap["Google"] = "谷歌"
//		siteMap["Runoob"] = "菜鸟教程"
//		siteMap["Baidu"] = "百度"
//		siteMap["Wiki"] = "维基百科"
//
//		/*使用键输出地图值 */
//		for site := range siteMap {
//			fmt.Println(site, "是", siteMap[site])
//		}
//
//		/*查看元素在集合中是否存在 */
//		name, ok := siteMap["Facebook"] /*如果确定是真实的,则存在,否则不存在 */
//		/*fmt.Println(capital) */
//		/*fmt.Println(ok) */
//		if ok {
//			fmt.Println("Facebook 的 站点是", name)
//		} else {
//			fmt.Println("Facebook 站点不存在")
//		}
//	}
//
// 以上实例运行结果为：
//
// Wiki 首都是 维基百科
// Google 首都是 谷歌
// Runoob 首都是 菜鸟教程
// Baidu 首都是 百度
// Facebook 站点不存在

//delete() 函数
//delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key。实例如下：
//
//实例
//package main
//
//import "fmt"
//
//func main() {
//        /* 创建map */
//        countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
//
//        fmt.Println("原始地图")
//
//        /* 打印地图 */
//        for country := range countryCapitalMap {
//                fmt.Println(country, "首都是", countryCapitalMap [ country ])
//        }
//
//        /*删除元素*/ delete(countryCapitalMap, "France")
//        fmt.Println("法国条目被删除")
//
//        fmt.Println("删除元素后地图")
//
//        /*打印地图*/
//        for country := range countryCapitalMap {
//                fmt.Println(country, "首都是", countryCapitalMap [ country ])
//        }
//}
//以上实例运行结果为：
//
//原始地图
//India 首都是 New delhi
//France 首都是 Paris
//Italy 首都是 Rome
//Japan 首都是 Tokyo
//法国条目被删除
//删除元素后地图
//Italy 首都是 Rome
//Japan 首都是 Tokyo
//India 首都是 New delhi

// 基于 go 实现简单 HashMap，暂未做 key 值的校验。
// 采用拉链法解决哈希冲突，使用固定长度 16 的桶数组。
// 每个桶内按 nodeIndex (hashCode >> 4) 升序维护链表，以优化查找。

// HashMap 定义链表节点结构
// 每个节点存储一个键值对，以及键的哈希码和指向下一个节点的指针
type HashMap struct {
	key      string   // 键
	value    string   // 值
	hashCode int      // 键的哈希码（由 genHashCode 计算）
	next     *HashMap // 链表后继指针
}

// table 是哈希表的桶数组，固定容量为 16，每个元素是指向 HashMap 节点的指针
var table [16]*HashMap

// initTable 初始化桶数组，为每个桶预置一个空节点（哨兵节点）
// 哨兵节点的 key 为空字符串，value 为空，hashCode 为索引值，next 为 nil
// 这样设计便于后续操作，避免空指针判断
func initTable() {
	for i := range table {
		table[i] = &HashMap{"", "", i, nil}
	}
}

// getInstance 返回桶数组，并确保已初始化（懒加载）
// 检查 table[0] 是否为 nil 来判断是否初始化，若未初始化则调用 initTable
func getInstance() [16](*HashMap) {
	if table[0] == nil {
		initTable()
	}
	return table
}

// genHashCode 生成字符串的哈希码
// 算法：从第一个字符开始，若未到最后一个字符，则 hashCode = (hashCode + 字符值) * 31
// 最后一个字符只加不乘，以保持一定分布性
// 注意：空字符串返回 0
func genHashCode(k string) int {
	if len(k) == 0 {
		return 0
	}
	var hashCode int = 0
	var lastIndex int = len(k) - 1
	for i := range k {
		if i == lastIndex {
			hashCode += int(k[i])
			break
		}
		hashCode += (hashCode + int(k[i])) * 31
	}
	return hashCode
}

// indexTable 根据哈希码计算桶索引（0 ~ 15）
// 使用取模运算，保证落在桶数组范围内
func indexTable(hashCode int) int {
	return hashCode % 16
}

// indexNode 计算节点的排序索引，用于在桶内链表中保持有序
// 取 hashCode 右移 4 位，相当于除以 16，即“桶内层级”
// 这样同一桶内的节点按此值升序排列，可加速查找（理论上）
func indexNode(hashCode int) int {
	return hashCode >> 4
}

// put 插入或更新键值对
// 若 key 已存在，则更新 value 并返回旧值；否则插入新节点，返回空字符串 ""
// 插入策略：
//  1. 计算哈希码、桶索引和节点排序索引
//  2. 若桶的头节点是空哨兵（key == ""），则直接替换为新的节点
//  3. 否则，遍历桶内链表，按 nodeIndex 升序找到插入位置
//  4. 若找到相同 hashCode 的节点（认为 key 相同），则更新 value
//  5. 否则插入到适当位置（保持链表有序）
func put(k string, v string) string {
	var hashCode = genHashCode(k)               // 计算哈希码
	var thisNode = HashMap{k, v, hashCode, nil} // 待插入的节点

	var tableIndex = indexTable(hashCode) // 桶索引
	var nodeIndex = indexNode(hashCode)   // 排序索引

	var headPtr [16](*HashMap) = getInstance() // 获取桶数组
	var headNode = headPtr[tableIndex]         // 当前桶的头节点（哨兵）

	// 若头节点是空的哨兵（key == ""），直接替换为新的节点
	if (*headNode).key == "" {
		*headNode = thisNode
		return ""
	}

	// 否则遍历链表，寻找插入位置（按 nodeIndex 升序）
	var lastNode *HashMap = headNode         // 前驱节点
	var nextNode *HashMap = (*headNode).next // 后继节点

	// 移动 lastNode 和 nextNode，直到 nextNode 为 nil 或 nextNode 的 nodeIndex 不小于当前 nodeIndex
	for nextNode != nil && (indexNode((*nextNode).hashCode) < nodeIndex) {
		lastNode = nextNode
		nextNode = (*nextNode).next
	}

	// 若找到相同哈希码的节点（认为 key 相同），则更新 value
	if (*lastNode).hashCode == thisNode.hashCode {
		var oldValue string = lastNode.value
		lastNode.value = thisNode.value
		return oldValue // 返回旧值
	}

	// 若 lastNode 的哈希码小于当前节点，则插入到 lastNode 之后
	if lastNode.hashCode < thisNode.hashCode {
		lastNode.next = &thisNode
	}
	// 若有后继节点，则将当前节点的 next 指向后继，形成链接
	if nextNode != nil {
		thisNode.next = nextNode
	}
	return "" // 插入成功，无旧值返回
}

// get 根据键查找对应的值
// 若找到则返回该值，否则返回空字符串 ""（注意：空字符串也可能作为有效值，此处未区分）
func get(k string) string {
	var hashCode = genHashCode(k)         // 计算哈希码
	var tableIndex = indexTable(hashCode) // 桶索引

	var headPtr [16](*HashMap) = getInstance() // 获取桶数组
	var node *HashMap = headPtr[tableIndex]    // 桶头节点

	// 检查头节点是否匹配
	if (*node).key == k {
		return (*node).value
	}

	// 遍历链表，查找匹配的 key
	for (*node).next != nil {
		if k == (*node).key {
			return (*node).value
		}
		node = (*node).next
	}
	return "" // 未找到返回空串
}

// 示例用法
func main() {
	getInstance()     // 初始化哈希表
	put("a", "a_put") // 插入键值对
	put("b", "b_put")
	fmt.Println(get("a")) // 输出: a_put
	fmt.Println(get("b")) // 输出: b_put
	put("p", "p_put")
	fmt.Println(get("p")) // 输出: p_put
}
