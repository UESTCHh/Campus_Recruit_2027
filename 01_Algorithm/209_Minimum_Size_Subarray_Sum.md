# 209. 长度最小的子数组学习笔记

## 一、 题目分析

### 1.1 题目描述

给定一个含有 `n` 个正整数的数组和一个正整数 `target`。找出该数组中满足其和 $\ge target$ 的长度最小的 **连续子数组**，并返回其长度。如果不存在符合条件的子数组，返回 0。

### 1.2 题目链接

- [LeetCode 209. 长度最小的子数组](https://leetcode.cn/problems/minimum-size-subarray-sum/description/)
- [代码随想录题解](https://leetcode.cn/problems/minimum-size-subarray-sum/solutions/1706223/by-carlsun-2-iiee/)
- [官方题解](https://leetcode.cn/problems/minimum-size-subarray-sum/solutions/305704/chang-du-zui-xiao-de-zi-shu-zu-by-leetcode-solutio/)

### 1.3 示例

**示例 1：**

```
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
```

**示例 2：**

```
输入：target = 4, nums = [1,4,4]
输出：1
```

**示例 3：**

```
输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0
```

### 1.4 提示

- 1 ≤ target ≤ 10⁹
- 1 ≤ nums.length ≤ 10⁵
- 1 ≤ nums[i] ≤ 10⁴

### 1.5 题目本质

长度最小的子数组的本质是：**在满足 "和 ≥ target" 的约束条件下，找到一个最短的连续区间**。

换句话说，我们需要：
1. 利用数组中**元素都是正整数**的关键特性
2. 动态调整窗口大小，在满足条件的前提下尽可能缩小窗口
3. 记录所有满足条件的窗口中最小的那个长度

### 1.6 题目进阶

如果你已经实现 $O(n)$ 时间复杂度的解法，请尝试设计一个 $O(n \log n)$ 时间复杂度的解法（前缀和 + 二分查找）。

---

## 二、 解题思路分析

### 2.1 核心思路

解决长度最小的子数组问题的核心思路是：**使用滑动窗口（Sliding Window）技术，动态调整窗口大小**。

在解题过程中，我们需要注意以下几点：

| 关键点 | 说明 |
|-------|------|
| 元素都是正整数 | 窗口扩大，和单调递增；窗口缩小，和单调递减 |
| 窗口指针的选择 | for 循环的指针必须是右边界（终点），不能是左边界 |
| 窗口缩小条件 | 必须用 while 而不是 if，确保尽可能缩小窗口 |
| 结果初始值 | 初始值设为 INT_MAX 或类似的极大值，最后判断是否更新过 |

### 2.2 滑动窗口原理选择

对于滑动窗口问题，我们需要使用**右指针不断探索，左指针动态收缩**的方式，原因如下：

- 我们需要先让窗口"扩大"到满足条件，然后再尝试"缩小"
- 因为元素都是正整数，窗口大小的变化与和的变化是单调关系
- 这种方式可以保证我们以 $O(n)$ 的时间复杂度完成搜索

### 2.3 滑动窗口思路

#### 2.3.1 滑动窗口三步曲

**1. 初始化变量**
- `result`：记录最小长度，初始为 `INT_MAX`
- `sum`：记录当前窗口内的元素和，初始为 0
- `left`：窗口左边界，初始为 0

**2. 窗口滑动过程**

在每次循环中：

| 步骤 | 操作 | 目的 |
|------|------|------|
| 步骤 1 | `sum += nums[right]` | 右边界向右移动，扩大窗口，加入新元素 |
| 步骤 2 | `while (sum >= target)` | 判断当前窗口是否满足条件 |
| 步骤 3 | `subLength = right - left + 1` | 计算当前窗口长度 |
| 步骤 4 | `result = min(result, subLength)` | 更新最小长度 |
| 步骤 5 | `sum -= nums[left]` | 吐出左边界元素，缩小窗口 |
| 步骤 6 | `left++` | 左边界向右移动 |

**3. 返回结果**
- 如果 `result` 还是初始值，返回 0，否则返回 `result`

#### 2.3.2 滑动窗口图解

```
示例：target = 7, nums = [2,3,1,2,4,3]

初始状态：
sum = 0, left = 0, result = ∞
窗口：[ ]

第 1 步：right = 0 (元素 2)
sum = 2
sum < 7，继续
窗口：[2]

第 2 步：right = 1 (元素 3)
sum = 5
sum < 7，继续
窗口：[2,3]

第 3 步：right = 2 (元素 1)
sum = 6
sum < 7，继续
窗口：[2,3,1]

第 4 步：right = 3 (元素 2)
sum = 8 ≥ 7
计算长度：3-0+1 = 4，result = 4
开始缩小窗口：
sum -= nums[0] (2), sum = 6, left = 1
sum < 7，停止缩小
窗口：[3,1,2]

第 5 步：right = 4 (元素 4)
sum = 6+4 = 10 ≥ 7
计算长度：4-1+1 = 4，result 仍为 4
开始缩小窗口：
sum -= nums[1] (3), sum = 7, left = 2
计算长度：4-2+1 = 3，result = 3
sum 仍 ≥ 7，继续缩小：
sum -= nums[2] (1), sum = 6, left = 3
sum < 7，停止缩小
窗口：[2,4]

第 6 步：right = 5 (元素 3)
sum = 6+3 = 9 ≥ 7
计算长度：5-3+1 = 3，result 仍为 3
开始缩小窗口：
sum -= nums[3] (2), sum = 7, left = 4
计算长度：5-4+1 = 2，result = 2
sum 仍 ≥ 7，继续缩小：
sum -= nums[4] (4), sum = 3, left = 5
sum < 7，停止缩小
窗口：[3]

最终：result = 2，返回 2
```

### 2.4 前缀和 + 二分查找思路

#### 2.4.1 前缀和思路

**1. 构造前缀和数组**
- `preSum[i]` 表示前 `i` 个元素的和
- `preSum[0] = 0`，`preSum[1] = nums[0]`，`preSum[2] = nums[0] + nums[1]`，依此类推

**2. 二分查找**
- 对于每个位置 `i`，我们需要找到最小的 `j > i`，使得 `preSum[j] - preSum[i] ≥ target`
- 即 `preSum[j] ≥ preSum[i] + target`
- 因为原数组元素都是正整数，前缀和数组是严格递增的，可以使用二分查找

### 2.5 核心思想：动态伸缩的毛毛虫

暴力解法是用两个 `for` 循环寻找所有可能的子数组，时间复杂度是极其糟糕的 $O(n^2)$。

滑动窗口的核心思想是：**像一个拥有弹性的毛毛虫，在数组上爬行。** 我们只用一个 `for` 循环，就能完成原先需要两个循环才能做完的事。这里的核心痛点在于：这个唯一的循环指针，到底该指向窗口的**起始位置**，还是**结束位置**？

---

## 三、 算法实现

### 3.1 滑动窗口法（推荐）

```cpp
class Solution {
public:
    int minSubArrayLen(int target, vector<int>& nums) {
        int result = INT32_MAX; // 记录最终的最小长度，初始设为一个极大的数
        int sum = 0;            // 记录当前滑动窗口内元素的总和
        int left = 0;           // 滑动窗口的起始位置（慢指针）

        // right 是滑动窗口的结束位置（快指针）
        // 💡 记住：for 循环的变量永远是探路的先锋（右边界）！
        for (int right = 0; right < nums.size(); right++) {
            sum += nums[right]; // 将当前右边界的元素吃进窗口

            // 💡 核心挤压逻辑：一旦窗口内的总和满足条件，就开始尝试缩小窗口
            // 必须是 while，不能是 if！
            while (sum >= target) {
                // 1. 抓取当前满足条件的窗口长度，并更新最短记录
                int subLength = right - left + 1;
                result = result < subLength ? result : subLength;
                
                // 2. 核心挤压动作：把窗口最左边的元素吐出来，左指针向右缩进
                sum -= nums[left]; 
                left++; 
            }
        }

        // 如果 result 还是初始的最大值，说明遍历完都没找到，按题意返回 0
        return result == INT32_MAX ? 0 : result;
    }
};
```

### 3.2 前缀和 + 二分查找法

```cpp
class Solution {
public:
    int minSubArrayLen(int target, vector<int>& nums) {
        int n = nums.size();
        if (n == 0) return 0;
        
        int result = INT32_MAX;
        vector<int> preSum(n + 1, 0); // preSum[0] = 0
        
        // 构造前缀和数组
        for (int i = 0; i < n; i++) {
            preSum[i + 1] = preSum[i] + nums[i];
        }
        
        // 对于每个 i，查找最小的 j 使得 preSum[j] >= preSum[i] + target
        for (int i = 0; i <= n; i++) {
            int targetSum = preSum[i] + target;
            
            // 使用 lower_bound 在 preSum 数组中二分查找
            auto bound = lower_bound(preSum.begin(), preSum.end(), targetSum);
            
            if (bound != preSum.end()) {
                // 找到了，计算长度
                int subLength = bound - preSum.begin() - i;
                result = min(result, subLength);
            }
        }
        
        return result == INT32_MAX ? 0 : result;
    }
};
```

### 3.3 手写二分查找版本

```cpp
class Solution {
public:
    int minSubArrayLen(int target, vector<int>& nums) {
        int n = nums.size();
        if (n == 0) return 0;
        
        int result = INT32_MAX;
        vector<int> preSum(n + 1, 0);
        
        // 构造前缀和数组
        for (int i = 0; i < n; i++) {
            preSum[i + 1] = preSum[i] + nums[i];
        }
        
        // 对于每个 i，查找最小的 j 使得 preSum[j] >= preSum[i] + target
        for (int i = 0; i <= n; i++) {
            int targetSum = preSum[i] + target;
            
            // 手写二分查找
            int left = i + 1, right = n;
            while (left <= right) {
                int mid = left + (right - left) / 2;
                if (preSum[mid] >= targetSum) {
                    result = min(result, mid - i);
                    right = mid - 1; // 找更小的 j
                } else {
                    left = mid + 1;
                }
            }
        }
        
        return result == INT32_MAX ? 0 : result;
    }
};
```

---

## 四、 代码分析

### 4.1 滑动窗口代码分析

#### 4.1.1 实现思路说明

**滑动窗口的设计思路：**
1. 使用 `right` 指针不断向右扩展窗口
2. 当窗口内的和满足条件时，使用 `left` 指针尽可能向右缩小窗口
3. 每次找到满足条件的窗口时，更新最小长度
4. 最终返回找到的最小长度

#### 4.1.2 核心逻辑讲解

**关键代码段 1：窗口初始化**
```cpp
int result = INT32_MAX;
int sum = 0;
int left = 0;
```
- `result` 初始设为极大值，用于后续比较更新
- `sum` 记录窗口内元素的和
- `left` 是窗口的左边界

**关键代码段 2：窗口扩展**
```cpp
sum += nums[right];
```
- 将右指针指向的元素加入窗口
- 窗口向右扩展

**关键代码段 3：窗口收缩（核心逻辑）**
```cpp
while (sum >= target) {
    int subLength = right - left + 1;
    result = min(result, subLength);
    sum -= nums[left];
    left++;
}
```
- 必须用 `while` 而不是 `if`，确保能把窗口缩到最小
- 每次都要记录当前窗口长度并更新最小值
- 然后缩小窗口，继续尝试

#### 4.1.3 初学者引导

**为什么外层循环用右指针而不是左指针？**
- 如果外层循环是左指针，那么为了找到满足条件的右指针，你就不得不再嵌套一个循环
- 这样就退化成了 $O(n^2)$ 的暴力解法
- 外层循环必须是右指针在前面"探路"

**为什么窗口收缩必须用 while？**
- 假设 `target = 7`，当前窗口是 `[1,1,1,1,100]`
- 用 `if` 的话，只吐出一个元素，和还是 103，仍然满足条件
- 但我们错失了继续缩小窗口的机会
- 用 `while` 才能保证尽可能地缩小窗口

### 4.2 前缀和+二分查找代码分析

#### 4.2.1 实现思路说明

**前缀和 + 二分查找的设计思路：**
1. 先构造前缀和数组
2. 对于每个起始位置 `i`，用二分查找找最小的结束位置 `j`
3. 找到满足 `preSum[j] - preSum[i] >= target` 的最小 `j`
4. 记录所有可能的窗口长度中的最小值

#### 4.2.2 核心逻辑讲解

**关键代码段 1：构造前缀和**
```cpp
vector<int> preSum(n + 1, 0);
for (int i = 0; i < n; i++) {
    preSum[i + 1] = preSum[i] + nums[i];
}
```
- 前缀和数组长度为 `n+1`，`preSum[0] = 0`
- `preSum[i]` 表示前 `i` 个元素的和

**关键代码段 2：二分查找**
```cpp
int targetSum = preSum[i] + target;
auto bound = lower_bound(preSum.begin(), preSum.end(), targetSum);
```
- `lower_bound` 找第一个大于等于 `targetSum` 的位置
- 这个位置就是满足条件的最小 `j`

---

## 五、 时间与空间复杂度分析

### 5.1 滑动窗口复杂度分析

| 复杂度类型 | 时间复杂度 | 空间复杂度 |
|-----------|-----------|-----------|
| **滑动窗口法** | $O(n)$ | $O(1)$ |

**时间复杂度推导（极其重要：为什么嵌套了 `while` 依然是线性的？）：**
- 不要被外层的 `for` 循环和内层的 `while` 循环骗了
- 分析时间复杂度的核心是看 **"每一个元素被操作的真正次数"**
- 在滑动窗口算法中：
  - **进窗口：** 右指针从左向右滑动，每个元素最多被加入一次
  - **出窗口：** 左指针从左向右滑动，每个元素最多被移出一次
- **结论：** 每个元素最多被操作两次（一进一出），总体操作次数是 $2N$，时间复杂度 $O(n)$

**空间复杂度推导：**
- 只使用了几个整型变量：`result`、`sum`、`left`、`right`
- 没有使用与数据规模相关的额外空间
- 空间复杂度为严格的 $O(1)$

### 5.2 前缀和 + 二分查找复杂度分析

| 复杂度类型 | 时间复杂度 | 空间复杂度 |
|-----------|-----------|-----------|
| **前缀和 + 二分查找** | $O(n \log n)$ | $O(n)$ |

**时间复杂度推导：**
- 构造前缀和数组：$O(n)$
- 对于每个位置 $i$，二分查找：$O(\log n)$
- 总共有 $n$ 个位置，所以总时间复杂度：$O(n \log n)$

**空间复杂度推导：**
- 需要额外的前缀和数组，大小为 $n+1$
- 空间复杂度为 $O(n)$

### 5.3 复杂度对比表格

| 方法 | 时间复杂度 | 空间复杂度 | 优点 | 缺点 |
|------|-----------|-----------|------|------|
| 滑动窗口法 | $O(n)$ | $O(1)$ | 时空复杂度最优，实现简单 | 只适用于正整数数组 |
| 前缀和+二分查找 | $O(n \log n)$ | $O(n)$ | 思路巧妙，不需要正整数限制（只要前缀和单调） | 需要额外空间，实现稍复杂 |

---

## 六、 测试用例验证

### 6.1 测试用例 1：正常情况

**输入：**
```
target = 7, nums = [2,3,1,2,4,3]
```

**执行过程（滑动窗口法）：**
1. 右指针移动，窗口扩大，直到 `sum = 8 ≥ 7`
2. 记录长度 4，开始缩小窗口
3. 继续移动右指针，找到更优解
4. 最终找到长度为 2 的窗口 `[4,3]`

**输出：** `2`

### 6.2 测试用例 2：单个元素满足

**输入：**
```
target = 4, nums = [1,4,4]
```

**执行过程：**
1. 右指针移动到索引 1（元素 4）
2. `sum = 4 ≥ 4`，记录长度 1
3. 继续移动右指针到索引 2（元素 4），同样满足
4. 最终最小长度为 1

**输出：** `1`

### 6.3 测试用例 3：无解

**输入：**
```
target = 11, nums = [1,1,1,1,1,1,1,1]
```

**执行过程：**
1. 右指针移动到末尾，`sum = 8 < 11`
2. 从未满足条件，`result` 保持 `INT_MAX`
3. 最终返回 0

**输出：** `0`

### 6.4 测试用例 4：单个元素

**输入：**
```
target = 5, nums = [5]
```

**执行过程：**
1. 右指针移动到索引 0
2. `sum = 5 ≥ 5`，记录长度 1
3. 返回 1

**输出：** `1`

---

## 七、 常见错误与注意事项

### 7.1 窗口指针选择错误

**错误示例：**
```cpp
// 错误：for 循环用左指针，内部循环找右指针
for (int left = 0; left < nums.size(); left++) {
    for (int right = left; right < nums.size(); right++) {
        // 退化成 O(n^2) 暴力解法
    }
}
```

**后果：**
- 时间复杂度退化成 $O(n^2)$
- 完全失去了滑动窗口的意义

**正确做法：**
```cpp
for (int right = 0; right < nums.size(); right++) {
    // ...
}
```

### 7.2 使用 if 而不是 while 来缩小窗口

**错误示例：**
```cpp
// 错误：只用 if，窗口没有缩到最小
if (sum >= target) {
    // 只缩小一次
    sum -= nums[left];
    left++;
}
```

**后果：**
- 比如 `nums = [1,1,1,100], target = 7`
- 当 `right` 到 100 时，`sum = 103`，只用 `if` 只会缩小到 102
- 错失了继续缩小的机会

**正确做法：**
```cpp
while (sum >= target) {
    // 持续缩小直到不满足条件
    sum -= nums[left];
    left++;
}
```

### 7.3 忘记处理无解情况

**错误示例：**
```cpp
// 错误：直接返回 result，可能是 INT_MAX
return result;
```

**后果：**
- 当没有满足条件的子数组时，会返回一个很大的数
- 不符合题目要求返回 0

**正确做法：**
```cpp
return result == INT32_MAX ? 0 : result;
```

### 7.4 初始值设置错误

**错误示例：**
```cpp
// 错误：初始值设为 0
int result = 0;
```

**后果：**
- 最小值会永远是 0
- 无法正确更新

**正确做法：**
```cpp
int result = INT32_MAX; // 或者 INT_MAX
```

### 7.5 窗口长度计算错误

**错误示例：**
```cpp
// 错误：忘记 +1
int subLength = right - left;
```

**后果：**
- 长度少算了 1
- 比如 `left = 0, right = 1`，正确长度是 2，错误计算成 1

**正确做法：**
```cpp
int subLength = right - left + 1;
```

---

## 八、 与其他相关题目的对比分析

### 8.1 与 3. 无重复字符的最长子串对比

| 题目 | 核心问题 | 约束条件 | 窗口操作 |
|------|---------|---------|---------|
| 209. 长度最小的子数组 | 找满足和 ≥ target 的最短子数组 | 和 ≥ target | 先扩再缩 |
| 3. 无重复字符的最长子串 | 找无重复的最长子串 | 无重复 | 先扩，遇到重复时缩 |

**相同点：**
- 都是滑动窗口问题
- 都要求 $O(n)$ 时间复杂度

**不同点：**
- 209 题要求最短，3 题要求最长
- 209 题约束条件是和 ≥ target，3 题约束条件是无重复

### 8.2 与 76. 最小覆盖子串对比

| 题目 | 核心问题 | 难度 |
|------|---------|------|
| 209. 长度最小的子数组 | 找和 ≥ target 的最短子数组 | 中等 |
| 76. 最小覆盖子串 | 找包含 t 所有字符的最短子串 | 困难 |

**相同点：**
- 都是"最小窗口"类型问题
- 都是滑动窗口问题
- 都是先扩再缩

**不同点：**
- 209 题条件简单（和 ≥ target）
- 76 题条件复杂（需要包含所有字符）

### 8.3 与 713. 乘积小于 K 的子数组对比

| 题目 | 核心问题 | 窗口收缩条件 |
|------|---------|-----------|
| 209. 长度最小的子数组 | 找和 ≥ target 的最短子数组 | sum ≥ target |
| 713. 乘积小于 K 的子数组 | 计数乘积 < K 的子数组个数 | product ≥ K |

**相同点：**
- 都是滑动窗口问题
- 都要求正整数数组
- 都是 $O(n)$ 时间复杂度

**不同点：**
- 209 题求最短，713 题求计数
- 209 题是和 ≥ target，713 题是积 < K

---

## 九、 面试高频问题与解答

### 9.1 基础问题

#### Q1：滑动窗口的核心思想是什么？

**答案：**

**核心思想：**
1. 用两个指针表示一个窗口的左右边界
2. 右指针不断向右扩展窗口
3. 当窗口满足条件时，尝试用左指针缩小窗口
4. 在窗口变化过程中记录最优解

**关键：**
- 对于"最小窗口"问题，是**先扩再缩**
- 对于"最大窗口"问题，是**扩到不能再扩，再移动左边界**

#### Q2：为什么滑动窗口的时间复杂度是 $O(n)$ 而不是 $O(n^2)$？

**答案：**

**关键：** 不要被嵌套的循环结构骗了！

- 滑动窗口中，每个元素最多被**右指针访问一次**（加入窗口）
- 每个元素最多被**左指针访问一次**（移出窗口）
- 每个元素总共被访问 2 次
- 总操作次数是 $2N$，忽略常数后是 $O(n)$

**类比：**
- 想象两个指针都从左向右走，都只走一趟
- 虽然有两个循环，但实际上是两个指针各走一遍

### 9.2 进阶问题

#### Q3：为什么滑动窗口只适用于正整数数组？

**答案：**

**核心原因：单调性！**

- 当数组元素都是正数时：
  - 窗口扩大 → 和一定单调递增
  - 窗口缩小 → 和一定单调递减

- 有了单调性，我们才能保证：
  - 一旦窗口满足条件，就可以放心地缩小窗口
  - 当窗口不满足条件时，才需要扩大窗口

**如果有负数：**
- 窗口扩大，和可能减小
- 窗口缩小，和可能增大
- 滑动窗口的单调性就被破坏了

#### Q4：滑动窗口中，左指针为什么不会往回走？

**答案：**

**因为单调性！**

- 一旦我们移动左指针缩小了窗口，说明之前更大的窗口已经被考虑过了
- 对于"最小窗口"问题，我们是在满足条件的前提下寻找更小的窗口
- 往回走只会得到更大的窗口，不是我们想要的

**总结：**
- 左右指针都只向右移动，从不回头
- 每个指针最多移动 n 次
- 保证了 $O(n)$ 的时间复杂度

#### Q5：如果题目要求的是 "等于 target" 而不是 "大于等于 target"，还能用滑动窗口吗？

**答案：**

**可以，但需要注意边界处理！**

- 条件变成 `sum == target` 时才记录
- 但是当 `sum > target` 时，依然要移动左指针缩小窗口
- 同时要注意：可能存在多个满足条件的窗口，需要都考虑到

**不过要注意：**
- 如果有负数，滑动窗口不一定适用
- 这时候可能需要前缀和 + 哈希表的方法

### 9.3 实际问题

#### Q6：在实际工程中，什么场景会用到滑动窗口？

**答案：**

**常见场景：**
1. **网络流量控制**：计算过去一段时间内的平均流量
2. **日志分析**：分析最近 N 条日志中的错误率
3. **金融分析**：计算移动平均线（Moving Average）
4. **数据流处理**：在无限数据流中找满足条件的窗口

**滑动窗口的优势：**
- 实时处理，不需要保存所有历史数据
- 效率高，O(n) 时间复杂度
- 空间复杂度低，通常 O(1) 或 O(k)，k 是窗口大小

#### Q7：如何处理环形数组的滑动窗口问题？

**答案：**

**常用技巧：**
1. **数组翻倍**：将数组复制一份接在后面，变成 `nums + nums`
2. **取模运算**：用 `right % n` 和 `left % n` 来处理环形索引
3. **特殊处理边界**：根据问题特性特殊处理跨越边界的情况

**注意：**
- 具体处理方式要看题目要求
- 有些问题可能不能简单地用数组翻倍的方法

---

## 十、 完整可运行代码

```cpp
#include <iostream>
#include <vector>
#include <climits>
#include <algorithm>
using namespace std;

// 方法一：滑动窗口法（推荐）
class SlidingWindowSolution {
public:
    int minSubArrayLen(int target, vector<int>& nums) {
        int result = INT_MAX; // 记录最终的最小长度，初始设为一个极大的数
        int sum = 0;            // 记录当前滑动窗口内元素的总和
        int left = 0;           // 滑动窗口的起始位置（慢指针）

        // right 是滑动窗口的结束位置（快指针）
        for (int right = 0; right < nums.size(); right++) {
            sum += nums[right]; // 将当前右边界的元素吃进窗口

            // 核心挤压逻辑：一旦窗口内的总和满足条件，就开始尝试缩小窗口
            while (sum >= target) {
                // 1. 抓取当前满足条件的窗口长度，并更新最短记录
                int subLength = right - left + 1;
                result = min(result, subLength);
                
                // 2. 核心挤压动作：把窗口最左边的元素吐出来，左指针向右缩进
                sum -= nums[left]; 
                left++; 
            }
        }

        // 如果 result 还是初始的最大值，说明遍历完都没找到，按题意返回 0
        return result == INT_MAX ? 0 : result;
    }
};

// 方法二：前缀和 + 二分查找法
class PrefixSumSolution {
public:
    int minSubArrayLen(int target, vector<int>& nums) {
        int n = nums.size();
        if (n == 0) return 0;
        
        int result = INT_MAX;
        vector<int> preSum(n + 1, 0); // preSum[0] = 0
        
        // 构造前缀和数组
        for (int i = 0; i < n; i++) {
            preSum[i + 1] = preSum[i] + nums[i];
        }
        
        // 对于每个 i，查找最小的 j 使得 preSum[j] >= preSum[i] + target
        for (int i = 0; i <= n; i++) {
            int targetSum = preSum[i] + target;
            
            // 使用 lower_bound 在 preSum 数组中二分查找
            auto bound = lower_bound(preSum.begin(), preSum.end(), targetSum);
            
            if (bound != preSum.end()) {
                // 找到了，计算长度
                int subLength = bound - preSum.begin() - i;
                result = min(result, subLength);
            }
        }
        
        return result == INT_MAX ? 0 : result;
    }
};

// 辅助函数：打印数组
void printArray(const vector<int>& nums) {
    cout << "[";
    for (int i = 0; i < nums.size(); i++) {
        cout << nums[i];
        if (i < nums.size() - 1) cout << ",";
    }
    cout << "]";
}

// 主函数测试
int main() {
    cout << "========================================" << endl;
    cout << "     209. 长度最小的子数组测试" << endl;
    cout << "========================================" << endl;

    // 创建解法实例
    SlidingWindowSolution slidingWindowSol;
    PrefixSumSolution prefixSumSol;

    // 测试用例 1
    cout << "\n--- 测试用例 1 ---" << endl;
    int target1 = 7;
    vector<int> nums1 = {2, 3, 1, 2, 4, 3};
    cout << "target = " << target1 << ", nums = ";
    printArray(nums1);
    cout << endl;
    cout << "滑动窗口法结果：" << slidingWindowSol.minSubArrayLen(target1, nums1) << endl;
    cout << "前缀和+二分结果：" << prefixSumSol.minSubArrayLen(target1, nums1) << endl;
    cout << "期望：2" << endl;

    // 测试用例 2
    cout << "\n--- 测试用例 2 ---" << endl;
    int target2 = 4;
    vector<int> nums2 = {1, 4, 4};
    cout << "target = " << target2 << ", nums = ";
    printArray(nums2);
    cout << endl;
    cout << "滑动窗口法结果：" << slidingWindowSol.minSubArrayLen(target2, nums2) << endl;
    cout << "前缀和+二分结果：" << prefixSumSol.minSubArrayLen(target2, nums2) << endl;
    cout << "期望：1" << endl;

    // 测试用例 3
    cout << "\n--- 测试用例 3 ---" << endl;
    int target3 = 11;
    vector<int> nums3 = {1, 1, 1, 1, 1, 1, 1, 1};
    cout << "target = " << target3 << ", nums = ";
    printArray(nums3);
    cout << endl;
    cout << "滑动窗口法结果：" << slidingWindowSol.minSubArrayLen(target3, nums3) << endl;
    cout << "前缀和+二分结果：" << prefixSumSol.minSubArrayLen(target3, nums3) << endl;
    cout << "期望：0" << endl;

    // 测试用例 4
    cout << "\n--- 测试用例 4 ---" << endl;
    int target4 = 5;
    vector<int> nums4 = {5};
    cout << "target = " << target4 << ", nums = ";
    printArray(nums4);
    cout << endl;
    cout << "滑动窗口法结果：" << slidingWindowSol.minSubArrayLen(target4, nums4) << endl;
    cout << "前缀和+二分结果：" << prefixSumSol.minSubArrayLen(target4, nums4) << endl;
    cout << "期望：1" << endl;

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
| 滑动窗口 | 双指针动态调整窗口大小 |
| 窗口指针 | for 循环变量必须是右边界 |
| 窗口收缩 | 使用 while 而不是 if |
| 单调性 | 数组元素为正，窗口大小变化与和的变化单调 |
| 时间复杂度 | 滑动窗口 O(n)，前缀和+二分 O(n log n) |
| 空间复杂度 | 滑动窗口 O(1)，前缀和+二分 O(n) |

### 11.2 学习收获

1. **理解了滑动窗口的核心原理**
   - 为什么双指针可以完成需要两个循环的事
   - 为什么虽然有嵌套循环，但时间复杂度还是 O(n)
   - 窗口指针的选择为什么很重要

2. **掌握了两种解题方法**
   - 滑动窗口法：最优解，实现简单
   - 前缀和+二分查找：巧妙思路，扩展性强

3. **学会了如何处理边界情况**
   - 没有解的情况如何处理
   - 单个元素的情况
   - 窗口长度的计算

4. **理解了滑动窗口的适用条件**
   - 为什么需要元素都是正整数
   - 单调性的重要性

### 11.3 后续学习建议

1. **练习更多滑动窗口题目**
   - 3. 无重复字符的最长子串
   - 76. 最小覆盖子串
   - 713. 乘积小于 K 的子数组
   - 239. 滑动窗口最大值

2. **深入理解二分查找**
   - lower_bound 的使用
   - 手写二分查找
   - 二分查找的各种变体

3. **学习非正整数数组的子数组问题**
   - 当有负数时，滑动窗口不再适用
   - 这时候要用前缀和 + 哈希表
   - 题目：560. 和为 K 的子数组

4. **学习动态窗口大小的其他变体**
   - 固定窗口大小的滑动窗口
   - 窗口大小变化，但有其他约束条件

---

## 💡 面试核心考点（三大直击灵魂的拷问）

* **直击灵魂的拷问一：为什么 for 循环里的指针 right 代表的是窗口的"终点"？能写成"起点"吗？**
  绝对不能！如果 for 循环的指针代表起点，那么你为了找到满足条件的终点，就不得不再嵌套一个内部循环去向后遍历。这样代码瞬间就退化成了 $O(n^2)$ 的暴力解法，滑动窗口的意义就完全丧失了。所以，外层循环必须是终点指针在前面"探路"。

* **直击灵魂的拷问二：缩小窗口时，为什么必须用 while 而不是 if？** 
  这是新手极其容易踩坑的地方！假设目标 target = 7，当前数组是 [1, 1, 1, 1, 100]。当右指针走到 100 时，窗口内的和突然变成了 104。满足 sum >= 7，触发条件。如果我们只用 if，左指针 left 只会向右移动一次，吐出一个 1，此时和为 103，依然大于 7！但代码却跳出了判断，错失了继续缩短窗口的绝佳机会。用 while 才能保证左边界持续向右挤压，直到把多余的元素全部吐干净。

* **直击灵魂的拷问三：代码里 for 循环嵌套了一个 while 循环，为什么时间复杂度依然是 $O(n)$ 而不是 $O(n^2)$？**
  不要被嵌套的循环代码结构骗了。分析时间复杂度要看每个元素被操作的次数。在这个算法中，数组里的每一个元素，最多被右指针 right 吃进去一次（进窗口），最多被左指针 left 吐出来一次（出窗口）。每个元素的操作次数是常量级别的（最多进出各一次），总操作次数约为 $2N$。因此，时间复杂度依然是极其优秀的 $O(n)$。

---

通过本课程，我们深入理解了滑动窗口算法的核心逻辑，掌握了滑动窗口和前缀和+二分查找两种解法，也学习了常见的错误和注意事项。这些知识对于理解双指针和滑动窗口非常重要！