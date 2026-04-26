# 二叉树层序遍历扩展学习笔记

> **学习内容**：二叉树的层序遍历（广度优先搜索）及其扩展应用
> **学习资源**：代码随想录二叉树层序遍历部分、LeetCode相关题目
> **相关题目**：
> - 102. 二叉树的层序遍历
> - 107. 二叉树的层序遍历 II
> - 199. 二叉树的右视图
> - 637. 二叉树的层平均值
> - 429. N叉树的层序遍历
> - 515. 在每个树行中找最大值
> - 116. 填充每个节点的下一个右侧节点指针
> - 117. 填充每个节点的下一个右侧节点指针 II
> - 104. 二叉树的最大深度
> - 111. 二叉树的最小深度
> **学习目标**：掌握二叉树层序遍历的原理和实现方法，理解其时间空间复杂度，以及相关题目的解题思路和扩展应用

---

## 一、 层序遍历核心原理

### 1. 基本概念

**层序遍历**（Level Order Traversal）是二叉树的一种遍历方式，它按照从上到下、从左到右的顺序逐层访问二叉树中的所有节点。

**广度优先搜索**（Breadth-First Search, BFS）是一种图遍历算法，从根节点开始，先访问根节点，然后访问其所有相邻节点，再依次访问这些节点的相邻节点，以此类推。二叉树的层序遍历本质上就是广度优先搜索的一种应用。

### 2. 实现原理

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

### 3. 算法实现

#### 3.1 队列实现（迭代法）

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
    vector<vector<int>> levelOrder(TreeNode* root) {
        vector<vector<int>> result;
        queue<TreeNode*> que;
        if (root != nullptr) que.push(root);
        
        while (!que.empty()) {
            int size = que.size();
            vector<int> vec;
            // 这里一定要使用固定大小size，不要使用que.size()，因为que.size是不断变化的
            for (int i = 0; i < size; i++) {
                TreeNode* node = que.front();
                que.pop();
                vec.push_back(node->val);
                if (node->left) que.push(node->left);
                if (node->right) que.push(node->right);
            }
            result.push_back(vec);
        }
        return result;
    }
};
```

**代码分析**：
- **边界情况处理**：首先检查根节点是否为 `nullptr`，如果是则直接返回空结果
- **队列初始化**：创建一个 `queue<TreeNode*>` 类型的队列，并将根节点入队
- **外层循环**：当队列不为空时，继续遍历
- **层的处理**：
  - 记录当前队列的大小 `size`，这表示当前层的节点数
  - 创建临时向量 `vec` 用于存储当前层的节点值
  - 内层循环遍历当前层的所有节点
- **节点处理**：
  - 获取并弹出队首节点
  - 将节点值加入当前层的临时向量
  - 将左子节点入队（如果存在）
  - 将右子节点入队（如果存在）
- **结果存储**：将当前层的临时向量加入最终结果

#### 3.2 递归实现

```cpp
class Solution {
public:
    void order(TreeNode* cur, vector<vector<int>>& result, int depth) {
        if (cur == nullptr) return;
        if (result.size() == depth) result.push_back(vector<int>());
        result[depth].push_back(cur->val);
        order(cur->left, result, depth + 1);
        order(cur->right, result, depth + 1);
    }
    
    vector<vector<int>> levelOrder(TreeNode* root) {
        vector<vector<int>> result;
        int depth = 0;
        order(root, result, depth);
        return result;
    }
};
```

**代码分析**：
- **递归函数设计**：设计一个递归函数 `order`，参数包括当前节点、结果向量和当前深度
- **递归终止条件**：当当前节点为 `nullptr` 时，直接返回
- **层的创建**：当 `result.size() == depth` 时，说明当前深度对应的层还未创建，需要创建新层
- **节点处理**：将当前节点值加入对应深度的层
- **递归调用**：递归遍历左子树和右子树，深度加 1
- **入口函数**：初始化结果向量和深度，调用递归函数

**优缺点分析**：
- **优点**：代码简洁，逻辑清晰
- **缺点**：对于深度较大的二叉树，可能会导致栈溢出
- **时间复杂度**：O(n)，每个节点都会被访问一次
- **空间复杂度**：O(n)，最坏情况下（树退化为链表）递归栈深度为 n

---

## 二、 时间与空间复杂度分析

### 1. 时间复杂度

对于层序遍历，无论是队列实现还是递归实现，时间复杂度都是 **O(n)**，其中 n 是二叉树的节点数。这是因为每个节点都会被访问一次，入队和出队操作都是 O(1) 时间复杂度。

**详细分析**：
- 每个节点都会被入队一次，出队一次，因此入队和出队操作的总时间复杂度是 O(n)
- 每个节点的值都会被处理一次，因此处理节点值的时间复杂度是 O(n)
- 因此，总的时间复杂度是 O(n)

### 2. 空间复杂度

空间复杂度的分析需要考虑树的结构，不同的树结构会导致不同的空间复杂度。

#### 2.1 队列实现的空间复杂度

队列实现的空间复杂度是 **O(n)**，最坏情况下（完全二叉树），队列最多存储 n/2 个节点（最后一层的节点数）。

**详细分析**：
- **完全二叉树**：最后一层的节点数接近总节点数的一半，因此队列在处理最后一层时会达到最大容量，空间复杂度为 O(n)
- **平衡二叉树**：每层的节点数相对均衡，队列的最大容量约为 n/2，空间复杂度为 O(n)
- **斜树**（退化为链表）：每层只有一个节点，队列的最大容量为 1，空间复杂度为 O(1)

#### 2.2 递归实现的空间复杂度

递归实现的空间复杂度是 **O(n)**，最坏情况下（树退化为链表）递归栈深度为 n。

**详细分析**：
- **完全二叉树**：递归栈的深度为树的高度，即 log(n)，空间复杂度为 O(log n)
- **平衡二叉树**：递归栈的深度为树的高度，即 log(n)，空间复杂度为 O(log n)
- **斜树**（退化为链表）：递归栈的深度为树的高度，即 n，空间复杂度为 O(n)

### 3. 复杂度比较

| 实现方式 | 时间复杂度 | 空间复杂度（最坏情况） | 空间复杂度（平均情况） |
|---------|-----------|-----------------------|-----------------------|
| 队列实现 | O(n) | O(n) | O(n) |
| 递归实现 | O(n) | O(n) | O(log n) |

**结论**：
- 队列实现的空间复杂度在各种情况下相对稳定，始终为 O(n)
- 递归实现的空间复杂度在树平衡时较好，但在树退化为链表时较差
- 对于深度较大的二叉树，递归实现可能会导致栈溢出，而队列实现则不会

---

## 三、 相关题目详细分析

### 1. LeetCode 102. 二叉树的层序遍历

#### 1.1 题目描述

给你二叉树的根节点 `root`，返回其按 **层序遍历** 得到的节点值。（即逐层地，从左到右访问所有节点）。

#### 1.2 解题思路

- **核心思路**：使用队列实现广度优先搜索，按层遍历二叉树
- **步骤**：
  1. 初始化队列，将根节点入队
  2. 当队列不为空时，记录当前队列大小（当前层节点数）
  3. 遍历当前层的所有节点，将节点值加入结果集，并将其左右子节点入队
  4. 重复步骤2-3，直到队列为空

#### 1.3 代码实现

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
    vector<vector<int>> levelOrder(TreeNode* root) {
        vector<vector<int>> result;
        queue<TreeNode*> que;
        if (root != nullptr) que.push(root);
        
        while (!que.empty()) {
            int size = que.size();
            vector<int> vec;
            // 这里一定要使用固定大小size，不要使用que.size()，因为que.size是不断变化的
            for (int i = 0; i < size; i++) {
                TreeNode* node = que.front();
                que.pop();
                vec.push_back(node->val);
                if (node->left) que.push(node->left);
                if (node->right) que.push(node->right);
            }
            result.push_back(vec);
        }
        return result;
    }
};
```

#### 1.4 代码分析

- **边界情况处理**：首先检查根节点是否为 `nullptr`，如果是则直接返回空结果，避免后续操作出现空指针异常
- **队列初始化**：创建一个 `queue<TreeNode*>` 类型的队列，并将根节点入队，作为遍历的起点
- **外层循环**：当队列不为空时，继续遍历，直到所有节点都被处理
- **层的处理**：
  - 记录当前队列的大小 `size`，这表示当前层的节点数，确保我们只处理当前层的节点
  - 创建临时向量 `vec` 用于存储当前层的节点值
  - 内层循环遍历当前层的所有节点
- **节点处理**：
  - 获取并弹出队首节点，注意弹出操作会从队列中移除该节点
  - 将节点值加入当前层的临时向量
  - 将左子节点入队（如果存在），确保左子节点先被处理
  - 将右子节点入队（如果存在），确保右子节点后被处理
- **结果存储**：将当前层的临时向量加入最终结果，完成一层的处理

#### 1.5 执行过程示例

对于二叉树：
```
    3
   / 
  9  20
    /  
   15   7
```

1. 初始化：`que = [3]`，`result = []`
2. 第一次外层循环：`size = 1`
   - 内层循环：`i = 0`
     - 弹出 3，`vec = [3]`
     - 入队 9, 20，`que = [9, 20]`
   - `result = [[3]]`
3. 第二次外层循环：`size = 2`
   - 内层循环：`i = 0`
     - 弹出 9，`vec = [9]`
     - 9 无左右子节点，`que = [20]`
   - 内层循环：`i = 1`
     - 弹出 20，`vec = [9, 20]`
     - 入队 15, 7，`que = [15, 7]`
   - `result = [[3], [9, 20]]`
4. 第三次外层循环：`size = 2`
   - 内层循环：`i = 0`
     - 弹出 15，`vec = [15]`
     - 15 无左右子节点，`que = [7]`
   - 内层循环：`i = 1`
     - 弹出 7，`vec = [15, 7]`
     - 7 无左右子节点，`que = []`
   - `result = [[3], [9, 20], [15, 7]]`
5. 队列为空，结束遍历，返回 `result`

#### 1.6 测试用例

**输入**：root = [3,9,20,null,null,15,7]
**输出**：[[3],[9,20],[15,7]]

**输入**：root = [1]
**输出**：[[1]]

**输入**：root = []
**输出**：[]

### 2. LeetCode 107. 二叉树的层序遍历 II

#### 2.1 题目描述

给你二叉树的根节点 `root`，返回其节点值 **自底向上的层序遍历**。（即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

#### 2.2 解题思路

- **核心思路**：先按正常层序遍历得到结果，然后将结果反转
- **步骤**：
  1. 按照正常层序遍历的方法获取结果，得到从上到下的层次顺序
  2. 使用 `reverse` 函数将结果向量反转，得到自底向上的顺序
  3. 返回反转后的结果

#### 2.3 代码实现

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
    vector<vector<int>> levelOrderBottom(TreeNode* root) {
        vector<vector<int>> result;
        queue<TreeNode*> que;
        if (root != nullptr) que.push(root);
        
        while (!que.empty()) {
            int size = que.size();
            vector<int> vec;
            for (int i = 0; i < size; i++) {
                TreeNode* node = que.front();
                que.pop();
                vec.push_back(node->val);
                if (node->left) que.push(node->left);
                if (node->right) que.push(node->right);
            }
            result.push_back(vec);
        }
        reverse(result.begin(), result.end());
        return result;
    }
};
```

#### 2.4 代码分析

- **与基础层序遍历的相同点**：
  - 使用队列实现广度优先搜索
  - 按层处理节点，记录每层的节点数
  - 将子节点入队，确保层序遍历的顺序
- **与基础层序遍历的不同点**：
  - 在返回结果之前，使用 `reverse` 函数将结果向量反转
  - 反转操作将从上到下的顺序转换为自底向上的顺序
- **reverse 操作**：
  - `reverse(result.begin(), result.end())` 会将向量中的元素顺序完全反转
  - 时间复杂度为 O(m)，其中 m 是树的高度，相对于整体 O(n) 的时间复杂度，这部分开销可以忽略不计

#### 2.5 执行过程示例

对于二叉树：
```
    3
   / 
  9  20
    /  
   15   7
```

1. 首先执行与基础层序遍历相同的步骤，得到 `result = [[3], [9, 20], [15, 7]]`
2. 执行 `reverse(result.begin(), result.end())` 后，`result = [[15, 7], [9, 20], [3]]`
3. 返回反转后的结果

#### 2.6 测试用例

**输入**：root = [3,9,20,null,null,15,7]
**输出**：[[15,7],[9,20],[3]]

**输入**：root = [1]
**输出**：[[1]]

**输入**：root = []
**输出**：[]

### 3. LeetCode 199. 二叉树的右视图

#### 3.1 题目描述

给定一个二叉树的 根节点 `root`，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

#### 3.2 解题思路

- **核心思路**：层序遍历，记录每层的最后一个节点
- **步骤**：
  1. 初始化队列，将根节点入队
  2. 当队列不为空时，记录当前队列大小（当前层节点数）
  3. 遍历当前层的所有节点，当遍历到最后一个节点时，将其值加入结果集
  4. 将当前节点的左右子节点入队（如果存在），确保左子节点先入队
  5. 重复步骤2-4，直到队列为空

#### 3.3 代码实现

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
    vector<int> rightSideView(TreeNode* root) {
        vector<int> result;
        queue<TreeNode*> que;
        if (root != nullptr) que.push(root);
        
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
};
```

#### 3.4 代码分析

- **结果类型**：返回 `vector<int>` 而不是 `vector<vector<int>>`，因为只需要每层的最后一个节点
- **关键逻辑**：
  - 通过 `i == size - 1` 判断当前节点是否是当前层的最后一个节点
  - 当遍历到最后一个节点时，将其值加入结果集
- **子节点入队顺序**：
  - 仍然按照左子节点先入队，右子节点后入队的顺序
  - 这样可以确保每层的节点按从左到右的顺序处理，最后一个节点就是最右侧的节点
- **边界情况处理**：
  - 当根节点为 `nullptr` 时，直接返回空结果

#### 3.5 执行过程示例

对于二叉树：
```
    1
   / 
  2   3
    \   
     5   4
```

1. 初始化：`que = [1]`，`result = []`
2. 第一次外层循环：`size = 1`
   - 内层循环：`i = 0`
     - 弹出 1，`i == 0 == size - 1`，将 1 加入 `result`，`result = [1]`
     - 入队 2, 3，`que = [2, 3]`
3. 第二次外层循环：`size = 2`
   - 内层循环：`i = 0`
     - 弹出 2，`i == 0 != size - 1`，不加入结果
     - 入队 5，`que = [3, 5]`
   - 内层循环：`i = 1`
     - 弹出 3，`i == 1 == size - 1`，将 3 加入 `result`，`result = [1, 3]`
     - 入队 4，`que = [5, 4]`
4. 第三次外层循环：`size = 2`
   - 内层循环：`i = 0`
     - 弹出 5，`i == 0 != size - 1`，不加入结果
     - 无左右子节点，`que = [4]`
   - 内层循环：`i = 1`
     - 弹出 4，`i == 1 == size - 1`，将 4 加入 `result`，`result = [1, 3, 4]`
     - 无左右子节点，`que = []`
5. 队列为空，结束遍历，返回 `result`

#### 3.6 测试用例

**输入**：root = [1,2,3,null,5,null,4]
**输出**：[1,3,4]

**输入**：root = [1,null,3]
**输出**：[1,3]

**输入**：root = []
**输出**：[]

### 4. LeetCode 637. 二叉树的层平均值

#### 4.1 题目描述

给定一个非空二叉树的根节点 `root`，以数组的形式返回每一层节点的平均值。与实际答案相差 10-5 以内的答案可以被接受。

#### 4.2 解题思路

- **核心思路**：层序遍历，计算每层节点的平均值
- **步骤**：
  1. 初始化队列，将根节点入队
  2. 当队列不为空时，记录当前队列大小（当前层节点数）
  3. 初始化当前层的总和为0
  4. 遍历当前层的所有节点，将节点值累加到总和中
  5. 计算当前层的平均值（总和除以节点数），将其加入结果集
  6. 将当前节点的左右子节点入队（如果存在）
  7. 重复步骤2-6，直到队列为空

#### 4.3 代码实现

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
    vector<double> averageOfLevels(TreeNode* root) {
        vector<double> result;
        queue<TreeNode*> que;
        if (root != nullptr) que.push(root);
        
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
            result.push_back(sum / size);
        }
        return result;
    }
};
```

#### 4.4 代码分析

- **结果类型**：返回 `vector<double>`，因为需要存储每层的平均值
- **数据类型选择**：
  - 使用 `double` 类型存储总和和平均值，确保计算精度
  - 避免使用 `int` 类型可能导致的整数除法问题
- **计算逻辑**：
  - 每层开始时初始化 `sum = 0`
  - 遍历每层的所有节点，将节点值累加到 `sum`
  - 遍历完成后，计算 `sum / size` 得到平均值
  - 将平均值加入结果集
- **边界情况处理**：
  - 当根节点为 `nullptr` 时，直接返回空结果
  - 当每层只有一个节点时，平均值就是该节点的值

#### 4.5 执行过程示例

对于二叉树：
```
    3
   / 
  9  20
    /  
   15   7
```

1. 初始化：`que = [3]`，`result = []`
2. 第一次外层循环：`size = 1`
   - `sum = 0`
   - 内层循环：`i = 0`
     - 弹出 3，`sum = 3`
     - 入队 9, 20，`que = [9, 20]`
   - 计算平均值：`3 / 1 = 3.0`，`result = [3.0]`
3. 第二次外层循环：`size = 2`
   - `sum = 0`
   - 内层循环：`i = 0`
     - 弹出 9，`sum = 9`
     - 无左右子节点，`que = [20]`
   - 内层循环：`i = 1`
     - 弹出 20，`sum = 29`
     - 入队 15, 7，`que = [15, 7]`
   - 计算平均值：`29 / 2 = 14.5`，`result = [3.0, 14.5]`
4. 第三次外层循环：`size = 2`
   - `sum = 0`
   - 内层循环：`i = 0`
     - 弹出 15，`sum = 15`
     - 无左右子节点，`que = [7]`
   - 内层循环：`i = 1`
     - 弹出 7，`sum = 22`
     - 无左右子节点，`que = []`
   - 计算平均值：`22 / 2 = 11.0`，`result = [3.0, 14.5, 11.0]`
5. 队列为空，结束遍历，返回 `result`

#### 4.6 测试用例

**输入**：root = [3,9,20,null,null,15,7]
**输出**：[3.00000,14.50000,11.00000]

**输入**：root = [1]
**输出**：[1.00000]

**输入**：root = [2147483647,2147483647,2147483647]
**输出**：[2147483647.00000,2147483647.00000]

### 5. LeetCode 429. N叉树的层序遍历

#### 5.1 题目描述

给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。

树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。

#### 5.2 解题思路

- **核心思路**：层序遍历，处理每个节点的所有子节点
- **步骤**：
  1. 初始化队列，将根节点入队
  2. 当队列不为空时，记录当前队列大小（当前层节点数）
  3. 创建临时向量用于存储当前层的节点值
  4. 遍历当前层的所有节点，将节点值加入临时向量
  5. 遍历当前节点的所有子节点，将它们入队（如果存在）
  6. 将临时向量加入结果集
  7. 重复步骤2-6，直到队列为空

#### 5.3 代码实现

```cpp
/*
// Definition for a Node.
class Node {
public:
    int val;
    vector<Node*> children;

    Node() {}

    Node(int _val) {
        val = _val;
    }

    Node(int _val, vector<Node*> _children) {
        val = _val;
        children = _children;
    }
};
*/

class Solution {
public:
    vector<vector<int>> levelOrder(Node* root) {
        vector<vector<int>> result;
        queue<Node*> que;
        if (root != nullptr) que.push(root);
        
        while (!que.empty()) {
            int size = que.size();
            vector<int> vec;
            for (int i = 0; i < size; i++) {
                Node* node = que.front();
                que.pop();
                vec.push_back(node->val);
                for (auto child : node->children) {
                    if (child != nullptr) que.push(child);
                }
            }
            result.push_back(vec);
        }
        return result;
    }
};
```

#### 5.4 代码分析

- **与二叉树层序遍历的异同**：
  - **相同点**：使用队列实现层序遍历，按层处理节点
  - **不同点**：二叉树每个节点只有左右两个子节点，而N叉树每个节点有多个子节点
- **子节点处理**：
  - 使用范围for循环（`for (auto child : node->children)`）遍历每个节点的子节点
  - 将每个非空的子节点入队
  - 入队顺序保持了从左到右的遍历顺序
- **边界情况处理**：
  - 当根节点为 `nullptr` 时，直接返回空结果
  - 当节点没有子节点时，跳过子节点入队步骤

#### 5.5 执行过程示例

对于N叉树：
```
      1
    / | \
   3  2  4
  / \
 5   6
```

1. 初始化：`que = [1]`，`result = []`
2. 第一次外层循环：`size = 1`
   - 内层循环：`i = 0`
     - 弹出 1，`vec = [1]`
     - 遍历子节点 3, 2, 4，入队后 `que = [3, 2, 4]`
   - `result = [[1]]`
3. 第二次外层循环：`size = 3`
   - 内层循环：`i = 0`
     - 弹出 3，`vec = [3]`
     - 遍历子节点 5, 6，入队后 `que = [2, 4, 5, 6]`
   - 内层循环：`i = 1`
     - 弹出 2，`vec = [3, 2]`
     - 无左右子节点，`que = [4, 5, 6]`
   - 内层循环：`i = 2`
     - 弹出 4，`vec = [3, 2, 4]`
     - 无左右子节点，`que = [5, 6]`
   - `result = [[1], [3, 2, 4]]`
4. 第三次外层循环：`size = 2`
   - 内层循环：`i = 0`
     - 弹出 5，`vec = [5]`
     - 无左右子节点，`que = [6]`
   - 内层循环：`i = 1`
     - 弹出 6，`vec = [5, 6]`
     - 无左右子节点，`que = []`
   - `result = [[1], [3, 2, 4], [5, 6]]`
5. 队列为空，结束遍历，返回 `result`

#### 5.6 测试用例

**输入**：root = [1,null,3,2,4,null,5,6]
**输出**：[[1],[3,2,4],[5,6]]

**输入**：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
**输出**：[[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]

**输入**：root = []
**输出**：[]

### 6. LeetCode 515. 在每个树行中找最大值

#### 6.1 题目描述

给定一棵二叉树的根节点 `root`，请找出该二叉树中每一层的最大值。

#### 6.2 解题思路

- **核心思路**：层序遍历，记录每层的最大值
- **步骤**：
  1. 初始化队列，将根节点入队
  2. 当队列不为空时，记录当前队列大小（当前层节点数）
  3. 初始化当前层的最大值为 `INT_MIN`
  4. 遍历当前层的所有节点，更新最大值
  5. 将最大值加入结果集
  6. 将当前节点的左右子节点入队（如果存在）
  7. 重复步骤2-6，直到队列为空

#### 6.3 代码实现

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
    vector<int> largestValues(TreeNode* root) {
        vector<int> result;
        queue<TreeNode*> que;
        if (root != nullptr) que.push(root);
        
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
            result.push_back(maxVal);
        }
        return result;
    }
};
```

#### 6.4 代码分析

- **结果类型**：返回 `vector<int>`，因为需要存储每层的最大值
- **最大值初始化**：
  - 使用 `INT_MIN` 作为初始值，确保即使节点值为负数也能正确比较
  - `INT_MIN` 是 C++ 中 `limits.h` 头文件定义的整数类型的最小值
- **最大值更新**：
  - 使用 `max(maxVal, node->val)` 函数更新最大值
  - 遍历每层的所有节点，确保找到该层的最大值
- **边界情况处理**：
  - 当根节点为 `nullptr` 时，直接返回空结果
  - 当每层只有一个节点时，该节点的值就是最大值

#### 6.5 执行过程示例

对于二叉树：
```
    1
   / 
  3   2
 / \   \
5   3   9
```

1. 初始化：`que = [1]`，`result = []`
2. 第一次外层循环：`size = 1`
   - `maxVal = INT_MIN`
   - 内层循环：`i = 0`
     - 弹出 1，`maxVal = max(INT_MIN, 1) = 1`
     - 入队 3, 2，`que = [3, 2]`
   - `result = [1]`
3. 第二次外层循环：`size = 2`
   - `maxVal = INT_MIN`
   - 内层循环：`i = 0`
     - 弹出 3，`maxVal = max(INT_MIN, 3) = 3`
     - 入队 5, 3，`que = [2, 5, 3]`
   - 内层循环：`i = 1`
     - 弹出 2，`maxVal = max(3, 2) = 3`
     - 入队 9，`que = [5, 3, 9]`
   - `result = [1, 3]`
4. 第三次外层循环：`size = 3`
   - `maxVal = INT_MIN`
   - 内层循环：`i = 0`
     - 弹出 5，`maxVal = max(INT_MIN, 5) = 5`
     - 无左右子节点，`que = [3, 9]`
   - 内层循环：`i = 1`
     - 弹出 3，`maxVal = max(5, 3) = 5`
     - 无左右子节点，`que = [9]`
   - 内层循环：`i = 2`
     - 弹出 9，`maxVal = max(5, 9) = 9`
     - 无左右子节点，`que = []`
   - `result = [1, 3, 9]`
5. 队列为空，结束遍历，返回 `result`

#### 6.6 测试用例

**输入**：root = [1,3,2,5,3,null,9]
**输出**：[1,3,9]

**输入**：root = [1,2,3]
**输出**：[1,3]

**输入**：root = [9,-1,-2]
**输出**：[9,-1]

### 7. LeetCode 116. 填充每个节点的下一个右侧节点指针

#### 7.1 题目描述

给定一个 **完美二叉树**，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

```
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
```

填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。

#### 7.2 解题思路

- **核心思路**：层序遍历，连接每层的节点
- **步骤**：
  1. 初始化队列，将根节点入队
  2. 当队列不为空时，记录当前队列大小（当前层节点数）
  3. 初始化前一个节点指针 `prev` 为 `nullptr`
  4. 遍历当前层的所有节点：
     - 如果 `prev` 不为 `nullptr`，将 `prev->next` 指向当前节点
     - 更新 `prev` 为当前节点
     - 将当前节点的左右子节点入队（如果存在）
  5. 重复步骤2-4，直到队列为空
  6. 返回根节点

#### 7.3 代码实现

```cpp
/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;
    Node* next;

    Node() : val(0), left(NULL), right(NULL), next(NULL) {}

    Node(int _val) : val(_val), left(NULL), right(NULL), next(NULL) {}

    Node(int _val, Node* _left, Node* _right, Node* _next)
        : val(_val), left(_left), right(_right), next(_next) {}
};
*/

class Solution {
public:
    Node* connect(Node* root) {
        queue<Node*> que;
        if (root != nullptr) que.push(root);
        
        while (!que.empty()) {
            int size = que.size();
            Node* prev = nullptr;
            for (int i = 0; i < size; i++) {
                Node* node = que.front();
                que.pop();
                if (prev != nullptr) {
                    prev->next = node;
                }
                prev = node;
                if (node->left) que.push(node->left);
                if (node->right) que.push(node->right);
            }
            // 每层的最后一个节点的 next 指针默认为 NULL，无需处理
        }
        return root;
    }
};
```

#### 7.4 代码分析

- **数据结构**：使用队列进行层序遍历
- **前一个节点指针**：
  - 使用 `prev` 指针记录当前层的前一个节点
  - 用于连接当前节点和前一个节点
- **连接逻辑**：
  - 对于每层的第一个节点，`prev` 为 `nullptr`，不需要连接
  - 对于后续节点，将 `prev->next` 指向当前节点
  - 每次循环结束后，更新 `prev` 为当前节点
- **子节点入队**：
  - 按照左子节点先入队，右子节点后入队的顺序
  - 确保每层的节点按从左到右的顺序处理
- **边界情况处理**：
  - 当根节点为 `nullptr` 时，直接返回 `nullptr`
  - 对于每层的最后一个节点，其 `next` 指针保持为 `nullptr`

#### 7.5 执行过程示例

对于完美二叉树：
```
     1
   /   
  2     3
 / \   / \
4   5 6   7
```

1. 初始化：`que = [1]`
2. 第一次外层循环：`size = 1`
   - `prev = nullptr`
   - 内层循环：`i = 0`
     - 弹出 1，`prev == nullptr`，不连接
     - `prev = 1`
     - 入队 2, 3，`que = [2, 3]`
3. 第二次外层循环：`size = 2`
   - `prev = nullptr`
   - 内层循环：`i = 0`
     - 弹出 2，`prev == nullptr`，不连接
     - `prev = 2`
     - 入队 4, 5，`que = [3, 4, 5]`
   - 内层循环：`i = 1`
     - 弹出 3，`prev != nullptr`，`2->next = 3`
     - `prev = 3`
     - 入队 6, 7，`que = [4, 5, 6, 7]`
4. 第三次外层循环：`size = 4`
   - `prev = nullptr`
   - 内层循环：`i = 0`
     - 弹出 4，`prev == nullptr`，不连接
     - `prev = 4`
     - 无左右子节点，`que = [5, 6, 7]`
   - 内层循环：`i = 1`
     - 弹出 5，`prev != nullptr`，`4->next = 5`
     - `prev = 5`
     - 无左右子节点，`que = [6, 7]`
   - 内层循环：`i = 2`
     - 弹出 6，`prev != nullptr`，`5->next = 6`
     - `prev = 6`
     - 无左右子节点，`que = [7]`
   - 内层循环：`i = 3`
     - 弹出 7，`prev != nullptr`，`6->next = 7`
     - `prev = 7`
     - 无左右子节点，`que = []`
5. 队列为空，结束遍历，返回根节点

#### 7.6 测试用例

**输入**：root = [1,2,3,4,5,6,7]
**输出**：[1,#,2,3,#,4,5,6,7,#]

**输入**：root = []
**输出**：[]

**输入**：root = [1]
**输出**：[1,#]

### 8. LeetCode 117. 填充每个节点的下一个右侧节点指针 II

#### 8.1 题目描述

给定一个二叉树：

```
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
```

填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL 。

初始状态下，所有 next 指针都被设置为 NULL 。

#### 8.2 解题思路

- **核心思路**：层序遍历，连接每层的节点
- **步骤**：
  1. 初始化队列，将根节点入队
  2. 当队列不为空时，记录当前队列大小（当前层节点数）
  3. 初始化前一个节点指针 `prev` 为 `nullptr`
  4. 遍历当前层的所有节点：
     - 如果 `prev` 不为 `nullptr`，将 `prev->next` 指向当前节点
     - 更新 `prev` 为当前节点
     - 将当前节点的左右子节点入队（如果存在）
  5. 重复步骤2-4，直到队列为空
  6. 返回根节点

#### 8.3 代码实现

```cpp
/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;
    Node* next;

    Node() : val(0), left(NULL), right(NULL), next(NULL) {}

    Node(int _val) : val(_val), left(NULL), right(NULL), next(NULL) {}

    Node(int _val, Node* _left, Node* _right, Node* _next)
        : val(_val), left(_left), right(_right), next(_next) {}
};
*/

class Solution {
public:
    Node* connect(Node* root) {
        queue<Node*> que;
        if (root != nullptr) que.push(root);
        
        while (!que.empty()) {
            int size = que.size();
            Node* prev = nullptr;
            for (int i = 0; i < size; i++) {
                Node* node = que.front();
                que.pop();
                if (prev != nullptr) {
                    prev->next = node;
                }
                prev = node;
                if (node->left) que.push(node->left);
                if (node->right) que.push(node->right);
            }
            // 每层的最后一个节点的 next 指针默认为 NULL，无需处理
        }
        return root;
    }
};
```

#### 8.4 代码分析

- **与 LeetCode 116 的区别**：
  - 本题不要求是完美二叉树，可能是任意二叉树
  - 但层序遍历的逻辑完全相同，因为队列会自动处理不同层的节点数
- **核心逻辑**：
  - 使用队列进行层序遍历
  - 使用 `prev` 指针记录前一个节点，用于连接当前节点
  - 按层处理节点，确保每层的节点按从左到右的顺序连接
- **边界情况处理**：
  - 当根节点为 `nullptr` 时，直接返回 `nullptr`
  - 当节点没有子节点时，跳过子节点入队步骤
  - 对于每层的最后一个节点，其 `next` 指针保持为 `nullptr`

#### 8.5 执行过程示例

对于二叉树：
```
     1
   /   
  2     3
 / \     \
4   5     7
```

1. 初始化：`que = [1]`
2. 第一次外层循环：`size = 1`
   - `prev = nullptr`
   - 内层循环：`i = 0`
     - 弹出 1，`prev == nullptr`，不连接
     - `prev = 1`
     - 入队 2, 3，`que = [2, 3]`
3. 第二次外层循环：`size = 2`
   - `prev = nullptr`
   - 内层循环：`i = 0`
     - 弹出 2，`prev == nullptr`，不连接
     - `prev = 2`
     - 入队 4, 5，`que = [3, 4, 5]`
   - 内层循环：`i = 1`
     - 弹出 3，`prev != nullptr`，`2->next = 3`
     - `prev = 3`
     - 入队 7，`que = [4, 5, 7]`
4. 第三次外层循环：`size = 3`
   - `prev = nullptr`
   - 内层循环：`i = 0`
     - 弹出 4，`prev == nullptr`，不连接
     - `prev = 4`
     - 无左右子节点，`que = [5, 7]`
   - 内层循环：`i = 1`
     - 弹出 5，`prev != nullptr`，`4->next = 5`
     - `prev = 5`
     - 无左右子节点，`que = [7]`
   - 内层循环：`i = 2`
     - 弹出 7，`prev != nullptr`，`5->next = 7`
     - `prev = 7`
     - 无左右子节点，`que = []`
5. 队列为空，结束遍历，返回根节点

#### 8.6 测试用例

**输入**：root = [1,2,3,4,5,null,7]
**输出**：[1,#,2,3,#,4,5,7,#]

**输入**：root = []
**输出**：[]

**输入**：root = [1,2,null,3,null,4,null,5]
**输出**：[1,#,2,#,3,#,4,#,5,#]

### 9. LeetCode 104. 二叉树的最大深度

#### 9.1 题目描述

给定一个二叉树 root ，返回其最大深度。

二叉树的 **最大深度** 是指从根节点到最远叶子节点的最长路径上的节点数。

#### 9.2 解题思路

- **核心思路**：层序遍历，记录层数
- **步骤**：
  1. 检查根节点是否为 `nullptr`，如果是则返回 0
  2. 初始化队列，将根节点入队
  3. 初始化深度变量 `depth` 为 0
  4. 当队列不为空时：
     - 记录当前队列大小（当前层节点数）
     - 深度加 1
     - 遍历当前层的所有节点，将其左右子节点入队（如果存在）
  5. 重复步骤4，直到队列为空
  6. 返回深度 `depth`

#### 9.3 代码实现

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
    int maxDepth(TreeNode* root) {
        if (root == nullptr) return 0;
        queue<TreeNode*> que;
        que.push(root);
        int depth = 0;
        
        while (!que.empty()) {
            int size = que.size();
            depth++;
            for (int i = 0; i < size; i++) {
                TreeNode* node = que.front();
                que.pop();
                if (node->left) que.push(node->left);
                if (node->right) que.push(node->right);
            }
        }
        return depth;
    }
};
```

#### 9.4 代码分析

- **边界情况处理**：
  - 当根节点为 `nullptr` 时，直接返回 0，避免后续操作
- **深度计算**：
  - 每处理完一层节点，深度加 1
  - 层序遍历确保了所有节点都被处理，因此最终的深度就是树的最大深度
- **队列操作**：
  - 每处理一个节点，将其左右子节点入队（如果存在）
  - 这样可以确保按层处理所有节点

#### 9.5 执行过程示例

对于二叉树：
```
    3
   / 
  9  20
    /  
   15   7
```

1. 初始化：`root != nullptr`，`que = [3]`，`depth = 0`
2. 第一次循环：`que != empty`
   - `size = 1`，`depth = 1`
   - 内层循环：`i = 0`
     - 弹出 3，入队 9, 20，`que = [9, 20]`
3. 第二次循环：`que != empty`
   - `size = 2`，`depth = 2`
   - 内层循环：`i = 0`
     - 弹出 9，无左右子节点，`que = [20]`
   - 内层循环：`i = 1`
     - 弹出 20，入队 15, 7，`que = [15, 7]`
4. 第三次循环：`que != empty`
   - `size = 2`，`depth = 3`
   - 内层循环：`i = 0`
     - 弹出 15，无左右子节点，`que = [7]`
   - 内层循环：`i = 1`
     - 弹出 7，无左右子节点，`que = []`
5. 队列为空，结束循环，返回 `depth = 3`

#### 9.6 测试用例

**输入**：root = [3,9,20,null,null,15,7]
**输出**：3

**输入**：root = [1]
**输出**：1

**输入**：root = []
**输出**：0

### 10. LeetCode 111. 二叉树的最小深度

#### 10.1 题目描述

给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

**说明**：叶子节点是指没有子节点的节点。

#### 10.2 解题思路

- **核心思路**：层序遍历，遇到叶子节点时返回当前层数
- **步骤**：
  1. 检查根节点是否为 `nullptr`，如果是则返回 0
  2. 初始化队列，将根节点入队
  3. 初始化深度变量 `depth` 为 0
  4. 当队列不为空时：
     - 记录当前队列大小（当前层节点数）
     - 深度加 1
     - 遍历当前层的所有节点：
       - 如果节点是叶子节点（左右子节点都为 `nullptr`），返回当前深度
       - 将节点的左右子节点入队（如果存在）
  5. 重复步骤4，直到队列为空
  6. 返回深度 `depth`

#### 10.3 代码实现

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
    int minDepth(TreeNode* root) {
        if (root == nullptr) return 0;
        queue<TreeNode*> que;
        que.push(root);
        int depth = 0;
        
        while (!que.empty()) {
            int size = que.size();
            depth++;
            for (int i = 0; i < size; i++) {
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
};
```

#### 10.4 代码分析

- **边界情况处理**：
  - 当根节点为 `nullptr` 时，直接返回 0
- **叶子节点判断**：
  - 当节点的左右子节点都为 `nullptr` 时，该节点为叶子节点
  - 遇到叶子节点时，立即返回当前深度，因为层序遍历是按层进行的，第一次遇到的叶子节点一定是深度最小的
- **队列操作**：
  - 每处理一个节点，将其左右子节点入队（如果存在）
  - 这样可以确保按层处理所有节点

#### 10.5 执行过程示例

对于二叉树：
```
    3
   / 
  9  20
    /  
   15   7
```

1. 初始化：`root != nullptr`，`que = [3]`，`depth = 0`
2. 第一次循环：`que != empty`
   - `size = 1`，`depth = 1`
   - 内层循环：`i = 0`
     - 弹出 3，不是叶子节点
     - 入队 9, 20，`que = [9, 20]`
3. 第二次循环：`que != empty`
   - `size = 2`，`depth = 2`
   - 内层循环：`i = 0`
     - 弹出 9，是叶子节点，返回 `depth = 2`

#### 10.6 测试用例

**输入**：root = [3,9,20,null,null,15,7]
**输出**：2

**输入**：root = [2,null,3,null,4,null,5,null,6]
**输出**：5

**输入**：root = []
**输出**：0

---

## 四、 面试八股内容

### 1. 基础概念类问题

#### 1.1 什么是二叉树的层序遍历？
**答案**：二叉树的层序遍历是按照从上到下、从左到右的顺序逐层访问二叉树中的所有节点，是广度优先搜索（BFS）在二叉树上的应用。层序遍历的核心思想是使用队列来保证节点处理的顺序，确保同一层的节点按顺序处理完毕后，再处理下一层的节点。

#### 1.2 层序遍历和深度优先遍历有什么区别？
**答案**：
| 特性 | 层序遍历（BFS） | 深度优先遍历（DFS） |
|------|----------------|-------------------|
| 实现方式 | 使用队列 | 使用栈（递归或显式栈） |
| 遍历顺序 | 按层从上到下，每层从左到右 | 先深入到叶子节点，再回溯 |
| 适用场景 | 层次相关问题（如树的高度、宽度） | 路径相关问题（如前序、中序、后序遍历） |
| 空间复杂度 | O(n)，最坏情况为完全二叉树的最后一层节点数 | O(n)，最坏情况为树退化为链表时的递归栈深度 |
| 时间复杂度 | O(n)，每个节点访问一次 | O(n)，每个节点访问一次 |

#### 1.3 层序遍历的时间复杂度和空间复杂度是多少？
**答案**：
- **时间复杂度**：O(n)，其中 n 是二叉树的节点数。每个节点都会被访问一次，入队和出队操作都是 O(1) 时间复杂度。
- **空间复杂度**：O(n)，最坏情况下（完全二叉树），队列最多存储 n/2 个节点（最后一层的节点数）。对于平衡二叉树，空间复杂度为 O(n)；对于斜树（退化为链表），空间复杂度为 O(1)。

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
    vector<vector<int>> result;
    queue<TreeNode*> que;
    if (root != nullptr) que.push(root);
    while (!que.empty()) {
        int size = que.size();
        vector<int> vec;
        for (int i = 0; i < size; i++) {
            TreeNode* node = que.front();
            que.pop();
            vec.push_back(node->val);
            if (node->left) que.push(node->left);
            if (node->right) que.push(node->right);
        }
        result.push_back(vec);
    }
    return result;
}
```

#### 2.2 如何实现自底向上的层序遍历？
**答案**：
1. **正常层序遍历**：按照标准层序遍历的方法获取结果，得到从上到下的层次顺序。
2. **反转结果**：使用 `reverse` 函数将结果向量反转，得到自底向上的顺序。

**代码框架**：
```cpp
vector<vector<int>> levelOrderBottom(TreeNode* root) {
    vector<vector<int>> result;
    queue<TreeNode*> que;
    if (root != nullptr) que.push(root);
    while (!que.empty()) {
        int size = que.size();
        vector<int> vec;
        for (int i = 0; i < size; i++) {
            TreeNode* node = que.front();
            que.pop();
            vec.push_back(node->val);
            if (node->left) que.push(node->left);
            if (node->right) que.push(node->right);
        }
        result.push_back(vec);
    }
    reverse(result.begin(), result.end());
    return result;
}
```

#### 2.3 如何实现二叉树的右视图？
**答案**：
1. **层序遍历**：使用队列进行层序遍历。
2. **记录每层最后节点**：在遍历每层节点时，记录当前层的最后一个节点，将其值加入结果集。

**代码框架**：
```cpp
vector<int> rightSideView(TreeNode* root) {
    vector<int> result;
    queue<TreeNode*> que;
    if (root != nullptr) que.push(root);
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

### 3. 扩展问题

#### 3.1 除了队列，还可以用什么数据结构实现层序遍历？
**答案**：
- **递归**：通过记录节点的深度，将节点值加入对应深度的结果集中。
- **双端队列**：在某些情况下，可以使用双端队列来实现层序遍历，特别是在需要同时处理两端节点的场景。
- **数组**：对于完全二叉树，可以使用数组来模拟层序遍历，利用完全二叉树的性质（第 i 个节点的左子节点是 2i+1，右子节点是 2i+2）。

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

#### 3.3 如何在层序遍历中处理每层的信息？
**答案**：通过在每层开始时记录队列的大小，确保只处理当前层的节点，然后将它们的子节点入队。具体步骤如下：
1. 在每次循环开始时，获取当前队列的大小 `size`，这表示当前层的节点数。
2. 然后进行 `size` 次循环，处理当前层的所有节点。
3. 在处理每个节点时，将其左右子节点入队。
4. 这样可以确保在处理完当前层的所有节点后，队列中恰好存储了下一层的所有节点。

### 4. 面试技巧

#### 4.1 如何在面试中回答层序遍历相关的问题？
**答案**：
1. **理解问题**：首先要明确问题的要求，是标准层序遍历还是其变种（如自底向上、右视图等）。
2. **选择方法**：根据问题要求选择合适的实现方法，通常使用队列实现迭代版本。
3. **代码实现**：写出清晰、正确的代码，注意边界情况的处理。
4. **分析复杂度**：分析时间复杂度和空间复杂度，并解释原因。
5. **优化思路**：如果可能，提出优化方案，例如空间复杂度的优化。
6. **测试示例**：使用示例二叉树验证代码的正确性。

#### 4.2 层序遍历相关的常见陷阱
**答案**：
- **队列大小错误**：在循环中使用 `que.size()` 作为循环条件，而不是在循环开始时记录大小，导致循环次数错误。
- **空指针处理**：不检查子节点是否为 `nullptr` 就入队，导致空指针异常。
- **边界情况**：忘记处理空树的情况，导致程序崩溃。
- **结果存储**：对于自底向上的层序遍历，忘记反转结果，导致返回的是正常顺序的结果。
- **右视图逻辑**：错误地只考虑右子节点，而不是记录每层的最后一个节点，导致遗漏某些情况。

---

## 五、 代码分析与知识点讲解

### 1. 核心代码分析

#### 1.1 队列实现的核心代码

```cpp
while (!que.empty()) {
    int size = que.size();
    vector<int> vec;
    for (int i = 0; i < size; i++) {
        TreeNode* node = que.front();
        que.pop();
        vec.push_back(node->val);
        if (node->left) que.push(node->left);
        if (node->right) que.push(node->right);
    }
    result.push_back(vec);
}
```

**代码分析**：
- **外层循环**：当队列不为空时，继续遍历
- **层的处理**：
  - 记录当前队列的大小 `size`，这表示当前层的节点数
  - 创建临时向量 `vec` 用于存储当前层的节点值
  - 内层循环遍历当前层的所有节点
- **节点处理**：
  - 获取并弹出队首节点
  - 将节点值加入当前层的临时向量
  - 将左子节点入队（如果存在）
  - 将右子节点入队（如果存在）
- **结果存储**：将当前层的临时向量加入最终结果

**关键知识点**：
- **队列的使用**：利用队列的先进先出特性，确保节点按层处理
- **层的分离**：通过记录每层开始时的队列大小，确保只处理当前层的节点
- **边界情况处理**：在入队前检查子节点是否为 `nullptr`，避免处理空指针

#### 1.2 递归实现的核心代码

```cpp
void order(TreeNode* cur, vector<vector<int>>& result, int depth) {
    if (cur == nullptr) return;
    if (result.size() == depth) result.push_back(vector<int>());
    result[depth].push_back(cur->val);
    order(cur->left, result, depth + 1);
    order(cur->right, result, depth + 1);
}
```

**代码分析**：
- **递归终止条件**：当当前节点为 `nullptr` 时，直接返回
- **层的创建**：当 `result.size() == depth` 时，说明当前深度对应的层还未创建，需要创建新层
- **节点处理**：将当前节点值加入对应深度的层
- **递归调用**：递归遍历左子树和右子树，深度加 1

**关键知识点**：
- **深度参数**：通过深度参数将节点映射到对应的层
- **层的动态创建**：根据深度动态创建新层，确保结果结构正确
- **递归终止条件**：处理空节点，避免无限递归

### 2. 扩展知识点

#### 2.1 层序遍历的变种

- **自底向上的层序遍历**：先按正常顺序遍历，然后反转结果
- **二叉树的右视图**：层序遍历，记录每层的最后一个节点
- **二叉树的左视图**：层序遍历，记录每层的第一个节点
- **之字形层序遍历**：奇数层从左到右，偶数层从右到左，使用双端队列实现

#### 2.2 层序遍历的应用

- **计算树的高度**：层序遍历，记录层数
- **计算每层的平均值**：层序遍历，计算每层的平均值
- **找到每层的最大值**：层序遍历，记录每层的最大值
- **填充每个节点的下一个右侧节点指针**：层序遍历，将每层的节点连接起来
- **N叉树的层序遍历**：层序遍历，处理每个节点的所有子节点

#### 2.3 层序遍历的优化

- **空间优化**：对于大规模二叉树，可以考虑使用迭代器或生成器模式，按需生成节点值，减少内存使用
- **时间优化**：对于频繁访问的二叉树，可以考虑缓存层序遍历结果
- **可读性优化**：使用命名清晰的变量和函数，添加详细的注释，提取重复代码为辅助函数

---

## 六、 总结

### 1. 核心知识点

- **层序遍历原理**：使用队列实现广度优先搜索，按层访问节点
- **基本步骤**：初始化队列 → 记录层大小 → 处理当前层节点 → 入队子节点 → 重复直到队列为空
- **实现方式**：队列实现（迭代法）和递归实现
- **时间复杂度**：O(n)，每个节点都会被访问一次
- **空间复杂度**：O(n)，最坏情况下（完全二叉树），队列最多存储 n/2 个节点

### 2. 相关题目总结

| 题目 | 核心要求 | 特殊处理 |
|------|---------|----------|
| 102. 二叉树的层序遍历 | 按层返回节点值 | 直接返回层序遍历结果 |
| 107. 二叉树的层序遍历 II | 自底向上返回节点值 | 反转层序遍历结果 |
| 199. 二叉树的右视图 | 返回每层最右侧节点 | 只记录每层的最后一个节点 |
| 637. 二叉树的层平均值 | 返回每层节点的平均值 | 计算每层的平均值 |
| 429. N叉树的层序遍历 | 按层返回N叉树节点值 | 处理每个节点的所有子节点 |
| 515. 在每个树行中找最大值 | 返回每层的最大值 | 记录每层的最大值 |
| 116. 填充每个节点的下一个右侧节点指针 | 填充每个节点的next指针 | 将每层的节点连接起来 |
| 117. 填充每个节点的下一个右侧节点指针 II | 填充每个节点的next指针 | 将每层的节点连接起来 |
| 104. 二叉树的最大深度 | 返回树的最大深度 | 记录层数 |
| 111. 二叉树的最小深度 | 返回树的最小深度 | 遇到叶子节点时返回当前层数 |

### 3. 学习收获

- 掌握了二叉树层序遍历的原理和实现方法
- 理解了广度优先搜索在二叉树上的应用
- 学会了如何处理层序遍历相关的变种问题
- 掌握了队列的基本操作和应用场景
- 了解了层序遍历的扩展应用和优化方法

### 4. 后续学习建议

- 学习图的广度优先搜索算法
- 探索层序遍历的其他应用场景
- 研究其他树的遍历算法（如Morris遍历）
- 练习更多层序遍历相关的题目，巩固所学知识

通过本课程的学习，我们掌握了二叉树层序遍历的原理和实现方法，理解了其在解决相关问题中的应用。希望这份学习笔记能帮助你更好地理解和应用这些知识点！