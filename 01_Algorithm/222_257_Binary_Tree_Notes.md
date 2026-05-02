# 222. 完全二叉树的节点个数 & 257. 二叉树的所有路径 学习笔记

## 一、 题目分析

### 1.1 222. 完全二叉树的节点个数

#### 1.1.1 题目描述

给出一个**完全二叉树**，求出该树的节点个数。

**完全二叉树的定义：**
- 除了最底层节点可能没填满外，其余每层节点数都达到最大值
- 最下面一层的节点都集中在该层最左边的若干位置
- 若最底层为第 h 层，则该层包含 1~2^(h-1) 个节点

#### 1.1.2 示例

**示例 1：**
```
输入：root = [1,2,3,4,5,6]
输出：6
```

**图示：**
```
     1
    / \
   2   3
  / \ /
 4  5 6
```

**示例 2：**
```
输入：root = []
输出：0
```

**示例 3：**
```
输入：root = [1]
输出：1
```

#### 1.1.3 提示

- 树中节点的数目范围是 `[0, 5 * 10^4]`
- 0 <= Node.val <= 5 * 10^4
- 题目数据保证输入的树是**完全二叉树**

---

### 1.2 257. 二叉树的所有路径

#### 1.2.1 题目描述

给定一个二叉树，返回所有从根节点到叶子节点的路径。

**说明：** 叶子节点是指没有子节点的节点。

#### 1.2.2 示例

**示例 1：**
```
输入：root = [1,2,3,null,5]
输出：["1->2->5","1->3"]
```

**图示：**
```
     1
    / \
   2   3
    \
     5
```

**示例 2：**
```
输入：root = [1]
输出：["1"]
```

#### 1.2.3 提示

- 树中节点的数目在范围 `[1, 100]` 内
- -100 <= Node.val <= 100

---

## 二、 解题思路分析

### 2.1 222. 完全二叉树的节点个数

#### 2.1.1 方法一：普通二叉树解法（递归）

**思路：**
- 后序遍历（左-右-中）
- 节点数 = 左子树节点数 + 右子树节点数 + 1（当前节点）

**递归三要素：**

1. **参数和返回值**：
   - 参数：当前节点指针
   - 返回值：以该节点为根的子树的节点数

2. **终止条件**：
   - 当前节点为空，返回 0

3. **单层逻辑**：
   - 递归求左子树节点数
   - 递归求右子树节点数
   - 返回左 + 右 + 1

#### 2.1.2 方法二：普通二叉树解法（迭代）

**思路：**
- 使用队列进行层序遍历
- 遍历过程中计数

**步骤：**
1. 创建队列，将根节点入队
2. 初始化计数器为 0
3. 当队列不为空时：
   - 取出队首节点
   - 计数器 +1
   - 将左右子节点入队
4. 返回计数器

#### 2.1.3 方法三：利用完全二叉树性质（最优）

**思路：**
- 完全二叉树要么是满二叉树，要么最后一层不满
- 如果是满二叉树，节点数 = 2^h - 1（h 为深度，根为 1）
- 如果不是满二叉树，递归处理左右子树，直到遇到满二叉树

**如何判断满二叉树：**
- 沿着左边界和右边界同时向下遍历
- 如果深度相同，则是满二叉树
- 如果不同，则不是

#### 2.1.4 三种方法对比

| 方法 | 时间复杂度 | 空间复杂度 | 特点 |
|------|-----------|-----------|------|
| 递归法 | O(n) | O(logn) | 简单，但未利用完全二叉树性质 |
| 迭代法 | O(n) | O(n) | 需要额外空间，未利用性质 |
| 利用性质 | O(logn * logn) | O(logn) | 最优，充分利用完全二叉树性质 |

---

### 2.2 257. 二叉树的所有路径

#### 2.2.1 方法一：递归 + 回溯

**思路：**
- 前序遍历（中-左-右），因为需要从根节点开始记录路径
- 使用回溯来维护当前路径
- 到达叶子节点时，将路径加入结果集

**递归三要素：**

1. **参数和返回值**：
   - 参数：当前节点、当前路径、结果集
   - 返回值：无（通过引用修改结果集）

2. **终止条件**：
   - 当前节点是叶子节点（左右都为空）

3. **单层逻辑**：
   - 将当前节点加入路径
   - 如果是叶子节点，记录路径
   - 递归左子树，然后回溯
   - 递归右子树，然后回溯

**回溯的重要性：**
- 路径是共享的，递归返回后需要回溯（移除当前节点）
- 否则会影响其他分支的路径

#### 2.2.2 方法二：迭代法

**思路：**
- 使用两个栈：一个存储节点，一个存储对应路径
- 模拟递归过程
- 到达叶子节点时记录路径

**步骤：**
1. 创建两个栈：节点栈和路径栈
2. 将根节点和初始路径入栈
3. 当栈不为空时：
   - 弹出节点和对应的路径
   - 如果是叶子节点，记录路径
   - 将右子节点和新路径入栈
   - 将左子节点和新路径入栈

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

---

### 3.2 222. 完全二叉树的节点个数

#### 3.2.1 方法一：递归法

```cpp
class Solution1 {
public:
    int countNodes(TreeNode* root) {
        // 终止条件：空节点返回 0
        if (root == nullptr) {
            return 0;
        }
        
        // 后序遍历：先左，后右，最后中间
        int leftNum = countNodes(root->left);   // 左子树节点数
        int rightNum = countNodes(root->right); // 右子树节点数
        
        // 当前节点为根的子树节点数 = 左 + 右 + 1（自己）
        return leftNum + rightNum + 1;
    }
};
```

**代码分析：**
- 这是最直接的递归实现
- 时间复杂度 O(n)，每个节点访问一次
- 空间复杂度 O(logn)，递归栈深度等于树的高度

#### 3.2.2 方法二：迭代法（层序遍历）

```cpp
#include <queue>
using namespace std;

class Solution2 {
public:
    int countNodes(TreeNode* root) {
        // 如果根为空，直接返回 0
        if (root == nullptr) {
            return 0;
        }
        
        // 创建队列进行层序遍历
        queue<TreeNode*> que;
        que.push(root);
        
        int result = 0;  // 节点计数器
        
        while (!que.empty()) {
            int size = que.size();  // 当前层的节点数
            
            for (int i = 0; i < size; ++i) {
                TreeNode* node = que.front();
                que.pop();
                
                result++;  // 计数
                
                // 将子节点入队
                if (node->left != nullptr) {
                    que.push(node->left);
                }
                if (node->right != nullptr) {
                    que.push(node->right);
                }
            }
        }
        
        return result;
    }
};
```

**代码分析：**
- 使用队列进行层序遍历
- 时间复杂度 O(n)，每个节点访问一次
- 空间复杂度 O(n)，最坏情况队列中有 n/2 个节点

#### 3.2.3 方法三：利用完全二叉树性质（最优）

```cpp
class Solution3 {
public:
    int countNodes(TreeNode* root) {
        // 空节点返回 0
        if (root == nullptr) {
            return 0;
        }
        
        // ==========================================
        // 计算左边界深度
        // ==========================================
        TreeNode* left = root->left;
        int leftDepth = 0;
        while (left != nullptr) {
            left = left->left;
            leftDepth++;
        }
        
        // ==========================================
        // 计算右边界深度
        // ==========================================
        TreeNode* right = root->right;
        int rightDepth = 0;
        while (right != nullptr) {
            right = right->right;
            rightDepth++;
        }
        
        // ==========================================
        // 判断是否为满二叉树
        // ==========================================
        if (leftDepth == rightDepth) {
            // 满二叉树：节点数 = 2^深度 - 1
            // 注意：左边界深度是从 0 开始的，所以要 +1
            return (1 << (leftDepth + 1)) - 1;
        }
        
        // ==========================================
        // 不是满二叉树，递归处理左右子树
        // ==========================================
        return 1 + countNodes(root->left) + countNodes(root->right);
    }
};
```

**代码分析：**
- **核心优化**：如果是满二叉树，直接用公式计算
- **判断满二叉树**：比较左右边界深度
- **时间复杂度**：O(logn * logn)
  - 每次计算深度是 O(logn)
  - 最多递归 O(logn) 次
- **空间复杂度**：O(logn)，递归栈深度

---

### 3.3 257. 二叉树的所有路径

#### 3.3.1 方法一：递归 + 回溯

```cpp
#include <vector>
#include <string>
using namespace std;

class Solution4 {
private:
    // ==========================================
    // 递归辅助函数
    // cur: 当前节点
    // path: 当前路径（字符串形式）
    // result: 结果集
    // ==========================================
    void traversal(TreeNode* cur, string path, vector<string>& result) {
        // ==========================================
        // 将当前节点加入路径（中）
        // ==========================================
        path += to_string(cur->val);
        
        // ==========================================
        // 终止条件：到达叶子节点
        // ==========================================
        if (cur->left == nullptr && cur->right == nullptr) {
            result.push_back(path);  // 记录路径
            return;
        }
        
        // ==========================================
        // 递归左子树（左）
        // ==========================================
        if (cur->left != nullptr) {
            traversal(cur->left, path + "->", result);
            // 这里不需要手动回溯，因为 path 是值传递
            // 每次递归都会创建新的字符串副本
        }
        
        // ==========================================
        // 递归右子树（右）
        // ==========================================
        if (cur->right != nullptr) {
            traversal(cur->right, path + "->", result);
            // 同上，不需要手动回溯
        }
    }
    
public:
    vector<string> binaryTreePaths(TreeNode* root) {
        vector<string> result;
        string path;
        
        // 边界处理
        if (root == nullptr) {
            return result;
        }
        
        // 调用递归函数
        traversal(root, path, result);
        
        return result;
    }
};
```

**代码分析：**
- **路径传递方式**：使用值传递（string path），每次递归创建新副本
- **优点**：不需要手动回溯，简单直观
- **缺点**：字符串拷贝有一定开销

**另一种写法（使用引用 + 手动回溯）：**

```cpp
class Solution5 {
private:
    void traversal(TreeNode* cur, string& path, vector<string>& result) {
        // 将当前节点加入路径
        string valStr = to_string(cur->val);
        
        // 如果路径不为空，需要加 "->"
        if (!path.empty()) {
            path += "->";
        }
        path += valStr;
        
        // 终止条件：叶子节点
        if (cur->left == nullptr && cur->right == nullptr) {
            result.push_back(path);
            // 回溯：移除当前节点（包括可能的 "->"）
            if (path.size() >= valStr.size()) {
                path.resize(path.size() - valStr.size());
                if (!path.empty() && path.back() == '>') {
                    // 移除 "->"
                    path.resize(path.size() - 2);
                }
            }
            return;
        }
        
        // 递归左子树
        if (cur->left != nullptr) {
            traversal(cur->left, path, result);
        }
        
        // 递归右子树
        if (cur->right != nullptr) {
            traversal(cur->right, path, result);
        }
        
        // 回溯：移除当前节点
        if (path.size() >= valStr.size()) {
            path.resize(path.size() - valStr.size());
            if (!path.empty() && path.back() == '>') {
                path.resize(path.size() - 2);
            }
        }
    }
    
public:
    vector<string> binaryTreePaths(TreeNode* root) {
        vector<string> result;
        string path;
        
        if (root == nullptr) {
            return result;
        }
        
        traversal(root, path, result);
        return result;
    }
};
```

**代码分析：**
- **路径传递方式**：使用引用传递（string& path）
- **手动回溯**：需要在递归返回后移除当前节点
- **优点**：避免字符串拷贝，效率更高
- **缺点**：代码更复杂，容易出错

#### 3.3.2 方法二：迭代法

```cpp
#include <vector>
#include <string>
#include <stack>
using namespace std;

class Solution6 {
public:
    vector<string> binaryTreePaths(TreeNode* root) {
        vector<string> result;
        
        // 边界处理
        if (root == nullptr) {
            return result;
        }
        
        // 创建两个栈：一个存储节点，一个存储路径
        stack<TreeNode*> nodeStack;
        stack<string> pathStack;
        
        // 初始化：将根节点和空路径入栈
        nodeStack.push(root);
        pathStack.push("");
        
        while (!nodeStack.empty()) {
            // 弹出节点和对应的路径
            TreeNode* cur = nodeStack.top();
            nodeStack.pop();
            
            string path = pathStack.top();
            pathStack.pop();
            
            // 将当前节点值加入路径
            if (path.empty()) {
                path = to_string(cur->val);
            } else {
                path += "->" + to_string(cur->val);
            }
            
            // 终止条件：到达叶子节点
            if (cur->left == nullptr && cur->right == nullptr) {
                result.push_back(path);
                continue;
            }
            
            // 注意：栈是后入先出，所以先压右子树，再压左子树
            // 这样弹出时会先处理左子树
            if (cur->right != nullptr) {
                nodeStack.push(cur->right);
                pathStack.push(path);
            }
            
            if (cur->left != nullptr) {
                nodeStack.push(cur->left);
                pathStack.push(path);
            }
        }
        
        return result;
    }
};
```

**代码分析：**
- **使用两个栈**：节点栈和路径栈，保持同步
- **入栈顺序**：先右后左，保证左子树先被处理
- **时间复杂度**：O(n)，每个节点访问一次
- **空间复杂度**：O(n)，最坏情况栈中有 n 个元素

---

## 四、 时间与空间复杂度分析

### 4.1 222. 完全二叉树的节点个数

#### 4.1.1 递归法

| 复杂度 | 分析 |
|--------|------|
| **时间** | O(n) | 每个节点访问一次 |
| **空间** | O(logn) | 递归栈深度，完全二叉树高度为 logn |

#### 4.1.2 迭代法

| 复杂度 | 分析 |
|--------|------|
| **时间** | O(n) | 每个节点入队出队一次 |
| **空间** | O(n) | 队列最大容量为 n/2（最后一层） |

#### 4.1.3 利用完全二叉树性质

| 复杂度 | 分析 |
|--------|------|
| **时间** | O((logn)^2) | 每次递归计算深度 O(logn)，最多递归 O(logn) 次 |
| **空间** | O(logn) | 递归栈深度 |

**为什么是 O((logn)^2)？**
```
对于完全二叉树：
- 每次计算左右边界深度：O(logn)
- 递归次数：O(logn)（每次至少减少一半节点）
- 总时间：O(logn) * O(logn) = O((logn)^2)

对比普通递归的 O(n)，当 n = 5*10^4 时：
- O(n) ≈ 5*10^4 次操作
- O((logn)^2) ≈ (16)^2 = 256 次操作
```

### 4.2 257. 二叉树的所有路径

#### 4.2.1 递归 + 回溯

| 复杂度 | 分析 |
|--------|------|
| **时间** | O(n) | 每个节点访问一次 |
| **空间** | O(n) | 递归栈深度 + 路径存储 |

#### 4.2.2 迭代法

| 复杂度 | 分析 |
|--------|------|
| **时间** | O(n) | 每个节点入栈出栈一次 |
| **空间** | O(n) | 栈和路径存储 |

---

## 五、 测试用例验证

### 5.1 222. 完全二叉树的节点个数

#### 测试用例 1：普通完全二叉树
```
输入：root = [1,2,3,4,5,6]
预期输出：6
```

**执行过程：**
```
根节点 1，左深度=2，右深度=1（不相等）
递归左子树 2，左深度=1，右深度=1（相等），返回 3（2^2-1）
递归右子树 3，左深度=1，右深度=0（不相等）
  递归左子树 6，左右深度=0，返回 1
  递归右子树 null，返回 0
  右子树 3 返回 2
总节点数 = 1 + 3 + 2 = 6
```

#### 测试用例 2：空树
```
输入：root = []
预期输出：0
```

#### 测试用例 3：单节点
```
输入：root = [1]
预期输出：1
```

#### 测试用例 4：满二叉树
```
输入：root = [1,2,3,4,5,6,7]
预期输出：7（2^3 - 1）
```

### 5.2 257. 二叉树的所有路径

#### 测试用例 1：普通二叉树
```
输入：root = [1,2,3,null,5]
预期输出：["1->2->5", "1->3"]
```

**执行过程（递归）：**
```
1. 访问 1，路径 = "1"
2. 递归左子树 2，路径 = "1->2"
3. 访问 2，路径 = "1->2"
4. 递归左子树 null（跳过）
5. 递归右子树 5，路径 = "1->2->5"
6. 访问 5（叶子节点），加入结果集
7. 返回，回溯到 2，路径 = "1->2"
8. 返回，回溯到 1，路径 = "1"
9. 递归右子树 3，路径 = "1->3"
10. 访问 3（叶子节点），加入结果集
11. 返回，回溯到 1，路径 = "1"
```

#### 测试用例 2：单节点
```
输入：root = [1]
预期输出：["1"]
```

#### 测试用例 3：左斜树
```
输入：root = [1,2,null,3,null,4]
预期输出：["1->2->3->4"]
```

---

## 六、 常见错误与注意事项

### 6.1 222. 完全二叉树的节点个数

#### 错误 1：忘记判断空节点

```cpp
// ❌ 错误示例
int countNodes(TreeNode* root) {
    return 1 + countNodes(root->left) + countNodes(root->right);
    // 如果 root 为空，访问 root->left 会崩溃！
}
```

**正确做法：**
```cpp
// ✅ 正确
int countNodes(TreeNode* root) {
    if (root == nullptr) return 0;  // 先判空
    return 1 + countNodes(root->left) + countNodes(root->right);
}
```

#### 错误 2：位运算错误

```cpp
// ❌ 错误示例：忘记 +1
if (leftDepth == rightDepth) {
    return (1 << leftDepth) - 1;  // 应该是 leftDepth + 1
}
```

**正确做法：**
```cpp
// ✅ 正确
if (leftDepth == rightDepth) {
    return (1 << (leftDepth + 1)) - 1;
}
```

**为什么要 +1？**
- 左边界深度是从 0 开始计数的（空树深度为 -1，单节点深度为 0）
- 但满二叉树公式是 2^h - 1，其中 h 是从 1 开始的深度
- 所以需要 +1

### 6.2 257. 二叉树的所有路径

#### 错误 1：路径拼接错误

```cpp
// ❌ 错误示例：最后一个节点后也加 "->"
path += to_string(cur->val) + "->";  // 错误！
```

**正确做法：**
```cpp
// ✅ 正确：只在非叶子节点后加 "->"
path += to_string(cur->val);
if (cur->left || cur->right) {
    path += "->";
}
```

#### 错误 2：忘记回溯

```cpp
// ❌ 错误示例：使用引用传递但不回溯
void traversal(TreeNode* cur, string& path, vector<string>& result) {
    path += to_string(cur->val);
    if (cur->left == nullptr && cur->right == nullptr) {
        result.push_back(path);
        return;  // 没有回溯！
    }
    if (cur->left) {
        traversal(cur->left, path, result);
    }
    // path 中还包含左子树的节点！
}
```

**正确做法：**
```cpp
// ✅ 正确：使用值传递，自动回溯
void traversal(TreeNode* cur, string path, vector<string>& result) {
    path += to_string(cur->val);
    // ...
}
```

#### 错误 3：迭代法入栈顺序错误

```cpp
// ❌ 错误示例：先压左后压右
if (cur->left) {
    nodeStack.push(cur->left);
}
if (cur->right) {
    nodeStack.push(cur->right);
}
// 弹出顺序变成：右 -> 左，不符合前序遍历
```

**正确做法：**
```cpp
// ✅ 正确：先压右后压左
if (cur->right) {
    nodeStack.push(cur->right);
}
if (cur->left) {
    nodeStack.push(cur->left);
}
// 弹出顺序：左 -> 右
```

---

## 七、 与其他相关题目的对比分析

### 7.1 222 vs 104. 二叉树的最大深度

| 题目 | 目标 | 递归逻辑 |
|------|------|---------|
| 222 | 计数节点 | 左 + 右 + 1 |
| 104 | 计算深度 | max(左, 右) + 1 |

**代码对比：**
```cpp
// 222. 节点数
int countNodes(TreeNode* root) {
    if (!root) return 0;
    return 1 + countNodes(root->left) + countNodes(root->right);
}

// 104. 最大深度
int maxDepth(TreeNode* root) {
    if (!root) return 0;
    return 1 + max(maxDepth(root->left), maxDepth(root->right));
}
```

### 7.2 222 vs 257：递归 vs 回溯

| 题目 | 操作类型 | 是否需要回溯 |
|------|---------|------------|
| 222 | 计数 | 不需要，只返回数值 |
| 257 | 路径记录 | 需要，路径需要回溯 |

**回溯的本质：**
- 当需要维护一个**共享状态**（如当前路径）时，需要回溯
- 当只需要**返回值**时，不需要回溯

### 7.3 完全二叉树 vs 普通二叉树

| 特性 | 普通二叉树 | 完全二叉树 |
|------|-----------|-----------|
| 节点分布 | 任意 | 除最后一层外都满，最后一层靠左 |
| 节点数计算 | O(n) | 最优 O((logn)^2) |
| 高度计算 | 需要遍历 | 可通过左边界快速计算 |

---

## 八、 面试相关内容

### 8.1 222. 完全二叉树的节点个数

#### Q1: 完全二叉树和满二叉树有什么区别？

**答案：**

**满二叉树：**
- 所有层的节点数都达到最大值
- 第 h 层有 2^(h-1) 个节点
- 总节点数 = 2^h - 1

**完全二叉树：**
- 除最后一层外，所有层的节点数都达到最大值
- 最后一层的节点都集中在最左边
- 最后一层节点数范围：1 ~ 2^(h-1)

**关系：**
- 满二叉树一定是完全二叉树
- 完全二叉树不一定是满二叉树

#### Q2: 为什么完全二叉树可以用 O((logn)^2) 的时间？

**答案：**

因为完全二叉树具有以下性质：
1. 可以通过比较左右边界深度判断是否为满二叉树
2. 满二叉树可以直接用公式计算节点数（O(1)）
3. 如果不是满二叉树，递归处理左右子树时，至少有一个子树是满二叉树

这样每次递归的问题规模至少减少一半，总共递归 O(logn) 次，每次计算深度 O(logn)，总时间 O((logn)^2)。

#### Q3: 如何判断一棵完全二叉树是不是满二叉树？

**答案：**

**方法：**
1. 沿着左边界向下遍历，记录深度
2. 沿着右边界向下遍历，记录深度
3. 如果两个深度相等，则是满二叉树
4. 如果不相等，则不是满二叉树

```cpp
bool isPerfectBinaryTree(TreeNode* root) {
    if (!root) return true;
    
    int leftDepth = 0, rightDepth = 0;
    TreeNode* left = root;
    TreeNode* right = root;
    
    while (left) { left = left->left; leftDepth++; }
    while (right) { right = right->right; rightDepth++; }
    
    return leftDepth == rightDepth;
}
```

---

### 8.2 257. 二叉树的所有路径

#### Q1: 为什么需要回溯？

**答案：**

**因为路径是共享状态！**

在递归过程中，如果我们使用引用传递路径：
```cpp
void traversal(TreeNode* cur, string& path, ...) {
    path += to_string(cur->val);  // 修改共享路径
    // ...
    traversal(cur->left, path, ...);  // 左子树会修改 path
    // 如果不回溯，path 中还包含左子树的节点
    traversal(cur->right, path, ...);  // 右子树会错误地使用修改后的 path
}
```

**回溯的作用：**
- 在递归返回后，撤销对共享状态的修改
- 确保不同分支之间不互相影响

#### Q2: 递归中的回溯和迭代中的回溯有什么区别？

**答案：**

**递归中的回溯：**
- 自动发生：函数栈帧保存了调用前的状态
- 可以通过值传递避免手动回溯
- 如果使用引用传递，需要手动回溯

**迭代中的回溯：**
- 需要手动维护路径状态
- 使用栈保存每个节点对应的路径
- 弹出节点时，路径也随之弹出

#### Q3: 如果要求路径从叶子到根，应该怎么做？

**答案：**

**方法 1：先收集根到叶子的路径，然后反转**

```cpp
vector<string> binaryTreePathsReverse(TreeNode* root) {
    vector<string> result;
    // ... 收集根到叶子的路径 ...
    
    // 反转每条路径
    for (string& path : result) {
        reverse(path.begin(), path.end());
        // 还需要反转 "->"，变成 "<-"
        // 更简单的方法是收集时就按叶子到根的顺序
    }
    return result;
}
```

**方法 2：后序遍历，从叶子开始构建路径**

```cpp
void traversal(TreeNode* cur, string path, vector<string>& result) {
    if (!cur) return;
    
    path = to_string(cur->val) + (path.empty() ? "" : "->" + path);
    
    if (!cur->left && !cur->right) {
        result.push_back(path);
        return;
    }
    
    traversal(cur->left, path, result);
    traversal(cur->right, path, result);
}
```

---

## 九、 完整参考代码

### 9.1 222. 完全二叉树的节点个数

```cpp
#include <iostream>
#include <queue>
using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) 
        : val(x), left(left), right(right) {}
};

// 方法一：递归法
class Solution1 {
public:
    int countNodes(TreeNode* root) {
        if (root == nullptr) return 0;
        return 1 + countNodes(root->left) + countNodes(root->right);
    }
};

// 方法二：迭代法（层序遍历）
class Solution2 {
public:
    int countNodes(TreeNode* root) {
        if (root == nullptr) return 0;
        
        queue<TreeNode*> que;
        que.push(root);
        int result = 0;
        
        while (!que.empty()) {
            int size = que.size();
            for (int i = 0; i < size; ++i) {
                TreeNode* node = que.front();
                que.pop();
                result++;
                if (node->left) que.push(node->left);
                if (node->right) que.push(node->right);
            }
        }
        
        return result;
    }
};

// 方法三：利用完全二叉树性质（最优）
class Solution3 {
public:
    int countNodes(TreeNode* root) {
        if (root == nullptr) return 0;
        
        // 计算左边界深度
        TreeNode* left = root->left;
        int leftDepth = 0;
        while (left) {
            left = left->left;
            leftDepth++;
        }
        
        // 计算右边界深度
        TreeNode* right = root->right;
        int rightDepth = 0;
        while (right) {
            right = right->right;
            rightDepth++;
        }
        
        // 判断是否为满二叉树
        if (leftDepth == rightDepth) {
            return (1 << (leftDepth + 1)) - 1;
        }
        
        // 递归处理左右子树
        return 1 + countNodes(root->left) + countNodes(root->right);
    }
};

// 测试代码
void test222() {
    // 构建测试树：[1,2,3,4,5,6]
    TreeNode* root = new TreeNode(1);
    root->left = new TreeNode(2);
    root->right = new TreeNode(3);
    root->left->left = new TreeNode(4);
    root->left->right = new TreeNode(5);
    root->right->left = new TreeNode(6);
    
    Solution1 s1;
    Solution2 s2;
    Solution3 s3;
    
    cout << "====== 222. 完全二叉树节点个数测试 ======" << endl;
    cout << "递归法: " << s1.countNodes(root) << endl;     // 6
    cout << "迭代法: " << s2.countNodes(root) << endl;     // 6
    cout << "最优法: " << s3.countNodes(root) << endl;     // 6
    
    // 释放内存（省略）
}
```

### 9.2 257. 二叉树的所有路径

```cpp
#include <iostream>
#include <vector>
#include <string>
#include <stack>
using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) 
        : val(x), left(left), right(right) {}
};

// 方法一：递归（值传递，自动回溯）
class Solution4 {
private:
    void traversal(TreeNode* cur, string path, vector<string>& result) {
        path += to_string(cur->val);
        
        if (!cur->left && !cur->right) {
            result.push_back(path);
            return;
        }
        
        if (cur->left) {
            traversal(cur->left, path + "->", result);
        }
        if (cur->right) {
            traversal(cur->right, path + "->", result);
        }
    }
    
public:
    vector<string> binaryTreePaths(TreeNode* root) {
        vector<string> result;
        if (!root) return result;
        traversal(root, "", result);
        return result;
    }
};

// 方法二：迭代法
class Solution5 {
public:
    vector<string> binaryTreePaths(TreeNode* root) {
        vector<string> result;
        if (!root) return result;
        
        stack<TreeNode*> nodeStack;
        stack<string> pathStack;
        
        nodeStack.push(root);
        pathStack.push("");
        
        while (!nodeStack.empty()) {
            TreeNode* cur = nodeStack.top();
            nodeStack.pop();
            
            string path = pathStack.top();
            pathStack.pop();
            
            if (path.empty()) {
                path = to_string(cur->val);
            } else {
                path += "->" + to_string(cur->val);
            }
            
            if (!cur->left && !cur->right) {
                result.push_back(path);
                continue;
            }
            
            if (cur->right) {
                nodeStack.push(cur->right);
                pathStack.push(path);
            }
            if (cur->left) {
                nodeStack.push(cur->left);
                pathStack.push(path);
            }
        }
        
        return result;
    }
};

// 测试代码
void test257() {
    // 构建测试树：[1,2,3,null,5]
    TreeNode* root = new TreeNode(1);
    root->left = new TreeNode(2);
    root->right = new TreeNode(3);
    root->left->right = new TreeNode(5);
    
    Solution4 s4;
    Solution5 s5;
    
    vector<string> result4 = s4.binaryTreePaths(root);
    vector<string> result5 = s5.binaryTreePaths(root);
    
    cout << "\n====== 257. 二叉树的所有路径测试 ======" << endl;
    cout << "递归法: ";
    for (const string& path : result4) {
        cout << path << " ";
    }
    cout << endl;  // "1->2->5" "1->3"
    
    cout << "迭代法: ";
    for (const string& path : result5) {
        cout << path << " ";
    }
    cout << endl;  // "1->2->5" "1->3"
    
    // 释放内存（省略）
}

int main() {
    test222();
    test257();
    return 0;
}
```

---

## 十、 总结与反思

### 10.1 核心知识点总结

| 题目 | 核心知识点 | 关键技巧 |
|------|-----------|---------|
| 222 | 完全二叉树性质 | 判断满二叉树，利用公式计算 |
| 257 | 递归 + 回溯 | 前序遍历，路径维护，回溯处理 |

### 10.2 学习收获

**222. 完全二叉树的节点个数：**
1. 理解了完全二叉树和满二叉树的区别
2. 学会了利用完全二叉树性质优化时间复杂度
3. 掌握了判断满二叉树的方法（比较左右边界深度）

**257. 二叉树的所有路径：**
1. 理解了回溯的重要性
2. 掌握了递归中回溯的两种方式（值传递 vs 引用传递）
3. 学会了迭代法模拟递归过程

### 10.3 易错点回顾

**222：**
- 忘记判断空节点
- 位运算错误（忘记 +1）

**257：**
- 路径拼接错误（最后一个节点后加 "->"）
- 忘记回溯（使用引用传递但不回溯）
- 迭代法入栈顺序错误

### 10.4 后续学习建议

1. **练习更多二叉树路径问题**：
   - 路径总和问题
   - 从根到叶子的数字之和

2. **学习回溯算法**：
   - 回溯的本质
   - 排列组合问题
   - 子集问题

3. **深入理解完全二叉树**：
   - 堆数据结构
   - 二叉堆的实现

---

通过这两道题，我们深入理解了完全二叉树的性质和二叉树路径问题的解法，掌握了递归、迭代和回溯等重要技巧。这些知识对于解决更复杂的树问题非常重要！
