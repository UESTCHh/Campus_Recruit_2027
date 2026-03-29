# CoreLoop C++ 进阶核心复盘 - 智能指针 (Smart Pointers)

> **打卡日期**：2026-03-29 (Day 4 课外加餐)
> **核心主题**：RAII 机制、`unique_ptr` 独占所有权、`shared_ptr` 引用计数机制、裸指针的致命陷阱、手撕智能指针。
> **资料网址**: https://www.bilibili.com/video/BV1rdAVzAEqf/?spm_id_from=333.1007.tianma.1-2-2.click&vd_source=e01209adaee5cf8495fa73fbfde26dd1

---

## 一、 智能指针的灵魂：RAII 思想

智能指针根本不是什么神奇的“指针”，它的本质是一个**分配在栈上的局部对象**。
它利用了栈对象在离开作用域时**必定会自动调用析构函数**的 C++ 铁律（也就是我们在 Day 2 刚复习过的概念）。我们将裸指针（堆内存）包装进这个局部对象里，在析构函数里写上 `delete`。这样，无论程序是正常结束，还是中间抛出异常退出，内存都绝对会被安全释放。这种机制叫做 **RAII (Resource Acquisition Is Initialization)**。

---

## 二、 视频作业深度剖析 (极其经典的大厂面试坑题)

### 🎯 作业 1：同一裸指针初始化多个智能指针的灾难
**第一段代码：**
```cpp
int* p = new int(3);
unique_ptr<int> up1(p);
unique_ptr<int> up2(p);
```
能编译通过吗？ 能！ 编译器非常傻，它不会阻止你这么写。

有什么 Bug？ 极其严重的 Double Free（重复释放）导致程序崩溃！

* unique_ptr 的核心语义是“独占”。up1 认为自己是 p 的唯一主人。

* up2 被构建时，它也认为自己是 p 的唯一主人。

当这两个局部变量离开作用域时，up2 会先调用析构函数 delete p;，接着 up1 也会调用析构函数再次 delete p;。同一块堆内存被杀死了两次，程序直接 Core Dump。

**第二段代码：**

```cpp
int* p = new int(3);
shared_ptr<int> sp1(p);
shared_ptr<int> sp2(p);
```
能编译通过吗？ 能！

有什么 Bug？ 依然是 Double Free！ 很多人以为用了 shared_ptr 就会自动共享引用计数，这是大错特错！

当你用同一个裸指针分别去构造两个 shared_ptr 时，它们是相互完全不知道对方存在的。

* sp1 会在底层给自己建一个“控制块（Control Block）”，里面写着：引用计数 = 1。

* sp2 会在底层给自己建另一个全新的“控制块”，里面也写着：引用计数 = 1。

结局同样悲惨：它们销毁时都会把计数从 1 减到 0，然后各自去 delete p 一次。

#### 💡 面试核心铁律：
绝对、绝对、绝对不要用同一个裸指针去初始化多个智能指针！
* 最佳实践：永远使用 std::make_unique<int>(3) 和 std::make_shared<int>(3) 一步到位，彻底消灭裸指针。

### 🎯 作业 2：强引用计数的生命周期演变
这是一道考察 shared_ptr 底层引用计数（Use Count）变化的满分追踪题。

```cpp
void f(shared_ptr<int> sp) { // 👈 注意：这里是按值传递 (Pass by value)
    shared_ptr<int> sp2 = sp;
    cout << *sp2 << endl;
}

int main() {
    shared_ptr<int> sp = std::make_shared<int>(3); // 状态 A
    f(sp); // 状态 B & C
    return 0; // 状态 D
}
```
强引用计数的完整变化过程：

* 状态 A (main 中创建)： make_shared 分配内存，控制块中强引用计数 = 1。

* 状态 B (进入函数 f 的瞬间)： 注意函数 f 的参数是按值传递的！这会触发拷贝构造函数。形参 sp 是实参 sp 的一份拷贝，强引用计数 +1，此时计数 = 2。

* 状态 C (执行 sp2 = sp)： 内部又发生了一次拷贝构造。局部变量 sp2 也指向了同一块内存，强引用计数再次 +1，此时计数 = 3。

* 离开函数 f 的瞬间： 局部变量 sp2 和形参 sp 相继离开作用域，调用两次析构函数。计数减 2，此时强引用计数恢复 = 1。

* 状态 D (离开 main)： 最后的 sp 离开作用域，强引用计数减 1 变成 0。触发终极动作：安全释放底层的 int(3) 内存！无泄漏！

### 🎯 作业 3：手撕智能指针核心骨架 (大厂终极必杀)
如果面试官让你现场写一个 shared_ptr，核心就是搞懂那个单独存放在堆上的引用计数器（Control Block）。

```cpp

template <typename T>
class MySharedPtr {
private:
    T* ptr;         // 指向实际数据的指针
    int* ref_count; // 💡 灵魂所在：指向引用计数的指针（必须在堆上，才能让多个对象共享！）

public:
    // 1. 构造函数
    MySharedPtr(T* p = nullptr) {
        if (p) {
            ptr = p;
            ref_count = new int(1); // 第一次创建，计数初始化为 1
        } else {
            ptr = nullptr;
            ref_count = new int(0);
        }
    }

    // 2. 拷贝构造函数 (核心逻辑：计数 +1)
    MySharedPtr(const MySharedPtr& other) {
        ptr = other.ptr;
        ref_count = other.ref_count; // 指向同一个计数器
        (*ref_count)++;              // 大家共享的计数器加 1
        cout << "[Debug] Copy Construct! Current Count: " << *ref_count << endl;
    }

    // 3. 析构函数 (核心逻辑：计数 -1，为 0 时销毁)
    ~MySharedPtr() {
        (*ref_count)--;
        cout << "[Debug] Destruct! Current Count: " << *ref_count << endl;
        if (*ref_count == 0) {
            delete ptr;
            delete ref_count; // 别忘了把计数器自己的内存也释放掉！
            cout << "[Debug] Memory actually freed!" << endl;
        }
    }
    
    // (注：为保持简洁，此处省略了拷贝赋值 operator= 和移动语义的实现)
};
```

## 三、 智能指针全景解析 (大厂八股终极总结)

### 1. 核心定义与作用
* **定义：** 智能指针（引入 `<memory>` 头文件）本质上不是指针，而是**重载了 `->` 和 `*` 操作符的类模板**。它在栈上创建，通过 RAII（资源获取即初始化）机制，接管了堆内存的生杀大权。
* **作用：** 彻底终结 C/C++ 程序员的噩梦——内存泄漏（忘记 `delete`）、悬空指针（野指针）、重复释放（Double Free）。

### 2. 三大核心剑客与使用场景

#### 🗡️ 剑客一：`std::unique_ptr` (独占式指针)
* **核心机制：** 绝对的霸道总裁。同一时刻，只能有一个 `unique_ptr` 指向某块内存。它**禁用了拷贝构造和拷贝赋值**，只允许**移动（`std::move`）**。
* **性能开销：** **零开销！** 它的内存大小和运行速度与普通的裸指针一模一样，是大厂 C++ 代码中的首选。
* **使用场景：** * 函数内部产生的局部动态对象（函数结束自动销毁）。
  * 类的 Pimpl 惯用法（隐藏实现细节，正如咱们白天复习的“委托”关系）。
  * 明确所有权转移的场景（比如把对象塞进一个容器里，原指针就不再管了）。

#### 🛡️ 剑客二：`std::shared_ptr` (共享式指针)
* **核心机制：** 民主投票制。底层维护着一个分配在堆上的**控制块（包含强引用计数和弱引用计数）**。只要有一个人还在用，内存就不销毁；当最后一个人离开（强引用计数清零），内存释放。
* **性能开销：** 有代价。它的大小是裸指针的两倍（一个指针指数据，一个指针指控制块）。且多线程下修改引用计数需要加锁（原子操作），有轻微性能损耗。
* **使用场景：**
  * 多个模块、多个线程需要共享同一个大对象（如配置信息、全局缓存）。
  * 复杂的图状数据结构（但要极度小心循环引用）。

#### 🩹 剑客三：`std::weak_ptr` (弱引用指针 / 破局者)

* **核心机制：** 它是 `shared_ptr` 的小跟班。它**只观察，不拥有**。把 `weak_ptr` 绑定到 `shared_ptr` 上，**不会增加强引用计数**。
* **核心作用：彻底解决“循环引用”导致的神仙死锁。**
  * *灾难场景：* 对象 A 里有个 `shared_ptr` 指向 B，对象 B 里有个 `shared_ptr` 指向 A。两人互相拉扯，强引用计数永远降不到 0，导致极其隐蔽的内存泄漏。
  * *破局之法：* 把其中一个指针改成 `weak_ptr`，循环羁绊瞬间瓦解！
* **使用技巧：** `weak_ptr` 不能直接访问数据，必须先调用 `.lock()` 方法，尝试把它提升为 `shared_ptr`。如果原内存还在，就提升成功；如果原内存已经被销毁了，就返回一个空的 `shared_ptr`，极其安全。

### 3. 大厂开发避坑指南 (Precautions)
1. **绝不混用：** 千万不要把同一个裸指针交由多个智能指针管理（会引发 Double Free）。
2. **拒绝 `this` 诱惑：** 在类的内部，千万不要直接 `return shared_ptr<T>(this);`。必须让类继承 `std::enable_shared_from_this`，然后调用 `shared_from_this()`，否则会产生两个独立的控制块，同样导致 Double Free！
3. **拥抱 `make_xxx`：** 永远优先使用 `std::make_unique` 和 `std::make_shared` 来创建对象，不仅代码更简洁，而且对于 `shared_ptr` 来说，能把数据块和控制块合并成一次内存分配，性能更高，防内存碎片。

### 4. 面试高频连环炮 (必背)
* **Q1：`shared_ptr` 是线程安全的吗？**
  * **满分回答：** “控制块（引用计数）的操作是线程安全的（原子操作）。但是，**智能指针指向的那个数据对象本身，绝对不是线程安全的！** 如果多个线程同时读写那个对象的内容，必须自己加锁（如 `std::mutex`）。”
* **Q2：怎么实现一个 `unique_ptr`？**
  * **满分回答：** “重点在于两步：1. 把对象包装进类中并在析构函数中 `delete`。 2. 把拷贝构造函数和赋值运算符显式设为 `= delete`，然后提供移动构造函数（Move Constructor）来转移内部裸指针的所有权，并将原指针置空。”
* **Q3：为什么需要 `weak_ptr`？**
  * **满分回答：** “为了打破 `shared_ptr` 的循环引用问题。比如在观察者模式（Observer）中，Subject 拥有 Observer 的 `shared_ptr`，而 Observer 为了获取最新状态，如果也拿着 Subject 的 `shared_ptr` 就会死锁，此时 Observer 持有 Subject 的 `weak_ptr` 是最完美的解法。”
  
---