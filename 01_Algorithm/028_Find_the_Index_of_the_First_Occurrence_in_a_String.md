# 算法实战复盘：找出字符串中第一个匹配项的下标 (LeetCode 28)

> **打卡日期**：2026-04-07
> **核心主题**：字符串匹配、暴力法、KMP算法。

---

## 📝 一、 题目描述与核心要求

### 1. 中文描述
给你两个字符串 `haystack` 和 `needle`，请你在 `haystack` 字符串中找出 `needle` 字符串的第一个匹配项的下标（下标从 0 开始）。如果 `needle` 不是 `haystack` 的一部分，则返回 `-1`。

### 2. 英文描述
Given two strings `haystack` and `needle`, return the index of the first occurrence of `needle` in `haystack`, or `-1` if `needle` is not part of `haystack`.

### 3. 输入示例
- **示例 1**：
  - 输入：`haystack = "sadbutsad", needle = "sad"`
  - 输出：`0`
  - 解释："sad" 在下标 0 和 6 处匹配。第一个匹配项的下标是 0 。

- **示例 2**：
  - 输入：`haystack = "leetcode", needle = "leeto"`
  - 输出：`-1`
  - 解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。

### 4. 核心痛点
- 高效的字符串匹配算法
- 处理边界情况（如空字符串）
- 避免不必要的比较，提高匹配效率
- 处理大字符串时的性能问题

### 5. 题目提示
- 1 <= haystack.length, needle.length <= 10^4
- haystack 和 needle 仅由小写英文字符组成

---

## 📊 二、 复杂度深度剖析

### 1. 暴力法时间复杂度：O(n*m)
**推导过程**：对于 haystack 中的每个位置，都需要与 needle 进行最多 m 次比较，其中 n 是 haystack 的长度，m 是 needle 的长度。在最坏情况下，需要进行 n*m 次比较。

### 2. KMP 算法时间复杂度：O(n+m)
**推导过程**：KMP 算法通过预处理 needle 得到 next 数组，然后在匹配过程中利用已经匹配的信息，避免重复比较。预处理时间为 O(m)，匹配时间为 O(n)，所以总时间复杂度为 O(n+m)。

### 3. 空间复杂度
- 暴力法：O(1)，不需要额外空间
- KMP 算法：O(m)，需要存储 next 数组

---

## 🤔 三、 解题思路分析

### 1. 初学者的常见思路

当面对这个问题时，初学者可能会想到以下几种方法：

1. **暴力法**：从 haystack 的每个位置开始，与 needle 进行逐字符比较
2. **KMP 算法**：利用预处理的 next 数组，避免重复比较
3. **Sunday 算法**：从右向左匹配，利用模式串的信息加速匹配过程
4. **Boyer-Moore 算法**：结合坏字符规则和好后缀规则，进一步提高匹配效率

### 2. 各种方法的优缺点分析

| 方法 | 时间复杂度 | 空间复杂度 | 优点 | 缺点 |
|------|------------|------------|------|------|
| 暴力法 | O(n*m) | O(1) | 实现简单，容易理解 | 时间复杂度高，在大字符串上性能差 |
| KMP 算法 | O(n+m) | O(m) | 时间复杂度低，性能稳定 | 实现较复杂，需要理解 next 数组的构建 |
| Sunday 算法 | O(n+m) | O(1) | 实现相对简单，在某些情况下性能优于 KMP | 最坏情况下时间复杂度仍为 O(n*m) |
| Boyer-Moore 算法 | O(n+m) | O(m) | 在实际应用中性能往往最优，特别是对于长模式串 | 实现复杂，理解难度高 |

### 3. 数据大小对方法选择的影响

根据题目提示，字符串长度最大为 10^4：

- **暴力法**：对于长度为 10^4 的字符串，最坏情况下需要 10^8 次比较，可能会超时
- **KMP 算法**：对于长度为 10^4 的字符串，预处理和匹配的时间都在可接受范围内
- **Sunday 算法**：在大多数情况下性能良好，实现相对简单
- **Boyer-Moore 算法**：实现复杂，但在实际应用中性能可能最优

**实际测试验证**：
- 暴力法：在短字符串上性能尚可，但在长字符串上可能超时
- KMP 算法：在各种情况下性能稳定，适合处理长字符串
- Sunday 算法：在实践中表现良好，代码简洁
- Boyer-Moore 算法：实现复杂，但在某些情况下性能最优

### 4. 方法选择建议

**推荐使用 KMP 算法**，原因如下：
1. 时间复杂度 O(n+m)，性能稳定
2. 空间复杂度 O(m)，在内存允许范围内
3. 适用于各种字符串匹配场景，尤其是长字符串
4. 是字符串匹配的经典算法，理解 KMP 有助于掌握更复杂的字符串算法

---

## 🎯 四、 核心思想：KMP 算法

KMP 算法是解决字符串匹配问题的经典算法，它通过预处理模式串（needle）得到 next 数组，然后在匹配过程中利用已经匹配的信息，避免重复比较。

### 1. 核心原理

1. **预处理**：构建 next 数组，其中 next[i] 表示模式串中前 i 个字符的最长相等前后缀长度
2. **匹配**：利用 next 数组，当匹配失败时，将模式串向右移动尽可能多的位置，而不是从新开始

### 2. 算法步骤详解

1. **构建 next 数组**：
   - 初始化 next[0] = 0
   - 对于 i 从 1 到 m-1：
     - 初始化 j = next[i-1]
     - 当 j > 0 且 p[i] != p[j]，则 j = next[j-1]
     - 如果 p[i] == p[j]，则 j++
     - 设置 next[i] = j

2. **匹配过程**：
   - 初始化 i = 0（haystack 指针），j = 0（needle 指针）
   - 当 i < n 且 j < m：
     - 如果 haystack[i] == needle[j]，则 i++，j++
     - 否则，如果 j > 0，则 j = next[j-1]
     - 否则，i++
   - 如果 j == m，返回 i - m
   - 否则，返回 -1

### 3. 示例分析

以示例 1 为例：`haystack = "sadbutsad", needle = "sad"`

**步骤 1**：构建 next 数组
- next[0] = 0
- i = 1, j = next[0] = 0
  - p[1] = 'a' != p[0] = 's'，所以 next[1] = 0
- i = 2, j = next[1] = 0
  - p[2] = 'd' != p[0] = 's'，所以 next[2] = 0
- next 数组为 [0, 0, 0]

**步骤 2**：匹配过程
- i = 0, j = 0
  - haystack[0] = 's' == needle[0] = 's'，i=1, j=1
- i = 1, j = 1
  - haystack[1] = 'a' == needle[1] = 'a'，i=2, j=2
- i = 2, j = 2
  - haystack[2] = 'd' == needle[2] = 'd'，i=3, j=3
- j == m，返回 i - m = 0

**最终结果**：`0`

### 4. 边界情况处理

1. **空字符串**：如果 needle 为空，返回 0；如果 haystack 为空且 needle 不为空，返回 -1
2. **needle 长度大于 haystack**：直接返回 -1
3. **单个字符**：当 needle 长度为 1 时，直接遍历 haystack 寻找匹配

### 5. 与其他方法的对比

| 方法 | 时间复杂度 | 空间复杂度 | 优点 | 缺点 |
|------|------------|------------|------|------|
| 暴力法 | O(n*m) | O(1) | 实现简单，容易理解 | 时间复杂度高，在大字符串上性能差 |
| KMP 算法 | O(n+m) | O(m) | 时间复杂度低，性能稳定 | 实现较复杂，需要理解 next 数组的构建 |
| Sunday 算法 | O(n+m) | O(1) | 实现相对简单，在某些情况下性能优于 KMP | 最坏情况下时间复杂度仍为 O(n*m) |
| Boyer-Moore 算法 | O(n+m) | O(m) | 在实际应用中性能往往最优，特别是对于长模式串 | 实现复杂，理解难度高 |

### 6. 关键技术点

1. **next 数组的构建**：理解最长相等前后缀的概念，正确计算 next 数组
2. **匹配过程中的指针移动**：利用 next 数组，避免重复比较
3. **边界条件处理**：正确处理空字符串、长度不匹配等情况
4. **性能优化**：在构建 next 数组和匹配过程中，避免不必要的计算

---

## 💻 五、 核心代码

### 1. 暴力法实现

```cpp
#include <string>
using namespace std;

class Solution {
public:
    int strStr(string haystack, string needle) {
        int n = haystack.size(), m = needle.size();
        // 边界情况处理
        if (m == 0) return 0;
        if (n < m) return -1;
        
        // 遍历 haystack 的每个位置
        for (int i = 0; i <= n - m; i++) {
            // 从 i 位置开始比较
            int j = 0;
            for (; j < m; j++) {
                if (haystack[i + j] != needle[j]) {
                    break;
                }
            }
            // 如果完全匹配，返回 i
            if (j == m) {
                return i;
            }
        }
        // 没有找到匹配
        return -1;
    }
};
```

**代码解释**：
1. **边界情况处理**：如果 needle 为空，返回 0；如果 haystack 长度小于 needle，返回 -1
2. **遍历 haystack**：从每个可能的起始位置开始尝试匹配
3. **逐字符比较**：从当前位置开始，与 needle 进行逐字符比较
4. **返回结果**：如果找到匹配，返回起始位置；否则返回 -1

### 2. KMP 算法实现

```cpp
#include <string>
#include <vector>
using namespace std;

class Solution {
public:
    int strStr(string haystack, string needle) {
        int n = haystack.size(), m = needle.size();
        // 边界情况处理
        if (m == 0) return 0;
        if (n < m) return -1;
        
        // 构建 next 数组
        vector<int> next(m);
        next[0] = 0;
        for (int i = 1, j = 0; i < m; i++) {
            while (j > 0 && needle[i] != needle[j]) {
                j = next[j - 1];
            }
            if (needle[i] == needle[j]) {
                j++;
            }
            next[i] = j;
        }
        
        // 匹配过程
        for (int i = 0, j = 0; i < n; i++) {
            while (j > 0 && haystack[i] != needle[j]) {
                j = next[j - 1];
            }
            if (haystack[i] == needle[j]) {
                j++;
            }
            if (j == m) {
                return i - m + 1;
            }
        }
        
        return -1;
    }
};
```

**代码解释**：
1. **边界情况处理**：如果 needle 为空，返回 0；如果 haystack 长度小于 needle，返回 -1
2. **构建 next 数组**：计算每个位置的最长相等前后缀长度
3. **匹配过程**：利用 next 数组，避免重复比较，提高匹配效率
4. **返回结果**：如果找到匹配，返回起始位置；否则返回 -1

### 3. Sunday 算法实现

```cpp
#include <string>
#include <unordered_map>
using namespace std;

class Solution {
public:
    int strStr(string haystack, string needle) {
        int n = haystack.size(), m = needle.size();
        // 边界情况处理
        if (m == 0) return 0;
        if (n < m) return -1;
        
        // 构建偏移表
        unordered_map<char, int> shift;
        for (int i = 0; i < m; i++) {
            shift[needle[i]] = m - i;
        }
        
        // 匹配过程
        int i = 0;
        while (i <= n - m) {
            // 从左到右比较
            int j = 0;
            for (; j < m; j++) {
                if (haystack[i + j] != needle[j]) {
                    break;
                }
            }
            if (j == m) {
                return i;
            }
            // 计算偏移量
            if (i + m < n) {
                char c = haystack[i + m];
                if (shift.count(c)) {
                    i += shift[c];
                } else {
                    i += m + 1;
                }
            } else {
                break;
            }
        }
        
        return -1;
    }
};
```

**代码解释**：
1. **边界情况处理**：如果 needle 为空，返回 0；如果 haystack 长度小于 needle，返回 -1
2. **构建偏移表**：计算每个字符在 needle 中最右侧出现的位置，用于快速移动
3. **匹配过程**：从左到右比较，当匹配失败时，根据偏移表快速移动
4. **返回结果**：如果找到匹配，返回起始位置；否则返回 -1

### 4. Boyer-Moore 算法实现

```cpp
#include <string>
#include <vector>
#include <unordered_map>
using namespace std;

class Solution {
public:
    int strStr(string haystack, string needle) {
        int n = haystack.size(), m = needle.size();
        // 边界情况处理
        if (m == 0) return 0;
        if (n < m) return -1;
        
        // 构建坏字符表
        unordered_map<char, int> bad_char;
        for (int i = 0; i < m; i++) {
            bad_char[needle[i]] = i;
        }
        
        // 构建好后缀表
        vector<int> good_suffix(m, 0);
        vector<int> suffix(m, -1);
        
        // 计算 suffix 数组
        for (int i = m - 2; i >= 0; i--) {
            int j = i;
            int k = 0;
            while (j >= 0 && needle[j] == needle[m - 1 - k]) {
                j--;
                k++;
                suffix[k] = j + 1;
            }
        }
        
        // 计算 good_suffix 数组
        for (int i = 0; i < m; i++) {
            good_suffix[i] = m;
        }
        
        int j = 0;
        for (int i = m - 1; i >= 0; i--) {
            if (suffix[m - i] == -1) {
                while (j < m - i) {
                    if (good_suffix[j] == m) {
                        good_suffix[j] = m - i;
                    }
                    j++;
                }
            } else {
                good_suffix[m - i - 1] = m - i - suffix[m - i];
            }
        }
        
        // 匹配过程
        int i = 0;
        while (i <= n - m) {
            int j = m - 1;
            while (j >= 0 && haystack[i + j] == needle[j]) {
                j--;
            }
            if (j < 0) {
                return i;
            }
            
            // 计算坏字符规则的偏移量
            int bad_char_shift = 1;
            if (bad_char.count(haystack[i + j])) {
                bad_char_shift = j - bad_char[haystack[i + j]];
                if (bad_char_shift < 1) {
                    bad_char_shift = 1;
                }
            }
            
            // 计算好后缀规则的偏移量
            int good_suffix_shift = 1;
            if (j < m - 1) {
                good_suffix_shift = good_suffix[j + 1];
            }
            
            // 取较大的偏移量
            i += max(bad_char_shift, good_suffix_shift);
        }
        
        return -1;
    }
};
```

**代码解释**：
1. **边界情况处理**：如果 needle 为空，返回 0；如果 haystack 长度小于 needle，返回 -1
2. **构建坏字符表**：记录每个字符在 needle 中最右侧出现的位置
3. **构建好后缀表**：计算当匹配失败时，根据好后缀规则的偏移量
4. **匹配过程**：从右到左比较，当匹配失败时，根据坏字符规则和好后缀规则计算偏移量，取较大的偏移量进行移动
5. **返回结果**：如果找到匹配，返回起始位置；否则返回 -1

**Boyer-Moore 算法性能分析**：
- 时间复杂度：O(n+m)，在实际应用中通常比 KMP 算法更快
- 空间复杂度：O(m)，需要存储坏字符表和好后缀表
- 适用场景：在实际应用中性能往往最优，特别是对于长模式串

### 4. 详细步骤解析

以示例 1 为例：`haystack = "sadbutsad", needle = "sad"`

#### 4.1 暴力法步骤解析

**步骤 1**：边界情况处理
- m = 3 != 0，n = 9 >= m = 3，继续

**步骤 2**：遍历 haystack
- i = 0
  - j = 0: haystack[0] == needle[0] → j=1
  - j = 1: haystack[1] == needle[1] → j=2
  - j = 2: haystack[2] == needle[2] → j=3
  - j == m，返回 i = 0

**最终结果**：`0`

#### 4.2 KMP 算法步骤解析

**步骤 1**：边界情况处理
- m = 3 != 0，n = 9 >= m = 3，继续

**步骤 2**：构建 next 数组
- next[0] = 0
- i = 1, j = 0
  - needle[1] != needle[0] → next[1] = 0
- i = 2, j = 0
  - needle[2] != needle[0] → next[2] = 0
- next 数组为 [0, 0, 0]

**步骤 3**：匹配过程
- i = 0, j = 0
  - haystack[0] == needle[0] → i=1, j=1
- i = 1, j = 1
  - haystack[1] == needle[1] → i=2, j=2
- i = 2, j = 2
  - haystack[2] == needle[2] → i=3, j=3
- j == m，返回 i - m + 1 = 0

**最终结果**：`0`

#### 4.3 Sunday 算法步骤解析

**步骤 1**：边界情况处理
- m = 3 != 0，n = 9 >= m = 3，继续

**步骤 2**：构建偏移表
- shift['s'] = 3 - 0 = 3
- shift['a'] = 3 - 1 = 2
- shift['d'] = 3 - 2 = 1

**步骤 3**：匹配过程
- i = 0
  - j = 0: haystack[0] == needle[0] → j=1
  - j = 1: haystack[1] == needle[1] → j=2
  - j = 2: haystack[2] == needle[2] → j=3
  - j == m，返回 i = 0

**最终结果**：`0`

#### 4.4 Boyer-Moore 算法步骤解析

**步骤 1**：边界情况处理
- m = 3 != 0，n = 9 >= m = 3，继续

**步骤 2**：构建坏字符表
- bad_char['s'] = 0
- bad_char['a'] = 1
- bad_char['d'] = 2

**步骤 3**：构建好后缀表
- suffix 数组：[-1, -1, -1]
- good_suffix 数组：[3, 3, 3]

**步骤 4**：匹配过程
- i = 0
  - j = 2: haystack[2] == needle[2] → j=1
  - j = 1: haystack[1] == needle[1] → j=0
  - j = 0: haystack[0] == needle[0] → j=-1
  - j < 0，返回 i = 0

**最终结果**：`0`

---

## 🧪 六、 测试用例设计与验证

### 1. 测试用例

| 测试用例 | 输入 | 输出 | 说明 |
|---------|------|------|------|
| 示例 1 | haystack = "sadbutsad", needle = "sad" | 0 | 普通匹配 |
| 示例 2 | haystack = "leetcode", needle = "leeto" | -1 | 无匹配 |
| 测试用例 3 | haystack = "a", needle = "a" | 0 | 单个字符匹配 |
| 测试用例 4 | haystack = "abc", needle = "" | 0 | 空模式串 |
| 测试用例 5 | haystack = "", needle = "a" | -1 | 空主串 |
| 测试用例 6 | haystack = "hello", needle = "ll" | 2 | 中间位置匹配 |

### 2. 验证过程

以测试用例 6 为例：`haystack = "hello", needle = "ll"`

**步骤 1**：边界情况处理
- m = 2 != 0，n = 5 >= m = 2，继续

**步骤 2**：构建 next 数组
- next[0] = 0
- i = 1, j = 0
  - needle[1] != needle[0] → next[1] = 0
- next 数组为 [0, 0]

**步骤 3**：匹配过程
- i = 0, j = 0
  - haystack[0] != needle[0] → j=0, i=1
- i = 1, j = 0
  - haystack[1] != needle[0] → j=0, i=2
- i = 2, j = 0
  - haystack[2] == needle[0] → i=3, j=1
- i = 3, j = 1
  - haystack[3] == needle[1] → i=4, j=2
- j == m，返回 i - m + 1 = 2

**最终结果**：`2`

---

## ❌ 七、 常见错误与注意事项

### 1. 常见错误

1. **边界情况处理错误**：没有正确处理空字符串、长度不匹配等情况
2. **索引计算错误**：在返回匹配位置时，计算错误
3. **KMP 算法 next 数组构建错误**：没有正确理解最长相等前后缀的概念
4. **暴力法超时**：在长字符串上使用暴力法导致超时
5. **字符比较错误**：在比较字符时，没有考虑大小写或其他特殊情况

### 2. 注意事项

1. **边界情况**：必须处理空字符串、长度不匹配等边界情况
2. **索引计算**：确保返回的是正确的起始位置
3. **性能考虑**：对于长字符串，应使用 KMP 等高效算法
4. **代码可读性**：确保代码清晰易懂，便于维护
5. **测试覆盖**：测试各种边界情况和典型场景

---

## 🚀 八、 相关拓展问题及解题思路

### 1. 相关拓展问题

| 问题 | 描述 | 解题思路 |
|------|------|----------|
| 重复的子字符串 | 判断一个字符串是否由重复的子字符串组成 | 使用 KMP 算法，分析 next 数组 |
| 实现 strStr() | 与本题相同，实现字符串匹配 | 暴力法或 KMP 算法 |
| 最长前缀后缀 | 计算字符串的最长前缀后缀 | 使用 KMP 的 next 数组构建方法 |
| 字符串的排列 | 判断一个字符串是否是另一个字符串的排列 | 滑动窗口 + 哈希表 |
| 最小覆盖子串 | 寻找包含另一个字符串所有字符的最小子串 | 滑动窗口 + 哈希表 |

### 2. 解题思路

#### 2.1 重复的子字符串

**问题描述**：给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。

**解题思路**：
- 使用 KMP 算法构建 next 数组
- 计算字符串长度 n 和最长相等前后缀长度 len = next[n-1]
- 如果 len > 0 且 n % (n - len) == 0，则返回 true
- 否则返回 false

#### 2.2 实现 strStr()

**问题描述**：与本题相同，实现字符串匹配函数。

**解题思路**：
- 暴力法：逐位置比较
- KMP 算法：利用 next 数组提高效率
- Sunday 算法：利用偏移表提高效率

#### 2.3 最长前缀后缀

**问题描述**：计算字符串的最长前缀后缀长度。

**解题思路**：
- 使用 KMP 的 next 数组构建方法
- next 数组的最后一个元素即为整个字符串的最长前缀后缀长度

---

## 💡 九、 面试直击灵魂的防暴击拷问

* **灵魂拷问一：为什么 KMP 算法比暴力法更高效？**
    * **绝杀回答：**"KMP 算法通过预处理模式串得到 next 数组，在匹配失败时，利用已经匹配的信息，避免了重复比较。暴力法在匹配失败时，需要从新的位置开始重新比较，而 KMP 算法可以根据 next 数组直接跳到可能匹配的位置，从而减少了比较次数，时间复杂度从 O(n*m) 降低到 O(n+m)。"

* **灵魂拷问二：如何理解 KMP 算法中的 next 数组？**
    * **绝杀回答：**"next 数组的每个元素 next[i] 表示模式串中前 i+1 个字符的最长相等前后缀长度。最长相等前后缀是指既是前缀又是后缀的最长子串。在匹配失败时，next 数组告诉我们应该回退到模式串的哪个位置，而不是从头开始，从而避免了重复比较。"

* **灵魂拷问三：KMP 算法的空间复杂度是多少？如何优化？**
    * **绝杀回答：**"KMP 算法的空间复杂度是 O(m)，其中 m 是模式串的长度。这是因为需要存储 next 数组。对于模式串很长的情况，可以考虑使用滚动数组或其他优化方法，但通常情况下，O(m) 的空间复杂度是可以接受的。"

* **灵魂拷问四：除了 KMP 算法，还有哪些字符串匹配算法？**
    * **绝杀回答：**"除了 KMP 算法，常见的字符串匹配算法还有：
      1. 暴力法：最简单直接，但时间复杂度高
      2. Sunday 算法：从右向左匹配，利用偏移表加速
      3. Boyer-Moore 算法：结合坏字符规则和好后缀规则，在实际应用中性能往往最优
      4. Rabin-Karp 算法：使用哈希函数，将字符串比较转化为哈希值比较
      每种算法都有其适用场景，选择合适的算法取决于具体问题的特点。"

* **灵魂拷问五：如何处理大字符串的匹配问题？**
    * **绝杀回答：**"处理大字符串的匹配问题，应选择时间复杂度低的算法，如 KMP 算法、Sunday 算法或 Boyer-Moore 算法。此外，还可以考虑以下优化：
      1. 预处理模式串，减少重复计算
      2. 使用多线程或并行处理，提高处理速度
      3. 对于特定场景，如 DNA 序列匹配，可以使用更专门的算法
      4. 利用硬件加速，如 SIMD 指令，提高字符串比较的速度
      总之，选择合适的算法和优化方法，可以有效处理大字符串的匹配问题。"

---

## 📚 十、 总结

本题是一道经典的字符串匹配问题，通过使用 KMP 算法，我们可以在 O(n+m) 的时间复杂度和 O(m) 的空间复杂度下高效地找到字符串中第一个匹配项的下标。

### 核心要点：
1. **KMP 算法**：利用 next 数组避免重复比较，提高匹配效率
2. **边界情况处理**：正确处理空字符串、长度不匹配等情况
3. **算法选择**：根据字符串长度和具体场景选择合适的算法
4. **性能优化**：对于长字符串，应使用高效的字符串匹配算法

### 扩展思考：
- **字符串匹配的应用**：字符串匹配在文本搜索、DNA 序列分析、模式识别等领域有广泛应用
- **算法优化**：通过理解各种字符串匹配算法的原理，可以根据具体问题选择最优算法
- **代码实现**：掌握 KMP 等经典算法的实现，有助于提高编程能力和算法思维

### 方法比较：

| 方法 | 时间复杂度 | 空间复杂度 | 优点 | 缺点 |
|------|------------|------------|------|------|
| 暴力法 | O(n*m) | O(1) | 实现简单，容易理解 | 时间复杂度高，在大字符串上性能差 |
| KMP 算法 | O(n+m) | O(m) | 时间复杂度低，性能稳定 | 实现较复杂，需要理解 next 数组的构建 |
| Sunday 算法 | O(n+m) | O(1) | 实现相对简单，在某些情况下性能优于 KMP | 最坏情况下时间复杂度仍为 O(n*m) |
| Boyer-Moore 算法 | O(n+m) | O(m) | 在实际应用中性能往往最优，特别是对于长模式串 | 实现复杂，理解难度高 |

通过本题的学习，我们不仅掌握了字符串匹配的基本方法，更重要的是理解了 KMP 等高效算法的原理，为解决更复杂的字符串问题打下了基础。