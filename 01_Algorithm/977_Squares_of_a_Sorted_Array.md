# 977. 有序数组的平方学习笔记

## 一、 题目分析

### 1.1 题目描述

给你一个按 **非递减顺序** 排序的整数数组 `nums`，返回每个数字的平方组成的新数组，要求也按 **非递减顺序** 排序。

### 1.2 题目链接

- [LeetCode 977. 有序数组的平方](https://leetcode.cn/problems/squares-of-a-sorted-array/description/)
- [代码随想录题解](https://www.programmercarl.com/0977.%E6%9C%89%E5%BA%8F%E6%95%B0%E7%BB%84%E7%9A%84%E5%B9%B3%E6%96%B9.html)
- [官方题解](https://leetcode.cn/problems/squares-of-a-sorted-array/solutions/447736/you-xu-shu-zu-de-ping-fang-by-leetcode-solution/)
- [双指针详解](https://leetcode.cn/problems/squares-of-a-sorted-array/solutions/2806253/xiang-xiang-shuang-zhi-zhen-cong-da-dao-blda6/)
- [多种排序双指针](https://leetcode.cn/problems/squares-of-a-sorted-array/solutions/39242/ge-chong-pai-xu-shuang-zhi-zhen-by-toxic-3/)

### 1.3 示例

**示例 1：**

```
输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]。
排序后，数组变为 [0,1,9,16,100]。
```

**示例 2：**

```
输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]
```

### 1.4 提示

- `1 <= nums.length <= 10^4`
- `-10^4 <= nums[i] <= 10^4`
- `nums` 已按 **非递减顺序** 排序

### 1.5 题目本质

有序数组的平方的本质是：**数组本身已排序，但负数的平方可能会打破这个顺序，需要重新排序**。

关键观察：
- 数组是升序排列的
- 负数的平方是降序的（绝对值越大，平方越大）
- 正数的平方是升序的（值越大，平方越大）
- 最大平方只可能来自数组的两端（最小负数或最大正数）

### 1.6 解题关键点

| 关键点 | 说明 |
|-------|------|
| 数组已排序 | 原数组是升序的 |
| 平方特性 | 负数平方后绝对值越大平方越大 |
| 最大值位置 | 平方后的最大值只可能来自数组两端 |
| 双指针技巧 | 从两端向中间收缩，选择较大的平方 |

---

## 二、 解题思路分析

### 2.1 核心思路

解决有序数组平方问题的核心思路是：**利用数组已排序的特性，使用双指针从两端向中间遍历，选择较大的平方放入结果数组的末尾**。

### 2.2 暴力解法

#### 2.2.1 暴力解法思路

最直接的方法：
1. 遍历数组，将每个元素平方
2. 对平方后的数组进行排序
3. 返回排序后的结果

#### 2.2.2 暴力解法流程图

```
暴力解法流程：
┌─────────────────────────────────────────────────────┐
│  1. 遍历数组：                                        │
│     for num in nums:                                 │
│         squares.push_back(num * num)                 │
│                                                         │
│  2. 对平方数组排序：                                   │
│     sort(squares.begin(), squares.end())             │
│                                                         │
│  3. 返回结果                                           │
└─────────────────────────────────────────────────────┘
```

### 2.3 双指针解法

#### 2.3.1 双指针核心逻辑

利用数组已排序的特性：
1. 数组两端分别是**最小负数**和**最大正数**
2. 两端元素的平方一定是数组中最大的两个
3. 从两端向中间收缩，依次选择较大的平方
4. 结果数组从**末尾开始填充**（因为我们从大到小选择）

#### 2.3.2 双指针流程图

```
双指针流程：
┌─────────────────────────────────────────────────────┐
│  left = 0, right = n - 1                            │
│  result = [0] * n                                   │
│  index = n - 1                                       │
├─────────────────────────────────────────────────────┤
│  while left <= right:                               │
│      leftSquare = nums[left] * nums[left]            │
│      rightSquare = nums[right] * nums[right]        │
│                                                         │
│      if leftSquare > rightSquare:                    │
│          result[index] = leftSquare                   │
│          left++                                      │
│      else:                                           │
│          result[index] = rightSquare                  │
│          right--                                     │
│      index--                                         │
├─────────────────────────────────────────────────────┤
│  return result                                       │
└─────────────────────────────────────────────────────┘
```

#### 2.3.3 双指针示例演示

以 `nums = [-4, -1, 0, 3, 10]` 为例：

```
初始状态：
数组:    [-4, -1, 0, 3, 10]
         ↑            ↑
       left          right
       (0)           (4)

结果数组: [0, 0, 0, 0, 0]
                      ↑
                    index = 4

第1轮：
- left=0指向-4, 平方=16
- right=4指向10, 平方=100
- 16 < 100，选择100（较大值），result[4]=100
- right-- → right=3
- index-- → index=3
结果: [0, 0, 0, 0, 100]

第2轮：
数组: [-4, -1, 0, 3, 10]
              ↑     ↑
            left  right
            (0)   (3)

- left=0指向-4, 平方=16
- right=3指向3, 平方=9
- 16 > 9，选择16，result[3]=16
- left++ → left=1
- index-- → index=2
结果: [0, 0, 0, 16, 100]

第3轮：
数组: [-4, -1, 0, 3, 10]
              ↑  ↑
            left right
            (1) (3)

- left=1指向-1, 平方=1
- right=3指向3, 平方=9
- 1 < 9，选择9，result[2]=9
- right-- → right=2
- index-- → index=1
结果: [0, 0, 9, 16, 100]

第4轮：
数组: [-4, -1, 0, 3, 10]
                 ↑
               left right
               (1) (2)

- left=1指向-1, 平方=1
- right=2指向0, 平方=0
- 1 > 0，选择1，result[1]=1
- left++ → left=2
- index-- → index=0
结果: [0, 1, 9, 16, 100]

第5轮：
数组: [-4, -1, 0, 3, 10]
                    ↑
                  left,right
                    (2)

- left=2指向0, 平方=0
- right=2指向0, 平方=0
- 0 == 0（相等，选择right），result[0]=0
- right-- → right=1
- index-- → index=-1
结果: [0, 1, 9, 16, 100]

最终结果: [0, 1, 9, 16, 100]  ✓
```

**关键点说明：**
- 每次比较左右两端平方的大小，选择**较大的**放入结果数组的**末尾**
- 较大的平方来自哪端，就移动哪端的指针
- 从大到小依次填满结果数组，最终得到升序结果

---

## 三、 算法实现

### 3.1 暴力解法

```cpp
class Solution {
public:
    vector<int> sortedSquares(vector<int>& nums) {
        vector<int> squares;
        
        // 步骤1：计算每个元素的平方
        for (int num : nums) {
            squares.push_back(num * num);
        }
        
        // 步骤2：对平方数组排序
        sort(squares.begin(), squares.end());
        
        return squares;
    }
};
```

### 3.2 双指针解法

```cpp
class Solution {
public:
    vector<int> sortedSquares(vector<int>& nums) {
        int n = nums.size();
        
        // 创建结果数组
        vector<int> result(n);
        
        // 左指针指向数组开头，右指针指向数组结尾
        int left = 0;
        int right = n - 1;
        
        // 结果数组的填充位置（从末尾开始）
        int index = n - 1;
        
        // 遍历直到左右指针相遇
        while (left <= right) {
            // 计算左右指针指向元素的平方
            int leftSquare = nums[left] * nums[left];
            int rightSquare = nums[right] * nums[right];
            
            // 选择较大的平方放入结果数组末尾
            if (leftSquare > rightSquare) {
                result[index] = leftSquare;
                left++;  // 左指针右移
            } else {
                result[index] = rightSquare;
                right--;  // 右指针左移
            }
            
            // 填充位置左移
            index--;
        }
        
        return result;
    }
};
```

### 3.3 双指针简洁版

```cpp
class Solution {
public:
    vector<int> sortedSquares(vector<int>& nums) {
        int n = nums.size();
        vector<int> result(n);
        
        int left = 0, right = n - 1;
        for (int i = n - 1; i >= 0; --i) {
            // 比较左右两端平方的大小
            if (abs(nums[left]) > abs(nums[right])) {
                result[i] = nums[left] * nums[left];
                left++;
            } else {
                result[i] = nums[right] * nums[right];
                right--;
            }
        }
        
        return result;
    }
};
```

---

## 四、 代码分析

### 4.1 暴力解法分析

**核心逻辑：**
1. **平方计算**：遍历数组，将每个元素平方
2. **排序**：使用标准库的排序算法

**优点：**
- 代码简单直观
- 容易理解和实现
- 无需考虑复杂的指针逻辑

**缺点：**
- 时间复杂度较高：平方 O(n) + 排序 O(nlogn)
- 没有利用数组已排序的特性

### 4.2 双指针解法分析

**核心逻辑：**
1. **指针初始化**：left 指向开头，right 指向结尾
2. **比较选择**：比较两端元素的平方，选择较大的
3. **逆序填充**：结果从末尾开始填充
4. **指针移动**：选择哪端，哪端指针向中间移动

**为什么从末尾填充？**
- 因为我们从大到小选择元素
- 最大的平方首先被确定
- 从末尾填充保证了结果的正确顺序

**为什么双指针有效？**
- 数组两端分别是绝对值最大和最小的元素
- 它们的平方一定是最大的两个
- 每次选择较大的后，较大的平方已经被确定
- 继续收缩，最终可以确定所有平方的顺序

### 4.3 两种方法对比

```
暴力解法：
数组: [-4, -1, 0, 3, 10]
平方: [16, 1, 0, 9, 100]
排序: [0, 1, 9, 16, 100]  ✓

双指针解法：
数组: [-4, -1, 0, 3, 10]
      ↑            ↑
    left          right

第1次: max(16, 100) = 100, 放最后
第2次: max(16, 9) = 16, 放倒数第二
第3次: max(16, 0) = 16, 放倒数第三
第4次: max(16, 1) = 16, 放倒数第四
第5次: 1, 放第一位

结果: [1, 16, 16, 16, 100]  ✓
```

---

## 五、 复杂度分析

### 5.1 暴力解法复杂度

| 复杂度类型 | 时间复杂度 | 空间复杂度 |
|-----------|-----------|-----------|
| **暴力解法** | $O(n \log n)$ | $O(n)$ |

**时间复杂度推导：**
- 计算平方：遍历数组，$O(n)$
- 排序：快速排序/归并排序等，$O(n \log n)$
- 总时间复杂度：$O(n \log n)$

**空间复杂度推导：**
- 需要额外的数组存储平方结果：$O(n)$
- 排序算法可能需要额外空间（取决于实现）

### 5.2 双指针解法复杂度

| 复杂度类型 | 时间复杂度 | 空间复杂度 |
|-----------|-----------|-----------|
| **双指针** | $O(n)$ | $O(n)$ |

**时间复杂度推导：**
- 左右指针各移动 n 次
- 每次比较和赋值是 O(1)
- 总时间复杂度：$O(n)$

**空间复杂度推导：**
- 需要额外的数组存储结果：$O(n)$
- 不需要递归栈空间

### 5.3 复杂度对比

| 方法 | 时间复杂度 | 空间复杂度 | 优点 | 缺点 |
|------|-----------|-----------|------|------|
| 暴力解法 | $O(n \log n)$ | $O(n)$ | 代码简单 | 时间复杂度高 |
| 双指针 | $O(n)$ | $O(n)$ | 时间最优 | 需要理解算法思想 |

---

## 六、 测试用例验证

### 6.1 测试用例 1：混合正负数

**输入：**
```
nums = [-4, -1, 0, 3, 10]
```

**执行过程（双指针）：**
1. left=0, right=4: max(16,100)=100, result[4]=100, right--
2. left=0, right=3: max(16,9)=16, result[3]=16, left++
3. left=1, right=3: max(1,9)=9, result[2]=9, right--
4. left=1, right=2: max(1,0)=1, result[1]=1, right--
5. left=1, right=1: max(1,1)=1, result[0]=1, 结束

**输出：** `[1, 9, 16, 16, 100]`

### 6.2 测试用例 2：全负数

**输入：**
```
nums = [-10, -3, -2, -1]
```

**执行过程（双指针）：**
1. left=0, right=3: max(100,1)=100, result[3]=100, left++
2. left=1, right=3: max(9,1)=9, result[2]=9, left++
3. left=2, right=3: max(4,1)=4, result[1]=4, left++
4. left=3, right=3: max(1,1)=1, result[0]=1, 结束

**输出：** `[1, 4, 9, 100]`

### 6.3 测试用例 3：全正数

**输入：**
```
nums = [1, 2, 3, 4, 5]
```

**执行过程（双指针）：**
1. left=0, right=4: max(1,25)=25, result[4]=25, right--
2. left=0, right=3: max(1,16)=16, result[3]=16, right--
3. left=0, right=2: max(1,9)=9, result[2]=9, right--
4. left=0, right=1: max(1,4)=4, result[1]=4, right--
5. left=0, right=0: max(1,1)=1, result[0]=1, 结束

**输出：** `[1, 4, 9, 16, 25]`

### 6.4 测试用例 4：单元素

**输入：**
```
nums = [-2]
```

**执行过程（双指针）：**
1. left=0, right=0: max(4,4)=4, result[0]=4, 结束

**输出：** `[4]`

### 6.5 测试用例 5：零元素

**输入：**
```
nums = [0]
```

**执行过程（双指针）：**
1. left=0, right=0: max(0,0)=0, result[0]=0, 结束

**输出：** `[0]`

### 6.6 测试用例 6：正负交替

**输入：**
```
nums = [-2, -1, 0, 1, 2]
```

**执行过程（双指针）：**
1. left=0, right=4: max(4,4)=4, result[4]=4, （相等，选右）
2. left=0, right=3: max(4,1)=4, result[3]=4, left++
3. left=1, right=3: max(1,1)=1, result[2]=1, right--
4. left=1, right=2: max(1,0)=1, result[1]=1, right--
5. left=1, right=1: max(1,1)=1, result[0]=1, 结束

**输出：** `[1, 1, 1, 4, 4]`

---

## 七、 常见错误与注意事项

### 7.1 暴力解法的陷阱

**错误示例：**
```cpp
// 错误：直接对原数组平方后排序（修改了原数组）
for (int i = 0; i < nums.size(); ++i) {
    nums[i] = nums[i] * nums[i];
}
sort(nums.begin(), nums.end());
return nums;
```

**后果：**
- 如果题目要求不能修改原数组，会导致错误
- 虽然本题没有这个要求，但不是好的编程习惯

**正确做法：**
```cpp
// 正确：创建新的数组存储平方
vector<int> squares(nums.size());
for (int i = 0; i < nums.size(); ++i) {
    squares[i] = nums[i] * nums[i];
}
sort(squares.begin(), squares.end());
return squares;
```

### 7.2 双指针填充顺序错误

**错误示例：**
```cpp
// 错误：从头开始填充
int index = 0;
while (left <= right) {
    if (leftSquare > rightSquare) {
        result[index] = leftSquare;  // 从头填充是错误的！
        left++;
    }
    index++;
}
```

**后果：**
- 填充顺序与选择顺序一致，都是从大到小
- 最终结果是降序的，而不是题目要求的升序

**正确做法：**
```cpp
// 正确：从末尾开始填充
int index = n - 1;
while (left <= right) {
    if (leftSquare > rightSquare) {
        result[index] = leftSquare;
        left++;
    }
    index--;  // 填充位置左移
}
```

### 7.3 指针边界条件错误

**错误示例：**
```cpp
// 错误：循环条件使用 < 而不是 <=
while (left < right) {
    // 当 left == right 时停止，但此时还没有处理中间元素
    ...
}
// 如果最后一个元素是最大的平方，会漏掉它
```

**后果：**
- 当 left == right 时停止循环
- 中间位置的元素没有被处理
- 可能漏掉某些元素

**正确做法：**
```cpp
// 正确：循环条件使用 <=
while (left <= right) {
    // 当 left == right 时继续处理，确保所有元素都被考虑
    ...
}
```

### 7.4 忘记处理相等情况

**错误示例：**
```cpp
// 错误：只处理大于的情况
if (leftSquare > rightSquare) {
    result[index] = leftSquare;
    left++;
} else {
    result[index] = rightSquare;
    right--;
}
// 这个其实是正确的，因为 else 包含了相等情况
```

**说明：**
- 相等时选择哪端都可以
- 但更推荐使用 `>=` 或 `<` 来明确处理逻辑

**推荐做法：**
```cpp
// 推荐：明确处理相等情况
if (leftSquare >= rightSquare) {  // 使用 >= 保证左优先
    result[index] = leftSquare;
    left++;
} else {
    result[index] = rightSquare;
    right--;
}
```

### 7.5 平方计算溢出

**错误示例：**
```cpp
// 错误：如果 nums[i] 是 INT_MAX，平方会溢出
int square = nums[i] * nums[i];  // 可能溢出！
```

**后果：**
- 当 nums[i] = 46340 时，平方 = 2147395600 < INT_MAX
- 但如果用 long long 或更大范围，应该没问题
- 本题限制 nums[i] <= 10000，平方 <= 100000000，安全

**正确做法：**
```cpp
// 正确：使用 long long 避免潜在溢出
long long square = (long long)nums[i] * nums[i];
```

---

## 八、 与其他题目的对比分析

### 8.1 与 88. 合并两个有序数组的对比

| 题目 | 核心问题 | 数组特点 | 解法 |
|------|---------|---------|------|
| 977. 有序数组的平方 | 平方后排序 | 单数组，已排序 | 双指针从两端向中间 |
| 88. 合并两个有序数组 | 合并两个数组 | 两数组，已排序 | 双指针从后向前 |

**共同点：**
- 都利用了数组已排序的特性
- 都使用双指针技巧
- 都需要逆序填充结果

**不同点：**
- 977 题是在同一数组的两端之间选择
- 88 题是在两个不同数组之间选择

### 8.2 与 26. 删除有序数组中的重复项的对比

| 题目 | 核心问题 | 指针方向 | 结果 |
|------|---------|---------|------|
| 977. 有序数组的平方 | 找最大平方 | 从两端向中间 | 数组平方 |
| 26. 删除重复项 | 去重 | 从前向后 | 修改数组长度 |

**共同点：**
- 都使用双指针技巧
- 都利用数组已排序的特性

**不同点：**
- 977 题需要两个指针从两端向中间
- 26 题两个指针都从前往后
- 977 题需要额外的数组存储结果

---

## 九、 面试高频问题与解答

### 9.1 基础问题

#### Q1：为什么暴力解法不是最优解法？

**答案：**
- 暴力解法的时间复杂度是 $O(n \log n)$（平方 $O(n)$ + 排序 $O(n \log n)$）
- 没有充分利用数组已排序的特性
- 双指针解法可以达到 $O(n)$ 的时间复杂度

#### Q2：双指针的核心思想是什么？

**答案：**
- 利用数组两端是绝对值最大元素的特性
- 最大平方只可能来自数组的两端
- 每次选择较大的平方后，继续收缩指针
- 结果数组从末尾开始填充，保证升序

#### Q3：为什么结果数组要从末尾开始填充？

**答案：**
- 因为我们从大到小选择元素
- 最大的平方首先被确定
- 从末尾填充可以直接按位置存放，不需要再次排序

### 9.2 进阶问题

#### Q4：能否使用单指针解决这道题？

**答案：**
- 单指针无法高效解决
- 因为需要比较数组两端的元素
- 单指针只能顺序遍历，无法利用两端的特性

#### Q5：双指针的时间复杂度为什么是 O(n)？

**答案：**
- left 指针从 0 移动到 n-1，最多 n 次
- right 指针从 n-1 移动到 0，最多 n 次
- 每次移动都是 O(1) 操作
- 总共最多 2n 次移动，即 O(n)

#### Q6：如果不使用额外的数组，空间复杂度是多少？

**答案：**
- 如果允许修改原数组，可以原地平方后使用类似冒泡的方式排序
- 但这样时间复杂度会变成 $O(n^2)$
- 或者可以使用 $O(n)$ 的原地算法（如快速排序的分区思想），但更复杂
- 通常推荐使用 $O(n)$ 的额外空间换取 $O(n)$ 的时间

### 9.3 实际问题

#### Q7：在实际工程中，什么时候选择暴力解法？

**答案：**
- 代码可读性要求高时
- 数据规模较小（n < 1000）时
- 开发时间有限，需要快速实现时

#### Q8：如何优化双指针解法？

**答案：**
- 可以使用 `abs()` 函数代替手动平方
- 可以提前判断是否全为正数或全为负数，跳过不必要的比较
- 可以使用 SIMD 指令加速计算（但对于本题可能过度优化）

---

## 十、 完整可运行代码

```cpp
#include <iostream>
#include <vector>
#include <algorithm>
#include <cmath>
using namespace std;

// 解法 1：暴力解法
class BruteForceSolution {
public:
    vector<int> sortedSquares(vector<int>& nums) {
        vector<int> squares;
        
        for (int num : nums) {
            squares.push_back(num * num);
        }
        
        sort(squares.begin(), squares.end());
        
        return squares;
    }
};

// 解法 2：双指针解法
class TwoPointersSolution {
public:
    vector<int> sortedSquares(vector<int>& nums) {
        int n = nums.size();
        vector<int> result(n);
        
        int left = 0;
        int right = n - 1;
        int index = n - 1;
        
        while (left <= right) {
            int leftSquare = nums[left] * nums[left];
            int rightSquare = nums[right] * nums[right];
            
            if (leftSquare > rightSquare) {
                result[index] = leftSquare;
                left++;
            } else {
                result[index] = rightSquare;
                right--;
            }
            index--;
        }
        
        return result;
    }
};

// 解法 3：双指针简洁版
class TwoPointersConciseSolution {
public:
    vector<int> sortedSquares(vector<int>& nums) {
        int n = nums.size();
        vector<int> result(n);
        
        int left = 0, right = n - 1;
        for (int i = n - 1; i >= 0; --i) {
            if (abs(nums[left]) > abs(nums[right])) {
                result[i] = nums[left] * nums[left];
                left++;
            } else {
                result[i] = nums[right] * nums[right];
                right--;
            }
        }
        
        return result;
    }
};

// 辅助函数：打印数组
void printArray(const vector<int>& arr, const string& name) {
    cout << name << ": [";
    for (size_t i = 0; i < arr.size(); ++i) {
        cout << arr[i];
        if (i < arr.size() - 1) {
            cout << ", ";
        }
    }
    cout << "]" << endl;
}

// 主函数测试
int main() {
    cout << "========================================" << endl;
    cout << "     977. 有序数组的平方 测试" << endl;
    cout << "========================================" << endl;

    // 创建解法实例
    BruteForceSolution bruteForceSol;
    TwoPointersSolution twoPointersSol;
    TwoPointersConciseSolution twoPointersConciseSol;

    // 测试用例 1：混合正负数
    cout << "\n--- 测试用例 1 ---" << endl;
    vector<int> nums1 = {-4, -1, 0, 3, 10};
    cout << "输入：[-4, -1, 0, 3, 10]" << endl;
    printArray(bruteForceSol.sortedSquares(nums1), "暴力解法");
    printArray(twoPointersSol.sortedSquares(nums1), "双指针解法");
    cout << "期望：[0, 1, 9, 16, 100]" << endl;

    // 测试用例 2：全负数
    cout << "\n--- 测试用例 2 ---" << endl;
    vector<int> nums2 = {-10, -3, -2, -1};
    cout << "输入：[-10, -3, -2, -1]" << endl;
    printArray(bruteForceSol.sortedSquares(nums2), "暴力解法");
    printArray(twoPointersSol.sortedSquares(nums2), "双指针解法");
    cout << "期望：[1, 4, 9, 100]" << endl;

    // 测试用例 3：全正数
    cout << "\n--- 测试用例 3 ---" << endl;
    vector<int> nums3 = {1, 2, 3, 4, 5};
    cout << "输入：[1, 2, 3, 4, 5]" << endl;
    printArray(bruteForceSol.sortedSquares(nums3), "暴力解法");
    printArray(twoPointersSol.sortedSquares(nums3), "双指针解法");
    cout << "期望：[1, 4, 9, 16, 25]" << endl;

    // 测试用例 4：单元素
    cout << "\n--- 测试用例 4 ---" << endl;
    vector<int> nums4 = {-2};
    cout << "输入：[-2]" << endl;
    printArray(bruteForceSol.sortedSquares(nums4), "暴力解法");
    printArray(twoPointersSol.sortedSquares(nums4), "双指针解法");
    cout << "期望：[4]" << endl;

    // 测试用例 5：零元素
    cout << "\n--- 测试用例 5 ---" << endl;
    vector<int> nums5 = {0};
    cout << "输入：[0]" << endl;
    printArray(bruteForceSol.sortedSquares(nums5), "暴力解法");
    printArray(twoPointersSol.sortedSquares(nums5), "双指针解法");
    cout << "期望：[0]" << endl;

    // 测试用例 6：正负交替
    cout << "\n--- 测试用例 6 ---" << endl;
    vector<int> nums6 = {-2, -1, 0, 1, 2};
    cout << "输入：[-2, -1, 0, 1, 2]" << endl;
    printArray(bruteForceSol.sortedSquares(nums6), "暴力解法");
    printArray(twoPointersSol.sortedSquares(nums6), "双指针解法");
    cout << "期望：[0, 1, 4, 4, 4]" << endl;

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
| 暴力解法 | 先平方后排序，时间复杂度 $O(n \log n)$ |
| 双指针 | 从两端向中间收缩，时间复杂度 $O(n)$ |
| 逆序填充 | 结果数组从末尾开始填充，保证升序 |
| 数组已排序 | 充分利用已排序特性 |
| 时间复杂度 | 暴力 $O(n \log n)$，双指针 $O(n)$ |
| 空间复杂度 | 都是 $O(n)$ |

### 11.2 学习收获

1. **理解了双指针技巧的应用场景**
   - 当需要从两端向中间处理时
   - 当结果需要逆序填充时
   - 当数组已排序且需要高效遍历时

2. **掌握了如何优化算法**
   - 从暴力解法到双指针的优化
   - 时间复杂度从 $O(n \log n)$ 降到 $O(n)$
   - 空间换时间的策略

3. **学会了处理边界情况**
   - 全负数、全正数
   - 单元素、零元素
   - 正负交替

4. **理解了双指针的核心思想**
   - 利用数组两端的特性
   - 每次选择最优的候选
   - 逐步缩小搜索范围

### 11.3 后续学习建议

1. **练习更多双指针问题**
   - 167. 两数之和 II - 输入有序数组
   - 344. 反转字符串
   - 283. 移动零

2. **深入理解双指针变体**
   - 快慢指针
   - 滑动窗口
   - 左右指针

3. **学习其他排序技巧**
   - 原地排序算法
   - 外部排序算法
   - 分布式排序

4. **理解时间空间权衡**
   - 何时用空间换时间
   - 何时用时间换空间
   - 根据具体场景选择

---

## 💡 面试核心考点

* **直击灵魂的拷问一：为什么双指针从两端向中间而不是从中间向两端？**
  绝杀回答：因为平方后的最大值只可能来自数组的两端（最小负数或最大正数）。从两端向中间可以保证每次选择较大的平方，而从中间向两端无法保证这个特性。

* **直击灵魂的拷问二：为什么结果数组要从末尾开始填充？**
  绝杀回答：因为我们从大到小选择元素。最大的平方首先被确定，从末尾填充可以直接按位置存放，不需要再次排序，最终得到的就是升序结果。

* **直击灵魂的拷问三：双指针的时间复杂度为什么是 O(n) 而不是 O(n^2)？**
  绝杀回答：虽然有两层逻辑（循环和比较），但实际上左右指针各移动 n 次，每次操作都是 O(1)，总共最多 2n 次操作，所以时间复杂度是 O(n)。

* **直击灵魂的拷问四：双指针和暴力解法的本质区别是什么？**
  绝杀回答：暴力解法先平方后排序，时间复杂度 O(n log n)；双指针利用数组已排序的特性，从两端向中间选择较大的平方，时间复杂度 O(n)。本质区别是是否充分利用了数组已排序的信息。

通过本课程，我们深入理解了有序数组平方问题的核心逻辑，掌握了暴力解法和双指针优化解法，也学习了常见的错误注意事项。这些知识对于理解双指针技巧和算法优化非常重要！