# 513. 找树左下角的值学习笔记

## 一、 题目分析

### 1.1 题目描述

给定一个二叉树的根节点 `root`，请找出该二叉树**最底层最左边**节点的值。

假设二叉树中至少有一个节点。

### 1.2 题目链接

- [LeetCode 513. 找树左下角的值](https://leetcode.cn/problems/find-bottom-left-tree-value/description/)
- [官方题解](https://leetcode.cn/problems/find-bottom-left-tree-value/solutions/1614779/zhao-shu-zuo-xia-jiao-de-zhi-by-leetcode-weeh/)
- [代码随想录题解](https://www.programmercarl.com/0513.%E6%89%BE%E6%A0%91%E5%B7%A6%E4%B8%8B%E8%A7%92%E7%9A%84%E5%80%BC.html)
- [BFS为什么要用队列](https://leetcode.cn/problems/find-bottom-left-tree-value/solutions/2049776/bfs-wei-shi-yao-yao-yong-dui-lie-yi-ge-s-f34y/)

### 1.3 示例

**示例 1：**

```
输入：root = [2,1,3]
输出：1
```

**示例 2：**

```
输入：root = [1,2,3,4,null,5,6,null,null,7]
输出：7
```

### 1.4 提示

- 树中节点数在 `[1, 10^4]` 范围内
- `-2^31 <= Node.val <= 2^31 - 1`

### 1.5 题目本质

找树左下角的值的本质是：**找到二叉树最后一层（最深层）的第一个节点的值**。

这是一个典型的**树遍历**问题，可以使用广度优先搜索（BFS）或深度优先搜索（DFS）来解决。

### 1.6 解题关键点

| 关键点 | 说明 |
|-------|------|
| 目标节点 | 最后一层的最左边节点 |
| 层序遍历 | 按层遍历，记录每层第一个节点 |
| 深度优先 | 优先访问左子树，记录最大深度的第一个节点 |
| 空树处理 | 题目保证至少有一个节点，无需处理空树 |

---

## 二、 解题思路分析

### 2.1 核心思路

解决找树左下角的值问题的核心思路是：**遍历二叉树，找到最后一层的最左边节点**。

主要有两种方法：
1. **BFS（广度优先搜索）**：按层遍历，记录每层的第一个节点，最后一层的第一个节点即为答案
2. **DFS（深度优先搜索）**：优先访问左子树，记录到达的最大深度和对应节点值

### 2.2 BFS思路

#### 2.2.1 BFS核心逻辑

BFS 使用队列进行层序遍历：
1. 将根节点加入队列
2. 遍历队列中的所有节点（当前层）
3. 记录当前层的第一个节点值
4. 将下一层节点加入队列
5. 重复直到队列为空，最后记录的即为答案

#### 2.2.2 BFS流程图

```
BFS流程：
┌─────────────────────────────────────────────────────┐
│  queue = [root]                                     │
│  result = root->val                                 │
├─────────────────────────────────────────────────────┤
│  while queue not empty:                             │
│      level_size = queue.size()                      │
│      for i from 0 to level_size-1:                 │
│          node = queue.front(); queue.pop()          │
│          if i == 0: result = node->val             │  // 记录每层第一个
│          if node->left: queue.push(node->left)      │
│          if node->right: queue.push(node->right)    │
├─────────────────────────────────────────────────────┤
│  return result                                      │
└─────────────────────────────────────────────────────┘
```

### 2.3 DFS思路

#### 2.3.1 DFS核心逻辑

DFS 需要记录两个关键信息：
1. 当前深度
2. 最大深度及对应的节点值

通过前序遍历（根-左-右），优先访问左子树，这样第一次到达某个深度时，该节点就是该层最左边的节点。

#### 2.3.2 DFS流程图

```
DFS流程：
┌─────────────────────────────────────────────────────┐
│  max_depth = -1                                     │
│  result = 0                                         │
├─────────────────────────────────────────────────────┤
│  dfs(node, depth):                                  │
│      if node is null: return                        │
│      if depth > max_depth:                          │
│          max_depth = depth                          │
│          result = node->val                         │
│      dfs(node->left, depth + 1)                    │  // 先左后右
│      dfs(node->right, depth + 1)                   │
├─────────────────────────────────────────────────────┤
│  dfs(root, 0)                                       │
│  return result                                      │
└─────────────────────────────────────────────────────┘
```

---

## 三、 算法实现

### 3.1 BFS解法

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
    int findBottomLeftValue(TreeNode* root) {
        // 使用队列进行层序遍历
        queue<TreeNode*> q;
        q.push(root);
        
        int result = 0;
        
        // 遍历每一层
        while (!q.empty()) {
            // 获取当前层的节点数量
            int levelSize = q.size();
            
            // 遍历当前层的所有节点
            for (int i = 0; i < levelSize; ++i) {
                TreeNode* node = q.front();
                q.pop();
                
                // 记录每层的第一个节点
                if (i == 0) {
                    result = node->val;
                }
                
                // 将子节点加入队列
                if (node->left != nullptr) {
                    q.push(node->left);
                }
                if (node->right != nullptr) {
                    q.push(node->right);
                }
            }
        }
        
        return result;
    }
};
```

### 3.2 DFS递归解法

```cpp
class Solution {
public:
    int findBottomLeftValue(TreeNode* root) {
        int maxDepth = -1;
        int result = 0;
        
        // 深度优先搜索，前序遍历
        function<void(TreeNode*, int)> dfs = [&](TreeNode* node, int depth) {
            // 终止条件：空节点
            if (node == nullptr) {
                return;
            }
            
            // 如果当前深度大于最大深度，更新结果
            if (depth > maxDepth) {
                maxDepth = depth;
                result = node->val;
            }
            
            // 先递归左子树（保证优先访问左边）
            dfs(node->left, depth + 1);
            // 再递归右子树
            dfs(node->right, depth + 1);
        };
        
        // 从根节点开始，深度为0
        dfs(root, 0);
        
        return result;
    }
};
```

### 3.3 DFS迭代解法

```cpp
class Solution {
public:
    int findBottomLeftValue(TreeNode* root) {
        if (root == nullptr) {
            return 0;
        }
        
        int maxDepth = -1;
        int result = 0;
        
        // 栈中存储：节点、深度
        stack<pair<TreeNode*, int>> st;
        st.push({root, 0});
        
        while (!st.empty()) {
            auto [node, depth] = st.top();
            st.pop();
            
            // 更新最大深度和结果
            if (depth > maxDepth) {
                maxDepth = depth;
                result = node->val;
            }
            
            // 栈是后进先出，所以先压右子树，再压左子树
            if (node->right != nullptr) {
                st.push({node->right, depth + 1});
            }
            if (node->left != nullptr) {
                st.push({node->left, depth + 1});
            }
        }
        
        return result;
    }
};
```

---

## 四、 代码分析

### 4.1 BFS解法分析

**核心逻辑：**
1. **初始化队列**：将根节点加入队列
2. **层序遍历**：使用 `while` 循环遍历每一层
3. **记录首节点**：每层第一个节点的值可能是答案（最后一层的首节点就是答案）
4. **子节点入队**：按左到右顺序将子节点加入队列

**BFS的优势：**
- 直观易懂，按层处理符合问题描述
- 不需要额外维护深度信息
- 空间复杂度较低（平均情况）

### 4.2 DFS递归解法分析

**核心逻辑：**
1. **前序遍历**：先访问根节点，再左子树，最后右子树
2. **深度记录**：维护当前深度和最大深度
3. **结果更新**：当深度超过最大深度时，更新结果

**DFS的优势：**
- 空间复杂度更低（递归栈深度为树高）
- 不需要存储整层节点

**递归调用示例（示例 2）：**
```
输入：root = [1,2,3,4,null,5,6,null,null,7]

DFS过程：
dfs(1, 0)
  depth=0 > maxDepth(-1), 更新 maxDepth=0, result=1
  dfs(2, 1)
    depth=1 > maxDepth(0), 更新 maxDepth=1, result=2
    dfs(4, 2)
      depth=2 > maxDepth(1), 更新 maxDepth=2, result=4
      dfs(null, 3) -> 返回
      dfs(null, 3) -> 返回
    dfs(null, 2) -> 返回
  dfs(3, 1)
    dfs(5, 2)
      depth=2 <= maxDepth(2), 不更新
      dfs(null, 3) -> 返回
      dfs(7, 3)
        depth=3 > maxDepth(2), 更新 maxDepth=3, result=7
        ... 返回
    dfs(6, 2)
      depth=2 <= maxDepth(3), 不更新
      ... 返回

最终 result = 7
```

### 4.3 DFS迭代解法分析

**核心逻辑：**
- 使用栈模拟递归过程
- 栈中存储节点和对应的深度
- 注意栈的顺序：先压右子树，再压左子树（保证左子树先被处理）

**空间开销：**
- 栈的空间取决于树的高度
- 最坏情况空间复杂度为 O(n)

---

## 五、 复杂度分析

### 5.1 BFS解法复杂度

| 复杂度类型 | 时间复杂度 | 空间复杂度 |
|-----------|-----------|-----------|
| **BFS** | $O(n)$ | $O(n)$ |

**时间复杂度推导：**
- 每个节点被访问一次
- 每层遍历需要 O(k) 时间（k为当前层节点数）
- 总时间复杂度为 O(n)

**空间复杂度推导：**
- 队列最多存储一层的节点
- 最坏情况（完全二叉树最后一层）：O(n/2) ≈ O(n)

### 5.2 DFS解法复杂度

| 复杂度类型 | 时间复杂度 | 空间复杂度 |
|-----------|-----------|-----------|
| **DFS递归** | $O(n)$ | $O(h)$ |
| **DFS迭代** | $O(n)$ | $O(h)$ |

**时间复杂度推导：**
- 每个节点被访问一次
- 时间复杂度为 O(n)

**空间复杂度推导：**
- 递归栈/栈的深度取决于树的高度 h
- 最好情况（平衡树）：h = log(n)
- 最坏情况（链式树）：h = n

### 5.3 复杂度对比

| 方法 | 时间复杂度 | 空间复杂度 | 优点 | 缺点 |
|------|-----------|-----------|------|------|
| BFS | $O(n)$ | $O(n)$ | 直观，按层处理 | 空间开销较大 |
| DFS递归 | $O(n)$ | $O(h)$ | 空间效率高 | 可能栈溢出 |
| DFS迭代 | $O(n)$ | $O(h)$ | 避免栈溢出 | 代码稍复杂 |

---

## 六、 测试用例验证

### 6.1 测试用例 1：简单二叉树

**输入：**
```
root = [2,1,3]
```

**执行过程（BFS）：**
1. 第一层：[2]，记录 result=2
2. 第二层：[1,3]，记录 result=1
3. 队列为空，返回 result=1

**输出：** `1`

### 6.2 测试用例 2：深度二叉树

**输入：**
```
root = [1,2,3,4,null,5,6,null,null,7]
```

**执行过程（DFS）：**
1. 访问 1（深度0），更新 result=1
2. 访问 2（深度1），更新 result=2
3. 访问 4（深度2），更新 result=4
4. 回到 2，右子树为空
5. 回到 1，访问 3（深度1）
6. 访问 5（深度2），不更新
7. 访问 7（深度3），更新 result=7
8. 访问 6（深度2），不更新

**输出：** `7`

### 6.3 测试用例 3：单节点树

**输入：**
```
root = [1]
```

**执行过程：**
1. 只有一个节点，直接返回 1

**输出：** `1`

### 6.4 测试用例 4：链式树（只有左子树）

**输入：**
```
root = [1,2,null,3,null,4,null,5]
```

**执行过程：**
1. 每层只有一个节点
2. 最后一层节点值为 5

**输出：** `5`

### 6.5 测试用例 5：链式树（只有右子树）

**输入：**
```
root = [1,null,2,null,3,null,4]
```

**执行过程：**
1. 每层只有一个节点（最左边就是唯一节点）
2. 最后一层节点值为 4

**输出：** `4`

---

## 七、 常见错误与注意事项

### 7.1 DFS中先访问右子树

**错误示例：**
```cpp
// 错误：先访问右子树会导致记录的是最右边节点
dfs(node->right, depth + 1);
dfs(node->left, depth + 1);
```

**后果：**
- 会记录每层最右边的节点，而不是最左边的
- 最终结果错误

**正确做法：**
```cpp
// 正确：先访问左子树
dfs(node->left, depth + 1);
dfs(node->right, depth + 1);
```

### 7.2 迭代DFS中栈的顺序错误

**错误示例：**
```cpp
// 错误：先压左子树会导致右子树先被处理
st.push({node->left, depth + 1});
st.push({node->right, depth + 1});
```

**后果：**
- 栈是后进先出，左子树后入栈会先出栈
- 实际访问顺序变成右-左，与递归顺序相反

**正确做法：**
```cpp
// 正确：先压右子树，再压左子树
st.push({node->right, depth + 1});
st.push({node->left, depth + 1});
```

### 7.3 忘记更新最大深度

**错误示例：**
```cpp
// 错误：只记录结果，不更新最大深度
if (depth > maxDepth) {
    result = node->val;
    // 缺少 maxDepth = depth;
}
```

**后果：**
- 可能会被更深层的节点覆盖
- 但如果后续没有更深节点，结果可能正确（运气好）

**正确做法：**
```cpp
if (depth > maxDepth) {
    maxDepth = depth;
    result = node->val;
}
```

### 7.4 BFS中没有按层处理

**错误示例：**
```cpp
// 错误：没有记录每层的第一个节点
while (!q.empty()) {
    TreeNode* node = q.front();
    q.pop();
    result = node->val;  // 记录了最后一个节点，不是最左
    if (node->left) q.push(node->left);
    if (node->right) q.push(node->right);
}
```

**后果：**
- 最终记录的是最后一个被访问的节点
- 而不是最后一层的第一个节点

**正确做法：**
```cpp
while (!q.empty()) {
    int levelSize = q.size();
    for (int i = 0; i < levelSize; ++i) {
        TreeNode* node = q.front();
        q.pop();
        if (i == 0) result = node->val;  // 只记录每层第一个
        if (node->left) q.push(node->left);
        if (node->right) q.push(node->right);
    }
}
```

---

## 八、 与其他题目的对比分析

### 8.1 与 102. 二叉树的层序遍历的对比

| 题目 | 核心问题 | 输出 | 解法差异 |
|------|---------|------|---------|
| 102. 层序遍历 | 输出每层节点 | 二维列表 | 需要完整记录每层 |
| 513. 找左下角值 | 找最后一层第一个 | 整数 | 只需记录每层第一个 |

**共同点：**
- 都使用层序遍历（BFS）
- 都需要按层处理

**不同点：**
- 102 题需要保存所有层的所有节点
- 513 题只需记录每层第一个节点，最后更新的就是答案

### 8.2 与 199. 二叉树的右视图的对比

| 题目 | 核心问题 | 输出 | 解法 |
|------|---------|------|------|
| 513. 找左下角值 | 最后一层第一个 | 整数 | BFS/DFS |
| 199. 二叉树的右视图 | 每层最后一个 | 列表 | BFS/DFS |

**共同点：**
- 都涉及层序遍历
- 都关注每层的特定位置节点

**不同点：**
- 513 题关注每层第一个，199 题关注每层最后一个
- 513 题只需返回一个值，199 题需要返回列表

---

## 九、 面试高频问题与解答

### 9.1 基础问题

#### Q1：为什么BFS要用队列？

**答案：**
- 队列是先进先出（FIFO）的数据结构
- 层序遍历需要按顺序处理同一层的所有节点
- 使用队列可以保证先入队的节点先被处理，符合层序遍历的顺序

#### Q2：DFS为什么要先访问左子树？

**答案：**
- 因为我们要找最左边的节点
- 前序遍历（根-左-右）先访问左子树，保证第一次到达某个深度时，该节点就是该层最左边的节点
- 如果先访问右子树，会记录最右边的节点

#### Q3：递归和迭代的区别是什么？

**答案：**
- 递归使用系统调用栈，代码简洁但可能栈溢出
- 迭代使用手动栈/队列，代码稍复杂但更可控
- 对于深度很大的树，迭代更安全

### 9.2 进阶问题

#### Q4：如果树的深度很大，哪种方法更优？

**答案：**
- BFS 的空间复杂度是 O(n)，与树的宽度相关
- DFS 的空间复杂度是 O(h)，与树的高度相关
- 对于深树（h ≈ n），两者空间复杂度相近
- 对于宽树（宽度大但高度小），DFS 更优

#### Q5：如何优化BFS的空间复杂度？

**答案：**
- 使用双端队列或两个队列交替处理
- 记录当前层和下一层的节点数，避免存储整层节点
- 但实际上，对于找左下角值的问题，BFS的空间复杂度已经是最优的

#### Q6：如果树有重复值怎么办？

**答案：**
- 本题只关心节点的位置（最后一层最左边），不关心节点值是否重复
- 无论节点值是否重复，算法逻辑都不变

### 9.3 实际问题

#### Q7：在实际工程中，什么时候选择BFS，什么时候选择DFS？

**答案：**
- 当问题与层序相关时（如找某一层的节点），选择BFS
- 当问题与路径相关时（如找路径和），选择DFS
- BFS更适合处理树宽较小的情况，DFS更适合处理树深较小的情况

#### Q8：如何验证算法的正确性？

**答案：**
- 测试各种边界情况（单节点、链式树、完全二叉树）
- 手动模拟算法执行过程
- 对比不同解法的结果是否一致

---

## 十、 完整可运行代码

```cpp
#include <iostream>
#include <queue>
#include <stack>
#include <functional>
#include <utility>
using namespace std;

// 二叉树节点定义
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

// 解法 1：BFS
class BFSSolution {
public:
    int findBottomLeftValue(TreeNode* root) {
        queue<TreeNode*> q;
        q.push(root);
        
        int result = 0;
        
        while (!q.empty()) {
            int levelSize = q.size();
            for (int i = 0; i < levelSize; ++i) {
                TreeNode* node = q.front();
                q.pop();
                
                if (i == 0) {
                    result = node->val;
                }
                
                if (node->left != nullptr) {
                    q.push(node->left);
                }
                if (node->right != nullptr) {
                    q.push(node->right);
                }
            }
        }
        
        return result;
    }
};

// 解法 2：DFS递归
class DFSRecursiveSolution {
public:
    int findBottomLeftValue(TreeNode* root) {
        int maxDepth = -1;
        int result = 0;
        
        function<void(TreeNode*, int)> dfs = [&](TreeNode* node, int depth) {
            if (node == nullptr) {
                return;
            }
            
            if (depth > maxDepth) {
                maxDepth = depth;
                result = node->val;
            }
            
            dfs(node->left, depth + 1);
            dfs(node->right, depth + 1);
        };
        
        dfs(root, 0);
        
        return result;
    }
};

// 解法 3：DFS迭代
class DFSIterativeSolution {
public:
    int findBottomLeftValue(TreeNode* root) {
        if (root == nullptr) {
            return 0;
        }
        
        int maxDepth = -1;
        int result = 0;
        
        stack<pair<TreeNode*, int>> st;
        st.push({root, 0});
        
        while (!st.empty()) {
            auto [node, depth] = st.top();
            st.pop();
            
            if (depth > maxDepth) {
                maxDepth = depth;
                result = node->val;
            }
            
            if (node->right != nullptr) {
                st.push({node->right, depth + 1});
            }
            if (node->left != nullptr) {
                st.push({node->left, depth + 1});
            }
        }
        
        return result;
    }
};

// 辅助函数：创建二叉树
TreeNode* createTree(const vector<int>& vals, int& index) {
    if (index >= vals.size() || vals[index] == -1) {
        index++;
        return nullptr;
    }
    TreeNode* node = new TreeNode(vals[index++]);
    node->left = createTree(vals, index);
    node->right = createTree(vals, index);
    return node;
}

// 主函数测试
int main() {
    cout << "========================================" << endl;
    cout << "    513. 找树左下角的值 测试" << endl;
    cout << "========================================" << endl;

    // 创建解法实例
    BFSSolution bfsSol;
    DFSRecursiveSolution dfsRecSol;
    DFSIterativeSolution dfsIterSol;

    // 测试用例 1：简单二叉树
    cout << "\n--- 测试用例 1 ---" << endl;
    vector<int> vals1 = {2,1,-1,-1,3,-1,-1};
    int index1 = 0;
    TreeNode* root1 = createTree(vals1, index1);
    cout << "输入：root = [2,1,3]" << endl;
    cout << "BFS解法：" << bfsSol.findBottomLeftValue(root1) << endl;
    cout << "期望：1" << endl;

    // 测试用例 2：深度二叉树
    cout << "\n--- 测试用例 2 ---" << endl;
    vector<int> vals2 = {1,2,4,-1,-1,-1,3,5,-1,7,-1,-1,6,-1,-1};
    int index2 = 0;
    TreeNode* root2 = createTree(vals2, index2);
    cout << "输入：root = [1,2,3,4,null,5,6,null,null,7]" << endl;
    cout << "BFS解法：" << bfsSol.findBottomLeftValue(root2) << endl;
    cout << "DFS递归：" << dfsRecSol.findBottomLeftValue(root2) << endl;
    cout << "DFS迭代：" << dfsIterSol.findBottomLeftValue(root2) << endl;
    cout << "期望：7" << endl;

    // 测试用例 3：单节点树
    cout << "\n--- 测试用例 3 ---" << endl;
    TreeNode* root3 = new TreeNode(1);
    cout << "输入：root = [1]" << endl;
    cout << "BFS解法：" << bfsSol.findBottomLeftValue(root3) << endl;
    cout << "期望：1" << endl;

    // 测试用例 4：链式树（只有左子树）
    cout << "\n--- 测试用例 4 ---" << endl;
    TreeNode* root4 = new TreeNode(1);
    root4->left = new TreeNode(2);
    root4->left->left = new TreeNode(3);
    root4->left->left->left = new TreeNode(4);
    root4->left->left->left->left = new TreeNode(5);
    cout << "输入：root = [1,2,null,3,null,4,null,5]" << endl;
    cout << "BFS解法：" << bfsSol.findBottomLeftValue(root4) << endl;
    cout << "期望：5" << endl;

    // 测试用例 5：链式树（只有右子树）
    cout << "\n--- 测试用例 5 ---" << endl;
    TreeNode* root5 = new TreeNode(1);
    root5->right = new TreeNode(2);
    root5->right->right = new TreeNode(3);
    root5->right->right->right = new TreeNode(4);
    cout << "输入：root = [1,null,2,null,3,null,4]" << endl;
    cout << "BFS解法：" << bfsSol.findBottomLeftValue(root5) << endl;
    cout << "期望：4" << endl;

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
| BFS层序遍历 | 使用队列，按层处理，记录每层第一个节点 |
| DFS深度优先 | 使用递归或栈，前序遍历，记录最大深度的第一个节点 |
| 时间复杂度 | $O(n)$，每个节点访问一次 |
| 空间复杂度 | BFS: $O(n)$，DFS: $O(h)$ |
| 关键技巧 | BFS按层遍历，DFS先左后右 |

### 11.2 学习收获

1. **理解了BFS和DFS的应用场景**
   - BFS适合层序相关问题
   - DFS适合路径相关问题

2. **掌握了两种主流算法**
   - BFS：队列实现，按层处理
   - DFS：递归/栈实现，深度优先

3. **学会了如何处理边界情况**
   - 单节点树
   - 链式树
   - 完全二叉树

4. **理解了算法优化的思路**
   - 根据树的结构选择合适的算法
   - 平衡时间和空间复杂度

### 11.3 后续学习建议

1. **练习更多树遍历问题**
   - 102. 二叉树的层序遍历
   - 199. 二叉树的右视图
   - 107. 二叉树的层序遍历 II

2. **深入理解BFS和DFS的区别**
   - 时间空间复杂度分析
   - 适用场景对比
   - 实际应用案例

3. **学习树的其他遍历方式**
   - 中序遍历
   - 后序遍历
   - 莫里斯遍历

4. **练习迭代写法**
   - 使用栈模拟递归
   - 理解迭代和递归的对应关系

---

## 💡 面试核心考点

* **直击灵魂的拷问一：为什么BFS要用队列？**
  绝杀回答：队列是先进先出（FIFO）的数据结构，层序遍历需要按顺序处理同一层的所有节点，使用队列可以保证先入队的节点先被处理，符合层序遍历的顺序。

* **直击灵魂的拷问二：DFS为什么要先访问左子树？**
  绝杀回答：因为我们要找最左边的节点。前序遍历（根-左-右）先访问左子树，保证第一次到达某个深度时，该节点就是该层最左边的节点。

* **直击灵魂的拷问三：BFS和DFS的时间空间复杂度各是多少？**
  绝杀回答：BFS的时间复杂度是O(n)，空间复杂度是O(n)（队列最多存储一层节点）；DFS的时间复杂度是O(n)，空间复杂度是O(h)（递归栈深度为树高）。

* **直击灵魂的拷问四：对于深树和宽树，分别选择哪种算法更优？**
  绝杀回答：对于深树，DFS的空间复杂度O(h)更优；对于宽树，BFS的空间复杂度可能更高，但如果只需要找最后一层第一个节点，BFS更直观。

通过本课程，我们深入理解了找树左下角的值问题的核心逻辑，掌握了BFS和DFS两种主流解法，也学习了常见的错误注意事项。这些知识对于理解树遍历和算法选择非常重要！