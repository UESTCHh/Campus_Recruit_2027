# 侯捷C++课程7-8集学习笔记：Array new 与 Placement new 深度解析

## 一、 课程概述

### 1.1 课程主题

本课程深入讲解了C++中的**Array new/Array delete**和**Placement new**的底层实现机制，重点分析了**Cookie机制**、**内存块布局**以及**重载operator new/delete[]**的注意事项。

### 1.2 课程内容结构

| 集数 | 主题 | 核心内容 |
|-----|------|---------|
| 第7集 | Array new/Array delete | 数组版本的new/delete，Cookie机制，默认构造函数要求 |
| 第8集 | Replacement new（Placement new） | 在已分配内存上构造对象，重载operator new的特殊形式 |

### 1.3 学习目标

1. 理解Array new/Array delete的工作原理和Cookie机制
2. 掌握Placement new的正确用法和注意事项
3. 了解内存块的具体布局和调试信息
4. 理解重载operator new/delete[]时的注意事项
5. 掌握常见错误和避免方法

---

## 二、 Array new 与 Array delete

### 2.1 Array new 的基本用法

#### 2.1.1 代码示例

```cpp
#include <iostream>

class Complex {
private:
    int re, im;
public:
    Complex(int r = 0, int i = 0) : re(r), im(i) {
        std::cout << "ctor: Complex(" << r << "," << i << ")" << std::endl;
    }
    
    ~Complex() {
        std::cout << "dtor: ~Complex(" << re << "," << im << ")" << std::endl;
    }
};

int main() {
    // Array new：创建3个Complex对象
    std::cout << "=== Array new ===" << std::endl;
    Complex* pca = new Complex[3];
    
    // 注意：Array new无法给初值，只能使用默认构造函数
    
    std::cout << "\n=== Array delete ===" << std::endl;
    delete[] pca;
    
    return 0;
}
```

#### 2.1.2 代码分析

**Array new的特点：**

1. **多次调用构造函数**：new Complex[3]会调用3次构造函数
2. **无法给初值**：Array new无法像单个对象那样给初始值，只能调用默认构造函数
3. **Cookie机制**：编译器会在分配的内存块头部记录数组大小

#### 2.1.3 执行流程

```
new Complex[3] 的编译转换：
  ↓
1. operator new[] (sizeof(Complex)*3 + sizeof(cookie))
  ↓
2. 调用3次 Complex::Complex()
  ↓
3. 返回指向第一个对象的指针

delete[] pca 的编译转换：
  ↓
1. 调用3次 Complex::~Complex()（按逆序）
  ↓
2. operator delete[] (pca)
```

### 2.2 Cookie 机制详解

#### 2.2.1 内存块布局

```
┌─────────────────────────────────────┐
│ Cookie（记录数组大小）                 │ ← 编译器使用，不暴露给用户
├─────────────────────────────────────┤
│ Complex 对象1                        │ ← pca 指向这里
├─────────────────────────────────────┤
│ Complex 对象2                        │
├─────────────────────────────────────┤
│ Complex 对象3                        │
└─────────────────────────────────────┘
```

#### 2.2.2 Cookie 的作用

**Cookie 的主要功能：**

1. **记录数组大小**：告诉delete[]需要调用多少次析构函数
2. **内存块信息**：可能包含其他内存管理信息（如调试信息）

**不同情况下的 Cookie：**

| 情况 | 是否记录 Cookie | 原因 |
|-----|---------------|------|
| 有non-trivial析构函数的类 | 是 | 需要知道调用多少次析构函数 |
| 没有析构函数的简单类型 | 否 | 不需要调用析构函数 |

#### 2.2.3 代码示例：观察 Cookie 行为

```cpp
#include <iostream>

// 简单类型：没有析构函数（或trivial析构）
class SimpleType {
private:
    int data;
public:
    SimpleType() : data(0) {
        std::cout << "SimpleType ctor" << std::endl;
    }
    // 注意：没有显式析构函数，或析构函数是trivial的
};

// 有复杂析构函数的类
class ComplexType {
private:
    int* data;
public:
    ComplexType() {
        data = new int[10];
        std::cout << "ComplexType ctor" << std::endl;
    }
    
    ~ComplexType() {
        delete[] data;
        std::cout << "ComplexType dtor" << std::endl;
    }
};

int main() {
    // 简单类型数组：可能不记录 cookie
    std::cout << "=== SimpleType array ===" << std::endl;
    SimpleType* pa = new SimpleType[3];
    delete[] pa;
    
    // 复杂类型数组：必须记录 cookie
    std::cout << "\n=== ComplexType array ===" << std::endl;
    ComplexType* pb = new ComplexType[3];
    delete[] pb;
    
    return 0;
}
```

#### 2.2.4 初学者引导

**为什么需要 Cookie？**

因为：
1. delete[] 需要知道调用多少次析构函数
2. 从指针本身无法看出数组大小
3. 编译器需要用 Cookie 来记录这个信息

**什么时候需要 Cookie？**

只有当类有**non-trivial析构函数**时才需要，因为：
- trivial析构函数：什么都不做，不需要调用
- non-trivial析构函数：有实际的清理工作，必须调用

### 2.3 不匹配 delete 的危害

#### 2.3.1 错误示例

```cpp
#include <iostream>

class String {
private:
    char* data;
public:
    String() {
        data = new char[10];
        std::cout << "String ctor" << std::endl;
    }
    
    ~String() {
        delete[] data;
        std::cout << "String dtor" << std::endl;
    }
};

int main() {
    // 错误示例：Array new 之后用了普通 delete
    std::cout << "=== 错误示例 ===" << std::endl;
    String* psa = new String[3];
    
    // delete psa;  // 错误！只会调用1次析构函数，内存泄漏
    
    // 正确做法
    std::cout << "\n=== 正确做法 ===" << std::endl;
    delete[] psa;
    
    return 0;
}
```

#### 2.3.2 不匹配的后果分析

**情况1：类没有指针成员**

```cpp
// 这种情况下，误用 delete 可能不会立即崩溃
class NoPtr {
    int id;
public:
    NoPtr() : id(0) {}
    ~NoPtr() {}  // trivial析构
};

NoPtr* pa = new NoPtr[3];
delete pa;  // 可能不会崩溃，但仍然是错误的！
```

**情况2：类有指针成员**

```cpp
// 这种情况下，误用 delete 会造成内存泄漏
class HasPtr {
    int* data;
public:
    HasPtr() : data(new int[10]) {}
    ~HasPtr() { delete[] data; }
};

HasPtr* pa = new HasPtr[3];
delete pa;  // 只有第一个对象的析构函数被调用，其他两个的内存泄漏！
```

#### 2.3.3 正确做法总结

| 分配方式 | 释放方式 |
|---------|---------|
| new Type | delete |
| new Type[] | delete[] |
| new (ptr) Type | ptr->~Type() |

**重要原则：**

- new 对应 delete
- new[] 对应 delete[]
- placement new 对应手动调用析构函数

### 2.4 Array new 对默认构造函数的要求

#### 2.4.1 代码示例

```cpp
#include <iostream>

class A {
private:
    int id;
public:
    // 默认构造函数
    A() : id(0) {
        std::cout << "default ctor: this=" << this << " id=" << id << std::endl;
    }
    
    // 带参数的构造函数
    A(int i) : id(i) {
        std::cout << "ctor: this=" << this << " id=" << id << std::endl;
    }
    
    ~A() {
        std::cout << "dtor: this=" << this << " id=" << id << std::endl;
    }
};

int main() {
    // Array new 要求类必须有默认构造函数
    std::cout << "=== Array new ===" << std::endl;
    A* buf = new A[3];
    
    // 如果没有默认构造函数，会编译错误：
    // Error: no matching function for call to 'A::A()'
    
    // 演示：如何在已有内存上构造对象（稍后讲解）
    
    delete[] buf;
    
    return 0;
}
```

#### 2.4.2 为什么需要默认构造函数？

因为：
1. Array new 会创建多个对象
2. 需要逐个调用构造函数
3. 无法给每个对象单独传参数
4. 只能调用默认构造函数

**如果没有默认构造函数怎么办？**

方法1：使用 placement new 手动构造（见第3章）

```cpp
// 方法2：C++11 之前没有好办法，C++11 之后：
#include <vector>

// 使用 vector 可以给初值
std::vector<A> vec;
vec.reserve(3);
vec.emplace_back(1);  // 调用 A(1)
vec.emplace_back(2);  // 调用 A(2)
vec.emplace_back(3);  // 调用 A(3)
```

#### 2.4.3 分步骤构造对象

```cpp
#include <iostream>

class A {
private:
    int id;
public:
    A() : id(0) {
        std::cout << "default ctor: this=" << this << " id=" << id << std::endl;
    }
    
    A(int i) : id(i) {
        std::cout << "ctor: this=" << this << " id=" << id << std::endl;
    }
    
    ~A() {
        std::cout << "dtor: this=" << this << " id=" << id << std::endl;
    }
};

int main() {
    const int size = 3;
    
    // 1. 先分配内存（不调用构造函数）
    char* raw_buf = new char[sizeof(A) * size];
    A* buf = reinterpret_cast<A*>(raw_buf);
    
    // 2. 逐个调用构造函数
    std::cout << "=== 手动构造 ===" << std::endl;
    for (int i = 0; i < size; ++i) {
        new (buf + i) A(i);
    }
    
    // 3. 使用对象
    std::cout << "buf=" << reinterpret_cast<void*>(buf) << std::endl;
    
    // 4. 逐个调用析构函数
    std::cout << "=== 手动析构 ===" << std::endl;
    for (int i = size - 1; i >= 0; --i) {
        (buf + i)->~A();
    }
    
    // 5. 释放内存
    delete[] raw_buf;
    
    return 0;
}
```

**输出说明：**

```
=== 手动构造 ===
ctor: this=0x... id=0
ctor: this=0x... id=1
ctor: this=0x... id=2
buf=0x...
=== 手动析构 ===
dtor: this=0x... id=2  ← 逆序析构
dtor: this=0x... id=1
dtor: this=0x... id=0
```

### 2.5 内存块的详细布局（VC6 示例）

#### 2.5.1 内存布局图示

```
调试模式下的内存块布局：

┌─────────────────────────────────────────┐
│ Cookie (包含大小信息)                      │ 4/8字节
├─────────────────────────────────────────┤
│ Debugger Header (调试信息)                │ 32字节
├─────────────────────────────────────────┤
│ int                                     │
│ int                                     │
│ int                                     │
│ int                                     │
│ int                                     │
│ int                                     │
│ int                                     │
│ int                                     │ ← 用户数据区
│ int                                     │
│ int                                     │
├─────────────────────────────────────────┤
│ No Man's Land (边界标记)                  │ 4字节
├─────────────────────────────────────────┤
│ Padding (对齐填充)                        │ 12字节
├─────────────────────────────────────────┤
│ Cookie                                  │ 4/8字节
└─────────────────────────────────────────┘
```

#### 2.5.2 各部分的作用

| 区域 | 大小 | 作用 |
|-----|------|------|
| Cookie | 4/8字节 | 记录内存块大小和其他信息 |
| Debugger Header | 32字节 | 调试信息（仅在调试模式） |
| 用户数据区 | 按需 | 实际存放对象的地方 |
| No Man's Land | 4字节 | 边界标记，检测越界访问 |
| Padding | 可变 | 对齐填充 |

#### 2.5.3 代码示例：观察内存布局

```cpp
#include <iostream>
#include <new>

class Demo {
private:
    int a, b, c;
public:
    Demo() : a(1), b(2), c(3) {
        std::cout << "Demo ctor" << std::endl;
    }
    ~Demo() {
        std::cout << "Demo dtor" << std::endl;
    }
};

int main() {
    // 分配数组
    Demo* arr = new Demo[3];
    
    std::cout << "arr address: " << static_cast<void*>(arr) << std::endl;
    
    // 释放数组
    delete[] arr;
    
    // 栈上分配
    Demo d[3];
    std::cout << "sizeof(d): " << sizeof(d) << std::endl;
    
    return 0;
}
```

### 2.6 重载 operator new/delete[] 的注意事项

#### 2.6.1 错误示例

```cpp
#include <iostream>
#include <new>

class Demo {
private:
    int data;
public:
    Demo() : data(0) {
        std::cout << "Demo ctor" << std::endl;
    }
    
    ~Demo() {
        std::cout << "Demo dtor" << std::endl;
    }
    
    // 重载 operator new
    void* operator new(size_t size) {
        std::cout << "operator new: " << size << " bytes" << std::endl;
        return malloc(size);
    }
    
    // 重载 operator delete
    void operator delete(void* p) noexcept {
        std::cout << "operator delete" << std::endl;
        free(p);
    }
    
    // 重载 operator new[]
    void* operator new[](size_t size) {
        std::cout << "operator new[]: " << size << " bytes" << std::endl;
        return malloc(size);
    }
    
    // 重载 operator delete[]
    void operator delete[](void* p) noexcept {
        std::cout << "operator delete[]" << std::endl;
        free(p);
    }
};

int main() {
    // 注意：使用重载的 operator new[] 时
    // 传入的指针位置可能包含 cookie，需要特别小心
    Demo* pd = new Demo[3];
    delete[] pd;
    
    return 0;
}
```

#### 2.6.2 为什么会有问题？

**原因：**

1. **指针偏移**：Array new 返回的指针可能不是分配块的开始位置
2. **Cookie位置**：Cookie 可能在指针之前
3. **调试模式**：调试模式下布局更复杂

**VC6 下的调试模式错误：**

如果只写了 `delete pd` 而不是 `delete[] pd`，在调试模式下可能会触发 `Assertion Failed`，错误信息可能是 `_BLOCK_TYPE_IS_VALID(pHead->nBlockUse)`。

#### 2.6.3 安全的重载方法

```cpp
#include <iostream>
#include <new>

class SafeDemo {
private:
    int data;
public:
    SafeDemo() : data(0) {}
    
    ~SafeDemo() {}
    
    // 使用全局 operator new/delete，不要重载数组版本
    // 或者非常小心地处理指针偏移
    
    // 如果确实需要重载，建议使用内存池等更安全的方案
};

// 更好的方案：使用自定义 allocator
template <typename T>
class MyAllocator {
public:
    T* allocate(size_t n) {
        return static_cast<T*>(malloc(sizeof(T) * n));
    }
    
    void deallocate(T* p) {
        free(p);
    }
    
    template <typename... Args>
    void construct(T* p, Args&&... args) {
        new (p) T(std::forward<Args>(args)...);
    }
    
    void destroy(T* p) {
        p->~T();
    }
};
```

---

## 三、 Placement new（Replacement new）

### 3.1 Placement new 的基本概念

#### 3.1.1 什么是 Placement new？

**定义：**

Placement new 是 operator new 的特殊版本，允许我们在**已分配的内存**上构造对象，而不分配新内存。

**特征：**

1. **不分配内存**：只构造对象，不调用 malloc
2. **需要指定位置**：传入一个指针，告诉它在哪里构造
3. **需要手动析构**：没有对应的 placement delete，需要手动调用析构函数

#### 3.1.2 基本用法示例

```cpp
#include <iostream>
#include <new>  // 需要包含这个头文件

class Complex {
private:
    int re, im;
public:
    Complex(int r = 0, int i = 0) : re(r), im(i) {
        std::cout << "ctor: Complex(" << r << "," << i << ")" << std::endl;
    }
    
    ~Complex() {
        std::cout << "dtor: ~Complex(" << re << "," << im << ")" << std::endl;
    }
    
    void print() const {
        std::cout << re << "+" << im << "i" << std::endl;
    }
};

int main() {
    // 1. 预分配内存缓冲区
    char* buf = new char[sizeof(Complex) * 3];
    
    // 2. 在缓冲区上构造对象
    std::cout << "=== Placement new ===" << std::endl;
    Complex* pc = new (buf) Complex(1, 2);
    
    // 3. 使用对象
    pc->print();
    
    // 4. 手动调用析构函数
    std::cout << "\n=== Manual dtor ===" << std::endl;
    pc->~Complex();
    
    // 5. 释放原始内存
    delete[] buf;
    
    return 0;
}
```

**输出结果：**

```
=== Placement new ===
ctor: Complex(1,2)
1+2i
=== Manual dtor ===
dtor: ~Complex(1,2)
```

#### 3.1.3 编译转换过程

```
Complex* pc = new (buf) Complex(1,2);
  ↓
编译器转换为：
  ↓
1. 调用 operator new (sizeof(Complex), buf)
  ↓
2. static_cast<Complex*>(返回值)
  ↓
3. 调用 Complex::Complex(1,2)

operator new 的特殊版本：
void* operator new(size_t size, void* loc) {
    return loc;  // 直接返回传入的指针，不分配新内存
}
```

### 3.2 Placement new 的完整流程

#### 3.2.1 代码示例

```cpp
#include <iostream>
#include <new>

class Complex {
private:
    int re, im;
public:
    Complex(int r = 0, int i = 0) : re(r), im(i) {
        std::cout << "ctor: this=" << this << std::endl;
    }
    
    ~Complex() {
        std::cout << "dtor: this=" << this << std::endl;
    }
};

// 模拟编译器转换
int main() {
    // 预分配内存
    char* buf = new char[sizeof(Complex) * 3];
    
    std::cout << "=== Placement new step by step ===" << std::endl;
    
    // 模拟编译器的转换过程
    Complex* pc;
    try {
        // 步骤1：调用特殊的 operator new
        void* mem = operator new(sizeof(Complex), buf);
        
        // 步骤2：类型转换
        pc = static_cast<Complex*>(mem);
        
        // 步骤3：调用构造函数
        pc->Complex::Complex(1, 2);
    }
    catch (std::bad_alloc&) {
        // allocation失败，不执行构造
        // placement new 通常不会失败，因为不分配内存
    }
    
    std::cout << "buf addr: " << static_cast<void*>(buf) << std::endl;
    std::cout << "pc  addr: " << static_cast<void*>(pc) << std::endl;
    
    // 手动析构
    pc->~Complex();
    
    // 释放原始内存
    delete[] buf;
    
    return 0;
}
```

#### 3.2.2 内存布局图示

```
buf 指向的内存：
  ↓
┌─────────────────────────────────────┐
│ Cookie (optional)                   │
├─────────────────────────────────────┤
│ Complex 对象                       │ ← pc 指向这里
├─────────────────────────────────────┤
│ (空余空间)                           │
├─────────────────────────────────────┤
│ (空余空间)                           │
└─────────────────────────────────────┘
```

### 3.3 Placement new 的不同叫法

#### 3.3.1 多种称呼

**称呼1：Placement new**

"placement" 意思是"放置"，表示在某个位置放置对象。

**称呼2：new(p)**

语法形式的叫法：`new (pointer) Type`。

**称呼3：operator new(size, void*)**

指对应的重载版本的函数签名。

**称呼4：Placement delete（不准确）**

注意：没有真正的 placement delete，这是一种非正式的称呼，指的是手动调用析构函数。

#### 3.3.2 注意事项

**重要：**

- **没有 placement delete**：没有对应的表达式来自动析构和释放
- **必须手动析构**：用完对象后要手动调用析构函数
- **内存分开释放**：原始内存的释放要单独处理

**错误做法：**

```cpp
Complex* pc = new (buf) Complex(1, 2);
delete pc;  // 错误！会释放原始内存，可能导致崩溃
```

**正确做法：**

```cpp
Complex* pc = new (buf) Complex(1, 2);
pc->~Complex();  // 先析构
delete[] buf;   // 再释放原始内存
```

### 3.4 Placement new 的应用场景

#### 3.4.1 场景1：内存池

```cpp
#include <iostream>
#include <vector>

template <typename T>
class MemoryPool {
private:
    std::vector<char*> blocks;
    size_t block_size;
    char* current_ptr;
    size_t current_used;
    
public:
    MemoryPool(size_t block_size = 1024) : block_size(block_size) {
        allocate_new_block();
    }
    
    ~MemoryPool() {
        for (char* block : blocks) {
            delete[] block;
        }
    }
    
    void allocate_new_block() {
        char* new_block = new char[block_size];
        blocks.push_back(new_block);
        current_ptr = new_block;
        current_used = 0;
    }
    
    void* allocate(size_t size) {
        if (current_used + size > block_size) {
            allocate_new_block();
        }
        void* result = current_ptr + current_used;
        current_used += size;
        return result;
    }
    
    // 注意：简单实现，没有回收功能
};

class MyClass {
private:
    int data;
public:
    MyClass(int d) : data(d) {
        std::cout << "MyClass ctor(" << d << ")" << std::endl;
    }
    
    ~MyClass() {
        std::cout << "MyClass dtor" << std::endl;
    }
};

int main() {
    MemoryPool<MyClass> pool;
    
    // 在内存池中构造对象
    std::cout << "=== Construct in pool ===" << std::endl;
    void* p1 = pool.allocate(sizeof(MyClass));
    MyClass* obj1 = new (p1) MyClass(1);
    
    void* p2 = pool.allocate(sizeof(MyClass));
    MyClass* obj2 = new (p2) MyClass(2);
    
    // 手动析构
    std::cout << "\n=== Manual dtor ===" << std::endl;
    obj1->~MyClass();
    obj2->~MyClass();
    
    // 内存池在析构时自动释放
    return 0;
}
```

#### 3.4.2 场景2：栈上构造对象

```cpp
#include <iostream>

class StackObject {
private:
    int data[100];  // 较大的对象
public:
    StackObject(int d) {
        data[0] = d;
        std::cout << "StackObject ctor" << std::endl;
    }
    
    ~StackObject() {
        std::cout << "StackObject dtor" << std::endl;
    }
    
    void set(int d) {
        data[0] = d;
    }
};

int main() {
    // 在栈上预分配缓冲区
    char buffer[sizeof(StackObject)];
    
    // 使用 placement new 构造
    std::cout << "=== Stack placement new ===" << std::endl;
    StackObject* obj = new (buffer) StackObject(42);
    
    // 使用对象
    obj->set(100);
    
    // 手动析构
    obj->~StackObject();
    
    // 栈内存自动释放
    return 0;
}
```

#### 3.4.3 场景3：对象重用

```cpp
#include <iostream>

class Reusable {
private:
    int id;
public:
    Reusable(int i) : id(i) {
        std::cout << "ctor: Reusable(" << i << ")" << std::endl;
    }
    
    ~Reusable() {
        std::cout << "dtor: ~Reusable(" << id << ")" << std::endl;
    }
    
    void reset(int i) {
        id = i;
    }
};

int main() {
    // 预分配内存
    char* buffer = new char[sizeof(Reusable)];
    
    // 第一次使用
    std::cout << "=== First use ===" << std::endl;
    Reusable* obj = new (buffer) Reusable(1);
    obj->~Reusable();
    
    // 重用同一块内存
    std::cout << "\n=== Reuse same memory ===" std::endl;
    obj = new (buffer) Reusable(2);
    obj->~Reusable();
    
    // 再次重用
    std::cout << "\n=== Reuse again ===" std::endl;
    obj = new (buffer) Reusable(3);
    obj->~Reusable();
    
    // 释放内存
    delete[] buffer;
    
    return 0;
}
```

### 3.5 常见错误与注意事项

#### 3.5.1 错误1：忘记手动析构

```cpp
#include <iostream>
#include <new>

class HasResource {
private:
    int* data;
public:
    HasResource() {
        data = new int[10];
        std::cout << "HasResource ctor" << std::endl;
    }
    
    ~HasResource() {
        delete[] data;
        std::cout << "HasResource dtor" << std::endl;
    }
};

void bad_example() {
    char* buf = new char[sizeof(HasResource)];
    HasResource* obj = new (buf) HasResource;
    
    // 错误：忘记调用析构函数
    delete[] buf;  // 内存泄漏！data 没有被释放
}

void good_example() {
    char* buf = new char[sizeof(HasResource)];
    HasResource* obj = new (buf) HasResource;
    
    // 正确：先析构
    obj->~HasResource();
    delete[] buf;
}
```

#### 3.5.2 错误2：对齐问题

```cpp
#include <iostream>
#include <new>

class AlignedObject {
private:
    double data[10];  // 需要8字节对齐
public:
    AlignedObject() {
        std::cout << "AlignedObject ctor" << std::endl;
    }
};

void bad_alignment() {
    // 可能对齐不正确
    char buffer[sizeof(AlignedObject)];
    AlignedObject* obj = new (buffer) AlignedObject;  // 可能崩溃
    
    obj->~AlignedObject();
}

void good_alignment() {
    // 正确做法：使用 C++11 的 alignas
    alignas(AlignedObject) char buffer[sizeof(AlignedObject)];
    AlignedObject* obj = new (buffer) AlignedObject;
    
    obj->~AlignedObject();
}
```

#### 3.5.3 错误3：使用 delete 代替手动析构

```cpp
#include <iostream>

class WrongUsage {
public:
    WrongUsage() {
        std::cout << "ctor" << std::endl;
    }
    ~WrongUsage() {
        std::cout << "dtor" << std::endl;
    }
};

void wrong_delete() {
    char* buf = new char[sizeof(WrongUsage)];
    WrongUsage* obj = new (buf) WrongUsage;
    
    // 非常错误！
    delete obj;  // 试图释放不是从 new 得到的内存
    
    // 这可能导致双重释放或内存破坏
}

void correct_usage() {
    char* buf = new char[sizeof(WrongUsage)];
    WrongUsage* obj = new (buf) WrongUsage;
    
    // 正确做法
    obj->~WrongUsage();
    delete[] buf;
}
```

---

## 四、 C++新标准的更新与改进

### 4.1 C++11：std::allocator 的改进

#### 4.1.1 std::allocator 的 construct/destroy

```cpp
#include <memory>
#include <vector>

// C++11 之前：需要自己处理 placement new
template <typename T>
class OldAllocator {
public:
    void construct(T* p, const T& val) {
        new (p) T(val);  // 只能拷贝构造
    }
};

// C++11 之后：支持变参模板和完美转发
template <typename T>
class NewAllocator {
public:
    template <typename... Args>
    void construct(T* p, Args&&... args) {
        new (p) T(std::forward<Args>(args)...);  // 完美转发
    }
};

// 使用示例
void use_construct() {
    std::allocator<int> alloc;
    int* p = alloc.allocate(1);
    
    // C++11 之前
    alloc.construct(p, 42);  // 只能用现有对象构造
    
    // C++11 之后
    alloc.construct(p, 42);  // 直接传参数给构造函数
    
    alloc.destroy(p);
    alloc.deallocate(p, 1);
}
```

### 4.2 C++11：std::aligned_storage

#### 4.2.1 使用 std::aligned_storage 实现对齐

```cpp
#include <iostream>
#include <type_traits>

class MyType {
private:
    double data;
public:
    MyType(double d) : data(d) {
        std::cout << "MyType ctor(" << d << ")" << std::endl;
    }
    ~MyType() {
        std::cout << "MyType dtor" << std::endl;
    }
};

void use_aligned_storage() {
    // 使用 std::aligned_storage
    std::aligned_storage<sizeof(MyType), alignof(MyType)>::type buffer;
    
    // 在对齐的内存上构造
    MyType* obj = new (&buffer) MyType(3.14);
    
    // 使用
    // ...
    
    // 析构
    obj->~MyType();
}

// C++14 的简化写法
#include <type_traits>

void cxx14_aligned_storage() {
    // C++14 提供了 alias template
    std::aligned_storage_t<sizeof(MyType), alignof(MyType)> buffer;
    
    MyType* obj = new (&buffer) MyType(2.718);
    obj->~MyType();
}
```

### 4.3 C++11：std::vector 的 emplace_back

#### 4.3.1 使用 emplace_back 避免 placement new

```cpp
#include <vector>
#include <iostream>

class Complex {
private:
    int re, im;
public:
    Complex(int r, int i) : re(r), im(i) {
        std::cout << "Complex(" << r << "," << i << ")" << std::endl;
    }
};

void use_emplace() {
    std::vector<Complex> vec;
    vec.reserve(3);  // 预分配空间
    
    // 直接在向量内存中构造，不需要显式使用 placement new
    vec.emplace_back(1, 2);
    vec.emplace_back(3, 4);
    vec.emplace_back(5, 6);
}
```

### 4.4 C++17：std::pmr 内存资源

#### 4.4.1 使用 pmr（Polymorphic Memory Resources）

```cpp
#include <iostream>
#include <memory_resource>
#include <vector>

void use_pmr() {
    // 使用单调缓冲区资源（适合频繁分配）
    std::pmr::monotonic_buffer_resource pool;
    
    // 使用 pmr 分配器的 vector
    std::pmr::vector<int> vec(&pool);
    
    // 所有内存分配都来自 pool
    for (int i = 0; i < 100; ++i) {
        vec.push_back(i);
    }
}
```

---

## 五、 面试高频问题与解答

### 5.1 基础问题

#### 5.1.1 Array new 和普通 new 的区别？

**答案：**

| 特性 | new Type | new Type[] |
|------|---------|-----------|
| 分配大小 | sizeof(Type) | sizeof(Type)*N + overhead |
| 构造函数调用 | 1次 | N次 |
| 析构函数调用 | 1次 | N次 |
| Cookie | 通常不需要 | 通常需要（用于记录N） |
| 内存分配 | operator new | operator new[] |
| 内存释放 | operator delete | operator delete[] |

#### 5.1.2 为什么 new[] 需要默认构造函数？

**答案：**

因为：
1. Array new 需要一次构造多个对象
2. 无法在语法上给每个对象分别提供参数
3. 所以只能逐个调用默认构造函数

**补充：**

- C++11 之前：如果没有默认构造函数，就无法使用 Array new
- C++11 之后：可以用 vector 和 emplace_back 来绕过这个限制
- 也可以手动分配内存后用 placement new 逐个构造

#### 5.1.3 Placement new 是什么？有什么用？

**答案：**

**定义：**

Placement new 是 operator new 的特殊重载形式，允许在**已分配的内存**上构造对象，而不分配新内存。

**签名：**

```cpp
void* operator new(size_t size, void* ptr);
```

**作用：**

1. **内存池/对象池**：预分配内存，反复使用
2. **栈上构造**：在栈上构造非 POD 类型对象
3. **自定义内存管理**：特殊的内存布局需求
4. **性能优化**：避免频繁的内存分配

### 5.2 进阶问题

#### 5.2.1 Cookie 是什么？什么时候需要？

**答案：**

**定义：**

Cookie 是编译器在数组分配的内存块头部记录的额外信息，主要用于记录数组大小。

**什么时候需要：**

| 情况 | 是否需要 Cookie | 原因 |
|-----|---------------|------|
| 类有 non-trivial 析构函数 | 是 | delete[] 需要知道调用多少次析构函数 |
| 类只有 trivial 析构函数 | 可能不需要 | 不需要调用析构函数 |
| 简单类型（int, char） | 通常不需要 | 不需要调用析构函数 |

**信息包含：**

- 数组大小（元素个数）
- 其他内存管理信息
- 调试信息（调试模式下）

#### 5.2.2 如果 new[] 和 delete 不匹配会怎样？

**答案：**

**后果取决于类的类型：**

1. **简单类型（int等）**：
   - 可能不会立即崩溃
   - 但仍然是未定义行为

2. **有指针成员、有non-trivial析构的类**：
   - 只调用第一个对象的析构函数
   - 其他对象的资源泄漏
   - 可能导致内存破坏

3. **重载了 operator new/delete[] 的类**：
   - 可能造成严重的内存错误
   - 调试模式下可能触发断言

**正确做法：**

- 始终用 delete[] 来释放 new[] 分配的内存
- 始终用 delete 来释放 new 分配的内存

#### 5.2.3 如何实现安全的 Array new 替代方案？

**答案：**

**方案1：使用 std::vector（推荐）**

```cpp
#include <vector>

std::vector<MyClass> vec(100);  // 自动管理，安全
```

**方案2：使用 std::unique_ptr（C++11）**

```cpp
#include <memory>

auto arr = std::make_unique<MyClass[]>(100);  // 自动调用 delete[]
```

**方案3：手写包装类**

```cpp
template <typename T>
class Array {
private:
    T* data;
    size_t count;
public:
    Array(size_t n) : count(n) {
        data = static_cast<T*>(operator new[](sizeof(T) * n));
        for (size_t i = 0; i < n; ++i) {
            new (&data[i]) T();
        }
    }
    
    ~Array() {
        for (size_t i = count; i-- > 0;) {
            data[i].~T();
        }
        operator delete[](data);
    }
    
    T& operator[](size_t i) { return data[i]; }
    const T& operator[](size_t i) const { return data[i]; }
};
```

### 5.3 实际问题

#### 5.3.1 如何检测 new[] 和 delete 是否匹配？

**答案：**

**方法1：代码审查**

- 检查所有 new[] 是否都匹配了 delete[]
- 使用代码检查工具

**方法2：调试工具**

- 使用内存检查工具（如 AddressSanitizer）
- 查看内存泄漏和错误

**方法3：使用智能指针（C++11）**

```cpp
// 使用 unique_ptr 自动匹配
auto arr = std::unique_ptr<int[]>(new int[10]);
```

#### 5.3.2 在什么情况下需要手动使用 placement new？

**答案：**

**场景1：内存池**

需要反复创建销毁对象，避免频繁 malloc/free。

**场景2：自定义内存布局**

需要在特定位置或特定类型的内存上构造对象。

**场景3：对象重用**

销毁对象后不释放内存，直接在原地构造新对象。

**场景4：在共享内存中构造对象**

多个进程间的共享内存区域需要手动管理对象生命周期。

**场景5：性能关键路径**

需要精确控制内存分配时机，避免动态分配开销。

---

## 六、 代码示例与分析

### 6.1 完整的内存池实现

#### 6.1.1 代码实现

```cpp
#include <iostream>
#include <list>
#include <new>

template <typename T>
class MemoryPool {
private:
    struct Block {
        Block* next;
    };
    
    std::list<char*> allocated_blocks;  // 管理所有分配的大内存块
    Block* free_list;                   // 空闲对象链表
    size_t block_size;
    const size_t POOL_SIZE = 10;         // 每次分配的对象数量
    
public:
    MemoryPool() : free_list(nullptr), block_size(sizeof(T)) {
        allocate_block();
    }
    
    ~MemoryPool() {
        // 释放所有大内存块
        for (char* block : allocated_blocks) {
            delete[] block;
        }
    }
    
    void allocate_block() {
        // 分配一大块内存
        size_t total_size = block_size * POOL_SIZE;
        char* new_block = new char[total_size];
        allocated_blocks.push_back(new_block);
        
        // 将块分割成单个对象，加入空闲链表
        for (size_t i = 0; i < POOL_SIZE; ++i) {
            Block* b = reinterpret_cast<Block*>(new_block + i * block_size);
            b->next = free_list;
            free_list = b;
        }
    }
    
    void* allocate() {
        if (free_list == nullptr) {
            allocate_block();
        }
        
        Block* result = free_list;
        free_list = free_list->next;
        return result;
    }
    
    void deallocate(void* p) {
        Block* b = reinterpret_cast<Block*>(p);
        b->next = free_list;
        free_list = b;
    }
};

class PoolTest {
private:
    int id;
public:
    static MemoryPool<PoolTest> pool;
    
    PoolTest(int i) : id(i) {
        std::cout << "ctor: PoolTest(" << i << ")" << std::endl;
    }
    
    ~PoolTest() {
        std::cout << "dtor: ~PoolTest(" << id << ")" << std::endl;
    }
    
    void* operator new(size_t size) {
        std::cout << "operator new (pool)" << std::endl;
        return pool.allocate();
    }
    
    void operator delete(void* p) noexcept {
        std::cout << "operator delete (pool)" << std::endl;
        pool.deallocate(p);
    }
    
    int getId() const { return id; }
};

// 静态成员初始化
MemoryPool<PoolTest> PoolTest::pool;

int main() {
    std::cout << "=== Memory Pool Test ===" << std::endl;
    
    // 创建对象
    PoolTest* p1 = new PoolTest(1);
    PoolTest* p2 = new PoolTest(2);
    PoolTest* p3 = new PoolTest(3);
    
    std::cout << "p1: " << p1->getId() << std::endl;
    std::cout << "p2: " << p2->getId() << std::endl;
    std::cout << "p3: " << p3->getId() << std::endl;
    
    // 释放对象（放回内存池）
    delete p1;
    delete p2;
    delete p3;
    
    // 再次创建（复用内存）
    std::cout << "\n=== Reuse objects ===" << std::endl;
    PoolTest* p4 = new PoolTest(4);
    PoolTest* p5 = new PoolTest(5);
    
    std::cout << "p4: " << p4->getId() << std::endl;
    std::cout << "p5: " << p5->getId() << std::endl;
    
    delete p4;
    delete p5;
    
    return 0;
}
```

#### 6.1.2 代码分析

**实现思路：**

1. **预分配大内存块**：一次分配多个对象的空间，减少 malloc 调用
2. **空闲链表管理**：将未使用的内存块组织成链表，快速分配
3. **自定义 operator new/delete**：让类自动使用内存池

**核心逻辑：**

- `allocate_block()`：分配大块内存，分割成多个小对象
- `allocate()`：从空闲链表取一个对象
- `deallocate()`：释放对象，放回空闲链表

**优势：**

1. **高性能**：避免频繁的 malloc/free
2. **内存复用**：同一个内存可以反复使用
3. **减少碎片**：固定大小的块不容易产生碎片

#### 6.1.3 优化建议

**当前实现的改进空间：**

1. **线程安全**：加入互斥锁，支持多线程
2. **动态扩展**：根据需要自动调整预分配大小
3. **内存对齐**：确保对齐正确
4. **调试支持**：加入分配计数和泄漏检测

**线程安全版本：**

```cpp
template <typename T>
class ThreadSafeMemoryPool {
private:
    std::list<char*> allocated_blocks;
    Block* free_list;
    std::mutex mutex;
    
public:
    void* allocate() {
        std::lock_guard<std::mutex> lock(mutex);
        // 原有的 allocate 逻辑...
    }
    
    void deallocate(void* p) {
        std::lock_guard<std::mutex> lock(mutex);
        // 原有的 deallocate 逻辑...
    }
};
```

### 6.2 Array new 的完整示例

#### 6.2.1 代码实现

```cpp
#include <iostream>
#include <new>
#include <cassert>

// 演示类
class MyClass {
private:
    int id;
    char* buffer;
public:
    // 默认构造函数
    MyClass() : id(0) {
        buffer = new char[10];
        std::cout << "ctor: MyClass() id=" << id << std::endl;
    }
    
    MyClass(int i) : id(i) {
        buffer = new char[10];
        std::cout << "ctor: MyClass(" << i << ")" << std::endl;
    }
    
    ~MyClass() {
        delete[] buffer;
        std::cout << "dtor: ~MyClass() id=" << id << std::endl;
    }
    
    void setId(int i) {
        id = i;
    }
    
    int getId() const {
        return id;
    }
};

// 测试基本用法
void test_basic_array() {
    std::cout << "=== Basic Array new ===" << std::endl;
    MyClass* arr = new MyClass[5];
    
    for (int i = 0; i < 5; ++i) {
        arr[i].setId(i);
        std::cout << "arr[" << i << "].id = " << arr[i].getId() << std::endl;
    }
    
    delete[] arr;
}

// 测试 placement new 逐个构造
void test_placement_array() {
    std::cout << "\n=== Placement new Array ===" << std::endl;
    
    const int size = 3;
    
    // 1. 分配原始内存
    char* raw = new char[sizeof(MyClass) * size];
    MyClass* arr = reinterpret_cast<MyClass*>(raw);
    
    // 2. 逐个构造
    for (int i = 0; i < size; ++i) {
        new (&arr[i]) MyClass(i + 100);  // 使用带参数的构造函数
    }
    
    // 3. 使用
    for (int i = 0; i < size; ++i) {
        std::cout << "arr[" << i << "].id = " << arr[i].getId() << std::endl;
    }
    
    // 4. 逐个析构（逆序）
    for (int i = size - 1; i >= 0; --i) {
        arr[i].~MyClass();
    }
    
    // 5. 释放原始内存
    delete[] raw;
}

// 测试错误用法
void test_bad_usage() {
    std::cout << "\n=== Bad Usage (Demonstration) ===" << std::endl;
    
    // 错误示例（实际中不要这么做）
    MyClass* arr = new MyClass[3];
    
    // delete arr;  // 错误！只调用一个析构函数
    // 应该使用 delete[]
    delete[] arr;
}

int main() {
    test_basic_array();
    test_placement_array();
    test_bad_usage();
    
    return 0;
}
```

#### 6.2.2 代码分析

**设计思路：**

1. **MyClass 设计**：有默认构造函数，也有带参数的构造函数
2. **Basic Array new**：展示标准的 Array new/delete 用法
3. **Placement Array new**：展示如何手动分配内存，逐个构造带参数的对象
4. **错误用法演示**：展示不匹配使用的危害

**关键点：**

1. **逆序析构**：析构函数通常按构造的逆序调用
2. **原始内存管理**：使用 placement new 时，需要自己管理原始内存
3. **带参数构造**：通过 placement new 可以调用任意构造函数

### 6.3 Placement new 的实际应用

#### 6.3.1 对象池实现

```cpp
#include <iostream>
#include <new>
#include <vector>

template <typename T>
class ObjectPool {
private:
    struct Slot {
        bool in_use;
        union {
            T data;
            char dummy;
        };
    };
    
    std::vector<Slot*> pool;
    size_t grow_size;
    
public:
    ObjectPool(size_t grow_size = 10) : grow_size(grow_size) {}
    
    ~ObjectPool() {
        // 先析构所有正在使用的对象
        for (Slot* slot : pool) {
            if (slot->in_use) {
                slot->data.~T();
            }
        }
        // 再释放内存
        for (Slot* slot : pool) {
            delete slot;
        }
    }
    
    void grow() {
        for (size_t i = 0; i < grow_size; ++i) {
            Slot* slot = new Slot;
            slot->in_use = false;
            pool.push_back(slot);
        }
    }
    
    T* allocate() {
        // 找空闲的 slot
        for (Slot* slot : pool) {
            if (!slot->in_use) {
                slot->in_use = true;
                return &slot->data;
            }
        }
        
        // 没有空闲的，需要扩展
        grow();
        return allocate();
    }
    
    void deallocate(T* p) {
        // 找到对应的 slot
        for (Slot* slot : pool) {
            if (&slot->data == p) {
                p->~T();  // 先析构
                slot->in_use = false;
                return;
            }
        }
    }
};

// 演示类
class PooledObject {
private:
    int data[100];
public:
    PooledObject(int d) {
        data[0] = d;
        std::cout << "PooledObject ctor" << std::endl;
    }
    
    ~PooledObject() {
        std::cout << "PooledObject dtor" << std::endl;
    }
    
    void setData(int d) {
        data[0] = d;
    }
};

int main() {
    ObjectPool<PooledObject> pool;
    
    std::cout << "=== Object Pool Test ===" << std::endl;
    
    // 从池中获取对象
    PooledObject* p1 = pool.allocate();
    new (p1) PooledObject(1);
    
    PooledObject* p2 = pool.allocate();
    new (p2) PooledObject(2);
    
    // 使用
    p1->setData(100);
    p2->setData(200);
    
    // 释放回池中
    pool.deallocate(p1);
    pool.deallocate(p2);
    
    // 再次获取（重用）
    PooledObject* p3 = pool.allocate();
    new (p3) PooledObject(3);
    
    pool.deallocate(p3);
    
    return 0;
}
```

---

## 七、 总结与反思

### 7.1 核心知识点总结

| 知识点 | 要点 |
|--------|------|
| Array new | 分配数组，调用多次构造函数，需要默认构造 |
| Array delete | 逆序调用析构函数，释放内存 |
| Cookie | 记录数组大小，用于正确调用析构函数 |
| Placement new | 在已分配内存上构造对象，不分配新内存 |
| 对齐问题 | 内存必须正确对齐，C++11 提供 alignas |
| 内存池 | 使用 placement new 实现对象重用 |
| 重载 operator new[] | 需要特别注意指针偏移问题 |

### 7.2 学习收获

1. **深入理解 Array new 的底层机制**：Cookie 的作用和布局
2. **掌握 Placement new 的正确用法**：理解手动析构的必要性
3. **了解内存块布局**：调试模式下的内存结构
4. **避免常见错误**：不匹配的 delete，忘记手动析构等

### 7.3 后续学习建议

1. **深入学习 STL 分配器**：了解 std::allocator 的实现原理
2. **学习现代 C++ 内存管理**：C++11/14/17 的智能指针和 pmr
3. **实现自己的内存池**：巩固 placement new 的使用
4. **学习内存检查工具**：AddressSanitizer, Valgrind 等

### 7.4 实践建议

1. **避免直接使用 Array new**：优先使用 std::vector 等容器
2. **使用智能指针**：std::unique_ptr<T[]> 可以自动匹配 delete[]
3. **谨慎使用 placement new**：只有在必要时才用，确保正确性
4. **内存对齐**：使用 C++11 的 alignas 确保对齐正确

---

通过本课程的学习，我们深入理解了 C++ 中数组内存管理和 placement new 的底层机制，掌握了相关的最佳实践和常见错误的避免方法。这些知识对于编写高效、健壮的 C++ 程序至关重要，也是面试中的高频考点。
