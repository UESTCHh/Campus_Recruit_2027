# 算法实战复盘：两个数组的交集 (LeetCode 349)

> **打卡日期**：2026-04-02
> **核心主题**：HashSet、集合操作、时间复杂度优化。

---

## 📝 一、 题目描述与核心要求

* **题目内容：** 给定两个整数数组 `nums1` 和 `nums2`，返回它们的交集。输出结果中的每个元素必须是唯一的，且顺序不限。
* **核心痛点：** 暴力解法是两层 `for` 循环，时间复杂度 $O(nm)$，其中 $n$ 和 $m$ 分别是两个数组的长度。
* **大厂面试限制：** 要求必须在 $O(n + m)$ 的时间复杂度内解决，且输出结果中的元素必须唯一。

---

## 📊 二、 复杂度深度剖析

### 1. 时间复杂度：$O(n + m)$
**推导过程：** 我们首先将第一个数组的元素存入 HashSet，时间复杂度为 $O(n)$。然后遍历第二个数组，对于每个元素，检查它是否在 HashSet 中，时间复杂度为 $O(m)$。因此整体时间复杂度为 $O(n + m)$。

### 2. 空间复杂度：$O(n)$ 或 $O(m)$
**推导过程：** 我们需要一个 HashSet 来存储第一个数组的元素，空间复杂度为 $O(n)$。如果第二个数组的长度比第一个数组小，我们可以选择存储第二个数组的元素，此时空间复杂度为 $O(m)$。因此空间复杂度为 $O(\min(n, m))$。

---

## 🎯 三、 核心思想：利用 HashSet 去重和快速查询

1. **去重：** 使用 HashSet 存储第一个数组的元素，自动去重，确保每个元素只出现一次。
2. **快速查询：** 遍历第二个数组，对于每个元素，使用 HashSet 的 `contains` 方法（平均时间复杂度 $O(1)$）检查它是否在第一个数组中出现过。
3. **结果收集：** 将同时存在于两个数组中的元素添加到结果集合中，再次使用 HashSet 确保结果中的元素唯一。

---

## 💻 四、 核心代码

```cpp
#include <vector>
#include <unordered_set>

class Solution {
public:
    std::vector<int> intersection(std::vector<int>& nums1, std::vector<int>& nums2) {
        // 1. 将第一个数组的元素存入 HashSet，自动去重
        std::unordered_set<int> set1(nums1.begin(), nums1.end());
        std::unordered_set<int> result_set;
        
        // 2. 遍历第二个数组，检查元素是否在 set1 中
        for (int num : nums2) {
            if (set1.count(num)) {
                result_set.insert(num);
            }
        }
        
        // 3. 将结果集合转换为向量
        return std::vector<int>(result_set.begin(), result_set.end());
    }
};
```

**另一种更简洁的实现：**

```cpp
#include <vector>
#include <unordered_set>

class Solution {
public:
    std::vector<int> intersection(std::vector<int>& nums1, std::vector<int>& nums2) {
        std::unordered_set<int> set1(nums1.begin(), nums1.end());
        std::vector<int> result;
        
        for (int num : nums2) {
            if (set1.erase(num)) { // 如果元素存在，erase 会返回 1，同时删除该元素避免重复
                result.push_back(num);
            }
        }
        
        return result;
    }
};
```
---

## 💡 五、 面试直击灵魂的防暴击拷问
* **灵魂拷问一：为什么使用 HashSet 而不是数组？**
    * **绝杀回答：**“HashSet 的查询时间复杂度是 $O(1)$，而数组的查询时间复杂度是 $O(n)$。使用 HashSet 可以将时间复杂度从 $O(nm)$ 降到 $O(n + m)$，显著提高效率。此外，HashSet 还能自动去重，确保结果中的元素唯一。”

* **灵魂拷问二：如果数组中的元素范围很小，比如在 0-1000 之间，有没有更高效的方法？**
    * **绝杀回答：**“如果元素范围很小，可以使用布尔数组来代替 HashSet。例如，创建一个大小为 1001 的布尔数组，标记第一个数组中出现的元素，然后遍历第二个数组检查对应位置是否为 true。这种方法的时间复杂度仍然是 $O(n + m)$，但空间复杂度可能更低，且访问速度更快。”

* **灵魂拷问三：如何处理大数据量的情况？**
    * **绝杀回答：**“对于大数据量的情况，HashSet 仍然是合适的选择，因为它的时间复杂度是线性的。但需要注意内存使用，如果其中一个数组非常大，我们应该选择存储较小的那个数组，以减少内存消耗。此外，在 C++ 中，`std::unordered_set` 的实现通常会自动处理哈希冲突，确保在大数据量下仍然保持良好的性能。”

* **灵魂拷问四：输出结果的顺序有要求吗？**
    * **绝杀回答：**“题目要求输出结果的顺序不限，因此我们不需要对结果进行排序。如果需要按特定顺序输出，比如升序，我们可以在收集完结果后对向量进行排序，时间复杂度会增加到 $O((n + m) \log (n + m))$。”
---