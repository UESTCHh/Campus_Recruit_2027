# 226. 翻转二叉树学习笔记

## 一、 题目分析

### 1.1 题目描述

给你一棵二叉树的根节点 `root` ，翻转这棵二叉树，并返回其根节点。

### 1.2 题目链接

- [LeetCode 226. 翻转二叉树](https://leetcode.cn/problems/invert-binary-tree/description/)
- [代码随想录题解](https://www.programmercarl.com/0226.%E7%BF%BB%E8%BD%AC%E4%BA%8C%E5%8F%89%E6%A0%91.html)

### 1.3 示例

**示例 1：**

```
输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]
```

**示例 2：**

```
输入：root = [2,1,3]
输出：[2,3,1]
```

**示例 3：**

```
输入：root = []
输出：[]
```

### 1.4 提示

- 树中节点数目范围在 `[0, 100]` 内
- `-100 <= Node.val <= 100`

### 1.5 题目本质

翻转二叉树的本质是：**将每个节点的左右子树进行交换**。

### 1.6 有趣的小插曲

这道题目背后有一个让程序员心酸的故事：听说 Homebrew 的作者 Max Howell，就是因为没在白板上写出翻转二叉树，最后被 Google 拒绝了。（真假不做判断，全当一个乐子哈）

这个故事告诉我们：
- 即使是简单的题目，也需要认真对待
- 基础数据结构和算法知识非常重要
- 面试前一定要复习常见的树操作题目

## 二、 解题思路

### 2.1 核心思路

翻转二叉树的核心思路非常简单：**遍历二叉树的每个节点，交换其左右子节点**。

关键在于选择合适的遍历顺序。常见的遍历方式包括：

| 遍历方式 | 是否适用 | 说明 |
|---------|---------|------|
| 前序遍历 | 适用 | 先交换当前节点的左右子节点，再递归处理左右子树 |
| 后序遍历 | 适用 | 先递归处理左右子树，再交换当前节点的左右子节点 |
| 中序遍历 | 不推荐 | 会导致某些节点的左右子节点被翻转两次 |
| 层序遍历 | 适用 | 按层处理每个节点，交换其左右子节点 |

### 2.2 思路详解

1. **递归法（前序遍历）**：
   - 确定递归函数的参数和返回值：参数为当前节点指针，返回值为翻转后的节点指针
   - 确定终止条件：当前节点为空时返回
   - 确定单层递归逻辑：先交换当前节点的左右子节点，再递归处理左子树和右子树

2. **迭代法（栈实现前序遍历）**：
   - 使用栈来模拟递归过程
   - 每次弹出一个节点，交换其左右子节点
   - 将右子节点和左子节点依次入栈（注意顺序）

3. **迭代法（队列实现层序遍历）**：
   - 使用队列进行层序遍历
   - 每次从队列中取出一个节点，交换其左右子节点
   - 将交换后的左右子节点入队

## 三、 算法实现

### 3.1 递归法（前序遍历）

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
    TreeNode* invertTree(TreeNode* root) {
        // 终止条件：当前节点为空，直接返回
        if (root == nullptr) return root;
        
        // 前序遍历：先处理当前节点，交换左右子节点
        swap(root->left, root->right);
        
        // 递归处理左子树
        invertTree(root->left);
        
        // 递归处理右子树
        invertTree(root->right);
        
        // 返回翻转后的根节点
        return root;
    }
};
```

**代码分析：**
- **终止条件**：当当前节点为空时，直接返回，这是递归的边界
- **交换逻辑**：使用 `swap` 函数交换当前节点的左右子节点
- **递归调用**：分别递归处理左子树和右子树
- **返回值**：返回翻转后的根节点

### 3.2 迭代法（栈实现前序遍历）

```cpp
class Solution {
public:
    TreeNode* invertTree(TreeNode* root) {
        // 边界条件：根节点为空，直接返回
        if (root == nullptr) return root;
        
        // 使用栈进行迭代遍历
        stack<TreeNode*> st;
        st.push(root);
        
        while (!st.empty()) {
            // 弹出当前节点（中）
            TreeNode* node = st.top();
            st.pop();
            
            // 交换当前节点的左右子节点
            swap(node->left, node->right);
            
            // 注意：栈是后进先出，所以先压右子节点，再压左子节点
            if (node->right) st.push(node->right);  // 右
            if (node->left) st.push(node->left);    // 左
        }
        
        return root;
    }
};
```

**代码分析：**
- **栈的使用**：栈用于模拟递归调用的过程
- **遍历顺序**：前序遍历（中-左-右），但由于栈是后进先出，所以先压右子节点，再压左子节点
- **交换逻辑**：每次弹出节点后立即交换其左右子节点

### 3.3 迭代法（队列实现层序遍历）

```cpp
class Solution {
public:
    TreeNode* invertTree(TreeNode* root) {
        // 边界条件：根节点为空，直接返回
        if (root == nullptr) return root;
        
        // 使用队列进行层序遍历
        queue<TreeNode*> que;
        que.push(root);
        
        while (!que.empty()) {
            // 取出队首节点
            TreeNode* node = que.front();
            que.pop();
            
            // 交换当前节点的左右子节点
            swap(node->left, node->right);
            
            // 将交换后的左右子节点入队（如果存在）
            if (node->left) que.push(node->left);
            if (node->right) que.push(node->right);
        }
        
        return root;
    }
};
```

**代码分析：**
- **队列的使用**：队列用于按层处理节点
- **遍历顺序**：层序遍历，按层从上到下，每层从左到右
- **交换逻辑**：处理每个节点时交换其左右子节点

## 四、 时间复杂度与空间复杂度分析

### 4.1 时间复杂度

三种解法的时间复杂度都是 **O(n)**，其中 n 是二叉树的节点数。

**原因分析：**
- 每个节点都被访问一次
- 每个节点的交换操作是 O(1)
- 所以整体时间复杂度为 O(n)

### 4.2 空间复杂度

| 解法 | 空间复杂度 | 分析 |
|------|-----------|------|
| 递归法 | O(h) | h 是树的高度，递归栈的深度 |
| 栈迭代法 | O(h) | 栈的最大深度为树的高度 |
| 队列迭代法 | O(n) | 队列最多存储一层的所有节点，最坏情况为 n/2 |

**详细分析：**

1. **递归法**：
   - 空间复杂度为 O(h)，其中 h 是树的高度
   - 对于平衡二叉树，h = log(n)
   - 对于斜树（退化为链表），h = n

2. **栈迭代法**：
   - 空间复杂度为 O(h)，栈的最大深度为树的高度
   - 与递归法类似

3. **队列迭代法**：
   - 空间复杂度为 O(n)
   - 在最坏情况下（完全二叉树），队列需要存储最后一层的所有节点
   - 最后一层最多有 n/2 个节点，所以空间复杂度为 O(n)

## 五、 关键代码注释

### 5.1 递归法关键点

```cpp
// 1. 终止条件：当前节点为空时返回
if (root == nullptr) return root;

// 2. 前序遍历的核心：先处理当前节点
swap(root->left, root->right);

// 3. 递归处理左子树和右子树
invertTree(root->left);
invertTree(root->right);
```

### 5.2 栈迭代法关键点

```cpp
// 1. 使用栈存储待处理的节点
stack<TreeNode*> st;
st.push(root);

while (!st.empty()) {
    // 2. 弹出当前节点
    TreeNode* node = st.top();
    st.pop();
    
    // 3. 交换左右子节点
    swap(node->left, node->right);
    
    // 4. 注意入栈顺序：先右后左，保证出栈时先左后右
    if (node->right) st.push(node->right);
    if (node->left) st.push(node->left);
}
```

### 5.3 队列迭代法关键点

```cpp
// 1. 使用队列进行层序遍历
queue<TreeNode*> que;
que.push(root);

while (!que.empty()) {
    // 2. 取出队首节点
    TreeNode* node = que.front();
    que.pop();
    
    // 3. 交换左右子节点
    swap(node->left, node->right);
    
    // 4. 将子节点入队，保持层序顺序
    if (node->left) que.push(node->left);
    if (node->right) que.push(node->right);
}
```

## 六、 测试用例验证

### 6.1 测试用例 1

**输入**：`root = [4,2,7,1,3,6,9]`

**预期输出**：`[4,7,2,9,6,3,1]`

**执行过程**：
```
原始树：
      4
    /   
   2     7
  / \   / \
 1   3 6   9

翻转后：
      4
    /   
   7     2
  / \   / \
 9   6 3   1
```

### 6.2 测试用例 2

**输入**：`root = [2,1,3]`

**预期输出**：`[2,3,1]`

**执行过程**：
```
原始树：
    2
   / 
  1   3

翻转后：
    2
   / 
  3   1
```

### 6.3 测试用例 3

**输入**：`root = []`

**预期输出**：`[]`

**执行过程**：空树直接返回空。

### 6.4 测试用例 4

**输入**：`root = [1]`

**预期输出**：`[1]`

**执行过程**：单节点树无需翻转。

## 七、 常见错误与注意事项

### 7.1 中序遍历的陷阱

中序遍历不适合翻转二叉树，因为会导致某些节点被翻转两次。

```cpp
// 错误示例：中序遍历会导致错误结果
TreeNode* invertTree(TreeNode* root) {
    if (root == nullptr) return root;
    invertTree(root->left);      // 左
    swap(root->left, root->right); // 中（此时左子树已经被翻转）
    invertTree(root->right);     // 右（此时右子树实际上是原来的左子树）
    return root;
}
```

**问题分析**：在中序遍历中，当执行到 `swap` 后，`root->right` 实际上指向的是原来的左子树，导致该子树被翻转两次。

### 7.2 边界条件处理

务必处理根节点为空的情况：

```cpp
if (root == nullptr) return root;
```

### 7.3 入栈顺序

使用栈实现迭代时，注意入栈顺序：

```cpp
// 正确：先压右，再压左
if (node->right) st.push(node->right);
if (node->left) st.push(node->left);

// 错误：先压左，再压右，会改变遍历顺序
// if (node->left) st.push(node->left);
// if (node->right) st.push(node->right);
```

## 八、 面试相关内容

### 8.1 面试常见问题

#### 8.1.1 翻转二叉树的几种方法？

**答案**：
- **递归法**：使用前序或后序遍历，递归交换每个节点的左右子节点
- **迭代法（栈）**：使用栈模拟递归过程
- **迭代法（队列）**：使用队列进行层序遍历

#### 8.1.2 为什么中序遍历不适合翻转二叉树？

**答案**：
- 中序遍历的顺序是左-中-右
- 在交换当前节点的左右子节点后，原来的左子树变成了右子树
- 后续递归处理右子树时，实际上处理的是原来的左子树
- 这会导致某些节点被翻转两次，最终结果不正确

#### 8.1.3 时间复杂度和空间复杂度是多少？

**答案**：
- **时间复杂度**：O(n)，每个节点访问一次
- **空间复杂度**：递归法和栈迭代法为 O(h)，队列迭代法为 O(n)

### 8.2 扩展问题

#### 8.2.1 翻转 N 叉树

**思路**：对于 N 叉树，需要交换所有子节点的顺序。

```cpp
class Node {
public:
    int val;
    vector<Node*> children;
    
    Node() {}
    Node(int _val) : val(_val) {}
    Node(int _val, vector<Node*> _children) : val(_val), children(_children) {}
};

Node* invertTree(Node* root) {
    if (root == nullptr) return root;
    
    // 反转子节点顺序
    reverse(root->children.begin(), root->children.end());
    
    // 递归处理每个子节点
    for (auto child : root->children) {
        invertTree(child);
    }
    
    return root;
}
```

#### 8.2.2 判断两棵树是否互为镜像

**思路**：可以翻转其中一棵树，然后判断是否与另一棵树相同。

```cpp
bool isMirror(TreeNode* p, TreeNode* q) {
    if (p == nullptr && q == nullptr) return true;
    if (p == nullptr || q == nullptr) return false;
    
    // 翻转 p
    invertTree(p);
    
    // 判断是否相等
    return isSameTree(p, q);
}

bool isSameTree(TreeNode* p, TreeNode* q) {
    if (p == nullptr && q == nullptr) return true;
    if (p == nullptr || q == nullptr) return false;
    return p->val == q->val && isSameTree(p->left, q->left) && isSameTree(p->right, q->right);
}
```

## 九、 个人总结与反思

### 9.1 学习收获

1. **问题本质理解**：翻转二叉树的本质是交换每个节点的左右子节点，这是一个非常直观的问题。

2. **遍历顺序的重要性**：选择正确的遍历顺序对于解决问题至关重要，中序遍历在本题中不适用。

3. **多种解法的掌握**：掌握了递归法、栈迭代法和队列迭代法三种解法，理解了它们之间的联系和区别。

4. **边界条件处理**：深刻认识到处理边界条件（如空树）的重要性。

### 9.2 解题技巧

1. **递归思路**：对于树相关的问题，递归是一种非常自然的思路，需要掌握递归三部曲：
   - 确定递归函数的参数和返回值
   - 确定终止条件
   - 确定单层递归的逻辑

2. **迭代思路**：当递归深度过大可能导致栈溢出时，需要使用迭代方法。栈用于深度优先遍历，队列用于广度优先遍历。

3. **代码简洁性**：本题的代码非常简洁，核心逻辑就是 `swap(root->left, root->right)`，体现了算法之美。

### 9.3 后续学习建议

1. **练习更多树相关的题目**：树是面试中的高频考点，需要多练习各种遍历方式和操作。

2. **深入理解递归**：递归是树操作的核心，需要深入理解递归的执行过程。

3. **尝试优化空间复杂度**：思考是否可以在 O(1) 空间复杂度内完成翻转操作（原地翻转）。

4. **学习相关问题**：如判断对称二叉树、合并二叉树等，加深对树操作的理解。

### 9.4 经典故事

这道题目背后有一个有趣的故事：Homebrew 的作者 Max Howell 据说因为没能在白板上写出翻转二叉树，最后被 Google 拒绝了。这说明即使是简单的问题，也需要认真对待，确保自己真正理解了问题的本质。

---

通过本次学习，我对二叉树的翻转操作有了全面的理解，掌握了多种解法，并深入分析了时间复杂度和空间复杂度。这道题虽然简单，但包含了很多重要的知识点，是学习树操作的很好入门题目。