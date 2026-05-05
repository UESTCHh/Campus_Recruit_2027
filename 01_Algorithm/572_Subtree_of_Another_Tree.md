# 572. 另一棵树的子树学习笔记

## 一、题目分析

### 1.1 题目描述

给你两棵二叉树 `root` 和 `subRoot`。检验 `root` 中是否包含和 `subRoot` 具有相同结构和节点值的子树。如果存在，返回 `true`；否则，返回 `false`。

二叉树 `tree` 的一棵子树包括 `tree` 的某个节点和这个节点的所有后代节点。`tree` 也可以看做它自身的一棵子树。

### 1.2 题目链接

- [LeetCode 572. 另一棵树的子树](https://leetcode.cn/problems/subtree-of-another-tree/description/)
- [代码随想录题解](https://www.programmercarl.com/)
- [官方题解](https://leetcode.cn/problems/subtree-of-another-tree/solutions/233896/ling-yi-ge-shu-de-zi-shu-by-leetcode-solution/)

### 1.3 示例

**示例 1：**

```
输入：root = [3,4,5,1,2], subRoot = [4,1,2]
输出：true
```

**图示：**
```
      root                    subRoot
        3                       4
       / \                     / \
      4   5                   1   2
     / \
    1   2
```

**示例 2：**

```
输入：root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
输出：false
```

**图示：**
```
      root                    subRoot
        3                       4
       / \                     / \
      4   5                   1   2
     / \
    1   2
       /
      0
```

### 1.4 提示

- `root` 树上的节点数量范围是 `[1, 2000]`
- `subRoot` 树上的节点数量范围是 `[1, 1000]`
- `-10^4 <= root.val <= 10^4`
- `-10^4 <= subRoot.val <= 10^4`

### 1.5 题目本质

判断一棵树是否是另一棵树的子树，本质上是：**在 `root` 树中查找是否存在一个节点，使得以该节点为根的子树与 `subRoot` 完全相同**。

这需要两个核心操作：
1. **遍历 `root` 树**：找到所有可能的候选节点
2. **比较两棵树是否相同**：判断候选节点的子树是否与 `subRoot` 相同

### 1.6 题目进阶

你可以运用递归和迭代两种方法解决这个问题吗？

---

## 二、解题思路分析

### 2.1 核心思路

判断 `subRoot` 是否是 `root` 的子树，可以分为以下步骤：

```
┌─────────────────────────────────────────────────────────┐
│                    解题流程                            │
├─────────────────────────────────────────────────────────┤
│  1. 如果 root 为空，返回 false（除非 subRoot 也为空）   │
│  2. 检查当前 root 是否与 subRoot 相同                  │
│  3. 如果相同，返回 true                                │
│  4. 如果不同，递归检查左子树                           │
│  5. 如果左子树中找到，返回 true                        │
│  6. 如果没找到，递归检查右子树                         │
│  7. 返回右子树的检查结果                              │
└─────────────────────────────────────────────────────────┘
```

### 2.2 关键子问题

**问题：如何判断两棵树是否完全相同？**

这是本问题的核心子问题，需要判断两棵树的结构和节点值是否完全一致。

判断两棵树相同的逻辑：
1. 如果两棵树都为空，返回 `true`
2. 如果其中一棵为空，另一棵不为空，返回 `false`
3. 如果两棵树都不为空，但值不同，返回 `false`
4. 递归比较左子树和右子树

### 2.3 递归法思路

#### 2.3.1 递归函数设计

**主函数：isSubtree**
- 参数：`root`（主树的当前节点），`subRoot`（子树的根节点）
- 返回值：`bool`，表示 `subRoot` 是否是 `root` 的子树

**辅助函数：isSameTree**
- 参数：`p`（第一棵树的节点），`q`（第二棵树的节点）
- 返回值：`bool`，表示两棵树是否相同

#### 2.3.2 递归终止条件

**isSubtree 终止条件：**
- 如果 `root` 为空，返回 `false`（因为题目保证 `subRoot` 不为空）

**isSameTree 终止条件：**
- 如果 `p` 和 `q` 都为空，返回 `true`
- 如果 `p` 或 `q` 其中一个为空，返回 `false`
- 如果 `p->val != q->val`，返回 `false`

#### 2.3.3 单层递归逻辑

**isSubtree 单层逻辑：**
1. 检查当前节点的子树是否与 `subRoot` 相同
2. 如果相同，返回 `true`
3. 如果不同，递归检查左子树
4. 如果左子树中找到，返回 `true`
5. 如果没找到，递归检查右子树
6. 返回右子树的检查结果

**isSameTree 单层逻辑：**
1. 比较当前节点的值
2. 递归比较左子树
3. 递归比较右子树
4. 只有左右子树都相同，当前树才相同

### 2.4 迭代法思路

#### 2.4.1 使用队列进行层序遍历

迭代法的思路是：
1. 使用队列对 `root` 进行层序遍历
2. 对于每个出队的节点，检查以该节点为根的子树是否与 `subRoot` 相同
3. 如果找到匹配的，返回 `true`
4. 如果遍历完所有节点都没有找到，返回 `false`

#### 2.4.2 迭代法步骤

```
┌─────────────────────────────────────────────────────────┐
│                   迭代法流程                           │
├─────────────────────────────────────────────────────────┤
│  1. 创建队列，将 root 入队                            │
│  2. 当队列不为空时：                                  │
│     a. 出队一个节点                                   │
│     b. 检查该节点的子树是否与 subRoot 相同             │
│     c. 如果相同，返回 true                            │
│     d. 如果该节点有左孩子，入队                        │
│     e. 如果该节点有右孩子，入队                        │
│  3. 遍历完成，返回 false                             │
└─────────────────────────────────────────────────────────┘
```

### 2.5 优化思路（树哈希）

#### 2.5.1 问题分析

上述方法的时间复杂度是 O(n*m)，其中 n 是 `root` 的节点数，m 是 `subRoot` 的节点数。对于每个节点，我们都需要进行一次树比较。

当树很大时，这种方法效率较低。我们可以使用**树哈希**的方法将时间复杂度优化到 O(n + m)。

#### 2.5.2 树哈希的思路

**核心思想：**
1. 对 `root` 和 `subRoot` 分别进行哈希计算
2. 将 `root` 中每个节点的子树哈希值存储起来
3. 检查 `subRoot` 的哈希值是否存在于 `root` 的哈希集合中

**哈希计算方法：**
- 使用后序遍历计算每个子树的哈希值
- 哈希值 = f(left_hash, right_hash, val)
- 使用一个大质数作为基数，另一个大质数作为模数

---

## 三、算法实现代码

### 3.1 递归法实现

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
    // ==========================================
    // 主函数：判断 subRoot 是否是 root 的子树
    // ==========================================
    bool isSubtree(TreeNode* root, TreeNode* subRoot) {
        // 边界条件：如果 root 为空，直接返回 false
        if (root == nullptr) {
            return false;
        }
        
        // 检查当前节点的子树是否与 subRoot 相同
        if (isSameTree(root, subRoot)) {
            return true;
        }
        
        // 递归检查左子树和右子树
        return isSubtree(root->left, subRoot) || isSubtree(root->right, subRoot);
    }

private:
    // ==========================================
    // 辅助函数：判断两棵树是否完全相同
    // ==========================================
    bool isSameTree(TreeNode* p, TreeNode* q) {
        // 情况1：两棵树都为空，相同
        if (p == nullptr && q == nullptr) {
            return true;
        }
        
        // 情况2：其中一棵树为空，另一棵不为空，不相同
        if (p == nullptr || q == nullptr) {
            return false;
        }
        
        // 情况3：两棵树都不为空，但值不同，不相同
        if (p->val != q->val) {
            return false;
        }
        
        // 情况4：递归比较左子树和右子树
        return isSameTree(p->left, q->left) && isSameTree(p->right, q->right);
    }
};
```

### 3.2 迭代法实现

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
#include <queue>
using namespace std;

class Solution {
public:
    // ==========================================
    // 主函数：使用迭代法判断 subRoot 是否是 root 的子树
    // ==========================================
    bool isSubtree(TreeNode* root, TreeNode* subRoot) {
        // 如果 root 为空，直接返回 false
        if (root == nullptr) {
            return false;
        }
        
        // 创建队列，进行层序遍历
        queue<TreeNode*> q;
        q.push(root);
        
        // 遍历队列中的每个节点
        while (!q.empty()) {
            // 出队一个节点
            TreeNode* current = q.front();
            q.pop();
            
            // 检查当前节点的子树是否与 subRoot 相同
            if (isSameTree(current, subRoot)) {
                return true;
            }
            
            // 如果有左孩子，入队
            if (current->left != nullptr) {
                q.push(current->left);
            }
            
            // 如果有右孩子，入队
            if (current->right != nullptr) {
                q.push(current->right);
            }
        }
        
        // 遍历完所有节点都没有找到，返回 false
        return false;
    }

private:
    // ==========================================
    // 辅助函数：判断两棵树是否完全相同（迭代版）
    // ==========================================
    bool isSameTree(TreeNode* p, TreeNode* q) {
        // 如果两棵树都为空，相同
        if (p == nullptr && q == nullptr) {
            return true;
        }
        
        // 如果其中一棵树为空，另一棵不为空，不相同
        if (p == nullptr || q == nullptr) {
            return false;
        }
        
        // 使用队列进行层序遍历比较
        queue<TreeNode*> q1, q2;
        q1.push(p);
        q2.push(q);
        
        while (!q1.empty() && !q2.empty()) {
            TreeNode* node1 = q1.front();
            TreeNode* node2 = q2.front();
            q1.pop();
            q2.pop();
            
            // 比较节点值
            if (node1->val != node2->val) {
                return false;
            }
            
            // 处理左孩子
            bool left1 = (node1->left != nullptr);
            bool left2 = (node2->left != nullptr);
            
            if (left1 != left2) {
                return false;  // 一个有左孩子，一个没有
            }
            if (left1) {
                q1.push(node1->left);
                q2.push(node2->left);
            }
            
            // 处理右孩子
            bool right1 = (node1->right != nullptr);
            bool right2 = (node2->right != nullptr);
            
            if (right1 != right2) {
                return false;  // 一个有右孩子，一个没有
            }
            if (right1) {
                q1.push(node1->right);
                q2.push(node2->right);
            }
        }
        
        // 如果两个队列都为空，说明完全相同
        return q1.empty() && q2.empty();
    }
};
```

### 3.3 树哈希优化实现（O(n + m) 时间复杂度）

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
#include <unordered_set>
using namespace std;

class Solution {
public:
    // ==========================================
    // 主函数：使用树哈希判断 subRoot 是否是 root 的子树
    // ==========================================
    bool isSubtree(TreeNode* root, TreeNode* subRoot) {
        // 计算 subRoot 的哈希值
        string subHash = hashTree(subRoot);
        
        // 存储 root 中所有子树的哈希值
        unordered_set<string> hashSet;
        hashTreeWithSet(root, hashSet);
        
        // 检查 subRoot 的哈希值是否存在
        return hashSet.count(subHash);
    }

private:
    // ==========================================
    // 计算树的哈希值（后序遍历）
    // ==========================================
    string hashTree(TreeNode* node) {
        if (node == nullptr) {
            return "#";  // 空节点标记
        }
        
        // 后序遍历：左-右-中
        string left = hashTree(node->left);
        string right = hashTree(node->right);
        
        // 构造哈希字符串
        return "(" + left + ")" + to_string(node->val) + "(" + right + ")";
    }
    
    // ==========================================
    // 计算树的哈希值并存储到集合中
    // ==========================================
    string hashTreeWithSet(TreeNode* node, unordered_set<string>& hashSet) {
        if (node == nullptr) {
            return "#";
        }
        
        // 后序遍历
        string left = hashTreeWithSet(node->left, hashSet);
        string right = hashTreeWithSet(node->right, hashSet);
        
        // 构造哈希字符串
        string hash = "(" + left + ")" + to_string(node->val) + "(" + right + ")";
        
        // 将当前子树的哈希值加入集合
        hashSet.insert(hash);
        
        return hash;
    }
};
```

### 3.4 优化的树哈希实现（使用数字哈希）

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
#include <unordered_set>
using namespace std;

class Solution {
public:
    // 大质数，用于哈希计算
    const int BASE = 911382629;
    const int MOD = 1000000007;
    
    bool isSubtree(TreeNode* root, TreeNode* subRoot) {
        // 计算 subRoot 的哈希值
        long long subHash = hashTree(subRoot);
        
        // 存储 root 中所有子树的哈希值
        unordered_set<long long> hashSet;
        hashTreeWithSet(root, hashSet);
        
        // 检查 subRoot 的哈希值是否存在
        return hashSet.count(subHash);
    }

private:
    // ==========================================
    // 计算树的数字哈希值（后序遍历）
    // ==========================================
    long long hashTree(TreeNode* node) {
        if (node == nullptr) {
            return 1;  // 空节点的哈希值
        }
        
        // 后序遍历
        long long left = hashTree(node->left);
        long long right = hashTree(node->right);
        
        // 计算哈希值：(left * BASE + val) * BASE + right
        return ((left * BASE % MOD + (node->val + 10000)) % MOD * BASE % MOD + right) % MOD;
    }
    
    // ==========================================
    // 计算树的哈希值并存储到集合中
    // ==========================================
    long long hashTreeWithSet(TreeNode* node, unordered_set<long long>& hashSet) {
        if (node == nullptr) {
            return 1;
        }
        
        // 后序遍历
        long long left = hashTreeWithSet(node->left, hashSet);
        long long right = hashTreeWithSet(node->right, hashSet);
        
        // 计算哈希值
        long long hash = ((left * BASE % MOD + (node->val + 10000)) % MOD * BASE % MOD + right) % MOD;
        
        // 将当前子树的哈希值加入集合
        hashSet.insert(hash);
        
        return hash;
    }
};
```

---

## 四、代码分析

### 4.1 递归法代码分析

#### 4.1.1 isSubtree 函数

```cpp
bool isSubtree(TreeNode* root, TreeNode* subRoot) {
    if (root == nullptr) {
        return false;
    }
    
    if (isSameTree(root, subRoot)) {
        return true;
    }
    
    return isSubtree(root->left, subRoot) || isSubtree(root->right, subRoot);
}
```

**分析：**
- **边界检查**：如果 `root` 为空，直接返回 `false`
- **当前节点检查**：调用 `isSameTree` 检查当前节点的子树是否与 `subRoot` 相同
- **递归检查**：如果当前节点不匹配，递归检查左子树和右子树

#### 4.1.2 isSameTree 函数

```cpp
bool isSameTree(TreeNode* p, TreeNode* q) {
    if (p == nullptr && q == nullptr) {
        return true;
    }
    
    if (p == nullptr || q == nullptr) {
        return false;
    }
    
    if (p->val != q->val) {
        return false;
    }
    
    return isSameTree(p->left, q->left) && isSameTree(p->right, q->right);
}
```

**分析：**
- **空节点处理**：两棵树都为空则相同，只有一棵为空则不同
- **值比较**：节点值不同则不同
- **递归比较**：需要左右子树都相同才返回 `true`

### 4.2 迭代法代码分析

#### 4.2.1 主函数

```cpp
bool isSubtree(TreeNode* root, TreeNode* subRoot) {
    if (root == nullptr) {
        return false;
    }
    
    queue<TreeNode*> q;
    q.push(root);
    
    while (!q.empty()) {
        TreeNode* current = q.front();
        q.pop();
        
        if (isSameTree(current, subRoot)) {
            return true;
        }
        
        if (current->left != nullptr) {
            q.push(current->left);
        }
        
        if (current->right != nullptr) {
            q.push(current->right);
        }
    }
    
    return false;
}
```

**分析：**
- **队列初始化**：将 `root` 入队
- **层序遍历**：依次处理每个节点
- **子树比较**：对每个节点调用 `isSameTree`
- **子节点入队**：将左右子节点加入队列

#### 4.2.2 迭代版 isSameTree

```cpp
bool isSameTree(TreeNode* p, TreeNode* q) {
    if (p == nullptr && q == nullptr) {
        return true;
    }
    
    if (p == nullptr || q == nullptr) {
        return false;
    }
    
    queue<TreeNode*> q1, q2;
    q1.push(p);
    q2.push(q);
    
    while (!q1.empty() && !q2.empty()) {
        TreeNode* node1 = q1.front();
        TreeNode* node2 = q2.front();
        q1.pop();
        q2.pop();
        
        if (node1->val != node2->val) {
            return false;
        }
        
        bool left1 = (node1->left != nullptr);
        bool left2 = (node2->left != nullptr);
        
        if (left1 != left2) {
            return false;
        }
        if (left1) {
            q1.push(node1->left);
            q2.push(node2->left);
        }
        
        bool right1 = (node1->right != nullptr);
        bool right2 = (node2->right != nullptr);
        
        if (right1 != right2) {
            return false;
        }
        if (right1) {
            q1.push(node1->right);
            q2.push(node2->right);
        }
    }
    
    return q1.empty() && q2.empty();
}
```

**分析：**
- **双队列比较**：使用两个队列分别存储两棵树的节点
- **节点值比较**：每次出队两个节点进行比较
- **子节点处理**：检查两个节点的子节点情况是否一致

### 4.3 树哈希代码分析

#### 4.3.1 字符串哈希

```cpp
string hashTree(TreeNode* node) {
    if (node == nullptr) {
        return "#";
    }
    
    string left = hashTree(node->left);
    string right = hashTree(node->right);
    
    return "(" + left + ")" + to_string(node->val) + "(" + right + ")";
}
```

**分析：**
- **空节点标记**：使用 `"#"` 表示空节点
- **后序遍历**：左-右-中的顺序
- **哈希字符串构造**：使用括号包裹子树哈希，避免歧义

#### 4.3.2 数字哈希

```cpp
long long hashTree(TreeNode* node) {
    if (node == nullptr) {
        return 1;
    }
    
    long long left = hashTree(node->left);
    long long right = hashTree(node->right);
    
    return ((left * BASE % MOD + (node->val + 10000)) % MOD * BASE % MOD + right) % MOD;
}
```

**分析：**
- **空节点哈希**：返回 1
- **哈希公式**：`((left * BASE + val) * BASE + right) % MOD`
- **值调整**：`val + 10000` 将负数转为正数

---

## 五、时间与空间复杂度分析

### 5.1 递归法复杂度分析

| 复杂度类型 | 时间复杂度 | 空间复杂度 |
|-----------|-----------|-----------|
| **最好情况** | O(m) | O(min(n, m)) |
| **最坏情况** | O(n*m) | O(min(n, m)) |

**说明：**
- **最好情况**：`root` 的根节点就是 `subRoot`，只需比较一次
- **最坏情况**：`root` 的每个节点都需要与 `subRoot` 比较，每次比较需要 O(m) 时间
- **空间复杂度**：递归调用栈的深度，取决于两棵树中较浅的那棵

### 5.2 迭代法复杂度分析

| 复杂度类型 | 时间复杂度 | 空间复杂度 |
|-----------|-----------|-----------|
| **最好情况** | O(m) | O(n) |
| **最坏情况** | O(n*m) | O(n) |

**说明：**
- **时间复杂度**：与递归法相同
- **空间复杂度**：队列存储的节点数，最坏情况下为 O(n)

### 5.3 树哈希复杂度分析

| 复杂度类型 | 时间复杂度 | 空间复杂度 |
|-----------|-----------|-----------|
| **最好情况** | O(n + m) | O(n) |
| **最坏情况** | O(n + m) | O(n) |

**说明：**
- **时间复杂度**：遍历 `root` 一次（O(n)），遍历 `subRoot` 一次（O(m)），哈希查找 O(1)
- **空间复杂度**：存储 `root` 中所有子树的哈希值，O(n)

### 5.4 复杂度对比表格

| 方法 | 时间复杂度（平均） | 时间复杂度（最坏） | 空间复杂度 |
|------|-------------------|-------------------|-----------|
| 递归法 | O(n*m) | O(n*m) | O(min(n, m)) |
| 迭代法 | O(n*m) | O(n*m) | O(n) |
| 树哈希 | O(n + m) | O(n + m) | O(n) |

---

## 六、测试用例验证

### 6.1 测试用例 1

**输入：**
```
root = [3,4,5,1,2], subRoot = [4,1,2]
```

**执行过程：**
1. 检查 `root`（值为 3）与 `subRoot`（值为 4），不相同
2. 递归检查左子树（值为 4）与 `subRoot`
3. 值相同，继续比较左子树（1 vs 1）和右子树（2 vs 2）
4. 所有节点都匹配，返回 `true`

**输出：** `true`

### 6.2 测试用例 2

**输入：**
```
root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
```

**执行过程：**
1. 检查 `root`（值为 3）与 `subRoot`（值为 4），不相同
2. 递归检查左子树（值为 4）与 `subRoot`
3. 值相同，比较左子树（1 vs 1），相同
4. 比较右子树：`root` 的右子树有子节点 0，`subRoot` 的右子树为空
5. 不匹配，返回 `false`
6. 继续检查其他节点，都不匹配
7. 最终返回 `false`

**输出：** `false`

### 6.3 测试用例 3

**输入：**
```
root = [1], subRoot = [1]
```

**执行过程：**
1. 检查 `root`（值为 1）与 `subRoot`（值为 1）
2. 两个节点都没有子节点，完全相同
3. 返回 `true`

**输出：** `true`

### 6.4 测试用例 4

**输入：**
```
root = [1,2], subRoot = [1]
```

**执行过程：**
1. 检查 `root`（值为 1）与 `subRoot`（值为 1）
2. `root` 有左子节点 2，`subRoot` 没有子节点
3. 不匹配
4. 检查左子树（值为 2）与 `subRoot`（值为 1），不匹配
5. 返回 `false`

**输出：** `false`

---

## 七、常见错误与注意事项

### 7.1 空节点处理

**错误示例：**
```cpp
// 错误：没有处理 subRoot 为空的情况
bool isSameTree(TreeNode* p, TreeNode* q) {
    if (p->val != q->val) {  // 如果 p 或 q 为空，这里会崩溃！
        return false;
    }
    // ...
}
```

**正确做法：**
```cpp
bool isSameTree(TreeNode* p, TreeNode* q) {
    if (p == nullptr && q == nullptr) {
        return true;
    }
    if (p == nullptr || q == nullptr) {
        return false;
    }
    if (p->val != q->val) {
        return false;
    }
    // ...
}
```

### 7.2 子树与子结构的区别

**重要区分：**
- **子树**：必须包含节点的所有后代节点
- **子结构**：只需要匹配部分结构

**示例：**
```
root:      3
          / \
         4   5
        / \
       1   2
          /
         0

subRoot:   4
          / \
         1   2
```

- `subRoot` **不是** `root` 的子树（因为 root 的 4 节点有额外的子节点 0）
- `subRoot` **是** `root` 的子结构

### 7.3 哈希冲突问题

**问题：**
使用数字哈希时可能发生哈希冲突，导致误判。

**解决方案：**
1. 使用双哈希：计算两个不同的哈希值，只有两个都匹配才认为相同
2. 使用大质数作为模数
3. 字符串哈希相对不容易冲突，但效率较低

### 7.4 负数节点值处理

**问题：**
节点值可能为负数，直接用于哈希计算可能导致问题。

**解决方案：**
```cpp
// 将负数转为正数
long long val = node->val + 10000;  // 因为 val >= -10^4
```

---

## 八、与其他相关题目的对比分析

### 8.1 与 100. 相同的树对比

| 题目 | 核心问题 | 解法 |
|------|---------|------|
| 100. 相同的树 | 判断两棵树是否完全相同 | 递归/迭代比较 |
| 572. 另一棵树的子树 | 判断一棵树是否是另一棵树的子树 | 在主树中遍历，对每个节点调用相同树判断 |

**关系：**
- 572 题可以看作是 100 题的扩展
- 572 题需要多次调用 100 题的解法

### 8.2 与 101. 对称二叉树对比

| 题目 | 核心问题 | 比较方式 |
|------|---------|---------|
| 101. 对称二叉树 | 判断一棵树是否对称 | 比较左子树的左节点与右子树的右节点 |
| 572. 另一棵树的子树 | 判断一棵树是否是另一棵树的子树 | 比较两棵树是否完全相同 |

**关系：**
- 两道题都涉及树的比较
- 101 题是内部比较，572 题是外部比较

### 8.3 与 543. 二叉树的直径对比

| 题目 | 核心问题 | 解法 |
|------|---------|------|
| 543. 二叉树的直径 | 求二叉树的最长路径 | 后序遍历计算深度 |
| 572. 另一棵树的子树 | 判断子树关系 | 遍历 + 树比较 |

**关系：**
- 两道题都涉及树的遍历
- 树哈希优化方法使用后序遍历，与 543 题有相似之处

---

## 九、面试高频问题与解答

### 9.1 基础问题

#### Q1: 如何判断一棵树是否是另一棵树的子树？

**答案：**

**方法：**
1. 遍历主树的每个节点
2. 对于每个节点，检查以该节点为根的子树是否与目标子树相同
3. 如果找到匹配的，返回 `true`；否则返回 `false`

**核心代码：**
```cpp
bool isSubtree(TreeNode* root, TreeNode* subRoot) {
    if (root == nullptr) return false;
    if (isSameTree(root, subRoot)) return true;
    return isSubtree(root->left, subRoot) || isSubtree(root->right, subRoot);
}
```

#### Q2: 如何判断两棵树是否完全相同？

**答案：**

**递归方法：**
1. 如果两棵树都为空，返回 `true`
2. 如果其中一棵为空，另一棵不为空，返回 `false`
3. 如果节点值不同，返回 `false`
4. 递归比较左子树和右子树

```cpp
bool isSameTree(TreeNode* p, TreeNode* q) {
    if (p == nullptr && q == nullptr) return true;
    if (p == nullptr || q == nullptr) return false;
    if (p->val != q->val) return false;
    return isSameTree(p->left, q->left) && isSameTree(p->right, q->right);
}
```

### 9.2 进阶问题

#### Q3: 如何优化子树判断的时间复杂度？

**答案：**

**方法：树哈希**

**思路：**
1. 对主树的每个子树计算哈希值
2. 对目标子树计算哈希值
3. 检查目标子树的哈希值是否存在于主树的哈希集合中

**时间复杂度：** O(n + m)

**代码：**
```cpp
string hashTree(TreeNode* node) {
    if (node == nullptr) return "#";
    string left = hashTree(node->left);
    string right = hashTree(node->right);
    return "(" + left + ")" + to_string(node->val) + "(" + right + ")";
}
```

#### Q4: 子树和子结构有什么区别？

**答案：**

**子树：**
- 必须包含某个节点及其所有后代节点
- 结构和值必须完全匹配

**子结构：**
- 只需要匹配部分结构
- 不需要包含所有后代节点

**示例：**
```
root:      3
          / \
         4   5
        / \
       1   2
          /
         0

subRoot:   4
          / \
         1   2
```

- `subRoot` **不是** `root` 的子树
- `subRoot` **是** `root` 的子结构

### 9.3 实际问题

#### Q5: 在实际项目中，你会选择哪种方法？

**答案：**

**根据场景选择：**

| 场景 | 推荐方法 | 原因 |
|------|---------|------|
| 树较小 | 递归法 | 实现简单，代码清晰 |
| 树较大 | 树哈希 | 时间复杂度更低 |
| 需要避免递归栈溢出 | 迭代法 | 适用于深度很大的树 |

**实际建议：**
- 对于大多数面试场景，递归法足够
- 如果面试官要求优化，可以提到树哈希方法

#### Q6: 如何处理哈希冲突？

**答案：**

**方法：**
1. **双哈希**：使用两个不同的哈希函数，只有两个哈希值都匹配才认为相同
2. **大质数模数**：使用较大的质数作为模数，减少冲突概率
3. **字符串哈希**：字符串哈希的冲突概率较低
4. **二次检查**：如果哈希值匹配，再进行一次完整的树比较确认

---

## 十、完整可运行代码

```cpp
#include <iostream>
#include <queue>
#include <unordered_set>
#include <string>
using namespace std;

// ==========================================
// 二叉树节点定义
// ==========================================
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
class RecursiveSolution {
public:
    bool isSubtree(TreeNode* root, TreeNode* subRoot) {
        if (root == nullptr) {
            return false;
        }
        
        if (isSameTree(root, subRoot)) {
            return true;
        }
        
        return isSubtree(root->left, subRoot) || isSubtree(root->right, subRoot);
    }

private:
    bool isSameTree(TreeNode* p, TreeNode* q) {
        if (p == nullptr && q == nullptr) {
            return true;
        }
        
        if (p == nullptr || q == nullptr) {
            return false;
        }
        
        if (p->val != q->val) {
            return false;
        }
        
        return isSameTree(p->left, q->left) && isSameTree(p->right, q->right);
    }
};

// ==========================================
// 方法二：迭代法
// ==========================================
class IterativeSolution {
public:
    bool isSubtree(TreeNode* root, TreeNode* subRoot) {
        if (root == nullptr) {
            return false;
        }
        
        queue<TreeNode*> q;
        q.push(root);
        
        while (!q.empty()) {
            TreeNode* current = q.front();
            q.pop();
            
            if (isSameTree(current, subRoot)) {
                return true;
            }
            
            if (current->left != nullptr) {
                q.push(current->left);
            }
            
            if (current->right != nullptr) {
                q.push(current->right);
            }
        }
        
        return false;
    }

private:
    bool isSameTree(TreeNode* p, TreeNode* q) {
        if (p == nullptr && q == nullptr) {
            return true;
        }
        
        if (p == nullptr || q == nullptr) {
            return false;
        }
        
        queue<TreeNode*> q1, q2;
        q1.push(p);
        q2.push(q);
        
        while (!q1.empty() && !q2.empty()) {
            TreeNode* node1 = q1.front();
            TreeNode* node2 = q2.front();
            q1.pop();
            q2.pop();
            
            if (node1->val != node2->val) {
                return false;
            }
            
            bool left1 = (node1->left != nullptr);
            bool left2 = (node2->left != nullptr);
            
            if (left1 != left2) {
                return false;
            }
            if (left1) {
                q1.push(node1->left);
                q2.push(node2->left);
            }
            
            bool right1 = (node1->right != nullptr);
            bool right2 = (node2->right != nullptr);
            
            if (right1 != right2) {
                return false;
            }
            if (right1) {
                q1.push(node1->right);
                q2.push(node2->right);
            }
        }
        
        return q1.empty() && q2.empty();
    }
};

// ==========================================
// 方法三：树哈希优化（字符串哈希）
// ==========================================
class HashSolution {
public:
    bool isSubtree(TreeNode* root, TreeNode* subRoot) {
        string subHash = hashTree(subRoot);
        
        unordered_set<string> hashSet;
        hashTreeWithSet(root, hashSet);
        
        return hashSet.count(subHash);
    }

private:
    string hashTree(TreeNode* node) {
        if (node == nullptr) {
            return "#";
        }
        
        string left = hashTree(node->left);
        string right = hashTree(node->right);
        
        return "(" + left + ")" + to_string(node->val) + "(" + right + ")";
    }
    
    string hashTreeWithSet(TreeNode* node, unordered_set<string>& hashSet) {
        if (node == nullptr) {
            return "#";
        }
        
        string left = hashTreeWithSet(node->left, hashSet);
        string right = hashTreeWithSet(node->right, hashSet);
        
        string hash = "(" + left + ")" + to_string(node->val) + "(" + right + ")";
        hashSet.insert(hash);
        
        return hash;
    }
};

// ==========================================
// 辅助函数：创建二叉树
// ==========================================
TreeNode* createTree(const int* arr, int size, int index = 0) {
    if (index >= size || arr[index] == -1) {
        return nullptr;
    }
    
    TreeNode* root = new TreeNode(arr[index]);
    root->left = createTree(arr, size, 2 * index + 1);
    root->right = createTree(arr, size, 2 * index + 2);
    
    return root;
}

// ==========================================
// 辅助函数：打印二叉树
// ==========================================
void printTree(TreeNode* root, int depth = 0) {
    if (root == nullptr) {
        return;
    }
    
    printTree(root->right, depth + 1);
    
    for (int i = 0; i < depth; ++i) {
        cout << "    ";
    }
    cout << root->val << endl;
    
    printTree(root->left, depth + 1);
}

// ==========================================
// 主函数测试
// ==========================================
int main() {
    cout << "========================================" << endl;
    cout << "        572. 另一棵树的子树测试" << endl;
    cout << "========================================" << endl;
    
    // ==========================================
    // 测试用例 1
    // ==========================================
    cout << "\n--- 测试用例 1 ---" << endl;
    int root1[] = {3, 4, 5, 1, 2};
    int subRoot1[] = {4, 1, 2};
    
    TreeNode* tree1 = createTree(root1, 5);
    TreeNode* subTree1 = createTree(subRoot1, 3);
    
    cout << "root 树：" << endl;
    printTree(tree1);
    
    cout << "\nsubRoot 树：" << endl;
    printTree(subTree1);
    
    RecursiveSolution sol1;
    bool result1 = sol1.isSubtree(tree1, subTree1);
    cout << "\n结果：" << (result1 ? "true" : "false") << endl;
    cout << "期望：true" << endl;
    
    // ==========================================
    // 测试用例 2
    // ==========================================
    cout << "\n--- 测试用例 2 ---" << endl;
    int root2[] = {3, 4, 5, 1, 2, -1, -1, -1, -1, 0};
    int subRoot2[] = {4, 1, 2};
    
    TreeNode* tree2 = createTree(root2, 10);
    TreeNode* subTree2 = createTree(subRoot2, 3);
    
    cout << "root 树：" << endl;
    printTree(tree2);
    
    cout << "\nsubRoot 树：" << endl;
    printTree(subTree2);
    
    bool result2 = sol1.isSubtree(tree2, subTree2);
    cout << "\n结果：" << (result2 ? "true" : "false") << endl;
    cout << "期望：false" << endl;
    
    // ==========================================
    // 测试用例 3
    // ==========================================
    cout << "\n--- 测试用例 3 ---" << endl;
    int root3[] = {1};
    int subRoot3[] = {1};
    
    TreeNode* tree3 = createTree(root3, 1);
    TreeNode* subTree3 = createTree(subRoot3, 1);
    
    bool result3 = sol1.isSubtree(tree3, subTree3);
    cout << "结果：" << (result3 ? "true" : "false") << endl;
    cout << "期望：true" << endl;
    
    // ==========================================
    // 测试用例 4
    // ==========================================
    cout << "\n--- 测试用例 4 ---" << endl;
    int root4[] = {1, 2};
    int subRoot4[] = {1};
    
    TreeNode* tree4 = createTree(root4, 2);
    TreeNode* subTree4 = createTree(subRoot4, 1);
    
    bool result4 = sol1.isSubtree(tree4, subTree4);
    cout << "结果：" << (result4 ? "true" : "false") << endl;
    cout << "期望：false" << endl;
    
    cout << "\n========================================" << endl;
    cout << "            测试完成！" << endl;
    cout << "========================================" << endl;
    
    return 0;
}
```

---

## 十一、总结与反思

### 11.1 核心知识点总结

| 知识点 | 要点 |
|--------|------|
| 子树判断 | 在主树中查找与目标子树完全相同的子树 |
| 递归法 | 遍历主树，对每个节点调用相同树判断 |
| 迭代法 | 使用队列进行层序遍历 |
| 树哈希 | 将时间复杂度优化到 O(n + m) |
| 子树 vs 子结构 | 子树必须包含所有后代节点 |

### 11.2 学习收获

1. **理解了子树判断的核心逻辑**
   - 需要遍历主树的每个节点
   - 对每个节点进行树比较

2. **掌握了三种不同的解法**
   - 递归法：实现简单，代码清晰
   - 迭代法：避免递归栈溢出
   - 树哈希：时间复杂度最优

3. **理解了哈希优化的原理**
   - 使用后序遍历计算子树哈希
   - 通过哈希集合快速查找

4. **区分了子树和子结构**
   - 子树必须完全匹配
   - 子结构只需部分匹配

### 11.3 后续学习建议

1. **学习更多树的比较算法**
   - 树的同构判断
   - 树的相似度计算

2. **深入学习哈希算法**
   - 字符串哈希
   - 数字哈希
   - 哈希冲突处理

3. **学习其他树问题**
   - 树的直径
   - 树的最大路径和
   - 树的序列化与反序列化

4. **练习更多相关题目**
   - LeetCode 100. 相同的树
   - LeetCode 101. 对称二叉树
   - LeetCode 543. 二叉树的直径

---

通过本课程，我们深入理解了子树判断的核心逻辑，掌握了递归法、迭代法和树哈希优化三种解法，也学习了子树与子结构的区别。这些知识对于理解树的操作和算法优化非常重要！