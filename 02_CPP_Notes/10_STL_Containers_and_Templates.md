# 侯捷 C++ STL标准库和泛型编程 - 容器与模板复盘 10

> **打卡日期**：2026-04-06
> **核心主题**：STL容器分类与测试、分配器原理、源代码分布、OOP vs GP、操作符重载、模板（泛化、全特化、偏特化）。
> **资料出处**：https://www.bilibili.com/video/BV1kBh5zrEYh?spm_id_from=333.788.videopod.sections&vd_source=e01209adaee5cf8495fa73fbfde26dd1

---

## 一、 STL容器分类与测试

### 1. 容器的基本分类

STL容器可以分为以下几类：

#### 1.1 序列容器 (Sequence Containers)
- **vector**：动态数组，支持快速随机访问
- **list**：双向链表，支持快速插入和删除
- **deque**：双端队列，支持两端的快速插入和删除
- **array**：固定大小的数组（C++11新增）
- **forward_list**：单向链表（C++11新增）

#### 1.2 关联容器 (Associative Containers)
- **set**：有序集合，元素唯一
- **multiset**：有序集合，元素可重复
- **map**：有序键值对，键唯一
- **multimap**：有序键值对，键可重复

#### 1.3 无序关联容器 (Unordered Associative Containers)（C++11新增）
- **unordered_set**：无序集合，元素唯一，基于哈希表
- **unordered_multiset**：无序集合，元素可重复，基于哈希表
- **unordered_map**：无序键值对，键唯一，基于哈希表
- **unordered_multimap**：无序键值对，键可重复，基于哈希表

#### 1.4 容器适配器 (Container Adapters)
- **stack**：栈，后进先出
- **queue**：队列，先进先出
- **priority_queue**：优先队列，基于堆

### 2. 容器的特性对比

| 容器 | 随机访问 | 插入/删除（中间） | 插入/删除（两端） | 内存布局 | 迭代器类型 |
|------|----------|-----------------|------------------|----------|------------|
| vector | O(1) | O(n) | 尾端O(1)，前端O(n) | 连续 | 随机访问 |
| list | O(n) | O(1) | O(1) | 非连续 | 双向 |
| deque | O(1) | O(n) | O(1) | 分段连续 | 随机访问 |
| array | O(1) | 不支持 | 不支持 | 连续 | 随机访问 |
| forward_list | O(n) | O(1) | 前端O(1)，后端O(n) | 非连续 | 前向 |
| set/map | O(logn) | O(logn) | 不适用 | 树结构 | 双向 |
| unordered_set/map | 平均O(1)，最坏O(n) | 平均O(1)，最坏O(n) | 不适用 | 哈希表 | 前向 |

### 3. 容器测试

#### 3.1 基本操作测试

```cpp
#include <iostream>
#include <vector>
#include <list>
#include <deque>
#include <set>
#include <map>
#include <unordered_set>
#include <unordered_map>
#include <stack>
#include <queue>

using namespace std;

void testSequenceContainers() {
    // vector
    vector<int> v = {1, 2, 3};
    v.push_back(4);
    cout << "Vector: ";
    for (auto i : v) cout << i << " ";
    cout << endl;
    
    // list
    list<int> l = {1, 2, 3};
    l.push_front(0);
    l.push_back(4);
    cout << "List: ";
    for (auto i : l) cout << i << " ";
    cout << endl;
    
    // deque
    deque<int> d = {1, 2, 3};
    d.push_front(0);
    d.push_back(4);
    cout << "Deque: ";
    for (auto i : d) cout << i << " ";
    cout << endl;
}

void testAssociativeContainers() {
    // set
    set<int> s = {3, 1, 4, 1, 5, 9};
    cout << "Set: ";
    for (auto i : s) cout << i << " ";
    cout << endl;
    
    // map
    map<string, int> m;
    m["Alice"] = 25;
    m["Bob"] = 30;
    m["Charlie"] = 35;
    cout << "Map: ";
    for (auto& pair : m) cout << pair.first << ":" << pair.second << " ";
    cout << endl;
}

void testUnorderedContainers() {
    // unordered_set
    unordered_set<int> us = {3, 1, 4, 1, 5, 9};
    cout << "Unordered Set: ";
    for (auto i : us) cout << i << " ";
    cout << endl;
    
    // unordered_map
    unordered_map<string, int> um;
    um["Alice"] = 25;
    um["Bob"] = 30;
    um["Charlie"] = 35;
    cout << "Unordered Map: ";
    for (auto& pair : um) cout << pair.first << ":" << pair.second << " ";
    cout << endl;
}

void testContainerAdapters() {
    // stack
    stack<int> st;
    st.push(1);
    st.push(2);
    st.push(3);
    cout << "Stack: ";
    while (!st.empty()) {
        cout << st.top() << " ";
        st.pop();
    }
    cout << endl;
    
    // queue
    queue<int> q;
    q.push(1);
    q.push(2);
    q.push(3);
    cout << "Queue: ";
    while (!q.empty()) {
        cout << q.front() << " ";
        q.pop();
    }
    cout << endl;
    
    // priority_queue
    priority_queue<int> pq;
    pq.push(3);
    pq.push(1);
    pq.push(4);
    pq.push(1);
    pq.push(5);
    cout << "Priority Queue: ";
    while (!pq.empty()) {
        cout << pq.top() << " ";
        pq.pop();
    }
    cout << endl;
}

int main() {
    cout << "=== Sequence Containers ===" << endl;
    testSequenceContainers();
    
    cout << "\n=== Associative Containers ===" << endl;
    testAssociativeContainers();
    
    cout << "\n=== Unordered Containers ===" << endl;
    testUnorderedContainers();
    
    cout << "\n=== Container Adapters ===" << endl;
    testContainerAdapters();
    
    return 0;
}
```

**执行结果：**
```
=== Sequence Containers ===
Vector: 1 2 3 4 
List: 0 1 2 3 4 
Deque: 0 1 2 3 4 

=== Associative Containers ===
Set: 1 3 4 5 9 
Map: Alice:25 Bob:30 Charlie:35 

=== Unordered Containers ===
Unordered Set: 5 9 3 1 4 
Unordered Map: Charlie:35 Bob:30 Alice:25 

=== Container Adapters ===
Stack: 3 2 1 
Queue: 1 2 3 
Priority Queue: 5 4 3 1 1 
```

#### 3.2 性能测试

**插入性能测试：**

```cpp
#include <iostream>
#include <vector>
#include <list>
#include <deque>
#include <chrono>

using namespace std;
using namespace chrono;

void testInsertPerformance() {
    const int N = 100000;
    
    // 测试vector插入性能
    vector<int> v;
    auto start = high_resolution_clock::now();
    for (int i = 0; i < N; ++i) {
        v.push_back(i);
    }
    auto end = high_resolution_clock::now();
    auto duration = duration_cast<milliseconds>(end - start).count();
    cout << "Vector push_back: " << duration << " ms" << endl;
    
    // 测试list插入性能
    list<int> l;
    start = high_resolution_clock::now();
    for (int i = 0; i < N; ++i) {
        l.push_back(i);
    }
    end = high_resolution_clock::now();
    duration = duration_cast<milliseconds>(end - start).count();
    cout << "List push_back: " << duration << " ms" << endl;
    
    // 测试deque插入性能
    deque<int> d;
    start = high_resolution_clock::now();
    for (int i = 0; i < N; ++i) {
        d.push_back(i);
    }
    end = high_resolution_clock::now();
    duration = duration_cast<milliseconds>(end - start).count();
    cout << "Deque push_back: " << duration << " ms" << endl;
}

int main() {
    testInsertPerformance();
    return 0;
}
```

**执行结果（示例）：**
```
Vector push_back: 2 ms
List push_back: 4 ms
Deque push_back: 3 ms
```

### 4. 容器选择指南

| 场景 | 推荐容器 | 原因 |
|------|----------|------|
| 需要快速随机访问 | vector, array | 支持O(1)随机访问 |
| 需要频繁在中间插入/删除 | list | 支持O(1)中间插入/删除 |
| 需要频繁在两端插入/删除 | deque | 支持O(1)两端操作 |
| 需要有序且唯一的元素 | set | 自动排序且元素唯一 |
| 需要有序的键值对 | map | 自动按键排序 |
| 需要快速查找（平均O(1)） | unordered_set, unordered_map | 基于哈希表，查找速度快 |
| 需要后进先出操作 | stack | 栈结构 |
| 需要先进先出操作 | queue | 队列结构 |
| 需要优先级操作 | priority_queue | 自动按优先级排序 |

---

## 二、 分配器 (Allocator) 测试与原理

### 1. 分配器的基本概念

**分配器**是STL中的一个重要组件，负责容器的内存管理，包括内存的分配、释放和对象的构造、析构。

### 2. 标准分配器

C++标准库提供了默认的分配器 `std::allocator<T>`，它是所有容器的默认模板参数。

```cpp
// 显式指定分配器
vector<int, allocator<int>> v;
```

### 3. 分配器的测试

```cpp
#include <iostream>
#include <vector>
#include <memory>

using namespace std;

void testAllocator() {
    // 使用默认分配器
    allocator<int> alloc;
    
    // 分配5个int的内存
    int* p = alloc.allocate(5);
    
    // 构造对象
    for (int i = 0; i < 5; ++i) {
        alloc.construct(p + i, i);
    }
    
    // 使用对象
    cout << "Allocated and constructed elements: ";
    for (int i = 0; i < 5; ++i) {
        cout << p[i] << " ";
    }
    cout << endl;
    
    // 析构对象
    for (int i = 0; i < 5; ++i) {
        alloc.destroy(p + i);
    }
    
    // 释放内存
    alloc.deallocate(p, 5);
}

int main() {
    testAllocator();
    return 0;
}
```

**执行结果：**
```
Allocated and constructed elements: 0 1 2 3 4 
```

### 4. 分配器的底层实现与内存池（Memory Pool）黑魔法

标准分配器的底层实现通常调用 `operator new` 和 `operator delete`，但为了提高性能，许多实现会采用内存池技术。

标准分配器（如 VC++ 或 GnuC++ 4.9+ 的 `std::allocator`）底层只是简单地封装了 `::operator new` 和 `::operator delete`，本质上还是调用 C 语言的 `malloc` 和 `free`。

**🚨 致命缺陷：Cookie 内存碎片**
每次调用 `malloc` 分配内存时，操作系统不仅会分配你请求的内存，还会在内存块的头部和尾部加上额外的管理信息（称为 **Cookie**，包含区块大小和状态标志）。
如果容器（如 `list`、`map`）频繁分配极小的节点对象（比如只有 8 字节），附带的 Cookie 也会占 8 字节，导致高达 50% 的内存浪费！

**🧠 侯捷 PPT 核心知识点：GnuC++ 2.9 的 `alloc` 内存池设计**
为了消除小区块的 Cookie 开销，GnuC++ 2.9 设计了一个极其精妙的次层配置器（`__default_alloc_template`，即著名的 `alloc`）：
1. **16 条自由链表（Free-lists）**：它维护了一个包含 16 个指针的数组，分别管理大小为 8, 16, 24, ..., 128 字节的小内存块链表。
2. **内存池拔河**：当容器申请内存时，会将申请大小上调至 8 的倍数，然后直接去对应的链表中拔出一个区块。
3. **消除 Cookie**：因为一大块内存是一次性用 `malloc` 申请来放入内存池的，只有这一大块有 Cookie。切分成小块挂在链表上给容器使用时，小块**完全没有 Cookie**，极大地提升了内存利用率和分配速度！
### 5. 自定义分配器

在性能要求极高的场景（如高频交易、大模型算子等），可以自定义分配器来提高内存管理效率。

```cpp
#include <iostream>
#include <vector>

using namespace std;

// 简单的内存池分配器
template <typename T>
class MyAllocator {
public:
    using value_type = T;
    
    MyAllocator() noexcept {}
    
    template <typename U>
    MyAllocator(const MyAllocator<U>&) noexcept {}
    
    T* allocate(size_t n) {
        cout << "MyAllocator::allocate(" << n << ")" << endl;
        return static_cast<T*>(::operator new(n * sizeof(T)));
    }
    
    void deallocate(T* p, size_t n) noexcept {
        cout << "MyAllocator::deallocate(" << p << ", " << n << ")" << endl;
        ::operator delete(p);
    }
};

// 必须为不同类型的分配器提供相等比较
template <typename T, typename U>
bool operator==(const MyAllocator<T>&, const MyAllocator<U>&) noexcept {
    return true;
}

template <typename T, typename U>
bool operator!=(const MyAllocator<T>&, const MyAllocator<U>&) noexcept {
    return false;
}

int main() {
    vector<int, MyAllocator<int>> v;
    v.push_back(1);
    v.push_back(2);
    v.push_back(3);
    
    cout << "Vector elements: ";
    for (auto i : v) {
        cout << i << " ";
    }
    cout << endl;
    
    return 0;
}
```

**执行结果：**
```
MyAllocator::allocate(1)
MyAllocator::allocate(2)
MyAllocator::deallocate(000001F5A9C8A4E0, 1)
MyAllocator::allocate(4)
MyAllocator::deallocate(000001F5A9C8A4F0, 2)
Vector elements: 1 2 3 
MyAllocator::deallocate(000001F5A9C8A500, 4)
```

---

## 三、 STL源代码分布

### 1. 主要实现版本

- **HP STL**：原始版本，由Alexander Stepanov和Meng Lee开发
- **SGI STL**：Silicon Graphics Inc. 实现，被广泛使用
- **STLport**：跨平台实现
- **PJ STL**：Microsoft Visual C++ 实现
- **GNU STL**：GCC 实现 (libstdc++)

### 2. 源代码分布

#### 2.1 Visual C++ (VC) 中的STL源码

在Visual Studio中，STL源码通常位于：
- `C:\Program Files (x86)\Microsoft Visual Studio\[版本]\VC\Tools\MSVC\[版本]\include`

主要文件：
- `vector`：vector容器实现
- `list`：list容器实现
- `map`：map容器实现
- `algorithm`：算法实现
- `memory`：分配器实现

#### 2.2 GCC中的STL源码

在GCC中，STL源码通常位于：
- `/usr/include/c++/[版本]`

主要文件：
- `bits/stl_vector.h`：vector容器实现
- `bits/stl_list.h`：list容器实现
- `bits/stl_map.h`：map容器实现
- `bits/stl_algobase.h`：基础算法实现
- `bits/stl_alloc.h`：分配器实现

### 3. 源码阅读技巧

1. **从容器的接口入手**：先了解容器的公共接口，再深入内部实现
2. **关注内存管理**：了解容器如何分配和管理内存
3. **理解迭代器实现**：迭代器是容器和算法的桥梁
4. **分析性能特性**：了解不同操作的时间复杂度
5. **对比不同实现**：比较不同编译器对STL的实现差异

---

## 四、 OOP (面向对象编程) vs. GP (泛型编程)

### 1. 基本概念

- **OOP (面向对象编程)**：将数据和操作封装在一起，通过继承和多态实现代码复用
- **GP (泛型编程)**：将算法和数据结构分离，通过模板实现代码复用

### 2. 核心差异

| 特性 | OOP | GP |
|------|-----|----|
| 代码组织 | 数据和操作封装在一起 | 算法和数据结构分离 |
| 复用方式 | 继承和多态 | 模板实例化 |
| 类型检查 | 运行时多态 | 编译时多态 |
| 性能 | 虚函数调用有开销 | 模板实例化，无运行时开销 |
| 灵活性 | 运行时动态绑定 | 编译时静态绑定 |
| 代码膨胀 | 较少 | 可能导致代码膨胀 |

### 3. 实例对比

#### 3.1 OOP实现

```cpp
#include <iostream>

using namespace std;

// 基类
class Shape {
public:
    virtual void draw() = 0;
    virtual ~Shape() {}
};

// 派生类
class Circle : public Shape {
public:
    void draw() override {
        cout << "Drawing a circle" << endl;
    }
};

class Rectangle : public Shape {
public:
    void draw() override {
        cout << "Drawing a rectangle" << endl;
    }
};

// 使用
void drawShape(Shape* shape) {
    shape->draw(); // 运行时多态
}

int main() {
    Circle circle;
    Rectangle rectangle;
    
    drawShape(&circle);
    drawShape(&rectangle);
    
    return 0;
}
```

**执行结果：**
```
Drawing a circle
Drawing a rectangle
```

#### 3.2 GP实现

```cpp
#include <iostream>

using namespace std;

// 圆形
class Circle {
public:
    void draw() {
        cout << "Drawing a circle" << endl;
    }
};

// 矩形
class Rectangle {
public:
    void draw() {
        cout << "Drawing a rectangle" << endl;
    }
};

// 泛型函数
template <typename T>
void drawShape(T& shape) {
    shape.draw(); // 编译时多态
}

int main() {
    Circle circle;
    Rectangle rectangle;
    
    drawShape(circle);
    drawShape(rectangle);
    
    return 0;
}
```

**执行结果：**
```
Drawing a circle
Drawing a rectangle
```

### 4. STL中的GP应用

STL是泛型编程的典范，它将算法和数据结构分离，通过迭代器连接：

- **算法**：如`sort`、`find`等，通过模板实现，不依赖具体容器
- **容器**：如`vector`、`list`等，提供迭代器接口
- **迭代器**：作为算法和容器的桥梁，提供统一的访问接口

这种设计使得STL具有高度的可扩展性和复用性，是现代C++编程的重要工具。

### 5. 架构师避坑：OOP 与 GP 在 STL 中的边界碰撞（以 sort 为例）

既然 GP 提倡算法和容器完全解耦，提供全局的 `::sort()`，那为什么 `std::list` 还要在自己内部实现一个成员函数 `list::sort()` 呢？这不是退化回 OOP 了吗？

**底层原因揭秘：**
- 全局算法 `::sort()` 的底层实现是**快速排序（Quick Sort）**，它在进行指针偏移时（如 `mid = first + (last - first) / 2`），要求传入的迭代器必须是**随机访问迭代器（RandomAccessIterator）**。
- 但是，`list` 是双向链表，内存不连续，它提供的迭代器只是**双向迭代器（BidirectionalIterator）**，不支持 `+` 和 `-` 操作！
- 如果强行写 `::sort(list.begin(), list.end())`，在编译期就会报出极其恐怖的模板推导错误。
- 因此，STL 必须妥协，为 `list` 单独提供一个基于**归并排序（Merge Sort）**的 OOP 成员方法 `list::sort()`。

**结论**：在 STL 中，如果全局算法的迭代器要求高于容器所能提供的迭代器类型，容器就必须自己实现 OOP 版本的专属方法。

---

## 五、 操作符重载

### 1. 基本概念

**操作符重载**是C++的一个特性，允许用户为自定义类型定义操作符的行为。

### 2. 重载规则

- 不能重载的操作符：`.`、`.*`、`::`、`? :`、`sizeof`、`typeid`、`const_cast`、`dynamic_cast`、`reinterpret_cast`、`static_cast`
- 重载操作符必须有一个操作数是用户定义类型
- 重载操作符不能改变操作符的优先级和结合性
- 重载操作符不能改变操作数的个数

### 3. 重载方式

#### 3.1 成员函数重载

```cpp
class Complex {
public:
    double real, imag;
    
    Complex(double r = 0, double i = 0) : real(r), imag(i) {}
    
    // 成员函数重载+
    Complex operator+(const Complex& other) const {
        return Complex(real + other.real, imag + other.imag);
    }
    
    // 成员函数重载=
    Complex& operator=(const Complex& other) {
        if (this != &other) {
            real = other.real;
            imag = other.imag;
        }
        return *this;
    }
};
```

#### 3.2 非成员函数重载

```cpp
// 非成员函数重载<<
ostream& operator<<(ostream& os, const Complex& c) {
    os << "(" << c.real << ", " << c.imag << ")";
    return os;
}

// 非成员函数重载+
Complex operator+(const Complex& c1, const Complex& c2) {
    return Complex(c1.real + c2.real, c1.imag + c2.imag);
}
```

### 4. 常用操作符重载

#### 4.1 算术操作符

```cpp
class Vector {
public:
    double x, y;
    
    Vector(double x = 0, double y = 0) : x(x), y(y) {}
    
    // 加法
    Vector operator+(const Vector& other) const {
        return Vector(x + other.x, y + other.y);
    }
    
    // 减法
    Vector operator-(const Vector& other) const {
        return Vector(x - other.x, y - other.y);
    }
    
    // 乘法（标量）
    Vector operator*(double scalar) const {
        return Vector(x * scalar, y * scalar);
    }
    
    // 除法（标量）
    Vector operator/(double scalar) const {
        return Vector(x / scalar, y / scalar);
    }
};

// 标量乘法（左操作数是标量）
Vector operator*(double scalar, const Vector& v) {
    return Vector(v.x * scalar, v.y * scalar);
}
```

#### 4.2 关系操作符

```cpp
class Student {
public:
    string name;
    int score;
    
    Student(string n, int s) : name(n), score(s) {}
    
    // 小于
    bool operator<(const Student& other) const {
        return score < other.score;
    }
    
    // 等于
    bool operator==(const Student& other) const {
        return name == other.name && score == other.score;
    }
};
```

#### 4.3 下标操作符

```cpp
class Array {
private:
    int* data;
    int size;
public:
    Array(int s) : size(s) {
        data = new int[size];
    }
    
    ~Array() {
        delete[] data;
    }
    
    // 下标操作符
    int& operator[](int index) {
        if (index < 0 || index >= size) {
            throw out_of_range("Index out of range");
        }
        return data[index];
    }
    
    // 常量版本下标操作符
    const int& operator[](int index) const {
        if (index < 0 || index >= size) {
            throw out_of_range("Index out of range");
        }
        return data[index];
    }
};
```

#### 4.4 递增/递减操作符

```cpp
class Counter {
private:
    int value;
public:
    Counter(int v = 0) : value(v) {}
    
    // 前置递增
    Counter& operator++() {
        ++value;
        return *this;
    }
    
    // 后置递增
    Counter operator++(int) {
        Counter temp = *this;
        ++value;
        return temp;
    }
    
    // 前置递减
    Counter& operator--() {
        --value;
        return *this;
    }
    
    // 后置递减
    Counter operator--(int) {
        Counter temp = *this;
        --value;
        return temp;
    }
};
```

---

## 六、 模板（Templates）

### 1. 函数模板

#### 1.1 基本语法

```cpp
template <typename T>
T max(T a, T b) {
    return a > b ? a : b;
}

// 使用
auto result = max(10, 20); // T = int
auto result2 = max(3.14, 2.71); // T = double
```

#### 1.2 模板参数推导

编译器会根据函数参数自动推导模板参数类型：

```cpp
// 自动推导T为int
int a = 10, b = 20;
auto result = max(a, b);

// 自动推导T为double
double x = 3.14, y = 2.71;
auto result2 = max(x, y);
```

#### 1.3 显式指定模板参数

```cpp
// 显式指定T为double
auto result = max<double>(10, 20.5); // 10会被转换为double
```

### 2. 类模板

#### 2.1 基本语法

```cpp
template <typename T>
class Stack {
private:
    vector<T> elements;
public:
    void push(const T& item) {
        elements.push_back(item);
    }
    
    void pop() {
        if (!elements.empty()) {
            elements.pop_back();
        }
    }
    
    T& top() {
        return elements.back();
    }
    
    bool empty() const {
        return elements.empty();
    }
};

// 使用
Stack<int> intStack;
intStack.push(10);
intStack.push(20);

Stack<string> stringStack;
stringStack.push("Hello");
stringStack.push("World");
```

#### 2.2 模板特化

##### 2.2.1 全特化

```cpp
// 主模板
template <typename T>
class MyClass {
public:
    void print() {
        cout << "General template" << endl;
    }
};

// 全特化（针对int类型）
template <>
class MyClass<int> {
public:
    void print() {
        cout << "Specialized for int" << endl;
    }
};

// 使用
MyClass<double> d; // 使用主模板
d.print(); // 输出：General template

MyClass<int> i; // 使用特化版本
i.print(); // 输出：Specialized for int
```

##### 2.2.2 偏特化 (Partial Specialization)

侯捷老师强调，偏特化在底层设计上有两个维度的“偏”：
1. **个数的偏**：模板有多个参数，我只绑定其中一部分。例如 `template <typename T1, typename T2>`，我特化为 `template <typename T2> class MyClass<int, T2>`（锁定第一个参数为 int）。
2. **范围的偏**：模板参数的个数没变，但我缩小了参数的接收范围。例如主模板接收任意 `T`，我特化一个版本专门接收指针类型 `T*`，或者专门接收常量指针 `const T*`。这在 STL 中被大量用于萃取器（Traits）机制，用来区分传入的是普通对象还是原生指针。

```cpp
// 主模板
template <typename T1, typename T2>
class MyClass {
public:
    void print() {
        cout << "General template" << endl;
    }
};

// 偏特化（第一个参数为int）
template <typename T2>
class MyClass<int, T2> {
public:
    void print() {
        cout << "Specialized for int and " << typeid(T2).name() << endl;
    }
};

// 偏特化（两个参数都是指针）
template <typename T1, typename T2>
class MyClass<T1*, T2*> {
public:
    void print() {
        cout << "Specialized for pointers" << endl;
    }
};

// 使用
MyClass<double, string> d; // 使用主模板
d.print(); // 输出：General template

MyClass<int, string> i; // 使用第一个偏特化
i.print(); // 输出：Specialized for int and class std::basic_string<char,struct std::char_traits<char>,class std::allocator<char> >

MyClass<int*, double*> p; // 使用第二个偏特化
p.print(); // 输出：Specialized for pointers
```

### 3. 可变参数模板 (C++11)

```cpp
// 递归终止函数
void print() {
    cout << endl;
}

// 可变参数模板
template <typename T, typename... Args>
void print(T first, Args... rest) {
    cout << first << " ";
    print(rest...); // 递归调用
}

// 使用
print(1, 2, 3, "Hello", 3.14); // 输出：1 2 3 Hello 3.14 
```

### 4. 模板的模板参数

```cpp
template <typename T, template <typename> class Container>
class MyContainer {
private:
    Container<T> container;
public:
    void push(const T& item) {
        container.push_back(item);
    }
    
    T& get(int index) {
        return container[index];
    }
};

// 使用
MyContainer<int, vector> intVector;
intVector.push(10);
intVector.push(20);
cout << intVector.get(0) << endl; // 输出：10
```

---

## 七、 代码示例与实践

### 1. 容器与算法综合应用

```cpp
#include <iostream>
#include <vector>
#include <algorithm>
#include <functional>
#include <string>

using namespace std;

// 自定义比较函数
bool compareDescending(int a, int b) {
    return a > b;
}

// 自定义仿函数
class EvenPredicate {
public:
    bool operator()(int x) const {
        return x % 2 == 0;
    }
};

int main() {
    // 创建并初始化vector
    vector<int> v = {5, 2, 8, 1, 9, 3, 7, 4, 6};
    
    // 排序（默认升序）
    sort(v.begin(), v.end());
    cout << "Sorted in ascending order: ";
    for (auto i : v) cout << i << " ";
    cout << endl;
    
    // 排序（降序）
    sort(v.begin(), v.end(), compareDescending);
    cout << "Sorted in descending order: ";
    for (auto i : v) cout << i << " ";
    cout << endl;
    
    // 排序（使用lambda）
    sort(v.begin(), v.end(), [](int a, int b) { return a < b; });
    cout << "Sorted using lambda: ";
    for (auto i : v) cout << i << " ";
    cout << endl;
    
    // 查找
    auto it = find(v.begin(), v.end(), 5);
    if (it != v.end()) {
        cout << "Found 5 at position: " << distance(v.begin(), it) << endl;
    }
    
    // 计数
    int evenCount = count_if(v.begin(), v.end(), EvenPredicate());
    cout << "Number of even numbers: " << evenCount << endl;
    
    // 使用lambda计数
    int oddCount = count_if(v.begin(), v.end(), [](int x) { return x % 2 != 0; });
    cout << "Number of odd numbers: " << oddCount << endl;
    
    // 变换
    vector<int> v2(v.size());
    transform(v.begin(), v.end(), v2.begin(), [](int x) { return x * 2; });
    cout << "Transformed vector (x2): ";
    for (auto i : v2) cout << i << " ";
    cout << endl;
    
    // 复制_if
    vector<int> v3;
    copy_if(v.begin(), v.end(), back_inserter(v3), [](int x) { return x > 5; });
    cout << "Numbers greater than 5: ";
    for (auto i : v3) cout << i << " ";
    cout << endl;
    
    return 0;
}
```

**执行结果：**
```
Sorted in ascending order: 1 2 3 4 5 6 7 8 9 
Sorted in descending order: 9 8 7 6 5 4 3 2 1 
Sorted using lambda: 1 2 3 4 5 6 7 8 9 
Found 5 at position: 4
Number of even numbers: 4
Number of odd numbers: 5
Transformed vector (x2): 2 4 6 8 10 12 14 16 18 
Numbers greater than 5: 6 7 8 9 
```

### 2. 模板与操作符重载综合应用

```cpp
#include <iostream>

using namespace std;

// 模板类：三维向量
template <typename T>
class Vector3D {
public:
    T x, y, z;
    
    Vector3D(T x = 0, T y = 0, T z = 0) : x(x), y(y), z(z) {}
    
    // 加法操作符重载
    Vector3D operator+(const Vector3D& other) const {
        return Vector3D(x + other.x, y + other.y, z + other.z);
    }
    
    // 减法操作符重载
    Vector3D operator-(const Vector3D& other) const {
        return Vector3D(x - other.x, y - other.y, z - other.z);
    }
    
    // 标量乘法操作符重载
    Vector3D operator*(T scalar) const {
        return Vector3D(x * scalar, y * scalar, z * scalar);
    }
    
    // 点积操作符重载
    T operator*(const Vector3D& other) const {
        return x * other.x + y * other.y + z * other.z;
    }
    
    // 输出操作符重载
    friend ostream& operator<<(ostream& os, const Vector3D& v) {
        os << "(" << v.x << ", " << v.y << ", " << v.z << ")";
        return os;
    }
};

// 标量乘法（左操作数是标量）
template <typename T>
Vector3D<T> operator*(T scalar, const Vector3D<T>& v) {
    return Vector3D<T>(v.x * scalar, v.y * scalar, v.z * scalar);
}

int main() {
    Vector3D<int> v1(1, 2, 3);
    Vector3D<int> v2(4, 5, 6);
    
    cout << "v1 = " << v1 << endl;
    cout << "v2 = " << v2 << endl;
    cout << "v1 + v2 = " << v1 + v2 << endl;
    cout << "v1 - v2 = " << v1 - v2 << endl;
    cout << "v1 * 2 = " << v1 * 2 << endl;
    cout << "2 * v1 = " << 2 * v1 << endl;
    cout << "v1 * v2 = " << v1 * v2 << endl;
    
    Vector3D<double> v3(1.1, 2.2, 3.3);
    Vector3D<double> v4(4.4, 5.5, 6.6);
    
    cout << "\nv3 = " << v3 << endl;
    cout << "v4 = " << v4 << endl;
    cout << "v3 + v4 = " << v3 + v4 << endl;
    cout << "v3 - v4 = " << v3 - v4 << endl;
    cout << "v3 * 2.5 = " << v3 * 2.5 << endl;
    cout << "2.5 * v3 = " << 2.5 * v3 << endl;
    cout << "v3 * v4 = " << v3 * v4 << endl;
    
    return 0;
}
```

**执行结果：**
```
v1 = (1, 2, 3)
v2 = (4, 5, 6)
v1 + v2 = (5, 7, 9)
v1 - v2 = (-3, -3, -3)
v1 * 2 = (2, 4, 6)
2 * v1 = (2, 4, 6)
v1 * v2 = 32

v3 = (1.1, 2.2, 3.3)
v4 = (4.4, 5.5, 6.6)
v3 + v4 = (5.5, 7.7, 9.9)
v3 - v4 = (-3.3, -3.3, -3.3)
v3 * 2.5 = (2.75, 5.5, 8.25)
2.5 * v3 = (2.75, 5.5, 8.25)
v3 * v4 = 39.6
```

---

## 八、 面试高频考点

### 1. 灵魂拷问：STL容器的分类及特点

**详细回答：**
STL容器分为四类：

1. **序列容器**：vector（动态数组，快速随机访问）、list（双向链表，快速插入删除）、deque（双端队列，两端快速操作）、array（固定大小数组）、forward_list（单向链表）

2. **关联容器**：set（有序集合，元素唯一）、multiset（有序集合，元素可重复）、map（有序键值对，键唯一）、multimap（有序键值对，键可重复）

3. **无序关联容器**：unordered_set（无序集合，基于哈希表）、unordered_multiset（无序集合，基于哈希表，元素可重复）、unordered_map（无序键值对，基于哈希表）、unordered_multimap（无序键值对，基于哈希表，键可重复）

4. **容器适配器**：stack（栈，后进先出）、queue（队列，先进先出）、priority_queue（优先队列，基于堆）

### 2. 灵魂拷问：OOP和GP的区别

**详细回答：**
- **OOP（面向对象编程）**：将数据和操作封装在一起，通过继承和多态实现代码复用。运行时多态，有虚函数调用开销，灵活性高。

- **GP（泛型编程）**：将算法和数据结构分离，通过模板实现代码复用。编译时多态，无运行时开销，性能好。

- **STL的应用**：STL是泛型编程的典范，通过迭代器连接算法和容器，实现高度解耦和复用。

### 3. 灵魂拷问：操作符重载的规则

**详细回答：**
- 不能重载的操作符：`.`、`.*`、`::`、`? :`、`sizeof`、`typeid`、`const_cast`、`dynamic_cast`、`reinterpret_cast`、`static_cast`
- 重载操作符必须有一个操作数是用户定义类型
- 重载操作符不能改变操作符的优先级和结合性
- 重载操作符不能改变操作数的个数
- 可以作为成员函数或非成员函数重载

### 4. 灵魂拷问：模板特化的作用

**详细回答：**
模板特化允许为特定类型提供定制化的实现，分为全特化和偏特化：

- **全特化**：为所有模板参数提供具体类型，完全覆盖主模板
- **偏特化**：为部分模板参数提供具体类型，部分覆盖主模板

特化的作用：
1. 为特定类型提供优化实现
2. 处理特殊类型的特殊逻辑
3. 解决模板无法处理的类型

### 5. 灵魂拷问：allocator的作用

**详细回答：**
allocator是STL中的内存管理组件，负责：

1. **内存分配与释放**：为容器分配和释放内存
2. **对象构造与析构**：在分配的内存上构造和析构对象
3. **自定义内存管理**：允许用户提供自定义分配器，如内存池，以提高性能

在性能要求极高的场景（如高频交易、大模型算子），自定义分配器可以显著提高内存管理效率，避免标准分配器的系统调用开销和内存碎片问题。

---

## 九、 总结

本章节介绍了STL的核心组件和泛型编程的基础，包括：

- **STL容器**：序列容器、关联容器、无序关联容器和容器适配器，以及它们的特点和使用场景
- **分配器**：内存管理组件，负责容器的内存分配和释放，支持自定义分配器以提高性能
- **STL源代码分布**：不同编译器中STL源码的位置和结构
- **OOP vs GP**：面向对象编程和泛型编程的区别与应用
- **操作符重载**：允许为自定义类型定义操作符行为的特性
- **模板**：函数模板、类模板、模板特化和可变参数模板

STL是C++编程中非常重要的工具，掌握STL的使用和原理，对于提高编程效率和代码质量至关重要。泛型编程是现代C++的核心特性之一，通过模板和STL，我们可以编写更加灵活、高效和可维护的代码。

在后续的学习中，我们将深入探讨STL各个组件的具体实现和高级应用，进一步提升C++编程能力。