# 206. 反转链表学习笔记

## 一、 题目分析

### 1.1 题目描述

给你单链表的头节点 `head` ，请你反转链表，并返回反转后的链表。

### 1.2 题目链接

- [LeetCode 206. 反转链表](https://leetcode.cn/problems/reverse-linked-list/description/)
- [代码随想录题解](https://www.programmercarl.com/0206.%E7%BF%BB%E8%BD%AC%E9%93%BE%E8%A1%A8.html)
- [官方题解](https://leetcode.cn/problems/reverse-linked-list/solutions/551596/fan-zhuan-lian-biao-by-leetcode-solution-d1k2/)

### 1.3 示例

**示例 1：**

```
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]
```

**示例 2：**

```
输入：head = [1,2]
输出：[2,1]
```

**示例 3：**

```
输入：head = []
输出：[]
```

### 1.4 提示

- 链表中节点的数目范围是 `[0, 5000]`
- `-5000 <= Node.val <= 5000`

### 1.5 题目本质

反转链表的本质是：**在不使用额外空间的前提下，原地（In-place）修改每个节点的指针指向**，将原本指向后继的指针改为指向前驱。

换句话说，我们需要：
1. 改变每个节点的 `next` 指针方向
2. 同时保证链表不会断裂
3. 最终返回新的头节点

### 1.6 题目进阶

链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？

---

## 二、 解题思路分析

### 2.1 核心思路

反转链表的核心思路是：**使用双指针来改变节点的指针方向**。

在反转过程中，我们需要注意以下几点：

| 关键点 | 说明 |
|-------|------|
| 断链保护 | 必须先保存下一个节点，再修改当前节点的指针 |
| 双指针配合 | 使用 `pre` 和 `cur` 两个指针配合工作 |
| 终止条件 | 当 `cur` 为 `nullptr` 时，说明链表已全部反转 |
| 返回值 | 最终返回 `pre`，即新的头节点 |

### 2.2 遍历顺序选择

对于反转链表，我们需要使用**从前往后**的顺序遍历，原因如下：

- 我们需要先处理前面的节点，再处理后面的节点
- 每个节点的反转只依赖于前一个节点，不依赖后一个节点
- 这种顺序可以保证我们能正确地保存和访问下一个节点

### 2.3 迭代法思路

#### 2.3.1 迭代三步曲

**1. 初始化三个指针**
- `pre`：前驱节点，初始为 `nullptr`
- `cur`：当前节点，初始为 `head`
- `temp`：临时节点，用于保存下一个节点

**2. 循环反转**

在每次循环中：

| 步骤 | 操作 | 目的 |
|------|------|------|
| 步骤 1 | `temp = cur->next` | 保存下一个节点，防止断链 |
| 步骤 2 | `cur->next = pre` | 核心反转操作，改变指针方向 |
| 步骤 3 | `pre = cur` | 前驱指针向前移动 |
| 步骤 4 | `cur = temp` | 当前指针向前移动 |

**3. 循环终止条件**
- 当 `cur == nullptr` 时，说明链表已全部反转

#### 2.3.2 迭代法图解

```
初始状态：
pre = nullptr  cur = 1 -> 2 -> 3 -> 4 -> 5 -> nullptr

第 1 次循环：
temp = cur->next (保存 2)
cur->next = pre (1 -> nullptr)
pre = cur (pre = 1)
cur = temp (cur = 2)
状态：pre = 1 -> nullptr  cur = 2 -> 3 -> 4 -> 5 -> nullptr

第 2 次循环：
temp = cur->next (保存 3)
cur->next = pre (2 -> 1)
pre = cur (pre = 2)
cur = temp (cur = 3)
状态：pre = 2 -> 1 -> nullptr  cur = 3 -> 4 -> 5 -> nullptr

... 以此类推 ...

最终状态：
pre = 5 -> 4 -> 3 -> 2 -> 1 -> nullptr  cur = nullptr
```

### 2.4 递归法思路

#### 2.4.1 递归三部曲

**1. 确定递归函数的参数和返回值**
- 参数：当前节点 `head`
- 返回值：反转后的链表头节点

```cpp
ListNode* reverseList(ListNode* head)
```

**2. 确定终止条件**
- 当 `head == nullptr` 或 `head->next == nullptr` 时，直接返回 `head`
- 说明已经到达链表尾部，或者链表为空

**3. 确定单层递归逻辑**

在单层递归逻辑中：
- 先递归反转后续的子链表
- 然后处理当前节点的指针反转
- 最后返回新的头节点

#### 2.4.2 递归法图解

```
初始调用：reverseList(1 -> 2 -> 3 -> 4 -> 5 -> nullptr)
         ↓
      递归进入：reverseList(2 -> 3 -> 4 -> 5 -> nullptr)
              ↓
           递归进入：reverseList(3 -> 4 -> 5 -> nullptr)
                   ↓
                递归进入：reverseList(4 -> 5 -> nullptr)
                        ↓
                     递归进入：reverseList(5 -> nullptr)
                             ↓
                          终止条件满足，返回 5
                     ↓
                  处理节点 4：
                  4->next->next = 4  (5->4)
                  4->next = nullptr
                  返回 5
           ↓
        处理节点 3：
        3->next->next = 3  (4->3)
        3->next = nullptr
        返回 5
↓
处理节点 2：
2->next->next = 2  (3->2)
2->next = nullptr
返回 5
↓
处理节点 1：
1->next->next = 1  (2->1)
1->next = nullptr
返回 5

最终返回：5 -> 4 -> 3 -> 2 -> 1 -> nullptr
```

### 2.5 核心思想：多指针贴身肉搏与断链保护

想象两个人手拉手，你想让第二个人回过头来牵第一个人的手。

**核心痛点：** 一旦第二个人（`cur`）松开牵着第三个人（原本的 `next`）的手，转而去牵第一个人（`pre`）的手，**第三个人就彻底在茫茫内存中走丢了！链表当场断裂！**

**破局思路：引入 `temp` 探路指针！**
在 `cur` 转身（改变指向）之前，先派 `temp` 去抓住原本的下一个节点。确保大部队不走丢后，`cur` 再放心地完成华丽的转身。

---

## 三、 算法实现

### 3.1 迭代法（双指针法）

```cpp
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* reverseList(ListNode* head) {
        ListNode* pre = nullptr; // 前驱节点，反转后它会成为新的头。一开始真实头节点的前驱是空。
        ListNode* cur = head;    // 当前正在处理的节点
        ListNode* temp;          // 探路兵：负责在断链前，死死抓住下一个节点

        // 当 cur 为 nullptr 时，说明链表已经全部反转完毕
        while (cur != nullptr) {
            // 步骤 1：探路兵先走！保存 cur 的下一个节点，防止转身后彻底失联
            temp = cur->next; 

            // 步骤 2：核心反转动作！cur 放开后面的手，转头指向前面的 pre
            cur->next = pre;

            // 步骤 3：整体舰队向前平移，准备处理下一个节点
            pre = cur;  // pre 移到 cur 的位置
            cur = temp; // cur 移到刚刚探路兵保存的位置
        }
        
        // 循环结束时，cur 指向了黑洞 (nullptr)，而 pre 刚好停在原链表的最后一个节点！
        // 此时的 pre，就是反转后的新链表的真实头节点！
        return pre; 
    }
};
```

### 3.2 递归法

```cpp
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* reverseList(ListNode* head) {
        // 终止条件：如果链表为空或只有一个节点，直接返回
        if (head == nullptr || head->next == nullptr) {
            return head;
        }
        
        // 递归反转后续的子链表
        ListNode* last = reverseList(head->next);
        
        // 反转当前节点的指针
        head->next->next = head;  // 让下一个节点指向当前节点
        head->next = nullptr;     // 当前节点指向 null，防止环
        
        // 返回反转后的头节点
        return last;
    }
};
```

### 3.3 递归法（另一种写法，带 pre 参数）

```cpp
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* reverseList(ListNode* head) {
        return reverse(nullptr, head);
    }
    
private:
    ListNode* reverse(ListNode* pre, ListNode* cur) {
        // 终止条件：当前节点为空，返回前驱节点
        if (cur == nullptr) {
            return pre;
        }
        
        // 保存下一个节点
        ListNode* temp = cur->next;
        
        // 反转当前节点的指针
        cur->next = pre;
        
        // 递归反转后续节点
        return reverse(cur, temp);
    }
};
```

---

## 四、 代码分析

### 4.1 迭代法代码分析

#### 4.1.1 实现思路说明

**迭代法的设计思路：**
1. 使用三个指针：`pre`（前驱）、`cur`（当前）、`temp`（临时保存）
2. 从前往后遍历链表，逐个反转每个节点的指针
3. 每次反转前先保存下一个节点，防止断链
4. 最终返回 `pre` 作为新的头节点

#### 4.1.2 核心逻辑讲解

**关键代码段 1：保存下一个节点**
```cpp
temp = cur->next;
```

**为什么需要这一步？**
- 一旦我们修改了 `cur->next` 的指向，就无法再找到原来的下一个节点
- 必须在修改之前先保存下来

**关键代码段 2：核心反转操作**
```cpp
cur->next = pre;
```

**这一步做了什么？**
- 将当前节点的 `next` 指针从指向后继改为指向前驱
- 完成了单个节点的反转

**关键代码段 3：指针向前移动**
```cpp
pre = cur;
cur = temp;
```

**为什么是这个顺序？**
- 先移动 `pre`，因为 `pre` 要跟上 `cur` 的步伐
- 后移动 `cur`，因为我们需要用之前保存的 `temp`

#### 4.1.3 初学者引导

**为什么循环条件是 `cur != nullptr` 而不是 `cur->next != nullptr`？**
- 如果使用 `cur->next != nullptr`，最后一个节点不会进入循环
- 最后一个节点的 `next` 指针就不会被反转
- 使用 `cur != nullptr` 才能保证每个节点都被处理

**为什么最后返回 `pre` 而不是 `cur`？**
- 循环结束时，`cur` 已经指向 `nullptr`
- `pre` 正好停在原链表的最后一个节点
- 这个节点现在就是新链表的头节点

### 4.2 递归法代码分析

#### 4.2.1 实现思路说明

**递归法的设计思路：**
1. 先递归处理后续的子链表
2. 子链表反转完成后，再处理当前节点
3. 让当前节点的下一个节点指向自己
4. 让当前节点指向 `nullptr`
5. 最后返回子链表反转后的头节点

#### 4.2.2 核心逻辑讲解

**关键代码段 1：终止条件**
```cpp
if (head == nullptr || head->next == nullptr) {
    return head;
}
```

**什么时候触发终止？**
- 链表为空
- 只有一个节点
- 到达链表尾部

**关键代码段 2：递归调用**
```cpp
ListNode* last = reverseList(head->next);
```

**这一步做了什么？**
- 递归进入子链表
- 子链表反转完成后，返回子链表的头节点
- 这个头节点也是整个链表的新头节点

**关键代码段 3：反转当前节点**
```cpp
head->next->next = head;
head->next = nullptr;
```

**为什么这么写？**
- `head->next` 已经是子链表的最后一个节点
- 让它指向 `head`，完成连接
- 让 `head->next` 指向 `nullptr`，防止形成环

#### 4.2.3 初学者引导

**递归会不会栈溢出？**
- 对于本题限制的 5000 个节点，递归深度就是 5000
- 一般情况下不会栈溢出（默认栈大小通常足够）
- 如果担心栈溢出，可以使用迭代法

**递归的优势是什么？**
- 代码简洁，逻辑清晰
- 不需要手动管理指针移动
- 符合"分而治之"的思想

---

## 五、 时间与空间复杂度分析

### 5.1 迭代法复杂度分析

| 复杂度类型 | 时间复杂度 | 空间复杂度 |
|-----------|-----------|-----------|
| **迭代法** | O(n) | O(1) |

**时间复杂度推导：**
- 使用 `cur` 指针从链表头走到尾
- 每个节点刚好被访问一次
- 每次循环的操作是常数时间 O(1)
- 总耗时随节点数量线性增长

**空间复杂度推导：**
- 只使用了 `pre`、`cur`、`temp` 三个临时指针
- 没有使用任何随数据规模增长的额外空间
- 完全符合原地反转的要求

### 5.2 递归法复杂度分析

| 复杂度类型 | 时间复杂度 | 空间复杂度 |
|-----------|-----------|-----------|
| **递归法** | O(n) | O(n) |

**时间复杂度推导：**
- 每个节点都被递归函数访问一次
- 每次递归调用的操作是常数时间 O(1)
- 总时间复杂度为 O(n)

**空间复杂度推导：**
- 递归调用栈的深度等于链表的长度
- 最坏情况下（链表完全线性），空间复杂度为 O(n)
- 这是递归法相对于迭代法的主要劣势

### 5.3 复杂度对比表格

| 方法 | 时间复杂度 | 空间复杂度 | 优点 | 缺点 |
|------|-----------|-----------|------|------|
| 迭代法 | O(n) | O(1) | 空间复杂度最优 | 需要手动管理指针 |
| 递归法 | O(n) | O(n) | 代码简洁，逻辑清晰 | 可能栈溢出，空间复杂度较高 |

---

## 六、 测试用例验证

### 6.1 测试用例 1：正常链表

**输入：**
```
head = [1,2,3,4,5]
```

**执行过程（迭代法）：**
1. 初始化：`pre = nullptr`, `cur = 1`
2. 第 1 次循环：保存 `2`，`1->nullptr`，`pre = 1`，`cur = 2`
3. 第 2 次循环：保存 `3`，`2->1`，`pre = 2`，`cur = 3`
4. 第 3 次循环：保存 `4`，`3->2`，`pre = 3`，`cur = 4`
5. 第 4 次循环：保存 `5`，`4->3`，`pre = 4`，`cur = 5`
6. 第 5 次循环：保存 `nullptr`，`5->4`，`pre = 5`，`cur = nullptr`
7. 循环结束，返回 `pre = 5`

**输出：** `[5,4,3,2,1]`

### 6.2 测试用例 2：两个节点

**输入：**
```
head = [1,2]
```

**执行过程：**
1. 初始化：`pre = nullptr`, `cur = 1`
2. 第 1 次循环：保存 `2`，`1->nullptr`，`pre = 1`，`cur = 2`
3. 第 2 次循环：保存 `nullptr`，`2->1`，`pre = 2`，`cur = nullptr`
4. 循环结束，返回 `pre = 2`

**输出：** `[2,1]`

### 6.3 测试用例 3：空链表

**输入：**
```
head = []
```

**执行过程：**
1. 初始化：`pre = nullptr`, `cur = nullptr`
2. 循环条件 `cur != nullptr` 不满足，直接返回 `pre = nullptr`

**输出：** `[]`

### 6.4 测试用例 4：单个节点

**输入：**
```
head = [1]
```

**执行过程：**
1. 初始化：`pre = nullptr`, `cur = 1`
2. 第 1 次循环：保存 `nullptr`，`1->nullptr`，`pre = 1`，`cur = nullptr`
3. 循环结束，返回 `pre = 1`

**输出：** `[1]`

---

## 七、 常见错误与注意事项

### 7.1 断链问题

**错误示例：**
```cpp
// 错误：没有保存下一个节点
ListNode* cur = head;
while (cur != nullptr) {
    cur->next = pre;  // 直接修改指针，导致后面的节点丢失！
    pre = cur;
    cur = cur->next;  // 此时 cur->next 已经变了！
}
```

**正确做法：**
```cpp
ListNode* cur = head;
while (cur != nullptr) {
    ListNode* temp = cur->next;  // 先保存
    cur->next = pre;             // 再修改
    pre = cur;
    cur = temp;
}
```

### 7.2 循环条件错误

**错误示例：**
```cpp
// 错误：循环条件不正确
while (cur->next != nullptr) {  // 最后一个节点不会进入循环！
    // ...
}
```

**后果：**
- 最后一个节点的 `next` 指针不会被反转
- 链表反转不完全

**正确做法：**
```cpp
while (cur != nullptr) {  // 每个节点都要处理
    // ...
}
```

### 7.3 返回值错误

**错误示例：**
```cpp
// 错误：返回了 cur
return cur;  // cur 现在是 nullptr！
```

**正确做法：**
```cpp
return pre;  // pre 是新的头节点
```

### 7.4 递归形成环

**错误示例：**
```cpp
// 错误：没有设置 head->next = nullptr
head->next->next = head;
// 缺少 head->next = nullptr;
```

**后果：**
- 链表会形成环
- 后续操作可能导致死循环

**正确做法：**
```cpp
head->next->next = head;
head->next = nullptr;  // 防止形成环
```

### 7.5 空指针访问

**错误示例：**
```cpp
// 错误：没有检查空指针
ListNode* next = head->next;  // 如果 head 是 nullptr，这里会崩溃！
```

**正确做法：**
```cpp
if (head == nullptr || head->next == nullptr) {
    return head;  // 先检查空指针
}
```

---

## 八、 与其他相关题目的对比分析

### 8.1 与 92. 反转链表 II 对比

| 题目 | 核心问题 | 解法 |
|------|---------|------|
| 206. 反转链表 | 反转整个链表 | 双指针从前往后 |
| 92. 反转链表 II | 反转链表的一部分 | 先定位，再反转 |

**关系：**
- 92 题是 206 题的扩展
- 206 题可以看作是 92 题的特例（left=1, right=n）

### 8.2 与 234. 回文链表对比

| 题目 | 核心问题 | 解法 |
|------|---------|------|
| 206. 反转链表 | 反转整个链表 | 双指针/递归 |
| 234. 回文链表 | 判断链表是否是回文 | 快慢指针 + 反转后半部分 |

**关系：**
- 234 题可以用到 206 题的解法
- 反转后半部分链表后，与前半部分比较

### 8.3 与 25. K 个一组翻转链表对比

| 题目 | 核心问题 | 解法 |
|------|---------|------|
| 206. 反转链表 | 反转整个链表 | 双指针/递归 |
| 25. K 个一组翻转链表 | 每 K 个节点翻转一次 | 分组翻转 + 拼接 |

**关系：**
- 25 题需要多次调用 206 题的解法
- 每个 K 长度的小组内部就是一个 206 问题

---

## 九、 面试高频问题与解答

### 9.1 基础问题

#### Q1: 反转链表的核心思路是什么？

**答案：**

**核心思路：**
1. 使用双指针（`pre` 和 `cur`）来改变节点的指针方向
2. 必须先保存下一个节点，再修改当前节点的指针（断链保护）
3. 从前往后遍历，逐个反转每个节点
4. 最终返回 `pre` 作为新的头节点

**核心代码：**
```cpp
ListNode* pre = nullptr;
ListNode* cur = head;
while (cur != nullptr) {
    ListNode* temp = cur->next;
    cur->next = pre;
    pre = cur;
    cur = temp;
}
return pre;
```

#### Q2: 循环的条件为什么是 while(cur != nullptr) 而不是 while(cur->next != nullptr)？

**答案：**

**如果使用 cur->next != nullptr：**
- 当 cur 走到原链表的最后一个节点时，循环就会终止
- 最后一个节点的 next 指针根本没有被反转
- 它依然指向 nullptr 而没有指向倒数第二个节点

**使用 cur != nullptr：**
- 保证每一个节点、包括尾节点都被拉进循环里
- 每个节点都能完成一次转身

### 9.2 进阶问题

#### Q3: 最后为什么要 return pre;，而不是 return cur;？

**答案：**

**因为：**
- while 循环终止的唯一条件就是 cur 变成了 nullptr
- 如果你返回 cur，你返回的就是一个空指针
- 而此时的 pre 恰好是紧紧跟在 cur 身后的
- 它停在了原链表的最后一个节点上
- 在全部反转完成后，这个节点正式加冕为新链表的头节点

**所以：必须返回 pre！**

#### Q4: 递归法和迭代法各有什么优缺点？

**答案：**

**迭代法：**
- **优点：** 空间复杂度 O(1)，不会栈溢出，效率较高
- **缺点：** 代码相对复杂，需要手动管理指针

**递归法：**
- **优点：** 代码简洁，逻辑清晰，易于理解
- **缺点：** 空间复杂度 O(n)，可能栈溢出

**选择建议：**
- 面试时两种都要会写
- 实际使用时迭代法更常用

#### Q5: 如何证明反转后的链表是正确的？

**答案：**

**验证方法：**
1. 检查新的头节点是否是原链表的尾节点
2. 检查每个节点的 next 指针是否正确
3. 检查链表长度是否不变
4. 检查节点值是否都在，且顺序相反

**边界测试：**
- 空链表
- 单个节点
- 两个节点
- 大量节点

### 9.3 实际问题

#### Q6: 在实际项目中，为什么很少需要手动反转链表？

**答案：**

**原因：**
1. **STL 容器**：实际项目中更多使用 std::list 或 std::vector
2. **反向迭代器**：如果需要反向遍历，直接使用反向迭代器即可
3. **设计模式**：一般会使用适配器模式或其他方式，而不是直接反转数据结构

**什么时候需要手动反转？**
- 算法题
- 性能要求极高的底层代码
- 实现某些特定的数据结构（如 LRU Cache）

#### Q7: 如何实现双向链表的反转？

**答案：**

**双向链表反转的思路：**
1. 对于每个节点，交换 `next` 和 `prev` 指针
2. 遍历完成后，交换头指针和尾指针

**关键代码：**
```cpp
while (cur != nullptr) {
    Node* temp = cur->next;
    cur->next = cur->prev;
    cur->prev = temp;
    if (temp == nullptr) {
        newHead = cur;
    }
    cur = temp;
}
```

---

## 十、 完整可运行代码

```cpp
#include <iostream>
#include <vector>
using namespace std;

// ==========================================
// 链表节点定义
// ==========================================
struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

// ==========================================
// 方法一：迭代法（双指针法）
// ==========================================
class IterativeSolution {
public:
    ListNode* reverseList(ListNode* head) {
        ListNode* pre = nullptr; // 前驱节点，反转后它会成为新的头
        ListNode* cur = head;    // 当前正在处理的节点
        ListNode* temp;          // 探路兵：负责在断链前，死死抓住下一个节点

        // 当 cur 为 nullptr 时，说明链表已经全部反转完毕
        while (cur != nullptr) {
            // 步骤 1：探路兵先走！保存 cur 的下一个节点，防止转身后彻底失联
            temp = cur->next; 

            // 步骤 2：核心反转动作！cur 放开后面的手，转头指向前面的 pre
            cur->next = pre;

            // 步骤 3：整体舰队向前平移，准备处理下一个节点
            pre = cur;  // pre 移到 cur 的位置
            cur = temp; // cur 移到刚刚探路兵保存的位置
        }
        
        // 循环结束时，cur 指向了黑洞 (nullptr)，而 pre 刚好停在原链表的最后一个节点！
        // 此时的 pre，就是反转后的新链表的真实头节点！
        return pre; 
    }
};

// ==========================================
// 方法二：递归法
// ==========================================
class RecursiveSolution {
public:
    ListNode* reverseList(ListNode* head) {
        // 终止条件：如果链表为空或只有一个节点，直接返回
        if (head == nullptr || head->next == nullptr) {
            return head;
        }
        
        // 递归反转后续的子链表
        ListNode* last = reverseList(head->next);
        
        // 反转当前节点的指针
        head->next->next = head;  // 让下一个节点指向当前节点
        head->next = nullptr;     // 当前节点指向 null，防止环
        
        // 返回反转后的头节点
        return last;
    }
};

// ==========================================
// 方法三：递归法（另一种写法，带 pre 参数）
// ==========================================
class RecursiveSolution2 {
public:
    ListNode* reverseList(ListNode* head) {
        return reverse(nullptr, head);
    }
    
private:
    ListNode* reverse(ListNode* pre, ListNode* cur) {
        // 终止条件：当前节点为空，返回前驱节点
        if (cur == nullptr) {
            return pre;
        }
        
        // 保存下一个节点
        ListNode* temp = cur->next;
        
        // 反转当前节点的指针
        cur->next = pre;
        
        // 递归反转后续节点
        return reverse(cur, temp);
    }
};

// ==========================================
// 辅助函数：创建链表
// ==========================================
ListNode* createList(const vector<int>& vals) {
    ListNode dummy;
    ListNode* cur = &dummy;
    for (int val : vals) {
        cur->next = new ListNode(val);
        cur = cur->next;
    }
    return dummy.next;
}

// ==========================================
// 辅助函数：打印链表
// ==========================================
void printList(ListNode* head) {
    cout << "[";
    while (head != nullptr) {
        cout << head->val;
        if (head->next != nullptr) {
            cout << ",";
        }
        head = head->next;
    }
    cout << "]" << endl;
}

// ==========================================
// 主函数测试
// ==========================================
int main() {
    cout << "========================================" << endl;
    cout << "        206. 反转链表测试" << endl;
    cout << "========================================" << endl;
    
    // ==========================================
    // 测试用例 1：正常链表
    // ==========================================
    cout << "\n--- 测试用例 1 ---" << endl;
    vector<int> vals1 = {1, 2, 3, 4, 5};
    ListNode* list1 = createList(vals1);
    cout << "原始链表：";
    printList(list1);
    
    IterativeSolution sol1;
    ListNode* reversed1 = sol1.reverseList(list1);
    cout << "反转链表：";
    printList(reversed1);
    cout << "期望：[5,4,3,2,1]" << endl;
    
    // ==========================================
    // 测试用例 2：两个节点
    // ==========================================
    cout << "\n--- 测试用例 2 ---" << endl;
    vector<int> vals2 = {1, 2};
    ListNode* list2 = createList(vals2);
    cout << "原始链表：";
    printList(list2);
    
    RecursiveSolution sol2;
    ListNode* reversed2 = sol2.reverseList(list2);
    cout << "反转链表：";
    printList(reversed2);
    cout << "期望：[2,1]" << endl;
    
    // ==========================================
    // 测试用例 3：空链表
    // ==========================================
    cout << "\n--- 测试用例 3 ---" << endl;
    vector<int> vals3 = {};
    ListNode* list3 = createList(vals3);
    cout << "原始链表：";
    printList(list3);
    
    ListNode* reversed3 = sol1.reverseList(list3);
    cout << "反转链表：";
    printList(reversed3);
    cout << "期望：[]" << endl;
    
    // ==========================================
    // 测试用例 4：单个节点
    // ==========================================
    cout << "\n--- 测试用例 4 ---" << endl;
    vector<int> vals4 = {1};
    ListNode* list4 = createList(vals4);
    cout << "原始链表：";
    printList(list4);
    
    RecursiveSolution2 sol3;
    ListNode* reversed4 = sol3.reverseList(list4);
    cout << "反转链表：";
    printList(reversed4);
    cout << "期望：[1]" << endl;
    
    cout << "\n========================================" << endl;
    cout << "            测试完成！" << endl;
    cout << "========================================" << endl;
    
    return 0;
}
```

---

## 十一、 总结与反思

### 11.1 核心知识点总结

| 知识点 | 要点 |
|--------|------|
| 反转链表 | 在原链表上原地修改指针指向 |
| 迭代法 | 双指针从前往后，逐个反转 |
| 递归法 | 先处理子链表，再处理当前节点 |
| 断链保护 | 必须先保存下一个节点，再修改指针 |
| 返回值 | 返回 pre 而不是 cur |
| 时间复杂度 | O(n)，两种方法都是 |
| 空间复杂度 | 迭代法 O(1)，递归法 O(n) |

### 11.2 学习收获

1. **理解了反转链表的核心逻辑**
   - 如何使用双指针配合工作
   - 为什么需要断链保护
   - 为什么返回 pre 而不是 cur

2. **掌握了两种不同的解法**
   - 迭代法：空间复杂度最优
   - 递归法：代码简洁，逻辑清晰

3. **理解了递归的本质**
   - 先深入底层，再逐层返回
   - 子问题的解如何构成父问题的解

4. **掌握了常见的错误点**
   - 断链问题
   - 循环条件错误
   - 返回值错误
   - 空指针访问

### 11.3 后续学习建议

1. **学习更多链表操作**
   - 反转链表 II（局部反转）
   - K 个一组翻转链表
   - 回文链表

2. **深入学习递归**
   - 递归的调用栈
   - 尾递归优化
   - 如何将递归改为迭代

3. **练习更多链表题目**
   - LeetCode 19. 删除链表的倒数第 N 个节点
   - LeetCode 21. 合并两个有序链表
   - LeetCode 234. 回文链表

4. **学习双向链表**
   - 双向链表的反转
   - 双向链表的插入和删除

---

通过本课程，我们深入理解了反转链表的核心逻辑，掌握了迭代法和递归法两种解法，也学习了常见的错误和注意事项。这些知识对于理解链表操作和指针操作非常重要！