# 112. 路径总和学习笔记

## 一、 题目分析

### 1.1 题目描述

给定一个二叉树的根节点 `root` 和一个整数 `targetSum`，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和 `targetSum`。

**叶子节点**：指没有子节点的节点。

### 1.2 题目链接

- [LeetCode 112. 路径总和](https://leetcode.cn/problems/path-sum/description/)
- [代码随想录题解](https://www.programmercarl.com/0112.%E8%B7%AF%E5%BE%84%E6%80%BB%E5%92%8C.html)
- [官方题解](https://leetcode.cn/problems/path-sum/solutions/318487/lu-jing-zong-he-by-leetcode-solution/)
- [简洁解法](https://leetcode.cn/problems/path-sum/solutions/2731531/jian-ji-xie-fa-pythonjavacgojsrust-by-en-icwe/)

### 1.3 示例

**示例 1：**

```
输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
输出：true
解释：存在目标和为 22 的路径：5 -> 4 -> 11 -> 2
```

**示例 2：**

```
输入：root = [1,2,3], targetSum = 5
输出：false
解释：不存在目标和为 5 的路径
```

**示例 3：**

```
输入：root = [], targetSum = 0
输出：false
解释：空树没有路径
```

### 1.4 提示

- 树中节点的数目在范围 `[0, 5000]` 内
- `-1000 <= Node.val <= 1000`
- `-1000 <= targetSum <= 1000`

### 1.5 题目本质

路径总和问题的本质是：**在二叉树中搜索从根到叶子的路径，判断是否存在一条路径的节点值之和等于目标值**。

这是一个典型的**深度优先搜索（DFS）**问题，需要遍历所有可能的根到叶子路径，并计算路径和。

### 1.6 解题关键点

| 关键点 | 说明 |
|-------|------|
| 路径定义 | 必须从根节点到叶子节点 |
| 叶子节点 | 左右子节点都为空的节点 |
| 递归终止条件 | 当前节点为空时返回 false；当前节点是叶子且值等于剩余目标和时返回 true |
| 递归逻辑 | 递归检查左子树和右子树，只要有一条路径满足即可 |

---

## 二、 解题思路分析

### 2.1 核心思路

解决路径总和问题的核心思路是：**使用深度优先搜索（DFS）遍历二叉树，递归检查每条从根到叶子的路径**。

### 2.2 递归思路

#### 2.2.1 递归三步曲

**1. 确定递归函数参数和返回值**
- 参数：当前节点 `root`、当前剩余目标和 `targetSum`
- 返回值：布尔值，表示是否存在满足条件的路径

**2. 确定终止条件**
- 如果当前节点为空，返回 `false`
- 如果当前节点是叶子节点，判断 `root->val == targetSum`

**3. 确定单层递归逻辑**
- 递归检查左子树：`hasPathSum(root->left, targetSum - root->val)`
- 递归检查右子树：`hasPathSum(root->right, targetSum - root->val)`
- 返回左右子树结果的**逻辑或**（只要有一条路径满足即可）

#### 2.2.2 递归流程图

```
递归调用流程：
hasPathSum(root, targetSum)
    │
    ▼
判断：root == null? ──是──▶ 返回 false
    │否
    ▼
判断：root 是叶子? ──是──▶ 返回 (root->val == targetSum)
    │否
    ▼
递归：hasPathSum(left, targetSum - root->val)
    │
    ▼
递归：hasPathSum(right, targetSum - root->val)
    │
    ▼
返回：left_result || right_result
```

### 2.3 迭代思路

使用广度优先搜索（BFS），维护两个队列：
1. 节点队列：存储待访问的节点
2. 路径和队列：存储从根到当前节点的路径和

**步骤：**
1. 如果根节点为空，返回 false
2. 将根节点和根节点的值分别加入两个队列
3. 遍历队列：
   - 取出队首节点和对应的路径和
   - 如果是叶子节点且路径和等于目标和，返回 true
   - 否则将左右子节点加入队列，更新路径和
4. 遍历结束未找到，返回 false

---

## 三、 算法实现

### 3.1 递归解法（DFS）

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
    bool hasPathSum(TreeNode* root, int targetSum) {
        // 终止条件：空节点，返回 false
        if (root == nullptr) {
            return false;
        }
        
        // 终止条件：叶子节点，判断值是否等于剩余目标和
        if (root->left == nullptr && root->right == nullptr) {
            return root->val == targetSum;
        }
        
        // 递归检查左子树和右子树
        int remaining = targetSum - root->val;
        return hasPathSum(root->left, remaining) || hasPathSum(root->right, remaining);
    }
};
```

### 3.2 简洁递归解法

```cpp
class Solution {
public:
    bool hasPathSum(TreeNode* root, int targetSum) {
        if (!root) return false;
        if (!root->left && !root->right) return root->val == targetSum;
        return hasPathSum(root->left, targetSum - root->val) 
            || hasPathSum(root->right, targetSum - root->val);
    }
};
```

### 3.3 迭代解法（BFS）

```cpp
class Solution {
public:
    bool hasPathSum(TreeNode* root, int targetSum) {
        if (root == nullptr) {
            return false;
        }
        
        queue<TreeNode*> nodeQueue;
        queue<int> sumQueue;
        
        nodeQueue.push(root);
        sumQueue.push(root->val);
        
        while (!nodeQueue.empty()) {
            TreeNode* curr = nodeQueue.front();
            nodeQueue.pop();
            
            int currSum = sumQueue.front();
            sumQueue.pop();
            
            // 判断是否是叶子节点且路径和等于目标和
            if (curr->left == nullptr && curr->right == nullptr) {
                if (currSum == targetSum) {
                    return true;
                }
                continue;
            }
            
            // 将左子节点加入队列
            if (curr->left != nullptr) {
                nodeQueue.push(curr->left);
                sumQueue.push(currSum + curr->left->val);
            }
            
            // 将右子节点加入队列
            if (curr->right != nullptr) {
                nodeQueue.push(curr->right);
                sumQueue.push(currSum + curr->right->val);
            }
        }
        
        return false;
    }
};
```

### 3.4 迭代解法（DFS - 使用栈）

```cpp
class Solution {
public:
    bool hasPathSum(TreeNode* root, int targetSum) {
        if (root == nullptr) {
            return false;
        }
        
        stack<pair<TreeNode*, int>> st;
        st.push({root, root->val});
        
        while (!st.empty()) {
            auto [curr, currSum] = st.top();
            st.pop();
            
            // 判断是否是叶子节点
            if (curr->left == nullptr && curr->right == nullptr) {
                if (currSum == targetSum) {
                    return true;
                }
                continue;
            }
            
            // 右子节点先入栈（栈是后进先出）
            if (curr->right != nullptr) {
                st.push({curr->right, currSum + curr->right->val});
            }
            
            // 左子节点后入栈
            if (curr->left != nullptr) {
                st.push({curr->left, currSum + curr->left->val});
            }
        }
        
        return false;
    }
};
```

---

## 四、 代码分析

### 4.1 递归解法分析

**核心逻辑：**
1. **空节点处理**：如果当前节点为空，直接返回 `false`
2. **叶子节点判断**：如果是叶子节点，检查节点值是否等于剩余目标和
3. **递归调用**：分别递归检查左右子树，将目标和减去当前节点值传递下去
4. **结果合并**：只要左子树或右子树有一条路径满足条件，就返回 `true`

**递归调用示例（示例 1）：**
```
输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22

递归过程：
hasPathSum(root(5), 22)
  remaining = 22 - 5 = 17
  调用 hasPathSum(4, 17) || hasPathSum(8, 17)
  
hasPathSum(4, 17)
  remaining = 17 - 4 = 13
  调用 hasPathSum(11, 13) || hasPathSum(null, 13)
  
hasPathSum(11, 13)
  remaining = 13 - 11 = 2
  调用 hasPathSum(7, 2) || hasPathSum(2, 2)
  
hasPathSum(7, 2)
  7 是叶子节点，7 != 2，返回 false
  
hasPathSum(2, 2)
  2 是叶子节点，2 == 2，返回 true
  
最终：true || false = true
```

### 4.2 迭代解法分析

**BFS 思路：**
- 使用两个队列分别存储节点和对应的路径和
- 每次从队列中取出节点，检查是否是叶子节点且路径和等于目标和
- 如果不是叶子节点，将左右子节点加入队列，更新路径和

**DFS（栈）思路：**
- 使用栈存储节点和对应的路径和（pair）
- 栈是后进先出，所以右子节点先入栈，左子节点后入栈
- 这样可以保证左子树先被访问（与递归顺序一致）

---

## 五、 复杂度分析

### 5.1 递归解法复杂度

| 复杂度类型 | 时间复杂度 | 空间复杂度 |
|-----------|-----------|-----------|
| **递归解法** | $O(n)$ | $O(h)$ |

**时间复杂度推导：**
- 每个节点最多被访问一次
- $n$ 是树中节点的数量
- 时间复杂度为 $O(n)$

**空间复杂度推导：**
- 递归调用栈的深度取决于树的高度
- 最好情况（平衡树）：$h = \log n$
- 最坏情况（链状树）：$h = n$
- 空间复杂度为 $O(h)$

### 5.2 迭代解法复杂度

| 复杂度类型 | 时间复杂度 | 空间复杂度 |
|-----------|-----------|-----------|
| **BFS 迭代** | $O(n)$ | $O(n)$ |
| **DFS 迭代** | $O(n)$ | $O(n)$ |

**时间复杂度推导：**
- 每个节点最多被访问一次
- 时间复杂度为 $O(n)$

**空间复杂度推导：**
- BFS：队列最多存储一层的节点，最坏情况 $O(n)$
- DFS：栈最多存储一条路径的节点，最坏情况 $O(n)$

### 5.3 复杂度对比

| 方法 | 时间复杂度 | 空间复杂度 | 优点 | 缺点 |
|------|-----------|-----------|------|------|
| 递归 | $O(n)$ | $O(h)$ | 代码简洁，逻辑清晰 | 可能栈溢出（树过深） |
| BFS 迭代 | $O(n)$ | $O(n)$ | 避免栈溢出 | 代码稍复杂 |
| DFS 迭代 | $O(n)$ | $O(n)$ | 避免栈溢出 | 代码稍复杂 |

---

## 六、 测试用例验证

### 6.1 测试用例 1：存在路径

**输入：**
```
root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
```

**执行过程：**
1. 递归检查路径 5 -> 4 -> 11 -> 7，和为 27，不等于 22
2. 递归检查路径 5 -> 4 -> 11 -> 2，和为 22，等于目标和
3. 返回 true

**输出：** `true`

### 6.2 测试用例 2：不存在路径

**输入：**
```
root = [1,2,3], targetSum = 5
```

**执行过程：**
1. 路径 1 -> 2，和为 3，不等于 5
2. 路径 1 -> 3，和为 4，不等于 5
3. 返回 false

**输出：** `false`

### 6.3 测试用例 3：空树

**输入：**
```
root = [], targetSum = 0
```

**执行过程：**
1. 根节点为空，直接返回 false

**输出：** `false`

### 6.4 测试用例 4：单节点树

**输入：**
```
root = [1], targetSum = 1
```

**执行过程：**
1. 根节点是叶子节点
2. 节点值 1 等于目标和 1
3. 返回 true

**输出：** `true`

### 6.5 测试用例 5：负数节点值

**输入：**
```
root = [-2,null,-3], targetSum = -5
```

**执行过程：**
1. 路径 -2 -> -3，和为 -5，等于目标和
2. 返回 true

**输出：** `true`

---

## 七、 常见错误与注意事项

### 7.1 忘记检查叶子节点

**错误示例：**
```cpp
// 错误：没有判断是否是叶子节点
bool hasPathSum(TreeNode* root, int targetSum) {
    if (root == nullptr) return false;
    if (root->val == targetSum) return true;  // 错误！非叶子节点也可能满足
    return hasPathSum(root->left, targetSum - root->val) 
        || hasPathSum(root->right, targetSum - root->val);
}
```

**后果：**
- 如果中间节点的值等于目标和，会错误地返回 true
- 例如：`root = [1,2], targetSum = 1`，会错误返回 true

**正确做法：**
```cpp
// 正确：必须检查是叶子节点
if (root->left == nullptr && root->right == nullptr) {
    return root->val == targetSum;
}
```

### 7.2 目标和传递错误

**错误示例：**
```cpp
// 错误：没有减去当前节点的值
return hasPathSum(root->left, targetSum) || hasPathSum(root->right, targetSum);
```

**后果：**
- 每次递归都使用原始的目标和，无法正确计算路径和

**正确做法：**
```cpp
int remaining = targetSum - root->val;
return hasPathSum(root->left, remaining) || hasPathSum(root->right, remaining);
```

### 7.3 空树处理

**错误示例：**
```cpp
// 错误：没有处理空树
bool hasPathSum(TreeNode* root, int targetSum) {
    // 直接递归，没有检查 root == nullptr
    if (root->left == nullptr && root->right == nullptr) {
        return root->val == targetSum;
    }
    // ...
}
```

**后果：**
- 如果输入是空树，会导致空指针异常

**正确做法：**
```cpp
if (root == nullptr) {
    return false;
}
```

### 7.4 负数节点值的处理

**特殊情况：**
- 节点值可能为负数
- 需要注意目标和也可能为负数
- 递归时传递剩余目标和，而不是绝对值

---

## 八、 与其他题目的对比分析

### 8.1 与 113. 路径总和 II 的对比

| 题目 | 核心问题 | 输出 | 解法差异 |
|------|---------|------|---------|
| 112. 路径总和 | 判断是否存在 | 布尔值 | 找到一条路径即可返回 |
| 113. 路径总和 II | 找出所有路径 | 路径列表 | 需要回溯收集所有路径 |

**共同点：**
- 都需要遍历从根到叶子的路径
- 都可以使用 DFS 或 BFS
- 都需要检查叶子节点

**不同点：**
- 112 题找到一条满足条件的路径即可返回
- 113 题需要收集所有满足条件的路径
- 113 题需要使用回溯法来维护路径状态

### 8.2 与 257. 二叉树的所有路径的对比

| 题目 | 核心问题 | 输出 | 解法 |
|------|---------|------|------|
| 112. 路径总和 | 判断路径和是否等于目标值 | 布尔值 | DFS/BFS |
| 257. 二叉树的所有路径 | 找出所有根到叶子的路径 | 字符串列表 | DFS+回溯 |

**共同点：**
- 都需要遍历所有根到叶子的路径

**不同点：**
- 112 题关注路径和
- 257 题关注路径本身

---

## 九、 面试高频问题与解答

### 9.1 基础问题

#### Q1：什么是叶子节点？

**答案：**
- 叶子节点是指左右子节点都为空的节点
- 在路径总和问题中，路径必须从根节点到叶子节点

#### Q2：递归解法的终止条件是什么？

**答案：**
1. 当前节点为空，返回 `false`
2. 当前节点是叶子节点，判断节点值是否等于剩余目标和

#### Q3：为什么递归时要减去当前节点的值？

**答案：**
- 因为我们要检查从当前节点到叶子节点的路径和是否等于剩余目标和
- 减去当前节点值后，递归调用检查子树时就只需要考虑剩余的目标和

### 9.2 进阶问题

#### Q4：如果树的深度很大（比如 10000 层），递归解法会有什么问题？

**答案：**
- 递归解法会导致栈溢出
- 因为递归调用栈的深度等于树的高度
- 对于深树，应该使用迭代解法（BFS 或 DFS 用栈）

#### Q5：BFS 和 DFS 迭代解法有什么区别？

**答案：**
| 特性 | BFS | DFS |
|------|-----|-----|
| 遍历顺序 | 层序遍历 | 深度优先 |
| 数据结构 | 队列 | 栈 |
| 空间复杂度 | 最坏 $O(n)$（一层节点数） | 最坏 $O(n)$（路径长度） |
| 找到路径的顺序 | 最短路径先找到 | 某条深度路径先找到 |

#### Q6：如何处理节点值为负数的情况？

**答案：**
- 节点值为负数不影响算法逻辑
- 因为我们是用目标和减去节点值，而不是比较绝对值
- 需要注意目标和也可能为负数

### 9.3 实际问题

#### Q7：在实际工程中，什么时候选择递归，什么时候选择迭代？

**答案：**
- **递归**：代码简洁，逻辑清晰，适合树不深的情况
- **迭代**：避免栈溢出，适合树很深或不确定深度的情况

#### Q8：如果二叉树是完全二叉树，哪种解法更高效？

**答案：**
- 时间复杂度都是 $O(n)$
- 空间复杂度：递归是 $O(\log n)$，迭代是 $O(n)$
- 递归更节省空间

---

## 十、 完整可运行代码

```cpp
#include <iostream>
#include <queue>
#include <stack>
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

// 解法 1：递归
class RecursiveSolution {
public:
    bool hasPathSum(TreeNode* root, int targetSum) {
        if (root == nullptr) {
            return false;
        }
        if (root->left == nullptr && root->right == nullptr) {
            return root->val == targetSum;
        }
        int remaining = targetSum - root->val;
        return hasPathSum(root->left, remaining) || hasPathSum(root->right, remaining);
    }
};

// 解法 2：BFS 迭代
class BFSIterativeSolution {
public:
    bool hasPathSum(TreeNode* root, int targetSum) {
        if (root == nullptr) {
            return false;
        }
        
        queue<TreeNode*> nodeQueue;
        queue<int> sumQueue;
        
        nodeQueue.push(root);
        sumQueue.push(root->val);
        
        while (!nodeQueue.empty()) {
            TreeNode* curr = nodeQueue.front();
            nodeQueue.pop();
            
            int currSum = sumQueue.front();
            sumQueue.pop();
            
            if (curr->left == nullptr && curr->right == nullptr) {
                if (currSum == targetSum) {
                    return true;
                }
                continue;
            }
            
            if (curr->left != nullptr) {
                nodeQueue.push(curr->left);
                sumQueue.push(currSum + curr->left->val);
            }
            
            if (curr->right != nullptr) {
                nodeQueue.push(curr->right);
                sumQueue.push(currSum + curr->right->val);
            }
        }
        
        return false;
    }
};

// 解法 3：DFS 迭代（栈）
class DFSIterativeSolution {
public:
    bool hasPathSum(TreeNode* root, int targetSum) {
        if (root == nullptr) {
            return false;
        }
        
        stack<pair<TreeNode*, int>> st;
        st.push({root, root->val});
        
        while (!st.empty()) {
            auto [curr, currSum] = st.top();
            st.pop();
            
            if (curr->left == nullptr && curr->right == nullptr) {
                if (currSum == targetSum) {
                    return true;
                }
                continue;
            }
            
            if (curr->right != nullptr) {
                st.push({curr->right, currSum + curr->right->val});
            }
            
            if (curr->left != nullptr) {
                st.push({curr->left, currSum + curr->left->val});
            }
        }
        
        return false;
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
    cout << "     112. 路径总和测试" << endl;
    cout << "========================================" << endl;

    // 创建解法实例
    RecursiveSolution recursiveSol;
    BFSIterativeSolution bfsSol;
    DFSIterativeSolution dfsSol;

    // 测试用例 1：存在路径
    cout << "\n--- 测试用例 1 ---" << endl;
    vector<int> vals1 = {5,4,11,7,-1,-1,2,-1,-1,-1,8,13,-1,-1,4,-1,1,-1,-1};
    int index1 = 0;
    TreeNode* root1 = createTree(vals1, index1);
    cout << "输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22" << endl;
    cout << "递归解法：" << boolalpha << recursiveSol.hasPathSum(root1, 22) << endl;
    cout << "BFS 解法：" << boolalpha << bfsSol.hasPathSum(root1, 22) << endl;
    cout << "DFS 解法：" << boolalpha << dfsSol.hasPathSum(root1, 22) << endl;
    cout << "期望：true" << endl;

    // 测试用例 2：不存在路径
    cout << "\n--- 测试用例 2 ---" << endl;
    vector<int> vals2 = {1,2,-1,-1,3,-1,-1};
    int index2 = 0;
    TreeNode* root2 = createTree(vals2, index2);
    cout << "输入：root = [1,2,3], targetSum = 5" << endl;
    cout << "递归解法：" << boolalpha << recursiveSol.hasPathSum(root2, 5) << endl;
    cout << "期望：false" << endl;

    // 测试用例 3：空树
    cout << "\n--- 测试用例 3 ---" << endl;
    TreeNode* root3 = nullptr;
    cout << "输入：root = [], targetSum = 0" << endl;
    cout << "递归解法：" << boolalpha << recursiveSol.hasPathSum(root3, 0) << endl;
    cout << "期望：false" << endl;

    // 测试用例 4：单节点树
    cout << "\n--- 测试用例 4 ---" << endl;
    TreeNode* root4 = new TreeNode(1);
    cout << "输入：root = [1], targetSum = 1" << endl;
    cout << "递归解法：" << boolalpha << recursiveSol.hasPathSum(root4, 1) << endl;
    cout << "期望：true" << endl;

    // 测试用例 5：负数节点值
    cout << "\n--- 测试用例 5 ---" << endl;
    TreeNode* root5 = new TreeNode(-2);
    root5->right = new TreeNode(-3);
    cout << "输入：root = [-2,null,-3], targetSum = -5" << endl;
    cout << "递归解法：" << boolalpha << recursiveSol.hasPathSum(root5, -5) << endl;
    cout << "期望：true" << endl;

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
| 递归思路 | 检查每个节点，传递剩余目标和 |
| 终止条件 | 空节点返回 false；叶子节点检查值 |
| 迭代思路 | 使用队列或栈存储节点和路径和 |
| 叶子节点 | 左右子节点都为空的节点 |
| 时间复杂度 | $O(n)$，每个节点访问一次 |
| 空间复杂度 | 递归 $O(h)$，迭代 $O(n)$ |

### 11.2 学习收获

1. **理解了递归在树问题中的应用**
   - 递归是解决树问题的常用方法
   - 需要明确递归函数的参数、返回值和终止条件

2. **掌握了多种解法**
   - 递归解法代码简洁
   - 迭代解法避免栈溢出
   - BFS 和 DFS 各有适用场景

3. **学会了如何处理边界条件**
   - 空树的处理
   - 叶子节点的判断
   - 负数节点值的处理

4. **理解了路径的定义**
   - 路径必须从根节点到叶子节点
   - 叶子节点的定义是左右子节点都为空

### 11.3 后续学习建议

1. **练习更多树的递归问题**
   - 113. 路径总和 II
   - 257. 二叉树的所有路径
   - 129. 求根节点到叶节点数字之和

2. **深入理解回溯法**
   - 学习如何收集所有满足条件的路径
   - 理解回溯过程中的状态维护

3. **学习其他遍历方式**
   - 层序遍历（BFS）
   - 前序、中序、后序遍历（DFS）

4. **练习迭代写法**
   - 使用栈实现 DFS
   - 使用队列实现 BFS

---

## 💡 面试核心考点

* **直击灵魂的拷问一：什么是叶子节点？**
  绝杀回答：叶子节点是指左右子节点都为空的节点。在路径总和问题中，路径必须从根节点到叶子节点，不能是中间节点。

* **直击灵魂的拷问二：递归解法的终止条件是什么？**
  绝杀回答：有两个终止条件。第一，如果当前节点为空，返回 false；第二，如果当前节点是叶子节点，判断节点值是否等于剩余目标和。

* **直击灵魂的拷问三：如果树很深，递归会有什么问题？**
  绝杀回答：递归会导致栈溢出。因为递归调用栈的深度等于树的高度，如果树很深（比如 10000 层），会超出系统栈的限制。此时应该使用迭代解法。

* **直击灵魂的拷问四：BFS 和 DFS 的区别是什么？**
  绝杀回答：BFS 使用队列，按层遍历；DFS 使用栈，按深度遍历。BFS 适合找最短路径，DFS 代码更简洁。空间复杂度都是 O(n)，但 BFS 的空间取决于一层的节点数，DFS 的空间取决于路径长度。

通过本课程，我们深入理解了路径总和问题的核心逻辑，掌握了递归和迭代两种解法，也学习了常见的错误注意事项。这些知识对于理解树的遍历和递归思想非常重要！
