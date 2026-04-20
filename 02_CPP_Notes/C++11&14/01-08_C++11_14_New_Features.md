# 侯捷 C++ 新标准 C++11&14 学习笔记（1-8集）

> **学习内容**：C++11/14 新特性，包括 Variadic Templates、nullptr、auto、Uniform Initialization、Initializer Lists、explicit、Range-based for
> **学习资源**：侯捷老师 C++ 系列视频课程
> **课程链接**：https://www.bilibili.com/video/BV1wBh5zkE2Y?spm_id_from=333.788.videopod.sections&vd_source=e01209adaee5cf8495fa73fbfde26dd1
> **学习目标**：掌握 C++11/14 的核心新特性，理解其设计意图和使用方法

---

## 一、 C++ 标准的演化

### 1. 标准演化历程
- **C++ 98 (1.0)**：第一个正式标准
- **C++ 03 (TR1, Technical Report 1)**：对 C++ 98 的小幅修正
- **C++ 11 (2.0)**：重大更新，引入大量新特性
- **C++ 14**：对 C++ 11 的补充和完善

### 2. 头文件的变化
- **C++ 标准库的头文件**：不带后缀名 (.h)，例如 `#include <vector>`
- **新式 C 头文件**：不带后缀名 .h，例如 `#include <cstdio>`
- **旧式 C 头文件**：带有后缀名 .h，仍可用，例如 `#include <stdio.h>`

### 3. C++ 2.0 的重磅特性

#### 语言层面
- **Variadic Templates**：可变参数模板
- **Move Semantics**：移动语义
- **auto**：自动类型推导
- **Range-based for loop**：基于范围的 for 循环
- **Initializer List**：初始化列表
- **Lambdas**：lambda 表达式

#### 标准库层面
- **type_traits**：类型特性
- **Unordered 容器**：无序容器
- **forward_list**：前向链表
- **array**：固定大小数组
- **tuple**：元组
- **Concurrency**：并发
- **RegEx**：正则表达式

---

## 二、 第一颗震撼弹：Variadic Templates

### 1. 基本概念
- **Variadic Templates**：数量不定的模板参数
- **参数包 (pack)**：
  - 用于 template parameters：template parameters pack
  - 用于 function parameter types：function parameter types pack
  - 用于 function parameters：function parameters pack

### 2. 基本语法

```cpp
// 基础版本（递归终止条件）
void printX() {
}

// 可变参数模板版本
template <typename T, typename... Types>
void printX(const T &firstArg, const Types&... args) {
    cout << firstArg << endl;  // 打印第一个参数
    printX(args...);            // 递归调用剩余参数
}
```

### 3. 工作原理
- **递归调用**：每次处理一个参数，然后递归处理剩余参数
- **参数包展开**：使用 `args...` 展开参数包
- **递归终止**：提供一个无参数的基础版本作为递归终止条件

### 4. 与 C 风格可变参数的对比
- **C 风格**：使用 `va_list`、`va_start`、`va_end` 等宏，类型不安全
- **C++ 风格**：类型安全，编译时检查，代码更清晰

### 5. 应用场景

#### 5.1 递归函数调用
```cpp
// 示例：打印任意数量的参数
printX(7.5, "hello", bitset<16>(377), 42);
// 输出：
// 7.5
// hello
// 0000000101111001
// 42
```

#### 5.2 哈希函数实现
```cpp
// 哈希组合函数
template <typename T>
inline void hash_combine(size_t &seed, const T &val) {
    seed ^= std::hash<T>()(val) + 0x9e3779b9
           + (seed<<6) + (seed>>2);
}

// 辅助函数（处理单个参数）
template <typename T>
inline void hash_val(size_t &seed, const T &val) {
    hash_combine(seed, val);
}

// 可变参数版本
template <typename T, typename... Types>
inline void hash_val(size_t &seed, const T &val, const Types&... args) {
    hash_combine(seed, val);
    hash_val(seed, args...);
}

// 主函数
template <typename... Types>
inline size_t hash_val(const Types&... args) {
    size_t seed = 0;
    hash_val(seed, args...);
    return seed;
}

// 使用示例
class CustomerHash {
public:
    std::size_t operator()(const Customer& c) const {
        return hash_val(c.fname, c.lname, c.no);
    }
};
```

#### 5.3 递归继承
```cpp
// tuple 的实现原理
template<typename... Values> class tuple;
template<> class tuple<> {};

template<typename Head, typename... Tail>
class tuple<Head, Tail...> : private tuple<Tail...> {
    typedef tuple<Tail...> inherited;
public:
    tuple() {}
    tuple(Head v, Tail... vtail) 
        : m_head(v), inherited(vtail...) {}
    
    typename Head::type head() { return m_head; }
    inherited& tail() { return *this; }
    
protected:
    Head m_head;
};

// 使用示例
tuple<int, float, string> t(41, 6.3, "nico");
t.head();           // 获得 41
(t.tail()).head();  // 获得 6.3
&(t.tail())        // 指向包含 6.3 和 "nico" 的 tuple
```

### 6. 面试相关问题

#### 6.1 基础问题
- **什么是 Variadic Templates？**
  - 答：Variadic Templates 是 C++11 引入的特性，允许模板接受数量不定的参数。

- **Variadic Templates 与 C 风格可变参数的区别？**
  - 答：C++ 的 Variadic Templates 是类型安全的，编译时检查，而 C 风格可变参数使用宏，类型不安全。

- **如何实现一个接受任意数量参数的函数？**
  - 答：使用 Variadic Templates，结合递归调用和参数包展开。

#### 6.2 进阶问题
- **Variadic Templates 的递归终止条件如何设计？**
  - 答：提供一个无参数的基础版本作为递归终止条件。

- **如何在 Variadic Templates 中获取参数包的大小？**
  - 答：使用 `sizeof...(args)` 操作符获取参数包的大小。

- **Variadic Templates 在标准库中的应用？**
  - 答：tuple、function、bind 等都使用了 Variadic Templates。

---

## 三、 模板表达式中的空格

### 1. 问题背景
- 在 C++98 中，两个右尖括号 `>>` 会被解析为右移运算符，而不是模板参数的结束
- 因此需要在两个右尖括号之间添加空格，如 `vector<list<int> >`

### 2. C++11 的改进
- C++11 中，编译器可以正确解析 `vector<list<int>>` 这种形式
- 不再需要在两个右尖括号之间添加空格

### 3. 示例
```cpp
// C++98 风格（需要空格）
vector<list<int> > v1;

// C++11 风格（不需要空格）
vector<list<int>> v2;
```

### 4. 标准库中的应用
- G4.5.3 版本的标准库已经使用了新的语法：
  ```cpp
  // G4.5.3/include/c++/profile/vector
  template<typename _Alloc>
  struct hash<__profile::vector<bool, _Alloc>>
  
  // G4.5.3/include/c++/bits/stl_bvector.h
  struct hash<__GLIBCXX_STD_D::vector<bool, _Alloc>>
  ```

### 5. 面试相关问题
- **C++11 中模板表达式的空格有什么变化？**
  - 答：C++11 允许 `vector<list<int>>` 这种形式，不再需要在两个右尖括号之间添加空格。

- **为什么 C++98 中需要添加空格？**
  - 答：因为 C++98 编译器会将 `>>` 解析为右移运算符，而不是模板参数的结束。

---

## 四、 nullptr 和 std::nullptr_t

### 1. 问题背景
- 在 C++98 中，使用 `0` 或 `NULL` 表示空指针
- 但 `0` 是整数，`NULL` 通常被定义为 `0` 或 `(void*)0`
- 这会导致类型混淆，例如函数重载时的歧义

### 2. C++11 的解决方案
- 引入新关键字 `nullptr`
- `nullptr` 的类型是 `std::nullptr_t`
- `std::nullptr_t` 是一种基本数据类型

### 3. 特性
- `nullptr` 自动转换为任何指针类型
- `nullptr` 不会转换为任何整数类型
- 可以用于函数重载

### 4. 示例
```cpp
void f(int);
void f(void*);

f(0);         // 调用 f(int)
f(NULL);      // 调用 f(int)（如果 NULL 是 0），否则歧义
f(nullptr);   // 调用 f(void*)
```

### 5. 实现细节
- 在 `<cstddef>` 头文件中定义：
  ```cpp
  #if defined(__cplusplus) && __cplusplus >= 201103L
  #ifndef _GXX_NULLPTR_T
  #define _GXX_NULLPTR_T
  typedef decltype(nullptr) nullptr_t;
  #endif
  #endif
  ```

### 6. 面试相关问题
- **什么是 nullptr？**
  - 答：nullptr 是 C++11 引入的关键字，表示空指针，类型为 std::nullptr_t。

- **nullptr 与 0 和 NULL 的区别？**
  - 答：nullptr 是类型安全的，自动转换为任何指针类型，但不会转换为整数类型；而 0 是整数，NULL 通常被定义为 0 或 (void*)0，可能导致类型混淆。

- **什么时候应该使用 nullptr？**
  - 答：当需要表示空指针时，应该使用 nullptr 而不是 0 或 NULL，以避免类型混淆。

---

## 五、 自动类型推导：auto 关键字

### 1. 基本概念
- `auto` 关键字允许编译器自动推导变量的类型
- 从 C++11 开始引入

### 2. 基本用法
```cpp
auto i = 42;      // i 类型为 int
double f();
auto d = f();      // d 类型为 double
```

### 3. 适用场景
- **类型较长或复杂**：例如迭代器类型
  ```cpp
  vector<string> v;
  // C++98 风格
  vector<string>::iterator pos = v.begin();
  // C++11 风格
  auto pos = v.begin();
  ```

- **lambda 表达式**：lambda 表达式的类型是匿名的，只能用 auto 推导
  ```cpp
  auto I = [](int x) -> bool {
      // ...
  };
  ```

### 4. 标准库中的应用
- G4.5.3 版本的标准库已经使用了 auto：
  ```cpp
  #if __cplusplus >= 201103L
  // DR 685.
  inline auto
  operator-(const reverse_iterator<_IteratorL>& __x,
            const reverse_iterator<_IteratorR>& __y)
  -> decltype(__y.base() - __x.base())
  #else
  inline typename reverse_iterator<_IteratorL>::difference_type
  operator-(const reverse_iterator<_IteratorL>& __x,
            const reverse_iterator<_IteratorR>& __y)
  #endif
  ```

### 5. 面试相关问题
- **什么是 auto 关键字？**
  - 答：auto 是 C++11 引入的关键字，允许编译器自动推导变量的类型。

- **auto 的适用场景？**
  - 答：当类型较长或复杂时（如迭代器类型），或当类型是匿名的时（如 lambda 表达式）。

- **auto 与 decltype 的区别？**
  - 答：auto 用于变量声明时的类型推导，decltype 用于获取表达式的类型，不用于变量声明。

---

## 六、 统一初始化：Uniform Initialization

### 1. 问题背景
- C++98 中，初始化方式多样，容易混淆：
  - 圆括号：`int i(42);`
  - 赋值：`int i = 42;`
  - 花括号：`int a[] = {1, 2, 3};`

### 2. C++11 的解决方案
- 引入统一初始化语法，使用花括号 `{}`
- 适用于任何类型的初始化

### 3. 基本用法
```cpp
int values[] {1, 2, 3};
vector<int> v {2, 3, 5, 7, 11, 13, 17};
vector<string> cities {
    "Berlin", "New York", "London", "Braunschweig", "Cairo", "Cologne"
};
complex<double> c {4.0, 3.0};  // 等同于 c(4.0, 3.0)
```

### 4. 工作原理
- 编译器看到 `{t1, t2, ..., tn}` 时，会创建一个 `initializer_list<T>`
- 背后有一个 `array<T, n>`
- 调用函数（如构造函数）时，数组内的元素会被分解逐一传给函数

### 5. 注意事项
- 如果函数参数是 `initializer_list<T>`，调用者不能给予数值 T 参数然后以为它们会被自动转为一个 `initializer_list<T>` 传入

### 6. 面试相关问题
- **什么是统一初始化？**
  - 答：统一初始化是 C++11 引入的特性，使用花括号 `{}` 进行初始化，适用于任何类型。

- **统一初始化的优势？**
  - 答：语法统一，减少混淆；强制值初始化；防止窄化转换。

- **统一初始化与传统初始化的区别？**
  - 答：统一初始化使用花括号，传统初始化使用圆括号或赋值运算符。统一初始化更安全，能防止窄化转换。

---

## 七、 初始化列表：Initializer Lists

### 1. 基本概念
- `initializer_list<T>` 是 C++11 提供的类模板
- 用于支持值列表初始化
- 可以在任何需要处理值列表的地方使用

### 2. 特性
- **值初始化**：即使是基本数据类型的局部变量，也会被初始化为零（或 nullptr）
  ```cpp
  int i;          // i 未初始化，值不确定
  int j{};        // j 被初始化为 0
  int* p;         // p 未初始化，值不确定
  int* q{};       // q 被初始化为 nullptr
  ```

- **防止窄化初始化**：不允许精度降低或值被修改的初始化
  ```cpp
  int x1(5.3);     // OK，但 OUCH: x1 变为 5
  int x2 = 5.3;    // OK，但 OUCH: x2 变为 5
  int x3{5.0};     // ERROR: narrowing
  int x4 = {5.3};  // ERROR: narrowing
  char c1{7};      // OK: 即使 7 是 int，也不是窄化
  char c2{99999};  // ERROR: narrowing (99999 不适合 char)
  ```

### 3. 实现细节
- `initializer_list<T>` 的内部实现：
  ```cpp
  template<class _E>
  class initializer_list {
  public:
    typedef _E          value_type;
    typedef const _E&   reference;
    typedef const _E&   const_reference;
    typedef size_t      size_type;
    typedef const _E*   iterator;
    typedef const _E*   const_iterator;
  
  private:
    iterator          _M_array;
    size_type         _M_len;
  
    // The compiler can call a private constructor.
    constexpr initializer_list(const_iterator __a, size_type __l)
      : _M_array(__a), _M_len(__l) {}
  
  public:
    constexpr initializer_list() noexcept
      : _M_array(0), _M_len(0) {}
  
    // Number of elements.
    constexpr size_type size() const noexcept { return _M_len; }
  
    // First element.
    constexpr const_iterator begin() const noexcept { return _M_array; }
  
    // One past the last element.
    constexpr const_iterator end() const noexcept { return begin() + size(); }
  };
  ```

### 4. 使用场景

#### 4.1 函数参数
```cpp
void print(std::initializer_list<int> vals) {
    for (auto p = vals.begin(); p != vals.end(); ++p) {
        std::cout << *p << "\n";
    }
}

// 调用
print({12, 3, 5, 7, 11, 13, 17});
```

#### 4.2 构造函数优先级
- 当同时有特定参数个数的构造函数和 `initializer_list` 版本时，优先选择 `initializer_list` 版本
  ```cpp
  class P {
  public:
    P(int a, int b) {
        cout << "P(int, int), a=" << a << ", b=" << b << endl;
    }
    P(initializer_list<int> initlist) {
        cout << "P(initializer_list<int>), values=";
        for (auto i : initlist)
            cout << i << ' ';
        cout << endl;
    }
  };

  P p1(77, 5);      // P(int, int), a=77, b=5
  P p2{77, 5};      // P(initializer_list<int>), values=77 5
  P p3{77, 5, 42};  // P(initializer_list<int>), values=77 5 42
  P p4 = {77, 5};   // P(initializer_list<int>), values=77 5
  ```

#### 4.3 容器操作
- 所有容器都接受 `initializer_list` 用于构造或赋值
  ```cpp
  // 构造
  vector<int> v1 {2, 5, 7, 13, 69, 83, 50};
  
  // 赋值
  vector<int> v2;
  v2 = {2, 5, 7, 13, 69, 83, 50};
  
  // 插入
  v3.insert(v3.begin()+2, {0, 1, 2, 3, 4});
  ```

#### 4.4 标准库算法
- 标准库算法如 `max`、`min` 也接受 `initializer_list`
  ```cpp
  cout << max({string("Ace"), string("Stacy"), string("Sabrina"), string("Barkley")});  // Stacy
  cout << min({string("Ace"), string("Stacy"), string("Sabrina"), string("Barkley")});  // Ace
  cout << min({54, 16, 48, 5});  // 5
  ```

### 5. 面试相关问题
- **什么是 initializer_list？**
  - 答：initializer_list 是 C++11 提供的类模板，用于支持值列表初始化。

- **initializer_list 的特性？**
  - 答：强制值初始化；防止窄化初始化；当同时有特定参数个数的构造函数和 initializer_list 版本时，优先选择 initializer_list 版本。

- **initializer_list 的内部实现？**
  - 答：内部包含一个指向数组的指针和数组长度，编译器会自动创建临时数组并初始化 initializer_list。

---

## 八、 explicit 关键字用于多参数构造函数

### 1. 问题背景
- C++98 中，`explicit` 关键字只能用于单参数构造函数
- 防止隐式类型转换

### 2. C++11 的改进
- `explicit` 关键字可以用于多参数构造函数
- 防止使用花括号进行隐式类型转换

### 3. 示例
```cpp
struct Complex {
    int real, imag;
    explicit Complex(int re, int im=0) : real(re), imag(im) {}
    Complex operator+(const Complex& x) {
        return Complex(real + x.real, imag + x.imag);
    }
};

Complex c1(12,5);
Complex c2 = c1 + 5;  // 错误：没有匹配的 operator+（操作数类型为 'Complex' 和 'int'）
```

### 4. 使用场景
```cpp
class P {
public:
    P(int a, int b) {
        cout << "P(int a, int b)\n";
    }
    P(initializer_list<int>) {
        cout << "P(initializer_list<int>)\n";
    }
    explicit P(int a, int b, int c) {
        cout << "explicit P(int a, int b, int c)\n";
    }
};

void fp(const P&) {}

P p1(77, 5);          // P(int a, int b)
P p2 {77, 5};          // P(initializer_list<int>)
P p3 {77, 5, 42};      // P(initializer_list<int>)
// P p4 = {77, 5, 42};  // 错误：从 initializer list 转换到 'P' 会使用 explicit 构造函数 'P::P'
P p6 (77,5,42);        // explicit P(int a, int b, int c)

// fp({47,11,3});       // 错误：从 initializer list 转换到 'const P' 会使用 explicit 构造函数
// fp(P{47,11,3});       // 错误：从 initializer list 转换到 'P' 会使用 explicit 构造函数
fp(P(47,11,3));        // P(initializer_list<int>)

P p11 {77, 5, 42, 500};  // P(initializer_list<int>)
P p12 = {77, 5, 42, 500}; // P(initializer_list<int>)
P p13 {10};              // P(initializer_list<int>)
```

### 5. 面试相关问题
- **explicit 关键字的作用？**
  - 答：explicit 关键字用于防止隐式类型转换，C++11 中可以用于多参数构造函数。

- **什么时候应该使用 explicit 关键字？**
  - 答：当你不希望构造函数被用于隐式类型转换时，应该使用 explicit 关键字。

- **explicit 构造函数与 initializer_list 构造函数的优先级？**
  - 答：当使用花括号初始化时，如果存在 initializer_list 构造函数，会优先选择它；如果不存在，才会考虑其他构造函数，包括 explicit 构造函数。

---

## 九、 基于范围的 for 循环：Range-based for statement

### 1. 基本概念
- C++11 引入的新循环语法
- 用于遍历容器或数组中的所有元素

### 2. 基本语法
```cpp
for (decl : coll) {
    statement
}
```

### 3. 使用示例
```cpp
// 遍历初始化列表
for (int i : { 2, 3, 5, 7, 9, 13, 17, 19 }) {
    cout << i << endl;
}

// 遍历容器
vector<double> vec;
// ...
for (auto elem : vec) {
    cout << elem << endl;
}

// 遍历容器并修改元素
for (auto& elem : vec) {
    elem *= 3;
}
```

### 4. 工作原理
- 编译器会将范围 for 循环转换为传统的 for 循环：
  ```cpp
  // 转换方式 1：使用成员函数 begin() 和 end()
  for(auto _pos=coll.begin(), _end=coll.end(); _pos!=_end; ++_pos) {
      decl = *_pos;
      statement
  }
  
  // 转换方式 2：使用全局函数 begin() 和 end()
  for(auto _pos=begin(coll), _end=end(coll); _pos!=_end; ++_pos) {
      decl = *_pos;
      statement
  }
  ```

### 5. 示例转换
```cpp
// 原始代码
template <typename T>
void printElements(const T& coll) {
    for(const auto& elem : coll) {
        cout << elem << endl;
    }
}

// 转换后的代码
template <typename T>
void printElements(const T& coll) {
    for(auto _pos=coll.begin(); _pos != coll.end(); ++_pos) {
        const auto& elem = *_pos;
        cout << elem << endl;
    }
}
```

### 6. 注意事项
- 范围 for 循环中，元素初始化时不允许显式类型转换
  ```cpp
  class C {
  public:
      explicit C(const string& s);  // explicit(!) type conversion from strings
      ...
  };

  vector<string> vs;
  for (const C& elem : vs) {  // 错误：没有从 string 到 C 的转换
      cout << elem << endl;
  }
  ```

### 7. 面试相关问题
- **什么是基于范围的 for 循环？**
  - 答：基于范围的 for 循环是 C++11 引入的新循环语法，用于遍历容器或数组中的所有元素。

- **基于范围的 for 循环的工作原理？**
  - 答：编译器会将范围 for 循环转换为传统的 for 循环，使用 begin() 和 end() 函数获取范围的起始和结束迭代器。

- **如何在基于范围的 for 循环中修改元素？**
  - 答：使用引用类型，如 `for (auto& elem : vec)`。

---

## 十、 容器 array

### 1. 基本概念
- C++11 引入的固定大小数组容器
- 是 TR1 中的特性，现在成为标准的一部分
- 内部实现是一个固定大小的数组

### 2. 实现细节
- 没有构造函数和析构函数
- 迭代器是原生指针
- 支持随机访问

### 3. 示例代码
```cpp
// 定义
array<int, 10> myArray;

// 使用
auto ite = myArray.begin();  // array<int, 10>::iterator ite = ...
ite += 3;
cout << *ite;
```

### 4. 与传统数组的对比
- **传统数组**：`int a[100];`
- **array 容器**：`array<int, 100> a;`
- **相同点**：内部实现相同，都是固定大小的数组
- **不同点**：array 容器提供了 STL 容器的接口，如 begin()、end()、size() 等

### 5. 面试相关问题
- **array 容器与传统数组的区别？**
  - 答：array 容器提供了 STL 容器的接口，如 begin()、end()、size() 等，而传统数组没有这些接口。

- **array 容器的特点？**
  - 答：固定大小，内部实现是数组，没有构造函数和析构函数，迭代器是原生指针。

- **什么时候应该使用 array 容器？**
  - 答：当需要固定大小的数组，并且希望使用 STL 算法和接口时。

---

## 十一、 总结与思考

### 1. 核心知识点总结
- **Variadic Templates**：可变参数模板，用于处理数量不定的模板参数
- **nullptr**：类型安全的空指针，避免类型混淆
- **auto**：自动类型推导，简化代码，减少类型错误
- **Uniform Initialization**：统一初始化语法，使用花括号 `{}`
- **Initializer Lists**：初始化列表，支持值列表初始化，防止窄化转换
- **explicit**：用于多参数构造函数，防止隐式类型转换
- **Range-based for**：基于范围的 for 循环，简化遍历代码
- **array**：固定大小数组容器，提供 STL 接口

### 2. 设计思想与技术要点
- **类型安全**：如 nullptr、explicit、initializer_list 的窄化检查
- **代码简化**：如 auto、range-based for
- **统一语法**：如 uniform initialization
- **编译时计算**：如 variadic templates
- **性能优化**：如 move semantics（后续课程）

### 3. 实际应用价值
- **提高代码可读性**：如 auto、range-based for
- **减少错误**：如 nullptr、explicit、initializer_list
- **增强表达能力**：如 variadic templates、lambda 表达式（后续课程）
- **提升性能**：如 move semantics、右值引用（后续课程）

### 4. 学习建议与进阶方向
- **基础巩固**：掌握 C++11/14 的核心特性
- **实践应用**：在项目中使用这些新特性
- **深入理解**：学习标准库的实现细节
- **拓展学习**：学习 C++17/20 的新特性

### 5. 面试准备
- **核心概念**：理解每个新特性的设计意图和使用方法
- **代码示例**：准备一些简洁的代码示例
- **对比分析**：与 C++98 的对比，理解改进之处
- **应用场景**：掌握每个特性的适用场景

### 6. 个人感悟
通过学习 C++11/14 的新特性，我深刻体会到 C++ 语言的不断进化和完善。这些新特性不仅简化了代码，提高了可读性，还增强了类型安全性和表达能力。

特别是：
- **Variadic Templates** 让模板编程更加灵活，能够处理任意数量的参数
- **auto** 关键字大大简化了代码，尤其是在处理复杂类型时
- **Uniform Initialization** 统一了初始化语法，减少了混淆
- **Initializer Lists** 提供了更安全的初始化方式，防止窄化转换
- **Range-based for** 简化了容器的遍历，使代码更清晰

这些特性的引入，使得 C++ 语言在保持高性能的同时，也变得更加现代和易用。作为 C++ 程序员，掌握这些新特性是必不可少的，它们将帮助我们编写更简洁、更安全、更高效的代码。

---

## 十二、 面试八股问题汇总

### 1. 基础问题

#### 1.1 C++ 标准的演化
- **问题**：C++ 标准的演化历程是怎样的？
  - **答案**：C++ 98 (1.0) → C++ 03 (TR1) → C++ 11 (2.0) → C++ 14 → C++ 17 → C++ 20

#### 1.2 头文件的变化
- **问题**：C++11 中头文件有什么变化？
  - **答案**：C++ 标准库的头文件不带后缀名 (.h)，如 `#include <vector>`；新式 C 头文件也不带后缀名 .h，如 `#include <cstdio>`。

#### 1.3 Variadic Templates
- **问题**：什么是 Variadic Templates？
  - **答案**：Variadic Templates 是 C++11 引入的特性，允许模板接受数量不定的参数。

- **问题**：Variadic Templates 与 C 风格可变参数的区别？
  - **答案**：C++ 的 Variadic Templates 是类型安全的，编译时检查，而 C 风格可变参数使用宏，类型不安全。

#### 1.4 nullptr
- **问题**：什么是 nullptr？
  - **答案**：nullptr 是 C++11 引入的关键字，表示空指针，类型为 std::nullptr_t。

- **问题**：nullptr 与 0 和 NULL 的区别？
  - **答案**：nullptr 是类型安全的，自动转换为任何指针类型，但不会转换为整数类型；而 0 是整数，NULL 通常被定义为 0 或 (void*)0，可能导致类型混淆。

#### 1.5 auto
- **问题**：什么是 auto 关键字？
  - **答案**：auto 是 C++11 引入的关键字，允许编译器自动推导变量的类型。

- **问题**：auto 的适用场景？
  - **答案**：当类型较长或复杂时（如迭代器类型），或当类型是匿名的时（如 lambda 表达式）。

#### 1.6 Uniform Initialization
- **问题**：什么是统一初始化？
  - **答案**：统一初始化是 C++11 引入的特性，使用花括号 `{}` 进行初始化，适用于任何类型。

- **问题**：统一初始化的优势？
  - **答案**：语法统一，减少混淆；强制值初始化；防止窄化转换。

#### 1.7 Initializer Lists
- **问题**：什么是 initializer_list？
  - **答案**：initializer_list 是 C++11 提供的类模板，用于支持值列表初始化。

- **问题**：initializer_list 的特性？
  - **答案**：强制值初始化；防止窄化初始化；当同时有特定参数个数的构造函数和 initializer_list 版本时，优先选择 initializer_list 版本。

#### 1.8 explicit
- **问题**：explicit 关键字的作用？
  - **答案**：explicit 关键字用于防止隐式类型转换，C++11 中可以用于多参数构造函数。

#### 1.9 Range-based for
- **问题**：什么是基于范围的 for 循环？
  - **答案**：基于范围的 for 循环是 C++11 引入的新循环语法，用于遍历容器或数组中的所有元素。

- **问题**：如何在基于范围的 for 循环中修改元素？
  - **答案**：使用引用类型，如 `for (auto& elem : vec)`。

#### 1.10 array
- **问题**：array 容器与传统数组的区别？
  - **答案**：array 容器提供了 STL 容器的接口，如 begin()、end()、size() 等，而传统数组没有这些接口。

### 2. 进阶问题

#### 2.1 Variadic Templates
- **问题**：Variadic Templates 的递归终止条件如何设计？
  - **答案**：提供一个无参数的基础版本作为递归终止条件。

- **问题**：如何在 Variadic Templates 中获取参数包的大小？
  - **答案**：使用 `sizeof...(args)` 操作符获取参数包的大小。

- **问题**：Variadic Templates 在标准库中的应用？
  - **答案**：tuple、function、bind 等都使用了 Variadic Templates。

#### 2.2 auto 与 decltype
- **问题**：auto 与 decltype 的区别？
  - **答案**：auto 用于变量声明时的类型推导，decltype 用于获取表达式的类型，不用于变量声明。

- **问题**：auto 推导的类型与原类型有什么关系？
  - **答案**：auto 推导的类型会忽略顶层 const 和引用，如需保留，需要显式指定，如 `const auto&`。

#### 2.3 Initializer Lists
- **问题**：initializer_list 的内部实现？
  - **答案**：内部包含一个指向数组的指针和数组长度，编译器会自动创建临时数组并初始化 initializer_list。

- **问题**：initializer_list 对象的生命周期？
  - **答案**：initializer_list 对象引用的临时数组的生命周期与 initializer_list 对象相同。

#### 2.4 性能问题
- **问题**：使用 auto 关键字会影响性能吗？
  - **答案**：不会，auto 只是在编译时推导类型，不会影响运行时性能。

- **问题**：基于范围的 for 循环与传统 for 循环的性能对比？
  - **答案**：性能基本相同，编译器会将范围 for 循环转换为传统 for 循环。

#### 2.5 标准库应用
- **问题**：C++11 中标准库有哪些新容器？
  - **答案**：array、forward_list、unordered_map、unordered_set 等。

- **问题**：C++11 中标准库算法有哪些改进？
  - **答案**：许多算法现在接受 initializer_list 参数，如 max、min、sort 等。

### 3. 实际应用问题

#### 3.1 代码风格
- **问题**：在项目中如何统一使用 C++11 新特性？
  - **答案**：制定编码规范，明确哪些新特性可以使用，如何使用。

- **问题**：如何处理 C++11 与旧代码的兼容性？
  - **答案**：使用条件编译，或逐步迁移旧代码。

#### 3.2 性能优化
- **问题**：如何利用 C++11 新特性提高代码性能？
  - **答案**：使用 move semantics、右值引用、lambda 表达式等。

- **问题**：如何选择合适的容器？
  - **答案**：根据具体需求选择，如需要固定大小数组使用 array，需要快速查找使用 unordered_map 等。

#### 3.3 调试与错误处理
- **问题**：如何调试使用 C++11 新特性的代码？
  - **答案**：使用支持 C++11 的调试工具，如 GDB 7.0+、Visual Studio 2012+。

- **问题**：C++11 中有哪些新的错误处理机制？
  - **答案**： noexcept 关键字、强类型枚举等。

### 4. 代码优化问题

#### 4.1 类型推导
- **问题**：如何优化使用 auto 的代码？
  - **答案**：在适当的地方使用 auto，提高代码可读性；在需要明确类型的地方，仍然使用显式类型声明。

#### 4.2 初始化
- **问题**：如何选择合适的初始化方式？
  - **答案**：优先使用统一初始化语法 `{}`，但在某些情况下（如构造函数重载歧义）可能需要使用圆括号。

#### 4.3 循环优化
- **问题**：如何优化循环代码？
  - **答案**：使用基于范围的 for 循环简化代码；对于大型容器，考虑使用迭代器而不是下标访问。

#### 4.4 模板编程
- **问题**：如何优化模板代码？
  - **答案**：使用 Variadic Templates 简化代码；使用 SFINAE 技术进行编译时类型检查；使用 type traits 进行类型特性判断。

---

## 十三、 代码示例与实践

### 1. Variadic Templates 示例

#### 1.1 打印任意数量的参数
```cpp
#include <iostream>
using namespace std;

// 基础版本（递归终止条件）
void print() {
    cout << endl;
}

// 可变参数模板版本
template <typename T, typename... Types>
void print(const T& firstArg, const Types&... args) {
    cout << firstArg << " ";
    print(args...);
}

int main() {
    print(1, 2, 3, "hello", 3.14);
    return 0;
}
// 输出：1 2 3 hello 3.14
```

#### 1.2 计算参数个数
```cpp
#include <iostream>
using namespace std;

template <typename... Types>
size_t countArgs(const Types&... args) {
    return sizeof...(args);
}

int main() {
    cout << countArgs(1, 2, 3) << endl;          // 3
    cout << countArgs("hello", 3.14) << endl;      // 2
    cout << countArgs() << endl;                  // 0
    return 0;
}
```

### 2. nullptr 示例

```cpp
#include <iostream>
using namespace std;

void f(int x) {
    cout << "f(int): " << x << endl;
}

void f(void* p) {
    cout << "f(void*): " << p << endl;
}

int main() {
    f(0);         // 调用 f(int)
    f(NULL);      // 调用 f(int)（如果 NULL 是 0）
    f(nullptr);   // 调用 f(void*)
    return 0;
}
```

### 3. auto 示例

```cpp
#include <iostream>
#include <vector>
#include <string>
using namespace std;

int main() {
    // 基本类型
    auto i = 42;              // int
    auto d = 3.14;            // double
    auto s = "hello";         // const char*
    
    // 复杂类型
    vector<string> v = {"a", "b", "c"};
    auto it = v.begin();       // vector<string>::iterator
    
    // lambda 表达式
    auto add = [](int a, int b) { return a + b; };
    cout << add(1, 2) << endl; // 3
    
    return 0;
}
```

### 4. Uniform Initialization 示例

```cpp
#include <iostream>
#include <vector>
#include <complex>
using namespace std;

int main() {
    // 基本类型
    int i {42};
    double d {3.14};
    
    // 数组
    int arr[] {1, 2, 3, 4, 5};
    
    // 容器
    vector<int> v {1, 2, 3, 4, 5};
    
    // 复杂类型
    complex<double> c {1.0, 2.0};
    
    return 0;
}
```

### 5. Initializer Lists 示例

```cpp
#include <iostream>
#include <vector>
#include <initializer_list>
using namespace std;

// 接受 initializer_list 的函数
void printList(initializer_list<int> list) {
    for (auto it = list.begin(); it != list.end(); ++it) {
        cout << *it << " ";
    }
    cout << endl;
}

// 接受 initializer_list 的类
class MyClass {
public:
    MyClass(initializer_list<int> list) {
        for (auto it = list.begin(); it != list.end(); ++it) {
            data.push_back(*it);
        }
    }
    
    void print() {
        for (auto it = data.begin(); it != data.end(); ++it) {
            cout << *it << " ";
        }
        cout << endl;
    }
    
private:
    vector<int> data;
};

int main() {
    // 调用函数
    printList({1, 2, 3, 4, 5});
    
    // 初始化对象
    MyClass obj {10, 20, 30, 40, 50};
    obj.print();
    
    return 0;
}
```

### 6. Range-based for 示例

```cpp
#include <iostream>
#include <vector>
#include <string>
using namespace std;

int main() {
    // 遍历数组
    int arr[] = {1, 2, 3, 4, 5};
    for (int i : arr) {
        cout << i << " ";
    }
    cout << endl;
    
    // 遍历容器
    vector<string> v = {"a", "b", "c", "d", "e"};
    for (const auto& s : v) {
        cout << s << " ";
    }
    cout << endl;
    
    // 遍历初始化列表
    for (auto i : {10, 20, 30, 40, 50}) {
        cout << i << " ";
    }
    cout << endl;
    
    // 修改元素
    vector<int> nums = {1, 2, 3, 4, 5};
    for (auto& num : nums) {
        num *= 2;
    }
    for (auto num : nums) {
        cout << num << " ";
    }
    cout << endl;
    
    return 0;
}
```

### 7. array 示例

```cpp
#include <iostream>
#include <array>
#include <algorithm>
using namespace std;

int main() {
    // 定义 array
    array<int, 5> arr = {1, 2, 3, 4, 5};
    
    // 访问元素
    cout << arr[0] << endl;          // 1
    cout << arr.at(2) << endl;       // 3
    
    // 遍历
    for (auto i : arr) {
        cout << i << " ";
    }
    cout << endl;
    
    // 使用 STL 算法
    sort(arr.begin(), arr.end(), greater<int>());
    for (auto i : arr) {
        cout << i << " ";
    }
    cout << endl;
    
    // 其他操作
    cout << "size: " << arr.size() << endl;          // 5
    cout << "empty: " << arr.empty() << endl;       // 0
    
    return 0;
}
```

---

## 十四、 总结

通过学习侯捷老师的 C++ 新标准 C++11&14 课程的 1-8 集，我们深入了解了 C++11 引入的一系列重要新特性，包括：

1. **Variadic Templates**：可变参数模板，为模板编程带来了更大的灵活性
2. **nullptr**：类型安全的空指针，避免了类型混淆
3. **auto**：自动类型推导，简化了代码，提高了可读性
4. **Uniform Initialization**：统一初始化语法，减少了混淆
5. **Initializer Lists**：初始化列表，提供了更安全的初始化方式
6. **explicit**：用于多参数构造函数，防止隐式类型转换
7. **Range-based for**：基于范围的 for 循环，简化了遍历代码
8. **array**：固定大小数组容器，提供了 STL 接口

这些新特性不仅简化了代码，提高了可读性，还增强了类型安全性和表达能力，使得 C++ 语言在保持高性能的同时，也变得更加现代和易用。

作为 C++ 程序员，掌握这些新特性是必不可少的，它们将帮助我们编写更简洁、更安全、更高效的代码。同时，这些特性也是面试中的常见考点，了解它们的设计意图和使用方法，对于通过技术面试也非常重要。

在后续的课程中，我们还将学习更多 C++11/14 的新特性，如 move semantics、lambda 表达式、智能指针等，这些都将进一步提升我们的 C++ 编程能力。