# 侯捷C++课程9-11集学习笔记：重载 operator new/operator delete 深度解析

## 一、 课程概述

### 1.1 课程主题

本课程深入讲解了 C++ 中**重载 operator new/operator delete** 的各种形式和机制，包括成员函数版本、全局函数版本、placement new/delete、以及在实际项目中的应用（如 basic_string）。

### 1.2 课程内容结构

| 集数 | 主题 | 核心内容 |
|-----|------|---------|
| 第9集 | 重载 | operator new/operator delete 的基本重载形式 |
| 第10集 | 重载示例（上） | 成员函数和全局函数的重载示例 |
| 第11集 | 重载示例（下） | placement new/delete 的详细示例 |

### 1.3 学习目标

1. 理解 operator new/operator delete 的两种重载方式（成员函数和全局函数）
2. 掌握 placement new 的重载和使用场景
3. 理解 placement delete 的特殊调用时机
4. 了解标准库（如 basic_string）中 placement new 的实际应用
5. 掌握面试中的高频问题和最佳实践

---

## 二、 C++ 应用程序分配内存的途径

### 2.1 内存分配路径图

```
C++ 应用程序
    ↓
expression (new Foo())  ← 不可改变，不可重载
    ↓
    ├─────────────────────────────────┐
    ↓                                 ↓
member functions              global functions
Foo::operator new()        ::operator new()
Foo::operator delete()      ::operator delete()
    ↓ (可重载)                 ↓ (可重载但少见)
    └─────────────────────────────────┘
                  ↓
            CRT functions
            malloc(size_t)
            free(void*)
                  ↓
            操作系统 API
```

### 2.2 两种改变内存分配的机会

我们有**两个机会**可以改变内存分配机制：

1. **成员函数版本**
   - 在类内部重载 `operator new` 和 `operator delete`
   - 只对该类生效
   - 适合自定义内存管理（如内存池）

2. **全局函数版本**
   - 在全局作用域重载 `::operator new` 和 `::operator delete`
   - 对所有类生效
   - 较少使用，影响范围太大

### 2.3 expression 的不可改变性

**重要：** expression（如 `new Foo()`）是**不可改变，不可重载**的。

编译器会将 expression 固定转换为两步：

```cpp
// 我们写的：
Foo* p = new Foo();

// 编译器转换为：
try {
    void* mem = operator new(sizeof(Foo));   // 第1步：分配内存
    p = static_cast<Foo*>(mem);
    p->Foo::Foo();                          // 第2步：调用构造函数
} catch (...) {
    // 处理异常
}
```

### 2.4 自定义实现方式

我们也可以模仿 new expression，手动拆分步骤：

```cpp
// 模仿 new expression 的手动实现
Foo* p = (Foo*)malloc(sizeof(Foo));  // 只分配内存
new (p) Foo(x);                     // 在已分配内存上构造对象
// ... 使用 ...
p->~Foo();                           // 手动调用析构
free(p);                            // 释放内存
```

---

## 三、 重载 operator new/operator delete（成员函数版本）

### 3.1 基本形式

```cpp
class Foo {
public:
    // 重载 operator new
    static void* operator new(size_t size);
    
    // 重载 operator delete
    // 第二个参数 size_t 是可选的
    static void operator delete(void* p, size_t size);
};
```

**关键点：**
- `operator new` 和 `operator delete` **必须是 static 成员函数**
- 因为它们在对象构造前/析构后调用，没有 this 指针

### 3.2 完整示例代码

```cpp
#include <iostream>
#include <cstdlib>
#include <string>

class Foo {
private:
    int _id;
    long _data;
    std::string _str;
public:
    // 默认构造函数
    Foo() : _id(0) {
        std::cout << "default ctor. this=" << this 
                  << " id=" << _id << std::endl;
    }
    
    // 带参数构造函数
    Foo(int i) : _id(i) {
        std::cout << "ctor. this=" << this 
                  << " id=" << _id << std::endl;
    }
    
    // 析构函数
    ~Foo() {
        std::cout << "dtor. this=" << this 
                  << " id=" << _id << std::endl;
    }
    
    // ==========================================
    // 重载 operator new（单个对象版本）
    // ==========================================
    static void* operator new(size_t size) {
        std::cout << "Foo::operator new(), size=" << size << std::endl;
        void* p = malloc(size);  // 调用 malloc
        std::cout << "  return: " << p << std::endl;
        return p;
    }
    
    // ==========================================
    // 重载 operator delete（单个对象版本）
    // ==========================================
    static void operator delete(void* p, size_t size) {
        std::cout << "Foo::operator delete(), p=" << p 
                  << " size=" << size << std::endl;
        free(p);  // 调用 free
    }
    
    // ==========================================
    // 重载 operator new[]（数组版本）
    // ==========================================
    static void* operator new[](size_t size) {
        std::cout << "Foo::operator new[](), size=" << size << std::endl;
        void* p = malloc(size);  // 调用 malloc
        std::cout << "  return: " << p << std::endl;
        return p;
    }
    
    // ==========================================
    // 重载 operator delete[]（数组版本）
    // ==========================================
    static void operator delete[](void* p, size_t size) {
        std::cout << "Foo::operator delete[](), p=" << p 
                  << " size=" << size << std::endl;
        free(p);  // 调用 free
    }
};

// 测试单个对象
void test_single_object() {
    std::cout << "\n=== 测试单个对象 ===" << std::endl;
    
    // 使用成员函数的 operator new
    Foo* pf = new Foo(7);
    
    // 使用成员函数的 operator delete
    delete pf;
}

// 测试数组
void test_array() {
    std::cout << "\n=== 测试数组 ===" << std::endl;
    
    // 使用成员函数的 operator new[]
    Foo* pArray = new Foo[5];
    
    // 使用成员函数的 operator delete[]
    delete[] pArray;
}

int main() {
    // 查看对象大小
    std::cout << "sizeof(Foo) = " << sizeof(Foo) << std::endl;
    
    test_single_object();
    test_array();
    
    return 0;
}
```

### 3.3 代码分析

**实现思路说明：**
- 在类内部重载 `operator new` 和 `operator delete`
- `operator new` 只负责分配内存，不构造对象
- `operator delete` 只负责释放内存，不析构对象
- 数组版本也是类似的逻辑

**核心逻辑：**
- 所有重载函数都是 `static` 的
- 使用 `malloc` 和 `free` 作为底层实现
- 第二个 `size_t` 参数是可选的，提供了额外的调试信息

**为什么采用这种实现方式：**
- 灵活性：可以自定义内存分配策略
- 效率：可以实现内存池等优化
- 调试：可以添加日志和跟踪

**初学者引导：**
- 理解 expression 是固定的，不能改
- 能改变的只有 `operator new` 和 `operator delete` 这两个函数
- 成员函数版本只影响该类

### 3.4 输出示例分析

**输出（Foo 没有虚析构）：**
```
sizeof(Foo) = 12

=== 测试单个对象 ===
Foo::operator new(), size=12
  return: 0x3e3988
ctor. this=0x3e3988 id=7
dtor. this=0x3e3988 id=7
Foo::operator delete(), p=0x3e3988 size=12

=== 测试数组 ===
Foo::operator new[](), size=64
  return: 0x3e5cf8
default ctor. this=0x3e5d08 id=0
default ctor. this=0x3e5d14 id=0
default ctor. this=0x3e5d20 id=0
default ctor. this=0x3e5d2c id=0
default ctor. this=0x3e5d38 id=0
dtor. this=0x3e5d38 id=0
dtor. this=0x3e5d2c id=0
dtor. this=0x3e5d20 id=0
dtor. this=0x3e5d14 id=0
dtor. this=0x3e5d08 id=0
Foo::operator delete[](), p=0x3e5cf8 size=64
```

**关键点：**
- 单个对象：size = 12（三个成员变量）
- 数组：size = 64（5个对象 + cookie）

**输出（Foo 有虚析构）：**
```
sizeof(Foo) = 16  ← 增加了虚函数表指针
```

### 3.5 优化建议

**当前实现的改进空间：**
1. **内存池**：预分配内存，减少 malloc 调用
2. **对齐控制**：确保内存正确对齐
3. **统计信息**：记录分配次数和泄漏检测
4. **线程安全**：加入互斥锁

**内存池版本示例：**
```cpp
class FooWithPool {
public:
    static void* operator new(size_t size) {
        if (free_list == nullptr) {
            // 预分配多个对象
            // ...
        }
        // 从空闲链表取一个
        void* p = free_list;
        free_list = free_list->next;
        return p;
    }
private:
    FooWithPool* next;
    static FooWithPool* free_list;
};
```

---

## 四、 重载 ::operator new/::operator delete（全局函数版本）

### 4.1 基本形式

```cpp
// 不能放在 namespace 内！
inline void* operator new(size_t size) {
    // 自定义实现
}

inline void operator delete(void* ptr) noexcept {
    // 自定义实现
}

inline void* operator new[](size_t size) {
    // 自定义实现
}

inline void operator delete[](void* ptr) noexcept {
    // 自定义实现
}
```

### 4.2 完整示例代码

```cpp
#include <iostream>
#include <cstdlib>

// ==========================================
// 自定义的内存分配和释放函数
// ==========================================
void* myAlloc(size_t size) {
    std::cout << "myAlloc(size=" << size << ")" << std::endl;
    return malloc(size);
}

void myFree(void* ptr) {
    std::cout << "myFree(ptr=" << ptr << ")" << std::endl;
    free(ptr);
}

// ==========================================
// 重载全局 ::operator new
// ==========================================
// 注意：不能被声明于一个 namespace 内
inline void* operator new(size_t size) {
    std::cout << "::operator new(size=" << size << ")" << std::endl;
    return myAlloc(size);
}

// ==========================================
// 重载全局 ::operator new[]
// ==========================================
inline void* operator new[](size_t size) {
    std::cout << "::operator new[](size=" << size << ")" << std::endl;
    return myAlloc(size);
}

// ==========================================
// 重载全局 ::operator delete
// ==========================================
inline void operator delete(void* ptr) noexcept {
    std::cout << "::operator delete(ptr=" << ptr << ")" << std::endl;
    myFree(ptr);
}

// ==========================================
// 重载全局 ::operator delete[]
// ==========================================
inline void operator delete[](void* ptr) noexcept {
    std::cout << "::operator delete[](ptr=" << ptr << ")" << std::endl;
    myFree(ptr);
}

// 测试类（没有重载成员函数版本）
class Bar {
private:
    int _id;
public:
    Bar() : _id(0) {
        std::cout << "Bar::Bar()" << std::endl;
    }
    ~Bar() {
        std::cout << "Bar::~Bar()" << std::endl;
    }
};

// 测试全局重载
void test_global_overload() {
    std::cout << "\n=== 测试全局 operator new/delete ===" << std::endl;
    
    // 使用全局版本
    Bar* p = new Bar();
    delete p;
    
    Bar* pArray = new Bar[3];
    delete[] pArray;
}

int main() {
    test_global_overload();
    return 0;
}
```

### 4.3 代码分析

**关键点：**
- **不能放在 namespace 内**
- `inline` 关键字可以避免多重定义问题
- 影响范围很大，对所有类都生效

**为什么这种方式少见：**
- 全局影响太大，可能破坏其他库
- 调试困难
- 推荐使用成员函数版本或自定义 allocator

---

## 五、 使用 :: 强制调用全局版本

### 5.1 基本语法

```cpp
Foo* p = ::new Foo(7);  // 强制使用全局 operator new
::delete p;             // 强制使用全局 operator delete

Foo* pArray = ::new Foo[5];  // 数组版本
::delete[] pArray;
```

### 5.2 完整示例代码

```cpp
#include <iostream>
#include <cstdlib>
#include <string>

class Foo {
private:
    int _id;
public:
    Foo() : _id(0) {
        std::cout << "default ctor. this=" << this 
                  << " id=" << _id << std::endl;
    }
    
    Foo(int i) : _id(i) {
        std::cout << "ctor. this=" << this 
                  << " id=" << _id << std::endl;
    }
    
    ~Foo() {
        std::cout << "dtor. this=" << this 
                  << " id=" << _id << std::endl;
    }
    
    // 成员函数版本
    static void* operator new(size_t size) {
        std::cout << "Foo::operator new(), size=" << size << std::endl;
        return malloc(size);
    }
    
    static void operator delete(void* p, size_t size) {
        std::cout << "Foo::operator delete(), p=" << p << std::endl;
        free(p);
    }
    
    static void* operator new[](size_t size) {
        std::cout << "Foo::operator new[](), size=" << size << std::endl;
        return malloc(size);
    }
    
    static void operator delete[](void* p, size_t size) {
        std::cout << "Foo::operator delete[](), p=" << p << std::endl;
        free(p);
    }
};

void test_force_global() {
    std::cout << "\n=== 测试 :: 强制使用全局版本 ===" << std::endl;
    
    // 使用成员函数版本
    std::cout << "\n--- 使用成员函数版本 ---" << std::endl;
    Foo* pf1 = new Foo(7);
    delete pf1;
    
    Foo* pArray1 = new Foo[5];
    delete[] pArray1;
    
    // 强制使用全局版本
    std::cout << "\n--- 强制使用全局版本 ---" << std::endl;
    Foo* pf2 = ::new Foo(7);  // 注意 :: 前缀
    ::delete pf2;
    
    Foo* pArray2 = ::new Foo[5];
    ::delete[] pArray2;
}

int main() {
    test_force_global();
    return 0;
}
```

### 5.3 代码分析

**关键点：**
- `::new` 和 `::delete` 会绕过所有成员函数重载
- 强制使用全局版本（即使是标准库版本）
- 这样调用，不会进入我们重载的成员函数版本

**输出示例：**
```
=== 测试 :: 强制使用全局版本 ===

--- 使用成员函数版本 ---
Foo::operator new(), size=...
ctor. this=...
dtor. this=...
Foo::operator delete(), p=...

--- 强制使用全局版本 ---
(没有 Foo::operator new 的输出，直接调用全局版本)
ctor. this=...
dtor. this=...
(直接调用全局版本的 delete)
```

---

## 六、 重载 placement new/placement delete

### 6.1 placement new 的重载规则

我们可以重载 class member `operator new()`，写出多个版本，前提是：
- **每一个版本的声明都必须有独特的参数列**
- **其中第一个参数必须是 `size_t`**
- **其余参数以 new 所指定的 placement arguments 为初值**
- **出现在 `new (...)` 小括号内的便是所谓 placement arguments**

### 6.2 基本示例代码

```cpp
#include <iostream>
#include <cstdlib>
#include <stdexcept>

class Bad {};  // 用于测试异常

class Foo {
public:
    Foo() {
        std::cout << "Foo::Foo()" << std::endl;
    }
    
    Foo(int) {
        std::cout << "Foo::Foo(int)" << std::endl;
        throw Bad();  // 故意抛出异常，测试 placement delete
    }
    
    // ==========================================
    // (1) 一般的 operator new 的重载
    // ==========================================
    static void* operator new(size_t size) {
        std::cout << "operator new(size_t), size=" << size << std::endl;
        return malloc(size);
    }
    
    // ==========================================
    // (2) 标准库已提供的 placement new 的重载形式
    // ==========================================
    static void* operator new(size_t size, void* start) {
        std::cout << "operator new(size_t, void*), size=" << size 
                  << " start=" << start << std::endl;
        return start;  // 只是返回传入的指针，不分配新内存
    }
    
    // ==========================================
    // (3) 新的 placement new（接受一个 long）
    // ==========================================
    static void* operator new(size_t size, long extra) {
        std::cout << "operator new(size_t, long), size=" << size 
                  << " extra=" << extra << std::endl;
        return malloc(size + extra);  // 多分配一些
    }
    
    // ==========================================
    // (4) 新的 placement new（接受 long 和 char）
    // ==========================================
    static void* operator new(size_t size, long extra, char init) {
        std::cout << "operator new(size_t, long, char), size=" << size 
                  << " extra=" << extra << " init=" << init << std::endl;
        return malloc(size + extra);
    }
    
    // ==========================================
    // (5) 错误示范：第一个参数不是 size_t
    // ==========================================
    // 错误！
    // static void* operator new(long extra, char init) {
    //     return malloc(extra);
    // }
    // 编译器报错：
    // [Error] 'operator new' takes type 'size_t' 
    // ('unsigned int') as first parameter
    
    
    // ==========================================
    // placement delete：只在 ctor 抛出异常时调用
    // ==========================================
    
    // (1) 对应一般的 operator new
    static void operator delete(void* p, size_t size) {
        std::cout << "operator delete(void*, size_t)" << std::endl;
        free(p);
    }
    
    // (2) 对应 (2) 的 placement new
    static void operator delete(void* p, void*) {
        std::cout << "operator delete(void*, void*)" << std::endl;
        free(p);
    }
    
    // (3) 对应 (3) 的 placement new
    static void operator delete(void* p, long) {
        std::cout << "operator delete(void*, long)" << std::endl;
        free(p);
    }
    
    // (4) 对应 (4) 的 placement new
    static void operator delete(void* p, long, char) {
        std::cout << "operator delete(void*, long, char)" << std::endl;
        free(p);
    }
    
private:
    int m_i;
};

void test_placement_new() {
    std::cout << "\n=== 测试 placement new 各种形式 ===" << std::endl;
    
    Foo start;  // 用于 placement new
    
    // (1) 一般形式
    std::cout << "\n--- (1) 一般形式 ---" << std::endl;
    Foo* p1 = new Foo;
    
    // (2) placement new（接受 void*）
    std::cout << "\n--- (2) placement new (void*) ---" << std::endl;
    Foo* p2 = new (&start) Foo;
    
    // (3) placement new（接受 long）
    std::cout << "\n--- (3) placement new (long) ---" << std::endl;
    Foo* p3 = new (100L) Foo;
    
    // (4) placement new（接受 long 和 char）
    std::cout << "\n--- (4) placement new (long, char) ---" << std::endl;
    Foo* p4 = new (100L, 'a') Foo;
    
    // (5) 测试异常（会调用 placement delete）
    std::cout << "\n--- (5) 测试异常 ---" << std::endl;
    try {
        Foo* p5 = new (100L, 'a') Foo(1);  // Foo(int) 会抛异常
    } catch (Bad&) {
        std::cout << "caught Bad exception" << std::endl;
    }
}

int main() {
    test_placement_new();
    return 0;
}
```

### 6.3 代码分析

**关键点：**
- **第一个参数必须是 size_t**（不能改变）
- **多个版本通过后面的参数区分**
- **placement delete 只有在构造函数抛出异常时调用**

**placement delete 的调用时机：**
```
new (args) Foo() 调用流程：
    ↓
调用 operator new(size, args)  ← 分配内存
    ↓
调用 Foo::Foo()              ← 构造对象
    ↓
    ├─ 成功 → 继续执行，不调用 placement delete
    └─ 失败（抛出异常） → 调用 operator delete(p, args)
```

**重要：**
- **正常 delete 不会调用 placement delete**
- **只有在 new 成功分配内存但构造失败时才调用**
- **目的：释放已经分配但未完全构造的对象的内存**

### 6.4 输出示例分析

```
=== 测试 placement new 各种形式 ===

Foo::Foo()

--- (1) 一般形式 ---
operator new(size_t), size=4
Foo::Foo()

--- (2) placement new (void*) ---
operator new(size_t, void*), size=4 start=0x...
Foo::Foo()

--- (3) placement new (long) ---
operator new(size_t, long), size=4 extra=100
Foo::Foo()

--- (4) placement new (long, char) ---
operator new(size_t, long, char), size=4 extra=100 init=a
Foo::Foo()

--- (5) 测试异常 ---
operator new(size_t, long, char), size=4 extra=100 init=a
Foo::Foo(int)
operator delete(void*, long, char)  ← 因为构造失败，调用 placement delete
caught Bad exception
```

### 6.5 VC6 警告说明

如果提供了 placement new 但没有对应的 placement delete，VC6 会警告：

```
warning C4291: 'void *__cdecl Foo::operator new(~)' :
no matching operator delete found; memory will not be freed
if initialization throws an exception
```

**含义：** 如果构造函数抛出异常，内存将不会被释放！

**即使 operator delete(...) 未能一一对应于 operator new(...)，也不会出现任何报错。** 意思是：你放弃处理构造函数抛出的异常。

---

## 七、 C++ 容器分配内存的途径

### 7.1 容器内存分配路径图

```
容器 Container<T>
    ↓
T* p = allocate();
construct();
...
destroy();
deallocate p;
    ↓
    ├────────────────────────────────────────┐
    ↓                                         ↓
std::allocator<T>                    其他分配器
    ↓ (G4.9)                             ↓
new_allocator<T>                   _pool_alloc<T>
    ↓ (G2.9 叫 alloc)                  ↓
    └────────────────────────────────────────┘
                  ↓
        global functions (可重载但少见)
        ::operator new(size_t)
        ::operator delete(void*)
                  ↓
            CRT functions
            malloc(size_t)
            free(void*)
```

### 7.2 标准库分配器简介

| 分配器 | 说明 |
|--------|------|
| std::allocator | 标准默认分配器 |
| new_allocator (G4.9) | 简单包装 ::operator new |
| _pool_alloc (G4.9) | 内存池分配器（G2.9 叫 alloc） |

### 7.3 自定义分配器示例

```cpp
#include <iostream>
#include <memory>
#include <vector>

// 自定义分配器
template <typename T>
class MyAllocator {
public:
    typedef T value_type;
    typedef T* pointer;
    typedef const T* const_pointer;
    typedef T& reference;
    typedef const T& const_reference;
    typedef std::size_t size_type;
    typedef std::ptrdiff_t difference_type;
    
    template <typename U>
    struct rebind {
        typedef MyAllocator<U> other;
    };
    
    MyAllocator() noexcept {}
    
    template <typename U>
    MyAllocator(const MyAllocator<U>&) noexcept {}
    
    pointer allocate(size_type n) {
        std::cout << "MyAllocator::allocate, n=" << n << std::endl;
        return static_cast<pointer>(std::malloc(n * sizeof(T)));
    }
    
    void deallocate(pointer p, size_type n) noexcept {
        std::cout << "MyAllocator::deallocate, p=" << p << " n=" << n << std::endl;
        std::free(p);
    }
    
    template <typename U, typename... Args>
    void construct(U* p, Args&&... args) {
        new (p) U(std::forward<Args>(args)...);
    }
    
    template <typename U>
    void destroy(U* p) {
        p->~U();
    }
};

// 测试自定义分配器
void test_custom_allocator() {
    std::cout << "\n=== 测试自定义分配器 ===" << std::endl;
    
    std::vector<int, MyAllocator<int>> vec;
    vec.push_back(10);
    vec.push_back(20);
    vec.push_back(30);
    
    for (int x : vec) {
        std::cout << x << " ";
    }
    std::cout << std::endl;
}

int main() {
    test_custom_allocator();
    return 0;
}
```

---

## 八、 basic_string 使用 new(extra) 扩充申请量

### 8.1 实际应用案例

basic_string 使用 placement new 来一次性分配对象和字符串数据的内存：

```cpp
template <class charT, class traits, class Allocator>
class basic_string {
private:
    struct Rep {
        // ... 引用计数、长度等 ...
        
        void release() {
            if (--ref == 0) {
                delete this;  // 删除自己
            }
        }
        
        inline static void* operator new(size_t s, size_t extra) {
            // extra 用于字符串数据
            return Allocator::allocate(s + extra * sizeof(charT));
        }
        
        inline static void operator delete(void* ptr) {
            Allocator::deallocate(ptr, 
                sizeof(Rep) + reinterpret_cast<Rep*>(ptr)->res * sizeof(charT));
        }
        
        inline static Rep* create(size_t extra) {
            extra = frob_size(extra + 1);  // 调整大小
            Rep* p = new (extra) Rep;      // 使用 placement new
            return p;
        }
    };
};
```

### 8.2 设计思路

**为什么这样设计：**
1. **一次分配**：对象和数据一起分配，减少分配次数
2. **内存连续**：更好的缓存局部性
3. **灵活性**：extra 可以根据需要调整

**内存布局：**
```
┌─────────────────┐
│   Rep 对象     │  ← Rep* p
├─────────────────┤
│  字符串数据区  │  ← extra 部分
└─────────────────┘
```

---

## 九、 C++ 新标准的更新和改进

### 9.1 C++11 noexcept

```cpp
// C++11 之前
void operator delete(void* ptr) throw();

// C++11 之后
void operator delete(void* ptr) noexcept;
```

### 9.2 C++11 alignas 和 align_val_t

```cpp
// C++11：对齐说明符
struct alignas(64) MyClass {
    // ...
};

// C++17：operator new 支持对齐参数
void* operator new(size_t size, std::align_val_t align);
void* operator new[](size_t size, std::align_val_t align);
```

### 9.3 C++17 新的 operator new 重载

```cpp
// C++17 新增：带对齐参数的版本
void* operator new(size_t size, std::align_val_t align);
void* operator new(size_t size, std::align_val_t align,
                   const std::nothrow_t&) noexcept;
void* operator new[](size_t size, std::align_val_t align);
void* operator new[](size_t size, std::align_val_t align,
                     const std::nothrow_t&) noexcept;
```

---

## 十、 面试高频问题与解答

### 10.1 基础问题

#### Q1: operator new 和 new 有什么区别？

**答案：**

| 特性 | new | operator new |
|------|-----|--------------|
| 是什么 | expression（表达式） | 函数（function） |
| 功能 | 分配内存 + 调用构造函数 | 只分配内存，不构造对象 |
| 是否可以重载 | 不可以（expression 是固定的） | 可以（成员函数或全局函数） |
| 调用关系 | new 内部调用 operator new | operator new 被 new 调用 |

**详细说明：**
```cpp
// 我们写的：
Foo* p = new Foo();

// 编译器转换为：
void* mem = operator new(sizeof(Foo));  // 调用 operator new
Foo* p = static_cast<Foo*>(mem);
p->Foo::Foo();                          // 调用构造函数
```

#### Q2: operator new/operator delete 可以重载吗？有几种方式？

**答案：**

**可以重载！有两种方式：**

1. **成员函数版本（per-class allocator）**
   - 在类内部重载
   - 只对该类生效
   - 适合自定义内存管理

```cpp
class Foo {
public:
    static void* operator new(size_t size);
    static void operator delete(void* p);
};
```

2. **全局函数版本**
   - 在全局作用域重载 `::operator new`
   - 对所有类生效
   - 较少使用

```cpp
void* operator new(size_t size) { /* ... */ }
void operator delete(void* p) noexcept { /* ... */ }
```

#### Q3: operator new/operator delete 必须是 static 吗？

**答案：**

**是的，必须是 static！**

**原因：**
- operator new 在构造对象之前调用，此时对象还不存在，没有 this 指针
- operator delete 在析构对象之后调用，此时对象已经被销毁，没有 this 指针
- 因此它们必须是静态成员函数

### 10.2 placement new/delete 相关问题

#### Q4: placement new 是什么？可以重载吗？

**答案：**

**placement new 是 operator new 的一种重载形式，接受额外的参数。**

**重载规则：**
- 可以有多个版本
- 第一个参数必须是 `size_t`
- 后面的参数是 placement arguments

```cpp
class Foo {
public:
    // 普通版本
    static void* operator new(size_t size);
    
    // placement new 1：接受 void*
    static void* operator new(size_t size, void* start);
    
    // placement new 2：接受 long
    static void* operator new(size_t size, long extra);
    
    // placement new 3：接受 long 和 char
    static void* operator new(size_t size, long extra, char init);
};

// 使用
Foo* p1 = new Foo;                          // 普通版本
Foo* p2 = new (&start) Foo;                // placement new 1
Foo* p3 = new (100L) Foo;                  // placement new 2
Foo* p4 = new (100L, 'a') Foo;             // placement new 3
```

#### Q5: placement delete 什么时候调用？

**答案：**

**placement delete 只有在一种情况下调用：**

```
当使用 placement new 分配内存成功，但构造函数抛出异常时。
```

**详细流程：**
```
new (args) Foo()
    ↓
operator new(size, args)  ← 分配内存成功
    ↓
Foo::Foo()              ← 构造函数抛出异常
    ↓
operator delete(p, args) ← 调用 placement delete 清理内存
```

**重要：**
- **正常 delete 不会调用 placement delete**
- **只有在构造失败时才调用**
- **目的：防止内存泄漏**

#### Q6: 如果提供了 placement new 但没有对应 placement delete 会怎样？

**答案：**

**不会报错，但可能有内存泄漏！**

**VC6 警告：**
```
warning C4291: no matching operator delete found;
memory will not be freed if initialization throws an exception
```

**含义：** 如果构造函数抛出异常，已经分配的内存将不会被释放！

**最佳实践：**
- 每一个 placement new 都应该对应一个 placement delete
- 参数列表要匹配（除了第一个参数）

### 10.3 进阶问题

#### Q7: 如何强制使用全局的 operator new/delete？

**答案：**

**使用 :: 前缀！**

```cpp
Foo* p = ::new Foo();    // 强制使用全局 operator new
::delete p;               // 强制使用全局 operator delete

Foo* pArray = ::new Foo[5];  // 数组版本
::delete[] pArray;
```

**作用：**
- 绕过成员函数重载
- 强制使用全局版本（即使是标准库版本）

#### Q8: 重载 operator new/delete 有什么实际应用？

**答案：**

**常见应用场景：**

1. **内存池**：减少 malloc/free 调用
2. **对象池**：频繁创建销毁的对象
3. **调试**：跟踪内存分配，检测泄漏
4. **对齐控制**：特殊对齐要求
5. **统计**：收集内存使用信息
6. **安全检查**：边界检查、使用后检测

#### Q9: operator delete 的第二个 size_t 参数有什么用？

**答案：**

**可选，提供调试信息。**

```cpp
class Foo {
public:
    static void operator delete(void* p, size_t size) {
        // size 告诉我们要释放多少字节
        // 可以用于调试、统计、验证等
        std::cout << "delete " << size << " bytes at " << p << std::endl;
        free(p);
    }
};
```

**这个参数是可选的，如果提供，编译器会传正确的大小。**

### 10.4 最佳实践问题

#### Q10: 重载 operator new/delete 的最佳实践有哪些？

**答案：**

**1. 优先使用成员函数版本，避免全局版本**

**2. 配对使用：**
- 如果重载了 `operator new`，也要重载 `operator delete`
- 如果重载了 `operator new[]`，也要重载 `operator delete[]`

**3. placement new 和 placement delete 配对：**
- 每一个 placement new 都应该对应一个 placement delete

**4. 继承问题：**
- 派生类可能继承基类的 operator new/delete
- 注意处理不同大小的对象

**5. 异常安全：**
- operator new 应该支持 new_handler
- 分配失败应该抛出 bad_alloc（或使用 nothrow）

---

## 十一、 完整代码示例与分析

### 11.1 完整的重载示例

```cpp
#include <iostream>
#include <cstdlib>
#include <new>
#include <stdexcept>
#include <string>

// 异常类
class Bad {};

class Foo {
private:
    int _id;
    long _data;
    std::string _str;
    
public:
    Foo() : _id(0) {
        std::cout << "Foo::Foo() this=" << this << " id=" << _id << std::endl;
    }
    
    Foo(int i) : _id(i) {
        std::cout << "Foo::Foo(int) this=" << this << " id=" << _id << std::endl;
        if (i == 999) {
            throw Bad();  // 用于测试 placement delete
        }
    }
    
    ~Foo() {
        std::cout << "Foo::~Foo() this=" << this << " id=" << _id << std::endl;
    }
    
    // ==========================================
    // 单个对象的 operator new/delete
    // ==========================================
    static void* operator new(size_t size) {
        std::cout << "[1] Foo::operator new(size=" << size << ")" << std::endl;
        void* p = malloc(size);
        std::cout << "    return " << p << std::endl;
        return p;
    }
    
    static void operator delete(void* p, size_t size) noexcept {
        std::cout << "[1] Foo::operator delete(p=" << p << ", size=" << size << ")" << std::endl;
        free(p);
    }
    
    // ==========================================
    // 数组的 operator new[]/delete[]
    // ==========================================
    static void* operator new[](size_t size) {
        std::cout << "[2] Foo::operator new[](size=" << size << ")" << std::endl;
        void* p = malloc(size);
        std::cout << "    return " << p << std::endl;
        return p;
    }
    
    static void operator delete[](void* p, size_t size) noexcept {
        std::cout << "[2] Foo::operator delete[](p=" << p << ", size=" << size << ")" << std::endl;
        free(p);
    }
    
    // ==========================================
    // placement new 1: void*
    // ==========================================
    static void* operator new(size_t size, void* start) {
        std::cout << "[3] Foo::operator new(size=" << size << ", void*=" << start << ")" << std::endl;
        return start;
    }
    
    static void operator delete(void* p, void* start) {
        std::cout << "[3] Foo::operator delete(p=" << p << ", void*=" << start << ")" << std::endl;
        free(p);
    }
    
    // ==========================================
    // placement new 2: long
    // ==========================================
    static void* operator new(size_t size, long extra) {
        std::cout << "[4] Foo::operator new(size=" << size << ", long=" << extra << ")" << std::endl;
        void* p = malloc(size + extra);
        std::cout << "    return " << p << std::endl;
        return p;
    }
    
    static void operator delete(void* p, long extra) {
        std::cout << "[4] Foo::operator delete(p=" << p << ", long=" << extra << ")" << std::endl;
        free(p);
    }
    
    // ==========================================
    // placement new 3: long + char
    // ==========================================
    static void* operator new(size_t size, long extra, char init) {
        std::cout << "[5] Foo::operator new(size=" << size << ", long=" << extra << ", char=" << init << ")" << std::endl;
        void* p = malloc(size + extra);
        std::cout << "    return " << p << std::endl;
        return p;
    }
    
    static void operator delete(void* p, long extra, char init) {
        std::cout << "[5] Foo::operator delete(p=" << p << ", long=" << extra << ", char=" << init << ")" << std::endl;
        free(p);
    }
};

void test_all_cases() {
    std::cout << "========== 测试开始 ==========" << std::endl;
    std::cout << "sizeof(Foo) = " << sizeof(Foo) << std::endl;
    
    // 测试1：单个对象
    std::cout << "\n---------- 测试1：单个对象 ----------" << std::endl;
    Foo* p1 = new Foo(7);
    delete p1;
    
    // 测试2：数组
    std::cout << "\n---------- 测试2：数组 ----------" << std::endl;
    Foo* p2 = new Foo[3];
    delete[] p2;
    
    // 测试3：强制全局版本
    std::cout << "\n---------- 测试3：强制全局 ----------" << std::endl;
    Foo* p3 = ::new Foo(8);
    ::delete p3;
    
    // 测试4：placement new (void*)
    std::cout << "\n---------- 测试4：placement new (void*) ----------" << std::endl;
    Foo local;
    Foo* p4 = new (&local) Foo(10);
    p4->~Foo();  // 手动析构
    
    // 测试5：placement new (long)
    std::cout << "\n---------- 测试5：placement new (long) ----------" << std::endl;
    Foo* p5 = new (100L) Foo(20);
    delete p5;
    
    // 测试6：placement new (long, char)
    std::cout << "\n---------- 测试6：placement new (long, char) ----------" << std::endl;
    Foo* p6 = new (100L, 'a') Foo(30);
    delete p6;
    
    // 测试7：placement delete（构造函数抛异常）
    std::cout << "\n---------- 测试7：placement delete ----------" << std::endl;
    try {
        Foo* p7 = new (100L, 'a') Foo(999);  // i=999 会抛异常
        delete p7;  // 如果成功才会到这里
    } catch (Bad&) {
        std::cout << "caught Bad exception!" << std::endl;
    }
    
    std::cout << "\n========== 测试结束 ==========" << std::endl;
}

int main() {
    try {
        test_all_cases();
    } catch (std::exception& e) {
        std::cerr << "Exception: " << e.what() << std::endl;
    }
    return 0;
}
```

### 11.2 代码分析

**实现思路说明：**
- 展示了所有可能的重载形式
- 添加了详细的日志输出，方便观察调用流程
- 包含异常测试，展示 placement delete 的调用时机

**核心逻辑：**
- 每个重载都有对应的 delete 版本
- 构造函数可以抛出异常测试清理逻辑
- :: 前缀测试强制全局版本

**初学者引导：**
- 观察输出，理解调用顺序
- 注意 placement delete 只在构造失败时调用
- 理解 :: 的作用

---

## 十二、 总结与反思

### 12.1 核心知识点总结

| 知识点 | 要点 |
|--------|------|
| 两个重载机会 | 成员函数版本（per-class）和全局函数版本 |
| expression 不可变 | new Foo() 固定转换为 operator new + ctor |
| operator new/delete | 必须是 static，只分配/释放内存 |
| 第二个 size_t 参数 | 可选，用于调试信息 |
| placement new | 可以多个版本，第一个参数必须是 size_t |
| placement delete | 只在 ctor 抛出异常时调用，用于清理 |
| :: 前缀 | 强制使用全局版本，绕过成员函数重载 |
| 容器分配器 | 通过 allocator 自定义内存管理 |
| basic_string 应用 | 使用 new(extra) 一次分配对象和数据 |

### 12.2 学习收获

1. **理解了内存分配的层次结构**
   - expression → operator new → malloc → OS API
   - 哪些可以改变，哪些不能

2. **掌握了重载的各种形式**
   - 成员函数 vs 全局函数
   - placement new/delete 的各种形式

3. **了解了 placement delete 的特殊调用时机**
   - 只在构造失败时调用
   - 防止内存泄漏

4. **看到了实际应用案例**
   - basic_string 的巧妙设计
   - 可以借鉴到自己的项目中

### 12.3 后续学习建议

1. **深入学习 STL allocator**
   - 理解 std::allocator 的实现
   - 学习自定义分配器的编写

2. **研究内存池实现**
   - 高效的内存管理
   - 减少碎片和分配开销

3. **学习调试工具**
   - AddressSanitizer
   - Valgrind
   - 内存调试器

4. **阅读优秀源码**
   - 看 GCC、Clang 的实现
   - 看实际项目中如何应用

### 12.4 最佳实践回顾

1. **优先使用成员函数版本，避免全局重载**
2. **placement new 和 placement delete 要配对**
3. **注意继承带来的问题**
4. **考虑使用标准库的分配器而非自己重载**
5. **充分测试，特别是异常情况**

---

通过本课程的学习，我们全面深入地理解了 C++ 中重载 `operator new/operator delete` 的各种机制，包括成员函数版本、全局函数版本、placement new/delete 的使用时机和注意事项，以及在实际项目（如 basic_string）中的应用。这些知识对于编写高效、健壮的 C++ 代码非常重要，也是面试中的高频考点！
