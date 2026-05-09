# 142. 环形链表 II 学习笔记

## 一、 题目分析

### 1.1 题目描述

给定一个链表的头节点 `head`，返回链表开始入环的第一个节点。如果链表无环，则返回 `null`。

如果链表中有某个节点，可以通过连续跟踪 `next` 指针再次到达，则链表中存在环。为了表示给定链表中的环，评测系统内部使用整数 `pos` 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 `pos` 是 -1，则在该链表中没有环。注意：`pos` 不作为参数进行传递，仅仅是为了标识链表的实际情况。

不允许修改链表。

### 1.2 题目链接

- [LeetCode 142. 环形链表 II](https://leetcode.cn/problems/linked-list-cycle-ii/description/)
- [灵茶山艾府题解](https://leetcode.cn/problems/linked-list-cycle-ii/solutions/1999271/mei-xiang-ming-bai-yi-ge-shi-pin-jiang-t-nvsq/)
- [官方题解](https://leetcode.cn/problems/linked-list-cycle-ii/solutions/441131/huan-xing-lian-biao-ii-by-leetcode-solution/)
- [宫水三叶题解](https://leetcode.cn/problems/linked-list-cycle-ii/solutions/1974627/huan-xing-lian-biao-de-jie-dian-ding-wei-a68a/)
- [题解：快慢指针双指针](https://leetcode.cn/problems/linked-list-cycle-ii/solutions/12616/linked-list-cycle-ii-kuai-man-zhi-zhen-shuang-zhi-/)

### 1.3 示例

**示例 1：**

```
输入：head = [3,2,0,-4], pos = 1
输出：返回索引为 1 的链表节点
解释：链表中有一个环，其尾部连接到第二个节点。
```

**示例 2：**

```
输入：head = [1,2], pos = 0
输出：返回索引为 0 的链表节点
解释：链表中有一个环，其尾部连接到第一个节点。
```

**示例 3：**

```
输入：head = [1], pos = -1
输出：返回 null
解释：链表中没有环。
```

### 1.4 提示

- 链表中节点的数目范围在范围 [0, 10⁴] 内
- -10⁵ ≤ Node.val ≤ 10⁵
- pos 的值为 -1 或者链表中的一个有效索引

### 1.5 题目本质

环形链表 II 的本质是：**在判断链表是否有环的基础上，进一步找到环的入口节点**。

换句话说，我们需要：
1. 先判断链表是否有环（类似 141. 环形链表）
2. 如果有环，找到环的入口节点
3. 要求使用 O(1) 空间复杂度（进阶要求）

### 1.6 题目进阶

你是否可以使用 O(1) 空间解决此题？

---

## 二、 解题思路分析

### 2.1 核心思路

解决环形链表 II 问题的核心思路是：**使用快慢指针判断是否有环，然后利用数学推导找到入口节点**。

在解题过程中，我们需要注意以下几点：

| 关键点 | 说明 |
|-------|------|
| 快慢指针 | 慢指针走 1 步，快指针走 2 步 |
| 数学推导 | 相遇后，一个指针从头出发，一个从相遇点出发，同速前进 |
| 空间复杂度 | O(1) 空间，不需要额外数据结构 |
| 边界处理 | 空链表、单节点链表都要考虑 |

### 2.2 快慢指针原理选择

对于环形链表问题，我们选择**快慢指针法**而不是哈希表法的原因：

- 快慢指针法空间复杂度为 O(1)，符合进阶要求
- 虽然数学推导看起来复杂，但一旦理解，代码非常简洁
- 在面试中，面试官更希望看到 O(1) 空间的解法

### 2.3 快慢指针思路

#### 2.3.1 快慢指针三步曲

**1. 变量定义**
- `slow`：慢指针，每次走 1 步
- `fast`：快指针，每次走 2 步
- 两者都从头节点出发

**2. 阶段 1：判断是否有环并找到相遇点**

| 步骤 | 操作 | 目的 |
|------|------|------|
| 步骤 1 | `slow = slow-&gt;next` | 慢指针走 1 步 |
| 步骤 2 | `fast = fast-&gt;next-&gt;next` | 快指针走 2 步 |
| 步骤 3 | 检查 `slow == fast` | 判断是否相遇 |
| 步骤 4 | 如果相遇，进入阶段 2 | 准备找入口节点 |
| 步骤 5 | 如果 `fast` 或 `fast-&gt;next` 为 null | 说明无环，返回 null |

**3. 阶段 2：找到环的入口节点**
- `index1` 从头节点出发
- `index2` 从相遇点出发
- 两者每次都走 1 步
- 相遇点就是环的入口节点

#### 2.3.2 快慢指针图解

```
示例 1：head = [3,2,0,-4], pos = 1

链表结构：
3 -&gt; 2 -&gt; 0 -&gt; -4
          ^      |
          |______|

阶段 1：快慢指针移动过程

初始状态：
slow = 3, fast = 3

第 1 步：
slow = 2, fast = 0

第 2 步：
slow = 0, fast = 2

第 3 步：
slow = -4, fast = 0

第 4 步：
slow = 2, fast = 2 (相遇！)

阶段 2：找入口节点

index1 = 3, index2 = 2

第 1 步：
index1 = 2, index2 = 0

第 2 步：
index1 = 0, index2 = -4

第 3 步：
index1 = -4, index2 = 2

第 4 步：
index1 = 2, index2 = 2 (相遇！这就是入口节点！)
```

#### 2.3.3 极其严密的数学推导（白板必写）

这道题纯写代码没有意义，面试官要看的是你在白板上的推导过程。我们设三个未知的距离变量：
1. **x**：从链表头部 `head` 到**环入口节点**的距离
2. **y**：从环入口节点到快慢指针**相遇点**的距离
3. **z**：从相遇点继续走到**环入口节点**的距离（此时环的周长就是 `y + z`）

**【逻辑推演开始】**
- 当快慢指针相遇时：
  - 慢指针走过的距离：`x + y`
  - 快指针走过的距离：`x + y + n(y + z)`（其中 n 是快指针在环内绕的圈数，n ≥ 1）
- 因为快指针的速度是慢指针的 2 倍，所以相同时间内，快指针的路程是慢指针的 2 倍：
  - `2 × (x + y) = x + y + n(y + z)`
- 化简等式：
  - `x + y = n(y + z)`
  - **`x = n(y + z) - y`**
- 为了寻找规律，我们把 `n(y+z)` 拆出一个 `(y+z)` 来：
  - `x = (n - 1)(y + z) + y + z - y`
  - **`x = (n - 1)(y + z) + z`**

**【终极物理意义的顿悟】**
看上面这个公式！`x` 代表头节点到入口的距离，`z` 代表相遇点到入口的距离。`(n - 1)(y + z)` 只是在环里绕圈的无用功。
这意味着：**如果此时安排一个指针 `index1` 从头节点出发，另一个指针 `index2` 从相遇点出发，两人每次都只走 1 步。当 `index1` 走完 `x` 的距离时，`index2` 刚好在环里绕了 `n-1` 圈加上 `z` 的距离。两人一定会精确无误地在"环的入口节点"相遇！**

**【关于慢指针是否会绕圈的问题】**
有人会问：上面的公式推导中，为什么慢指针走过的路程肯定是 `x + y`，而不是慢指针也绕了环好几圈呢（比如 `x + m(y+z) + y`）？
绝杀回答：因为快指针的速度是慢指针的两倍。当慢指针刚刚走到环入口时，快指针一定已经在环内了。接下来，慢指针只需要走完不到一圈的距离，快指针就绝对能追上它。由于慢指针走完一圈的时间，足够快指针走两圈，所以快指针必将在慢指针走完第一圈之前将其截获。因此，相遇时慢指针走过的环内距离 `y` 一定小于环的周长，不可能绕圈。

### 2.4 哈希表思路

如果工程项目中不在乎空间复杂度，也可以使用哈希表：

**1. 哈希表思路**
- 遍历链表，把每个访问到的节点指针存入哈希表
- 每次插入前检查该指针是否已经存在于哈希表中
- 如果存在，该节点就是环的入口
- 如果遍历完都没有重复，说明无环

---

## 三、 算法实现

### 3.1 快慢指针法（推荐）

```cpp
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode *detectCycle(ListNode *head) {
        ListNode* slow = head;
        ListNode* fast = head;

        // 阶段 1：使用快慢指针判断是否有环，并找到相遇点
        while (fast != nullptr &amp;&amp; fast-&gt;next != nullptr) {
            slow = slow-&gt;next;
            fast = fast-&gt;next-&gt;next;

            // 发生相遇，说明有环！开始阶段 2！
            if (slow == fast) {
                ListNode* index1 = head; // 派遣指针 1 从头节点出发
                ListNode* index2 = fast; // 派遣指针 2 从相遇点出发（此时 fast 和 slow 在同一个位置）

                // 两人同速前进，按照公式推导，相遇点必定是环的入口！
                while (index1 != index2) {
                    index1 = index1-&gt;next;
                    index2 = index2-&gt;next;
                }
                
                return index1; // 返回环的入口节点
            }
        }

        return nullptr; // 跑完了都没相遇，说明无环
    }
};
```

### 3.2 哈希表法

```cpp
class Solution {
public:
    ListNode *detectCycle(ListNode *head) {
        unordered_set&lt;ListNode*&gt; visited;
        
        while (head != nullptr) {
            // 如果当前节点已经在哈希表中，说明是环的入口
            if (visited.count(head)) {
                return head;
            }
            visited.insert(head);
            head = head-&gt;next;
        }
        
        return nullptr; // 无环
    }
};
```

---

## 四、 代码分析

### 4.1 快慢指针代码分析

#### 4.1.1 实现思路说明

**快慢指针的设计思路：**
1. 快慢指针都从头节点出发
2. 慢指针每次走 1 步，快指针每次走 2 步
3. 如果有环，快慢指针一定会相遇
4. 相遇后，一个指针从头出发，一个从相遇点出发
5. 两者同速前进，相遇点就是环的入口

#### 4.1.2 核心逻辑讲解

**关键代码段 1：快慢指针初始化**
```cpp
ListNode* slow = head;
ListNode* fast = head;
```
- 快慢指针都从头节点出发

**关键代码段 2：阶段 1 - 找相遇点**
```cpp
while (fast != nullptr &amp;&amp; fast-&gt;next != nullptr) {
    slow = slow-&gt;next;
    fast = fast-&gt;next-&gt;next;
    if (slow == fast) {
        // 找到相遇点
    }
}
```
- 循环条件要检查 `fast` 和 `fast-&gt;next`，避免空指针
- 慢指针走 1 步，快指针走 2 步
- 如果相遇，说明有环

**关键代码段 3：阶段 2 - 找入口**
```cpp
ListNode* index1 = head;
ListNode* index2 = fast;
while (index1 != index2) {
    index1 = index1-&gt;next;
    index2 = index2-&gt;next;
}
return index1;
```
- 两个指针同速前进
- 相遇点就是环的入口

#### 4.1.3 初学者引导

**为什么快指针一定要走 2 步，不能走 3 步、4 步？**
- 走 2 步是最简单、最容易理解的
- 其实走 3 步、4 步也能相遇，但数学推导会更复杂
- 走 2 步可以保证快指针在慢指针走完第一圈前追上
- 面试中就用走 2 步的方案

**为什么要检查 `fast-&gt;next` 而不只是 `fast`？**
- 因为快指针要走 2 步：`fast-&gt;next-&gt;next`
- 如果 `fast-&gt;next` 是 null，访问 `fast-&gt;next-&gt;next` 就会空指针异常
- 所以必须同时检查 `fast` 和 `fast-&gt;next`

### 4.2 哈希表代码分析

#### 4.2.1 实现思路说明

**哈希表的设计思路：**
1. 遍历链表，记录每个访问过的节点
2. 如果某个节点已经被访问过，说明是环的入口
3. 如果遍历完都没有重复，说明无环

#### 4.2.2 核心逻辑讲解

**关键代码段 1：哈希表查找**
```cpp
if (visited.count(head)) {
    return head;
}
```
- 检查当前节点是否已经访问过
- 如果是，说明是环的入口

---

## 五、 时间与空间复杂度分析

### 5.1 快慢指针复杂度分析

| 复杂度类型 | 时间复杂度 | 空间复杂度 |
|-----------|-----------|-----------|
| **快慢指针法** | $O(n)$ | $O(1)$ |

**时间复杂度推导：**
- 阶段 1：快慢指针相遇，慢指针最多走 n 步
- 阶段 2：找入口，最多走 n 步
- 总共最多走 2n 步，时间复杂度 O(n)

**空间复杂度推导：**
- 只使用了几个指针变量
- 没有使用与数据规模相关的额外空间
- 空间复杂度为严格的 O(1)

### 5.2 哈希表复杂度分析

| 复杂度类型 | 时间复杂度 | 空间复杂度 |
|-----------|-----------|-----------|
| **哈希表法** | $O(n)$ | $O(n)$ |

**时间复杂度推导：**
- 遍历链表一次，O(n)
- 哈希表的插入和查找都是 O(1) 平均时间复杂度

**空间复杂度推导：**
- 需要存储所有访问过的节点
- 最坏情况下（无环），需要存储 n 个节点
- 空间复杂度 O(n)

### 5.3 复杂度对比表格

| 方法 | 时间复杂度 | 空间复杂度 | 优点 | 缺点 |
|------|-----------|-----------|------|------|
| 快慢指针法 | $O(n)$ | $O(1)$ | 空间复杂度最优，符合进阶要求 | 需要数学推导 |
| 哈希表法 | $O(n)$ | $O(n)$ | 思路简单，容易理解 | 需要额外空间 |

---

## 六、 测试用例验证

### 6.1 测试用例 1：正常有环情况

**输入：**
```
head = [3,2,0,-4], pos = 1
```

**执行过程（快慢指针法）：**
1. 快慢指针移动，最终在节点 2 相遇
2. index1 从头出发，index2 从相遇点出发
3. 两者同速前进，在节点 2 相遇

**输出：** 返回索引为 1 的节点（值为 2）

### 6.2 测试用例 2：环入口在头节点

**输入：**
```
head = [1,2], pos = 0
```

**执行过程：**
1. 快慢指针移动，会在环内相遇
2. index1 从头出发，index2 从相遇点出发
3. 两者会在头节点相遇

**输出：** 返回索引为 0 的节点（值为 1）

### 6.3 测试用例 3：无环

**输入：**
```
head = [1], pos = -1
```

**执行过程：**
1. 快指针很快走到 null
2. 没有相遇，返回 null

**输出：** null

### 6.4 测试用例 4：空链表

**输入：**
```
head = null
```

**执行过程：**
1. 直接返回 null

**输出：** null

---

## 七、 常见错误与注意事项

### 7.1 忘记检查 fast->next

**错误示例：**
```cpp
// 错误：只检查 fast，不检查 fast->next
while (fast != nullptr) {
    slow = slow->next;
    fast = fast->next->next; // 可能空指针异常
}
```

**后果：**
- 当 `fast` 指向最后一个节点时，`fast->next` 是 null
- 访问 `fast->next->next` 会导致空指针异常

**正确做法：**
```cpp
while (fast != nullptr && fast->next != nullptr) {
    // ...
}
```

### 7.2 快指针只走 1 步

**错误示例：**
```cpp
// 错误：快慢指针都走 1 步
slow = slow->next;
fast = fast->next;
```

**后果：**
- 两个指针永远同步，永远不会相遇
- 无法判断是否有环

**正确做法：**
```cpp
slow = slow->next;
fast = fast->next->next;
```

### 7.3 相遇后直接返回相遇点

**错误示例：**
```cpp
// 错误：相遇后直接返回
if (slow == fast) {
    return slow; // 这不是入口点！
}
```

**后果：**
- 返回的是相遇点，不是环的入口点
- 答案错误

**正确做法：**
```cpp
if (slow == fast) {
    ListNode* index1 = head;
    ListNode* index2 = fast;
    while (index1 != index2) {
        index1 = index1->next;
        index2 = index2->next;
    }
    return index1;
}
```

### 7.4 修改链表结构

**错误示例：**
```cpp
// 错误：修改了链表
ListNode* temp = head->next;
head->next = nullptr; // 题目要求不允许修改链表
```

**后果：**
- 违反题目要求
- 虽然可能通过测试，但不是正确解法

**正确做法：**
不修改链表，只进行遍历和指针操作。

---

## 八、 与其他相关题目的对比分析

### 8.1 与 141. 环形链表对比

| 题目 | 核心问题 | 难度 |
|------|---------|------|
| 141. 环形链表 | 判断链表是否有环 | 简单 |
| 142. 环形链表 II | 找到环的入口节点 | 中等 |

**相同点：**
- 都可以用快慢指针法
- 都是链表环的问题

**不同点：**
- 141 题只需要判断是否有环
- 142 题需要进一步找到入口节点
- 142 题需要数学推导

### 8.2 与 287. 寻找重复数对比

| 题目 | 数据结构 | 核心思想 |
|------|---------|---------|
| 142. 环形链表 II | 链表 | 快慢指针找环入口 |
| 287. 寻找重复数 | 数组 | 快慢指针找环入口（把数组看成链表） |

**相同点：**
- 都可以用快慢指针法
- 都是找"环入口"的问题

**不同点：**
- 142 题是真正的链表
- 287 题是把数组抽象成链表

---

## 九、 面试高频问题与解答

### 9.1 基础问题

#### Q1：快慢指针法的核心思想是什么？

**答案：**

**核心思想：**
1. 慢指针走 1 步，快指针走 2 步
2. 如果有环，快慢指针一定会相遇
3. 相遇后，一个指针从头出发，一个从相遇点出发，同速前进
4. 相遇点就是环的入口

**关键：**
- 数学推导是核心，一定要理解 `x = (n - 1)(y + z) + z`

#### Q2：为什么快慢指针一定会相遇？

**答案：**

**关键：** 相对速度！

- 快指针相对于慢指针的速度是 1 步/次
- 如果有环，快指针一定会追上慢指针
- 就像在环形跑道上跑步，快的人一定会追上慢的人

**类比：**
- 想象两个运动员在环形跑道上跑步
- 一个速度是 1，一个速度是 2
- 快的人一定会从后面追上慢的人

### 9.2 进阶问题

#### Q3：为什么慢指针在相遇时一定没有走完一圈？

**答案：**

**关键：** 时间和路程分析！

- 当慢指针走到环入口时，快指针已经在环内了
- 此时两者的距离最多是环长 - 1
- 快指针相对于慢指针的速度是 1
- 所以最多需要环长 - 1 步就能追上
- 而慢指针在这段时间内走的距离就是环长 - 1，小于环长
- 所以慢指针一定没有走完一圈

**绝杀回答：** 因为快指针的速度是慢指针的两倍。当慢指针刚刚走到环入口时，快指针一定已经在环内了。接下来，慢指针只需要走完不到一圈的距离，快指针就绝对能追上它。由于慢指针走完一圈的时间，足够快指针走两圈，所以快指针必将在慢指针走完第一圈之前将其截获。因此，相遇时慢指针走过的环内距离 y 一定小于环的周长，不可能绕圈。

#### Q4：快指针可以走 3 步吗？

**答案：**

**可以，但不推荐！**

- 走 3 步也能相遇，但数学推导更复杂
- 走 2 步是最简单、最容易理解的
- 面试中就用走 2 步的方案

**注意：**
- 如果走 3 步，快指针相对于慢指针的速度是 2
- 也能追上，但需要考虑更多边界情况

#### Q5：如何证明 x = z（当 n=1 时）？

**答案：**

**当 n=1 时（快指针只多走了一圈）：**
- 快指针路程：x + y + (y + z)
- 慢指针路程：x + y
- 2(x + y) = x + y + y + z
- 化简得：x = z

**这是最常见的情况，也是最容易理解的情况。**

### 9.3 实际问题

#### Q6：在实际工程中，什么场景会用到环形链表？

**答案：**

**常见场景：**
1. **循环缓冲区（Circular Buffer）**：用于数据流处理
2. **LRU Cache**：可以用环形链表实现
3. **定时器**：时间轮算法
4. **游戏开发**：循环动画、循环路径

**环形链表的优势：**
- 可以循环遍历
- 不需要特殊处理首尾边界

#### Q7：如何判断两个单链表是否相交？

**答案：**

**常见方法：**
1. **哈希表法**：遍历一个链表，记录节点，然后遍历另一个
2. **双指针法**：两个指针分别从两个链表头出发，走到末尾后换个头继续走，相遇点就是交点
3. **拼接法**：把第一个链表的尾接到第二个链表的头，然后用 142 题的方法找环入口

**这是一个经典的面试题，和 142 题有相似之处。**

---

## 十、 完整可运行代码

```cpp
#include <iostream>
#include <unordered_set>
using namespace std;

// 链表节点定义
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

// 方法一：快慢指针法（推荐）
class FastSlowPointerSolution {
public:
    ListNode *detectCycle(ListNode *head) {
        ListNode* slow = head;
        ListNode* fast = head;

        // 阶段 1：使用快慢指针判断是否有环，并找到相遇点
        while (fast != nullptr && fast->next != nullptr) {
            slow = slow->next;
            fast = fast->next->next;

            // 发生相遇，说明有环！开始阶段 2！
            if (slow == fast) {
                ListNode* index1 = head; // 派遣指针 1 从头节点出发
                ListNode* index2 = fast; // 派遣指针 2 从相遇点出发（此时 fast 和 slow 在同一个位置）

                // 两人同速前进，按照公式推导，相遇点必定是环的入口！
                while (index1 != index2) {
                    index1 = index1->next;
                    index2 = index2->next;
                }
                
                return index1; // 返回环的入口节点
            }
        }

        return nullptr; // 跑完了都没相遇，说明无环
    }
};

// 方法二：哈希表法
class HashTableSolution {
public:
    ListNode *detectCycle(ListNode *head) {
        unordered_set<ListNode*> visited;
        
        while (head != nullptr) {
            // 如果当前节点已经在哈希表中，说明是环的入口
            if (visited.count(head)) {
                return head;
            }
            visited.insert(head);
            head = head->next;
        }
        
        return nullptr; // 无环
    }
};

// 辅助函数：创建有环链表
ListNode* createCyclicList(const vector<int>& vals, int pos) {
    if (vals.empty()) return nullptr;
    
    ListNode* head = new ListNode(vals[0]);
    ListNode* curr = head;
    ListNode* cycleEntry = nullptr;
    
    for (int i = 1; i < vals.size(); i++) {
        curr->next = new ListNode(vals[i]);
        curr = curr->next;
        if (i == pos) {
            cycleEntry = curr;
        }
    }
    
    // 如果 pos >= 0，创建环
    if (pos >= 0) {
        if (pos == 0) {
            cycleEntry = head;
        }
        curr->next = cycleEntry;
    }
    
    return head;
}

// 辅助函数：打印链表信息
void printListInfo(ListNode* head, ListNode* entry) {
    if (head == nullptr) {
        cout << "空链表" << endl;
        return;
    }
    
    cout << "链表节点值：";
    ListNode* curr = head;
    unordered_set<ListNode*> visited;
    int count = 0;
    
    while (curr != nullptr && count < 10) { // 限制打印数量，避免无限循环
        cout << curr->val;
        if (visited.count(curr)) {
            cout << " (环入口)";
            break;
        }
        visited.insert(curr);
        if (curr->next != nullptr) cout << " -> ";
        curr = curr->next;
        count++;
    }
    cout << endl;
    
    if (entry != nullptr) {
        cout << "环入口节点值：" << entry->val << endl;
    } else {
        cout << "无环" << endl;
    }
}

// 主函数测试
int main() {
    cout << "========================================" << endl;
    cout << "     142. 环形链表 II 测试" << endl;
    cout << "========================================" << endl;

    // 创建解法实例
    FastSlowPointerSolution fastSlowSol;
    HashTableSolution hashTableSol;

    // 测试用例 1：正常有环情况
    cout << "\n--- 测试用例 1 ---" << endl;
    vector<int> vals1 = {3, 2, 0, -4};
    int pos1 = 1;
    ListNode* head1 = createCyclicList(vals1, pos1);
    cout << "输入：vals = [3,2,0,-4], pos = 1" << endl;
    ListNode* entry1_fast = fastSlowSol.detectCycle(head1);
    ListNode* entry1_hash = hashTableSol.detectCycle(head1);
    printListInfo(head1, entry1_fast);
    cout << "期望：环入口节点值为 2" << endl;

    // 测试用例 2：环入口在头节点
    cout << "\n--- 测试用例 2 ---" << endl;
    vector<int> vals2 = {1, 2};
    int pos2 = 0;
    ListNode* head2 = createCyclicList(vals2, pos2);
    cout << "输入：vals = [1,2], pos = 0" << endl;
    ListNode* entry2_fast = fastSlowSol.detectCycle(head2);
    ListNode* entry2_hash = hashTableSol.detectCycle(head2);
    printListInfo(head2, entry2_fast);
    cout << "期望：环入口节点值为 1" << endl;

    // 测试用例 3：无环
    cout << "\n--- 测试用例 3 ---" << endl;
    vector<int> vals3 = {1};
    int pos3 = -1;
    ListNode* head3 = createCyclicList(vals3, pos3);
    cout << "输入：vals = [1], pos = -1" << endl;
    ListNode* entry3_fast = fastSlowSol.detectCycle(head3);
    ListNode* entry3_hash = hashTableSol.detectCycle(head3);
    printListInfo(head3, entry3_fast);
    cout << "期望：无环" << endl;

    // 测试用例 4：空链表
    cout << "\n--- 测试用例 4 ---" << endl;
    ListNode* head4 = nullptr;
    cout << "输入：空链表" << endl;
    ListNode* entry4_fast = fastSlowSol.detectCycle(head4);
    ListNode* entry4_hash = hashTableSol.detectCycle(head4);
    printListInfo(head4, entry4_fast);
    cout << "期望：无环" << endl;

    cout << "\n========================================" << endl;
    cout << "           测试完成！" << endl;
    cout << "========================================" << endl;

    return 0;
}
```

---

## 十一、 总结与反思

### 11.1 核心知识点总结

| 知识点 | 要点 |
|--------|------|
| 快慢指针 | 慢指针走 1 步，快指针走 2 步 |
| 数学推导 | x = (n - 1)(y + z) + z |
| 找入口 | 一个从头出发，一个从相遇点出发，同速前进 |
| 时间复杂度 | O(n) |
| 空间复杂度 | 快慢指针 O(1)，哈希表 O(n) |

### 11.2 学习收获

1. **理解了快慢指针法的核心原理**
   - 为什么快慢指针一定会相遇
   - 数学推导的重要性
   - 如何从相遇点找到入口节点

2. **掌握了两种解题方法**
   - 快慢指针法：最优解，空间复杂度 O(1)
   - 哈希表法：思路简单，容易理解

3. **学会了如何进行数学推导**
   - 变量定义要清晰
   - 等式化简要仔细
   - 最终结论要有物理意义

4. **理解了环形链表问题的本质**
   - 相对速度的概念
   - 环形跑道的类比

### 11.3 后续学习建议

1. **练习更多链表题目**
   - 141. 环形链表
   - 206. 反转链表
   - 21. 合并两个有序链表
   - 23. 合并 K 个升序链表

2. **深入理解双指针技巧**
   - 快慢指针
   - 左右指针
   - 滑动窗口
   - 双指针在数组、链表中的应用

3. **学习其他找环入口的问题**
   - 287. 寻找重复数
   - 这道题可以用快慢指针法，把数组看成链表

4. **学习环形链表的变体**
   - 如何求环的长度
   - 如何判断两个环形链表是否相交
   - 如何找到两个环形链表的交点

---

## 💡 面试核心考点（四大直击灵魂的拷问）

* **直击灵魂的拷问一：为什么慢指针走过的路程肯定是 x + y，而不是慢指针也绕了环好几圈呢？**
  绝杀回答：因为快指针的速度是慢指针的两倍。当慢指针刚刚走到环入口时，快指针一定已经在环内了。接下来，慢指针只需要走完不到一圈的距离，快指针就绝对能追上它。由于慢指针走完一圈的时间，足够快指针走两圈，所以快指针必将在慢指针走完第一圈之前将其截获。因此，相遇时慢指针走过的环内距离 y 一定小于环的周长，不可能绕圈。

* **直击灵魂的拷问二：请在白板上完整推导 x = (n - 1)(y + z) + z 这个公式！**
  推导过程：
  1. 慢指针路程：x + y
  2. 快指针路程：x + y + n(y + z)
  3. 2(x + y) = x + y + n(y + z)
  4. x + y = n(y + z)
  5. x = n(y + z) - y
  6. x = (n - 1)(y + z) + y + z - y
  7. x = (n - 1)(y + z) + z

* **直击灵魂的拷问三：相遇后，为什么一个指针从头出发，一个从相遇点出发，同速前进，相遇点就是环的入口？**
  看公式 x = (n - 1)(y + z) + z！x 代表头节点到入口的距离，z 代表相遇点到入口的距离。(n - 1)(y + z) 只是在环里绕圈的无用功。这意味着：当 index1 走完 x 的距离时，index2 刚好在环里绕了 n-1 圈加上 z 的距离。两人一定会精确无误地在"环的入口节点"相遇！

* **直击灵魂的拷问四：如果工程项目中不在乎空间复杂度，用哈希表怎么做？**
  直接遍历链表，把每次访问到的节点指针存入 unordered_set 中。每次插入前检查该指针是否已经存在于 set 里。如果存在，那个节点就是环的入口。时间复杂度和空间复杂度都是 O(n)。但在对内存极度苛刻的基础架构组件中，快慢指针的 O(1) 空间解法才是真正的工业级标准。

---

通过本课程，我们深入理解了环形链表 II 的核心逻辑，掌握了快慢指针法和哈希表法两种解法，也学习了详细的数学推导和常见的错误注意事项。这些知识对于理解链表和双指针技巧非常重要！
