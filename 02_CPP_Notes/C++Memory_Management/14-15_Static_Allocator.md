# 侯捷 C++ 课程 14-15 集学习笔记：Static Allocator（静态分配器）

## 一、课程概述

### 1.1 课程主题

本课程深入讲解 **static allocator**（静态分配器），将内存分配器封装为独立的 `allocator` 类，通过静态成员的方式为应用类提供高效的内存管理。这比之前的 per-class allocator 更加优雅和可复用。

### 1.2 课程内容结构

| 集数 | 主题 | 核心内容 |
|------|------|---------|
| 第14集 | Static allocator | 将 allocator 封装为独立类，应用类通过 static 成员使用 |
| 第15集 | Macro for static allocator | 使用宏定义简化代码，提高复用性 |

### 1.3 学习目标

1. 理解 static allocator 的设计思想
2. 掌握 allocator 类的完整实现
3. 理解 free list 的工作原理
4. 掌握宏定义的使用技巧
5. 了解 G2.9 std::alloc 的设计雏形
6. 掌握面试中相关的高频问题

### 1.4 参考来源

- C++ Primer 3/e, p.765
- Effective C++ 2/e, item 10
- SGI STL 源码分析

---

## 二、Static Allocator 核心概念

### 2.1 为什么需要 Static Allocator？

#### 2.1.1 Per-class Allocator 的问题

回顾之前的 per-class allocator（如 Airplane 类）：

```cpp
class Airplane {
    union {
        AirplaneRep rep;
        Airplane* next;
    };
    static Airplane* headOfFreeList;  // 每个类都要自己写！
    // ...
};
```

**问题：**
1. **代码重复**：每个需要自定义内存管理的类都要重写类似的代码
2. **耦合度高**：应用类和内存分配逻辑混杂在一起
3. **难以复用**：内存分配的逻辑无法在多个类之间共享

#### 2.1.2 Static Allocator 的优势

**核心思想：将内存分配器封装为独立的类！**

```cpp
class allocator {
private:
    struct obj {
        struct obj* next;  // 嵌入式指针
    };
    obj* freeStore = nullptr;
    const int CHUNK = 5;
public:
    void* allocate(size_t);
    void deallocate(void*, size_t);
};
```

**优势：**
1. **代码解耦**：应用类不再与内存分配细节纠缠
2. **易于复用**：多个类可以共享同一个 allocator
3. **设计更干净**：职责分离，各司其职

### 2.2 设计理念对比

```
┌─────────────────────────────────────────────────────────────┐
│                   Per-class Allocator                       │
├─────────────────────────────────────────────────────────────┤
│ class Airplane {                                           │
│     // 应用逻辑 + 内存管理逻辑混在一起                       │
│     AirplaneRep rep;                                       │
│     static Airplane* headOfFreeList;  // 内存管理           │
│     static void* operator new(...);   // 内存管理           │
│     static void operator delete(...); // 内存管理           │
│ };                                                          │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│                    Static Allocator                         │
├─────────────────────────────────────────────────────────────┤
│ // 内存管理逻辑独立出来                                     │
│ class allocator {                                          │
│     void* allocate(size_t);                                 │
│     void deallocate(void*, size_t);                         │
│ };                                                          │
│                                                             │
│ // 应用类只关注业务逻辑                                     │
│ class Foo {                                                 │
│     long L;                                                 │
│     string str;                                             │
│     static allocator myAlloc;  // 复用 allocator             │
│ };                                                          │
└─────────────────────────────────────────────────────────────┘
```

---

## 三、Allocator 类的完整实现

### 3.1 核心数据结构

```cpp
#include <cstddef>
#include <iostream>
#include <cstdlib>  // for malloc/free
using namespace std;

// ==========================================
// allocator 类 - 独立的内存分配器
// ==========================================
class allocator {
private:
    // ==========================================
    // 内部节点结构：嵌入式指针
    // ==========================================
    struct obj {
        struct obj* next;  // 用于链接空闲块的指针
    };
    
private:
    // ==========================================
    // 静态成员：自由链表的头指针
    // ==========================================
    obj* freeStore = nullptr;
    
    // ==========================================
    // 每次分配的块数（CHUNK_SIZE）
    // 截图中是 5，用于演示方便观察
    // ==========================================
    const int CHUNK_SIZE = 5;

public:
    // ==========================================
    // 分配内存
    // ==========================================
    void* allocate(size_t size);
    
    // ==========================================
    // 释放内存
    // ==========================================
    void deallocate(void* p, size_t size);
};
```

### 3.2 allocate 方法实现

```cpp
void* allocator::allocate(size_t size) {
    obj* p = nullptr;
    
    // ==========================================
    // 情况1：自由链表为空，需要申请一大块内存
    // ==========================================
    if (!freeStore) {
        // 计算需要分配的总大小
        // CHUNK_SIZE 个对象，每个对象大小为 size
        size_t chunk = CHUNK_SIZE * size;
        
        // 申请一大块内存（用 malloc 避免调用构造函数）
        freeStore = p = reinterpret_cast<obj*>(malloc(chunk));
        
        // ==========================================
        // 将这一大块内存分割成多个小块，串成链表
        // ==========================================
        // 遍历每个小块，设置 next 指针
        for (int i = 0; i < CHUNK_SIZE - 1; ++i) {
            // 当前块的 next 指向下一个块
            // 需要跳过当前块的大小（size）才能到达下一块
            p->next = reinterpret_cast<obj*>(
                reinterpret_cast<char*>(p) + size
            );
            p = p->next;
        }
        
        // 最后一个块的 next 设为 nullptr
        p->next = nullptr;
    }
    
    // ==========================================
    // 情况2：自由链表不为空，直接取第一个块
    // ==========================================
    // 记录要返回的块
    p = freeStore;
    
    // 链表头向后移动一位
    freeStore = freeStore->next;
    
    return p;
}
```

### 3.3 deallocate 方法实现

```cpp
void allocator::deallocate(void* p, size_t size) {
    // 空指针检查
    if (p == nullptr) {
        return;
    }
    
    // ==========================================
    // 将释放的块插回自由链表的前端
    // ==========================================
    // 为什么插在前端？因为 O(1) 时间复杂度，最快！
    static_cast<obj*>(p)->next = freeStore;
    freeStore = static_cast<obj*>(p);
}
```

### 3.4 代码分析

#### 3.4.1 嵌入式指针的理解

**嵌入式指针（Embedded Pointer）** 是指在空闲内存块内部使用一部分空间来存储链表指针。

```
┌─────────────────────────────────────────────┐
│ 空闲块在 free list 上时：                   │
│ ┌─────────────────────────────────────┐    │
│ │         obj 结构                    │    │
│ │  ┌─────────────────────────────┐    │    │
│ │  │ next: 指向下一个空闲块的指针 │    │    │
│ │  └─────────────────────────────┘    │    │
│ └─────────────────────────────────────┘    │
│                 size 字节                    │
└─────────────────────────────────────────────┘

┌─────────────────────────────────────────────┐
│ 块被分配出去后：                           │
│ ┌─────────────────────────────────────┐    │
│ │         用户数据（如 Foo 对象）       │    │
│ │  ┌─────────────────────────────┐    │    │
│ │  │ long L;                     │    │    │
│ │  │ string str;                 │    │    │
│ │  └─────────────────────────────┘    │    │
│ └─────────────────────────────────────┘    │
│                 size 字节                    │
└─────────────────────────────────────────────┘
```

**关键点：**
- 块在空闲时，`next` 指针占用空间
- 块被分配后，用户可以使用完整的 `size` 字节
- 不需要额外的 cookie，节省内存！

#### 3.4.2 链表操作图解

**allocate 时的链表操作：**

```
初始状态（链表为空）：
freeStore → nullptr

申请大块后：
freeStore → [块0] → [块1] → [块2] → [块3] → [块4] → nullptr

第一次分配（返回块0）：
freeStore → [块1] → [块2] → [块3] → [块4] → nullptr
返回: [块0]

第二次分配（返回块1）：
freeStore → [块2] → [块3] → [块4] → nullptr
返回: [块1]
```

**deallocate 时的链表操作：**

```
当前链表：
freeStore → [块2] → [块3] → [块4] → nullptr

释放块1：
freeStore → [块1] → [块2] → [块3] → [块4] → nullptr

释放块0：
freeStore → [块0] → [块1] → [块2] → [块3] → [块4] → nullptr
```

#### 3.4.3 初学者引导

**为什么要用 malloc 而不是 new？**

```cpp
freeStore = p = reinterpret_cast<obj*>(malloc(chunk));
// 而不是
// freeStore = p = new obj[chunk / sizeof(obj)];
```

**原因：**
- `new` 会调用构造函数，而我们只是需要原始内存
- `malloc` 只分配内存，不调用构造函数
- 稍后用户会用 placement new 构造对象（如果需要的话）

**为什么 CHUNK_SIZE 设置为 5？**

- 截图中设置为 5 是为了方便演示和观察
- 实际项目中通常设置为更大的值（如 24、64、512 等）
- 值越大，malloc 调用次数越少，但内存可能浪费更多

---

## 四、应用类如何使用 Static Allocator

### 4.1 Foo 类示例

```cpp
// ==========================================
// Foo 类 - 使用 static allocator
// ==========================================
class Foo {
public:
    // 数据成员
    long L;
    string str;
    
    // ==========================================
    // 静态 allocator 成员
    // ==========================================
    static allocator myAlloc;
    
public:
    // 构造函数
    Foo(long l) : L(l) {}
    
    // ==========================================
    // 重载 operator new
    // ==========================================
    static void* operator new(size_t size) {
        // 委托给 allocator 处理
        return myAlloc.allocate(size);
    }
    
    // ==========================================
    // 重载 operator delete
    // ==========================================
    static void operator delete(void* p, size_t size) {
        // 委托给 allocator 处理
        return myAlloc.deallocate(p, size);
    }
};

// ==========================================
// 静态成员初始化
// ==========================================
allocator Foo::myAlloc;
```

### 4.2 Goo 类示例

```cpp
// ==========================================
// Goo 类 - 另一个使用 static allocator 的类
// ==========================================
class Goo {
public:
    // 数据成员（与 Foo 不同）
    complex<double> c;
    string str;
    
    // ==========================================
    // 静态 allocator 成员
    // ==========================================
    static allocator myAlloc;
    
public:
    // 构造函数
    Goo(const complex<double>& x) : c(x) {}
    
    // ==========================================
    // 重载 operator new
    // ==========================================
    static void* operator new(size_t size) {
        return myAlloc.allocate(size);
    }
    
    // ==========================================
    // 重载 operator delete
    // ==========================================
    static void operator delete(void* p, size_t size) {
        return myAlloc.deallocate(p, size);
    }
};

// ==========================================
// 静态成员初始化
// ==========================================
allocator Goo::myAlloc;
```

### 4.3 代码分析

#### 4.3.1 设计优势

**对比之前的 per-class allocator：**

| 特性 | Per-class Allocator | Static Allocator |
|------|---------------------|------------------|
| 代码重复 | 每个类都要写 operator new/delete | 只需要声明 static allocator |
| 逻辑复用 | 无法复用 | allocator 类可被多个类共享 |
| 耦合度 | 高（业务+内存混在一起） | 低（职责分离） |
| 扩展性 | 差 | 好（可扩展 allocator） |

#### 4.3.2 每个类有独立的 free list

**重要特点：**
- 每个类有自己的 `static allocator myAlloc`
- 不同类的 allocator 对象维护不同的 free list
- Foo 的对象和 Goo 的对象不会混在同一个 free list 中

```
Foo::myAlloc.freeStore → [Foo对象] → [Foo对象] → nullptr
Goo::myAlloc.freeStore → [Goo对象] → [Goo对象] → nullptr
```

**为什么要分开？**
- Foo 和 Goo 的大小可能不同
- 如果混在一起，分配时可能拿到大小不匹配的块
- 类型安全！

---

## 五、使用宏定义简化代码

### 5.1 宏定义的必要性

观察 Foo 和 Goo 类，它们的 operator new/delete 实现几乎相同：

```cpp
// Foo 类
static void* operator new(size_t size) {
    return myAlloc.allocate(size);
}
static void operator delete(void* p, size_t size) {
    return myAlloc.deallocate(p, size);
}

// Goo 类（几乎一样）
static void* operator new(size_t size) {
    return myAlloc.allocate(size);
}
static void operator delete(void* p, size_t size) {
    return myAlloc.deallocate(p, size);
}
```

**问题：代码重复！** 如果有 100 个类需要使用 allocator，就要写 100 次相同的代码。

### 5.2 宏定义的设计

```cpp
// ==========================================
// DECLARE_POOL_ALLOC - 在类定义中使用
// ==========================================
#define DECLARE_POOL_ALLOC() \
public: \
    static void* operator new(size_t size) { \
        return myAlloc.allocate(size); \
    } \
    static void operator delete(void* p) { \
        myAlloc.deallocate(p, 0); \
    } \
protected: \
    static allocator myAlloc;

// ==========================================
// IMPLEMENT_POOL_ALLOC - 在实现文件中使用
// ==========================================
#define IMPLEMENT_POOL_ALLOC(class_name) \
    allocator class_name::myAlloc;
```

### 5.3 使用宏定义简化后的类

```cpp
// ==========================================
// 使用宏定义后的 Foo 类
// ==========================================
class Foo {
    DECLARE_POOL_ALLOC()  // 一行搞定！
    
public:
    long L;
    string str;
    
    Foo(long l) : L(l) {}
};
IMPLEMENT_POOL_ALLOC(Foo)  // 在实现文件中

// ==========================================
// 使用宏定义后的 Goo 类
// ==========================================
class Goo {
    DECLARE_POOL_ALLOC()  // 一行搞定！
    
public:
    complex<double> c;
    string str;
    
    Goo(const complex<double>& x) : c(x) {}
};
IMPLEMENT_POOL_ALLOC(Goo)  // 在实现文件中
```

### 5.4 代码分析

#### 5.4.1 宏定义的优点

**优点：**
1. **代码简洁**：一行宏定义代替多行重复代码
2. **一致性**：所有类的实现方式统一
3. **易于维护**：修改宏定义即可修改所有类的行为
4. **减少错误**：避免手写时的拼写错误

#### 5.4.2 初学者注意事项

**宏定义的陷阱：**

1. **宏定义中的分号**：宏定义本身不需要分号，但使用时需要
2. **作用域问题**：宏定义中的 public/protected 会改变后续成员的访问权限
3. **调试困难**：宏展开后的代码可能难以调试

**使用建议：**
- 宏定义应尽量简单
- 复杂逻辑应放在函数中
- 使用 `#undef` 避免宏污染

---

## 六、测试示例与结果分析

### 6.1 测试代码

```cpp
// ==========================================
// 测试代码
// ==========================================
int main() {
    cout << "====== Foo 类测试 ======" << endl;
    cout << "sizeof(Foo) = " << sizeof(Foo) << endl;
    
    // 分配 23 个 Foo 对象（23 = 5*4 + 3，需要分配 5 次大块）
    Foo* p[100];
    for (int i = 0; i < 23; ++i) {
        p[i] = new Foo(i);
        cout << p[i] << ' ' << p[i]->L << endl;
    }
    
    // 释放所有对象
    for (int i = 0; i < 23; ++i) {
        delete p[i];
    }
    
    cout << "\n====== Goo 类测试 ======" << endl;
    cout << "sizeof(Goo) = " << sizeof(Goo) << endl;
    
    // 分配 17 个 Goo 对象
    Goo* q[100];
    for (int i = 0; i < 17; ++i) {
        q[i] = new Goo(complex<double>(i, i));
        cout << q[i] << ' ' << q[i]->c << endl;
    }
    
    // 释放所有对象
    for (int i = 0; i < 17; ++i) {
        delete q[i];
    }
    
    return 0;
}
```

### 6.2 输出示例

```
====== Foo 类测试 ======
sizeof(Foo) = 8
0x3e5f90 0
0x3e5f98 1
0x3e5fa0 2
0x3e5fa8 3
0x3e5fb0 4
0x3e5fb8 5
0x3e5fc0 6
0x3e5fc8 7
0x3e5fd0 8
0x3e5fd8 9
...

====== Goo 类测试 ======
sizeof(Goo) = 24
0x3e6200 (0,0)
0x3e6218 (1,1)
0x3e6230 (2,2)
0x3e6248 (3,3)
0x3e6260 (4,4)
0x3e6278 (5,5)
...
```

### 6.3 结果分析

**观察地址间隔：**

```
Foo 对象地址间隔：
0x3e5f90 → 0x3e5f98 → 0x3e5fa0 → ...
间隔 = 8（sizeof(Foo)）

Goo 对象地址间隔：
0x3e6200 → 0x3e6218 → 0x3e6230 → ...
间隔 = 24（sizeof(Goo)）
```

**结论：**
- 没有 cookie！地址间隔正好是对象大小
- 内存紧凑，没有额外开销
- 不同类的对象存储在不同的内存区域

---

## 七、多 Free-list 的全局 Allocator（G2.9 std::alloc 雏形）

### 7.1 设计思想

之前的 allocator 只有一个 free list，只能处理固定大小的对象。**SGI STL 的 std::alloc** 进一步发展为具有 **16 个 free lists** 的全局分配器。

### 7.2 架构设计

```
┌─────────────────────────────────────────────────────────────────┐
│              global allocator (16 个 free lists)                │
├─────────────────────────────────────────────────────────────────┤
│  #0: 8 字节对象的 free list  → [块] → [块] → [块] → nullptr   │
│  #1: 16 字节对象的 free list → [块] → [块] → nullptr          │
│  #2: 24 字节对象的 free list → [块] → [块] → [块] → [块] → ...│
│  #3: 32 字节对象的 free list → [块] → nullptr                  │
│  ...                                                           │
│ #15: 128 字节对象的 free list → [块] → [块] → ...              │
└─────────────────────────────────────────────────────────────────┘
```

### 7.3 分配策略

**分配时：**
1. 根据对象大小选择对应的 free list
2. 如果该 free list 不为空，直接取一块
3. 如果为空，向系统申请一大块内存，分割后加入 free list

**释放时：**
1. 根据对象大小找到对应的 free list
2. 将块插回该 free list 的前端

### 7.4 优势分析

| 特性 | 单 free list | 多 free lists |
|------|-------------|--------------|
| 适用对象大小 | 固定大小 | 多种大小（8~128 字节） |
| 内存利用率 | 较低（可能浪费） | 较高（按需分配） |
| 复杂度 | 简单 | 较复杂 |
| 适用场景 | 单一类型对象 | 多种类型对象 |

### 7.5 SGI STL std::alloc 的特点

**截图中的关键信息：**
> "將前述 allocator 進一步發展為具備 16 條 free-lists，並因此不再以 application classes 內的 static 呈現，而是一個 global allocator —— 這就是 G2.9 的 std::alloc 的雛形。"

**std::alloc 的特点：**
1. **全局共享**：不再是每个类一个 allocator
2. **16 个 free lists**：对应 8、16、24、...、128 字节
3. **内存池策略**：先从对应 free list 取，取不到再申请大块
4. **内存对齐**：所有块都按 8 字节对齐

---

## 八、C++ 新标准的更新与改进

### 8.1 C++11 的 std::allocator

C++11 标准库提供了更规范的 allocator 接口：

```cpp
#include <memory>

// C++11 allocator 的基本使用
std::allocator<int> alloc;

// 分配内存（不构造对象）
int* p = alloc.allocate(10);

// 构造对象（placement new）
for (int i = 0; i < 10; ++i) {
    alloc.construct(&p[i], i);  // 在 p[i] 处构造 int(i)
}

// 析构对象
for (int i = 0; i < 10; ++i) {
    alloc.destroy(&p[i]);
}

// 释放内存
alloc.deallocate(p, 10);
```

### 8.2 C++17 的 pmr（Polymorphic Memory Resources）

C++17 引入了多态内存资源，提供更灵活的内存管理：

```cpp
#include <memory_resource>
#include <vector>

// 创建一个单调缓冲区资源（类似我们的 allocator）
std::pmr::monotonic_buffer_resource pool;

// 使用该资源的 vector
std::pmr::vector<int> vec(&pool);

// 所有分配都从 pool 中获取
vec.push_back(1);
vec.push_back(2);
vec.push_back(3);
```

**pmr 的优势：**
- 运行时可替换内存资源
- 更好的性能（减少 malloc 调用）
- 便于测试和调试

### 8.3 C++20 的 std::allocator_traits

C++20 增强了 allocator 的类型特征支持：

```cpp
#include <memory>

template <typename T>
void process(std::allocator<T>& alloc) {
    // 使用 allocator_traits 获取类型信息
    using traits = std::allocator_traits<decltype(alloc)>;
    
    // 分配
    T* p = traits::allocate(alloc, 10);
    
    // 构造
    for (size_t i = 0; i < 10; ++i) {
        traits::construct(alloc, &p[i], i);
    }
    
    // 析构和释放
    for (size_t i = 0; i < 10; ++i) {
        traits::destroy(alloc, &p[i]);
    }
    traits::deallocate(alloc, p, 10);
}
```

---

## 九、面试高频问题与解答

### 9.1 基础问题

#### Q1: Static allocator 是什么？相比 per-class allocator 有什么改进？

**答案：**

**定义：**
Static allocator 是将内存分配器封装为独立的类，应用类通过静态成员使用该分配器。

**改进：**

| 特性 | Per-class Allocator | Static Allocator |
|------|---------------------|------------------|
| 代码复用 | 每个类都要重写分配逻辑 | allocator 类可被多个类共享 |
| 耦合度 | 业务逻辑与分配逻辑混杂 | 职责分离，代码更干净 |
| 可维护性 | 修改分配逻辑需要修改每个类 | 只需修改 allocator 类 |
| 扩展性 | 差 | 好（可轻松扩展新功能） |

#### Q2: allocator 类中的 freeStore 是什么？如何工作？

**答案：**

**定义：**
`freeStore` 是指向自由链表头的指针，用于管理空闲内存块。

**工作原理：**

```cpp
// 分配时：从链表头取一块
p = freeStore;
freeStore = freeStore->next;

// 释放时：插回链表头
static_cast<obj*>(p)->next = freeStore;
freeStore = static_cast<obj*>(p);
```

**图解：**

```
分配前：freeStore → [块0] → [块1] → [块2] → nullptr
分配后：freeStore → [块1] → [块2] → nullptr，返回 [块0]

释放前：freeStore → [块1] → [块2] → nullptr
释放后：freeStore → [块0] → [块1] → [块2] → nullptr
```

#### Q3: 为什么释放时要插在链表前端而不是后端？

**答案：**

**原因：**
- **时间复杂度**：插在前端是 O(1)，插在后端需要遍历链表 O(n)
- **实现简单**：不需要维护尾指针
- **局部性原理**：最近释放的块可能很快又被分配，放在前端更容易被取到

### 9.2 进阶问题

#### Q4: 嵌入式指针（embedded pointer）是什么？有什么好处？

**答案：**

**定义：**
嵌入式指针是指在空闲内存块内部使用一部分空间来存储链表指针，而不是额外分配空间。

**好处：**

```cpp
// 嵌入式指针（好）
struct obj {
    struct obj* next;  // 指针存在块内部
};

// 非嵌入式（不好，浪费空间）
struct Block {
    void* data;       // 实际数据
    struct Block* next; // 额外的指针，浪费空间
};
```

**优点：**
1. **节省内存**：不需要额外的指针空间
2. **内存紧凑**：没有额外开销
3. **实现简单**：直接在块内操作

#### Q5: 为什么每个应用类要有自己的 static allocator？

**答案：**

**原因：**
1. **类型安全**：不同类的对象大小可能不同，不能混用
2. **独立性**：一个类的内存管理不影响其他类
3. **灵活性**：每个类可以有不同的分配策略

```cpp
class Foo {  // sizeof = 8
    static allocator myAlloc;
};

class Goo {  // sizeof = 24
    static allocator myAlloc;
};

// 如果共用一个 allocator，分配 Foo 时可能拿到 Goo 大小的块！
```

#### Q6: SGI STL 的 std::alloc 有多少个 free lists？每个管理多大的对象？

**答案：**

**16 个 free lists**，每个管理的对象大小为：

| Free list 编号 | 对象大小 |
|---------------|---------|
| #0 | 8 字节 |
| #1 | 16 字节 |
| #2 | 24 字节 |
| ... | ... |
| #15 | 128 字节 |

**特点：**
- 每个 free list 管理的对象大小相差 8 字节
- 所有对象都按 8 字节对齐
- 大于 128 字节的对象直接使用 malloc

### 9.3 实际问题

#### Q7: 什么时候应该使用自定义 allocator？

**答案：**

**适用场景：**
1. **大量小对象频繁分配/释放**：如游戏中的粒子系统
2. **性能敏感场景**：如高频交易系统
3. **内存受限环境**：如嵌入式系统
4. **需要内存池**：如对象池模式

**不适用场景：**
1. **一般业务代码**：增加复杂度，得不偿失
2. **大对象**：直接用 malloc/free 更简单
3. **一次性分配**：不需要内存池

#### Q8: 自定义 allocator 需要注意什么？

**答案：**

**注意事项：**
1. **线程安全**：多线程环境下需要加锁
2. **内存对齐**：确保对象正确对齐
3. **异常安全**：分配失败时的处理
4. **继承处理**：子类对象大小可能不同
5. **内存泄漏**：程序结束前是否需要清理

**线程安全示例：**
```cpp
#include <mutex>

class thread_safe_allocator {
private:
    std::mutex mutex_;
    obj* freeStore = nullptr;
public:
    void* allocate(size_t size) {
        std::lock_guard<std::mutex> lock(mutex_);
        // ... 分配逻辑
    }
    
    void deallocate(void* p, size_t size) {
        std::lock_guard<std::mutex> lock(mutex_);
        // ... 释放逻辑
    }
};
```

---

## 十、完整可运行代码

```cpp
#include <cstddef>
#include <iostream>
#include <cstdlib>
#include <complex>
#include <string>
using namespace std;

// ==========================================
// allocator 类 - 独立的内存分配器
// ==========================================
class allocator {
private:
    struct obj {
        struct obj* next;
    };
    
    obj* freeStore = nullptr;
    const int CHUNK_SIZE = 5;

public:
    void* allocate(size_t size) {
        obj* p = nullptr;
        
        if (!freeStore) {
            size_t chunk = CHUNK_SIZE * size;
            freeStore = p = reinterpret_cast<obj*>(malloc(chunk));
            
            for (int i = 0; i < CHUNK_SIZE - 1; ++i) {
                p->next = reinterpret_cast<obj*>(
                    reinterpret_cast<char*>(p) + size
                );
                p = p->next;
            }
            p->next = nullptr;
        }
        
        p = freeStore;
        freeStore = freeStore->next;
        return p;
    }
    
    void deallocate(void* p, size_t /*size*/) {
        if (p == nullptr) return;
        static_cast<obj*>(p)->next = freeStore;
        freeStore = static_cast<obj*>(p);
    }
};

// ==========================================
// 宏定义
// ==========================================
#define DECLARE_POOL_ALLOC() \
public: \
    static void* operator new(size_t size) { \
        return myAlloc.allocate(size); \
    } \
    static void operator delete(void* p) { \
        myAlloc.deallocate(p, 0); \
    } \
protected: \
    static allocator myAlloc;

#define IMPLEMENT_POOL_ALLOC(class_name) \
    allocator class_name::myAlloc;

// ==========================================
// Foo 类
// ==========================================
class Foo {
    DECLARE_POOL_ALLOC()
    
public:
    long L;
    string str;
    
    Foo(long l) : L(l) {}
};
IMPLEMENT_POOL_ALLOC(Foo)

// ==========================================
// Goo 类
// ==========================================
class Goo {
    DECLARE_POOL_ALLOC()
    
public:
    complex<double> c;
    string str;
    
    Goo(const complex<double>& x) : c(x) {}
};
IMPLEMENT_POOL_ALLOC(Goo)

// ==========================================
// 测试代码
// ==========================================
int main() {
    cout << "========================================" << endl;
    cout << "     Static Allocator 测试" << endl;
    cout << "========================================" << endl;
    
    // ==========================================
    // 测试 Foo 类
    // ==========================================
    cout << "\n====== Foo 类 ======" << endl;
    cout << "sizeof(Foo) = " << sizeof(Foo) << endl;
    
    const int N_FOO = 23;
    Foo* foo[N_FOO];
    
    cout << "\n分配 " << N_FOO << " 个 Foo 对象：" << endl;
    for (int i = 0; i < N_FOO; ++i) {
        foo[i] = new Foo(i);
        cout << "p[" << i << "] = " << foo[i] 
             << ", L = " << foo[i]->L << endl;
    }
    
    cout << "\n释放 " << N_FOO << " 个 Foo 对象" << endl;
    for (int i = 0; i < N_FOO; ++i) {
        delete foo[i];
    }
    
    // ==========================================
    // 测试 Goo 类
    // ==========================================
    cout << "\n====== Goo 类 ======" << endl;
    cout << "sizeof(Goo) = " << sizeof(Goo) << endl;
    
    const int N_GOO = 17;
    Goo* goo[N_GOO];
    
    cout << "\n分配 " << N_GOO << " 个 Goo 对象：" << endl;
    for (int i = 0; i < N_GOO; ++i) {
        goo[i] = new Goo(complex<double>(i, i));
        cout << "p[" << i << "] = " << goo[i] 
             << ", c = " << goo[i]->c << endl;
    }
    
    cout << "\n释放 " << N_GOO << " 个 Goo 对象" << endl;
    for (int i = 0; i < N_GOO; ++i) {
        delete goo[i];
    }
    
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
| Static allocator | 将内存分配器封装为独立类，应用类通过 static 成员使用 |
| 嵌入式指针 | 在空闲块内部存储链表指针，节省内存 |
| Free list | 空闲块的链表管理，分配 O(1)，释放 O(1) |
| 宏定义 | DECLARE_POOL_ALLOC 和 IMPLEMENT_POOL_ALLOC |
| SGI std::alloc | 16 个 free lists，管理 8~128 字节对象 |
| 内存紧凑 | 无 cookie，地址间隔等于对象大小 |

### 11.2 学习收获

1. **理解了职责分离的重要性**
   - 将内存管理逻辑从应用类中分离出来
   - 提高代码的可维护性和复用性

2. **掌握了嵌入式指针的技巧**
   - 在空闲块内部存储链表指针
   - 节省内存，提高利用率

3. **理解了宏定义的使用场景**
   - 减少重复代码
   - 保持代码一致性

4. **了解了 SGI STL 的设计思想**
   - 多 free lists 的设计
   - 内存池策略

### 11.3 后续学习建议

1. **深入学习 SGI STL 源码**
   - 阅读 `stl_alloc.h` 的实现
   - 理解内存池的完整实现

2. **学习现代 C++ 的内存管理**
   - C++11 的 std::allocator
   - C++17 的 pmr
   - C++20 的 allocator_traits

3. **学习内存对齐**
   - 为什么需要对齐
   - 如何正确对齐

4. **学习线程安全的内存分配**
   - 多线程环境下的内存管理
   - 无锁数据结构

---

通过本课程，我们深入理解了 static allocator 的设计思想，掌握了 allocator 类的完整实现，理解了嵌入式指针和 free list 的工作原理，也了解了 SGI STL std::alloc 的设计雏形。这些知识对于理解 C++ 内存管理和写高性能代码非常重要！