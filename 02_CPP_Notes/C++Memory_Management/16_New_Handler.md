# 侯捷 C++ 课程第16集学习笔记：New Handler（new 失败处理器）

## 一、课程概述

### 1.1 课程主题

本课程深入讲解 **new handler**（new 失败处理器）机制，包括其概念、作用、设置方法以及实际应用。同时介绍 C++11 引入的 `=default` 和 `=delete` 语法。

### 1.2 课程内容结构

| 主题 | 核心内容 |
|------|---------|
| new handler 概念 | 当 operator new 分配内存失败时的处理机制 |
| set_new_handler | 设置自定义的 new handler 函数 |
| nothrow new | 不抛异常的 new 形式 |
| =default / =delete | C++11 控制默认函数生成 |

### 1.3 学习目标

1. 理解 new handler 的作用和工作机制
2. 掌握 set_new_handler 的使用方法
3. 理解 nothrow new 的用法和适用场景
4. 掌握 =default 和 =delete 的语法和应用
5. 了解 operator new 的底层实现
6. 掌握面试中相关的高频问题

### 1.4 参考来源

- C++ Primer 3/e, p.765
- Effective C++ 2/e, item 7
- SGI STL 源码分析

---

## 二、New Handler 核心概念

### 2.1 什么是 New Handler？

**定义：**
New handler 是一个函数指针类型，当 `operator new` 无法分配内存时，会在抛出 `std::bad_alloc` 异常之前调用这个函数。

```cpp
// new_handler 的类型定义
typedef void (*new_handler)();
```

**作用：**
- 给程序一个机会来释放一些内存，以便后续的 new 操作能够成功
- 或者记录日志、发送警告等
- 最后可以选择调用 `abort()` 或 `exit()` 终止程序

### 2.2 operator new 的底层实现

从截图中可以看到 operator new 的源码结构：

```cpp
void* operator new(size_t size, const std::nothrow_t&) THROW0() {
    void* p;
    // 循环尝试分配内存
    while ((p = malloc(size)) == 0) {
        // 尝试释放更多内存或返回空指针
        _TRY_BEGIN
            if (_callnewh(size) == 0) break;
        _CATCH(std::bad_alloc) return 0;
        _CATCH_END
    }
    return p;
}
```

**执行流程：**

```
┌─────────────────────────────────────────────────────────┐
│              operator new 执行流程                      │
├─────────────────────────────────────────────────────────┤
│  1. 尝试用 malloc 分配内存                              │
│     │                                                  │
│     ├── 成功 → 返回指针                                │
│     │                                                  │
│     └── 失败 → 调用 new_handler                        │
│           │                                            │
│           ├── new_handler 释放了内存 → 再次尝试 malloc │
│           │                                            │
│           └── new_handler 没释放内存 → 抛出异常或返回0  │
└─────────────────────────────────────────────────────────┘
```

### 2.3 设置 New Handler

使用 `std::set_new_handler` 函数设置自定义的 new handler：

```cpp
#include <new>  // 必须包含这个头文件

// 自定义的 new handler 函数
void myNewHandler() {
    std::cerr << "Out of memory!" << std::endl;
    std::abort();  // 终止程序
}

int main() {
    // 设置 new handler
    std::set_new_handler(myNewHandler);
    
    // 尝试分配大量内存
    int* p = new int[1000000000000000LL];
    
    return 0;
}
```

### 2.4 Nothrow New

**nothrow new** 是不抛出异常的 new 形式，当分配失败时返回 `nullptr` 而不是抛出 `std::bad_alloc` 异常。

```cpp
#include <new>

// nothrow new 的使用
int* p = new (std::nothrow) int[1000000000000000LL];

if (p == nullptr) {
    // 分配失败，处理错误
    std::cerr << "Memory allocation failed!" << std::endl;
}
```

**注意：**
- nothrow new 仍然会调用 new handler
- 只有当 new handler 无法释放内存时，才返回 nullptr

---

## 三、完整代码示例

### 3.1 New Handler 基本示例

```cpp
#include <new>
#include <iostream>
#include <cstdlib>  // for abort()
using namespace std;

// ==========================================
// 自定义 new handler 函数
// ==========================================
void noMoreMemory() {
    cerr << "Error: out of memory!" << endl;
    abort();  // 终止程序
}

int main() {
    // ==========================================
    // 设置 new handler
    // ==========================================
    set_new_handler(noMoreMemory);
    
    // ==========================================
    // 尝试分配大量内存（肯定会失败）
    // ==========================================
    try {
        // 分配 1000000000000000 个 int，约 4TB！
        int* p = new int[1000000000000000LL];
        assert(p);  // 不应该执行到这里
        delete[] p;
    } catch (const bad_alloc& e) {
        // 如果 new handler 没有调用 abort，会到这里
        cerr << "Caught bad_alloc: " << e.what() << endl;
    }
    
    return 0;
}
```

### 3.2 输出结果分析

**执行结果（BCB4 编译器）：**
```
Error: out of memory!
Abnormal program termination
```

**说明：**
- 当 new 无法分配内存时，会调用我们设置的 `noMoreMemory` 函数
- `noMoreMemory` 输出错误信息并调用 `abort()` 终止程序
- 这是符合预期的行为，表示 operator new 无法满足内存申请

---

## 四、设计良好的 New Handler

### 4.1 New Handler 的两个选择

根据截图中的内容，设计良好的 new handler 只有两个选择：

1. **让更多内存可用**
   - 释放一些不再需要的内存
   - 关闭一些文件句柄
   - 清理缓存等

2. **调用 abort() 或 exit()**
   - 如果无法释放内存，优雅地终止程序

### 4.2 循环调用机制

**重要特性：**
- new handler 可能被多次调用
- 每次调用后，operator new 会再次尝试分配内存
- 直到分配成功或 new handler 终止程序

**示例：释放缓存的 new handler**

```cpp
#include <new>
#include <iostream>
#include <vector>
using namespace std;

// 全局缓存
vector<int>* g_cache = nullptr;

// ==========================================
// 智能的 new handler：释放缓存
// ==========================================
void releaseCacheHandler() {
    cerr << "Releasing cache to free memory..." << endl;
    
    if (g_cache != nullptr) {
        delete g_cache;
        g_cache = nullptr;
        cerr << "Cache released successfully!" << endl;
    } else {
        // 没有更多内存可释放，终止程序
        cerr << "No more memory to release. Aborting..." << endl;
        abort();
    }
}

int main() {
    // 初始化缓存
    g_cache = new vector<int>(100000000);  // 约 400MB
    
    // 设置 new handler
    set_new_handler(releaseCacheHandler);
    
    // 第一次分配：可能成功（如果有足够内存）
    int* p1 = new(nothrow) int[500000000];  // 约 2GB
    if (p1 == nullptr) {
        cerr << "First allocation failed" << endl;
    } else {
        cerr << "First allocation succeeded" << endl;
        delete[] p1;
    }
    
    // 第二次分配：应该会触发 new handler
    int* p2 = new(nothrow) int[500000000];
    if (p2 == nullptr) {
        cerr << "Second allocation failed" << endl;
    } else {
        cerr << "Second allocation succeeded after cache release" << endl;
        delete[] p2;
    }
    
    return 0;
}
```

---

## 五、C++11 的 =default 和 =delete

### 5.1 基本语法

**=default：** 显式要求编译器生成默认版本的函数

**=delete：** 显式禁止编译器生成默认版本的函数

```cpp
class Foo {
public:
    // =default：使用编译器生成的默认构造函数
    Foo() = default;
    
    // =delete：禁止拷贝构造
    Foo(const Foo&) = delete;
    
    // =delete：禁止拷贝赋值
    Foo& operator=(const Foo&) = delete;
    
    // =default：使用编译器生成的析构函数
    ~Foo() = default;
};
```

### 5.2 应用场景

#### 场景1：禁止拷贝（单例模式）

```cpp
class Singleton {
private:
    Singleton() {}  // 私有构造函数
    
public:
    // 禁止拷贝
    Singleton(const Singleton&) = delete;
    Singleton& operator=(const Singleton&) = delete;
    
    static Singleton& getInstance() {
        static Singleton instance;
        return instance;
    }
};
```

#### 场景2：禁止某些 operator new/delete

```cpp
class Goo {
public:
    long x;
    
    Goo(long x = 0) : x(x) {}
    
    // 禁止单个对象的 new，但允许数组 new
    static void* operator new(size_t size) = delete;
    static void operator delete(void* p, size_t size) = delete;
    
    // 允许数组 new/delete
    static void* operator new[](size_t size) = default;
    static void operator delete[](void* p, size_t size) = default;
};
```

**注意：**
- `operator new` 和 `operator delete` 不能用 `=default`
- 截图中显示 `cannot be defaulted` 的错误

### 5.3 错误示例

根据截图中的内容，以下用法是错误的：

```cpp
class Foo {
public:
    long x;
    
    Foo(long x = 0) : x(x) {}
    
    // ❌ 错误：operator new 不能用 =default
    static void* operator new(size_t size) = default;
    
    // ❌ 错误：operator delete 不能用 =default
    static void operator delete(void* p, size_t size) = default;
    
    // ✅ 正确：可以用 =delete
    static void operator new[](size_t size) = delete;
    
    // ✅ 正确：可以用 =delete
    static void operator delete[](void* p, size_t size) = delete;
};
```

### 5.4 使用限制

| 函数类型 | =default | =delete |
|---------|---------|---------|
| 默认构造函数 | ✅ | ✅ |
| 拷贝构造函数 | ✅ | ✅ |
| 拷贝赋值运算符 | ✅ | ✅ |
| 移动构造函数 | ✅ | ✅ |
| 移动赋值运算符 | ✅ | ✅ |
| 析构函数 | ✅ | ❌（不允许） |
| operator new | ❌ | ✅ |
| operator delete | ❌ | ✅ |
| operator new[] | ❌ | ✅ |
| operator delete[] | ❌ | ✅ |

---

## 六、深入理解 operator new 的源码

### 6.1 源码分析

从截图中可以看到 operator new 的核心实现：

```cpp
void* operator new(size_t size, const std::nothrow_t&) THROW0() {
    void* p;
    
    // 循环尝试分配
    while ((p = malloc(size)) == 0) {
        // 尝试调用 new handler
        _TRY_BEGIN
            if (_callnewh(size) == 0) 
                break;  // new handler 返回 0，表示无法释放内存
        _CATCH(std::bad_alloc) 
            return 0;  // 捕获异常，返回 null
        _CATCH_END
    }
    
    return p;
}
```

### 6.2 _callnewh 函数

`_callnewh` 函数负责调用用户设置的 new handler：

```cpp
// 简化版实现
int _callnewh(size_t size) {
    // 获取当前的 new handler
    new_handler handler = get_new_handler();
    
    if (handler != nullptr) {
        // 调用 new handler
        handler();
        // 如果返回，说明 handler 没有终止程序
        // 返回非 0 表示继续尝试
        return 1;
    }
    
    // 没有设置 handler，返回 0
    return 0;
}
```

### 6.3 执行流程图解

```
┌────────────────────────────────────────────────────────────┐
│                  operator new 执行流程                      │
├────────────────────────────────────────────────────────────┤
│                                                            │
│  ┌──────────────┐                                          │
│  │ malloc(size) │                                          │
│  └──────┬───────┘                                          │
│         │                                                  │
│    ┌────┴────┐                                             │
│    │ 成功    │ 失败                                        │
│    ↓         ↓                                             │
│  ┌────────┐ ┌─────────────────────┐                        │
│  │返回指针│ │_callnewh(size)调用   │                        │
│  └────────┘ └──────────┬──────────┘                        │
│                        │                                   │
│              ┌─────────┴─────────┐                         │
│              │ handler 返回非 0   │ handler 返回 0 或终止   │
│              ↓                   ↓                         │
│        ┌─────────────┐    ┌──────────────┐                │
│        │再次尝试malloc│    │抛出异常或返回0│                │
│        └──────┬──────┘    └──────────────┘                │
│               │                                            │
│               ↓                                            │
│        ┌─────────────┐                                     │
│        │回到循环开始 │                                     │
│        └─────────────┘                                     │
│                                                            │
└────────────────────────────────────────────────────────────┘
```

---

## 七、C++ 新标准的更新与改进

### 7.1 C++11 的改进

#### 7.1.1 =default 和 =delete

C++11 引入的这两个语法让程序员能够更好地控制默认函数的生成：

```cpp
class MyClass {
public:
    // C++11：显式默认
    MyClass() = default;
    ~MyClass() = default;
    
    // C++11：显式删除
    MyClass(const MyClass&) = delete;
    MyClass& operator=(const MyClass&) = delete;
    
    // C++11：禁用某些操作
    void* operator new(size_t) = delete;
};
```

#### 7.1.2 noexcept 说明符

C++11 引入 `noexcept` 说明符，表示函数不会抛出异常：

```cpp
// 不会抛出异常的 new handler
void myHandler() noexcept {
    std::cerr << "Out of memory" << std::endl;
    std::abort();
}
```

### 7.2 C++17 的改进

#### 7.2.1 std::bad_alloc 的改进

C++17 为 `std::bad_alloc` 添加了更多信息：

```cpp
#include <new>
#include <iostream>

void handler() {
    // 获取更多内存分配失败的信息
    std::cerr << "Allocation failed" << std::endl;
    std::abort();
}

int main() {
    std::set_new_handler(handler);
    
    try {
        while (true) {
            new int[1000000000];
        }
    } catch (const std::bad_alloc& e) {
        // C++17 可以获取更多信息
        std::cerr << "Caught: " << e.what() << std::endl;
    }
}
```

### 7.3 C++20 的改进

#### 7.3.1 std::new_handler 类型别名

C++20 将 `new_handler` 定义为类型别名：

```cpp
// C++20
namespace std {
    using new_handler = void(*)();
}
```

---

## 八、面试高频问题与解答

### 8.1 基础问题

#### Q1: New handler 是什么？有什么作用？

**答案：**

**定义：**
New handler 是一个函数指针，当 `operator new` 无法分配内存时，会在抛出异常之前调用它。

**作用：**
1. **释放内存**：给程序一个机会释放不再需要的内存
2. **记录日志**：记录内存分配失败的信息
3. **优雅终止**：如果无法释放内存，优雅地终止程序

**核心代码：**
```cpp
typedef void (*new_handler)();
std::new_handler set_new_handler(std::new_handler p) throw();
```

#### Q2: Nothrow new 和普通 new 有什么区别？

**答案：**

| 特性 | 普通 new | Nothrow new |
|------|---------|-------------|
| 分配失败 | 抛出 std::bad_alloc | 返回 nullptr |
| 语法 | `new T` | `new (std::nothrow) T` |
| 是否调用 new handler | 是 | 是 |
| 需要的头文件 | 不需要 | `<new>` |

**使用场景：**
- **普通 new**：需要异常处理的场景
- **Nothrow new**：不需要异常处理，或者在不能使用异常的环境中

#### Q3: =default 和 =delete 的作用是什么？

**答案：**

**=default：**
- 显式要求编译器生成默认版本的函数
- 适用于构造函数、析构函数、拷贝/移动操作

**=delete：**
- 显式禁止编译器生成默认版本的函数
- 适用于禁止拷贝、禁止某些操作

**示例：**
```cpp
class NoCopy {
public:
    NoCopy() = default;
    NoCopy(const NoCopy&) = delete;
    NoCopy& operator=(const NoCopy&) = delete;
};
```

### 8.2 进阶问题

#### Q4: operator new 在分配失败时会做什么？

**答案：**

**执行流程：**
1. 尝试用 `malloc` 分配内存
2. 如果失败，调用 `new_handler`
3. `new_handler` 可以选择：
   - 释放一些内存，然后返回（operator new 会再次尝试分配）
   - 调用 `abort()` 或 `exit()` 终止程序
4. 如果 `new_handler` 返回但没有释放足够内存，重复步骤 1-3
5. 如果 `new_handler` 为空或返回 0，抛出 `std::bad_alloc` 异常

**关键点：**
- new handler 可能被多次调用
- 这是一个循环过程，直到分配成功或程序终止

#### Q5: 为什么 operator new/delete 不能用 =default？

**答案：**

**原因：**
- `operator new` 和 `operator delete` 是特殊的静态成员函数
- 它们的实现涉及底层内存分配（如 malloc/free）
- 编译器无法自动生成它们的默认实现
- 必须由程序员提供实现或使用全局版本

**错误示例：**
```cpp
class Foo {
    // ❌ 错误：不能用 =default
    static void* operator new(size_t size) = default;
};
```

**正确做法：**
```cpp
class Foo {
    // ✅ 使用全局版本
    static void* operator new(size_t size) {
        return ::operator new(size);
    }
};
```

#### Q6: 如何为特定类设置专属的 new handler？

**答案：**

**方法：** 在类的 `operator new` 中设置专属的 new handler

```cpp
#include <new>

class MyClass {
private:
    static std::new_handler old_handler;
    static void myHandler() {
        // 专属的处理逻辑
        std::cerr << "MyClass out of memory!" << std::endl;
        std::abort();
    }
    
public:
    static void* operator new(size_t size) {
        // 保存旧的 handler
        old_handler = std::set_new_handler(myHandler);
        
        void* p = nullptr;
        try {
            p = ::operator new(size);
        } catch (...) {
            // 恢复旧的 handler
            std::set_new_handler(old_handler);
            throw;
        }
        
        // 恢复旧的 handler
        std::set_new_handler(old_handler);
        return p;
    }
};

std::new_handler MyClass::old_handler = nullptr;
```

### 8.3 实际问题

#### Q7: 在实际项目中应该如何使用 new handler？

**答案：**

**建议：**
1. **记录日志**：在 new handler 中记录详细的错误信息
2. **释放缓存**：释放不再需要的内存缓存
3. **优雅降级**：如果无法分配内存，尝试使用更小的数据结构
4. **通知用户**：向用户显示友好的错误消息
5. **避免无限循环**：确保 new handler 最终会终止程序

**示例：**
```cpp
void myNewHandler() {
    // 1. 记录日志
    logError("Memory allocation failed");
    
    // 2. 释放缓存
    releaseCache();
    
    // 3. 检查是否有足够内存
    if (!hasEnoughMemory()) {
        // 4. 通知用户
        showErrorMessage("Out of memory");
        
        // 5. 终止程序
        std::abort();
    }
}
```

#### Q8: 在嵌入式系统中应该如何处理内存分配失败？

**答案：**

**嵌入式系统的特点：**
- 内存有限
- 通常不使用异常
- 需要确定性的行为

**建议：**
1. **使用 nothrow new**：避免异常
2. **检查返回值**：每次 new 后检查是否为 nullptr
3. **预分配内存**：在启动时分配好所需的内存
4. **使用内存池**：避免碎片化
5. **设置合理的 new handler**：在内存不足时优雅处理

**示例：**
```cpp
// 嵌入式系统中的内存分配
int* buffer = new (std::nothrow) int[BUFFER_SIZE];
if (buffer == nullptr) {
    // 处理内存不足的情况
    handleOutOfMemory();
}
```

---

## 九、完整可运行代码

```cpp
#include <new>
#include <iostream>
#include <cstdlib>
#include <cassert>
#include <vector>
using namespace std;

// ==========================================
// 全局缓存（用于演示 new handler 释放内存）
// ==========================================
vector<int>* g_globalCache = nullptr;

// ==========================================
// 自定义 new handler 函数
// ==========================================
void customNewHandler() {
    cerr << "\n=== New Handler Called ===" << endl;
    
    // 尝试释放缓存
    if (g_globalCache != nullptr) {
        cerr << "Releasing global cache..." << endl;
        delete g_globalCache;
        g_globalCache = nullptr;
        cerr << "Cache released successfully!" << endl;
    } else {
        // 没有更多内存可释放
        cerr << "No more memory to release. Aborting..." << endl;
        abort();
    }
}

// ==========================================
// 禁止拷贝的类示例
// ==========================================
class NonCopyable {
public:
    // C++11：显式默认构造函数
    NonCopyable() = default;
    
    // C++11：显式删除拷贝构造和赋值
    NonCopyable(const NonCopyable&) = delete;
    NonCopyable& operator=(const NonCopyable&) = delete;
    
    // C++11：显式默认析构函数
    ~NonCopyable() = default;
};

// ==========================================
// 禁止单个对象 new 的类示例
// ==========================================
class ArrayOnly {
public:
    long value;
    
    ArrayOnly(long v = 0) : value(v) {}
    
    // 禁止单个对象的 new/delete
    static void* operator new(size_t size) = delete;
    static void operator delete(void* p, size_t size) = delete;
    
    // 允许数组的 new/delete
    static void* operator new[](size_t size) {
        cout << "ArrayOnly[] new: " << size << " bytes" << endl;
        return ::operator new[](size);
    }
    
    static void operator delete[](void* p, size_t size) {
        cout << "ArrayOnly[] delete: " << size << " bytes" << endl;
        ::operator delete[](p);
    }
};

// ==========================================
// 主函数测试
// ==========================================
int main() {
    cout << "========================================" << endl;
    cout << "         New Handler 测试" << endl;
    cout << "========================================" << endl;
    
    // ==========================================
    // 测试1：设置 new handler
    // ==========================================
    cout << "\n--- 测试1: New Handler 基本功能 ---" << endl;
    
    // 初始化缓存（占用大量内存）
    cout << "Initializing global cache..." << endl;
    g_globalCache = new vector<int>(10000000);  // 约 40MB
    cout << "Cache initialized. Size: " << g_globalCache->size() << endl;
    
    // 设置 new handler
    cout << "\nSetting new handler..." << endl;
    set_new_handler(customNewHandler);
    
    // 尝试分配大量内存（应该触发 new handler）
    cout << "\nAttempting to allocate large memory..." << endl;
    try {
        // 分配一个非常大的数组
        int* bigArray = new int[1000000000];  // 约 4GB
        cout << "Allocation succeeded! (unexpected)" << endl;
        delete[] bigArray;
    } catch (const bad_alloc& e) {
        cout << "Caught bad_alloc: " << e.what() << endl;
    }
    
    // ==========================================
    // 测试2：Nothrow new
    // ==========================================
    cout << "\n--- 测试2: Nothrow New ---" << endl;
    
    int* nothrowArray = new(nothrow) int[1000000000];
    if (nothrowArray == nullptr) {
        cout << "Nothrow new returned nullptr (expected)" << endl;
    } else {
        cout << "Nothrow new succeeded! (unexpected)" << endl;
        delete[] nothrowArray;
    }
    
    // ==========================================
    // 测试3：=delete 禁止拷贝
    // ==========================================
    cout << "\n--- 测试3: NonCopyable Class ---" << endl;
    
    NonCopyable obj1;
    // NonCopyable obj2 = obj1;  // ❌ 编译错误：拷贝构造被删除
    
    cout << "NonCopyable object created successfully" << endl;
    
    // ==========================================
    // 测试4：=delete 禁止单个 new
    // ==========================================
    cout << "\n--- 测试4: ArrayOnly Class ---" << endl;
    
    // ArrayOnly* single = new ArrayOnly(5);  // ❌ 编译错误：operator new 被删除
    
    ArrayOnly* array = new ArrayOnly[10];  // ✅ 可以
    cout << "ArrayOnly array created successfully" << endl;
    delete[] array;
    
    cout << "\n========================================" << endl;
    cout << "            测试完成！" << endl;
    cout << "========================================" << endl;
    
    return 0;
}
```

---

## 十、总结与反思

### 10.1 核心知识点总结

| 知识点 | 要点 |
|--------|------|
| new handler | 内存分配失败时的回调函数，可释放内存或终止程序 |
| set_new_handler | 设置自定义的 new handler |
| nothrow new | 不抛异常，失败返回 nullptr |
| =default | 显式要求编译器生成默认函数 |
| =delete | 显式禁止编译器生成默认函数 |
| operator new 流程 | malloc → 失败 → 调用 new handler → 循环或抛出异常 |

### 10.2 学习收获

1. **理解了 new handler 的工作机制**
   - operator new 分配失败时的处理流程
   - new handler 的两个选择（释放内存或终止程序）

2. **掌握了 nothrow new 的用法**
   - 在不需要异常的场景中使用
   - 分配失败返回 nullptr

3. **理解了 C++11 的 =default 和 =delete**
   - 控制默认函数的生成
   - 禁止拷贝、禁止某些操作

4. **了解了 operator new 的底层实现**
   - 循环调用 malloc 和 new handler
   - 直到分配成功或程序终止

### 10.3 后续学习建议

1. **学习异常安全**
   - RAII 原则
   - 异常安全的三个级别

2. **学习内存管理最佳实践**
   - 智能指针的使用
   - 内存池的实现

3. **学习嵌入式系统的内存管理**
   - 无异常环境下的内存分配
   - 确定性内存管理

4. **深入学习 STL 的内存分配**
   - std::allocator 的实现
   - SGI STL 的内存池策略

---

通过本课程，我们深入理解了 new handler 的工作机制，掌握了 set_new_handler 和 nothrow new 的使用方法，也学习了 C++11 的 =default 和 =delete 语法。这些知识对于编写健壮的 C++ 程序非常重要！