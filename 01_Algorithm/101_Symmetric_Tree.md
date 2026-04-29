# 101. 对称二叉树学习笔记

## 一、 题目分析

### 1.1 题目描述

给你一棵二叉树的根节点 `root` ，检查它是否轴对称。

### 1.2 题目链接

- [LeetCode 101. 对称二叉树](https://leetcode.cn/problems/symmetric-tree/description/)
- [代码随想录题解](https://www.programmercarl.com/0101.%E5%AF%B9%E7%A7%B0%E4%BA%8C%E5%8F%89%E6%A0%91.html)

### 1.3 示例

**示例 1：**

```
输入：root = [1,2,2,3,4,4,3]
输出：true
```

**示例 2：**

```
输入：root = [1,2,2,null,3,null,3]
输出：false
```

### 1.4 提示

- 树中节点数目范围在 `[1, 1000]` 内
- `-100 <= Node.val <= 100`

### 1.5 题目本质

对称二叉树的本质是：**判断根节点的左子树和右子树是否是相互翻转的**。

换句话说，我们需要比较的不是一个节点的左右孩子，而是两棵树（左子树和右子树），看它们是否镜像对称。

### 1.6 题目进阶

你可以运用递归和迭代两种方法解决这个问题吗？

## 二、 解题思路分析

### 2.1 核心思路

判断对称二叉树的核心思路是：**比较根节点的左子树和右子树是否是相互翻转的**。

在比较过程中，我们需要注意以下几点：

| 比较对象 | 说明 |
|---------|------|
| 左子树的左节点 vs 右子树的右节点 | 外侧比较 |
| 左子树的右节点 vs 右子树的左节点 | 内侧比较 |
| 两个节点的值是否相等 | 值比较 |
| 两个节点的为空情况 | 空节点处理 |

### 2.2 遍历顺序选择

对于对称二叉树，我们需要使用**后序遍历**的方式，原因如下：

- 我们需要先比较两个子树的外侧和内侧节点
- 然后根据子树的比较结果来判断当前节点是否对称
- 这符合后序遍历"左右中"的思想（左子树：左-右-中，右子树：右-左-中）

虽然不是严格的后序遍历（因为我们同时在遍历两棵树），但可以理解为后序遍历的思想。

### 2.3 递归法思路

#### 2.3.1 递归三部曲

**1. 确定递归函数的参数和返回值**
- 参数：需要比较的两个节点（左子树节点和右子树节点）
- 返回值：bool类型，表示是否对称

```cpp
bool compare(TreeNode* left, TreeNode* right)
```

**2. 确定终止条件**

终止条件需要处理以下几种情况：

| 左节点情况 | 右节点情况 | 结果 |
|------------|------------|------|
| 空 | 空 | true（对称） |
| 空 | 非空 | false（不对称） |
| 非空 | 空 | false（不对称） |
| 非空 | 非空 | 需要进一步比较值 |

**3. 确定单层递归逻辑**

在单层递归逻辑中：

- 比较外侧节点：左子树的左节点 vs 右子树的右节点
- 比较内侧节点：左子树的右节点 vs 右子树的左节点
- 如果外侧和内侧都对称，那么当前节点对称

### 2.4 迭代法思路

迭代法可以使用**队列**（层序遍历）或**栈**（深度优先遍历）来实现，核心思想是：

- 将需要比较的两个节点同时放入容器
- 每次取出两个节点进行比较
- 如果相等，则继续比较它们的子节点（注意放入顺序）

使用队列的优势：

- 按层处理节点，逻辑清晰
- 容易理解和实现

使用栈的优势：

- 模拟递归过程
- 空间复杂度可能更低（取决于树的高度）

## 三、 算法实现

### 3.1 递归法（后序遍历思想）

```cpp
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    // 递归函数：比较两个子树是否对称
    bool compare(TreeNode* left, TreeNode* right) {
        // 首先排除空节点的情况
        if (left == nullptr && right == nullptr) {
            // 两个节点都为空，对称
            return true;
        }
        else if (left == nullptr && right != nullptr) {
            // 左空右不空，不对称
            return false;
        }
        else if (left != nullptr && right == nullptr) {
            // 左不空右空，不对称
            return false;
        }
        // 排除了空节点，再排除数值不相同的情况
        else if (left->val != right->val) {
            // 节点值不同，不对称
            return false;
        }

        // 此时就是：左右节点都不为空，且数值相同的情况
        // 此时才做递归，做下一层的判断
        bool outside = compare(left->left, right->right);  // 左子树：左、右子树：右（外侧比较）
        bool inside = compare(left->right, right->left);   // 左子树：右、右子树：左（内侧比较）
        bool isSame = outside && inside;                   // 左子树：中、右子树：中（逻辑处理）
        return isSame;
    }

    bool isSymmetric(TreeNode* root) {
        // 边界条件：根节点为空，直接返回true
        if (root == nullptr) {
            return true;
        }
        // 比较根节点的左右子树是否对称
        return compare(root->left, root->right);
    }
};
```

**代码详细分析：**

1. **终止条件处理（关键）**
   ```cpp
   if (left == nullptr && right == nullptr) return true;
   else if (left == nullptr && right != nullptr) return false;
   else if (left != nullptr && right == nullptr) return false;
   else if (left->val != right->val) return false;
   ```
   - 注意这里使用了多个 `else if` 而不是嵌套的 `if-else`
   - 这样可以清晰地处理所有情况
   - 最后进入递归的一定是"左右节点都不为空且值相等"的情况

2. **单层递归逻辑**
   ```cpp
   bool outside = compare(left->left, right->right);  // 外侧比较
   bool inside = compare(left->right, right->left);   // 内侧比较
   bool isSame = outside && inside;                   // 只有内外都对称才返回true
   return isSame;
   ```
   - 先递归比较外侧节点
   - 再递归比较内侧节点
   - 最后根据两者的结果判断是否对称（这是后序遍历的"中"）

### 3.2 迭代法（队列实现）

```cpp
class Solution {
public:
    bool isSymmetric(TreeNode* root) {
        if (root == nullptr) return true;
        
        // 使用队列进行层序遍历
        queue<TreeNode*> que;
        
        // 将根节点的左右子节点加入队列
        que.push(root->left);
        que.push(root->right);
        
        while (!que.empty()) {
            // 取出两个节点进行比较
            TreeNode* leftNode = que.front();
            que.pop();
            TreeNode* rightNode = que.front();
            que.pop();
            
            // 左右都为空，对称
            if (leftNode == nullptr && rightNode == nullptr) {
                continue;
            }
            
            // 左右有一个为空，或者值不相等，不对称
            if (leftNode == nullptr || rightNode == nullptr || leftNode->val != rightNode->val) {
                return false;
            }
            
            // 注意入队顺序！
            // 外侧：左子树的左节点和右子树的右节点
            que.push(leftNode->left);
            que.push(rightNode->right);
            
            // 内侧：左子树的右节点和右子树的左节点
            que.push(leftNode->right);
            que.push(rightNode->left);
        }
        
        return true;
    }
};
```

**代码详细分析：**

1. **队列初始化**
   ```cpp
   queue<TreeNode*> que;
   que.push(root->left);
   que.push(root->right);
   ```
   - 将根节点的左右子节点同时入队
   - 这是迭代法的开始

2. **比较逻辑**
   ```cpp
   if (leftNode == nullptr && rightNode == nullptr) {
       continue;
   }
   if (leftNode == nullptr || rightNode == nullptr || leftNode->val != rightNode->val) {
       return false;
   }
   ```
   - 两个都为空，继续比较下一对
   - 一个为空一个不为空，或者值不相等，直接返回false

3. **入队顺序（关键）**
   ```cpp
   que.push(leftNode->left);   // 左子树的左
   que.push(rightNode->right); // 右子树的右（外侧）
   que.push(leftNode->right);  // 左子树的右
   que.push(rightNode->left);  // 右子树的左（内侧）
   ```
   - 外侧节点先入队，内侧节点后入队
   - 每次入队两个节点，确保成对取出

### 3.3 迭代法（栈实现）

```cpp
class Solution {
public:
    bool isSymmetric(TreeNode* root) {
        if (root == nullptr) return true;
        
        // 使用栈进行深度优先遍历
        stack<TreeNode*> st;
        
        // 将根节点的左右子节点入栈
        st.push(root->left);
        st.push(root->right);
        
        while (!st.empty()) {
            // 取出两个节点进行比较
            TreeNode* rightNode = st.top();
            st.pop();
            TreeNode* leftNode = st.top();
            st.pop();
            
            // 左右都为空，对称
            if (leftNode == nullptr && rightNode == nullptr) {
                continue;
            }
            
            // 左右有一个为空，或者值不相等，不对称
            if (leftNode == nullptr || rightNode == nullptr || leftNode->val != rightNode->val) {
                return false;
            }
            
            // 注意入栈顺序（与队列相反）！
            // 内侧：左子树的右节点和右子树的左节点
            st.push(leftNode->right);
            st.push(rightNode->left);
            
            // 外侧：左子树的左节点和右子树的右节点
            st.push(leftNode->left);
            st.push(rightNode->right);
        }
        
        return true;
    }
};
```

**代码详细分析：**

1. **栈的使用**
   - 栈的入栈和出栈顺序与队列相反
   - 需要注意入栈顺序，确保成对取出

2. **入栈顺序（关键）**
   ```cpp
   st.push(leftNode->right);  // 内侧先入栈
   st.push(rightNode->left);
   st.push(leftNode->left);   // 外侧后入栈
   st.push(rightNode->right);
   ```
   - 因为栈是后进先出，所以内侧先入栈，外侧后入栈
   - 这样出栈时，外侧先出，内侧后出

## 四、 时间与空间复杂度分析

### 4.1 递归法复杂度分析

**时间复杂度：O(n)**

- n 是二叉树的节点数
- 每个节点都被访问一次
- 所以时间复杂度为 O(n)

**空间复杂度：O(n)**

- 空间复杂度主要取决于递归调用的栈深度
- 对于平衡二叉树，栈深度为 O(log n)
- 对于斜树（退化为链表），栈深度为 O(n)
- 所以最坏情况下空间复杂度为 O(n)

### 4.2 迭代法（队列）复杂度分析

**时间复杂度：O(n)**

- 每个节点被访问一次
- 队列操作是 O(1) 的
- 所以时间复杂度为 O(n)

**空间复杂度：O(n)**

- 队列最多存储一层的节点
- 最坏情况下（完全二叉树），队列存储最后一层的所有节点
- 最后一层最多有 n/2 个节点，所以空间复杂度为 O(n)

### 4.3 迭代法（栈）复杂度分析

**时间复杂度：O(n)**

- 与队列方法相同，每个节点访问一次

**空间复杂度：O(n)**

- 与递归法类似，栈的最大深度为树的高度
- 最坏情况下空间复杂度为 O(n)

### 4.4 三种解法复杂度对比

| 解法 | 时间复杂度 | 空间复杂度 | 备注 |
|------|-----------|-----------|------|
| 递归法 | O(n) | O(n) | 最坏情况是斜树 |
| 队列迭代 | O(n) | O(n) | 最坏情况是完全二叉树最后一层 |
| 栈迭代 | O(n) | O(n) | 最坏情况是斜树 |

## 五、 测试用例验证

### 5.1 测试用例 1：对称的完全二叉树

**输入**：`root = [1,2,2,3,4,4,3]`

**预期输出**：`true`

**执行过程**：

```
原始树：
      1
    /   
   2     2
  / \   / \
 3   4 4   3

比较过程：
compare(2, 2):
  compare(3, 3): true
  compare(4, 4): true
  true && true = true

结果：true
```

### 5.2 测试用例 2：不对称的树

**输入**：`root = [1,2,2,null,3,null,3]`

**预期输出**：`false`

**执行过程**：

```
原始树：
      1
    /   
   2     2
    \     \
     3     3

比较过程：
compare(2, 2):
  compare(null, null): true
  compare(3, 3): true
  等等，不对！重新看：
  compare(2->left, 2->right) = compare(null, 3): false
  所以直接返回false

结果：false
```

### 5.3 测试用例 3：单节点树

**输入**：`root = [1]`

**预期输出**：`true`

**执行过程**：

```
原始树：
      1

比较过程：
根节点为空，直接返回true

结果：true
```

### 5.4 测试用例 4：空树

**输入**：`root = []`

**预期输出**：`true`

**执行过程**：

```
原始树：空

比较过程：
根节点为空，直接返回true

结果：true
```

### 5.5 测试用例 5：节点值不同的对称结构树

**输入**：`root = [1,2,2,3,4,5,3]`

**预期输出**：`false`

**执行过程**：

```
原始树：
      1
    /   
   2     2
  / \   / \
 3   4 5   3

比较过程：
compare(2, 2):
  compare(3, 5): false（值不相等）
  所以直接返回false

结果：false
```

## 六、 常见错误与注意事项

### 6.1 终止条件处理错误

**错误示例：**

```cpp
// 错误的终止条件处理
if (left == nullptr && right == nullptr) return true;
if (left->val != right->val) return false;
```

**问题分析：**
- 没有先处理"一个为空一个不为空"的情况
- 会导致空指针访问异常

**正确处理：**

```cpp
// 正确的终止条件处理
if (left == nullptr && right == nullptr) return true;
else if (left == nullptr && right != nullptr) return false;
else if (left != nullptr && right == nullptr) return false;
else if (left->val != right->val) return false;
```

### 6.2 比较顺序错误

**错误示例：**

```cpp
// 错误的比较顺序
bool inside = compare(left->left, right->left);   // 内侧：左左 vs 右左（错误）
bool outside = compare(left->right, right->right); // 外侧：左右 vs 右右（错误）
```

**问题分析：**
- 应该比较：左左 vs 右右（外侧）
- 应该比较：左右 vs 右左（内侧）
- 顺序完全反了

**正确做法：**

```cpp
// 正确的比较顺序
bool outside = compare(left->left, right->right);  // 外侧：左左 vs 右右
bool inside = compare(left->right, right->left);   // 内侧：左右 vs 右左
```

### 6.3 迭代法入队顺序错误

**错误示例（队列）：**

```cpp
// 错误的入队顺序
que.push(leftNode->left);
que.push(leftNode->right);  // 连续入队左子树的两个子节点
que.push(rightNode->right);
que.push(rightNode->left);   // 连续入队右子树的两个子节点
```

**问题分析：**
- 这样会导致外侧和内侧节点没有成对入队
- 比较时会出现错误

**正确做法：**

```cpp
// 正确的入队顺序：成对入队外侧和内侧节点
que.push(leftNode->left);
que.push(rightNode->right); // 外侧成对入队
que.push(leftNode->right);
que.push(rightNode->left);  // 内侧成对入队
```

### 6.4 混淆对称和相同

**错误理解：**
- 认为"对称二叉树"等于"左右子树相同"

**正确理解：**
- 对称二叉树的左右子树应该是"镜像翻转"的关系
- 不是完全相同，而是左右对称

**对比：**
- 相同二叉树：左子树的左 vs 右子树的左，左子树的右 vs 右子树的右
- 对称二叉树：左子树的左 vs 右子树的右，左子树的右 vs 右子树的左

## 七、 面试相关内容

### 7.1 面试常见问题

#### 7.1.1 判断对称二叉树用什么遍历顺序？

**答案：**
- 需要使用后序遍历的思想
- 因为我们需要先处理子树，再处理父节点
- 左子树按照左右中的顺序，右子树按照右左中的顺序

#### 7.1.2 递归法和迭代法的优缺点？

**答案：**

递归法：
- 优点：代码简洁易懂
- 缺点：可能栈溢出，有递归调用开销

迭代法：
- 优点：没有栈溢出风险，空间可优化
- 缺点：代码相对复杂，需要额外的数据结构

#### 7.1.3 如何不使用额外空间？

**答案：**
- 使用Morris遍历的思想
- 但实现会比较复杂
- 面试中可能不会要求

#### 7.1.4 时间复杂度和空间复杂度是多少？

**答案：**
- 时间复杂度：O(n)，每个节点访问一次
- 空间复杂度：O(n)，最坏情况下递归栈或队列/栈的大小为n

### 7.2 扩展问题

#### 7.2.1 判断两棵树是否相同

**题目：** 判断两棵二叉树是否完全相同

**解题思路：**
- 与对称二叉树类似，但比较顺序不同
- 比较：左左 vs 右右，左右 vs 右左 改成 左左 vs 右右，左右 vs 右左（不对，应该是左左 vs 右左，左右 vs 右右）

**代码示例：**

```cpp
bool isSameTree(TreeNode* p, TreeNode* q) {
    if (p == nullptr && q == nullptr) return true;
    if (p == nullptr || q == nullptr) return false;
    if (p->val != q->val) return false;
    
    bool left = isSameTree(p->left, q->left);
    bool right = isSameTree(p->right, q->right);
    return left && right;
}
```

#### 7.2.2 翻转二叉树（第226题）

**题目：** 翻转二叉树

**解题思路：**
- 可以通过翻转其中一棵树，然后比较是否相同来判断对称
- 但直接判断对称更高效

**与本题的关系：**
- 对称二叉树的左右子树应该是相互翻转的
- 翻转二叉树是本题的基础操作

#### 7.2.3 镜像二叉树的构造

**题目：** 构造一棵给定二叉树的镜像

**解题思路：**
- 递归交换每个节点的左右子树
- 与第226题翻转二叉树相同

### 7.3 面试技巧

#### 7.3.1 如何在面试中快速分析问题？

**步骤：**
1. 理解题意：明确对称的定义
2. 画图分析：画出对称树的结构，明确比较规则
3. 选择算法：先想递归，再想迭代
4. 处理边界：先考虑空树、单节点等简单情况
5. 写出代码：注意代码风格和注释
6. 测试验证：用几个测试用例验证

#### 7.3.2 如何优化代码？

**优化方向：**
1. **空间优化**：使用迭代法而不是递归法
2. **提前返回**：发现不对称立即返回
3. **代码简化**：合并条件判断

**优化示例：**

```cpp
// 简化版递归代码
bool compare(TreeNode* left, TreeNode* right) {
    if (left == nullptr && right == nullptr) return true;
    if (left == nullptr || right == nullptr) return false;
    if (left->val != right->val) return false;
    return compare(left->left, right->right) && compare(left->right, right->left);
}
```

## 八、 与其他相关题目的对比分析

### 8.1 对称二叉树 vs 相同二叉树

| 特性 | 对称二叉树 | 相同二叉树 |
|------|-----------|-----------|
| 比较对象 | 左子树 vs 翻转后的右子树 | 左子树 vs 右子树 |
| 外侧比较 | 左左 vs 右右 | 左左 vs 右左 |
| 内侧比较 | 左右 vs 右左 | 左右 vs 右右 |
| 遍历顺序 | 左：左右中，右：右左中 | 两棵树都是左右中 |

**代码对比：**

```cpp
// 对称二叉树比较
bool compareSymmetric(TreeNode* left, TreeNode* right) {
    // ...
    bool outside = compare(left->left, right->right);  // 外侧
    bool inside = compare(left->right, right->left);   // 内侧
    return outside && inside;
}

// 相同二叉树比较
bool compareSame(TreeNode* left, TreeNode* right) {
    // ...
    bool leftSame = compare(left->left, right->left);  // 左比较
    bool rightSame = compare(left->right, right->right); // 右比较
    return leftSame && rightSame;
}
```

### 8.2 对称二叉树 vs 翻转二叉树

| 特性 | 对称二叉树 | 翻转二叉树 |
|------|-----------|-----------|
| 目的 | 判断是否对称 | 翻转一棵树 |
| 操作 | 比较两个节点 | 交换一个节点的左右子树 |
| 返回值 | bool | TreeNode* |
| 关系 | 对称二叉树翻转后与原树相同 | 翻转是对称的基础操作 |

**代码对比：**

```cpp
// 翻转二叉树
TreeNode* invertTree(TreeNode* root) {
    if (root == nullptr) return root;
    swap(root->left, root->right);
    invertTree(root->left);
    invertTree(root->right);
    return root;
}

// 判断对称可以通过翻转实现（虽然效率不高）
bool isSymmetricByInvert(TreeNode* root) {
    if (root == nullptr) return true;
    TreeNode* left = root->left;
    TreeNode* right = root->right;
    TreeNode* invertedLeft = invertTree(left);
    return isSameTree(invertedLeft, right);
}
```

### 8.3 对称二叉树 vs 二叉树的右视图

| 特性 | 对称二叉树 | 二叉树右视图 |
|------|-----------|-----------|
| 目的 | 判断对称性 | 找出从右侧能看到的节点 |
| 核心操作 | 成对比较两个节点 | 只记录每层最右侧节点 |
| 数据结构 | 队列（成对存储） | 队列（按层存储） |
| 时间复杂度 | O(n) | O(n) |

**代码对比：**

```cpp
// 二叉树右视图
vector<int> rightSideView(TreeNode* root) {
    vector<int> result;
    if (root == nullptr) return result;
    queue<TreeNode*> que;
    que.push(root);
    while (!que.empty()) {
        int size = que.size();
        for (int i = 0; i < size; i++) {
            TreeNode* node = que.front();
            que.pop();
            if (i == size - 1) {
                result.push_back(node->val);
            }
            if (node->left) que.push(node->left);
            if (node->right) que.push(node->right);
        }
    }
    return result;
}
```

## 九、 知识点总结

### 9.1 核心知识点

1. **后序遍历思想**：先处理子节点，再处理父节点
2. **对称比较逻辑**：左左 vs 右右，左右 vs 右左
3. **终止条件处理**：空节点、值不同等情况
4. **递归与迭代**：两种实现方式的理解和转换
5. **数据结构选择**：队列和栈的使用

### 9.2 解题技巧

1. **画图分析**：先画图理解对称的含义
2. **明确比较规则**：外侧和内侧的比较方式
3. **边界条件优先处理**：空节点情况先处理
4. **迭代法成对处理**：队列或栈中成对存储节点
5. **注意顺序**：入队或入栈的顺序很重要

### 9.3 相关题目

- [226. 翻转二叉树](https://leetcode.cn/problems/invert-binary-tree/)：基础操作
- [100. 相同的树](https://leetcode.cn/problems/same-tree/)：类似的比较逻辑
- [104. 二叉树的最大深度](https://leetcode.cn/problems/maximum-depth-of-binary-tree/)：后序遍历思想
- [199. 二叉树的右视图](https://leetcode.cn/problems/binary-tree-right-side-view/)：层序遍历应用

## 十、 个人总结与反思

### 10.1 学习收获

1. **理解对称的本质**：深刻理解了对称二叉树的判断逻辑，不是简单比较每个节点的左右子树，而是比较两棵子树。

2. **递归思想的深化**：通过本题加深了对后序遍历思想的理解，先处理子问题，再根据子问题的结果解决当前问题。

3. **递归转迭代的技巧**：掌握了如何将递归算法转换为迭代算法，特别是使用队列或栈模拟递归过程。

4. **边界条件的重要性**：认识到边界条件（如空节点）的处理对于写出正确代码的重要性。

5. **多种解法的掌握**：掌握了递归法和两种迭代法（队列和栈），理解了它们之间的联系和区别。

### 10.2 解题过程中的思考

1. **问题分析阶段**：一开始可能会想简单比较每个节点的左右子节点，但很快发现这是错误的，需要比较两棵子树。

2. **遍历顺序选择**：在选择遍历顺序时，思考为什么需要后序遍历，理解了因为需要子问题的结果来解决当前问题。

3. **终止条件设计**：设计终止条件时，需要考虑各种空节点的情况，这是最容易出错的地方。

4. **迭代法入队顺序**：在实现迭代法时，入队顺序是关键，需要特别注意成对存储和取出节点。

### 10.3 常见错误回顾

1. **空指针错误**：没有先处理空节点的情况，导致空指针访问。
2. **比较顺序错误**：外侧和内侧比较顺序弄反了。
3. **入队顺序错误**：迭代法中没有成对入队节点。
4. **混淆对称和相同**：没有理解对称的真正含义。

### 10.4 代码优化思路

1. **简化条件判断**：合并一些条件，使代码更简洁。
2. **提前返回**：发现不对称立即返回，提高效率。
3. **选择合适的实现方式**：根据实际情况选择递归或迭代。

### 10.5 后续学习建议

1. **更多树的题目练习**：树是面试高频考点，需要多练习。
2. **深入理解递归与迭代**：掌握它们之间的转换方法。
3. **学习Morris遍历**：了解如何不使用额外空间遍历树。
4. **总结树题目的共性**：找出树题目的常见模式和技巧。

### 10.6 面试准备建议

1. **熟练写出三种解法**：递归、队列迭代、栈迭代。
2. **能讲解算法思路**：清晰解释为什么这样设计算法。
3. **分析复杂度**：时间复杂度和空间复杂度的分析。
4. **对比相关题目**：对称二叉树、相同二叉树、翻转二叉树的关系。
5. **快速写代码**：在面试中快速、准确地写出代码。

---

通过本次学习，我对对称二叉树问题有了全面的理解，掌握了递归和迭代两种实现方法，深入分析了时间复杂度和空间复杂度，并且与相关题目进行了对比。这道题虽然不是特别难，但包含了很多重要的知识点，是学习树操作的很好的题目。希望这份笔记对复习和面试准备有所帮助！