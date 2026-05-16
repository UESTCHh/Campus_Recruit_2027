# 113. 路径总和 II学习笔记

## 一、 题目分析

### 1.1 题目描述

给你二叉树的根节点 `root` 和一个整数目标和 `targetSum`，找出**所有**从根节点到叶子节点路径总和等于给定目标和的路径。

**叶子节点**：指没有子节点的节点。

### 1.2 题目链接

- [LeetCode 113. 路径总和 II](https://leetcode.cn/problems/path-sum-ii/description/)
- [官方题解](https://leetcode.cn/problems/path-sum-ii/solutions/427759/lu-jing-zong-he-ii-by-leetcode-solution/)
- [回溯常见问题](https://leetcode.cn/problems/path-sum-ii/solutions/3061294/hui-su-fu-chang-jian-wen-ti-ji-qi-jie-da-g8im/)
- [代码随想录题解](https://www.programmercarl.com/0112.%E8%B7%AF%E5%BE%84%E6%80%BB%E5%92%8C.html)

### 1.3 示例

**示例 1：**

```
输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：[[5,4,11,2],[5,8,4,5]]
解释：存在两条路径满足目标和为 22：
5 -> 4 -> 11 -> 2
5 -> 8 -> 4 -> 5
```

**示例 2：**

```
输入：root = [1,2,3], targetSum = 5
输出：[]
```

**示例 3：**

```
输入：root = [1,2], targetSum = 0
输出：[]
```

### 1.4 提示

- 树中节点总数在范围 `[0, 5000]` 内
- `-1000 <= Node.val <= 1000`
- `-1000 <= targetSum <= 1000`

### 1.5 题目本质

路径总和 II 的本质是：**在二叉树中搜索所有从根到叶子的路径，收集所有路径和等于目标值的路径**。

这是一个典型的**深度优先搜索（DFS）+ 回溯**问题，需要遍历所有可能的根到叶子路径，并在满足条件时记录路径。

### 1.6 解题关键点

| 关键点 | 说明 |
|-------|------|
| 路径定义 | 必须从根节点到叶子节点 |
| 叶子节点 | 左右子节点都为空的节点 |
| 回溯思想 | 需要维护当前路径，回溯时移除当前节点 |
| 递归终止条件 | 当前节点为空时返回；当前节点是叶子且路径和等于目标和时记录路径 |

---

## 二、 解题思路分析

### 2.1 核心思路

解决路径总和 II 问题的核心思路是：**使用深度优先搜索（DFS）遍历二叉树，结合回溯法维护当前路径，当路径和等于目标和且到达叶子节点时，记录该路径**。

### 2.2 递归+回溯思路

#### 2.2.1 递归+回溯三步曲

**1. 确定递归函数参数和返回值**
- 参数：当前节点 `root`、当前剩余目标和 `targetSum`、当前路径 `path`、结果列表 `result`
- 返回值：无（结果通过引用参数 `result` 收集）

**2. 确定终止条件**
- 如果当前节点为空，直接返回
- 如果当前节点是叶子节点且路径和等于目标和，将当前路径加入结果列表

**3. 确定单层递归逻辑**
- 将当前节点值加入路径
- 递归检查左子树：传入 `targetSum - root->val`
- 递归检查右子树：传入 `targetSum - root->val`
- **回溯**：从路径中移除当前节点值

#### 2.2.2 回溯流程图

```
回溯流程：
┌─────────────────────────────────────────────────────────┐
│  dfs(root, targetSum, path, result)                    │
├─────────────────────────────────────────────────────────┤
│                                                         │
│  1. 判空：if (root == nullptr) return                   │
│                                                         │
│  2. 加入路径：path.push_back(root->val)                 │
│                                                         │
│  3. 判断叶子：                                           │
│     if (root->left == nullptr && root->right == nullptr)│
│         if (root->val == targetSum)                     │
│             result.push_back(path)                      │
│         return                                          │
│                                                         │
│  4. 递归左子树：                                         │
│     dfs(root->left, targetSum - root->val, path, result)│
│                                                         │
│  5. 递归右子树：                                         │
│     dfs(root->right, targetSum - root->val, path, result)│
│                                                         │
│  6. 回溯：path.pop_back()                               │
│                                                         │
└─────────────────────────────────────────────────────────┘
```

### 2.3 迭代思路

使用栈模拟递归过程，同时维护当前路径和路径和：
1. 如果根节点为空，返回空列表
2. 将根节点、当前路径和、当前路径加入栈
3. 遍历栈：
   - 取出栈顶元素
   - 如果是叶子节点且路径和等于目标和，将路径加入结果列表
   - 否则将右子节点和左子节点加入栈（注意顺序）
4. 返回结果列表

---

## 三、 算法实现

### 3.1 递归+回溯解法

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
    vector<vector<int>> pathSum(TreeNode* root, int targetSum) {
        vector<vector<int>> result;
        vector<int> path;
        dfs(root, targetSum, path, result);
        return result;
    }
    
private:
    void dfs(TreeNode* root, int targetSum, vector<int>& path, vector<vector<int>>& result) {
        // 终止条件：空节点
        if (root == nullptr) {
            return;
        }
        
        // 将当前节点加入路径
        path.push_back(root->val);
        
        // 终止条件：叶子节点且路径和等于目标和
        if (root->left == nullptr && root->right == nullptr) {
            if (root->val == targetSum) {
                result.push_back(path);
            }
            // 回溯：移除当前节点
            path.pop_back();
            return;
        }
        
        // 递归左子树
        dfs(root->left, targetSum - root->val, path, result);
        
        // 递归右子树
        dfs(root->right, targetSum - root->val, path, result);
        
        // 回溯：移除当前节点
        path.pop_back();
    }
};
```

### 3.2 简洁版递归+回溯

```cpp
class Solution {
public:
    vector<vector<int>> pathSum(TreeNode* root, int targetSum) {
        vector<vector<int>> res;
        vector<int> path;
        function<void(TreeNode*, int)> dfs = [&](TreeNode* node, int sum) {
            if (!node) return;
            path.push_back(node->val);
            if (!node->left && !node->right && node->val == sum) {
                res.push_back(path);
            }
            dfs(node->left, sum - node->val);
            dfs(node->right, sum - node->val);
            path.pop_back();
        };
        dfs(root, targetSum);
        return res;
    }
};
```

### 3.3 迭代解法

```cpp
class Solution {
public:
    vector<vector<int>> pathSum(TreeNode* root, int targetSum) {
        vector<vector<int>> result;
        if (root == nullptr) {
            return result;
        }
        
        // 栈中存储：节点、当前路径和、当前路径
        stack<tuple<TreeNode*, int, vector<int>>> st;
        st.push({root, root->val, {root->val}});
        
        while (!st.empty()) {
            auto [node, currSum, path] = st.top();
            st.pop();
            
            // 判断是否是叶子节点
            if (node->left == nullptr && node->right == nullptr) {
                if (currSum == targetSum) {
                    result.push_back(path);
                }
                continue;
            }
            
            // 右子节点先入栈（栈是后进先出）
            if (node->right != nullptr) {
                vector<int> newPath = path;
                newPath.push_back(node->right->val);
                st.push({node->right, currSum + node->right->val, newPath});
            }
            
            // 左子节点后入栈
            if (node->left != nullptr) {
                vector<int> newPath = path;
                newPath.push_back(node->left->val);
                st.push({node->left, currSum + node->left->val, newPath});
            }
        }
        
        return result;
    }
};
```

---

## 四、 代码分析

### 4.1 递归+回溯解法分析

**核心逻辑：**
1. **判空处理**：如果当前节点为空，直接返回
2. **路径入栈**：将当前节点值加入路径
3. **叶子节点判断**：如果是叶子节点且路径和等于目标和，将路径加入结果列表
4. **递归调用**：分别递归检查左右子树
5. **路径出栈（回溯）**：从路径中移除当前节点，恢复状态

**回溯的重要性：**
- 回溯是一种"尝试-恢复"的思想
- 在递归返回后，需要将当前节点从路径中移除，以便尝试其他分支
- 这是收集所有路径的关键

**递归调用示例（示例 1）：**
```
输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22

递归过程：
pathSum(root(5), 22)
  dfs(5, 22, [], result)
    path = [5]
    dfs(4, 17, [5], result)
      path = [5, 4]
      dfs(11, 13, [5, 4], result)
        path = [5, 4, 11]
        dfs(7, 2, [5, 4, 11], result)
          path = [5, 4, 11, 7]
          7是叶子，7 != 2，返回
          path = [5, 4, 11]
        dfs(2, 2, [5, 4, 11], result)
          path = [5, 4, 11, 2]
          2是叶子，2 == 2，加入result
          result = [[5,4,11,2]]
          path = [5, 4, 11]
        path = [5, 4]
      path = [5]
    dfs(8, 17, [5], result)
      path = [5, 8]
      ... 继续递归找到第二条路径 [5,8,4,5]
```

### 4.2 迭代解法分析

**核心逻辑：**
- 使用栈存储节点、路径和、路径信息
- 每次从栈中取出元素，检查是否是叶子节点
- 如果不是叶子节点，将左右子节点加入栈
- 需要注意栈的顺序：右子节点先入栈，左子节点后入栈

**空间开销：**
- 每次入栈都需要复制路径，空间复杂度较高
- 适合树不深的情况

---

## 五、 复杂度分析

### 5.1 递归+回溯解法复杂度

| 复杂度类型 | 时间复杂度 | 空间复杂度 |
|-----------|-----------|-----------|
| **递归+回溯** | $O(n)$ | $O(n)$ |

**时间复杂度推导：**
- 每个节点最多被访问一次
- 每次访问需要 $O(1)$ 的操作（除了复制路径）
- 复制路径的时间：最坏情况下每条路径有 $O(n)$ 个节点，最多有 $O(n)$ 条路径
- 总时间复杂度为 $O(n)$（主要由节点访问决定）

**空间复杂度推导：**
- 递归调用栈的深度取决于树的高度 $h$
- 当前路径的空间最多 $O(n)$
- 结果列表的空间最多 $O(n^2)$（存储所有路径）
- 空间复杂度为 $O(n)$（不考虑结果存储）

### 5.2 迭代解法复杂度

| 复杂度类型 | 时间复杂度 | 空间复杂度 |
|-----------|-----------|-----------|
| **迭代解法** | $O(n)$ | $O(n^2)$ |

**时间复杂度推导：**
- 每个节点最多被访问一次
- 时间复杂度为 $O(n)$

**空间复杂度推导：**
- 栈中每个元素都包含完整路径的副本
- 最坏情况下空间复杂度为 $O(n^2)$

### 5.3 复杂度对比

| 方法 | 时间复杂度 | 空间复杂度 | 优点 | 缺点 |
|------|-----------|-----------|------|------|
| 递归+回溯 | $O(n)$ | $O(n)$ | 代码简洁，空间效率高 | 可能栈溢出 |
| 迭代 | $O(n)$ | $O(n^2)$ | 避免栈溢出 | 空间开销大 |

---

## 六、 测试用例验证

### 6.1 测试用例 1：存在两条路径

**输入：**
```
root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
```

**执行过程：**
1. 递归遍历路径 5 -> 4 -> 11 -> 7，和为 27，不等于 22
2. 递归遍历路径 5 -> 4 -> 11 -> 2，和为 22，加入结果
3. 递归遍历路径 5 -> 8 -> 13，和为 26，不等于 22
4. 递归遍历路径 5 -> 8 -> 4 -> 5，和为 22，加入结果
5. 递归遍历路径 5 -> 8 -> 4 -> 1，和为 18，不等于 22

**输出：** `[[5,4,11,2],[5,8,4,5]]`

### 6.2 测试用例 2：不存在路径

**输入：**
```
root = [1,2,3], targetSum = 5
```

**执行过程：**
1. 路径 1 -> 2，和为 3，不等于 5
2. 路径 1 -> 3，和为 4，不等于 5

**输出：** `[]`

### 6.3 测试用例 3：空树

**输入：**
```
root = [], targetSum = 0
```

**执行过程：**
1. 根节点为空，直接返回空列表

**输出：** `[]`

### 6.4 测试用例 4：单节点树

**输入：**
```
root = [1], targetSum = 1
```

**执行过程：**
1. 根节点是叶子节点，节点值 1 等于目标和 1
2. 加入结果

**输出：** `[[1]]`

### 6.5 测试用例 5：负数节点值

**输入：**
```
root = [-2,null,-3], targetSum = -5
```

**执行过程：**
1. 路径 -2 -> -3，和为 -5，等于目标和
2. 加入结果

**输出：** `[[-2,-3]]`

---

## 七、 常见错误与注意事项

### 7.1 忘记回溯

**错误示例：**
```cpp
// 错误：没有回溯，路径会不断增长
void dfs(TreeNode* root, int targetSum, vector<int>& path, vector<vector<int>>& result) {
    if (root == nullptr) return;
    path.push_back(root->val);  // 入栈
    if (root->left == nullptr && root->right == nullptr) {
        if (root->val == targetSum) {
            result.push_back(path);
        }
        return;  // 没有出栈！
    }
    dfs(root->left, targetSum - root->val, path, result);
    dfs(root->right, targetSum - root->val, path, result);
    // 没有出栈！
}
```

**后果：**
- 路径会不断累积所有访问过的节点
- 最终路径包含所有节点，而不是从根到叶子的路径

**正确做法：**
```cpp
path.push_back(root->val);
// ... 递归逻辑 ...
path.pop_back();  // 回溯：恢复状态
```

### 7.2 叶子节点判断错误

**错误示例：**
```cpp
// 错误：只判断左子节点为空
if (root->left == nullptr) {
    if (root->val == targetSum) {
        result.push_back(path);
    }
}
```

**后果：**
- 可能将非叶子节点当作叶子节点处理
- 导致错误的路径被加入结果

**正确做法：**
```cpp
if (root->left == nullptr && root->right == nullptr) {
    if (root->val == targetSum) {
        result.push_back(path);
    }
}
```

### 7.3 结果复制错误

**错误示例：**
```cpp
// 错误：直接保存路径的引用
result.push_back(&path);  // 错误！
```

**后果：**
- 保存的是路径的引用，回溯后路径内容会改变
- 最终结果列表中所有路径都指向同一个不断变化的路径

**正确做法：**
```cpp
result.push_back(path);  // 正确：复制路径
```

### 7.4 空树处理

**错误示例：**
```cpp
// 错误：没有处理空树
vector<vector<int>> pathSum(TreeNode* root, int targetSum) {
    vector<vector<int>> result;
    vector<int> path;
    dfs(root, targetSum, path, result);  // root 为空时会崩溃
    return result;
}
```

**后果：**
- 如果输入是空树，会导致空指针异常

**正确做法：**
```cpp
if (root == nullptr) {
    return result;
}
```

---

## 八、 与其他题目的对比分析

### 8.1 与 112. 路径总和的对比

| 题目 | 核心问题 | 输出 | 解法差异 |
|------|---------|------|---------|
| 112. 路径总和 | 判断是否存在 | 布尔值 | 找到一条路径即可返回 |
| 113. 路径总和 II | 找出所有路径 | 路径列表 | 需要回溯收集所有路径 |

**共同点：**
- 都需要遍历从根到叶子的路径
- 都可以使用 DFS 或 BFS

**不同点：**
- 112 题找到一条满足条件的路径即可返回
- 113 题需要收集所有满足条件的路径
- 113 题必须使用回溯法来维护路径状态

### 8.2 与 257. 二叉树的所有路径的对比

| 题目 | 核心问题 | 输出 | 解法 |
|------|---------|------|------|
| 113. 路径总和 II | 找出所有路径和等于目标值的路径 | 整数列表列表 | DFS+回溯 |
| 257. 二叉树的所有路径 | 找出所有根到叶子的路径 | 字符串列表 | DFS+回溯 |

**共同点：**
- 都需要遍历所有根到叶子的路径
- 都需要使用回溯法

**不同点：**
- 113 题需要判断路径和是否等于目标值
- 257 题需要将路径转换为字符串格式

---

## 九、 面试高频问题与解答

### 9.1 基础问题

#### Q1：什么是回溯算法？

**答案：**
- 回溯算法是一种通过探索所有可能的候选解来找出所有解的算法
- 如果候选解被确认不是一个有效的解（或至少不是最后一个解），回溯算法会通过在上一步进行一些变化来丢弃该解
- 在树的遍历中，回溯意味着在递归返回后恢复状态（如从路径中移除当前节点）

#### Q2：为什么需要回溯？

**答案：**
- 因为我们需要维护一条从根到当前节点的路径
- 当我们从一个分支返回到父节点时，需要移除当前节点，才能正确地探索其他分支
- 如果不回溯，路径会包含所有访问过的节点，而不是从根到叶子的路径

#### Q3：递归和回溯的关系是什么？

**答案：**
- 回溯通常与递归结合使用
- 递归负责深度优先搜索
- 回溯负责在递归返回时恢复状态

### 9.2 进阶问题

#### Q4：如果只需要找出一条满足条件的路径，如何优化？

**答案：**
- 可以让递归函数返回布尔值，表示是否找到了路径
- 一旦找到一条路径，就立即返回 true，不再继续搜索
- 这就是 112 题的解法

#### Q5：如果树的深度很大，递归解法会有什么问题？

**答案：**
- 递归解法会导致栈溢出
- 因为递归调用栈的深度等于树的高度
- 对于深树，应该使用迭代解法

#### Q6：迭代解法中为什么右子节点先入栈？

**答案：**
- 因为栈是后进先出的数据结构
- 右子节点先入栈，左子节点后入栈
- 这样可以保证左子树先被访问（与递归顺序一致）

### 9.3 实际问题

#### Q7：在实际工程中，什么时候使用递归+回溯？

**答案：**
- 需要收集所有满足条件的解时
- 需要维护路径或状态时
- 问题具有明显的树形结构时

#### Q8：如何优化空间复杂度？

**答案：**
- 使用引用传递路径，避免复制
- 在递归返回时进行回溯，而不是每次都复制路径
- 对于迭代解法，可以使用索引来标记路径，而不是存储完整路径

---

## 十、 完整可运行代码

```cpp
#include <iostream>
#include <vector>
#include <stack>
#include <tuple>
#include <functional>
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

// 解法 1：递归+回溯
class RecursiveSolution {
public:
    vector<vector<int>> pathSum(TreeNode* root, int targetSum) {
        vector<vector<int>> result;
        vector<int> path;
        dfs(root, targetSum, path, result);
        return result;
    }
    
private:
    void dfs(TreeNode* root, int targetSum, vector<int>& path, vector<vector<int>>& result) {
        if (root == nullptr) {
            return;
        }
        
        path.push_back(root->val);
        
        if (root->left == nullptr && root->right == nullptr) {
            if (root->val == targetSum) {
                result.push_back(path);
            }
            path.pop_back();
            return;
        }
        
        dfs(root->left, targetSum - root->val, path, result);
        dfs(root->right, targetSum - root->val, path, result);
        
        path.pop_back();
    }
};

// 解法 2：简洁版递归+回溯
class ConciseRecursiveSolution {
public:
    vector<vector<int>> pathSum(TreeNode* root, int targetSum) {
        vector<vector<int>> res;
        vector<int> path;
        function<void(TreeNode*, int)> dfs = [&](TreeNode* node, int sum) {
            if (!node) return;
            path.push_back(node->val);
            if (!node->left && !node->right && node->val == sum) {
                res.push_back(path);
            }
            dfs(node->left, sum - node->val);
            dfs(node->right, sum - node->val);
            path.pop_back();
        };
        dfs(root, targetSum);
        return res;
    }
};

// 解法 3：迭代解法
class IterativeSolution {
public:
    vector<vector<int>> pathSum(TreeNode* root, int targetSum) {
        vector<vector<int>> result;
        if (root == nullptr) {
            return result;
        }
        
        stack<tuple<TreeNode*, int, vector<int>>> st;
        st.push({root, root->val, {root->val}});
        
        while (!st.empty()) {
            auto [node, currSum, path] = st.top();
            st.pop();
            
            if (node->left == nullptr && node->right == nullptr) {
                if (currSum == targetSum) {
                    result.push_back(path);
                }
                continue;
            }
            
            if (node->right != nullptr) {
                vector<int> newPath = path;
                newPath.push_back(node->right->val);
                st.push({node->right, currSum + node->right->val, newPath});
            }
            
            if (node->left != nullptr) {
                vector<int> newPath = path;
                newPath.push_back(node->left->val);
                st.push({node->left, currSum + node->left->val, newPath});
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

// 辅助函数：打印路径列表
void printPaths(const vector<vector<int>>& paths) {
    cout << "[";
    for (size_t i = 0; i < paths.size(); ++i) {
        cout << "[";
        for (size_t j = 0; j < paths[i].size(); ++j) {
            cout << paths[i][j];
            if (j != paths[i].size() - 1) {
                cout << ",";
            }
        }
        cout << "]";
        if (i != paths.size() - 1) {
            cout << ",";
        }
    }
    cout << "]" << endl;
}

// 主函数测试
int main() {
    cout << "========================================" << endl;
    cout << "     113. 路径总和 II 测试" << endl;
    cout << "========================================" << endl;

    // 创建解法实例
    RecursiveSolution recursiveSol;
    ConciseRecursiveSolution conciseSol;
    IterativeSolution iterativeSol;

    // 测试用例 1：存在两条路径
    cout << "\n--- 测试用例 1 ---" << endl;
    vector<int> vals1 = {5,4,11,7,-1,-1,2,-1,-1,-1,8,13,-1,-1,4,5,-1,-1,1,-1,-1};
    int index1 = 0;
    TreeNode* root1 = createTree(vals1, index1);
    cout << "输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22" << endl;
    cout << "递归解法：";
    printPaths(recursiveSol.pathSum(root1, 22));
    cout << "期望：[[5,4,11,2],[5,8,4,5]]" << endl;

    // 测试用例 2：不存在路径
    cout << "\n--- 测试用例 2 ---" << endl;
    vector<int> vals2 = {1,2,-1,-1,3,-1,-1};
    int index2 = 0;
    TreeNode* root2 = createTree(vals2, index2);
    cout << "输入：root = [1,2,3], targetSum = 5" << endl;
    cout << "递归解法：";
    printPaths(recursiveSol.pathSum(root2, 5));
    cout << "期望：[]" << endl;

    // 测试用例 3：空树
    cout << "\n--- 测试用例 3 ---" << endl;
    TreeNode* root3 = nullptr;
    cout << "输入：root = [], targetSum = 0" << endl;
    cout << "递归解法：";
    printPaths(recursiveSol.pathSum(root3, 0));
    cout << "期望：[]" << endl;

    // 测试用例 4：单节点树
    cout << "\n--- 测试用例 4 ---" << endl;
    TreeNode* root4 = new TreeNode(1);
    cout << "输入：root = [1], targetSum = 1" << endl;
    cout << "递归解法：";
    printPaths(recursiveSol.pathSum(root4, 1));
    cout << "期望：[[1]]" << endl;

    // 测试用例 5：负数节点值
    cout << "\n--- 测试用例 5 ---" << endl;
    TreeNode* root5 = new TreeNode(-2);
    root5->right = new TreeNode(-3);
    cout << "输入：root = [-2,null,-3], targetSum = -5" << endl;
    cout << "递归解法：";
    printPaths(recursiveSol.pathSum(root5, -5));
    cout << "期望：[[-2,-3]]" << endl;

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
| 回溯思想 | 在递归返回时恢复状态 |
| 路径维护 | 使用向量存储当前路径，回溯时弹出 |
| 叶子节点判断 | 左右子节点都为空 |
| 递归终止条件 | 空节点返回；叶子节点检查路径和 |
| 时间复杂度 | $O(n)$，每个节点访问一次 |
| 空间复杂度 | $O(n)$（不考虑结果存储） |

### 11.2 学习收获

1. **理解了回溯算法的核心思想**
   - 回溯是一种"尝试-恢复"的思想
   - 在递归深入时记录状态，在返回时恢复状态
   - 是解决"收集所有解"类型问题的关键

2. **掌握了递归+回溯的写法**
   - 入栈：在递归前加入当前节点
   - 递归：处理左右子树
   - 出栈：在递归后移除当前节点

3. **学会了如何处理边界条件**
   - 空树的处理
   - 叶子节点的判断
   - 路径的正确维护

4. **理解了迭代解法的思路**
   - 使用栈模拟递归过程
   - 存储节点、路径和、路径信息
   - 注意栈的顺序

### 11.3 后续学习建议

1. **练习更多回溯问题**
   - 46. 全排列
   - 47. 全排列 II
   - 77. 组合
   - 39. 组合总和

2. **深入理解回溯的应用场景**
   - 子集问题
   - 排列问题
   - 组合问题
   - 切割问题

3. **学习剪枝优化**
   - 在回溯过程中提前终止无效路径
   - 优化时间复杂度

4. **练习迭代写法**
   - 使用栈模拟递归
   - 理解迭代和递归的对应关系

---

## 💡 面试核心考点

* **直击灵魂的拷问一：什么是回溯算法？**
  绝杀回答：回溯算法是一种通过探索所有可能的候选解来找出所有解的算法。如果候选解被确认不是有效的解，回溯算法会通过在上一步进行一些变化来丢弃该解。在树的遍历中，回溯意味着在递归返回后恢复状态，例如从路径中移除当前节点。

* **直击灵魂的拷问二：为什么需要回溯？**
  绝杀回答：因为我们需要维护一条从根到当前节点的路径。当我们从一个分支返回到父节点时，需要移除当前节点，才能正确地探索其他分支。如果不回溯，路径会包含所有访问过的节点，而不是从根到叶子的路径。

* **直击灵魂的拷问三：递归和回溯的关系是什么？**
  绝杀回答：回溯通常与递归结合使用。递归负责深度优先搜索，回溯负责在递归返回时恢复状态。两者相辅相成，缺一不可。

* **直击灵魂的拷问四：如果只需要找出一条路径，如何优化？**
  绝杀回答：可以让递归函数返回布尔值，表示是否找到了路径。一旦找到一条路径，就立即返回 true，不再继续搜索。这就是 112 题的解法。

通过本课程，我们深入理解了路径总和 II 问题的核心逻辑，掌握了递归+回溯的解法，也学习了常见的错误注意事项。这些知识对于理解回溯算法和树的遍历非常重要！
