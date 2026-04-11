# 侯捷 C++ STL标准库和泛型编程 - Set、Map和Hashtable深度探索
> **打卡日期**：2026-04-10
 > **核心主题**：set、multiset、map、multimap的实现原理以及hashtable的深度探索。
 > **资料出处**：https://www.bilibili.com/video/BV1kBh5zrEYh?spm_id_from=333.788.videopod.sections&vd_source=e01209adaee5cf8495fa73fbfde26dd1

---
## 一、set和multiset深度探索

### 1. 基本概念

**重要程度：基础**

set和multiset是关联式容器，它们的特点：

- **底层实现**：以rb_tree为底层结构
- **自动排序**：元素会根据键自动排序
- **键值关系**：value和key合一，value就是key
- **迭代器**：提供遍历操作及iterators，按正常规则(++ite)遍历可获得排序状态
- **不可修改**：无法使用iterators修改元素值，因为key有严格的排列规则，set/multiset的iterator是底层RB tree的const_iterator

### 2. set的实现

**重要程度：进阶**

#### 模板参数

```cpp
template <class Key,
          class Compare = less<Key>,
          class Alloc = alloc>
class set {
public:
    // typedefs:
    typedef Key key_type;
    typedef Key value_type;
    typedef Compare key_compare;
    typedef Compare value_compare;
private:
    typedef rb_tree<key_type, value_type,
                   identity<value_type>, key_compare, Alloc> rep_type;
    rep_type t;
public:
    typedef typename rep_type::const_iterator iterator;
    // ...
};
```

#### 关键特性

- **键唯一性**：set元素的key必须独一无二，因此其insert()使用rb_tree的insert_unique()
- **迭代器**：set的iterator是const_iterator，禁止用户修改元素
- **操作转发**：set的所有操作都转调用底层t的操作，从这个意义看，set可视为一个container adapter

### 3. multiset的实现

**重要程度：进阶**

- **键可重复**：multiset元素的key可以重复，因此其insert()使用rb_tree的insert_equal()
- **其他特性**：与set基本相同，唯一区别是允许键重复

### 4. VC6中的实现

**重要程度：进阶**

VC6不提供identity()，其set实现使用自定义的Kfn结构体：

```cpp
template<class K, class Pr = less<K>,
         class A = allocator<K>>
class set {
public:
    typedef set<K, Pr, A> _Myt;
    typedef K value_type;
    struct Kfn : public unary_function<value_type, K> {
        const K& operator()(const value_type& _X) const
        { return (_X); }
    };
    typedef K key_type;
    typedef Pr key_compare;
    typedef A allocator_type;
    typedef _Tree<K, value_type, Kfn, Pr, A> _Imp;
protected:
    _Imp _Tr;
    // ...
};
```

### 5. 实践应用

**重要程度：进阶**

#### 使用示例

```cpp
void test_multiset(long& value) {
    cout << "\ntest_multiset()............. \n";
    multiset<string> c;
    char buf[10];
    clock_t timeStart = clock();
    for(long i=0; i < value; ++i) {
        try {
            sprintf(buf, "%d", rand());
            c.insert(string(buf));
        } catch(exception& p) {
            cout << " " << p.what() << endl;  // out of memory
            abort();
        }
    }
    cout << "milli-seconds: " << (clock()-timeStart) << endl;  //
    cout << "multiset.size()= " << c.size() << endl;
    cout << "multiset.max_size()= " << c.max_size() << endl;
    
    string target = get_a_target_string();
    { // 此作用域用于测试find
        clock_t timeStart = clock();
        auto pItem = find(c.begin(), c.end(), target);  // 比c.find()慢很多
        cout << "::find(), milli-seconds: " << (clock()-timeStart) << endl;
        if (pItem != c.end())
            cout << "found, " << *pItem << endl;
        else
            cout << "not found!" << endl;
    }
    { // 此作用域用于测试find
        clock_t timeStart = clock();
        auto pItem = c.find(target);  // 比::find()快很多
        cout << "c.find(), milli-seconds: " << (clock()-timeStart) << endl;
        if (pItem != c.end())
            cout << "found, " << *pItem << endl;
        else
            cout << "not found!" << endl;
    }
}
```

### 6. 面试高频考点解析

**重要程度：高级**

#### 1. 问题：set和multiset的区别是什么？

**回答思路**：
- 键的唯一性
- 底层插入操作
- 应用场景

**参考答案**：
- set要求元素的key必须独一无二，使用rb_tree的insert_unique()
- multiset允许元素的key重复，使用rb_tree的insert_equal()
- 应用场景：set适合需要唯一键的场景，multiset适合需要允许键重复的场景

#### 2. 问题：为什么set的iterator是const_iterator？

**回答思路**：
- 元素排序的需要
- 底层实现的限制
- 设计意图

**参考答案**：
- set的元素是根据key自动排序的，如果允许修改元素值，会破坏排序规则
- set的iterator是底层rb_tree的const_iterator，这样设计是为了禁止用户对元素赋值
- 这种设计确保了set的有序性和一致性

#### 3. 问题：set和unordered_set的区别是什么？

**回答思路**：
- 底层实现
- 有序性
- 性能
- 内存开销

**参考答案**：
- 底层实现：set基于rb_tree，unordered_set基于hashtable
- 有序性：set是有序的，unordered_set是无序的
- 性能：
  - set：插入、查找、删除均为O(log n)
  - unordered_set：平均为O(1)，最坏为O(n)
- 内存开销：unordered_set更大，需要哈希表结构

### 7. 核心知识点强化

**重要程度：进阶**

#### 1. set的性能特性

- **时间复杂度**：
  - 插入：O(log n)
  - 查找：O(log n)
  - 删除：O(log n)
  - 遍历：O(n)

- **空间复杂度**：O(n)，每个元素需要额外的树节点开销

#### 2. 使用注意事项

- **不可修改元素**：set的iterator是const的，无法修改元素
- **自动排序**：插入元素会自动排序，影响插入顺序
- **键类型要求**：键类型必须支持比较操作
- **性能考量**：对于频繁查找的场景，set的O(log n)性能可能不如unordered_set的O(1)

### 8. 实践应用指南

**重要程度：进阶**

#### 1. 实际开发场景中的应用

**场景1：需要有序且唯一的数据**
- **应用**：存储不重复的有序数据
- **最佳实践**：使用set

**场景2：需要有序且允许重复的数据**
- **应用**：存储可能重复的有序数据
- **最佳实践**：使用multiset

**场景3：需要快速查找**
- **应用**：字典、集合操作
- **最佳实践**：使用set，利用其O(log n)的查找性能

### 9. 知识关联图谱

**重要程度：进阶**

```
set/multiset
├── 基本特性
│   ├── 底层实现：rb_tree
│   ├── 自动排序
│   ├── 键值合一
│   └── 不可修改元素
├── 核心操作
│   ├── insert
│   ├── find
│   ├── erase
│   └── size
├── 实现差异
│   ├── set：键唯一，使用insert_unique()
│   └── multiset：键可重复，使用insert_equal()
└── 应用场景
    ├── set：唯一键场景
    └── multiset：可重复键场景
```

## 二、map和multimap深度探索

### 1. 基本概念

**重要程度：基础**

map和multimap是关联式容器，它们的特点：

- **底层实现**：以rb_tree为底层结构
- **自动排序**：元素会根据key自动排序
- **键值对**：存储键值对(pair<const Key, T>)
- **迭代器**：提供遍历操作及iterators，按正常规则(++ite)遍历可获得排序状态
- **部分可修改**：无法修改元素的key，但可以修改元素的data

### 2. map的实现

**重要程度：进阶**

#### 模板参数

```cpp
template <class Key,
          class T,
          class Compare = less<Key>,
          class Alloc = alloc>
class map {
public:
    typedef Key key_type;
    typedef T data_type;
    typedef T mapped_type;
    typedef pair<const Key, T> value_type;
    typedef Compare key_compare;
private:
    typedef rb_tree<key_type, value_type,
                   select1st<value_type>, key_compare, Alloc> rep_type;
    rep_type t;
public:
    typedef typename rep_type::iterator iterator;
    // ...
};
```

#### 关键特性

- **键唯一性**：map元素的key必须独一无二，因此其insert()使用rb_tree的insert_unique()
- **迭代器**：map的iterator允许修改元素的data，但不能修改key
- **operator[]**：map提供独特的operator[]操作，方便访问和插入元素

### 3. multimap的实现

**重要程度：进阶**

- **键可重复**：multimap元素的key可以重复，因此其insert()使用rb_tree的insert_equal()
- **无operator[]**：multimap不提供operator[]操作，因为键可能重复
- **其他特性**：与map基本相同，唯一区别是允许键重复

### 4. VC6中的实现

**重要程度：进阶**

VC6不提供select1st()，其map实现使用自定义的Kfn结构体：

```cpp
template<class K, class Ty, class Pr = less<K>,
         class A = allocator<Ty>>
class map {
public:
    typedef map<K, Ty, Pr, A> _Myt;
    typedef pair<const K, Ty> value_type;
    struct Kfn : public unary_function<value_type, K> {
        const K& operator()(const value_type& _X) const
        { return (_X.first); }
    };
    // ...
    typedef _Tree<K, value_type, Kfn, Pr, A> _Imp;
protected:
    _Imp _Tr;
    // ...
};
```

### 5. map的operator[]

**重要程度：高级**

map的operator[]实现：

```cpp
mapped_type& operator[](const key_type& k) {
    // concept requirements
    __glibcxx_function_requires(_DefaultConstructibleConcept<mapped_type>)

    iterator i = lower_bound(k);
    // i is the first element >= k.
    if (i == end() || key_comp()(k, (*i).first))
#if __cplusplus >= 201103L
        i = _M_t._M_emplace_hint_unique(i, std::piecewise_construct,
                                         std::tuple<const key_type&>(k),
                                         std::tuple<>());
#else
        i = insert(i, value_type(k, mapped_type()));
#endif
    return (*i).second;
}
```

**特点**：
- 使用lower_bound进行二分查找
- 如果key不存在，会创建一个默认构造的value
- 返回对value的引用

### 6. 实践应用

**重要程度：进阶**

#### 使用示例

```cpp
void test_multimap(long& value) {
    cout << "\ntest_multimap()............. \n";
    multimap<long, string> c;
    char buf[10];
    clock_t timeStart = clock();
    for(long i=0; i < value; ++i) {
        try {
            sprintf(buf, "%d", rand());
            c.insert(pair<long, string>(i, string(buf)));
        } catch(exception& p) {
            cout << " " << i << " " << p.what() << endl;  // out of memory
            abort();
        }
    }
    cout << "milli-seconds: " << (clock()-timeStart) << endl;  //
    cout << "multimap.size()= " << c.size() << endl;
    cout << "multimap.max_size()= " << c.max_size() << endl;
    
    long target = get_a_target_long();
    { // 此作用域用于测试find
        clock_t timeStart = clock();
        auto pItem = c.find(target);
        cout << "c.find(), milli-seconds: " << (clock()-timeStart) << endl;
        if (pItem != c.end())
            cout << "found, value=" << (*pItem).second << endl;
        else
            cout << "not found!" << endl;
    }
}
```

### 7. 面试高频考点解析

**重要程度：高级**

#### 1. 问题：map和multimap的区别是什么？

**回答思路**：
- 键的唯一性
- 底层插入操作
- operator[]的支持
- 应用场景

**参考答案**：
- map要求元素的key必须独一无二，使用rb_tree的insert_unique()
- multimap允许元素的key重复，使用rb_tree的insert_equal()
- map提供operator[]操作，multimap不提供
- 应用场景：map适合需要唯一键值映射的场景，multimap适合需要允许键重复的映射场景

#### 2. 问题：map的operator[]是如何实现的？

**回答思路**：
- 查找机制
- 插入逻辑
- 返回值
- 性能分析

**参考答案**：
- map的operator[]使用lower_bound进行二分查找
- 如果key不存在，会创建一个默认构造的value并插入
- 返回对value的引用
- 时间复杂度：O(log n)，因为需要进行二分查找

#### 3. 问题：map和unordered_map的区别是什么？

**回答思路**：
- 底层实现
- 有序性
- 性能
- 内存开销

**参考答案**：
- 底层实现：map基于rb_tree，unordered_map基于hashtable
- 有序性：map是有序的，unordered_map是无序的
- 性能：
  - map：插入、查找、删除均为O(log n)
  - unordered_map：平均为O(1)，最坏为O(n)
- 内存开销：unordered_map更大，需要哈希表结构

### 8. 核心知识点强化

**重要程度：进阶**

#### 1. map的性能特性

- **时间复杂度**：
  - 插入：O(log n)
  - 查找：O(log n)
  - 删除：O(log n)
  - 遍历：O(n)
  - operator[]：O(log n)

- **空间复杂度**：O(n)，每个元素需要额外的树节点开销

#### 2. 使用注意事项

- **键的const性**：map的key是const的，无法修改
- **自动排序**：插入元素会根据key自动排序
- **operator[]的副作用**：当key不存在时，会自动插入默认构造的value
- **性能考量**：对于频繁查找的场景，map的O(log n)性能可能不如unordered_map的O(1)

### 9. 实践应用指南

**重要程度：进阶**

#### 1. 实际开发场景中的应用

**场景1：键值映射**
- **应用**：字典、配置管理、缓存
- **最佳实践**：使用map

**场景2：允许多个值对应同一个键**
- **应用**：多值映射、分组数据
- **最佳实践**：使用multimap

**场景3：需要有序的键值对**
- **应用**：需要按键排序的场景
- **最佳实践**：使用map或multimap

### 10. 知识关联图谱

**重要程度：进阶**

```
map/multimap
├── 基本特性
│   ├── 底层实现：rb_tree
│   ├── 自动排序
│   ├── 键值对
│   └── 可修改data，不可修改key
├── 核心操作
│   ├── insert
│   ├── find
│   ├── erase
│   ├── size
│   └── operator[] (仅map)
├── 实现差异
│   ├── map：键唯一，使用insert_unique()，支持operator[]
│   └── multimap：键可重复，使用insert_equal()，不支持operator[]
└── 应用场景
    ├── map：唯一键值映射
    └── multimap：可重复键值映射
```

## 三、hashtable深度探索

### 1. 基本概念

**重要程度：基础**

hashtable是一种高效的关联式容器底层实现，它的特点：

- **哈希函数**：将键映射到桶（bucket）
- **碰撞处理**：使用分离链接法（Separate Chaining）处理碰撞
- **动态调整**：当元素数量超过阈值时，会进行rehashing
- **性能**：平均情况下，插入、查找、删除操作的时间复杂度为O(1)

### 2. hashtable的内部结构

**重要程度：进阶**

#### 模板参数

```cpp
template <class Value, class Key, class HashFcn,
          class ExtractKey, class EqualKey,
          class Alloc = alloc>
class hashtable {
public:
    typedef HashFcn hasher;
    typedef EqualKey key_equal;
    typedef size_t size_type;
private:
    hasher hash;
    key_equal equals;
    ExtractKey get_key;
    typedef __hashtable_node<Value> node;
    vector<node*, Alloc> buckets;
    size_type num_elements;
public:
    size_type bucket_count() const { return buckets.size(); }
    // ...
};
```

#### 节点结构

```cpp
template <class Value>
struct __hashtable_node {
    __hashtable_node* next;
    Value val;
};
```

#### 迭代器

```cpp
template <class Value, class Key, class HashFcn,
          class ExtractKey, class EqualKey>
struct __hashtable_iterator {
    node* cur;
    hashtable* ht;
    // ...
};
```

### 3. 哈希函数

**重要程度：进阶**

#### 基本哈希函数

STL提供了基本类型的哈希函数特化：

```cpp
template <class Key> struct hash {};

// 特化版本
template<> struct hash<char> {
    size_t operator()(char x) const { return x; }
};

template<> struct hash<short> {
    size_t operator()(short x) const { return x; }
};

template<> struct hash<int> {
    size_t operator()(int x) const { return x; }
};

template<> struct hash<long> {
    size_t operator()(long x) const { return x; }
};
```

#### 字符串哈希函数

```cpp
inline size_t __stl_hash_string(const char* s) {
    unsigned long h = 0;
    for (; *s; ++s)
        h = 5*h + *s;
    return size_t(h);
}

template<> struct hash<char*> {
    size_t operator()(const char* s) const {
        return __stl_hash_string(s);
    }
};

template<> struct hash<const char*> {
    size_t operator()(const char* s) const {
        return __stl_hash_string(s);
    }
};
```

**注意**：标准库没有提供现成的`hash<std::string>`，需要用户自行实现或使用C++11及以后版本的标准库。

#### 哈希函数的目的

哈希函数的目的是希望根据元素值算出一个hash code（一个可进行modulus运算的值），使得元素经hash code映射之后能够「够乱够随机」地被置于hashtable内。愈是紊乱，愈不容易发生碰撞。

### 4. 哈希冲突处理

**重要程度：进阶**

hashtable使用分离链接法（Separate Chaining）处理哈希冲突：

- 每个桶（bucket）是一个链表的头指针
- 当多个键哈希到同一个桶时，它们被链接在同一个链表中
- 查找时，先通过哈希函数找到桶，然后在链表中线性查找

### 5. Modulus运算

**重要程度：进阶**

在hashtable中，modulus运算用于将哈希值映射到具体的桶索引：

```cpp
// 计算键对应的桶编号
size_type bkt_num_key(const key_type& key) const {
    return bkt_num_key(key, buckets.size());
}

// 计算值对应的桶编号
size_type bkt_num(const value_type& obj) const {
    return bkt_num_key(get_key(obj));
}

// 核心modulus运算
size_type bkt_num_key(const key_type& key, size_t n) const {
    return hash(key) % n;
}

// 带桶大小的桶编号计算
size_type bkt_num(const value_type& obj, size_t n) const {
    return bkt_num_key(get_key(obj), n);
}
```

**注意**：这里的`hash`不是前面出现的`struct hash`，而是`class hashtable`中的`hasher hash`成员变量。

### 6. Rehashing

**重要程度：高级**

当hashtable中的元素数量超过桶数量时，会进行rehashing：

1. 计算新的桶数量（通常是原桶数量的两倍左右，且为质数）
2. 分配新的桶数组
3. 将所有元素重新哈希到新的桶中
4. 释放旧的桶数组

#### 质数表

STL使用预定义的质数表来选择桶数量：

```cpp
static const unsigned long _stl_prime_list[_stl_num_primes] = {
    53,         97,         193,        389,        769,
    1543,       3079,       6151,       12289,      24593,
    49157,      98317,      196613,     393241,     786433,
    1572869,    3145739,    6291469,    12582917,   25165843,
    50331653,   100663319,  201326611,  402653189,  805306457,
    1610612741, 3221225473, 4294967291
};
```

### 6. 插入操作

**重要程度：进阶**

hashtable提供两种插入操作：

- **insert_unique**：插入唯一键，若键已存在则插入失败
- **insert_equal**：插入可重复键，允许键重复

### 7. 面试高频考点解析

**重要程度：高级**

#### 1. 问题：hashtable的工作原理是什么？

**回答思路**：
- 哈希函数的作用
- 桶的概念
- 碰撞处理
- Rehashing机制

**参考答案**：
- hashtable使用哈希函数将键映射到桶（bucket）
- 每个桶是一个链表的头指针
- 使用分离链接法处理碰撞，多个键哈希到同一个桶时，链接在同一个链表中
- 当元素数量超过阈值时，会进行rehashing，扩大桶数量并重新分配元素
- 平均情况下，插入、查找、删除操作的时间复杂度为O(1)

#### 2. 问题：哈希冲突的处理方法有哪些？

**回答思路**：
- 分离链接法
- 开放寻址法
- 再哈希法
- 建立公共溢出区

**参考答案**：
- **分离链接法**：每个桶维护一个链表，哈希到同一桶的元素都放在这个链表中
- **开放寻址法**：当发生冲突时，使用某种探测方法在哈希表中寻找下一个空位置
- **再哈希法**：使用多个哈希函数，当发生冲突时使用下一个哈希函数
- **建立公共溢出区**：将冲突的元素都放入一个公共的溢出区
- STL的hashtable使用分离链接法

#### 3. 问题：为什么hashtable的桶数量要选择质数？

**回答思路**：
- 减少碰撞
- 均匀分布
- 数学原理

**参考答案**：
- 选择质数作为桶数量可以减少哈希冲突，使元素分布更均匀
- 因为质数的因数较少，能减少哈希值的周期性重复
- 这样可以提高hashtable的查找效率
- STL使用预定义的质数表来选择桶数量

### 8. 核心知识点强化

**重要程度：进阶**

#### 1. hashtable的性能特性

- **时间复杂度**：
  - 插入：平均O(1)，最坏O(n)
  - 查找：平均O(1)，最坏O(n)
  - 删除：平均O(1)，最坏O(n)
  - 遍历：O(n)

- **空间复杂度**：O(n + m)，其中n是元素数量，m是桶数量

#### 2. 使用注意事项

- **哈希函数选择**：好的哈希函数应能将键均匀分布到桶中
- **键类型要求**：键类型必须支持哈希函数和相等比较
- **rehashing开销**：当元素数量增长时，会触发rehashing，可能影响性能
- **内存开销**：需要额外的桶数组和链表节点开销

### 9. 实践应用指南

**重要程度：进阶**

#### 1. 实际开发场景中的应用

**场景1：需要快速查找**
- **应用**：字典、缓存、集合
- **最佳实践**：使用基于hashtable的容器（如unordered_set、unordered_map）

**场景2：需要键值映射**
- **应用**：配置管理、计数器
- **最佳实践**：使用unordered_map

**场景3：需要处理大量数据**
- **应用**：大数据处理、索引
- **最佳实践**：使用hashtable，利用其O(1)的平均时间复杂度

### 10. 知识关联图谱

**重要程度：进阶**

```
hashtable
├── 基本特性
│   ├── 哈希函数
│   ├── 桶和链表
│   ├── 分离链接法处理碰撞
│   └── Rehashing机制
├── 内部结构
│   ├── 模板参数
│   ├── 节点结构
│   ├── 迭代器
│   └── 桶数组
├── 核心操作
│   ├── insert_unique
│   ├── insert_equal
│   ├── find
│   ├── erase
│   └── bucket_count
├── 性能特性
│   ├── 平均时间复杂度：O(1)
│   ├── 最坏时间复杂度：O(n)
│   └── 空间复杂度：O(n + m)
└── 应用场景
    ├── unordered_set
    ├── unordered_map
    ├── unordered_multiset
    └── unordered_multimap
```

## 四、unordered容器概念

### 1. 基本概念

**重要程度：基础**

unordered容器是C++11引入的关联式容器，它们的特点：

- **底层实现**：以hashtable为底层结构
- **无序性**：元素不按顺序存储
- **高效性**：平均情况下，插入、查找、删除操作的时间复杂度为O(1)
- **键要求**：键类型必须支持哈希函数和相等比较

### 2. 容器种类

**重要程度：基础**

- **unordered_set**：存储唯一键，无序
- **unordered_map**：存储键值对，键唯一，无序
- **unordered_multiset**：存储可重复键，无序
- **unordered_multimap**：存储键值对，键可重复，无序

### 3. C++11前后的变化

**重要程度：进阶**

| 时期 | 容器名称 |
|------|----------|
| Before C++11 | hash_set, hash_multiset, hash_map, hash_multimap |
| Since C++11 | unordered_set, unordered_multiset, unordered_map, unordered_multimap |

### 4. 模板定义

**重要程度：进阶**

```cpp
// unordered_set
template <typename T,
          typename Hash = hash<T>,
          typename EqPred = equal_to<T>,
          typename Allocator = allocator<T> >
class unordered_set;

// unordered_multiset
template <typename T,
          typename Hash = hash<T>,
          typename EqPred = equal_to<T>,
          typename Allocator = allocator<T> >
class unordered_multiset;

// unordered_map
template <typename Key, typename T,
          typename Hash = hash<Key>,
          typename EqPred = equal_to<Key>,
          typename Allocator = allocator<pair<const Key, T> > >
class unordered_map;

// unordered_multimap
template <typename Key, typename T,
          typename Hash = hash<Key>,
          typename EqPred = equal_to<Key>,
          typename Allocator = allocator<pair<const Key, T> > >
class unordered_multimap;
```

### 5. 与有序容器的对比

**重要程度：进阶**

| 特性 | 有序容器（set/map） | 无序容器（unordered_set/unordered_map） |
|------|-------------------|---------------------------------------|
| 底层实现 | rb_tree | hashtable |
| 有序性 | 有序 | 无序 |
| 时间复杂度 | O(log n) | 平均O(1)，最坏O(n) |
| 内存开销 | 较小 | 较大 |
| 键要求 | 支持比较操作 | 支持哈希函数和相等比较 |
| 适用场景 | 需要有序遍历 | 需要快速查找 |

### 4. 面试高频考点解析

**重要程度：高级**

#### 1. 问题：什么时候使用unordered容器而不是有序容器？

**回答思路**：
- 性能需求
- 有序性需求
- 内存限制
- 键类型特性

**参考答案**：
- 当需要快速的插入、查找、删除操作时，使用unordered容器
- 当不需要元素有序时，使用unordered容器
- 当键类型支持良好的哈希函数时，使用unordered容器
- 当内存不是主要限制因素时，使用unordered容器

#### 2. 问题：如何为自定义类型实现哈希函数？

**回答思路**：
- 特化hash模板
- 实现哈希函数的注意事项
- 示例代码

**参考答案**：
- 为自定义类型特化std::hash模板：
  ```cpp
  namespace std {
      template<> struct hash<MyType> {
          size_t operator()(const MyType& obj) const {
              // 实现哈希函数
              return hash<int>()(obj.id) ^ (hash<string>()(obj.name) << 1);
          }
      };
  }
  ```
- 注意事项：
  - 哈希函数应尽量均匀分布
  - 哈希函数应与operator==一致
  - 对于复杂类型，可组合多个成员的哈希值

### 5. 核心知识点强化

**重要程度：进阶**

#### 1. unordered容器的性能特性

- **时间复杂度**：
  - 插入：平均O(1)，最坏O(n)
  - 查找：平均O(1)，最坏O(n)
  - 删除：平均O(1)，最坏O(n)
  - 遍历：O(n)

- **空间复杂度**：O(n)，需要额外的哈希表结构

#### 2. 使用注意事项

- **哈希函数选择**：好的哈希函数对性能至关重要
- **键类型要求**：键类型必须支持哈希函数和相等比较
- **负载因子**：控制哈希表的密度，影响性能
- **rehashing**：当元素数量增长时，会触发rehashing

### 6. 实践应用指南

**重要程度：进阶**

#### 1. 实际开发场景中的应用

**场景1：需要快速查找**
- **应用**：字典、缓存、集合
- **最佳实践**：使用unordered_set或unordered_map

**场景2：不需要有序遍历**
- **应用**：计数、去重
- **最佳实践**：使用unordered容器

**场景3：处理大量数据**
- **应用**：大数据处理、索引
- **最佳实践**：使用unordered容器，利用其O(1)的平均时间复杂度

### 7. 知识关联图谱

**重要程度：进阶**

```
unordered容器
├── 基本特性
│   ├── 底层实现：hashtable
│   ├── 无序性
│   ├── 平均O(1)时间复杂度
│   └── 键要求：哈希函数和相等比较
├── 容器种类
│   ├── unordered_set：唯一键
│   ├── unordered_map：键值对，键唯一
│   ├── unordered_multiset：可重复键
│   └── unordered_multimap：键值对，键可重复
├── 与有序容器对比
│   ├── 性能：更快的平均时间复杂度
│   ├── 有序性：无序
│   ├── 内存开销：更大
│   └── 键要求：需要哈希函数
└── 应用场景
    ├── 快速查找
    ├── 不需要有序遍历
    └── 处理大量数据
```

## 五、代码示例与实践

### 1. set和multiset使用示例

```cpp
#include <iostream>
#include <set>

int main() {
    // 使用set
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
    
    // 使用multiset
    std::multiset<int> ms;
    ms.insert(3);
    ms.insert(8);
    ms.insert(5);
    ms.insert(5);  // 重复键，会插入
    
    std::cout << "multiset size: " << ms.size() << std::endl;  // 4
    std::cout << "count(5): " << ms.count(5) << std::endl;  // 2
    
    // 遍历multiset
    std::cout << "multiset elements: ";
    for (auto it = ms.begin(); it != ms.end(); ++it) {
        std::cout << *it << " ";  // 3 5 5 8
    }
    std::cout << std::endl;
    
    return 0;
}
```

### 2. map和multimap使用示例

```cpp
#include <iostream>
#include <map>

int main() {
    // 使用map
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
    
    // 使用multimap
    std::multimap<std::string, int> mm;
    mm.insert(std::make_pair("apple", 10));
    mm.insert(std::make_pair("banana", 20));
    mm.insert(std::make_pair("apple", 15));  // 重复键，会插入
    
    std::cout << "multimap size: " << mm.size() << std::endl;  // 3
    
    // 遍历multimap
    std::cout << "multimap elements: ";
    for (auto it = mm.begin(); it != mm.end(); ++it) {
        std::cout << it->first << ":" << it->second << " ";  // apple:10 apple:15 banana:20
    }
    std::cout << std::endl;
    
    return 0;
}
```

### 3. unordered容器使用示例

#### unordered_set使用示例

```cpp
namespace jj15 {
void test_unordered_set(long& value) {
    cout << "\ntest_unordered_set()............. \n";
    unordered_set<string> c;
    char buf[10];
    clock_t timeStart = clock();
    for(long i=0; i < value; ++i) {
        try {
            sprintf(buf, "%d", rand());
            c.insert(string(buf));
        } catch(exception& p) {
            cout << " " << i << " " << p.what() << endl;  // out of memory
            abort();
        }
    }
    cout << "milli-seconds: " << (clock()-timeStart) << endl;  //
    cout << "unordered_set.size()= " << c.size() << endl;
    cout << "unordered_set.max_size()= " << c.max_size() << endl;
    cout << "unordered_set.bucket_count()= " << c.bucket_count() << endl;
    cout << "unordered_set.load_factor()= " << c.load_factor() << endl;
    cout << "unordered_set.max_load_factor()= " << c.max_load_factor() << endl;
    cout << "unordered_set.max_bucket_count()= " << c.max_bucket_count() << endl;
    for (unsigned i=0; i < 20; ++i) {
        cout << "bucket " << i << " has " << c.bucket_size(i) << " elements.\n";
    }
    
    string target = get_a_target_string();
    { // 此作用域用于测试find
        clock_t timeStart = clock();
        auto pItem = find(c.begin(), c.end(), target);  // 比c.find()慢很多
        cout << "::find(), milli-seconds: " << (clock()-timeStart) << endl;
        if (pItem != c.end())
            cout << "found, " << *pItem << endl;
        else
            cout << "not found!" << endl;
    }
    { // 此作用域用于测试find
        clock_t timeStart = clock();
        auto pItem = c.find(target);  // 比::find()快很多
        cout << "c.find(), milli-seconds: " << (clock()-timeStart) << endl;
        if (pItem != c.end())
            cout << "found, " << *pItem << endl;
        else
            cout << "not found!" << endl;
    }
}
}
```

#### 运行结果示例

```
select: 15
how many elements: 1000000
test_unordered_set()............. 
milli-seconds: 2891
unordered_set.size()= 32768
unordered_set.max_size()= 357913941
unordered_set.bucket_count()= 62233
unordered_set.load_factor()= 0.526537
unordered_set.max_load_factor()= 1
unordered_set.max_bucket_count()= 357913941
bucket #0 has 1 elements.
bucket #1 has 1 elements.
bucket #2 has 1 elements.
bucket #3 has 3 elements.
bucket #4 has 0 elements.
bucket #5 has 0 elements.
bucket #6 has 1 elements.
bucket #7 has 0 elements.
bucket #8 has 0 elements.
bucket #9 has 0 elements.
bucket #10 has 1 elements.
bucket #11 has 0 elements.
bucket #12 has 1 elements.
bucket #13 has 2 elements.
bucket #14 has 1 elements.
bucket #15 has 1 elements.
bucket #16 has 0 elements.
bucket #17 has 1 elements.
bucket #18 has 0 elements.
bucket #19 has 0 elements.
Target (92376?): 23456
::find(), milli-seconds: 23456
found, 23456
c.find(), milli-seconds: 0
found, 23456
```

#### unordered_map使用示例

```cpp
#include <iostream>
#include <unordered_map>

int main() {
    // 使用unordered_map
    std::unordered_map<std::string, int> um;
    um["apple"] = 10;
    um["banana"] = 20;
    um["orange"] = 15;
    
    std::cout << "unordered_map size: " << um.size() << std::endl;  // 3
    std::cout << "apple: " << um["apple"] << std::endl;  // 10
    
    // 遍历unordered_map（无序）
    std::cout << "unordered_map elements: ";
    for (auto it = um.begin(); it != um.end(); ++it) {
        std::cout << it->first << ":" << it->second << " ";  // 顺序不确定
    }
    std::cout << std::endl;
    
    return 0;
}
```

## 六、总结

### 1. 核心知识点回顾

- **set和multiset**：基于rb_tree，有序，set键唯一，multiset键可重复
- **map和multimap**：基于rb_tree，有序，存储键值对，map键唯一且支持operator[]，multimap键可重复
- **hashtable**：底层实现，使用哈希函数和分离链接法处理碰撞，平均时间复杂度O(1)
- **unordered容器**：基于hashtable，无序，平均时间复杂度O(1)，包括unordered_set、unordered_map等

### 2. 容器选择指南

| 需求 | 推荐容器 | 原因 |
|------|----------|------|
| 有序且唯一键 | set | 自动排序，键唯一 |
| 有序且可重复键 | multiset | 自动排序，键可重复 |
| 有序键值映射 | map | 自动排序，键唯一，支持operator[] |
| 有序且可重复键值映射 | multimap | 自动排序，键可重复 |
| 快速查找且唯一键 | unordered_set | 平均O(1)查找，键唯一 |
| 快速查找且可重复键 | unordered_multiset | 平均O(1)查找，键可重复 |
| 快速键值映射 | unordered_map | 平均O(1)查找，键唯一 |
| 快速且可重复键值映射 | unordered_multimap | 平均O(1)查找，键可重复 |

### 3. 学习建议

1. **理解底层实现**：了解不同容器的底层实现有助于选择合适的容器
2. **掌握性能特性**：熟悉各容器的时间复杂度和空间复杂度
3. **实践应用**：通过实际代码练习加深对容器的理解
4. **关注C++标准**：了解C++11及以后版本引入的unordered容器
5. **性能优化**：根据具体场景选择最合适的容器，必要时优化哈希函数

### 4. 未来发展趋势

- **并行容器**：支持多线程并发操作的容器
- **内存高效容器**：针对小对象的优化容器
- **领域特定容器**：针对特定领域的专用容器
- **编译期容器**：利用constexpr实现编译期计算的容器

通过深入学习这些容器，我们可以更灵活地使用STL，编写更高效、更可靠的C++代码。容器是STL的基础，掌握它们的使用和实现原理，对于C++编程能力的提升至关重要。