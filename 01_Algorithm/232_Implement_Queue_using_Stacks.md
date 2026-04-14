# 算法实战复盘：用栈实现队列 (LeetCode 232)

> **打卡日期**：2026-04-14
> **核心主题**：栈、队列、数据结构设计、Amortized Time Complexity。

---

## 📝 一、 题目描述与核心要求

### 1. 中文描述
请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：

实现 MyQueue 类：
- `void push(int x)` 将元素 x 推到队列的末尾
- `int pop()` 从队列的开头移除并返回元素
- `int peek()` 返回队列开头的元素
- `boolean empty()` 如果队列为空，返回 true ；否则，返回 false

说明：
- 你 只能 使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
- 你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。

### 2. 英文描述
Implement a first-in-first-out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue (`push`, `pop`, `peek`, and `empty`).

Implement the `MyQueue` class:
- `void push(int x)` Pushes element x to the back of the queue.
- `int pop()` Removes the element from the front of the queue and returns it.
- `int peek()` Returns the element at the front of the queue.
- `boolean empty()` Returns true if the queue is empty, false otherwise.

Notes:
- You must use only standard stack operations — which means only `push to top`, `peek/pop from top`, `size`, and `is empty` operations are valid.
- Depending on your language, the stack may not be supported natively. You may simulate a stack using a list or deque (double-ended queue) as long as you use only a stack's standard operations.

### 3. 输入示例
- **示例 1**：
  - 输入：
    ```
    ["MyQueue", "push", "push", "peek", "pop", "empty"]
    [[], [1], [2], [], [], []]
    ```
  - 输出：
    ```
    [null, null, null, 1, 1, false]
    ```
  - 解释：
    ```
    MyQueue myQueue = new MyQueue();
    myQueue.push(1); // queue is: [1]
    myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
    myQueue.peek(); // return 1
    myQueue.pop(); // return 1, queue is [2]
    myQueue.empty(); // return false
    ```

### 4. 核心痛点
- 如何用后进先出（LIFO）的栈实现先进先出（FIFO）的队列
- 如何优化操作的时间复杂度，尤其是 pop 和 peek 操作
- 如何处理边界情况，如空队列的操作
- 如何设计数据结构，使得代码简洁且高效

### 5. 题目提示
- 1 <= x <= 9
- 最多调用 100 次 push、pop、peek 和 empty
- 假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）

---

## 📊 二、 复杂度深度剖析

### 1. 时间复杂度

#### 方法一：双栈优化法（推荐）

| 操作 | 时间复杂度 | 说明 |
|------|------------|------|
| push | O(1) | 直接将元素压入输入栈，常数时间操作 |
| pop | 均摊 O(1) | 当输出栈为空时，需要将输入栈的所有元素转移到输出栈，时间复杂度为 O(n)；但每个元素只会被转移一次，所以均摊时间复杂度为 O(1) |
| peek | 均摊 O(1) | 复用 pop 操作，然后将元素放回，均摊时间复杂度为 O(1) |
| empty | O(1) | 只需检查两个栈是否都为空 |

#### 方法二：暴力解法

| 操作 | 时间复杂度 | 说明 |
|------|------------|------|
| push | O(1) | 直接将元素压入主栈，常数时间操作 |
| pop | O(n) | 每次都需要将主栈的元素转移到辅助栈，弹出队首元素后再转移回来，时间复杂度为 O(n) |
| peek | O(n) | 类似 pop 操作，需要将主栈的元素转移到辅助栈，查看队首元素后再转移回来，时间复杂度为 O(n) |
| empty | O(1) | 只需检查主栈是否为空 |

#### 方法三：双栈法（peek 优化版）

| 操作 | 时间复杂度 | 说明 |
|------|------------|------|
| push | O(1) | 直接将元素压入输入栈，常数时间操作 |
| pop | 均摊 O(1) | 当输出栈为空时，需要将输入栈的所有元素转移到输出栈，时间复杂度为 O(n)；但每个元素只会被转移一次，所以均摊时间复杂度为 O(1) |
| peek | 均摊 O(1) | 当输出栈为空时需要转移元素，否则直接返回栈顶元素，均摊时间复杂度为 O(1) |
| empty | O(1) | 只需检查两个栈是否都为空 |

### 2. 空间复杂度

| 方法 | 空间复杂度 | 说明 |
|------|------------|------|
| 方法一 | O(n) | 两个栈最多存储 n 个元素，没有额外的空间开销 |
| 方法二 | O(n) | 主栈存储 n 个元素，辅助栈在操作时最多存储 n-1 个元素，总空间复杂度为 O(n) |
| 方法三 | O(n) | 两个栈最多存储 n 个元素，没有额外的空间开销 |

---

## 🤔 三、 解题思路分析

### 1. 初学者的常见思路

当面对这个问题时，初学者可能会想到以下几种方法：

1. **暴力解法**：每次 pop 或 peek 时，将所有元素从一个栈转移到另一个栈，操作完成后再转移回去
2. **优化解法**：使用两个栈，一个作为输入栈，一个作为输出栈，只有当输出栈为空时才进行元素转移

### 2. 各种方法的优缺点分析

| 方法 | 时间复杂度 | 空间复杂度 | 优点 | 缺点 |
|------|------------|------------|------|------|
| 暴力解法 | O(n) per pop/peek | O(n) | 实现简单，逻辑直观 | 时间复杂度高，每次操作都需要转移所有元素 |
| 优化解法 | 均摊 O(1) per pop/peek | O(n) | 时间复杂度低，操作高效 | 实现稍微复杂，需要理解元素转移的时机 |

### 3. 数据大小对方法选择的影响

根据题目提示，最多调用 100 次操作：

- **暴力解法**：对于 100 次操作，每次最多转移 100 个元素，总操作次数为 100*100=10000 次，在时间上是可行的
- **优化解法**：总操作次数约为 100 次，性能更优

**实际测试验证**：
- 暴力解法：在小规模数据上性能尚可，但在大规模数据上会明显变慢
- 优化解法：在各种规模的数据上性能都很稳定

### 4. 方法选择建议

**推荐使用优化解法**，原因如下：
1. 均摊时间复杂度 O(1)，性能更优
2. 实现相对简单，逻辑清晰
3. 适用于各种规模的数据
4. 是解决此类问题的经典方法，体现了数据结构设计的智慧

---

## 🎯 四、 核心思想：双栈实现队列

### 1. 基本原理
使用两个栈：
- **输入栈（stIn）**：用于处理 push 操作，新元素总是压入输入栈
- **输出栈（stOut）**：用于处理 pop 和 peek 操作

### 2. 操作流程
1. **push 操作**：直接将元素压入输入栈
2. **pop 操作**：
   - 如果输出栈为空，将输入栈的所有元素转移到输出栈
   - 从输出栈弹出栈顶元素并返回
3. **peek 操作**：
   - 如果输出栈为空，将输入栈的所有元素转移到输出栈
   - 返回输出栈的栈顶元素
4. **empty 操作**：检查两个栈是否都为空

### 3. 关键设计点
- **元素转移的时机**：只有当输出栈为空时才进行转移，避免重复转移
- **peek 操作的实现**：可以复用 pop 操作，然后将弹出的元素重新压回输出栈

---

## 💻 五、 代码实现与解析

### 1. C++ 实现方法一：双栈优化法（推荐）

```cpp
#include <stack>

class MyQueue {
private:
    std::stack<int> stIn;  // 输入栈，用于push操作
    std::stack<int> stOut; // 输出栈，用于pop和peek操作

public:
    /** Initialize your data structure here. */
    MyQueue() {
        // 构造函数，无需特殊初始化
    }
    
    /** Push element x to the back of queue. */
    void push(int x) {
        // 直接将元素压入输入栈
        stIn.push(x);
    }
    
    /** Removes the element from in front of queue and returns that element. */
    int pop() {
        // 如果输出栈为空，将输入栈的所有元素转移到输出栈
        if (stOut.empty()) {
            while (!stIn.empty()) {
                stOut.push(stIn.top());
                stIn.pop();
            }
        }
        // 从输出栈弹出元素
        int result = stOut.top();
        stOut.pop();
        return result;
    }
    
    /** Get the front element. */
    int peek() {
        // 复用pop操作，然后将元素放回
        int res = this->pop();
        stOut.push(res);
        return res;
    }
    
    /** Returns whether the queue is empty. */
    bool empty() {
        // 只有当两个栈都为空时，队列才为空
        return stIn.empty() && stOut.empty();
    }
};

/**
 * Your MyQueue object will be instantiated and called as such:
 * MyQueue* obj = new MyQueue();
 * obj->push(x);
 * int param_2 = obj->pop();
 * int param_3 = obj->peek();
 * bool param_4 = obj->empty();
 */
```

### 2. C++ 实现方法二：暴力解法

```cpp
#include <stack>

class MyQueue {
private:
    std::stack<int> stack1; // 主栈

public:
    /** Initialize your data structure here. */
    MyQueue() {
        // 构造函数，无需特殊初始化
    }
    
    /** Push element x to the back of queue. */
    void push(int x) {
        // 直接将元素压入主栈
        stack1.push(x);
    }
    
    /** Removes the element from in front of queue and returns that element. */
    int pop() {
        std::stack<int> stack2; // 辅助栈
        // 将主栈的元素转移到辅助栈，除了最后一个元素
        while (stack1.size() > 1) {
            stack2.push(stack1.top());
            stack1.pop();
        }
        // 弹出并保存主栈的最后一个元素（队首元素）
        int result = stack1.top();
        stack1.pop();
        // 将辅助栈的元素转移回主栈
        while (!stack2.empty()) {
            stack1.push(stack2.top());
            stack2.pop();
        }
        return result;
    }
    
    /** Get the front element. */
    int peek() {
        std::stack<int> stack2; // 辅助栈
        // 将主栈的元素转移到辅助栈，除了最后一个元素
        while (stack1.size() > 1) {
            stack2.push(stack1.top());
            stack1.pop();
        }
        // 保存主栈的最后一个元素（队首元素）
        int result = stack1.top();
        // 将辅助栈的元素转移回主栈
        while (!stack2.empty()) {
            stack1.push(stack2.top());
            stack2.pop();
        }
        return result;
    }
    
    /** Returns whether the queue is empty. */
    bool empty() {
        return stack1.empty();
    }
};

/**
 * Your MyQueue object will be instantiated and called as such:
 * MyQueue* obj = new MyQueue();
 * obj->push(x);
 * int param_2 = obj->pop();
 * int param_3 = obj->peek();
 * bool param_4 = obj->empty();
 */
```

### 3. C++ 实现方法三：双栈法（peek 优化版）

```cpp
#include <stack>

class MyQueue {
private:
    std::stack<int> stIn;  // 输入栈
    std::stack<int> stOut; // 输出栈

public:
    /** Initialize your data structure here. */
    MyQueue() {
        // 构造函数，无需特殊初始化
    }
    
    /** Push element x to the back of queue. */
    void push(int x) {
        stIn.push(x);
    }
    
    /** Removes the element from in front of queue and returns that element. */
    int pop() {
        // 确保输出栈有元素
        if (stOut.empty()) {
            moveElements();
        }
        int result = stOut.top();
        stOut.pop();
        return result;
    }
    
    /** Get the front element. */
    int peek() {
        // 确保输出栈有元素
        if (stOut.empty()) {
            moveElements();
        }
        return stOut.top();
    }
    
    /** Returns whether the queue is empty. */
    bool empty() {
        return stIn.empty() && stOut.empty();
    }
    
private:
    /** 将输入栈的元素转移到输出栈 */
    void moveElements() {
        while (!stIn.empty()) {
            stOut.push(stIn.top());
            stIn.pop();
        }
    }
};

/**
 * Your MyQueue object will be instantiated and called as such:
 * MyQueue* obj = new MyQueue();
 * obj->push(x);
 * int param_2 = obj->pop();
 * int param_3 = obj->peek();
 * bool param_4 = obj->empty();
 */
```

### 4. 代码解析

#### 4.1 方法一：双栈优化法（推荐）

**核心数据结构**：
- 使用两个栈：`stIn` 用于输入，`stOut` 用于输出
- 栈的特性是后进先出，通过两个栈的配合实现先进先出的队列特性

**操作分析**：
- **push 操作**：O(1)，直接将元素压入 `stIn` 栈
- **pop 操作**：均摊 O(1)，当 `stOut` 为空时，将 `stIn` 中的所有元素转移到 `stOut`，然后从 `stOut` 弹出栈顶元素
- **peek 操作**：均摊 O(1)，复用 pop 操作，然后将弹出的元素重新压回 `stOut`
- **empty 操作**：O(1)，检查两个栈是否都为空

**优点**：
- 时间复杂度低，均摊 O(1) per operation
- 代码简洁，逻辑清晰
- 适用于各种规模的数据

#### 4.2 方法二：暴力解法

**核心数据结构**：
- 使用一个主栈 `stack1`，每次操作时使用临时辅助栈 `stack2`

**操作分析**：
- **push 操作**：O(1)，直接将元素压入主栈
- **pop 操作**：O(n)，每次都需要将主栈的元素转移到辅助栈，弹出队首元素后再转移回来
- **peek 操作**：O(n)，类似 pop 操作，但不弹出元素
- **empty 操作**：O(1)，检查主栈是否为空

**优点**：
- 实现简单，逻辑直观

**缺点**：
- 时间复杂度高，每次 pop 和 peek 操作都需要 O(n) 时间
- 性能较差，不适合大规模数据

#### 4.3 方法三：双栈法（peek 优化版）

**核心数据结构**：
- 使用两个栈：`stIn` 用于输入，`stOut` 用于输出
- 提取了元素转移的逻辑到单独的 `moveElements` 方法

**操作分析**：
- **push 操作**：O(1)，直接将元素压入 `stIn` 栈
- **pop 操作**：均摊 O(1)，当 `stOut` 为空时，调用 `moveElements` 方法转移元素
- **peek 操作**：均摊 O(1)，当 `stOut` 为空时，调用 `moveElements` 方法转移元素，然后返回 `stOut` 的栈顶元素
- **empty 操作**：O(1)，检查两个栈是否都为空

**优点**：
- 时间复杂度低，均摊 O(1) per operation
- 代码结构更清晰，将元素转移逻辑提取为单独的方法
- peek 操作直接返回栈顶元素，避免了元素的弹出和压回操作

**与方法一的区别**：
- 方法三的 peek 操作直接返回输出栈的栈顶元素，而不是通过 pop 操作实现
- 方法三将元素转移逻辑提取为单独的方法，代码结构更清晰

---

## 🧪 六、 测试用例验证

### 1. 测试用例 1：基本操作

**输入**：
```
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
```

**预期输出**：
```
[null, null, null, 1, 1, false]
```

**执行过程**：
1. `MyQueue()` - 初始化队列
2. `push(1)` - 将 1 压入 stIn，stIn: [1]
3. `push(2)` - 将 2 压入 stIn，stIn: [1, 2]
4. `peek()` - stOut 为空，将 stIn 元素转移到 stOut，stOut: [2, 1]，返回 1
5. `pop()` - 从 stOut 弹出 1，stOut: [2]
6. `empty()` - stIn 为空，stOut 不为空，返回 false

**实际输出**：`[null, null, null, 1, 1, false]` ✅

### 2. 测试用例 2：连续 pop 操作

**输入**：
```
["MyQueue", "push", "push", "push", "pop", "pop", "peek", "pop", "empty"]
[[], [1], [2], [3], [], [], [], [], []]
```

**预期输出**：
```
[null, null, null, null, 1, 2, 3, 3, true]
```

**执行过程**：
1. `MyQueue()` - 初始化队列
2. `push(1)` - stIn: [1]
3. `push(2)` - stIn: [1, 2]
4. `push(3)` - stIn: [1, 2, 3]
5. `pop()` - stOut 为空，转移元素，stOut: [3, 2, 1]，弹出 1
6. `pop()` - 从 stOut 弹出 2
7. `peek()` - 返回 stOut 栈顶元素 3
8. `pop()` - 从 stOut 弹出 3
9. `empty()` - 两个栈都为空，返回 true

**实际输出**：`[null, null, null, null, 1, 2, 3, 3, true]` ✅

### 3. 测试用例 3：交替 push 和 pop 操作

**输入**：
```
["MyQueue", "push", "pop", "push", "push", "peek", "push", "pop", "empty"]
[[], [1], [], [2], [3], [], [4], [], []]
```

**预期输出**：
```
[null, null, 1, null, null, 2, null, 2, false]
```

**执行过程**：
1. `MyQueue()` - 初始化队列
2. `push(1)` - stIn: [1]
3. `pop()` - stOut 为空，转移元素，stOut: [1]，弹出 1
4. `push(2)` - stIn: [2]
5. `push(3)` - stIn: [2, 3]
6. `peek()` - stOut 为空，转移元素，stOut: [3, 2]，返回 2
7. `push(4)` - stIn: [4]
8. `pop()` - 从 stOut 弹出 2
9. `empty()` - stIn 不为空，返回 false

**实际输出**：`[null, null, 1, null, null, 2, null, 2, false]` ✅

---

## ⚠️ 七、 常见错误与注意事项

### 1. 常见错误

#### 1.1 每次操作都转移元素
**错误**：每次 pop 或 peek 时都将所有元素在两个栈之间转移
**后果**：时间复杂度变为 O(n)  per operation，性能较差
**避免方法**：只有当输出栈为空时才进行元素转移

#### 1.2 错误处理空队列
**错误**：在 pop 或 peek 操作时没有检查队列是否为空
**后果**：可能导致栈的空指针异常或未定义行为
**避免方法**：题目假设所有操作都是有效的，所以不需要额外的错误处理

#### 1.3 错误实现 peek 操作
**错误**：peek 操作没有保持队列状态不变
**后果**：可能导致元素丢失或重复
**避免方法**：peek 操作应该返回队首元素但不删除它，可以通过复用 pop 操作并将元素放回实现

#### 1.4 错误判断队列是否为空
**错误**：只检查输入栈是否为空
**后果**：可能在输出栈还有元素时错误地判断队列为空
**避免方法**：同时检查两个栈是否都为空

### 2. 注意事项

#### 2.1 元素转移的时机
- 只有当输出栈为空时才进行元素转移，这是实现均摊 O(1) 时间复杂度的关键
- 转移时需要将输入栈的所有元素全部转移到输出栈

#### 2.2 代码复用
- peek 操作可以复用 pop 操作，然后将弹出的元素重新压回输出栈
- 这样可以减少代码重复，提高代码的可维护性

#### 2.3 栈的选择
- 在某些语言中，栈可能不是内置数据结构，可以使用列表或双端队列来模拟
- 只要保证只使用栈的标准操作（push to top, peek/pop from top, size, is empty）即可

---

## 🚀 八、 相关拓展问题及解题思路

### 1. 用队列实现栈（LeetCode 225）

**题目链接**：https://leetcode.cn/problems/implement-stack-using-queues/

**解题思路**：
- 使用两个队列，一个主队列，一个辅助队列
- push 操作：将元素加入主队列
- pop 操作：将主队列的前 n-1 个元素转移到辅助队列，然后弹出主队列的最后一个元素，再将辅助队列的元素转移回主队列
- top 操作：类似 pop 操作，但需要将弹出的元素重新压回
- empty 操作：检查主队列是否为空

### 2. 设计循环队列（LeetCode 622）

**题目链接**：https://leetcode.cn/problems/design-circular-queue/

**解题思路**：
- 使用数组实现，维护 head 和 tail 指针
- 空队列条件：head == tail
- 满队列条件：(tail + 1) % capacity == head
- enQueue 操作：检查队列是否已满，然后将元素加入 tail 位置，tail 指针后移
- deQueue 操作：检查队列是否为空，然后将 head 指针后移
- Front 操作：返回 head 位置的元素
- Rear 操作：返回 tail-1 位置的元素

### 3. 设计双端队列（LeetCode 641）

**题目链接**：https://leetcode.cn/problems/design-circular-deque/

**解题思路**：
- 类似循环队列，但支持两端的插入和删除操作
- 需要实现：insertFront, insertLast, deleteFront, deleteLast, getFront, getRear, isEmpty, isFull 等操作
- 可以使用数组或链表实现

---

## 📌 九、 总结与思考

### 1. 核心知识点
- **栈和队列的基本特性**：栈是后进先出（LIFO），队列是先进先出（FIFO）
- **双栈实现队列**：使用两个栈，一个用于输入，一个用于输出，通过元素转移实现队列特性
- **均摊时间复杂度**：虽然某些操作的最坏时间复杂度为 O(n)，但通过均摊分析，整体时间复杂度为 O(1)
- **代码复用**：通过复用已有操作，减少代码重复，提高代码可维护性

### 2. 解题技巧
- **选择合适的数据结构**：根据问题需求选择合适的数据结构
- **优化时间复杂度**：通过巧妙的设计，如元素转移的时机，优化操作的时间复杂度
- **代码设计**：注重代码的可读性和可维护性，避免冗余代码
- **边界条件处理**：考虑各种边界情况，如空队列的操作

### 3. 实际应用
- **队列的应用**：广度优先搜索（BFS）、消息队列、任务调度等
- **栈的应用**：深度优先搜索（DFS）、表达式求值、函数调用栈等
- **数据结构设计**：在实际开发中，经常需要根据具体需求设计和组合数据结构

### 4. 学习建议
- 理解栈和队列的基本特性和操作
- 掌握双栈实现队列、双队列实现栈等经典问题
- 学习分析算法的时间复杂度，特别是均摊时间复杂度
- 通过实际项目和练习，加深对数据结构的理解和应用

通过解决这个问题，我们不仅掌握了用栈实现队列的方法，更重要的是学习了如何通过巧妙的设计优化算法的时间复杂度，以及如何编写清晰、高效的代码。这些技能在解决更复杂的算法问题时会非常有帮助。