# 侯捷 C++ 课程 12-13 集学习笔记：Per-class Allocator（类级别内存分配器）

## 一、 课程概述

### 1.1 课程主题

本课程深入讲解 **per-class allocator**（类级别的内存分配器），通过重载成员函数 `operator new/operator delete`，为特定类实现自定义内存管理，提高效率并减少内存碎片。

### 1.2 课程内容结构

| 集数 | 主题 | 核心内容 |
|------|------|---------|
| 第12集 | Per-class allocator 1 | Screen 类的实现（使用 next 指针作为成员） |
| 第13集 | Per-class allocator 2 | Airplane 类的实现（使用 union 更巧妙的设计） |

### 1.3 学习目标

1. 理解为什么需要 per-class allocator
2. 掌握两种 per-class allocator 的实现方式
3. 理解 union 在内存节省中的巧妙应用
4. 了解继承时的内存分配处理
5. 掌握面试中相关的高频问题

### 1.4 参考来源

- C++ Primer 3/e, p.765
- Effective C++ 2/e, item 10

---

## 二、 Per-class Allocator 核心概念

### 2.1 为什么需要 Per-class Allocator？

#### 2.1.1 全局 operator new 的问题

```cpp
// 默认的全局 operator new 会给每个对象加上 cookie
// 导致内存开销较大
```

**问题：**
1. **内存碎片**：频繁 malloc/free 产生碎片
2. **内存开销**：每个对象有 cookie（调试模式下还有其他开销）
3. **效率问题**：频繁调用系统函数较慢

#### 2.1.2 Per-class Allocator 的优势

**优势：**
1. **一次分配大块**：减少 malloc 调用次数
2. **无 cookie**：自定义管理，不需要 cookie
3. **内存紧凑**：对象之间间隔更小
4. **自由链表**：对象释放后回收到链表，可复用

### 2.2 内存间隔对比（重要！）

```
使用全局 operator new：
对象地址间隔 16（有 cookie）

使用 per-class allocator：
对象地址间隔 8（无 cookie，更紧凑）
```

**截图对比：**
```
左邊（自定義）：間隔 8
右邊（全局）：間隔 16
```

### 2.3 实现思路

**核心思想：**
1. **预分配一大块内存**
2. **分割成小对象**
3. **用 free list 链接起来**
4. **分配时从链表取，释放时回收到链表**

---

## 三、 版本一：Screen 类（Next 指针作为成员）

### 3.1 完整代码实现

```cpp
#include <cstddef>
#include <iostream>
using namespace std;

// ==========================================
// Screen 类 - Per-class allocator 版本一
// ==========================================
class Screen {
public:
    // 构造函数
    Screen(int x) : i(x) {};
    
    // 获取成员变量
    int get() const { return i; }
    
    // ==========================================
    // 重载 operator new
    // ==========================================
    static void* operator new(size_t size);
    
    // ==========================================
    // 重载 operator delete
    // ==========================================
    static void operator delete(void* p, size_t size);

private:
    // ==========================================
    // next 指针：用于构建自由链表
    // 【注意】每个对象都带这个指针，有内存开销！
    // ==========================================
    Screen* next;
    
    // 数据成员
    int i;
    
    // ==========================================
    // 静态成员：自由链表的头指针
    // ==========================================
    static Screen* freeStore;
    
    // 每次分配的对象数量
    static const int screenChunk;
};

// ==========================================
// 静态成员初始化
// ==========================================
Screen* Screen::freeStore = nullptr;  // 链表初始为空
const int Screen::screenChunk = 24;   // 每次分配 24 个对象

// ==========================================
// operator new 的实现
// ==========================================
void* Screen::operator new(size_t size) {
    Screen* p = nullptr;
    
    // ==========================================
    // 情况1：自由链表不为空
    // ==========================================
    if (!freeStore) {
        // 直接从链表头取一个对象
        p = freeStore;
        freeStore = freeStore->next;  // 链表头后移
        return p;
    }
    
    // ==========================================
    // 情况2：自由链表为空，需要申请一大块内存
    // ==========================================
    // 计算需要分配的总大小
    size_t chunk = screenChunk * size;
    
    // 分配一大块 char 数组（字节级别分配）
    // 注意：用 char 是为了避免调用构造函数
    char* pChunk = new char[chunk];
    
    // 将这块内存的起始位置转为 Screen 指针
    p = reinterpret_cast<Screen*>(pChunk);
    
    // ==========================================
    // 将这一大块分割成多个对象，串成链表
    // ==========================================
    for (int i = 0; i < screenChunk - 1; ++i) {
        // 当前对象的 next 指向下一个对象
        p[i].next = reinterpret_cast<Screen*>(&p[i + 1]);
    }
    
    // 最后一个对象的 next 设为 nullptr
    p[screenChunk - 1].next = nullptr;
    
    // 将 freeStore 设为这一块的第二个对象
    freeStore = &p[1];
    
    // 返回第一个对象
    return p;
}

// ==========================================
// operator delete 的实现
// ==========================================
void Screen::operator delete(void* p, size_t size) {
    if (p == nullptr) {
        return;
    }
    
    // 将对象插回 free list 的前端
    // 为什么是前端？因为快，O(1)
    Screen* carcass = static_cast<Screen*>(p);
    carcass->next = freeStore;
    freeStore = carcass;
}

// ==========================================
// 测试代码
// ==========================================
int main() {
    cout << "====== Screen 类测试 ======" << endl;
    cout << "sizeof(Screen) = " << sizeof(Screen) << endl;
    
    const int N = 100;
    Screen* p[N];  // 指针数组
    
    // ==========================================
    // 分配 100 个对象
    // ==========================================
    for (int i = 0; i < N; ++i) {
        p[i] = new Screen(i);
    }
    
    // ==========================================
    // 输出前 10 个对象的地址，观察间隔
    // ==========================================
    cout << "\n前 10 个对象的地址：" << endl;
    for (int i = 0; i < 10; ++i) {
        cout << p[i] << endl;
    }
    
    // ==========================================
    // 测试对象是否正常使用
    // ==========================================
    p[1]->get();  // 可以正常调用成员函数
    
    // ==========================================
    // 释放所有对象
    // ==========================================
    for (int i = 0; i < N; ++i) {
        delete p[i];
    }
    
    cout << "\n====== 测试完成 ======" << endl;
    return 0;
}
```

### 3.2 代码分析

#### 3.2.1 实现思路说明

**Screen 类的设计思路：**
1. **每个对象带一个 next 指针**
2. **预分配一大块内存**，分割成多个对象
3. **用 next 指针将对象串成链表**
4. **分配时从链表头取**，**释放时插回链表头**

#### 3.2.2 核心逻辑讲解

**关键代码段 1：链表为空时申请大块**
```cpp
if (!freeStore) {
    // 分配 screenChunk 个对象大小的内存
    size_t chunk = screenChunk * size;
    char* pChunk = new char[chunk];
    p = reinterpret_cast<Screen*>(pChunk);
```

**为什么用 char 数组？**
- 避免调用 Screen 的构造函数
- 我们只需要内存，不需要对象构造
- 稍后会用 placement new（如果需要的话）

**关键代码段 2：将大块串成链表**
```cpp
for (int i = 0; i < screenChunk - 1; ++i) {
    p[i].next = reinterpret_cast<Screen*>(&p[i + 1]);
}
p[screenChunk - 1].next = nullptr;
freeStore = &p[1];  // 从第 2 个开始
return p;           // 返回第 1 个
```

**关键代码段 3：释放时插回链表头**
```cpp
carcass->next = freeStore;  // 新对象的 next 指向旧链表头
freeStore = carcass;        // 链表头更新为新对象
```

**为什么插在前端？**
- O(1) 时间复杂度，最快
- 不需要遍历链表

#### 3.2.3 初学者引导

**为什么需要 next 指针？**
- 因为要把空闲的对象链接起来
- 方便快速找到下一个可用对象

**这个版本的缺点是什么？**
- 每个对象都带着 next 指针
- 即使对象在使用中，next 指针也占内存
- 这就是下一个版本要解决的问题！

### 3.3 输出示例

```
sizeof(Screen) = 8  // int(4) + pointer(4)，无虚函数

前 10 个对象的地址：
0x3e48f0
0x3e48f8  ← 间隔 8
0x3e4900  ← 间隔 8
0x3e4908
0x3e4910
0x3e4918
0x3e4920
0x3e4928
0x3e4930
0x3e4938
```

**观察：**
- 间隔 8，没有 cookie！
- 比全局 operator new 的 16 更省空间！

---

## 四、 版本二：Airplane 类（Union 的巧妙使用）

### 4.1 问题分析

Screen 类的问题：
```cpp
class Screen {
    Screen* next;  // 即使对象在使用中，这个指针也占空间！
    int i;
};
```

**内存浪费！**
- 对象在使用中时，next 指针没用
- 对象在 free list 上时，i 数据没用

**解决方案：使用 union！**

### 4.2 Union 的巧妙之处

```cpp
union {
    AirplaneRep rep;  // 对象使用中时：用这个
    Airplane* next;   // 对象在 free list 上时：用这个
};
```

**union 的特点：**
- 所有成员共用同一块内存
- 同一时间只有一个成员有效
- 大小等于最大成员的大小

**巧妙之处：**
- 节省内存！不会同时有两个指针
- 复用同一块内存！

### 4.3 完整代码实现

```cpp
#include <cstddef>
#include <iostream>
using namespace std;

// ==========================================
// Airplane 类 - Per-class allocator 版本二
// ==========================================
class Airplane {
private:
    // ==========================================
    // 飞机数据结构体
    // ==========================================
    struct AirplaneRep {
        unsigned long miles;  // 里程数
        char type;             // 机型
    };
    
private:
    // ==========================================
    // 【核心技巧】union 的使用！
    // ==========================================
    union {
        // 情况1：对象在使用中，用 rep
        AirplaneRep rep;
        
        // 情况2：对象在 free list 上，用 next
        Airplane* next;
    };
    
public:
    // 构造函数
    Airplane() {}
    
    // 成员函数
    unsigned long getMiles() { return rep.miles; }
    char getType() { return rep.type; }
    void set(unsigned long m, char t) {
        rep.miles = m;
        rep.type = t;
    }
    
    // ==========================================
    // 重载 operator new
    // ==========================================
    static void* operator new(size_t size);
    
    // ==========================================
    // 重载 operator delete
    // ==========================================
    static void operator delete(void* p, size_t size);
    
private:
    // 自由链表头指针
    static Airplane* headOfFreeList;
    
    // 每次分配的对象数量
    static const int BLOCK_SIZE;
};

// ==========================================
// 静态成员初始化
// ==========================================
Airplane* Airplane::headOfFreeList = nullptr;
const int Airplane::BLOCK_SIZE = 512;

// ==========================================
// operator new 的实现
// ==========================================
void* Airplane::operator new(size_t size) {
    // ==========================================
    // 继承时的特殊处理！（重要）
    // ==========================================
    // 如果大小不是 sizeof(Airplane)，说明是子类对象
    // 那么就转交给全局 operator new 处理
    if (size != sizeof(Airplane)) {
        return ::operator new(size);  // 调用全局版本
    }
    
    Airplane* p = nullptr;
    
    // ==========================================
    // 情况1：自由链表不为空
    // ==========================================
    if (headOfFreeList) {
        p = headOfFreeList;
        headOfFreeList = headOfFreeList->next;
    }
    // ==========================================
    // 情况2：自由链表为空，申请一大块
    // ==========================================
    else {
        // 分配一大块内存
        Airplane* newBlock = static_cast<Airplane*>(
            ::operator new(BLOCK_SIZE * sizeof(Airplane))
        );
        
        // ==========================================
        // 将这一大块串成链表
        // ==========================================
        // 跳过第 0 个，因为我们等下会返回它
        for (int i = 1; i < BLOCK_SIZE - 1; ++i) {
            newBlock[i].next = &newBlock[i + 1];
        }
        
        // 最后一个的 next 设为 nullptr
        newBlock[BLOCK_SIZE - 1].next = nullptr;
        
        // 从第 1 个开始作为 free list
        headOfFreeList = &newBlock[1];
        
        // 返回第 0 个
        p = newBlock;
    }
    
    return p;
}

// ==========================================
// operator delete 的实现
// ==========================================
void Airplane::operator delete(void* deadObject, size_t size) {
    // 空指针检查
    if (deadObject == nullptr) {
        return;
    }
    
    // ==========================================
    // 继承时的特殊处理
    // ==========================================
    if (size != sizeof(Airplane)) {
        ::operator delete(deadObject);  // 转交全局版本
        return;
    }
    
    // ==========================================
    // 正常情况：插回 free list 前端
    // ==========================================
    Airplane* carcass = static_cast<Airplane*>(deadObject);
    carcass->next = headOfFreeList;
    headOfFreeList = carcass;
}

// ==========================================
// 测试代码
// ==========================================
int main() {
    cout << "====== Airplane 类测试 ======" << endl;
    cout << "sizeof(Airplane) = " << sizeof(Airplane) << endl;
    
    const int N = 100;
    Airplane* p[N];
    
    // ==========================================
    // 分配 100 个对象
    // ==========================================
    for (int i = 0; i < N; ++i) {
        p[i] = new Airplane();
    }
    
    // ==========================================
    // 测试对象使用
    // ==========================================
    p[1]->set(1000, 'A');
    p[5]->set(2000, 'B');
    p[9]->set(500000, 'C');
    
    cout << "\np[1]: miles=" << p[1]->getMiles() 
         << ", type=" << p[1]->getType() << endl;
    cout << "p[5]: miles=" << p[5]->getMiles() 
         << ", type=" << p[5]->getType() << endl;
    cout << "p[9]: miles=" << p[9]->getMiles() 
         << ", type=" << p[9]->getType() << endl;
    
    // ==========================================
    // 输出前 10 个对象的地址
    // ==========================================
    cout << "\n前 10 个对象的地址：" << endl;
    for (int i = 0; i < 10; ++i) {
        cout << p[i] << endl;
    }
    
    // ==========================================
    // 释放所有对象
    // ==========================================
    for (int i = 0; i < N; ++i) {
        delete p[i];
    }
    
    cout << "\n====== 测试完成 ======" << endl;
    return 0;
}
```

### 4.4 代码分析

#### 4.4.1 实现思路说明

**核心改进：**
- 使用 **union** 让 **rep** 和 **next** 共用内存
- 继承时检查大小，如果不对就转交给全局 operator new

#### 4.4.2 核心逻辑讲解

**关键代码段 1：继承时的处理**
```cpp
if (size != sizeof(Airplane)) {
    return ::operator new(size);  // 交给全局版本
}
```

**为什么需要这个检查？**
```cpp
class JetPlane : public Airplane {  // 继承
    // JetPlane 可能比 Airplane 大
};

JetPlane* p = new JetPlane;  // 会调用父类的 operator new
// 但 size 是 sizeof(JetPlane)，不是 sizeof(Airplane)
// 所以需要转交给全局版本！
```

**关键代码段 2：union 的使用**
```cpp
union {
    AirplaneRep rep;  // 使用中
    Airplane* next;   // 空闲
};
```

**内存布局：**
```
对象使用中时：
┌─────────────────┐
│  AirplaneRep    │ ← 使用 union 的这一半
│  miles          │
│  type           │
└─────────────────┘

对象在 free list 上时：
┌─────────────────┐
│  Airplane*      │ ← 使用 union 的这一半
│  next           │
└─────────────────┘
```

**同一块内存，两种用法！**

#### 4.4.3 初学者引导

**union 为什么能节省空间？**
- 因为两个成员不会同时使用
- 所以它们可以共用同一块内存
- 大小就是最大成员的大小（在这里应该是指针的大小）

**为什么继承时需要转交给全局 operator new？**
- 子类对象可能比父类大
- 我们的 free list 只有父类大小的对象
- 所以不能用，必须用全局版本

---

## 五、 内存布局与间隔对比

### 5.1 对比示意图

```
┌───────────────────────────────────────┐
│  使用 per-class allocator（左）        │
├───────────────────────────────────────┤
│ 0x3e4ce0  [Screen 对象]  间隔 8      │
│ 0x3e4ce8  [Screen 对象]  间隔 8      │
│ 0x3e4cf0  [Screen 对象]  间隔 8      │
│ ...                                    │
│ （无 cookie，更紧凑）                  │
└───────────────────────────────────────┘

┌───────────────────────────────────────┐
│  使用全局 operator new（右）          │
├───────────────────────────────────────┤
│ 0x3e4900  [Screen 对象]  间隔 16     │
│ 0x3e4910  [Screen 对象]  间隔 16     │
│ 0x3e4920  [Screen 对象]  间隔 16     │
│ ...                                    │
│ （有 cookie，开销较大）                │
└───────────────────────────────────────┘
```

### 5.2 为什么有这个区别？

**全局 operator new：**
- 会给每个块加上 cookie
- Cookie 记录块的大小等信息
- 所以对象之间间隔更大

**Per-class allocator：**
- 自己管理内存
- 不需要 cookie
- 更紧凑，更省空间

---

## 六、 两种实现方式的对比

| 特性 | Screen（版本1） | Airplane（版本2） |
|------|---------------|-------------------|
| next 指针位置 | 类成员，每个对象都有 | 在 union 中，与 rep 共用内存 |
| 内存开销 | 较大（next 一直占空间） | 较小（union 共用内存） |
| 继承处理 | 没有考虑 | 有大小检查，交给全局版本 |
| 设计简单度 | 简单 | 较巧妙 |
| 适用场景 | 不需要继承的场景 | 需要考虑继承的场景 |

---

## 七、 C++ 新标准的更新与改进

### 7.1 C++11 的改进

#### 7.1.1 std::allocator 的改进

C++11 提供了更强大的 allocator 接口，也可以自定义 allocator。

```cpp
#include <memory>

// C++11 的 allocator 使用
std::allocator<Airplane> alloc;
Airplane* p = alloc.allocate(100);  // 分配 100 个对象的空间

// 使用 allocate 后，还需要用 construct 构造对象
for (int i = 0; i < 100; ++i) {
    alloc.construct(&p[i]);  // placement new
}

// 析构
for (int i = 0; i < 100; ++i) {
    alloc.destroy(&p[i]);
}

// 释放
alloc.deallocate(p, 100);
```

#### 7.1.2 std::aligned_storage

C++11 提供的对齐内存分配工具，比 union 更类型安全：

```cpp
#include <type_traits>

struct Airplane {
    // 使用 C++11 的 aligned_storage
    std::aligned_storage<sizeof(AirplaneRep), alignof(AirplaneRep)>::type storage;
};
```

### 7.2 C++17 的 pmr（Polymorphic Memory Resources）

C++17 引入的多态内存资源，提供了更灵活的内存管理：

```cpp
#include <memory_resource>
#include <vector>

// 使用 monotonic_buffer_resource
std::pmr::monotonic_buffer_resource pool;
std::pmr::vector<Airplane> vec(&pool);
```

---

## 八、 面试高频问题与解答

### 8.1 基础问题

#### Q1: Per-class allocator 是什么？有什么好处？

**答案：**

**定义：**
Per-class allocator 是在类内部重载 `operator new/operator delete`，为该类自定义内存分配策略。

**好处：**
1. **减少内存碎片**：一次分配大块，反复使用
2. **没有 cookie**：更省空间，对象间隔更小
3. **效率更高**：减少 malloc/free 调用次数
4. **内存复用**：对象释放后可以再次分配

#### Q2: Screen 类的 per-class allocator 实现思路是？

**答案：**

**思路：**
1. **预分配一大块内存**（24个对象的大小）
2. **用 next 指针将这些对象串成链表**
3. **分配时从链表头取一个对象**
4. **释放时将对象插回链表头**

**核心代码：**
```cpp
Screen* p = freeStore;
freeStore = freeStore->next;  // 分配
carcass->next = freeStore;
freeStore = carcass;  // 释放
```

#### Q3: Airplane 类为什么要用 union？相比 Screen 有什么改进？

**答案：**

**Screen 类的问题：**
- `next` 指针作为成员一直存在
- 即使对象在使用中，next 指针也占空间

**Airplane 类的改进：**
```cpp
union {
    AirplaneRep rep;  // 使用中时
    Airplane* next;   // 在 free list 上时
};
```

**union 的巧妙：**
- 两个成员共用同一块内存
- 不会同时使用，所以不冲突
- 节省内存！

### 8.2 进阶问题

#### Q4: Airplane 类的 operator new 为什么要检查 size？

**答案：**

**为了处理继承！**

```cpp
if (size != sizeof(Airplane)) {
    return ::operator new(size);
}
```

**原因：**
```cpp
class JetPlane : public Airplane {  // 继承
    // JetPlane 可能比 Airplane 大！
    // 假设 Airplane 8 字节，JetPlane 12 字节
};

JetPlane* p = new JetPlane;  // 会调用父类的 operator new
// 此时传入的 size 是 12，不是 8！
// 所以不能用我们的 free list，必须用全局版本！
```

**这是一个非常重要的边界处理！**

#### Q5: Per-class allocator 会不会造成内存泄漏？

**答案：**

**从截图中可以看到侯捷老师说的话：**
```
沒有對應的 delete，但這不算是 memory leak！
```

**为什么？**
- 因为程序结束后，操作系统会回收所有内存
- 我们只是让内存在 free list 上，程序结束后系统会回收
- 不是真正的内存泄漏（真正的泄漏是不再使用但未释放的内存）

**但更好的做法：**
- 可以写一个静态清理函数
- 在程序结束时释放所有大块

```cpp
static void cleanUp() {
    // 释放所有预分配的大块内存
}
```

#### Q6: 除了 free list，还有其他的内存池实现方式吗？

**答案：**

**其他方式：**
1. **Bitmap（位图）**：用位记录哪个位置空/已用
2. **固定大小分配器**：只分配固定大小的块
3. **Segregated Storage（隔离存储）**：多个大小的 free list

**C++ 标准库中的实现：**
- GCC 的 `__pool_alloc`
- Boost.Pool

### 8.3 实际问题

#### Q7: Per-class allocator 在实际项目中用得多吗？

**答案：**

**在 C++ 标准库和高性能项目中常用：**

**例子：**
1. **STL 容器都有 allocator 模板参数**
2. **游戏开发**：大量相同对象的分配/释放
3. **高频交易系统**：性能要求高
4. **嵌入式系统**：内存有限，需要紧凑管理

**但在一般项目中：**
- 不常用，增加复杂度
- 除非性能瓶颈确实在内存分配上

#### Q8: 写 per-class allocator 时需要注意什么？

**答案：**

**注意事项：**
1. **继承处理**：检查 size，转交全局版本
2. **空指针检查**：operator delete 检查 nullptr
3. **线程安全**（如果需要多线程）：加锁
4. **内存对齐**：保证正确对齐
5. **异常安全**：分配失败时的处理

---

## 九、 完整可运行代码（两个版本对比）

```cpp
#include <cstddef>
#include <iostream>
using namespace std;

// ==========================================
// 版本一：Screen 类（next 指针作为成员）
// ==========================================
class Screen {
public:
    Screen(int x) : i(x) {}
    int get() const { return i; }
    
    static void* operator new(size_t size);
    static void operator delete(void* p, size_t size);

private:
    Screen* next;
    int i;
    static Screen* freeStore;
    static const int screenChunk;
};

Screen* Screen::freeStore = nullptr;
const int Screen::screenChunk = 24;

void* Screen::operator new(size_t size) {
    Screen* p = nullptr;
    if (!freeStore) {
        size_t chunk = screenChunk * size;
        char* pChunk = new char[chunk];
        p = reinterpret_cast<Screen*>(pChunk);
        
        for (int i = 0; i < screenChunk - 1; ++i) {
            p[i].next = reinterpret_cast<Screen*>(&p[i + 1]);
        }
        p[screenChunk - 1].next = nullptr;
        freeStore = &p[1];
    } else {
        p = freeStore;
        freeStore = freeStore->next;
    }
    return p;
}

void Screen::operator delete(void* p, size_t size) {
    if (!p) return;
    Screen* carcass = static_cast<Screen*>(p);
    carcass->next = freeStore;
    freeStore = carcass;
}

// ==========================================
// 版本二：Airplane 类（union 的巧妙使用）
// ==========================================
class Airplane {
private:
    struct AirplaneRep {
        unsigned long miles;
        char type;
    };
    
    union {
        AirplaneRep rep;
        Airplane* next;
    };
    
public:
    Airplane() {}
    unsigned long getMiles() { return rep.miles; }
    char getType() { return rep.type; }
    void set(unsigned long m, char t) {
        rep.miles = m;
        rep.type = t;
    }
    
    static void* operator new(size_t size);
    static void operator delete(void* p, size_t size);
    
private:
    static Airplane* headOfFreeList;
    static const int BLOCK_SIZE;
};

Airplane* Airplane::headOfFreeList = nullptr;
const int Airplane::BLOCK_SIZE = 512;

void* Airplane::operator new(size_t size) {
    if (size != sizeof(Airplane)) {
        return ::operator new(size);
    }
    
    Airplane* p = nullptr;
    if (headOfFreeList) {
        p = headOfFreeList;
        headOfFreeList = headOfFreeList->next;
    } else {
        Airplane* newBlock = static_cast<Airplane*>(
            ::operator new(BLOCK_SIZE * sizeof(Airplane))
        );
        for (int i = 1; i < BLOCK_SIZE - 1; ++i) {
            newBlock[i].next = &newBlock[i + 1];
        }
        newBlock[BLOCK_SIZE - 1].next = nullptr;
        headOfFreeList = &newBlock[1];
        p = newBlock;
    }
    return p;
}

void Airplane::operator delete(void* deadObject, size_t size) {
    if (deadObject == nullptr) {
        return;
    }
    if (size != sizeof(Airplane)) {
        ::operator delete(deadObject);
        return;
    }
    
    Airplane* carcass = static_cast<Airplane*>(deadObject);
    carcass->next = headOfFreeList;
    headOfFreeList = carcass;
}

// ==========================================
// 测试代码
// ==========================================
int main() {
    cout << "========================================" << endl;
    cout << "  Screen 类 vs Airplane 类 对比测试" << endl;
    cout << "========================================" << endl;
    
    // ==========================================
    // 测试 Screen 类
    // ==========================================
    cout << "\n====== Screen 类 ======" << endl;
    cout << "sizeof(Screen) = " << sizeof(Screen) << endl;
    
    const int N = 10;
    Screen* s[N];
    for (int i = 0; i < N; ++i) {
        s[i] = new Screen(i);
    }
    
    cout << "\nScreen 地址：" << endl;
    for (int i = 0; i < N; ++i) {
        cout << s[i] << endl;
    }
    
    for (int i = 0; i < N; ++i) {
        delete s[i];
    }
    
    // ==========================================
    // 测试 Airplane 类
    // ==========================================
    cout << "\n====== Airplane 类 ======" << endl;
    cout << "sizeof(Airplane) = " << sizeof(Airplane) << endl;
    
    Airplane* a[N];
    for (int i = 0; i < N; ++i) {
        a[i] = new Airplane();
    }
    
    for (int i = 0; i < N; i += 2) {
        a[i]->set(i * 1000, 'A' + i % 26);
    }
    
    cout << "\nAirplane 地址：" << endl;
    for (int i = 0; i < N; ++i) {
        cout << a[i] << endl;
    }
    
    for (int i = 0; i < N; ++i) {
        delete a[i];
    }
    
    cout << "\n========================================" << endl;
    cout << "  测试完成！" << endl;
    cout << "========================================" << endl;
    return 0;
}
```

---

## 十、 总结与反思

### 10.1 核心知识点总结

| 知识点 | 要点 |
|--------|------|
| Per-class allocator | 类内部重载 operator new/delete |
| Free list | 空闲对象的链表，方便分配/释放 |
| Union 的巧妙 | 同一内存两种用途，节省空间 |
| 继承处理 | 检查 size，不对时转交全局版本 |
| 内存间隔对比 | 8 vs 16，自定义更紧凑 |
| 预分配大块 | 减少 malloc 调用次数 |

### 10.2 学习收获

1. **理解了自定义内存管理的重要性**
   - 为什么要自定义？
   - 什么时候需要自定义？

2. **掌握了两种 per-class allocator 的实现**
   - 简单版（Screen）
   - 巧妙版（Airplane with union）

3. **学会了 union 的实用技巧**
   - 同一块内存，两种用法
   - 节省空间的好方法

4. **理解了继承时的边界情况处理**
   - size 检查很重要
   - 不能处理时就交出去

### 10.3 后续学习建议

1. **学习 STL allocator**
   - 了解 std::allocator 的接口
   - 学习如何写自定义 allocator

2. **学习 Boost.Pool 库**
   - 现成的内存池实现
   - 学习工业级的代码

3. **学习内存对齐**
   - 为什么要对齐？
   - 如何正确对齐？

4. **学习线程安全版本**
   - 多线程环境下的内存分配
   - 如何加锁？

---

通过本课程，我们深入理解了 per-class allocator 的实现原理，掌握了两种不同的实现方式（Screen 类和 Airplane 类），理解了 union 在节省内存方面的巧妙应用，也学会了继承时的边界处理。这些知识对于理解 C++ 内存管理和写高性能代码非常重要！
