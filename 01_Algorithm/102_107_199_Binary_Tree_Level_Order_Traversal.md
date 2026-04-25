# 二叉树层序遍历综合学习笔记

> **学习内容**：二叉树的层序遍历（广度优先搜索）
> **学习资源**：代码随想录二叉树层序遍历部分、LeetCode相关题目
> **相关题目**：
> - 102. 二叉树的层序遍历
> - 107. 二叉树的层序遍历 II
> - 199. 二叉树的右视图
> **学习目标**：掌握二叉树层序遍历的原理和实现方法，理解其时间空间复杂度，以及相关题目的解题思路

---

## 一、 层序遍历概述

### 1. 核心概念

**层序遍历**（Level Order Traversal）是二叉树的一种遍历方式，它按照从上到下、从左到右的顺序逐层访问二叉树中的所有节点。

**广度优先搜索**（Breadth-First Search, BFS）是一种图遍历算法，从根节点开始，先访问根节点，然后访问其所有相邻节点，再依次访问这些节点的相邻节点，以此类推。二叉树的层序遍历本质上就是广度优先搜索的一种应用。

### 2. 与深度优先遍历的区别

| 遍历方式 | 特点 | 实现方式 | 适用场景 |
|---------|------|---------|----------|
| 深度优先遍历（DFS） | 先往深走，遇到叶子节点再往回走 | 使用栈（递归或显式栈） | 前序、中序、后序遍历 |
| 广度优先遍历（BFS） | 一层一层的去遍历 | 使用队列 | 层序遍历 |

### 3. 层序遍历示例

对于如下二叉树：
```
    3
   / 
  9  20
    /  
   15   7
```

**层序遍历结果**：[[3], [9, 20], [15, 7]]

---

## 二、 基于队列的层序遍历实现

### 1. 核心原理

层序遍历使用队列（Queue）来实现，利用队列的先进先出（FIFO）特性，确保按层访问节点。

**基本步骤**：
1. 初始化队列，将根节点入队
2. 当队列不为空时，执行以下操作：
   - 记录当前队列的大小（即当前层的节点数）
   - 依次取出队列中的节点，将其值加入结果集
   - 将取出节点的左右子节点（如果存在）入队
3. 重复步骤2，直到队列为空

**为什么使用队列**：
- 队列的先进先出特性保证了节点按层处理的顺序
- 每次处理完一层节点后，队列中恰好存储了下一层的所有节点
- 记录每层大小确保了我们可以清晰地分离不同层的节点

### 2. 详细代码实现与分析

#### 2.1 基础层序遍历（LeetCode 102）

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
    /**
     * @brief 二叉树的层序遍历
     * @param root 二叉树的根节点指针
     * @return 按层存储的节点值二维向量
     */
    vector<vector<int>> levelOrder(TreeNode* root) {
        vector<vector<int>> ans;  // 存储最终结果
        
        // 边界情况处理：空树直接返回空结果
        if (root == nullptr) return ans;
        
        // 创建队列用于层序遍历
        queue<TreeNode*> que;
        que.push(root);  // 将根节点入队
        
        // 当队列不为空时，继续遍历
        while (!que.empty()) {
            int size = que.size();  // 当前层的节点数
            vector<int> temp;  // 存储当前层的节点值
            
            // 遍历当前层的所有节点
            while (size--) {
                TreeNode* f = que.front();  // 获取队首节点
                que.pop();  // 弹出队首节点
                
                temp.push_back(f->val);  // 将节点值加入当前层结果
                
                // 将左子节点入队（如果存在）
                if (f->left) que.push(f->left);
                // 将右子节点入队（如果存在）
                if (f->right) que.push(f->right);
            }
            
            ans.push_back(temp);  // 将当前层结果加入最终结果
        }
        
        return ans;  // 返回层序遍历结果
    }
};
```

**代码分析**：
- **边界情况处理**：首先检查根节点是否为 `nullptr`，如果是则直接返回空结果，避免后续操作出现空指针异常
- **队列初始化**：创建一个 `queue<TreeNode*>` 类型的队列，并将根节点入队，作为遍历的起点
- **外层循环**：当队列不为空时，继续遍历
- **层的处理**：
  - 记录当前队列的大小 `size`，这表示当前层的节点数
  - 创建临时向量 `temp` 用于存储当前层的节点值
  - 内层循环遍历当前层的所有节点
- **节点处理**：
  - 获取并弹出队首节点
  - 将节点值加入当前层的临时向量
  - 将左子节点入队（如果存在）
  - 将右子节点入队（如果存在）
- **结果存储**：将当前层的临时向量加入最终结果

**执行过程示例**：
对于二叉树：
```
    3
   / 
  9  20
    /  
   15   7
```

1. 初始化：`que = [3]`，`ans = []`
2. 第一次外层循环：`size = 1`
   - 内层循环：`size = 1`
     - 弹出 3，`temp = [3]`
     - 入队 9, 20，`que = [9, 20]`
   - `ans = [[3]]`
3. 第二次外层循环：`size = 2`
   - 内层循环：`size = 2`
     - 弹出 9，`temp = [9]`
     - 9 无左右子节点，`que = [20]`
     - 弹出 20，`temp = [9, 20]`
     - 入队 15, 7，`que = [15, 7]`
   - `ans = [[3], [9, 20]]`
4. 第三次外层循环：`size = 2`
   - 内层循环：`size = 2`
     - 弹出 15，`temp = [15]`
     - 15 无左右子节点，`que = [7]`
     - 弹出 7，`temp = [15, 7]`
     - 7 无左右子节点，`que = []`
   - `ans = [[3], [9, 20], [15, 7]]`
5. 队列为空，结束遍历，返回 `ans`

**技术要点**：
- 队列的使用：利用队列的先进先出特性，确保节点按层处理
- 层的分离：通过记录每层开始时的队列大小，确保只处理当前层的节点
- 边界情况：处理空树的情况，避免空指针异常
- 子节点处理：只将非空的子节点入队，避免处理空指针

#### 2.2 自底向上的层序遍历（LeetCode 107）

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
    /**
     * @brief 自底向上的二叉树层序遍历
     * @param root 二叉树的根节点指针
     * @return 按自底向上顺序存储的节点值二维向量
     */
    vector<vector<int>> levelOrderBottom(TreeNode* root) {
        vector<vector<int>> ans;  // 存储最终结果
        
        // 边界情况处理：空树直接返回空结果
        if (root == nullptr) return ans;
        
        // 创建队列用于层序遍历
        queue<TreeNode*> que;
        que.push(root);  // 将根节点入队
        
        // 当队列不为空时，继续遍历
        while (!que.empty()) {
            int size = que.size();  // 当前层的节点数
            vector<int> temp;  // 存储当前层的节点值
            
            // 遍历当前层的所有节点
            while (size--) {
                TreeNode* f = que.front();  // 获取队首节点
                que.pop();  // 弹出队首节点
                
                temp.push_back(f->val);  // 将节点值加入当前层结果
                
                // 将左子节点入队（如果存在）
                if (f->left) que.push(f->left);
                // 将右子节点入队（如果存在）
                if (f->right) que.push(f->right);
            }
            
            ans.push_back(temp);  // 将当前层结果加入最终结果
        }
        
        // 反转结果，实现自底向上的顺序
        reverse(ans.begin(), ans.end());
        
        return ans;  // 返回自底向上的层序遍历结果
    }
};
```

**代码分析**：
- **与基础层序遍历的区别**：此代码与基础层序遍历的实现几乎相同，唯一的区别是在返回结果之前，使用 `reverse` 函数将结果向量反转，从而实现自底向上的顺序
- **reverse 操作**：`reverse(ans.begin(), ans.end())` 会将向量中的元素顺序完全反转，使得最后一层变成第一层，倒数第二层变成第二层，以此类推
- **时间复杂度影响**：反转操作的时间复杂度为 O(m)，其中 m 是树的高度，相对于整体 O(n) 的时间复杂度，这部分开销可以忽略不计

**执行过程示例**：
对于与上面相同的二叉树：
1. 首先执行与基础层序遍历相同的步骤，得到 `ans = [[3], [9, 20], [15, 7]]`
2. 执行 `reverse(ans.begin(), ans.end())` 后，`ans = [[15, 7], [9, 20], [3]]`
3. 返回反转后的结果

**技术要点**：
- 利用现有算法：先按正常顺序遍历，再反转结果，避免重新设计算法
- 高效实现：反转操作简单高效，时间复杂度低
- 代码复用：大部分代码与基础层序遍历相同，只增加了一行反转代码

#### 2.3 二叉树的右视图（LeetCode 199）

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
    /**
     * @brief 二叉树的右视图
     * @param root 二叉树的根节点指针
     * @return 从右侧能看到的节点值向量
     */
    vector<int> rightSideView(TreeNode* root) {
        vector<int> ans;  // 存储最终结果
        
        // 边界情况处理：空树直接返回空结果
        if (root == nullptr) return ans;
        
        // 创建队列用于层序遍历
        queue<TreeNode*> que;
        que.push(root);  // 将根节点入队
        
        // 当队列不为空时，继续遍历
        while (!que.empty()) {
            int size = que.size();  // 当前层的节点数
            
            // 遍历当前层的所有节点
            while (size--) {
                TreeNode* f = que.front();  // 获取队首节点
                que.pop();  // 弹出队首节点
                
                // 当为当前层的最后一个节点时，将其值加入结果
                if (size == 0) {
                    ans.push_back(f->val);
                }
                
                // 将左子节点入队（如果存在）
                if (f->left) que.push(f->left);
                // 将右子节点入队（如果存在）
                if (f->right) que.push(f->right);
            }
        }
        
        return ans;  // 返回右视图结果
    }
};
```

**代码分析**：
- **与基础层序遍历的区别**：
  - 结果类型不同：返回 `vector<int>` 而不是 `vector<vector<int>>`
  - 只记录每层的最后一个节点：当 `size == 0` 时，说明当前节点是当前层的最后一个节点，将其值加入结果
- **关键逻辑**：通过 `size--` 的方式，当 `size` 减到 0 时，当前处理的节点就是当前层的最后一个节点，也就是从右侧能看到的节点
- **子节点入队顺序**：仍然按照左子节点先入队，右子节点后入队的顺序，这样可以确保每层的节点按从左到右的顺序处理

**执行过程示例**：
对于二叉树：
```
    1
   / 
  2   3
    \   
     5   4
```

1. 初始化：`que = [1]`，`ans = []`
2. 第一次外层循环：`size = 1`
   - 内层循环：`size = 1`
     - 弹出 1，`size` 变为 0，将 1 加入 `ans`，`ans = [1]`
     - 入队 2, 3，`que = [2, 3]`
3. 第二次外层循环：`size = 2`
   - 内层循环：`size = 2`
     - 弹出 2，`size` 变为 1，不是最后一个节点
     - 入队 5，`que = [3, 5]`
     - 弹出 3，`size` 变为 0，将 3 加入 `ans`，`ans = [1, 3]`
     - 入队 4，`que = [5, 4]`
4. 第三次外层循环：`size = 2`
   - 内层循环：`size = 2`
     - 弹出 5，`size` 变为 1，不是最后一个节点
     - 无左右子节点，`que = [4]`
     - 弹出 4，`size` 变为 0，将 4 加入 `ans`，`ans = [1, 3, 4]`
     - 无左右子节点，`que = []`
5. 队列为空，结束遍历，返回 `ans`

**技术要点**：
- 层的最后节点：通过记录每层的大小，识别并记录每层的最后一个节点
- 顺序处理：保持左子节点先入队的顺序，确保正确处理每层节点的顺序
- 结果存储：使用一维向量存储结果，只保留每层的最后一个节点

### 3. 递归实现层序遍历

除了迭代实现外，层序遍历也可以通过递归来实现：

```cpp
class Solution {
public:
    /**
     * @brief 递归辅助函数，实现层序遍历
     * @param cur 当前遍历的节点指针
     * @param result 存储遍历结果的二维向量引用
     * @param depth 当前节点的深度（从0开始）
     */
    void order(TreeNode* cur, vector<vector<int>>& result, int depth) {
        // 递归终止条件：当前节点为空
        if (cur == nullptr) return;
        
        // 如果当前深度对应的层还未创建，创建新层
        if (result.size() == depth) {
            result.push_back(vector<int>());
        }
        
        // 将当前节点值加入对应深度的层
        result[depth].push_back(cur->val);
        
        // 递归遍历左子树，深度+1
        order(cur->left, result, depth + 1);
        // 递归遍历右子树，深度+1
        order(cur->right, result, depth + 1);
    }
    
    /**
     * @brief 层序遍历入口函数
     * @param root 二叉树的根节点指针
     * @return 按层存储的节点值二维向量
     */
    vector<vector<int>> levelOrder(TreeNode* root) {
        vector<vector<int>> result;  // 存储遍历结果
        int depth = 0;  // 初始深度为0
        order(root, result, depth);  // 调用递归辅助函数
        return result;  // 返回遍历结果
    }
};
```

**代码分析**：
- **递归原理**：通过记录节点的深度，将同一深度的节点值存储在结果向量的同一层中
- **深度管理**：每次递归调用时，深度加 1，确保子节点存储在正确的层中
- **层的创建**：当 `result.size() == depth` 时，说明当前深度对应的层还未创建，需要创建新层
- **遍历顺序**：先递归左子树，后递归右子树，确保同一层的节点按从左到右的顺序存储

**优缺点分析**：
- **优点**：代码简洁，逻辑清晰
- **缺点**：对于深度较大的二叉树，可能会导致栈溢出
- **时间复杂度**：O(n)，每个节点都会被访问一次
- **空间复杂度**：O(n)，最坏情况下（树退化为链表）递归栈深度为 n

**技术要点**：
- 深度参数：通过深度参数将节点映射到对应的层
- 层的动态创建：根据深度动态创建新层，确保结果结构正确
- 递归终止条件：处理空节点，避免无限递归

---

## 三、 相关题目解析

### 1. LeetCode 102. 二叉树的层序遍历

#### 1.1 题目描述

给你二叉树的根节点 `root`，返回其节点值的 **层序遍历**。（即逐层地，从左到右访问所有节点）。

#### 1.2 解题思路

- **核心思路**：使用队列实现广度优先搜索，按层遍历二叉树
- **步骤**：
  1. 初始化队列，将根节点入队
  2. 当队列不为空时，记录当前队列大小（当前层节点数）
  3. 遍历当前层的所有节点，将节点值加入结果集，并将其左右子节点入队
  4. 重复步骤2-3，直到队列为空

#### 1.3 复杂度分析

- **时间复杂度**：O(n)，其中 n 是二叉树的节点数。每个节点都会被访问一次
- **空间复杂度**：O(n)，最坏情况下（完全二叉树），队列最多存储 n/2 个节点

#### 1.4 代码实现

见上文基础层序遍历代码实现。

### 2. LeetCode 107. 二叉树的层序遍历 II

#### 2.1 题目描述

给你二叉树的根节点 `root`，返回其节点值 **自底向上的层序遍历**。（即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

#### 2.2 解题思路

- **核心思路**：先按正常层序遍历得到结果，然后将结果反转
- **步骤**：
  1. 按照正常层序遍历的方法获取结果
  2. 使用 `reverse` 函数将结果反转，得到自底向上的顺序

#### 2.3 复杂度分析

- **时间复杂度**：O(n)，其中 n 是二叉树的节点数。每个节点都会被访问一次，反转操作的时间复杂度为 O(m)，其中 m 是树的高度
- **空间复杂度**：O(n)，最坏情况下（完全二叉树），队列最多存储 n/2 个节点

#### 2.4 代码实现

见上文自底向上的层序遍历代码实现。

### 3. LeetCode 199. 二叉树的右视图

#### 3.1 题目描述

给定一个二叉树的 根节点 `root`，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

#### 3.2 解题思路

- **核心思路**：层序遍历，记录每层的最后一个节点
- **步骤**：
  1. 初始化队列，将根节点入队
  2. 当队列不为空时，记录当前队列大小（当前层节点数）
  3. 遍历当前层的所有节点，当遍历到最后一个节点时，将其值加入结果集
  4. 将当前节点的左右子节点入队（如果存在）
  5. 重复步骤2-4，直到队列为空

#### 3.3 复杂度分析

- **时间复杂度**：O(n)，其中 n 是二叉树的节点数。每个节点都会被访问一次
- **空间复杂度**：O(n)，最坏情况下（完全二叉树），队列最多存储 n/2 个节点

#### 3.4 代码实现

见上文二叉树的右视图代码实现。

---

## 四、 相关题目之间的联系与区别

### 1. 联系

- **核心算法**：三个题目都基于层序遍历（广度优先搜索），使用队列实现
- **实现思路**：都需要按层处理节点，记录每层的节点数
- **时间空间复杂度**：三个题目的时间复杂度都是 O(n)，空间复杂度都是 O(n)

### 2. 区别

| 题目 | 核心要求 | 特殊处理 |
|------|---------|----------|
| 102. 二叉树的层序遍历 | 按层返回节点值 | 直接返回层序遍历结果 |
| 107. 二叉树的层序遍历 II | 自底向上返回节点值 | 反转层序遍历结果 |
| 199. 二叉树的右视图 | 返回每层最右侧节点 | 只记录每层的最后一个节点 |

### 3. 扩展思考

- **层序遍历的通用性**：层序遍历是一种通用的遍历方法，可以通过不同的处理逻辑适应不同的问题需求
- **变种问题**：除了这三个题目外，还有许多基于层序遍历的变种问题，如计算每层的平均值、找到每层的最大值等
- **算法优化**：对于特定问题，可以针对层序遍历进行优化，如提前终止（如最小深度问题）

---

## 五、 学习过程中的注意事项与常见错误

### 1. 注意事项

1. **队列操作**：
   - 入队操作：使用 `push` 方法
   - 出队操作：使用 `front` 获取队首元素，然后使用 `pop` 弹出
   - 队列大小：使用 `size` 方法获取当前队列大小

2. **层的处理**：
   - 必须在循环开始时记录队列大小，因为队列大小会随节点入队而变化
   - 使用嵌套循环处理每层的节点

3. **边界情况**：
   - 空树处理：当根节点为 `nullptr` 时，直接返回空结果
   - 单节点树：正常处理，返回包含该节点的单层结果

4. **子节点入队**：
   - 只有当子节点不为 `nullptr` 时才入队，避免处理空指针

### 2. 常见错误

1. **队列大小错误**：
   - **错误**：在循环中使用 `que.size()` 作为循环条件，而不是在循环开始时记录大小
   - **后果**：由于队列大小会随节点入队而变化，导致循环次数错误
   - **解决方法**：在循环开始时记录队列大小，使用该大小作为循环条件

2. **子节点处理错误**：
   - **错误**：不检查子节点是否为 `nullptr` 就入队
   - **后果**：可能导致空指针异常
   - **解决方法**：在入队前检查子节点是否为 `nullptr`

3. **边界情况处理错误**：
   - **错误**：忘记处理空树情况
   - **后果**：当输入为空树时，程序可能会崩溃
   - **解决方法**：在函数开始时检查根节点是否为 `nullptr`

4. **结果存储错误**：
   - **错误**：对于 107 题，忘记反转结果
   - **后果**：返回的是正常层序遍历结果，而不是自底向上的结果
   - **解决方法**：在返回前使用 `reverse` 函数反转结果

5. **右视图逻辑错误**：
   - **错误**：对于 199 题，错误地只考虑右子节点
   - **后果**：当某层最右侧节点在左子树时，会遗漏该节点
   - **解决方法**：按正常层序遍历，记录每层的最后一个节点

6. **递归深度错误**：
   - **错误**：在递归实现中，深度参数传递错误
   - **后果**：节点存储在错误的层中
   - **解决方法**：确保每次递归调用时深度参数正确递增

---

## 六、 面试八股内容

### 1. 基础概念类问题

#### 1.1 什么是二叉树的层序遍历？
**答案**：二叉树的层序遍历是按照从上到下、从左到右的顺序逐层访问二叉树中的所有节点，是广度优先搜索（BFS）在二叉树上的应用。层序遍历的核心思想是使用队列来保证节点处理的顺序，确保同一层的节点按顺序处理完毕后，再处理下一层的节点。

**深入理解**：层序遍历与深度优先遍历（DFS）不同，它不是先深入到树的叶子节点，而是按照树的层次结构，从根节点开始，逐层向下遍历。这种遍历方式更符合人类对树结构的直观理解，也便于解决与树的层次相关的问题。

#### 1.2 层序遍历和深度优先遍历有什么区别？
**答案**：
| 特性 | 层序遍历（BFS） | 深度优先遍历（DFS） |
|------|----------------|-------------------|
| 实现方式 | 使用队列 | 使用栈（递归或显式栈） |
| 遍历顺序 | 按层从上到下，每层从左到右 | 先深入到叶子节点，再回溯 |
| 适用场景 | 层次相关问题（如树的高度、宽度） | 路径相关问题（如前序、中序、后序遍历） |
| 空间复杂度 | O(n)，最坏情况为完全二叉树的最后一层节点数 | O(n)，最坏情况为树退化为链表时的递归栈深度 |
| 时间复杂度 | O(n)，每个节点访问一次 | O(n)，每个节点访问一次 |

**面试技巧**：在回答时，可以举例说明两种遍历方式的不同，例如对于一个简单的二叉树，分别说明两种遍历的顺序和结果，这样会更直观。

#### 1.3 层序遍历的时间复杂度和空间复杂度是多少？
**答案**：
- **时间复杂度**：O(n)，其中 n 是二叉树的节点数。每个节点都会被访问一次，入队和出队操作都是 O(1) 时间复杂度。
- **空间复杂度**：O(n)，最坏情况下（完全二叉树），队列最多存储 n/2 个节点（最后一层的节点数）。对于平衡二叉树，空间复杂度为 O(n)；对于斜树（退化为链表），空间复杂度为 O(1)。

**深入理解**：空间复杂度的分析需要考虑树的结构。在完全二叉树中，最后一层的节点数接近总节点数的一半，因此队列在处理最后一层时会达到最大容量。而在斜树中，队列中最多只有一个节点，因此空间复杂度为 O(1)。

### 2. 实现类问题

#### 2.1 如何用队列实现二叉树的层序遍历？
**答案**：
1. **初始化**：创建一个队列，并将根节点入队。
2. **循环处理**：当队列不为空时，执行以下操作：
   - **记录层大小**：获取当前队列的大小，这表示当前层的节点数。
   - **处理当前层**：遍历当前层的所有节点，依次取出队首节点，将其值加入结果集，并将其左右子节点（如果存在）入队。
3. **结束条件**：当队列为空时，遍历结束，返回结果。

**代码框架**：
```cpp
vector<vector<int>> levelOrder(TreeNode* root) {
    vector<vector<int>> ans;
    if (!root) return ans;
    queue<TreeNode*> q;
    q.push(root);
    while (!q.empty()) {
        int size = q.size();
        vector<int> level;
        for (int i = 0; i < size; i++) {
            TreeNode* node = q.front();
            q.pop();
            level.push_back(node->val);
            if (node->left) q.push(node->left);
            if (node->right) q.push(node->right);
        }
        ans.push_back(level);
    }
    return ans;
}
```

**面试技巧**：在实现时，要注意先记录队列大小，再进行循环处理，因为队列大小会随着子节点的入队而变化。同时，要处理空树的边界情况。

#### 2.2 如何实现自底向上的层序遍历？
**答案**：
1. **正常层序遍历**：按照标准层序遍历的方法获取结果，得到从上到下的层次顺序。
2. **反转结果**：使用 `reverse` 函数将结果向量反转，得到自底向上的顺序。

**代码框架**：
```cpp
vector<vector<int>> levelOrderBottom(TreeNode* root) {
    vector<vector<int>> ans;
    if (!root) return ans;
    queue<TreeNode*> q;
    q.push(root);
    while (!q.empty()) {
        int size = q.size();
        vector<int> level;
        for (int i = 0; i < size; i++) {
            TreeNode* node = q.front();
            q.pop();
            level.push_back(node->val);
            if (node->left) q.push(node->left);
            if (node->right) q.push(node->right);
        }
        ans.push_back(level);
    }
    reverse(ans.begin(), ans.end());
    return ans;
}
```

**深入理解**：这种方法的时间复杂度仍然是 O(n)，因为反转操作的时间复杂度为 O(m)，其中 m 是树的高度，相对于整体 O(n) 的时间复杂度可以忽略不计。空间复杂度也仍然是 O(n)。

#### 2.3 如何实现二叉树的右视图？
**答案**：
1. **层序遍历**：使用队列进行层序遍历。
2. **记录每层最后节点**：在遍历每层节点时，记录当前层的最后一个节点，将其值加入结果集。

**代码框架**：
```cpp
vector<int> rightSideView(TreeNode* root) {
    vector<int> ans;
    if (!root) return ans;
    queue<TreeNode*> q;
    q.push(root);
    while (!q.empty()) {
        int size = q.size();
        for (int i = 0; i < size; i++) {
            TreeNode* node = q.front();
            q.pop();
            if (i == size - 1) { // 当前层的最后一个节点
                ans.push_back(node->val);
            }
            if (node->left) q.push(node->left);
            if (node->right) q.push(node->right);
        }
    }
    return ans;
}
```

**面试技巧**：实现右视图时，要注意不是只考虑右子节点，而是要记录每层的最后一个节点，因为最右侧的节点可能在左子树中。例如，当某层只有左子节点时，最右侧的节点就是左子节点。

#### 2.4 如何用递归实现层序遍历？
**答案**：
1. **递归函数设计**：设计一个递归函数，参数包括当前节点、结果向量和当前深度。
2. **深度管理**：在递归过程中，记录节点的深度，将同一深度的节点值存储在结果向量的同一层中。
3. **层的创建**：当当前深度对应的层还未创建时，创建新层。

**代码框架**：
```cpp
void dfs(TreeNode* node, vector<vector<int>>& ans, int depth) {
    if (!node) return;
    if (depth == ans.size()) {
        ans.push_back({});
    }
    ans[depth].push_back(node->val);
    dfs(node->left, ans, depth + 1);
    dfs(node->right, ans, depth + 1);
}

vector<vector<int>> levelOrder(TreeNode* root) {
    vector<vector<int>> ans;
    dfs(root, ans, 0);
    return ans;
}
```

**深入理解**：递归实现层序遍历的核心是通过深度参数来管理节点的层次，确保同一深度的节点存储在同一层中。这种方法的代码更简洁，但对于深度较大的二叉树可能会导致栈溢出。

### 3. 扩展问题

#### 3.1 除了队列，还可以用什么数据结构实现层序遍历？
**答案**：
- **递归**：通过记录节点的深度，将节点值加入对应深度的结果集中。
- **双端队列**：在某些情况下，可以使用双端队列来实现层序遍历，特别是在需要同时处理两端节点的场景。
- **数组**：对于完全二叉树，可以使用数组来模拟层序遍历，利用完全二叉树的性质（第 i 个节点的左子节点是 2i+1，右子节点是 2i+2）。

**深入理解**：不同的数据结构适用于不同的场景。队列是最常用的实现方式，因为它天然支持先进先出的特性，符合层序遍历的需求。递归实现代码更简洁，但可能存在栈溢出的风险。

#### 3.2 层序遍历有哪些应用场景？
**答案**：
- **二叉树的按层打印**：将二叉树按层打印出来，便于可视化树结构。
- **二叉树的层次相关问题**：
  - 计算树的最大深度和最小深度
  - 计算树的宽度（每层的最大节点数）
  - 找到树的某一层的所有节点
- **广度优先搜索算法**：层序遍历是图的广度优先搜索的特例，可用于解决图的最短路径问题。
- **二叉树的序列化和反序列化**：层序遍历可以用于二叉树的序列化和反序列化，便于存储和传输。
- **检查二叉树的性质**：例如判断二叉树是否是完全二叉树、满二叉树等。

**面试技巧**：在回答应用场景时，可以举例说明具体的问题，例如如何使用层序遍历计算树的最大深度，这样会更具体和有说服力。

#### 3.3 如何在层序遍历中处理每层的信息？
**答案**：通过在每层开始时记录队列的大小，确保只处理当前层的节点，然后将它们的子节点入队。具体步骤如下：
1. 在每次循环开始时，获取当前队列的大小 `size`，这表示当前层的节点数。
2. 然后进行 `size` 次循环，处理当前层的所有节点。
3. 在处理每个节点时，将其左右子节点入队。
4. 这样可以确保在处理完当前层的所有节点后，队列中恰好存储了下一层的所有节点。

**深入理解**：这种方法的关键是在每层开始时记录队列的大小，因为队列的大小会随着子节点的入队而变化。如果不记录大小，直接使用 `q.size()` 作为循环条件，会导致循环次数错误，因为队列大小在循环过程中会发生变化。

#### 3.4 如何判断一个二叉树是否是完全二叉树？
**答案**：可以使用层序遍历来判断一个二叉树是否是完全二叉树。具体步骤如下：
1. 使用队列进行层序遍历。
2. 当遇到第一个空节点时，标记为已遇到空节点。
3. 在遇到空节点之后，如果再遇到非空节点，则不是完全二叉树。
4. 如果遍历结束后，没有出现上述情况，则是完全二叉树。

**代码框架**：
```cpp
bool isCompleteTree(TreeNode* root) {
    if (!root) return true;
    queue<TreeNode*> q;
    q.push(root);
    bool encounteredNull = false;
    while (!q.empty()) {
        TreeNode* node = q.front();
        q.pop();
        if (!node) {
            encounteredNull = true;
        } else {
            if (encounteredNull) {
                return false;
            }
            q.push(node->left);
            q.push(node->right);
        }
    }
    return true;
}
```

**深入理解**：完全二叉树的定义是除了最后一层外，每一层都被完全填满，并且最后一层的节点都靠左排列。使用层序遍历可以有效地检查这个性质，因为层序遍历按照从左到右的顺序访问节点，一旦遇到空节点，后面的节点都应该是空的。

#### 3.5 如何在层序遍历中同时记录节点的父节点信息？
**答案**：可以在队列中存储节点及其父节点的信息，例如使用 `pair<TreeNode*, TreeNode*>` 来存储节点和其父节点。具体步骤如下：
1. 初始化队列，将根节点和 nullptr（表示根节点没有父节点）入队。
2. 当队列不为空时，取出队首元素，处理当前节点和其父节点的关系。
3. 将当前节点的左右子节点和当前节点（作为父节点）入队。

**代码框架**：
```cpp
void levelOrderWithParent(TreeNode* root) {
    if (!root) return;
    queue<pair<TreeNode*, TreeNode*>> q;
    q.push({root, nullptr});
    while (!q.empty()) {
        auto [node, parent] = q.front();
        q.pop();
        // 处理节点和父节点的关系
        if (parent) {
            cout << "Node " << node->val << "'s parent is " << parent->val << endl;
        } else {
            cout << "Node " << node->val << " is root" << endl;
        }
        if (node->left) {
            q.push({node->left, node});
        }
        if (node->right) {
            q.push({node->right, node});
        }
    }
}
```

**面试技巧**：这种方法可以用于解决需要父节点信息的问题，例如找到节点的所有祖先、构建父指针等。在面试中，要注意使用合适的数据结构来存储节点和其父节点的信息。

### 4. 高级问题

#### 4.1 如何实现二叉树的之字形层序遍历？
**答案**：二叉树的之字形层序遍历是指奇数层从左到右遍历，偶数层从右到左遍历。可以使用双端队列来实现：
1. 使用双端队列进行层序遍历。
2. 对于奇数层，从队列头部取出节点，将子节点从尾部入队（先左后右）。
3. 对于偶数层，从队列尾部取出节点，将子节点从头部入队（先右后左）。

**代码框架**：
```cpp
vector<vector<int>> zigzagLevelOrder(TreeNode* root) {
    vector<vector<int>> ans;
    if (!root) return ans;
    deque<TreeNode*> dq;
    dq.push_back(root);
    bool leftToRight = true;
    while (!dq.empty()) {
        int size = dq.size();
        vector<int> level;
        for (int i = 0; i < size; i++) {
            TreeNode* node;
            if (leftToRight) {
                node = dq.front();
                dq.pop_front();
                if (node->left) dq.push_back(node->left);
                if (node->right) dq.push_back(node->right);
            } else {
                node = dq.back();
                dq.pop_back();
                if (node->right) dq.push_front(node->right);
                if (node->left) dq.push_front(node->left);
            }
            level.push_back(node->val);
        }
        ans.push_back(level);
        leftToRight = !leftToRight;
    }
    return ans;
}
```

**深入理解**：之字形层序遍历的关键是根据当前层的奇偶性来决定遍历的方向和子节点的入队顺序。使用双端队列可以方便地从两端进行操作，实现不同方向的遍历。

#### 4.2 如何在层序遍历中处理节点的层次信息？
**答案**：可以在队列中存储节点及其层次信息，例如使用 `pair<TreeNode*, int>` 来存储节点和其所在的层次。具体步骤如下：
1. 初始化队列，将根节点和层次 0 入队。
2. 当队列不为空时，取出队首元素，处理当前节点和其层次。
3. 将当前节点的左右子节点和层次 + 1 入队。

**代码框架**：
```cpp
void levelOrderWithLevel(TreeNode* root) {
    if (!root) return;
    queue<pair<TreeNode*, int>> q;
    q.push({root, 0});
    while (!q.empty()) {
        auto [node, level] = q.front();
        q.pop();
        // 处理节点和层次
        cout << "Level " << level << ": " << node->val << endl;
        if (node->left) {
            q.push({node->left, level + 1});
        }
        if (node->right) {
            q.push({node->right, level + 1});
        }
    }
}
```

**面试技巧**：这种方法可以用于解决需要层次信息的问题，例如找到某一层次的所有节点、计算每层的平均值等。在面试中，要注意选择合适的数据结构来存储节点和层次信息。

#### 4.3 如何优化层序遍历的空间复杂度？
**答案**：对于某些特定的问题，可以通过优化空间复杂度来提高算法效率：
1. **使用迭代器或生成器**：对于大规模二叉树，可以使用迭代器或生成器模式，按需生成节点值，减少内存使用。
2. **使用父指针**：如果树的节点有父指针，可以利用父指针来实现层序遍历，而不需要使用队列。
3. **利用树的性质**：对于完全二叉树，可以使用数组来模拟层序遍历，利用完全二叉树的性质（第 i 个节点的左子节点是 2i+1，右子节点是 2i+2）。

**深入理解**：空间复杂度的优化需要根据具体的问题和树的结构来选择合适的方法。例如，对于完全二叉树，使用数组可以将空间复杂度降低到 O(1)，因为可以直接通过索引计算节点的位置。

### 5. 面试技巧

#### 5.1 如何在面试中回答层序遍历相关的问题？
**答案**：
1. **理解问题**：首先要明确问题的要求，是标准层序遍历还是其变种（如自底向上、右视图等）。
2. **选择方法**：根据问题要求选择合适的实现方法，通常使用队列实现迭代版本。
3. **代码实现**：写出清晰、正确的代码，注意边界情况的处理。
4. **分析复杂度**：分析时间复杂度和空间复杂度，并解释原因。
5. **优化思路**：如果可能，提出优化方案，例如空间复杂度的优化。
6. **测试示例**：使用示例二叉树验证代码的正确性。

**面试技巧**：在面试中，要保持思路清晰，代码结构合理，注释充分。如果遇到问题，可以先在脑海中或纸上模拟算法的执行过程，确保逻辑正确。

#### 5.2 层序遍历相关的常见陷阱
**答案**：
- **队列大小错误**：在循环中使用 `q.size()` 作为循环条件，而不是在循环开始时记录大小，导致循环次数错误。
- **空指针处理**：不检查子节点是否为 `nullptr` 就入队，导致空指针异常。
- **边界情况**：忘记处理空树的情况，导致程序崩溃。
- **结果存储**：对于自底向上的层序遍历，忘记反转结果，导致返回的是正常顺序的结果。
- **右视图逻辑**：错误地只考虑右子节点，而不是记录每层的最后一个节点，导致遗漏某些情况。

**面试技巧**：在面试中，要注意避免这些常见陷阱，确保代码的正确性和鲁棒性。可以通过测试边界情况来验证代码的正确性。

---

## 七、 代码优化与扩展

### 1. 代码优化

1. **空间优化**：
   - 对于大规模二叉树，可以考虑使用迭代器或生成器模式，按需生成节点值，减少内存使用
   - 对于递归实现，可以使用尾递归优化，减少栈空间使用

2. **时间优化**：
   - 对于频繁访问的二叉树，可以考虑缓存层序遍历结果
   - 对于平衡二叉树，可以利用其特性优化遍历过程

3. **可读性优化**：
   - 使用命名清晰的变量和函数
   - 添加详细的注释
   - 提取重复代码为辅助函数

### 2. 扩展应用

1. **二叉树的最大深度**：
   ```cpp
   int maxDepth(TreeNode* root) {
       if (root == nullptr) return 0;
       queue<TreeNode*> que;
       que.push(root);
       int depth = 0;
       while (!que.empty()) {
           int size = que.size();
           depth++;
           while (size--) {
               TreeNode* node = que.front();
               que.pop();
               if (node->left) que.push(node->left);
               if (node->right) que.push(node->right);
           }
       }
       return depth;
   }
   ```

2. **二叉树的最小深度**：
   ```cpp
   int minDepth(TreeNode* root) {
       if (root == nullptr) return 0;
       queue<TreeNode*> que;
       que.push(root);
       int depth = 0;
       while (!que.empty()) {
           int size = que.size();
           depth++;
           while (size--) {
               TreeNode* node = que.front();
               que.pop();
               if (node->left == nullptr && node->right == nullptr) {
                   return depth;
               }
               if (node->left) que.push(node->left);
               if (node->right) que.push(node->right);
           }
       }
       return depth;
   }
   ```

3. **N叉树的层序遍历**：
   ```cpp
   vector<vector<int>> levelOrder(Node* root) {
       vector<vector<int>> ans;
       if (root == nullptr) return ans;
       queue<Node*> que;
       que.push(root);
       while (!que.empty()) {
           int size = que.size();
           vector<int> temp;
           while (size--) {
               Node* node = que.front();
               que.pop();
               temp.push_back(node->val);
               for (auto child : node->children) {
                   if (child) que.push(child);
               }
           }
           ans.push_back(temp);
       }
       return ans;
   }
   ```

4. **二叉树的层平均值**：
   ```cpp
   vector<double> averageOfLevels(TreeNode* root) {
       vector<double> ans;
       if (root == nullptr) return ans;
       queue<TreeNode*> que;
       que.push(root);
       while (!que.empty()) {
           int size = que.size();
           double sum = 0;
           for (int i = 0; i < size; i++) {
               TreeNode* node = que.front();
               que.pop();
               sum += node->val;
               if (node->left) que.push(node->left);
               if (node->right) que.push(node->right);
           }
           ans.push_back(sum / size);
       }
       return ans;
   }
   ```

5. **二叉树的每层最大值**：
   ```cpp
   vector<int> largestValues(TreeNode* root) {
       vector<int> ans;
       if (root == nullptr) return ans;
       queue<TreeNode*> que;
       que.push(root);
       while (!que.empty()) {
           int size = que.size();
           int maxVal = INT_MIN;
           for (int i = 0; i < size; i++) {
               TreeNode* node = que.front();
               que.pop();
               maxVal = max(maxVal, node->val);
               if (node->left) que.push(node->left);
               if (node->right) que.push(node->right);
           }
           ans.push_back(maxVal);
       }
       return ans;
   }
   ```

---

## 八、 总结

### 1. 核心知识点

- **层序遍历原理**：使用队列实现广度优先搜索，按层访问节点
- **基本步骤**：初始化队列 → 记录层大小 → 处理当前层节点 → 入队子节点 → 重复直到队列为空
- **相关题目**：
  - 102. 二叉树的层序遍历：直接返回层序遍历结果
  - 107. 二叉树的层序遍历 II：反转层序遍历结果
  - 199. 二叉树的右视图：记录每层最后一个节点

### 2. 学习收获

- 掌握了二叉树层序遍历的原理和实现方法
- 理解了广度优先搜索在二叉树上的应用
- 学会了如何处理层序遍历相关的变种问题
- 掌握了队列的基本操作和应用场景
- 了解了层序遍历的扩展应用和优化方法

### 3. 应用价值

- **算法学习**：层序遍历是广度优先搜索的基础，理解它有助于学习更复杂的图算法
- **实际应用**：在树的层次相关问题中，层序遍历是一种高效的解决方案
- **面试准备**：层序遍历是面试中的常见考点，掌握它有助于应对相关问题

### 4. 后续学习建议

- 学习图的广度优先搜索算法
- 探索层序遍历的其他应用场景
- 研究其他树的遍历算法（如Morris遍历）
- 练习更多层序遍历相关的题目，如：
  - 637. 二叉树的层平均值
  - 515. 在每个树行中找最大值
  - 116. 填充每个节点的下一个右侧节点指针
  - 117. 填充每个节点的下一个右侧节点指针 II
  - 104. 二叉树的最大深度
  - 111. 二叉树的最小深度

通过本课程的学习，我们掌握了二叉树层序遍历的原理和实现方法，理解了其在解决相关问题中的应用。希望这份学习笔记能帮助你更好地理解和应用这些知识点！