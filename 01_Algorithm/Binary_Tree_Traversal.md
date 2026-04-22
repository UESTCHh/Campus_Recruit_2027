# 二叉树遍历算法综合学习笔记

> **学习内容**：二叉树的前序、中序、后序遍历
> **学习资源**：代码随想录二叉树基础部分和递归遍历部分
> **相关题目**：
> - 144. 二叉树的前序遍历
> - 94. 二叉树的中序遍历
> - 145. 二叉树的后序遍历
> **学习目标**：掌握二叉树遍历的递归和迭代实现方法，理解其时间空间复杂度，以及边界情况处理

---

## 一、 二叉树遍历概述

### 1. 遍历方式分类

二叉树主要有两种遍历方式：

- **深度优先遍历**（DFS）：先往深走，遇到叶子节点再往回走
  - 前序遍历（中左右）
  - 中序遍历（左中右）
  - 后序遍历（左右中）

- **广度优先遍历**（BFS）：一层一层的去遍历
  - 层次遍历

### 2. 前中后序遍历的定义

- **前序遍历**：先访问根节点，然后遍历左子树，最后遍历右子树
- **中序遍历**：先遍历左子树，然后访问根节点，最后遍历右子树
- **后序遍历**：先遍历左子树，然后遍历右子树，最后访问根节点

### 3. 遍历顺序示例

对于如下二叉树：
```
    1
   / \
  2   3
 / \   \
4   5   8
   / \
  6   7
```

- **前序遍历**：1, 2, 4, 5, 6, 7, 3, 8
- **中序遍历**：4, 2, 6, 5, 7, 1, 3, 8
- **后序遍历**：4, 6, 7, 5, 2, 8, 3, 1

---

## 二、 递归实现方法

### 1. 递归算法三要素

写递归算法有三大核心要素：

1. **确定递归函数的参数和返回值**：
   - 参数：当前遍历的节点 `TreeNode* cur`，存储结果的容器 `vector<int>& vec`
   - 返回值：`void`，因为结果通过引用传递

2. **确定终止条件**：
   - 当当前节点为空时，直接返回
   ```cpp
   if (cur == nullptr) return;
   ```

3. **确定单层递归的逻辑**：
   - 前序遍历：中左右
   - 中序遍历：左中右
   - 后序遍历：左右中

### 2. 代码实现

#### 2.1 前序遍历

```cpp
class Solution {
public:
    /**
     * @brief 递归辅助函数，实现前序遍历（中左右）
     * @param cur 当前遍历的节点指针
     * @param vec 存储遍历结果的向量引用，通过引用传递避免拷贝
     */
    void traversal(TreeNode* cur, vector<int>& vec) {
        // 递归终止条件：当节点为空时，直接返回，不再继续递归
        if (cur == nullptr) return;

        // 处理当前节点（中）：将当前节点的值加入结果向量
        vec.push_back(cur->val);

        // 递归遍历左子树（左）
        traversal(cur->left, vec);

        // 递归遍历右子树（右）
        traversal(cur->right, vec);
    }

    /**
     * @brief 前序遍历入口函数
     * @param root 二叉树的根节点指针
     * @return 包含前序遍历结果的向量
     */
    vector<int> preorderTraversal(TreeNode* root) {
        vector<int> result;  // 存储遍历结果
        traversal(root, result);  // 调用递归辅助函数
        return result;  // 返回遍历结果
    }
};
```

#### 2.2 中序遍历

```cpp
class Solution {
public:
    /**
     * @brief 递归辅助函数，实现中序遍历（左中右）
     * @param cur 当前遍历的节点指针
     * @param vec 存储遍历结果的向量引用
     */
    void traversal(TreeNode* cur, vector<int>& vec) {
        // 递归终止条件：当节点为空时，直接返回
        if (cur == nullptr) return;

        // 递归遍历左子树（左）
        traversal(cur->left, vec);

        // 处理当前节点（中）：将当前节点的值加入结果向量
        vec.push_back(cur->val);

        // 递归遍历右子树（右）
        traversal(cur->right, vec);
    }

    /**
     * @brief 中序遍历入口函数
     * @param root 二叉树的根节点指针
     * @return 包含中序遍历结果的向量
     */
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> result;  // 存储遍历结果
        traversal(root, result);  // 调用递归辅助函数
        return result;  // 返回遍历结果
    }
};
```

#### 2.3 后序遍历

```cpp
class Solution {
public:
    /**
     * @brief 递归辅助函数，实现后序遍历（左右中）
     * @param cur 当前遍历的节点指针
     * @param vec 存储遍历结果的向量引用
     */
    void traversal(TreeNode* cur, vector<int>& vec) {
        // 递归终止条件：当节点为空时，直接返回
        if (cur == nullptr) return;

        // 递归遍历左子树（左）
        traversal(cur->left, vec);

        // 递归遍历右子树（右）
        traversal(cur->right, vec);

        // 处理当前节点（中）：将当前节点的值加入结果向量
        vec.push_back(cur->val);
    }

    /**
     * @brief 后序遍历入口函数
     * @param root 二叉树的根节点指针
     * @return 包含后序遍历结果的向量
     */
    vector<int> postorderTraversal(TreeNode* root) {
        vector<int> result;  // 存储遍历结果
        traversal(root, result);  // 调用递归辅助函数
        return result;  // 返回遍历结果
    }
};
```

### 3. 递归实现分析

#### 3.1 时间复杂度
- **时间复杂度**：O(n)，其中 n 是二叉树的节点数。每个节点都会被访问一次

#### 3.2 空间复杂度
- **空间复杂度**：O(n)，其中 n 是二叉树的节点数。递归调用的栈深度最坏情况下为 n（当二叉树退化为链表时）

#### 3.3 优缺点
- **优点**：代码简洁易懂，逻辑清晰
- **缺点**：递归调用可能导致栈溢出，对于深度很大的二叉树不适用

---

## 三、 迭代实现方法

### 1. 前序遍历迭代实现

#### 1.1 思路
- 使用栈来模拟递归过程
- 由于栈的后进先出特性，前序遍历需要先压入右子节点，再压入左子节点
- 这样出栈顺序就是中左右，符合前序遍历的顺序

#### 1.2 代码实现

```cpp
class Solution {
public:
    /**
     * @brief 前序遍历迭代实现（使用栈）
     * @param root 二叉树的根节点指针
     * @return 包含前序遍历结果的向量
     */
    vector<int> preorderTraversal(TreeNode* root) {
        vector<int> result;  // 存储遍历结果

        // 边界情况处理：空树直接返回空结果
        if (root == nullptr) return result;

        // 创建栈用于模拟递归过程
        stack<TreeNode*> st;
        st.push(root);  // 将根节点入栈

        // 开始迭代过程
        while (!st.empty()) {
            // 弹出栈顶节点（相当于处理"中"节点）
            TreeNode* node = st.top();
            st.pop();

            // 将当前节点的值加入结果向量（中）
            result.push_back(node->val);

            // 注意：先压右子节点，后压左子节点
            // 因为栈是LIFO，这样左子节点会先出栈，符合前序遍历顺序
            if (node->right) st.push(node->right);  // 右子节点入栈
            if (node->left) st.push(node->left);   // 左子节点入栈
        }

        return result;  // 返回遍历结果
    }
};
```

### 2. 中序遍历迭代实现

#### 2.1 思路
- 使用栈来模拟递归过程
- 需要先将左子树全部压入栈中，然后弹出节点访问，再处理右子树
- 关键点：一直向左走，将所有左子节点入栈，直到最左边的节点
- 然后弹出节点访问，再向右走，处理右子树

#### 2.2 代码实现

```cpp
class Solution {
public:
    /**
     * @brief 中序遍历迭代实现（使用栈）
     * @param root 二叉树的根节点指针
     * @return 包含中序遍历结果的向量
     */
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> result;  // 存储遍历结果
        stack<TreeNode*> st;  // 创建栈用于模拟递归过程
        TreeNode* cur = root;  // 当前遍历指针，初始化为根节点

        // 迭代条件：当前节点不为空或栈不为空
        while (cur != nullptr || !st.empty()) {
            if (cur != nullptr) {
                // 当前节点不为空，将当前节点入栈
                // 然后继续向左走，将所有左子节点入栈
                st.push(cur);
                cur = cur->left;  // 移动到左子节点
            } else {
                // 当前节点为空（已经到达最左叶子节点的左孩子）
                // 弹出栈顶节点进行处理
                cur = st.top();
                st.pop();

                // 处理当前节点（中）：将节点值加入结果向量
                result.push_back(cur->val);

                // 处理右子树：移动到右子节点，继续循环
                cur = cur->right;
            }
        }

        return result;  // 返回遍历结果
    }
};
```

### 3. 后序遍历迭代实现

#### 3.1 思路
- 方法1（推荐）：利用前序遍历"中左右"的思想，先实现"中右左"，然后反转结果得到"左右中"
- 方法2：使用一个栈，记录节点是否被访问过

#### 3.2 代码实现（方法1：反转法）

```cpp
class Solution {
public:
    /**
     * @brief 后序遍历迭代实现（使用栈 + 反转）
     * @param root 二叉树的根节点指针
     * @return 包含后序遍历结果的向量
     */
    vector<int> postorderTraversal(TreeNode* root) {
        vector<int> result;  // 存储遍历结果

        // 边界情况处理：空树直接返回空结果
        if (root == nullptr) return result;

        // 创建栈用于模拟递归过程
        stack<TreeNode*> st;
        st.push(root);  // 将根节点入栈

        // 迭代过程：实现"中右左"的遍历顺序
        while (!st.empty()) {
            // 弹出栈顶节点（相当于处理"中"节点）
            TreeNode* node = st.top();
            st.pop();

            // 将当前节点的值加入结果向量（中）
            result.push_back(node->val);

            // 注意：与前序遍历不同，这里先压左子节点，后压右子节点
            // 这样出栈顺序是"中右左"，符合后序遍历的逆序
            if (node->left) st.push(node->left);   // 左子节点入栈
            if (node->right) st.push(node->right); // 右子节点入栈
        }

        // 反转结果：将"中右左"变为"左右中"，得到正确的后序遍历结果
        reverse(result.begin(), result.end());

        return result;  // 返回遍历结果
    }
};
```

#### 3.3 代码实现（方法2：标记法）

```cpp
class Solution {
public:
    /**
     * @brief 后序遍历迭代实现（使用栈 + 标记）
     * @param root 二叉树的根节点指针
     * @return 包含后序遍历结果的向量
     */
    vector<int> postorderTraversal(TreeNode* root) {
        vector<int> result;  // 存储遍历结果

        // 边界情况处理：空树直接返回空结果
        if (root == nullptr) return result;

        // 创建栈，栈元素为 pair<节点指针, 是否已访问>
        // visited = false 表示需要继续向下遍历
        // visited = true 表示需要处理该节点（输出其值）
        stack<pair<TreeNode*, bool>> st;
        st.push({root, false});  // 将根节点标记为"未访问"入栈

        // 迭代过程
        while (!st.empty()) {
            // 弹出栈顶元素
            auto [node, visited] = st.top();
            st.pop();

            if (visited) {
                // 如果节点已被访问（visited = true），则处理该节点
                result.push_back(node->val);
            } else {
                // 如果节点未被访问（visited = false），则将其标记为已访问
                // 并将其右、左子节点按顺序入栈

                // 入栈顺序：中 -> 右 -> 左（因为栈是LIFO，出栈顺序是左 -> 右 -> 中）
                // 这样最终处理顺序就是：左 -> 右 -> 中，符合后序遍历
                st.push({node, true});  // 标记当前节点为已访问

                if (node->right) st.push({node->right, false});  // 右子节点入栈（未访问）
                if (node->left) st.push({node->left, false});   // 左子节点入栈（未访问）
            }
        }

        return result;  // 返回遍历结果
    }
};
```

### 4. 迭代实现分析

#### 4.1 时间复杂度
- **时间复杂度**：O(n)，其中 n 是二叉树的节点数。每个节点都会被访问一次

#### 4.2 空间复杂度
- **空间复杂度**：O(n)，其中 n 是二叉树的节点数。栈的大小最坏情况下为 n（当二叉树退化为链表时）

#### 4.3 优缺点
- **优点**：避免了递归调用可能导致的栈溢出问题
- **缺点**：代码相对复杂，逻辑不如递归清晰

---

## 四、 统一迭代法

### 1. 思路
- 使用一个栈，通过标记法来统一处理前中后序遍历
- 使用 `nullptr` 作为标记，表示该节点已经被访问过，需要处理其值
- 核心思想：将"处理节点"和"遍历子树"分开，需要处理节点值时压入标记

### 2. 代码实现

#### 2.1 前序遍历

```cpp
class Solution {
public:
    /**
     * @brief 前序遍历统一迭代法（使用栈 + nullptr 标记）
     * @param root 二叉树的根节点指针
     * @return 包含前序遍历结果的向量
     */
    vector<int> preorderTraversal(TreeNode* root) {
        vector<int> result;  // 存储遍历结果

        // 边界情况处理：空树直接返回空结果
        if (root == nullptr) return result;

        // 创建栈，用于模拟递归过程
        stack<TreeNode*> st;
        st.push(root);  // 将根节点入栈

        // 开始迭代过程
        while (!st.empty()) {
            // 弹出栈顶元素
            TreeNode* node = st.top();
            st.pop();

            if (node != nullptr) {
                // node 不为 nullptr，表示这是一个需要处理的节点

                // 前序遍历顺序：中左右
                // 入栈顺序（逆序）：右 -> 左 -> 中
                // 这样出栈顺序就是：中 -> 左 -> 右，符合前序遍历顺序

                if (node->right) st.push(node->right);  // 右子节点入栈
                if (node->left) st.push(node->left);   // 左子节点入栈
                st.push(node);                          // 中节点入栈
                st.push(nullptr);                      // nullptr 标记，表示该节点需要被处理
            } else {
                // node 为 nullptr，表示这是一个标记，需要处理其后面的节点
                node = st.top();  // 获取标记后面的节点
                st.pop();
                result.push_back(node->val);  // 处理该节点（中）
            }
        }

        return result;  // 返回遍历结果
    }
};
```

#### 2.2 中序遍历

```cpp
class Solution {
public:
    /**
     * @brief 中序遍历统一迭代法（使用栈 + nullptr 标记）
     * @param root 二叉树的根节点指针
     * @return 包含中序遍历结果的向量
     */
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> result;  // 存储遍历结果

        // 边界情况处理：空树直接返回空结果
        if (root == nullptr) return result;

        // 创建栈，用于模拟递归过程
        stack<TreeNode*> st;
        st.push(root);  // 将根节点入栈

        // 开始迭代过程
        while (!st.empty()) {
            // 弹出栈顶元素
            TreeNode* node = st.top();
            st.pop();

            if (node != nullptr) {
                // node 不为 nullptr，表示这是一个需要遍历的节点

                // 中序遍历顺序：左中右
                // 入栈顺序（逆序）：右 -> 中 -> 左
                // 这样出栈顺序就是：左 -> 中 -> 右，符合中序遍历顺序

                if (node->right) st.push(node->right);  // 右子节点入栈
                st.push(node);                          // 中节点入栈
                st.push(nullptr);                      // nullptr 标记，表示该节点需要被处理
                if (node->left) st.push(node->left);   // 左子节点入栈
            } else {
                // node 为 nullptr，表示这是一个标记，需要处理其后面的节点
                node = st.top();  // 获取标记后面的节点
                st.pop();
                result.push_back(node->val);  // 处理该节点（中）
            }
        }

        return result;  // 返回遍历结果
    }
};
```

#### 2.3 后序遍历

```cpp
class Solution {
public:
    /**
     * @brief 后序遍历统一迭代法（使用栈 + nullptr 标记）
     * @param root 二叉树的根节点指针
     * @return 包含后序遍历结果的向量
     */
    vector<int> postorderTraversal(TreeNode* root) {
        vector<int> result;  // 存储遍历结果

        // 边界情况处理：空树直接返回空结果
        if (root == nullptr) return result;

        // 创建栈，用于模拟递归过程
        stack<TreeNode*> st;
        st.push(root);  // 将根节点入栈

        // 开始迭代过程
        while (!st.empty()) {
            // 弹出栈顶元素
            TreeNode* node = st.top();
            st.pop();

            if (node != nullptr) {
                // node 不为 nullptr，表示这是一个需要遍历的节点

                // 后序遍历顺序：左右中
                // 入栈顺序（逆序）：中 -> 右 -> 左
                // 这样出栈顺序就是：左 -> 右 -> 中，符合后序遍历顺序

                st.push(node);                          // 中节点入栈
                st.push(nullptr);                      // nullptr 标记，表示该节点需要被处理
                if (node->right) st.push(node->right);  // 右子节点入栈
                if (node->left) st.push(node->left);   // 左子节点入栈
            } else {
                // node 为 nullptr，表示这是一个标记，需要处理其后面的节点
                node = st.top();  // 获取标记后面的节点
                st.pop();
                result.push_back(node->val);  // 处理该节点（中）
            }
        }

        return result;  // 返回遍历结果
    }
};
```

### 3. 统一迭代法分析

#### 3.1 时间复杂度
- **时间复杂度**：O(n)，其中 n 是二叉树的节点数。每个节点都会被访问两次（入栈和出栈）

#### 3.2 空间复杂度
- **空间复杂度**：O(n)，其中 n 是二叉树的节点数。栈的大小最坏情况下为 n

#### 3.3 优缺点
- **优点**：统一了前中后序遍历的实现方式，代码结构相似，便于理解和记忆
- **缺点**：需要额外的标记，增加了空间开销（但仍是 O(n)）

---

## 五、 三种遍历方法的对比

### 1. 实现方式对比

| 遍历方式 | 递归实现 | 迭代实现 | 统一迭代法 |
|---------|---------|---------|-----------|
| 前序遍历 | 中左右 | 中左右（栈：右左中） | 中左右（栈：右左中） |
| 中序遍历 | 左中右 | 左中右（栈：左链入栈） | 左中右（栈：右中左） |
| 后序遍历 | 左右中 | 左右中（栈：中右左 + 反转） | 左右中（栈：中右左） |

### 2. 时间空间复杂度对比

| 遍历方式 | 递归实现 | 迭代实现 | 统一迭代法 |
|---------|---------|---------|-----------|
| 时间复杂度 | O(n) | O(n) | O(n) |
| 空间复杂度 | O(n) | O(n) | O(n) |

### 3. 适用场景

- **递归实现**：代码简洁，逻辑清晰，适用于树深度不大的情况
- **迭代实现**：避免栈溢出，适用于树深度较大的情况
- **统一迭代法**：代码结构统一，便于理解和记忆

---

## 六、 边界情况处理

### 1. 空树
- **处理方法**：直接返回空数组
- **代码示例**：
  ```cpp
  if (root == nullptr) return result;
  ```

### 2. 单节点树
- **处理方法**：直接返回包含该节点值的数组
- **代码示例**：递归和迭代实现都能正确处理

### 3. 退化为链表的树
- **处理方法**：
  - 递归实现：可能导致栈溢出
  - 迭代实现：能正常处理

### 4. 完全二叉树
- **处理方法**：递归和迭代实现都能正确处理

---

## 七、 算法优化思路

### 1. 空间优化
- **思路**：对于前序和中序遍历，可以使用 Morris 遍历算法，将空间复杂度降低到 O(1)
- **Morris 遍历**：利用树的空闲指针，实现线索化遍历

### 2. 时间优化
- **思路**：对于大规模二叉树，迭代实现可能比递归实现更快，因为减少了函数调用的开销

### 3. 代码可读性优化
- **思路**：使用统一迭代法，使三种遍历的代码结构相似，便于理解和维护

---

## 八、 面试八股内容

### 1. 基础概念类问题

#### 1.1 什么是二叉树的遍历？
**答案**：二叉树的遍历是指按照某种顺序访问二叉树中的所有节点，每个节点被访问且仅被访问一次。常见的遍历方式有深度优先遍历（前序、中序、后序）和广度优先遍历（层次遍历）。

#### 1.2 前序、中序、后序遍历的区别是什么？
**答案**：
- **前序遍历**（中左右）：先访问根节点，然后遍历左子树，最后遍历右子树
- **中序遍历**（左中右）：先遍历左子树，然后访问根节点，最后遍历右子树
- **后序遍历**（左右中）：先遍历左子树，然后遍历右子树，最后访问根节点

#### 1.3 深度优先遍历和广度优先遍历有什么区别？
**答案**：
- **深度优先遍历**（DFS）：先往深走，遇到叶子节点再往回走。使用栈（隐式或显式）实现
- **广度优先遍历**（BFS）：一层一层的去遍历。使用队列实现

### 2. 递归实现类问题

#### 2.1 递归实现二叉树遍历的三要素是什么？
**答案**：
1. **确定递归函数的参数和返回值**：参数通常包括当前遍历的节点和存储结果的容器，返回值为 void
2. **确定终止条件**：当当前节点为空时，直接返回
3. **确定单层递归的逻辑**：按照遍历顺序处理节点（左/右子树的递归调用顺序）

#### 2.2 递归实现的时间复杂度和空间复杂度是多少？
**答案**：
- **时间复杂度**：O(n)，其中 n 是二叉树的节点数，每个节点都会被访问一次
- **空间复杂度**：O(n)，最坏情况下（树退化为链表）递归栈深度为 n

#### 2.3 递归实现的优点和缺点是什么？
**答案**：
- **优点**：代码简洁易懂，逻辑清晰，容易实现
- **缺点**：递归调用可能导致栈溢出，对于深度很大的二叉树不适用

#### 2.4 什么情况下递归会导致栈溢出？
**答案**：当二叉树的深度很大时（如退化为链表的二叉树），递归调用次数过多，会导致栈空间耗尽，抛出栈溢出异常。

### 3. 迭代实现类问题

#### 3.1 迭代实现前序遍历的核心思路是什么？
**答案**：使用栈模拟递归过程。由于栈是后进先出（LIFO），前序遍历需要先压入右子节点，再压入左子节点，这样出栈顺序就是中左右。

#### 3.2 迭代实现中序遍历的核心思路是什么？
**答案**：使用栈模拟递归过程。关键点是先一直向左走，将所有左子节点入栈，直到最左边的节点。然后弹出节点访问（处理中），再向右走处理右子树。

#### 3.3 迭代实现后序遍历有哪些方法？
**答案**：
- **方法1（推荐）**：利用前序遍历"中左右"的思想，先实现"中右左"，然后反转结果得到"左右中"
- **方法2**：使用一个栈，记录节点是否被访问过（标记法）

#### 3.4 迭代实现和递归实现相比有什么优势？
**答案**：
- 避免了递归调用可能导致的栈溢出问题
- 对于深度很大的二叉树也能正常处理

### 4. 统一迭代法类问题

#### 4.1 什么是统一迭代法？
**答案**：统一迭代法是一种使用栈和标记（nullptr）来统一处理前中后序遍历的方法。通过在节点入栈时插入 nullptr 标记，将"处理节点"和"遍历子树"分开，使三种遍历的代码结构相似。

#### 4.2 统一迭代法中 nullptr 标记的作用是什么？
**答案**：nullptr 标记用于区分"需要遍历的节点"和"需要处理的节点"。当弹出 nullptr 时，表示其后面的节点需要被处理（输出其值）。

#### 4.3 统一迭代法的优缺点是什么？
**答案**：
- **优点**：统一了前中后序遍历的实现方式，代码结构相似，便于理解和记忆
- **缺点**：需要额外的 nullptr 标记，代码相对复杂

### 5. 复杂度分析类问题

#### 5.1 二叉树遍历的时间复杂度是多少？为什么？
**答案**：时间复杂度是 O(n)，其中 n 是二叉树的节点数。因为每个节点都会被访问且仅被访问一次。

#### 5.2 二叉树遍历的空间复杂度是多少？为什么？
**答案**：空间复杂度是 O(n)，其中 n 是二叉树的节点数。
- 递归实现：递归调用栈的深度最坏情况下为 n（树退化为链表时）
- 迭代实现：栈的大小最坏情况下为 n（树退化为链表时）

#### 5.3 有没有办法降低空间复杂度？
**答案**：可以使用 Morris 遍历算法，利用树的空闲指针实现 O(1) 的空间复杂度。但该算法会修改树的结构，且代码复杂。

### 6. 应用场景类问题

#### 6.1 二叉树遍历有哪些实际应用场景？
**答案**：
- **前序遍历**：常用于复制树、序列化二叉树
- **中序遍历**：在二叉搜索树中可以得到有序序列
- **后序遍历**：常用于删除树、计算目录大小等

#### 6.2 什么是二叉搜索树？中序遍历有什么用？
**答案**：二叉搜索树（BST）是一种特殊的二叉树，对于每个节点，其左子树的所有节点值小于该节点，其右子树的所有节点值大于该节点。对二叉搜索树进行中序遍历可以得到节点的升序序列。

#### 6.3 如何根据前序遍历和中序遍历构建二叉树？
**答案**：前序遍历的第一个节点是根节点，在中序遍历中找到该节点，其左边的节点是左子树，右边的节点是右子树。递归构建左右子树。

### 7. 进阶问题

#### 7.1 什么是 Morris 遍历？它有什么特点？
**答案**：Morris 遍历是一种遍历二叉树的方法，利用树的空闲指针（原本为空的右子节点指针）来记录遍历路径，可以在 O(1) 的空间复杂度下完成遍历。但会修改树的结构，且遍历顺序不是标准的前中后序。

#### 7.2 递归和迭代如何选择？
**答案**：
- 树深度不大（一般小于 1000）：选择递归实现，代码简洁
- 树深度很大：选择迭代实现，避免栈溢出
- 需要统一三种遍历的代码结构：选择统一迭代法

#### 7.3 如何在面试中回答二叉树遍历问题？
**答案**：
1. 先解释遍历的概念和三种遍历方式的区别
2. 写出递归实现，解释三要素（参数和返回值、终止条件、单层递归逻辑）
3. 写出迭代实现，解释思路和栈的作用
4. 分析时间复杂度和空间复杂度
5. 如果面试官问到优化，可以提到 Morris 遍历

---

## 九、 代码示例与实践

### 1. 递归实现示例

#### 1.1 前序遍历

```cpp
#include <iostream>
#include <vector>

using namespace std;

/**
 * @brief 二叉树节点结构体定义
 */
struct TreeNode {
    int val;                    // 节点值
    TreeNode *left;             // 左子节点指针
    TreeNode *right;            // 右子节点指针

    // 默认构造函数，节点值为0，左右子节点均为空
    TreeNode() : val(0), left(nullptr), right(nullptr) {}

    // 带值构造函数，节点值为指定值，左右子节点均为空
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}

    // 带值和左右子节点构造函数
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

class Solution {
public:
    /**
     * @brief 递归辅助函数，实现前序遍历（中左右）
     * @param cur 当前遍历的节点指针
     * @param vec 存储遍历结果的向量引用，通过引用传递避免拷贝
     */
    void traversal(TreeNode* cur, vector<int>& vec) {
        // 递归终止条件：当节点为空时，直接返回，不再继续递归
        if (cur == nullptr) return;

        // 处理当前节点（中）：将当前节点的值加入结果向量
        vec.push_back(cur->val);

        // 递归遍历左子树（左）
        traversal(cur->left, vec);

        // 递归遍历右子树（右）
        traversal(cur->right, vec);
    }

    /**
     * @brief 前序遍历入口函数
     * @param root 二叉树的根节点指针
     * @return 包含前序遍历结果的向量
     */
    vector<int> preorderTraversal(TreeNode* root) {
        vector<int> result;  // 存储遍历结果
        traversal(root, result);  // 调用递归辅助函数
        return result;  // 返回遍历结果
    }
};

/**
 * @brief 测试函数，验证前序遍历的正确性
 */
void testPreorder() {
    // 构建测试树
    //       1
    //      / \
    //     2   3
    //    / \   \
    //   4   5   8
    //      / \
    //     6   7
    TreeNode* root = new TreeNode(1);
    root->left = new TreeNode(2);
    root->right = new TreeNode(3);
    root->left->left = new TreeNode(4);
    root->left->right = new TreeNode(5);
    root->right->right = new TreeNode(8);
    root->left->right->left = new TreeNode(6);
    root->left->right->right = new TreeNode(7);
    root->right->right->left = new TreeNode(9);

    // 创建 Solution 对象并调用前序遍历
    Solution solution;
    vector<int> result = solution.preorderTraversal(root);

    // 打印结果
    cout << "前序遍历结果：";
    for (int val : result) {
        cout << val << " ";
    }
    cout << endl;

    // 释放内存（实际应用中需要实现析构函数或使用智能指针）
    // 这里简化处理，仅作为示例
}

int main() {
    testPreorder();  // 运行测试函数
    return 0;
}
```

#### 1.2 中序遍历

```cpp
/**
 * @brief 中序遍历递归实现
 */
class SolutionInorder {
public:
    /**
     * @brief 递归辅助函数，实现中序遍历（左中右）
     * @param cur 当前遍历的节点指针
     * @param vec 存储遍历结果的向量引用
     */
    void traversal(TreeNode* cur, vector<int>& vec) {
        // 递归终止条件：当节点为空时，直接返回
        if (cur == nullptr) return;

        // 递归遍历左子树（左）
        traversal(cur->left, vec);

        // 处理当前节点（中）：将当前节点的值加入结果向量
        vec.push_back(cur->val);

        // 递归遍历右子树（右）
        traversal(cur->right, vec);
    }

    /**
     * @brief 中序遍历入口函数
     * @param root 二叉树的根节点指针
     * @return 包含中序遍历结果的向量
     */
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> result;  // 存储遍历结果
        traversal(root, result);  // 调用递归辅助函数
        return result;  // 返回遍历结果
    }
};
```

#### 1.3 后序遍历

```cpp
/**
 * @brief 后序遍历递归实现
 */
class SolutionPostorder {
public:
    /**
     * @brief 递归辅助函数，实现后序遍历（左右中）
     * @param cur 当前遍历的节点指针
     * @param vec 存储遍历结果的向量引用
     */
    void traversal(TreeNode* cur, vector<int>& vec) {
        // 递归终止条件：当节点为空时，直接返回
        if (cur == nullptr) return;

        // 递归遍历左子树（左）
        traversal(cur->left, vec);

        // 递归遍历右子树（右）
        traversal(cur->right, vec);

        // 处理当前节点（中）：将当前节点的值加入结果向量
        vec.push_back(cur->val);
    }

    /**
     * @brief 后序遍历入口函数
     * @param root 二叉树的根节点指针
     * @return 包含后序遍历结果的向量
     */
    vector<int> postorderTraversal(TreeNode* root) {
        vector<int> result;  // 存储遍历结果
        traversal(root, result);  // 调用递归辅助函数
        return result;  // 返回遍历结果
    }
};
```

### 2. 迭代实现示例

#### 2.1 前序遍历

```cpp
/**
 * @brief 前序遍历迭代实现（使用栈）
 */
class SolutionPreorderIter {
public:
    /**
     * @brief 前序遍历迭代实现
     * @param root 二叉树的根节点指针
     * @return 包含前序遍历结果的向量
     */
    vector<int> preorderTraversal(TreeNode* root) {
        vector<int> result;  // 存储遍历结果

        // 边界情况处理：空树直接返回空结果
        if (root == nullptr) return result;

        // 创建栈用于模拟递归过程
        stack<TreeNode*> st;
        st.push(root);  // 将根节点入栈

        // 开始迭代过程
        while (!st.empty()) {
            // 弹出栈顶节点（相当于处理"中"节点）
            TreeNode* node = st.top();
            st.pop();

            // 将当前节点的值加入结果向量（中）
            result.push_back(node->val);

            // 注意：先压右子节点，后压左子节点
            // 因为栈是LIFO，这样左子节点会先出栈，符合前序遍历顺序
            if (node->right) st.push(node->right);  // 右子节点入栈
            if (node->left) st.push(node->left);   // 左子节点入栈
        }

        return result;  // 返回遍历结果
    }
};
```

#### 2.2 中序遍历

```cpp
/**
 * @brief 中序遍历迭代实现（使用栈）
 */
class SolutionInorderIter {
public:
    /**
     * @brief 中序遍历迭代实现
     * @param root 二叉树的根节点指针
     * @return 包含中序遍历结果的向量
     */
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> result;  // 存储遍历结果
        stack<TreeNode*> st;  // 创建栈用于模拟递归过程
        TreeNode* cur = root;  // 当前遍历指针，初始化为根节点

        // 迭代条件：当前节点不为空或栈不为空
        while (cur != nullptr || !st.empty()) {
            if (cur != nullptr) {
                // 当前节点不为空，将当前节点入栈
                // 然后继续向左走，将所有左子节点入栈
                st.push(cur);
                cur = cur->left;  // 移动到左子节点
            } else {
                // 当前节点为空（已经到达最左叶子节点的左孩子）
                // 弹出栈顶节点进行处理
                cur = st.top();
                st.pop();

                // 处理当前节点（中）：将节点值加入结果向量
                result.push_back(cur->val);

                // 处理右子树：移动到右子节点，继续循环
                cur = cur->right;
            }
        }

        return result;  // 返回遍历结果
    }
};
```

#### 2.3 后序遍历（反转法）

```cpp
/**
 * @brief 后序遍历迭代实现（使用栈 + 反转）
 */
class SolutionPostorderIter {
public:
    /**
     * @brief 后序遍历迭代实现
     * @param root 二叉树的根节点指针
     * @return 包含后序遍历结果的向量
     */
    vector<int> postorderTraversal(TreeNode* root) {
        vector<int> result;  // 存储遍历结果

        // 边界情况处理：空树直接返回空结果
        if (root == nullptr) return result;

        // 创建栈用于模拟递归过程
        stack<TreeNode*> st;
        st.push(root);  // 将根节点入栈

        // 迭代过程：实现"中右左"的遍历顺序
        while (!st.empty()) {
            // 弹出栈顶节点（相当于处理"中"节点）
            TreeNode* node = st.top();
            st.pop();

            // 将当前节点的值加入结果向量（中）
            result.push_back(node->val);

            // 注意：与前序遍历不同，这里先压左子节点，后压右子节点
            // 这样出栈顺序是"中右左"，符合后序遍历的逆序
            if (node->left) st.push(node->left);   // 左子节点入栈
            if (node->right) st.push(node->right); // 右子节点入栈
        }

        // 反转结果：将"中右左"变为"左右中"，得到正确的后序遍历结果
        reverse(result.begin(), result.end());

        return result;  // 返回遍历结果
    }
};
```

---

## 十、 总结与思考

### 1. 核心知识点总结

- **遍历顺序**：前序（中左右）、中序（左中右）、后序（左右中）
- **实现方法**：递归、迭代、统一迭代法
- **时间复杂度**：均为 O(n)，其中 n 是二叉树的节点数
- **空间复杂度**：均为 O(n)，递归实现的栈深度最坏情况下为 n

### 2. 设计思想与技术要点

- **递归思想**：利用函数调用栈实现深度优先遍历
- **栈的应用**：使用栈来模拟递归过程，实现迭代遍历
- **统一迭代法**：使用标记法统一三种遍历的实现方式
- **边界情况处理**：空树、单节点树、退化为链表的树等

### 3. 实际应用价值

- **二叉树遍历**是二叉树相关算法的基础，许多二叉树问题都需要基于遍历解决
- **递归与迭代**的转换是算法学习中的重要内容，理解其转换过程有助于掌握算法设计的本质
- **空间优化**：对于大规模数据，迭代实现和 Morris 遍历可以避免栈溢出问题

### 4. 学习建议与进阶方向

- **基础巩固**：掌握三种遍历的递归和迭代实现，理解其原理
- **实践应用**：通过解决二叉树相关问题，加深对遍历算法的理解
- **进阶学习**：学习 Morris 遍历算法，进一步优化空间复杂度
- **扩展应用**：将遍历算法应用到更复杂的二叉树问题中，如二叉树的序列化与反序列化、路径总和等

### 5. 个人感悟

二叉树的遍历是算法学习中的基础内容，虽然看似简单，但其中蕴含了深刻的算法思想。递归实现的简洁性和迭代实现的灵活性，以及统一迭代法的巧妙设计，都体现了算法设计的艺术。

通过学习二叉树的遍历，我不仅掌握了具体的实现方法，更重要的是理解了递归与迭代之间的联系和转换，这对于解决更复杂的算法问题具有重要意义。

在实际应用中，我会根据具体情况选择合适的遍历方法：对于简单问题，使用递归实现；对于深度较大的树，使用迭代实现；对于需要统一处理三种遍历的场景，使用统一迭代法。

---

## 十一、 参考文献

1. **代码随想录** - 二叉树理论基础篇
2. **代码随想录** - 二叉树的递归遍历
3. **LeetCode** - 144. 二叉树的前序遍历
4. **LeetCode** - 94. 二叉树的中序遍历
5. **LeetCode** - 145. 二叉树的后序遍历

---

希望这份学习笔记能帮助你更好地理解二叉树的遍历算法，掌握递归与迭代的实现方法，为解决更复杂的二叉树问题打下坚实的基础！