# 侯捷 C++ STL标准库和泛型编程 - 容器结构与Rb_tree深度探索
> **打卡日期**：2026-04-09
 > **核心主题**：STL容器结构与分类、Rb_tree实现原理与应用。
 > **资料出处**：https://www.bilibili.com/video/BV1kBh5zrEYh?spm_id_from=333.788.videopod.sections&vd_source=e01209adaee5cf8495fa73fbfde26dd1

---
## 一、STL容器结构与分类

### 1. 容器分类体系

**重要程度：基础**

STL容器分为两大类：序列式容器（Sequence Containers）和关联式容器（Associative Containers）。

#### 序列式容器

| 容器 | 特点 | 内存布局 | 引入时间 |
|------|------|----------|----------|
| array | 固定大小 | 连续空间 | C++11 |
| vector | 动态大小 | 连续空间 | C++98 |
| list | 双向链表 | 非连续空间 | C++98 |
| forward_list (slist) | 单向链表 | 非连续空间 | C++11 (非标准) |
| deque | 分段连续空间 | 分段连续 | C++98 |
| stack | 容器适配器 | 基于deque | C++98 |
| queue | 容器适配器 | 基于deque | C++98 |
| priority_queue | 容器适配器 | 基于vector (heap) | C++98 |

#### 关联式容器

| 容器 | 底层实现 | 特点 |
|------|----------|------|
| set | rb_tree | 键唯一，有序 |
| map | rb_tree | 键值对，键唯一，有序 |
| multiset | rb_tree | 键可重复，有序 |
| multimap | rb_tree | 键值对，键可重复，有序 |
| unordered_set | hashtable | 键唯一，无序 |
| unordered_map | hashtable | 键值对，键唯一，无序 |
| unordered_multiset | hashtable | 键可重复，无序 |
| unordered_multimap | hashtable | 键值对，键可重复，无序 |

### 2. 容器大小对比（G2.9 vs G4.9）

**重要程度：进阶**

| 容器 | G2.9大小 | G4.9大小 | 说明 |
|------|----------|----------|------|
| array | 12 | 12 | 固定大小，与元素类型无关 |
| vector | 12 | 24 | 包含三个指针/迭代器 |
| list | 4 | 8 | 包含一个指针/两个指针 |
| forward_list | 4 | 8 | 包含一个指针 |
| deque | 40 | 40 | 包含多个指针和大小信息 |
| stack | 40 | 40 | 基于deque |
| queue | 40 | 40 | 基于deque |
| set | 12 | 24 | 基于rb_tree |
| map | 12 | 24 | 基于rb_tree |
| multiset | 12 | 24 | 基于rb_tree |
| multimap | 12 | 24 | 基于rb_tree |
| hashtable | 20 | - | 非公开实现 |
| unordered_set | 20 | 28 | 基于hashtable |
| unordered_map | 20 | 28 | 基于hashtable |

### 3. 容器间的关系

**重要程度：进阶**

- **组合关系**：关联式容器（如set、map）内部使用rb_tree作为底层实现
- **容器适配器**：stack、queue、priority_queue是对其他容器的包装
- **底层实现**：
  - 序列式容器：直接实现或基于其他容器
  - 关联式容器：基于rb_tree或hashtable

### 4. 面试高频考点解析

**重要程度：高级**

#### 1. 问题：STL容器分为哪几类？各自的特点是什么？

**回答思路**：
- 分类标准
- 各类容器的特点
- 内存布局差异
- 适用场景

**参考答案**：
- STL容器分为序列式容器和关联式容器两大类。
- 序列式容器：
  - array：固定大小，连续空间，适合大小已知的场景
  - vector：动态大小，连续空间，适合随机访问频繁的场景
  - list：双向链表，适合频繁插入删除的场景
  - forward_list：单向链表，空间开销更小
  - deque：分段连续，适合两端操作频繁的场景
  - stack/queue：容器适配器，提供特定接口
- 关联式容器：
  - set/map：基于rb_tree，有序，键唯一
  - multiset/multimap：基于rb_tree，有序，键可重复
  - unordered_set/unordered_map：基于hashtable，无序，键唯一
  - unordered_multiset/unordered_multimap：基于hashtable，无序，键可重复

#### 2. 问题：vector和deque的区别是什么？

**回答思路**：
- 内存布局
- 两端操作性能
- 中间插入删除性能
- 随机访问性能
- 适用场景

**参考答案**：
- 内存布局：vector是连续内存，deque是分段连续内存
- 两端操作：vector的push_front为O(n)，deque的push_front为O(1)
- 中间插入删除：两者均为O(n)，但deque实际性能可能更好
- 随机访问：两者均为O(1)，但vector访问速度更快
- 适用场景：vector适合频繁随机访问，deque适合频繁两端操作

#### 3. 问题：set和unordered_set的区别是什么？

**回答思路**：
- 底层实现
- 有序性
- 插入查找删除性能
- 内存开销
- 适用场景

**参考答案**：
- 底层实现：set基于rb_tree，unordered_set基于hashtable
- 有序性：set是有序的，unordered_set是无序的
- 性能：
  - set：插入、查找、删除均为O(log n)
  - unordered_set：平均为O(1)，最坏为O(n)
- 内存开销：unordered_set更大，需要哈希表结构
- 适用场景：set适合需要有序遍历的场景，unordered_set适合需要快速查找的场景

### 5. 核心知识点强化

**重要程度：进阶**

#### 1. 容器适配器

容器适配器是对现有容器的包装，提供特定的接口：

- **stack**：后进先出（LIFO），默认基于deque
- **queue**：先进先出（FIFO），默认基于deque
- **priority_queue**：优先级队列，默认基于vector（使用heap算法）

#### 2. 容器的选择依据

选择容器时应考虑以下因素：

- **数据访问模式**：随机访问 vs 顺序访问
- **插入删除频率**：频繁插入删除 vs 较少修改
- **内存开销**：空间限制
- **性能要求**：时间复杂度需求
- **是否需要有序**：有序 vs 无序

### 6. 实践应用指南

**重要程度：进阶**

#### 1. 实际开发场景中的应用

**场景1：需要频繁随机访问**
- **应用**：存储大量数据并需要快速访问
- **最佳实践**：使用vector

**场景2：需要频繁插入删除**
- **应用**：实现链表结构、频繁修改的数据集
- **最佳实践**：使用list或forward_list

**场景3：需要两端操作**
- **应用**：实现队列、滑动窗口
- **最佳实践**：使用deque

**场景4：需要有序数据**
- **应用**：需要排序的数据、范围查询
- **最佳实践**：使用set或map

**场景5：需要快速查找**
- **应用**：字典、缓存
- **最佳实践**：使用unordered_set或unordered_map

### 7. 知识关联图谱

**重要程度：进阶**

```
STL容器
├── 序列式容器
│   ├── 连续空间
│   │   ├── array (固定大小)
│   │   └── vector (动态大小)
│   ├── 链表结构
│   │   ├── list (双向)
│   │   └── forward_list (单向)
│   ├── 分段连续
│   │   └── deque
│   └── 容器适配器
│       ├── stack (LIFO)
│       ├── queue (FIFO)
│       └── priority_queue (优先级)
└── 关联式容器
    ├── 有序
    │   ├── set (键唯一)
    │   ├── map (键值对，键唯一)
    │   ├── multiset (键可重复)
    │   └── multimap (键值对，键可重复)
    └── 无序
        ├── unordered_set (键唯一)
        ├── unordered_map (键值对，键唯一)
        ├── unordered_multiset (键可重复)
        └── unordered_multimap (键值对，键可重复)
```

### 8. 前沿技术拓展

**重要程度：高级**

#### 1. C++11/14/17/20中的容器改进

- **C++11**：
  - 引入forward_list（替换slist）
  - 引入unordered容器系列
  - 引入移动语义和右值引用

- **C++14**：
  - 增加make_shared等辅助函数
  - 改进哈希表实现

- **C++17**：
  - 引入std::optional
  - 引入std::any
  - 改进容器的emplace操作

- **C++20**：
  - 引入concepts，改进容器的模板约束
  - 引入span，提供非拥有的视图
  - 引入ranges库，提供更灵活的容器操作

#### 2. 性能优化趋势

- **内存池**：使用自定义分配器减少内存分配开销
- **小对象优化**：针对小对象的特殊处理
- **并行处理**：与并行算法库配合使用
- **编译期优化**：利用constexpr和模板元编程

## 二、Rb_tree深度探索

### 1. Rb_tree的基本概念

**重要程度：基础**

Rb_tree（红黑树）是一种平衡二叉搜索树，具有以下特性：

- **平衡性质**：任何节点的左右子树高度差不超过1
- **红黑性质**：
  - 每个节点要么是红色，要么是黑色
  - 根节点是黑色
  - 每个叶子节点（NIL）是黑色
  - 如果一个节点是红色，则其两个子节点都是黑色
  - 从任一节点到其每个叶子节点的所有路径都包含相同数目的黑色节点

### 2. Rb_tree的内部结构

**重要程度：进阶**

#### G4.9中的Rb_tree结构

```cpp
template<typename _Key, typename _Value, typename _KeyOfValue,
         typename _Compare, typename _Alloc = allocator<_Value>>
class _Rb_tree
{
protected:
    struct _Rb_tree_impl
    {
        _Rb_tree_node_base* _M_node;  // 根节点
        size_type _M_node_count;  // 节点数量
        _Compare _M_key_compare;  // 键比较函数
        // ...
    };
    
    _Rb_tree_impl _M_impl;
    
public:
    // 插入操作
    pair<iterator, bool> _M_insert_unique(const value_type& __v);
    iterator _M_insert_equal(const value_type& __v);
    // ...
};
```

#### 节点结构

```cpp
struct _Rb_tree_node_base
{
    typedef _Rb_tree_node_base* _Base_ptr;
    typedef _Rb_tree_color_type _Color_type;
    
    _Color_type _M_color;  // 节点颜色
    _Base_ptr _M_parent;  // 父节点指针
    _Base_ptr _M_left;  // 左子节点指针
    _Base_ptr _M_right;  // 右子节点指针
    
    static _Base_ptr minimum(_Base_ptr __x);
    static _Base_ptr maximum(_Base_ptr __x);
};

template<typename _Value>
struct _Rb_tree_node : public _Rb_tree_node_base
{
    typedef _Rb_tree_node<_Value>* _Link_type;
    _Value _M_value_field;  // 节点值
};
```

### 3. Rb_tree的迭代器

**重要程度：进阶**

```cpp
template<typename _Tp>
struct _Rb_tree_iterator
{
    typedef _Rb_tree_node<_Tp> _Node;
    typedef _Rb_tree_iterator<_Tp> _Self;
    
    _Node* _M_node;  // 当前节点指针
    
    reference operator*() const { return _M_node->_M_value_field; }
    pointer operator->() const { return &(operator*()); }
    
    _Self& operator++();  // 前序遍历
    _Self operator++(int);
    _Self& operator--();  // 后序遍历
    _Self operator--(int);
    // ...
};
```

### 4. Rb_tree的插入操作

**重要程度：高级**

Rb_tree提供两种插入操作：

- **_M_insert_unique**：插入唯一键，若键已存在则插入失败
- **_M_insert_equal**：插入可重复键，允许键重复

#### 插入流程

1. 寻找插入位置
2. 创建新节点
3. 插入节点
4. 调整红黑树平衡

### 5. Rb_tree的应用

**重要程度：进阶**

Rb_tree是set、map、multiset、multimap的底层实现：

- **set**：使用Rb_tree，键即值
- **map**：使用Rb_tree，值为键值对
- **multiset**：使用Rb_tree，允许键重复
- **multimap**：使用Rb_tree，允许键重复，值为键值对

### 6. 面试高频考点解析

**重要程度：高级**

#### 1. 问题：Rb_tree的特性是什么？

**回答思路**：
- 平衡二叉搜索树的特性
- 红黑树的特殊性质
- 与AVL树的对比

**参考答案**：
- Rb_tree是一种平衡二叉搜索树，具有以下特性：
  - 平衡性质：任何节点的左右子树高度差不超过1
  - 红黑性质：
    - 每个节点要么是红色，要么是黑色
    - 根节点是黑色
    - 每个叶子节点（NIL）是黑色
    - 如果一个节点是红色，则其两个子节点都是黑色
    - 从任一节点到其每个叶子节点的所有路径都包含相同数目的黑色节点
- 与AVL树相比，Rb_tree的平衡要求更宽松，插入和删除操作的开销更小

#### 2. 问题：Rb_tree如何实现set和map？

**回答思路**：
- set的实现原理
- map的实现原理
- 键的处理方式
- 迭代器的实现

**参考答案**：
- set的实现：
  - 使用Rb_tree，键即值
  - 使用identity函数作为KeyOfValue
  - 使用insert_unique确保键唯一
- map的实现：
  - 使用Rb_tree，值为键值对(pair<const Key, T>)
  - 使用select1st函数作为KeyOfValue，提取键
  - 使用insert_unique确保键唯一
- 两者都利用Rb_tree的有序性和平衡特性，提供O(log n)的插入、查找、删除操作

#### 3. 问题：Rb_tree的插入操作有哪些？

**回答思路**：
- 两种插入操作的区别
- 插入流程
- 时间复杂度
- 应用场景

**参考答案**：
- Rb_tree提供两种插入操作：
  - _M_insert_unique：插入唯一键，若键已存在则插入失败
  - _M_insert_equal：插入可重复键，允许键重复
- 插入流程：
  1. 寻找插入位置
  2. 创建新节点
  3. 插入节点
  4. 调整红黑树平衡
- 时间复杂度：O(log n)
- 应用场景：
  - _M_insert_unique：用于set和map
  - _M_insert_equal：用于multiset和multimap

### 7. 核心知识点强化

**重要程度：进阶**

#### 1. Rb_tree的性能特性

- **时间复杂度**：
  - 插入：O(log n)
  - 查找：O(log n)
  - 删除：O(log n)
  - 遍历：O(n)

- **空间复杂度**：O(n)，每个节点需要额外存储颜色和三个指针

#### 2. Rb_tree的优缺点

**优点**：
- 保持适度平衡，避免树过深
- 插入和删除操作的开销相对较小
- 提供有序遍历
- 适用于需要有序数据的场景

**缺点**：
- 实现复杂
- 内存开销较大
- 对于频繁插入删除的场景，性能不如哈希表

### 8. 实践应用指南

**重要程度：进阶**

#### 1. 实际开发场景中的应用

**场景1：需要有序数据**
- **应用**：需要按顺序遍历的数据集合
- **最佳实践**：使用基于Rb_tree的set或map

**场景2：需要范围查询**
- **应用**：查找某个范围内的元素
- **最佳实践**：使用基于Rb_tree的set或map，利用其有序性

**场景3：需要键值映射**
- **应用**：字典、配置管理
- **最佳实践**：使用map，提供键到值的映射

### 9. 知识关联图谱

**重要程度：进阶**

```
Rb_tree
├── 基本特性
│   ├── 平衡二叉搜索树
│   ├── 红黑性质
│   └── 适度平衡
├── 内部结构
│   ├── _Rb_tree_impl
│   ├── _Rb_tree_node_base
│   ├── _Rb_tree_node
│   └── _Rb_tree_iterator
├── 核心操作
│   ├── _M_insert_unique
│   ├── _M_insert_equal
│   ├── find
│   └── erase
├── 应用
│   ├── set
│   ├── map
│   ├── multiset
│   └── multimap
└── 性能特性
    ├── 时间复杂度：O(log n)
    ├── 空间复杂度：O(n)
    └── 适用场景
```

### 10. 前沿技术拓展

**重要程度：高级**

#### 1. C++11/14/17/20中的Rb_tree改进

- **C++11**：
  - 引入移动语义，优化节点插入操作
  - 改进迭代器实现

- **C++14**：
  - 增加emplace操作，直接在树中构造元素
  - 改进异常安全性

- **C++17**：
  - 增加try_emplace操作
  - 改进哈希表与Rb_tree的接口一致性

- **C++20**：
  - 增加constexpr支持
  - 与ranges库集成

#### 2. 性能优化趋势

- **内存池**：使用自定义分配器减少节点内存分配开销
- **批量操作**：优化批量插入和删除操作
- **并行处理**：探索并行化的Rb_tree实现
- **硬件加速**：利用现代CPU的特性优化树操作

#### 3. 相关技术发展

- **无锁红黑树**：提高并发性能
- **持久化红黑树**：支持版本控制
- **区间树**：基于红黑树的扩展，支持区间查询
- **Treap**：结合树和堆的特性，简化实现

## 三、代码示例与实践

### 1. Rb_tree使用示例

```cpp
#include <iostream>
#include <set>
#include <map>

int main() {
    // 使用set（基于Rb_tree）
    std::set<int> s;
    s.insert(3);
    s.insert(8);
    s.insert(5);
    s.insert(9);
    s.insert(13);
    s.insert(5);  // 重复键，不会插入
    
    std::cout << "set size: " << s.size() << std::endl;  // 5
    std::cout << "count(5): " << s.count(5) << std::endl;  // 1
    
    // 遍历set
    std::cout << "set elements: ";
    for (auto it = s.begin(); it != s.end(); ++it) {
        std::cout << *it << " ";  // 3 5 8 9 13
    }
    std::cout << std::endl;
    
    // 使用map（基于Rb_tree）
    std::map<std::string, int> m;
    m["apple"] = 10;
    m["banana"] = 20;
    m["orange"] = 15;
    
    std::cout << "map size: " << m.size() << std::endl;  // 3
    std::cout << "apple: " << m["apple"] << std::endl;  // 10
    
    // 遍历map
    std::cout << "map elements: ";
    for (auto it = m.begin(); it != m.end(); ++it) {
        std::cout << it->first << ":" << it->second << " ";  // apple:10 banana:20 orange:15
    }
    std::cout << std::endl;
    
    return 0;
}
```

### 2. 自定义比较器示例

```cpp
#include <iostream>
#include <set>

// 自定义比较器
struct ReverseCompare {
    bool operator()(int a, int b) const {
        return a > b;  // 降序排列
    }
};

int main() {
    // 使用自定义比较器的set
    std::set<int, ReverseCompare> s;
    s.insert(3);
    s.insert(8);
    s.insert(5);
    s.insert(9);
    s.insert(13);
    
    // 遍历set（降序）
    std::cout << "set elements (reverse): ";
    for (auto it = s.begin(); it != s.end(); ++it) {
        std::cout << *it << " ";  // 13 9 8 5 3
    }
    std::cout << std::endl;
    
    return 0;
}
```

## 四、总结

### 1. 核心知识点回顾

- **STL容器分类**：序列式容器和关联式容器，各有特点和适用场景
- **Rb_tree原理**：平衡二叉搜索树，通过红黑性质保持适度平衡
- **Rb_tree应用**：作为set、map、multiset、multimap的底层实现
- **容器选择**：根据数据访问模式、插入删除频率、内存开销等因素选择合适的容器

### 2. 学习建议

1. **理解容器底层实现**：了解不同容器的底层实现有助于选择合适的容器
2. **掌握容器特性**：熟悉各容器的时间复杂度和适用场景
3. **实践应用**：通过实际代码练习加深对容器的理解
4. **关注C++标准更新**：了解最新C++标准对容器的改进
5. **性能优化**：根据具体场景选择最合适的容器，必要时使用自定义分配器

### 3. 未来发展趋势

- **并行容器**：支持多线程并发操作的容器
- **内存高效容器**：针对小对象的优化容器
- **领域特定容器**：针对特定领域的专用容器
- **编译期容器**：利用constexpr实现编译期计算的容器

通过深入学习STL容器和Rb_tree，我们可以更灵活地使用STL，编写更高效、更可靠的C++代码。容器是STL的基础，掌握它们的使用和实现原理，对于C++编程能力的提升至关重要。