## KMP 算法学习笔记

> **打卡日期**：2026-04-11
> **核心主题**：KMP 算法原理、实现与应用
> **资料出处**：
> - https://www.programmercarl.com/0028.%E5%AE%9E%E7%8E%B0strStr.html#%E6%80%9D%E8%B7%AF
> - https://www.bilibili.com/video/BV1M5411j7Xx
> - https://www.bilibili.com/video/BV1PD4y1o7nd
> - https://www.bilibili.com/video/BV1K7DnBFEMe
> - https://www.bilibili.com/video/BV1AY4y157yL
> - https://www.bilibili.com/video/BV1jb411V78H

---

## 1. KMP 算法简介

### 1.1 什么是 KMP 算法

KMP 算法是一种高效的字符串匹配算法，由 Knuth、Morris 和 Pratt 三位学者发明，因此得名 KMP。它的核心思想是利用已经匹配的信息，避免模式串的重复比较，从而提高字符串匹配的效率。

### 1.2 应用场景

KMP 算法主要应用于以下场景：
- 字符串匹配：在一个文本串中查找一个模式串的出现位置
- 模式识别：在生物序列分析、文本编辑等领域识别特定模式
- 编译器设计：词法分析中的模式匹配
- 网络安全：入侵检测系统中的特征匹配

### 1.3 KMP 算法的优势

与朴素字符串匹配算法相比，KMP 算法具有以下优势：
- **时间复杂度**：KMP 算法的时间复杂度为 O(m + n)，其中 m 是模式串长度，n 是文本串长度，而朴素算法的时间复杂度为 O(m * n)
- **避免重复比较**：当匹配失败时，KMP 算法利用已匹配的信息，避免从头开始匹配
- **空间效率**：虽然需要额外的空间来存储 next 数组，但空间复杂度仅为 O(m)，对于大多数应用场景来说是可以接受的

---

## 2. 核心原理

### 2.1 朴素字符串匹配的问题

朴素字符串匹配算法在每次匹配失败时，都会将模式串向右移动一位，然后从头开始匹配。这种方法在模式串和文本串有很多重复前缀时会做很多无用功。

例如，当文本串是 "AAAAAAB"，模式串是 "AAAAAB" 时，朴素算法会在每次匹配失败后重新开始，而 KMP 算法则会利用已匹配的信息，直接跳过不必要的比较。

### 2.2 KMP 的核心思想

KMP 算法的核心思想是：当出现字符串不匹配时，可以利用已经匹配的部分信息，避免模式串的重复比较。

具体来说，KMP 算法通过构建一个前缀表（next 数组），记录模式串中每个位置的最长相等前后缀长度，从而在匹配失败时，知道模式串应该回退到哪个位置继续匹配。

### 2.3 前缀和后缀的定义

- **前缀**：指不包含最后一个字符的所有以第一个字符开头的连续子串
- **后缀**：指不包含第一个字符的所有以最后一个字符结尾的连续子串

例如，对于字符串 "ABABC"：
- 前缀包括："A", "AB", "ABA", "ABAB"
- 后缀包括："B", "CB", "ABC", "BABC"

### 2.4 最长相等前后缀

最长相等前后缀是指字符串中相同的前缀和后缀的最大长度。

例如，对于字符串 "ABAB"：
- 前缀："A", "AB", "ABA"
- 后缀："B", "AB", "BAB"
- 最长相等前后缀是 "AB"，长度为 2

---

## 3. 前缀表与 next 数组

### 3.1 前缀表的作用

前缀表（也称为 next 数组）的作用是：当模式串与文本串匹配失败时，记录模式串应该回退到哪个位置继续匹配，而不是从头开始。

### 3.2 next 数组的定义

next 数组的每个元素 next[i] 表示：模式串中以 i 结尾的子串的最长相等前后缀的长度。

### 3.3 构建 next 数组的示例

以模式串 "ABAB" 为例，构建 next 数组：

| 索引 i | 子串 | 最长相等前后缀 | next[i] |
|--------|------|----------------|---------|
| 0      | "A"  | 无             | 0       |
| 1      | "AB" | 无             | 0       |
| 2      | "ABA" | "A"            | 1       |
| 3      | "ABAB" | "AB"           | 2       |

以模式串 "AABAAF" 为例，构建 next 数组：

| 索引 i | 子串 | 最长相等前后缀 | next[i] |
|--------|------|----------------|---------|
| 0      | "A"  | 无             | 0       |
| 1      | "AA" | "A"            | 1       |
| 2      | "AAB" | 无             | 0       |
| 3      | "AABA" | "A"            | 1       |
| 4      | "AABAA" | "AA"           | 2       |
| 5      | "AABAAF" | 无             | 0       |

### 3.4 前缀表的不同表示方式

前缀表有两种常见的表示方式：
1. **原始前缀表**：直接存储最长相等前后缀的长度
2. **前缀表统一减一**：将原始前缀表中的每个值减一，这样在回退时可以直接使用 next[i] 作为索引

---

## 4. 构建 next 数组的步骤

### 4.1 初始化

- 初始化两个指针：j 指向已匹配的前缀末尾，i 指向当前处理的位置
- 初始化 next 数组，next[0] = 0

### 4.2 处理不匹配的情况

当 s[i] != s[j] 时，需要回退 j 到 next[j-1] 的位置，直到 j == 0 或者 s[i] == s[j]

### 4.3 处理匹配的情况

当 s[i] == s[j] 时，j 增加 1，然后 next[i] = j

### 4.4 完整的 next 数组构建代码

```cpp
// 构建原始前缀表（不减一）
void getNext(int* next, const string& s) {
    int j = 0; // j 指向前缀的末尾
    next[0] = 0;
    for (int i = 1; i < s.size(); i++) { // i 指向后缀的末尾
        // 处理不匹配的情况
        while (j > 0 && s[i] != s[j]) {
            j = next[j - 1]; // 回退 j
        }
        // 处理匹配的情况
        if (s[i] == s[j]) {
            j++;
        }
        next[i] = j;
    }
}

// 构建前缀表（统一减一）
void getNext(int* next, const string& s) {
    int j = -1;
    next[0] = j;
    for (int i = 1; i < s.size(); i++) {
        while (j >= 0 && s[i] != s[j + 1]) {
            j = next[j];
        }
        if (s[i] == s[j + 1]) {
            j++;
        }
        next[i] = j;
    }
}
```

### 4.5 构建 next 数组的详细示例

以模式串 "AABAAF" 为例，构建 next 数组（原始前缀表）：

1. 初始化：j = 0, next[0] = 0
2. i = 1：s[1] = 'A', s[j] = 'A'，匹配，j = 1，next[1] = 1
3. i = 2：s[2] = 'B', s[j] = 'A'，不匹配，j = next[j-1] = next[0] = 0，s[2] != s[0]，j 保持 0，next[2] = 0
4. i = 3：s[3] = 'A', s[j] = 'A'，匹配，j = 1，next[3] = 1
5. i = 4：s[4] = 'A', s[j] = 'A'，匹配，j = 2，next[4] = 2
6. i = 5：s[5] = 'F', s[j] = 'B'，不匹配，j = next[j-1] = next[1] = 1，s[5] != s[1]，j = next[j-1] = next[0] = 0，s[5] != s[0]，j 保持 0，next[5] = 0

最终 next 数组为：[0, 1, 0, 1, 2, 0]

---

## 5. KMP 算法的匹配过程

### 5.1 匹配步骤

1. 初始化两个指针：i 指向文本串的当前位置，j 指向模式串的当前位置
2. 当 i < 文本串长度且 j < 模式串长度时：
   - 如果 s[i] == p[j]，则 i 和 j 都增加 1
   - 如果 s[i] != p[j]：
     - 如果 j > 0，则 j = next[j-1]（使用原始前缀表）
     - 如果 j == 0，则 i 增加 1
3. 当 j == 模式串长度时，匹配成功，返回 i - j
4. 如果匹配失败，返回 -1

### 5.2 匹配过程的详细示例

以文本串 "AABAABAFA" 和模式串 "AABAAF" 为例：

1. 初始化：i = 0, j = 0
2. s[0] = 'A', p[0] = 'A'，匹配，i = 1, j = 1
3. s[1] = 'A', p[1] = 'A'，匹配，i = 2, j = 2
4. s[2] = 'B', p[2] = 'B'，匹配，i = 3, j = 3
5. s[3] = 'A', p[3] = 'A'，匹配，i = 4, j = 4
6. s[4] = 'A', p[4] = 'A'，匹配，i = 5, j = 5
7. s[5] = 'B', p[5] = 'F'，不匹配
8. j = next[j-1] = next[4] = 2
9. s[5] = 'B', p[2] = 'B'，匹配，i = 6, j = 3
10. s[6] = 'A', p[3] = 'A'，匹配，i = 7, j = 4
11. s[7] = 'F', p[4] = 'A'，不匹配
12. j = next[j-1] = next[3] = 1
13. s[7] = 'F', p[1] = 'A'，不匹配
14. j = next[j-1] = next[0] = 0
15. s[7] = 'F', p[0] = 'A'，不匹配，i = 8
16. i == 文本串长度，匹配失败，返回 -1

### 5.3 匹配过程的可视化流程图

```
文本串: A A B A A B A F A
       0 1 2 3 4 5 6 7 8
模式串: A A B A A F
       0 1 2 3 4 5
next数组: [0, 1, 0, 1, 2, 0]

步骤 1: i=0, j=0 → 匹配 → i=1, j=1
步骤 2: i=1, j=1 → 匹配 → i=2, j=2
步骤 3: i=2, j=2 → 匹配 → i=3, j=3
步骤 4: i=3, j=3 → 匹配 → i=4, j=4
步骤 5: i=4, j=4 → 匹配 → i=5, j=5
步骤 6: i=5, j=5 → 不匹配 → j=next[4]=2
步骤 7: i=5, j=2 → 匹配 → i=6, j=3
步骤 8: i=6, j=3 → 匹配 → i=7, j=4
步骤 9: i=7, j=4 → 不匹配 → j=next[3]=1
步骤10: i=7, j=1 → 不匹配 → j=next[0]=0
步骤11: i=7, j=0 → 不匹配 → i=8
步骤12: 匹配失败，返回 -1
```

---

## 6. 完整代码实现

### 6.1 使用原始前缀表的实现

```cpp
#include <iostream>
#include <string>
#include <vector>

using namespace std;

// 构建原始前缀表（不减一）
void getNext(vector<int>& next, const string& s) {
    int j = 0; // j 指向前缀的末尾
    next[0] = 0;
    for (int i = 1; i < s.size(); i++) { // i 指向后缀的末尾
        // 处理不匹配的情况
        while (j > 0 && s[i] != s[j]) {
            j = next[j - 1]; // 回退 j
        }
        // 处理匹配的情况
        if (s[i] == s[j]) {
            j++;
        }
        next[i] = j;
    }
}

// KMP 匹配算法
int strStr(string haystack, string needle) {
    if (needle.empty()) {
        return 0;
    }
    
    vector<int> next(needle.size());
    getNext(next, needle);
    
    int j = 0; // j 指向模式串的当前位置
    for (int i = 0; i < haystack.size(); i++) { // i 指向文本串的当前位置
        // 处理不匹配的情况
        while (j > 0 && haystack[i] != needle[j]) {
            j = next[j - 1]; // 回退 j
        }
        // 处理匹配的情况
        if (haystack[i] == needle[j]) {
            j++;
        }
        // 匹配成功
        if (j == needle.size()) {
            return i - j + 1;
        }
    }
    
    return -1;
}

int main() {
    string haystack = "AABAABAFA";
    string needle = "AABAAF";
    int result = strStr(haystack, needle);
    cout << "Result: " << result << endl; // 输出: -1
    
    haystack = "hello";
    needle = "ll";
    result = strStr(haystack, needle);
    cout << "Result: " << result << endl; // 输出: 2
    
    return 0;
}
```

### 6.2 使用前缀表统一减一的实现

```cpp
#include <iostream>
#include <string>
#include <vector>

using namespace std;

// 构建前缀表（统一减一）
void getNext(vector<int>& next, const string& s) {
    int j = -1;
    next[0] = j;
    for (int i = 1; i < s.size(); i++) {
        // 处理不匹配的情况
        while (j >= 0 && s[i] != s[j + 1]) {
            j = next[j];
        }
        // 处理匹配的情况
        if (s[i] == s[j + 1]) {
            j++;
        }
        next[i] = j;
    }
}

// KMP 匹配算法
int strStr(string haystack, string needle) {
    if (needle.empty()) {
        return 0;
    }
    
    vector<int> next(needle.size());
    getNext(next, needle);
    
    int j = -1; // j 指向模式串的当前位置
    for (int i = 0; i < haystack.size(); i++) { // i 指向文本串的当前位置
        // 处理不匹配的情况
        while (j >= 0 && haystack[i] != needle[j + 1]) {
            j = next[j]; // 回退 j
        }
        // 处理匹配的情况
        if (haystack[i] == needle[j + 1]) {
            j++;
        }
        // 匹配成功
        if (j == needle.size() - 1) {
            return i - j;
        }
    }
    
    return -1;
}

int main() {
    string haystack = "AABAABAFA";
    string needle = "AABAAF";
    int result = strStr(haystack, needle);
    cout << "Result: " << result << endl; // 输出: -1
    
    haystack = "hello";
    needle = "ll";
    result = strStr(haystack, needle);
    cout << "Result: " << result << endl; // 输出: 2
    
    return 0;
}
```

---

## 7. 常见问题与解决方案

### 7.1 如何处理空模式串

当模式串为空时，根据题目要求，应该返回 0。这与 C 语言的 `strstr()` 和 Java 的 `indexOf()` 定义相符。

### 7.2 如何处理模式串长度大于文本串长度的情况

当模式串长度大于文本串长度时，不可能匹配成功，直接返回 -1。

### 7.3 如何理解 next 数组的回退机制

next 数组的回退机制是 KMP 算法的核心。当匹配失败时，next 数组告诉我们模式串应该回退到哪个位置继续匹配，而不是从头开始。这是因为 next 数组记录了模式串中每个位置的最长相等前后缀长度，利用这个信息可以避免重复比较。

### 7.4 如何选择前缀表的表示方式

两种前缀表表示方式各有优缺点：
- **原始前缀表**：更直观，容易理解
- **前缀表统一减一**：在代码实现中更简洁，不需要额外的边界检查

选择哪种方式取决于个人偏好和具体应用场景。

---

## 8. 练习题与解答

### 8.1 练习 1：LeetCode 28. 实现 strStr()

**题目**：实现 `strStr()` 函数。给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置（从 0 开始）。如果不存在，则返回 -1。

**示例 1**：
输入: haystack = "hello", needle = "ll"
输出: 2

**示例 2**：
输入: haystack = "aaaaa", needle = "bba"
输出: -1

**解答**：使用上面的 KMP 算法实现即可。

### 8.2 练习 2：构建 next 数组

**题目**：为模式串 "ABABABC" 构建 next 数组（原始前缀表）。

**解答**：

| 索引 i | 子串 | 最长相等前后缀 | next[i] |
|--------|------|----------------|---------|
| 0      | "A"  | 无             | 0       |
| 1      | "AB" | 无             | 0       |
| 2      | "ABA" | "A"            | 1       |
| 3      | "ABAB" | "AB"           | 2       |
| 4      | "ABABA" | "ABA"          | 3       |
| 5      | "ABABAB" | "ABAB"         | 4       |
| 6      | "ABABABC" | 无             | 0       |

最终 next 数组为：[0, 0, 1, 2, 3, 4, 0]

### 8.3 练习 3：模拟 KMP 匹配过程

**题目**：使用 KMP 算法在文本串 "ABABABABCABABABABC" 中查找模式串 "ABABABC"，模拟匹配过程。

**解答**：
1. 构建模式串 "ABABABC" 的 next 数组：[0, 0, 1, 2, 3, 4, 0]
2. 匹配过程：
   - i=0, j=0 → 匹配 → i=1, j=1
   - i=1, j=1 → 匹配 → i=2, j=2
   - i=2, j=2 → 匹配 → i=3, j=3
   - i=3, j=3 → 匹配 → i=4, j=4
   - i=4, j=4 → 匹配 → i=5, j=5
   - i=5, j=5 → 匹配 → i=6, j=6
   - i=6, j=6 → 不匹配 → j=next[5]=4
   - i=6, j=4 → s[6]='A' != p[4]='A'？不，s[6]='A'，p[4]='A'，匹配 → i=7, j=5
   - i=7, j=5 → s[7]='B' == p[5]='B'，匹配 → i=8, j=6
   - i=8, j=6 → s[8]='C' == p[6]='C'，匹配，j=7 == 模式串长度，匹配成功，返回 i-j+1=8-7+1=2

---

## 9. 总结

KMP 算法是一种高效的字符串匹配算法，其核心在于利用前缀表（next 数组）来避免模式串的重复比较。通过理解 KMP 算法的原理和实现，我们可以：

1. **提高字符串匹配的效率**：时间复杂度从 O(m * n) 降低到 O(m + n)
2. **理解算法设计的思想**：利用已有的信息来优化算法性能
3. **解决实际应用中的字符串匹配问题**：如文本搜索、模式识别等

KMP 算法的难点在于理解 next 数组的构建和回退机制，但通过详细的例子和步骤分解，我们可以逐步掌握其核心原理。希望本笔记能够帮助初学者理解 KMP 算法，并在实际应用中灵活使用。

---

*本笔记基于代码随想录和相关视频教程整理，结合个人理解和实践经验。*