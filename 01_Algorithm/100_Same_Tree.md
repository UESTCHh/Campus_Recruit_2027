# 100. 相同的树 学习笔记

## 一、 题目分析

### 1.1 题目描述

给你两棵二叉树的根节点 `p` 和 `q`，编写一个函数来检验这两棵树是否相同。

如果两个树在**结构上相同**，并且**节点具有相同的值**，则认为它们是相同的。

### 1.2 题目示例

#### 示例 1：
```
输入：p = [1,2,3], q = [1,2,3]
输出：true
```

**图示：**
```
    p 树              q 树
     1                  1
    / \                / \
   2   3              2   3
```

#### 示例 2：
```
输入：p = [1,2], q = [1,null,2]
输出：false
```

**图示：**
```
    p 树              q 树
     1                  1
    /                    \
   2                      2
```

#### 示例 3：
```
输入：p = [1,2,1], q = [1,1,2]
输出：false
```

**图示：**
```
    p 树              q 树
     1                  1
    / \                / \
   2   1              1   2
```

### 1.3 提示信息

- 两棵树上的节点数目都在范围 `[0, 100]` 内
- `-10^4 <= Node.val <= 10^4`

### 1.4 题目本质分析

**核心判断条件：**

1. **结构相同**：在相同位置上，节点要么都存在，要么都不存在
2. **值相同**：在对应位置上的节点，值必须相等

**关键点：**
- 这是一个**两棵树同步遍历**的问题
- 需要同时访问两棵树的对应节点
- 判断条件的顺序很重要

---

## 二、 解题思路分析

### 2.1 问题核心

判断两棵树是否相同，本质上是一个**同步递归遍历**的问题：

1. 同时访问两棵树的当前节点
2. 判断两个节点是否满足相同的条件
3. 递归（或迭代）处理左右子树

### 2.2 判断条件的顺序（重要！）

判断两棵树相同的条件需要按正确的顺序检查：

```
条件 1：两个节点都为空 → 相同 ✓
条件 2：一个为空，一个不为空 → 不同 ✗
条件 3：两个节点值不同 → 不同 ✗
条件 4：递归比较左子树和右子树
```

**错误的顺序会导致空指针访问！**

### 2.3 递归法思路

#### 2.3.1 递归三要素

**1. 确定递归函数的参数和返回值：**
- 参数：两棵树的当前节点 `p` 和 `q`
- 返回值：bool，表示当前这两棵子树是否相同

**2. 确定终止条件：**
```
if (p == nullptr && q == nullptr) return true;    // 都为空，相同
if (p == nullptr || q == nullptr) return false;   // 一个为空一个不为空，不同
if (p->val != q->val) return false;               // 值不同，不同
```

**3. 确定单层递归的逻辑：**
```
bool leftSame = isSameTree(p->left, q->left);     // 比较左子树
bool rightSame = isSameTree(p->right, q->right);   // 比较右子树
return leftSame && rightSame;                     // 左右都相同才返回 true
```

#### 2.3.2 遍历顺序选择

这道题使用**前序遍历、中序遍历、后序遍历都可以**！

但最自然的是**前序遍历**（中→左→右）：
1. 先比较当前节点（中）
2. 再比较左子树
3. 最后比较右子树

**为什么任意顺序都可以？**
因为我们需要判断**所有**节点都相同，无论访问顺序如何，最终结果都要取 &&。

### 2.4 迭代法思路

#### 2.4.1 使用队列（层序遍历）

思路：使用一个队列，**同时存储两棵树的节点**，成对处理。

```
队列元素：(p的节点, q的节点)

处理流程：
1. 先将 (p根, q根) 加入队列
2. 当队列不为空时：
   a. 取出一对节点 (node1, node2)
   b. 判断这对节点是否相同
   c. 将 (node1->left, node2->left) 加入队列
   d. 将 (node1->right, node2->right) 加入队列
```

#### 2.4.2 使用栈（深度优先遍历）

与队列类似，只是访问顺序不同，但判断逻辑完全相同。

---

## 三、 算法实现代码

### 3.1 二叉树节点定义

```cpp
/**
 * Definition for a binary tree node.
 */
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    
    // 构造函数
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) 
        : val(x), left(left), right(right) {}
};
```

### 3.2 方法一：递归法（前序遍历）

```cpp
class Solution {
public:
    bool isSameTree(TreeNode* p, TreeNode* q) {
        // ==========================================
        // 终止条件1：两个节点都为空
        // ==========================================
        if (p == nullptr && q == nullptr) {
            return true;
        }
        
        // ==========================================
        // 终止条件2：一个为空，一个不为空
        // （注意：前面已经排除了都为空的情况）
        // ==========================================
        if (p == nullptr || q == nullptr) {
            return false;
        }
        
        // ==========================================
        // 终止条件3：值不相等
        // ==========================================
        if (p->val != q->val) {
            return false;
        }
        
        // ==========================================
        // 单层递归逻辑：比较左子树和右子树
        // ==========================================
        bool leftSame = isSameTree(p->left, q->left);   // 比较左子树
        bool rightSame = isSameTree(p->right, q->right); // 比较右子树
        
        // 左右子树都相同才返回 true
        return leftSame && rightSame;
    }
};
```

### 3.3 代码分析（递归法）

#### 3.3.1 实现思路说明

**整体思路：**
- 使用前序遍历的思想：先处理当前节点，再处理左子树，最后处理右子树
- 终止条件判断非常严格，按顺序排除不可能的情况
- 递归逻辑简洁明了

**为什么使用这种实现方式：**
- 递归代码最简洁，最符合树的问题思维
- 终止条件的顺序是经过精心设计的，可以避免空指针访问

#### 3.3.2 核心逻辑讲解

**终止条件的顺序（关键！）：**

```
1. 先检查是否都为空 → 返回 true
2. 再检查是否一个为空 → 返回 false
3. 最后才访问值进行比较 → 安全！
```

**如果顺序乱了会怎样？**
```cpp
// ❌ 错误示例：先访问值，可能空指针！
if (p->val != q->val) { ... }  // 如果 p 或 q 为空，崩溃！

// ✅ 正确：先判空，后访问值
if (p == nullptr && q == nullptr) { ... }  // 安全
if (p == nullptr || q == nullptr) { ... }  // 安全
if (p->val != q->val) { ... }              // 此时 p 和 q 一定都不为空！
```

#### 3.3.3 初学者引导

**写给初学者：**

1. **为什么要有三个 if 条件？**
   - 必须按顺序处理各种情况
   - 保证访问 p->val 时，p 一定不为空

2. **为什么是 (p == nullptr || q == nullptr) 而不是 &&？**
   - 前面已经排除了都为空的情况
   - 此时如果有一个为空，说明另一个一定不为空
   - 结构不同，直接返回 false

3. **最后的 return leftSame && rightSame 是什么意思？**
   - 左子树和右子树都必须相同
   - 用 && 连接，因为需要两者同时满足

### 3.4 方法一的简化写法

```cpp
class Solution {
public:
    bool isSameTree(TreeNode* p, TreeNode* q) {
        // 都为空
        if (!p && !q) return true;
        
        // 一个为空一个不为空，或值不相等
        if (!p || !q || p->val != q->val) return false;
        
        // 递归比较左右子树
        return isSameTree(p->left, q->left) && isSameTree(p->right, q->right);
    }
};
```

**分析：**
- 更简洁，但逻辑与上面完全相同
- `if (!p || !q || p->val != q->val)` 这一行利用了**短路求值**特性
- 前面的条件满足了，后面的条件不会再判断

### 3.5 方法二：迭代法（使用队列，层序遍历）

```cpp
#include <queue>
using namespace std;

class Solution {
public:
    bool isSameTree(TreeNode* p, TreeNode* q) {
        // 创建队列，存储成对的节点
        queue<pair<TreeNode*, TreeNode*>> que;
        
        // 将根节点加入队列
        que.push({p, q});
        
        // ==========================================
        // 循环处理队列
        // ==========================================
        while (!que.empty()) {
            // 取出一对节点
            auto [node1, node2] = que.front();
            que.pop();
            
            // ==========================================
            // 判断条件（与递归法相同）
            // ==========================================
            
            // 情况1：都为空，继续处理下一对
            if (node1 == nullptr && node2 == nullptr) {
                continue;
            }
            
            // 情况2：一个为空一个不为空，返回 false
            if (node1 == nullptr || node2 == nullptr) {
                return false;
            }
            
            // 情况3：值不相等，返回 false
            if (node1->val != node2->val) {
                return false;
            }
            
            // ==========================================
            // 将左右子节点加入队列
            // ==========================================
            que.push({node1->left, node2->left});
            que.push({node1->right, node2->right});
        }
        
        // 队列处理完了都没有返回 false，说明完全相同
        return true;
    }
};
```

### 3.6 代码分析（迭代法）

#### 3.6.1 实现思路说明

**队列法的核心思想：**
- 使用队列进行层序遍历
- 成对存储和处理节点：`(p的节点, q的节点)`
- 判断逻辑与递归法完全相同

**为什么采用这种方式：**
- 避免递归调用栈的开销（虽然对于这道题影响不大）
- 层序遍历在某些情况下更直观

#### 3.6.2 关键技术点解释

**1. pair 的使用：**
```cpp
queue<pair<TreeNode*, TreeNode*>> que;
```
- `pair` 可以把两个指针绑定在一起
- 保证两个指针的同步处理

**2. C++17 的结构化绑定：**
```cpp
auto [node1, node2] = que.front();  // C++17
```
- 方便地从 pair 中取出两个元素
- 旧版本可以用：
```cpp
TreeNode* node1 = que.front().first;
TreeNode* node2 = que.front().second;
```

**3. continue 的使用：**
```cpp
if (node1 == nullptr && node2 == nullptr) {
    continue;  // 都为空，不需要处理子节点
}
```
- 跳过这一对，继续处理队列中的下一对

### 3.7 方法三：迭代法（使用栈，深度优先遍历）

```cpp
#include <stack>
using namespace std;

class Solution {
public:
    bool isSameTree(TreeNode* p, TreeNode* q) {
        // 创建栈，存储成对的节点
        stack<pair<TreeNode*, TreeNode*>> stk;
        
        // 将根节点入栈
        stk.push({p, q});
        
        while (!stk.empty()) {
            auto [node1, node2] = stk.top();
            stk.pop();
            
            // 判断逻辑与队列法完全相同
            if (node1 == nullptr && node2 == nullptr) {
                continue;
            }
            if (node1 == nullptr || node2 == nullptr) {
                return false;
            }
            if (node1->val != node2->val) {
                return false;
            }
            
            // 注意：栈是后入先出，所以先压右再压左
            stk.push({node1->right, node2->right});
            stk.push({node1->left, node2->left});
        }
        
        return true;
    }
};
```

**栈方法的小细节：**
- 栈是**后入先出**的
- 所以入栈顺序是**先右后左**
- 这样弹出时才会是**先左后右**
- 但对这道题来说，顺序不影响结果，因为判断逻辑相同

---

## 四、 时间与空间复杂度分析

### 4.1 递归法复杂度分析

#### 4.1.1 时间复杂度

**O(n)，其中 n 是树的节点数。**

**分析：**
- 每个节点被访问且仅被访问一次
- 没有重复计算

**最好情况：O(1)**
- 根节点就不相同，直接返回 false

**最坏情况：O(n)**
- 两个树完全相同，需要遍历所有节点

#### 4.1.2 空间复杂度

**O(h)，其中 h 是树的高度。**

**分析：**
- 递归调用会占用系统栈
- 栈的深度等于树的高度

**不同情况下的 h：**
- **完全平衡树：** h = logn
- **链表式的树：** h = n

### 4.2 迭代法复杂度分析

#### 4.2.1 时间复杂度

**O(n)，与递归法相同。**

每个节点入队且仅入队一次。

#### 4.2.2 空间复杂度

**O(n)，最坏情况下队列大小为 n。**

**分析：**
- 最坏情况：完全平衡树，队列大小为 n/2 ≈ n
- 最好情况：链表式的树，队列大小为 1

**队列 vs 栈：**
- 空间复杂度相同
- 只是遍历顺序不同

### 4.3 方法对比总结

| 方法 | 时间复杂度 | 空间复杂度 | 优点 | 缺点 |
|------|-----------|-----------|------|------|
| 递归法 | O(n) | O(h) | 代码简洁，思路清晰 | 有栈溢出风险（但本题 n ≤ 100，不会） |
| 队列法 | O(n) | O(n) | 迭代实现，避免递归 | 代码稍长 |
| 栈法 | O(n) | O(n) | 迭代实现 | 代码稍长 |

---

## 五、 测试用例验证

### 5.1 测试用例设计

#### 测试用例 1：完全相同的树
```
输入：
p = [1,2,3]
q = [1,2,3]
输出：true
```

**执行过程：**
```
1. 比较根节点 1 和 1 → 相同 ✓
2. 比较左子树 2 和 2 → 相同 ✓
3. 比较右子树 3 和 3 → 相同 ✓
4. 返回 true
```

#### 测试用例 2：结构不同
```
输入：
p = [1,2]
q = [1,null,2]
输出：false
```

**执行过程：**
```
1. 比较根节点 1 和 1 → 相同 ✓
2. 比较左子树：
   p 的左子树是 2（非空）
   q 的左子树是 null（空）
   → 结构不同，返回 false
```

#### 测试用例 3：结构相同但值不同
```
输入：
p = [1,2,1]
q = [1,1,2]
输出：false
```

**执行过程：**
```
1. 比较根节点 1 和 1 → 相同 ✓
2. 比较左子树 2 和 1 → 值不同，返回 false
```

#### 测试用例 4：空树
```
输入：
p = []
q = []
输出：true
```

#### 测试用例 5：一个空一个非空
```
输入：
p = [1]
q = []
输出：false
```

### 5.2 测试用例代码

```cpp
#include <iostream>
using namespace std;

// 测试函数
void testCase(TreeNode* p, TreeNode* q, bool expected, const string& caseName) {
    Solution sol;
    bool result = sol.isSameTree(p, q);
    cout << "测试用例 " << caseName << ": ";
    if (result == expected) {
        cout << "✅ 通过" << endl;
    } else {
        cout << "❌ 失败" << endl;
        cout << "  期望: " << (expected ? "true" : "false") << endl;
        cout << "  实际: " << (result ? "true" : "false") << endl;
    }
}

// 辅助函数：创建树
TreeNode* createTree(const vector<int>& vals, int index) {
    if (index >= vals.size() || vals[index] == -1) {
        return nullptr;
    }
    TreeNode* root = new TreeNode(vals[index]);
    root->left = createTree(vals, 2 * index + 1);
    root->right = createTree(vals, 2 * index + 2);
    return root;
}

// 辅助函数：删除树（防止内存泄漏）
void deleteTree(TreeNode* root) {
    if (root == nullptr) {
        return;
    }
    deleteTree(root->left);
    deleteTree(root->right);
    delete root;
}

int main() {
    cout << "========== 100. 相同的树 测试 ==========" << endl;
    
    // 测试用例 1
    {
        vector<int> vals1 = {1, 2, 3};
        vector<int> vals2 = {1, 2, 3};
        TreeNode* p = createTree(vals1, 0);
        TreeNode* q = createTree(vals2, 0);
        testCase(p, q, true, "1. 完全相同的树");
        deleteTree(p);
        deleteTree(q);
    }
    
    // 测试用例 2
    {
        vector<int> vals1 = {1, 2, -1};
        vector<int> vals2 = {1, -1, 2};
        TreeNode* p = createTree(vals1, 0);
        TreeNode* q = createTree(vals2, 0);
        testCase(p, q, false, "2. 结构不同");
        deleteTree(p);
        deleteTree(q);
    }
    
    // 测试用例 3
    {
        vector<int> vals1 = {1, 2, 1};
        vector<int> vals2 = {1, 1, 2};
        TreeNode* p = createTree(vals1, 0);
        TreeNode* q = createTree(vals2, 0);
        testCase(p, q, false, "3. 值不同");
        deleteTree(p);
        deleteTree(q);
    }
    
    // 测试用例 4
    {
        testCase(nullptr, nullptr, true, "4. 两棵空树");
    }
    
    // 测试用例 5
    {
        TreeNode* p = new TreeNode(1);
        testCase(p, nullptr, false, "5. 一个空一个非空");
        deleteTree(p);
    }
    
    cout << "====================================" << endl;
    
    return 0;
}
```

---

## 六、 常见错误与注意事项

### 6.1 常见错误

#### 错误 1：判断条件的顺序错误（最重要！）

```cpp
// ❌ 错误示例：先访问值，可能导致空指针！
bool isSameTree(TreeNode* p, TreeNode* q) {
    if (p->val != q->val) {  // 崩溃！如果 p 或 q 为空
        return false;
    }
    // ...
}
```

**正确做法：**
```cpp
// ✅ 正确：先判空，再访问值
bool isSameTree(TreeNode* p, TreeNode* q) {
    if (p == nullptr && q == nullptr) {
        return true;
    }
    if (p == nullptr || q == nullptr) {
        return false;
    }
    if (p->val != q->val) {  // 此时 p 和 q 一定都不为空
        return false;
    }
    // ...
}
```

#### 错误 2：忘记处理空指针

```cpp
// ❌ 错误示例
bool isSameTree(TreeNode* p, TreeNode* q) {
    return (p->val == q->val) && 
           isSameTree(p->left, q->left) && 
           isSameTree(p->right, q->right);
}
```
- 完全没有处理空指针的情况
- 如果有空节点，一定会崩溃

#### 错误 3：迭代法中忘记将空指针加入队列

```cpp
// ❌ 错误示例
if (node1->left && node2->left) {  // 只在非空时才加入队列
    que.push({node1->left, node2->left});
}
```
- 空指针也应该加入队列，用于检查结构是否相同
- 因为我们需要判断：是都为空，还是一个空一个非空

**正确做法：**
```cpp
// ✅ 正确：空指针也要加入队列
que.push({node1->left, node2->left});  // 不管是否为空
que.push({node1->right, node2->right});
```

### 6.2 注意事项

#### 6.2.1 短路求值的利用

```cpp
if (!p || !q || p->val != q->val) {
    return false;
}
```

这段代码利用了 C++ 的**短路求值**特性：
1. 如果 `!p` 为 true，后面的条件不再判断
2. 所以 `p->val` 只在 `p` 和 `q` 都不为空时才会执行

#### 6.2.2 迭代法中 continue 的使用

```cpp
if (node1 == nullptr && node2 == nullptr) {
    continue;  // 都为空，不需要处理子节点
}
```

这里不能直接 return，因为后面可能还有其他节点需要判断。

#### 6.2.3 不要修改原树

在这道题中，我们只需要**遍历判断**，不需要修改树的结构。

---

## 七、 与其他相关题目的对比分析

### 7.1 对比 101. 对称二叉树

#### 7.1.1 题目对比

| 题目 | 判断对象 | 比较关系 |
|------|---------|---------|
| 100. 相同的树 | 两棵树 | p的左 ↔ q的左，p的右 ↔ q的右 |
| 101. 对称二叉树 | 一棵树的左右子树 | p的左 ↔ q的右，p的右 ↔ q的左 |

#### 7.1.2 代码对比

**100. 相同的树：**
```cpp
bool isSameTree(TreeNode* p, TreeNode* q) {
    // ...
    bool leftSame = isSameTree(p->left, q->left);
    bool rightSame = isSameTree(p->right, q->right);
    return leftSame && rightSame;
}
```

**101. 对称二叉树：**
```cpp
bool compare(TreeNode* p, TreeNode* q) {
    // ...
    bool outside = compare(p->left, q->right);  // 注意：左↔右
    bool inside = compare(p->right, q->left);   // 注意：右↔左
    return outside && inside;
}
```

**关键区别：**
- 相同的树：`左↔左，右↔右`
- 对称二叉树：`左↔右，右↔左`

#### 7.1.3 对比总结

这两道题**核心思路完全相同**！都是：
1. 同步遍历
2. 比较对应位置的节点
3. 递归处理子树

唯一的区别就是**比较的配对关系不同**。

### 7.2 对比 104. 二叉树的最大深度

这道题更简单，只需要处理一棵树，但可以帮助理解递归思想。

### 7.3 对比 101、100、104 的递进关系

```
104. 二叉树的最大深度
    ↓（基础：单树遍历）
100. 相同的树
    ↓（进阶：双树同步遍历）
101. 对称二叉树
    ↓（灵活：比较的配对关系不同）
```

---

## 八、 面试相关内容

### 8.1 面试常见问题

#### Q1: 判断相同二叉树，你有几种方法？

**答案：**

**至少有三种方法：**

1. **递归法（前/中/后序遍历都可以）**
   - 最简洁的方法
   - 时间 O(n)，空间 O(h)

2. **迭代法（队列-层序遍历）**
   - 成对存储节点
   - 时间 O(n)，空间 O(n)

3. **迭代法（栈-深度优先遍历）**
   - 与队列类似，只是遍历顺序不同
   - 时间 O(n)，空间 O(n)

#### Q2: 判断相同二叉树的时间和空间复杂度？

**答案：**

| 方法 | 时间复杂度 | 空间复杂度 |
|------|-----------|-----------|
| 递归法 | O(n) | O(h) |
| 队列法 | O(n) | O(n) |
| 栈法 | O(n) | O(n) |

**详细说明：**
- n：节点数
- h：树的高度（完全平衡树 h=logn，链表 h=n）

#### Q3: 判断相同二叉树，递归终止条件的顺序为什么重要？

**答案：**

**为了避免空指针访问！**

**正确的顺序：**
1. 先判断是否都为空
2. 再判断是否一个为空一个不为空
3. 最后才访问值进行比较

**原因：**
- 如果不先判空，直接访问 p->val，当 p 为空时会崩溃
- 正确的顺序保证了：在执行 p->val 时，p 和 q 一定都不为空

#### Q4: 如果不使用递归，如何用迭代实现？

**答案：**

使用队列或栈，成对存储和处理节点。

**队列法：**
```cpp
queue<pair<TreeNode*, TreeNode*>> que;
que.push({p, q});
while (!que.empty()) {
    auto [node1, node2] = que.front();
    que.pop();
    // 判断逻辑
    que.push({node1->left, node2->left});
    que.push({node1->right, node2->right});
}
```

#### Q5: 前序、中序、后序遍历都可以吗？

**答案：**

**都可以！**

**原因：**
- 因为我们需要比较所有节点
- 无论什么顺序，最终结果都是所有比较的 &&
- 只要能保证每个节点都比较到就行

**但前序遍历最自然：**
先比较当前节点，一旦不同立即返回，可能提前终止。

### 8.2 扩展问题

#### 扩展 1：如何判断一棵树是另一棵树的子树？

**LeetCode 572. 另一棵树的子树**

**思路：**
- 遍历第一棵树的每一个节点
- 判断以该节点为根的子树是否与第二棵树相同
- 就可以复用这道题的 isSameTree 函数

```cpp
bool isSubtree(TreeNode* root, TreeNode* subRoot) {
    if (root == nullptr && subRoot == nullptr) {
        return true;
    }
    if (root == nullptr || subRoot == nullptr) {
        return false;
    }
    if (isSameTree(root, subRoot)) {  // 复用本题函数
        return true;
    }
    return isSubtree(root->left, subRoot) || isSubtree(root->right, subRoot);
}
```

#### 扩展 2：如何判断两棵树是同构的？

**同构：可以通过交换左右子节点使两棵树相同。**

**思路：**
```cpp
bool isIsomorphic(TreeNode* p, TreeNode* q) {
    if (p == nullptr && q == nullptr) {
        return true;
    }
    if (p == nullptr || q == nullptr) {
        return false;
    }
    if (p->val != q->val) {
        return false;
    }
    // 两种可能：不交换 或 交换
    return (isIsomorphic(p->left, q->left) && isIsomorphic(p->right, q->right)) ||
           (isIsomorphic(p->left, q->right) && isIsomorphic(p->right, q->left));
}
```

#### 扩展 3：如何实现树的深拷贝？

**LeetCode 133. 克隆图 的思路可以借鉴到树。**

### 8.3 面试技巧

#### 技巧 1：先说思路，再写代码

面试时，不要一上来就写代码，先说思路：

1. **这道题可以用递归或迭代**
2. **判断条件有三个，顺序很重要**
3. **递归的终止条件是...，单层逻辑是...**

#### 技巧 2：先写递归终止条件

写递归代码时，先把终止条件写好，最后再写递归逻辑。

#### 技巧 3：注意空指针的处理

面试官非常看重这一点！

---

## 九、 知识点总结

### 9.1 核心知识点

| 知识点 | 内容 |
|--------|------|
| 双树同步遍历 | 同时访问两棵树的对应节点 |
| 判断条件的顺序 | 先判空，再访问值 |
| 递归三要素 | 参数和返回值、终止条件、单层逻辑 |
| 迭代实现 | 队列或栈，成对处理节点 |
| 时间复杂度 | O(n)，每个节点访问一次 |
| 空间复杂度 | 递归 O(h)，迭代 O(n) |

### 9.2 解题技巧

1. **递归法最简洁**，面试时优先考虑
2. **判断条件顺序不能错**，否则会有空指针访问
3. **迭代法使用 pair**，保证节点同步处理
4. **空指针也要加入队列/栈**，用于判断结构是否相同

### 9.3 相关题目

- **100. 相同的树**（本题）
- **101. 对称二叉树**（姊妹题）
- **104. 二叉树的最大深度**（基础）
- **572. 另一棵树的子树**（扩展）

---

## 十、 个人总结与反思

### 10.1 学习收获

1. **深入理解了双树同步遍历**
   - 不是只遍历一棵树
   - 而是同时处理两棵树的对应节点

2. **理解了判断条件顺序的重要性**
   - 这是一个很容易犯的错误
   - 但只要遵循"先判空，后访问值"的原则，就可以避免

3. **掌握了递归到迭代的转换**
   - 理解了如何用队列或栈来模拟递归过程
   - 成对处理节点的思想

4. **对比了与对称二叉树的关系**
   - 核心思路完全相同
   - 唯一区别是比较的配对关系不同

### 10.2 解题过程中的思考

**第一次做这道题时，可能的思考过程：**

```
1. 怎么比较两棵树？
   → 同时遍历吧！

2. 同时遍历，需要什么判断？
   → 先看节点是否都存在
   → 再看值是否相等
   → 再比较子树

3. 判断顺序很重要！
   → 如果先访问值，可能会空指针
   → 应该先判空

4. 子树怎么比较？
   → 递归啊！
   → 左子树比较左子树
   → 右子树比较右子树
```

### 10.3 常见错误回顾

1. **判断条件顺序错误** → 空指针访问
2. **迭代法中忘记加入空指针** → 结构检查不完整
3. **递归法中忘记返回值** → 逻辑错误

### 10.4 代码优化思路

这道题的时间和空间复杂度已经是最优的了，很难再优化。

但可以从代码简洁性上考虑：

```cpp
// 简洁版本
bool isSameTree(TreeNode* p, TreeNode* q) {
    if (!p && !q) return true;
    if (!p || !q || p->val != q->val) return false;
    return isSameTree(p->left, q->left) && isSameTree(p->right, q->right);
}
```

**优点：**
- 代码非常简洁
- 利用短路求值

**缺点：**
- 对初学者不够友好
- 逻辑不够清晰

### 10.5 后续学习建议

1. **先做 100. 相同的树，再做 101. 对称二叉树**
   - 两道题联系紧密
   - 理解相同，对称就简单了

2. **对比递归和迭代实现**
   - 理解递归的本质
   - 掌握如何用栈模拟递归

3. **尝试其他遍历顺序**
   - 前序可以，中序、后序呢？
   - 其实都可以，但前序最自然

4. **思考扩展问题**
   - 如何判断子树？
   - 如何判断同构？

---

## 十一、 完整参考代码（带测试）

```cpp
#include <iostream>
#include <queue>
#include <stack>
#include <vector>
using namespace std;

/**
 * Definition for a binary tree node.
 */
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

// ==========================================
// 方法一：递归法
// ==========================================
class Solution1 {
public:
    bool isSameTree(TreeNode* p, TreeNode* q) {
        // 都为空
        if (p == nullptr && q == nullptr) {
            return true;
        }
        // 一个为空一个不为空
        if (p == nullptr || q == nullptr) {
            return false;
        }
        // 值不相等
        if (p->val != q->val) {
            return false;
        }
        // 递归比较左右子树
        return isSameTree(p->left, q->left) && isSameTree(p->right, q->right);
    }
};

// ==========================================
// 方法二：迭代法（队列）
// ==========================================
class Solution2 {
public:
    bool isSameTree(TreeNode* p, TreeNode* q) {
        queue<pair<TreeNode*, TreeNode*>> que;
        que.push({p, q});
        
        while (!que.empty()) {
            auto [node1, node2] = que.front();
            que.pop();
            
            if (node1 == nullptr && node2 == nullptr) {
                continue;
            }
            if (node1 == nullptr || node2 == nullptr) {
                return false;
            }
            if (node1->val != node2->val) {
                return false;
            }
            
            que.push({node1->left, node2->left});
            que.push({node1->right, node2->right});
        }
        
        return true;
    }
};

// ==========================================
// 方法三：迭代法（栈）
// ==========================================
class Solution3 {
public:
    bool isSameTree(TreeNode* p, TreeNode* q) {
        stack<pair<TreeNode*, TreeNode*>> stk;
        stk.push({p, q});
        
        while (!stk.empty()) {
            auto [node1, node2] = stk.top();
            stk.pop();
            
            if (node1 == nullptr && node2 == nullptr) {
                continue;
            }
            if (node1 == nullptr || node2 == nullptr) {
                return false;
            }
            if (node1->val != node2->val) {
                return false;
            }
            
            stk.push({node1->right, node2->right});
            stk.push({node1->left, node2->left});
        }
        
        return true;
    }
};

// ==========================================
// 测试代码
// ==========================================

// 辅助函数：创建树（-1 表示空）
TreeNode* createTree(const vector<int>& vals, int index) {
    if (index >= vals.size() || vals[index] == -1) {
        return nullptr;
    }
    TreeNode* root = new TreeNode(vals[index]);
    root->left = createTree(vals, 2 * index + 1);
    root->right = createTree(vals, 2 * index + 2);
    return root;
}

// 辅助函数：删除树
void deleteTree(TreeNode* root) {
    if (root == nullptr) {
        return;
    }
    deleteTree(root->left);
    deleteTree(root->right);
    delete root;
}

// 测试函数
void testSolution(TreeNode* p, TreeNode* q, bool expected, const string& caseName, int method) {
    bool result;
    switch (method) {
        case 1:
            result = Solution1().isSameTree(p, q);
            break;
        case 2:
            result = Solution2().isSameTree(p, q);
            break;
        case 3:
            result = Solution3().isSameTree(p, q);
            break;
    }
    
    cout << "方法" << method << " " << caseName << ": ";
    if (result == expected) {
        cout << "✅ 通过" << endl;
    } else {
        cout << "❌ 失败" << endl;
    }
}

int main() {
    cout << "========== 100. 相同的树 完整测试 ==========" << endl;
    
    // 测试用例
    struct TestCase {
        vector<int> vals1;
        vector<int> vals2;
        bool expected;
        string name;
    };
    
    vector<TestCase> cases = {
        {{1, 2, 3}, {1, 2, 3}, true, "完全相同"},
        {{1, 2, -1}, {1, -1, 2}, false, "结构不同"},
        {{1, 2, 1}, {1, 1, 2}, false, "值不同"},
        {{}, {}, true, "两空树"},
        {{1}, {}, false, "一空一非空"},
    };
    
    // 测试三种方法
    for (int method = 1; method <= 3; method++) {
        cout << "\n-------- 方法" << method << " --------" << endl;
        for (auto& cs : cases) {
            TreeNode* p = createTree(cs.vals1, 0);
            TreeNode* q = createTree(cs.vals2, 0);
            testSolution(p, q, cs.expected, cs.name, method);
            deleteTree(p);
            deleteTree(q);
        }
    }
    
    cout << "\n========================================" << endl;
    
    return 0;
}
```

---

通过这道题，我们深入理解了**两棵树同步遍历**的思想，掌握了**递归和迭代两种实现方法**，理解了**判断条件顺序的重要性**，并且对比了与**对称二叉树**的关系。这些知识对于解决树的相关问题非常重要！
