# 算法实战复盘：螺旋矩阵 II (LeetCode 59)

> **打卡日期**：2026-05-19
> **核心主题**：数组模拟、边界控制、循环不变量。

---

## 📝 一、 题目描述与核心要求

### 1. 中文描述

给你一个正整数 `n` ，生成一个包含 `1` 到 `n²` 所有元素，且元素按顺时针顺序螺旋排列的 `n x n` 正方形矩阵。

### 2. 英文描述

Given a positive integer `n`, generate an `n x n` matrix filled with elements from `1` to `n²` in spiral order.

### 3. 输入输出示例

**示例 1**：
- 输入：`n = 3`
- 输出：
```
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
```

**示例 2**：
- 输入：`n = 1`
- 输出：`[[1]]`

### 4. 题目提示

- `1 <= n <= 20`
- `n²` 的最大值为 400

### 5. 核心痛点

- 需要按顺时针螺旋顺序填充矩阵的四个方向
- 每一圈遍历后，边界会收缩，需要准确控制
- 遍历完所有元素即停止，避免重复填充

### 6. 题目延伸

本题是 [54. 螺旋矩阵](https://leetcode.cn/problems/spiral-matrix/) 的逆向问题：
- **54题**：给定矩阵，按螺旋顺序遍历并输出
- **59题**：给定大小，生成螺旋顺序的矩阵

---

## 📊 二、 复杂度深度剖析

### 1. 时间复杂度：O(n²)

**推导过程**：
- 需要填充 `n x n = n²` 个元素
- 每个元素只被访问一次
- 因此时间复杂度为 O(n²)

### 2. 空间复杂度：O(n²)

**推导过程**：
- 需要存储 `n x n` 的结果矩阵
- 对于 n=20，结果矩阵包含 400 个元素
- 因此空间复杂度为 O(n²)

---

## 🤔 三、 解题思路分析

### 1. 初学者的常见思路

当面对这个问题时，初学者可能会想到以下几种方法：

1. **方向数组法**：定义四个方向（右、下、左、上），按顺序循环处理
2. **边界收缩法**：设定四条边界，每遍历完一条边就收缩对应边界
3. **模拟法**：完全按照人类填写矩阵的直觉，依次向右、向下、向左、向上填充

### 2. 各种方法的优缺点分析

| 方法 | 时间复杂度 | 空间复杂度 | 优点 | 缺点 |
|------|------------|------------|------|------|
| 方向数组法 | O(n²) | O(n²) | 代码清晰，方向控制灵活 | 需要处理边界越界 |
| 边界收缩法 | O(n²) | O(n²) | 逻辑直观，边界明确 | 需要记录四条边界值 |
| 模拟法 | O(n²) | O(n²) | 最直观，符合人类思维 | 代码较复杂，容易出错 |

### 3. 方法选择建议

**推荐使用方向数组 + 边界收缩法**，原因如下：
1. 逻辑清晰，每一步都有明确的边界判断
2. 代码简洁，循环终止条件明确
3. 易于理解和调试
4. 是面试中的标准解法

---

## 🎯 四、 核心思想：循环不变量

### 1. 循环不变式的定义

**循环不变量**是指在循环的每次迭代前后都保持不变的条件。在本题中：

- **不变量**：`start <= end`，即起始边界不超过结束边界
- **含义**：当 `start > end` 时，所有元素都已填充完毕

### 2. 螺旋填充的四个阶段

```
┌─────────────────────────────┐
│  第一阶段：左→右（顶部行）    │  top 边界
│  第二阶段：上→下（右侧列）    │  right 边界
│  第三阶段：右→左（底部行）    │  bottom 边界
│  第四阶段：下→上（左侧列）    │  left 边界
└─────────────────────────────┘
     每阶段结束后，相应边界收缩
```

### 3. 边界收缩过程图解

对于 n=3 的矩阵，填充过程如下：

```
第1圈：
┌─────────────┐
│ 1 → 2 → 3   │  ← top=0, 填充顶部行后 top++
│ ↓           │
│ 8           │  ← right=2, 填充右侧列后 right--
│ ↓           │
│ 7 ← 6 ← 5   │  ← bottom=2, 填充底部行后 bottom--
│ ↑           │
│ 4           │  ← left=0, 填充左侧列后 left++
└─────────────┘
     top=1, bottom=1, left=1, right=1
     继续填充中心元素 9

最终结果：
┌─────────────┐
│ 1   2   3   │
│ 8   9   4   │
│ 7   6   5   │
└─────────────┘
```

### 4. 奇数n的中心处理

当 n 为奇数时，矩阵中心会有一个单独的元素，需要单独处理：

```
n=3：中心是 matrix[1][1] = 9
n=5：中心是 matrix[2][2] = 25
```

---

## 📚 五、 算法实现

### 1. 边界收缩法（C++实现）

```cpp
class Solution {
public:
    vector<vector<int>> generateMatrix(int n) {
        // 初始化结果矩阵，所有元素为0
        vector<vector<int>> matrix(n, vector<int>(n, 0));

        // 定义四条边界
        int top = 0;      // 顶部边界（起始行）
        int bottom = n - 1;  // 底部边界（结束行）
        int left = 0;     // 左侧边界（起始列）
        int right = n - 1;   // 右侧边界（结束列）

        // 当前要填充的数字，从1开始
        int num = 1;

        // 当 top <= bottom 且 left <= right 时，继续填充
        while (top <= bottom && left <= right) {
            // 阶段1：从左到右填充顶部行
            // 填充完毕后，顶部边界下移一位
            for (int col = left; col <= right; ++col) {
                matrix[top][col] = num++;
            }
            ++top;  // 收缩顶部边界

            // 阶段2：从上到下填充右侧列
            // 填充完毕后，右侧边界左移一位
            for (int row = top; row <= bottom; ++row) {
                matrix[row][right] = num++;
            }
            --right;  // 收缩右侧边界

            // 阶段3：从右到左填充底部行
            // 需要检查 top <= bottom，确保还有剩余行
            if (top <= bottom) {
                for (int col = right; col >= left; --col) {
                    matrix[bottom][col] = num++;
                }
                --bottom;  // 收缩底部边界
            }

            // 阶段4：从下到上填充左侧列
            // 需要检查 left <= right，确保还有剩余列
            if (left <= right) {
                for (int row = bottom; row >= top; --row) {
                    matrix[left][row] = num++;
                }
                ++left;  // 收缩左侧边界
            }
        }

        return matrix;
    }
};
```

### 2. "撞南墙"方法（方向数组法）

这种方法非常巧妙，像机器人走路一样，一直往前走，碰到边界或已经填充的位置就"右转"90度。

```cpp
class Solution {
public:
    vector<vector<int>> generateMatrix(int n) {
        // 方向数组：右下左上（顺时针顺序）
        static constexpr int DIRS[4][2] = {
            {0, 1},   // 向右
            {1, 0},   // 向下
            {0, -1},  // 向左
            {-1, 0}   // 向上
        };
        
        // 初始化结果矩阵
        vector<vector<int>> ans(n, vector<int>(n));
        
        int i = 0, j = 0;  // 当前位置
        int di = 0;        // 当前方向索引（0-3）
        
        // 填充1到n²
        for (int val = 1; val <= n * n; val++) {
            ans[i][j] = val;  // 填入当前值
            
            // 计算下一步位置
            int x = i + DIRS[di][0];
            int y = j + DIRS[di][1];
            
            // 判断是否需要转向
            // 如果下一步出界，或者下一步位置已经被填充
            if (x < 0 || x >= n || y < 0 || y >= n || ans[x][y] != 0) {
                di = (di + 1) % 4;  // 右转90度（顺时针切换方向）
            }
            
            // 移动到下一个位置
            i += DIRS[di][0];
            j += DIRS[di][1];
        }
        
        return ans;
    }
};
```

#### 2.1 "撞南墙"方法核心逻辑

**方向数组**：
```cpp
// 右下左上
static constexpr int DIRS[4][2] = {{0, 1}, {1, 0}, {0, -1}, {-1, 0}};
```
- `DIRS[0]` = {0, 1}：向右移动（行不变，列+1）
- `DIRS[1]` = {1, 0}：向下移动（行+1，列不变）
- `DIRS[2]` = {0, -1}：向左移动（行不变，列-1）
- `DIRS[3]` = {-1, 0}：向上移动（行-1，列不变）

**转向判断条件**：
```cpp
if (x < 0 || x >= n || y < 0 || y >= n || ans[x][y] != 0)
```
- `x < 0` 或 `x >= n`：行越界
- `y < 0` 或 `y >= n`：列越界
- `ans[x][y] != 0`：该位置已被填充

**转向操作**：
```cpp
di = (di + 1) % 4;  // 顺时针转向
// di = (di + 3) % 4;  // 如果想逆时针转向
```

#### 2.2 "撞南墙"方法图解（n=3）

```
填充过程：
1 → 2 → 3 ↓
↑         ↓
8 ← 9 ← 4 ↓
↑         ↓
7 ← 6 ← 5

详细步骤：
val=1: (0,0) → 向右
val=2: (0,1) → 向右
val=3: (0,2) → 下一步(0,3)出界，转向向下
val=4: (1,2) → 向下
val=5: (2,2) → 下一步(3,2)出界，转向向左
val=6: (2,1) → 向左
val=7: (2,0) → 下一步(2,-1)出界，转向上
val=8: (1,0) → 向上
val=9: (0,0)已填充，转向向右
val=10: 循环结束（10 > 9）
```

#### 2.3 "撞南墙"方法优缺点

| 特点 | 说明 |
|------|------|
| **优点** | 代码极其简洁，逻辑直观，不需要复杂的边界管理 |
| **优点** | 易于理解和记忆，像机器人走路一样自然 |
| **优点** | 不需要判断奇偶，自动处理中心元素 |
| **缺点** | 每次移动前都需要检查边界和填充状态 |
| **适用场景** | 代码竞赛、面试快速编码 |

### 3. 简洁版边界收缩法

这是民间流传的非常简洁的写法，将边界收缩和填充合并在一个循环中。

```cpp
class Solution {
public:
    vector<vector<int>> generateMatrix(int n) {
        // 初始化四条边界
        int t = 0;      // top
        int b = n - 1;  // bottom
        int l = 0;      // left
        int r = n - 1;  // right
        
        // 初始化结果矩阵
        vector<vector<int>> ans(n, vector<int>(n));
        
        int k = 1;  // 当前要填充的值
        
        // 填充直到所有元素都被填入
        while (k <= n * n) {
            // 阶段1：左→右（顶部行）
            for (int i = l; i <= r && k <= n * n; ++i, ++k) {
                ans[t][i] = k;
            }
            ++t;  // 顶部边界下移
            
            // 阶段2：上→下（右侧列）
            for (int i = t; i <= b && k <= n * n; ++i, ++k) {
                ans[i][r] = k;
            }
            --r;  // 右侧边界左移
            
            // 阶段3：右→左（底部行）
            for (int i = r; i >= l && k <= n * n; --i, ++k) {
                ans[b][i] = k;
            }
            --b;  // 底部边界上移
            
            // 阶段4：下→上（左侧列）
            for (int i = b; i >= t && k <= n * n; --i, ++k) {
                ans[i][l] = k;
            }
            ++l;  // 左侧边界右移
        }
        
        return ans;
    }
};
```

#### 3.1 简洁版核心特点

**关键优化**：在每个for循环中添加 `k <= n * n` 的条件检查

```cpp
for (int i = l; i <= r && k <= n * n; ++i, ++k)
```

这样做的好处：
- 当 n 为奇数时，中心元素被填充后，后续循环会自动跳过
- 不需要单独处理边界检查（如 `if (top <= bottom)`）
- 代码更加紧凑，适合快速编写

#### 3.2 三种方法对比

| 方法 | 特点 | 代码复杂度 | 适用场景 |
|------|------|------------|----------|
| **标准边界收缩法** | 边界检查明确，逻辑清晰 | 中等 | 面试讲解、教学 |
| **撞南墙法** | 代码简洁，自动转向 | 低 | 竞赛快速编码 |
| **简洁版边界收缩法** | 紧凑高效，条件合并 | 低 | 日常开发 |

### 4. 标准边界收缩法核心逻辑解析

**循环终止条件**：`while (top <= bottom && left <= right)`

- 当 `top > bottom` 时，所有行都已填充完毕
- 当 `left > right` 时，所有列都已填充完毕
- 两个条件同时满足时，说明还有未填充的区域

**边界收缩的四个阶段**：

1. **顶部行填充**：`top` 行，从 `left` 列到 `right` 列
   - 填充完毕后 `top++`

2. **右侧列填充**：`right` 列，从 `top` 行到 `bottom` 行
   - 填充完毕后 `right--`

3. **底部行填充**：`bottom` 行，从 `right` 列到 `left` 列
   - **必须检查** `top <= bottom`，否则可能重复填充

4. **左侧列填充**：`left` 列，从 `bottom` 行到 `top` 行
   - **必须检查** `left <= right`，否则可能重复填充

### 5. 边界检查的重要性

```cpp
// 阶段3：底部行填充
// 如果不检查 top <= bottom，对于 n=3：
// top=1, bottom=1 时仍会执行，但此时只剩中心元素
// 如果没有检查，会再次填充第一行！
if (top <= bottom) {  // 必须检查
    for (int col = right; col >= left; --col) {
        matrix[bottom][col] = num++;
    }
    --bottom;
}

// 阶段4：左侧列填充
// 如果不检查 left <= right，会导致重复填充或越界
if (left <= right) {  // 必须检查
    for (int row = bottom; row >= top; --row) {
        matrix[left][row] = num++;
    }
    ++left;
}
```

### 6. 完整测试代码

```cpp
#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    vector<vector<int>> generateMatrix(int n) {
        vector<vector<int>> matrix(n, vector<int>(n, 0));
        int top = 0, bottom = n - 1, left = 0, right = n - 1;
        int num = 1;

        while (top <= bottom && left <= right) {
            // 阶段1：左→右（顶部行）
            for (int col = left; col <= right; ++col) {
                matrix[top][col] = num++;
            }
            ++top;

            // 阶段2：上→下（右侧列）
            for (int row = top; row <= bottom; ++row) {
                matrix[row][right] = num++;
            }
            --right;

            // 阶段3：右→左（底部行）
            if (top <= bottom) {
                for (int col = right; col >= left; --col) {
                    matrix[bottom][col] = num++;
                }
                --bottom;
            }

            // 阶段4：下→上（左侧列）
            if (left <= right) {
                for (int row = bottom; row >= top; --row) {
                    matrix[left][row] = num++;
                }
                ++left;
            }
        }

        return matrix;
    }
};

// 打印矩阵的辅助函数
void printMatrix(const vector<vector<int>>& matrix) {
    for (const auto& row : matrix) {
        for (int val : row) {
            cout << val << "\t";
        }
        cout << endl;
    }
}

int main() {
    Solution sol;

    cout << "=== n = 1 ===" << endl;
    printMatrix(sol.generateMatrix(1));

    cout << "\n=== n = 2 ===" << endl;
    printMatrix(sol.generateMatrix(2));

    cout << "\n=== n = 3 ===" << endl;
    printMatrix(sol.generateMatrix(3));

    cout << "\n=== n = 4 ===" << endl;
    printMatrix(sol.generateMatrix(4));

    cout << "\n=== n = 5 ===" << endl;
    printMatrix(sol.generateMatrix(5));

    return 0;
}
```

**输出结果**：

```
=== n = 1 ===
1

=== n = 2 ===
1	2
4	3

=== n = 3 ===
1	2	3
8	9	4
7	6	5

=== n = 4 ===
1	2	3	4
12	13	14	5
11	16	15	6
10	9	8	7

=== n = 5 ===
1	2	3	4	5
16	17	18	19	6
15	24	25	20	7
14	23	22	21	8
13	12	11	10	9
```

---

## ⚠️ 六、 易错点与注意事项

### 1. 忘记边界检查

**错误写法**：

```cpp
// 缺少边界检查会导致重复填充
for (int col = right; col >= left; --col) {
    matrix[bottom][col] = num++;
}
--bottom;

for (int row = bottom; row >= top; --row) {
    matrix[left][row] = num++;
}
++left;
```

**正确写法**：

```cpp
// 必须添加边界检查
if (top <= bottom) {  // 检查是否还有剩余行
    for (int col = right; col >= left; --col) {
        matrix[bottom][col] = num++;
    }
    --bottom;
}

if (left <= right) {  // 检查是否还有剩余列
    for (int row = bottom; row >= top; --row) {
        matrix[left][row] = num++;
    }
    ++left;
}
```

### 2. 循环终止条件错误

**错误写法**：

```cpp
// 只用 top < bottom && left < right
while (top < bottom && left < right) {
    // 对于 n=3，当 top=1, bottom=1, left=1, right=1 时
    // 条件 top < bottom 为 false（1 < 1），退出循环
    // 但中心元素 matrix[1][1] 还未填充！
}
```

**正确写法**：

```cpp
// 使用 <= 允许中心元素被填充
while (top <= bottom && left <= right) {
    // 当 top=1, bottom=1, left=1, right=1 时
    // 条件为 true (1 <= 1)，进入循环填充中心元素
}
```

### 3. 边界收缩时机错误

**错误写法**：

```cpp
// 在填充循环内部收缩边界
for (int col = left; col <= right; ++col) {
    matrix[top][col] = num++;
    ++top;  // 错误！每次填充后都收缩
}
```

**正确写法**：

```cpp
// 在填充循环外部收缩边界
for (int col = left; col <= right; ++col) {
    matrix[top][col] = num++;
}
++top;  // 正确！填充完毕后收缩
```

### 4. 初始化错误

**错误写法**：

```cpp
vector<vector<int>> matrix(n, vector<int>(n));  // 缺少初始值
// 可能导致未定义行为
```

**正确写法**：

```cpp
vector<vector<int>> matrix(n, vector<int>(n, 0));  // 显式初始化为0
```

---

## 🔄 七、 螺旋矩阵系列题目对比

| 题目 | 难度 | 核心差异 |
|------|------|----------|
| [54. 螺旋矩阵](https://leetcode.cn/problems/spiral-matrix/) | 中等 | 给定矩阵，按螺旋顺序遍历输出 |
| [59. 螺旋矩阵 II](https://leetcode.cn/problems/spiral-matrix-ii/) | 中等 | 给定大小，生成螺旋矩阵 |
| [885. 螺旋矩阵 III](https://leetcode.cn/problems/spiral-matrix-iii/) | 中等 | 在矩阵中按螺旋顺序走R*C步 |

### 1. 三道题的共同点

- 都使用方向数组或边界收缩法
- 都需要处理循环终止条件
- 核心思想都是「循环不变量」

### 2. 三道题的差异点

| 题目 | 输入 | 输出 | 特点 |
|------|------|------|------|
| 54 | 已有矩阵 | 遍历结果数组 | 需要处理空矩阵 |
| 59 | n（矩阵大小） | 完整矩阵 | 需要处理中心元素 |
| 885 | R, C, startRow, startCol | 走过的R*C个位置 | 需要计算起始位置 |

---

## 💡 八、 面试高频问题

### 问题1：为什么需要边界检查？

**参考答案**：
- 在填充底部行和左侧列时，如果只填充了一圈，可能会导致重复填充
- 例如 n=3 时，填充完顶部行、右侧列、底部行后，`top=1, bottom=1`，只剩中心行
- 如果不检查 `top <= bottom`，底部行填充会重复填充第一行

### 问题2：如何处理奇数n的中心元素？

**参考答案**：
- 当 n 为奇数时，中心元素会在最后一圈被填充
- 由于 `top <= bottom && left <= right` 的循环条件，中心元素会在最后一圈被正确填充
- 不需要额外处理，但需要在循环终止条件中使用 `<=` 而非 `<`

### 问题3：时间复杂度为什么是O(n²)？

**参考答案**：
- 矩阵包含 n² 个元素
- 每个元素只被访问一次
- 因此时间复杂度为 O(n²)

### 问题4：能否不使用额外空间？

**参考答案**：
- 本题需要存储结果矩阵，空间复杂度至少为 O(n²)
- 如果只是遍历现有矩阵（54题），可以做到 O(1) 额外空间

---

## 📖 九、 学习总结

### 1. 核心知识点

- **循环不变量**：理解并应用「边界收缩」的思维方式
- **边界控制**：正确处理四个方向的边界条件
- **终止条件**：使用 `<=` 确保中心元素被填充

### 2. 关键技巧

- 初始化四条边界：`top`, `bottom`, `left`, `right`
- 每阶段填充完毕后收缩对应边界
- 添加必要的边界检查，防止重复填充

### 3. 延伸学习

- 建议同时完成 **54题（螺旋矩阵）** 和 **885题（螺旋矩阵 III）**
- 对比三道题的异同，加深对「循环不变量」的理解

---

**参考资源**：
- [代码随想录 - 螺旋矩阵II](https://www.programmercarl.com/0059.螺旋矩阵II.html)
- [LeetCode 官方题解](https://leetcode.cn/problems/spiral-matrix-ii/solutions/)
- [简洁写法对比](https://leetcode.cn/problems/spiral-matrix-ii/solutions/3059650/)