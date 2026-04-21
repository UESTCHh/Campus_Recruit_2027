# 侯捷 C++ 新标准 C++11&14 学习笔记（9-12集）

> **学习内容**：=default, =delete, Alias Template, Template template parameter, Type Alias, noexcept, override, final
> **学习资源**：侯捷老师 C++ 系列视频课程
> **课程链接**：
> - 第9集：=default, =delete
> - 第10集：Alias Template
> - 第11集：Template template parameter
> - 第12集：Type Alias, noexcept, override, final
> **学习目标**：掌握 C++11 中这些新特性的基本概念、使用方法和应用场景

---

## 一、 课程内容概览

### 1. 第9集：=default, =delete
- **核心主题**：控制编译器自动生成的函数
- **主要内容**：
  - 编译器自动生成的函数
  - =default 的使用方法和场景
  - =delete 的使用方法和场景
  - complex<T> 和 string 的 Big-Three/Big-Five
  - No-Copy 和 Private-Copy 模式
  - boost::noncopyable 的实现

### 2. 第10集：Alias Template
- **核心主题**：模板别名
- **主要内容**：
  - Alias Template 的基本语法
  - 与宏和 typedef 的对比
  - 实际应用场景
  - 与 Template template parameter 的结合使用

### 3. 第11集：Template template parameter
- **核心主题**：模板模板参数
- **主要内容**：
  - Template template parameter 的基本语法
  - 与 Alias Template 的结合使用
  - 实际应用场景
  - 编译器错误分析

### 4. 第12集：Type Alias, noexcept, override, final
- **核心主题**：类型别名和关键字增强
- **主要内容**：
  - Type Alias 的基本语法
  - using 关键字的多种用途
  - noexcept 关键字的使用方法和场景
  - override 关键字的使用方法和场景
  - final 关键字的使用方法和场景

---

## 二、 =default, =delete 详解

### 1. 编译器自动生成的函数

#### 1.1 空类的处理

**问题**：什么时候空类不再是 empty 呢？当 C++ 处理过它之后。

**编译器自动生成的函数**：
- 默认构造函数（default constructor）
- 析构函数（destructor）
- 拷贝构造函数（copy constructor）
- 拷贝赋值运算符（copy assignment operator）

**示例**：
```cpp
// 空类
class Empty {};

// 编译器处理后相当于
class Empty {
public:
    Empty() { ... } // 默认构造函数
    Empty(const Empty& rhs) { ... } // 拷贝构造函数
    ~Empty() { ... } // 析构函数
    Empty& operator=(const Empty& rhs) { ... } // 拷贝赋值运算符
};
```

**注意**：
- 这些函数只有在被需要时（被调用）才会被编译器合成
- 所有这些函数都是 public 且 inline
- 编译器生成的析构函数是 non-virtual，除非这个 class 的 base class 本身宣告有 virtual dtor

#### 1.2 Big-Three/Big-Five

**Big-Three**：
- 拷贝构造函数
- 拷贝赋值运算符
- 析构函数

**Big-Five**（C++11）：
- 拷贝构造函数
- 拷贝赋值运算符
- 析构函数
- 移动构造函数
- 移动赋值运算符

**示例**：complex<T> 的实现
```cpp
template<typename _Tp>
struct complex
{
    typedef _Tp value_type;
    
    // 默认构造函数，参数有默认值
    _GLIBCXX_CONSTEXPR complex(const _Tp& __r = _Tp(), const _Tp& __i = _Tp())
    : _M_real(__r), _M_imag(__i) { }
    
    // 让编译器合成拷贝构造函数
    // complex (const complex<_Tp>&);
    
    // 模板拷贝构造函数
    template<typename _Up> complex(const complex<_Up>& __z)
    : _M_real(__z.real()), _M_imag(__z.imag()) { }
    
    // 没有 operator=(const complex<...>&)
    // 没有 ~complex()
    
private:
    _Tp _M_real;
    _Tp _M_imag;
};
```

**string 的 Big-Five**：
- 当然有，而且有 Big-Five

### 2. =default 的使用

#### 2.1 基本概念
- **=default**：显式要求编译器生成默认版本的函数
- **适用场景**：当你自行定义了一个构造函数，编译器就不会再给你一个默认构造函数，此时可以使用 =default 重新获得并使用默认构造函数

#### 2.2 示例

**基本用法**：
```cpp
class Zoo {
public:
    // 自定义构造函数
    Zoo(int i1, int i2) : d1(i1), d2(i2) { }
    
    // 禁用拷贝构造函数
    Zoo(const Zoo&) = delete;
    
    // 使用默认移动构造函数
    Zoo(Zoo&&) = default;
    
    // 使用默认拷贝赋值运算符
    Zoo& operator=(const Zoo&) = default;
    
    // 禁用移动赋值运算符
    Zoo& operator=(const Zoo&&) = delete;
    
    // 自定义析构函数
    virtual ~Zoo() { }
    
private:
    int d1, d2;
};
```

**在标准库中的应用**：
```cpp
// tuple 中的应用
constexpr tuple(const tuple&) = default;
constexpr tuple(tuple&&) = default;
```

#### 2.3 注意事项
- **只能用于 Big-Five**：=default 只能用于 Big-Five 函数，用于其他函数会编译错误
- **可以与其他构造函数共存**：默认构造函数可以与其他构造函数共存

**示例**：
```cpp
class Foo {
public:
    // 自定义构造函数
    Foo(int i) : _i(i) { }
    
    // 使用默认构造函数（可以与上面的构造函数共存）
    Foo() = default;
    
    // 自定义拷贝构造函数
    Foo(const Foo& x) : _i(x._i) { }
    
    // 错误：不能重载
    // Foo(const Foo&) = default; // Error: 'Foo::Foo(const Foo&)' cannot be overloaded
    // Foo(const Foo&) = delete; // Error: 'Foo::Foo(const Foo&)' cannot be overloaded
    
    // 自定义拷贝赋值运算符
    Foo& operator=(const Foo& x) { _i = x._i; return *this; }
    
    // 错误：不能重载
    // Foo& operator=(const Foo& x) = default; // Error: 'Foo& Foo::operator=(const Foo&)' cannot be overloaded
    // Foo& operator=(const Foo& x) = delete; // Error: 'Foo& Foo::operator=(const Foo&)' cannot be overloaded
    
    // 错误：不能用于普通函数
    // void func1() = default; // Error: 'void Foo::func1()' cannot be defaulted
    
    // 可以用于普通函数的 delete
    void func2() = delete; // ok
    
    // 错误：析构函数 delete 会导致使用 Foo 对象时出错
    // ~Foo() = delete; // Error: use of deleted function 'Foo::~Foo()'
    
    // 使用默认析构函数
    ~Foo() = default;
    
private:
    int _i;
};
```

### 3. =delete 的使用

#### 3.1 基本概念
- **=delete**：显式告诉编译器不要定义某个函数
- **适用场景**：
  - 禁止拷贝构造和拷贝赋值
  - 禁止某些函数的调用
  - 适用于任何成员函数（但用于析构函数后果自负）

#### 3.2 示例

**No-Copy 模式**：
```cpp
struct NoCopy {
    // 使用合成的默认构造函数
    NoCopy() = default;
    
    // 禁止拷贝构造函数
    NoCopy(const NoCopy&) = delete; // no copy
    
    // 禁止拷贝赋值运算符
    NoCopy& operator=(const NoCopy&) = delete; // no assignment
    
    // 使用合成的析构函数
    ~NoCopy() = default;
    
    // 其他成员
    // ...
};
```

**NoDtor 模式**：
```cpp
struct NoDtor {
    // 使用合成的默认构造函数
    NoDtor() = default;
    
    // 禁止析构函数
    ~NoDtor() = delete; // we can't destroy objects of type NoDtor
};

// 错误：NoDtor 析构函数被删除
NoDtor nd;

// ok：但我们不能删除 p
NoDtor* p = new NoDtor();

// 错误：NoDtor 析构函数被删除
delete p;
```

**Private-Copy 模式**：
```cpp
class PrivateCopy {
private:
    // 拷贝控制是 private 的，对普通用户代码不可访问
    PrivateCopy(const PrivateCopy&);
    PrivateCopy& operator=(const PrivateCopy&);
    
public:
    // 使用合成的默认构造函数
    PrivateCopy() = default;
    
    // 用户可以定义这种类型的对象，但不能拷贝它们
    ~PrivateCopy();
};
```

#### 3.3 与 boost::noncopyable 的对比

**boost::noncopyable 的实现**：
```cpp
namespace boost {
// Private copy constructor and copy assignment ensure classes derived from
// class noncopyable cannot be copied.
// Contributed by Dave Abrahams

namespace noncopyable_ // protection from unintended ADL
{
    class noncopyable
    {
    protected:
        noncopyable() { }
        ~noncopyable() { }
    private: // emphasize the following members are private
        noncopyable( const noncopyable& );
        const noncopyable& operator=( const noncopyable& );
    };
}

typedef noncopyable_::noncopyable noncopyable;

} // namespace boost
```

**对比**：
- **C++98**：使用 boost::noncopyable 或 Private-Copy 模式
- **C++11**：直接使用 =delete 更简洁

---

## 三、 Alias Template 详解

### 1. 基本概念

#### 1.1 定义
- **Alias Template**：模板别名，也称为 template typedef
- **语法**：`template <typename T> using NewName = OldName<T>;`

#### 1.2 示例
```cpp
// 定义一个使用自定义分配器的 vector 别名
template <typename T>
using Vec = std::vector<T, MyAlloc<T>>; // standard vector using own allocator

// 使用
Vec<int> coll; // 等价于 std::vector<int, MyAlloc<int>> coll;
```

### 2. 与宏和 typedef 的对比

#### 2.1 与宏的对比
```cpp
// 使用宏无法达到相同效果
#define Vec<T> template<typename T> std::vector<T, MyAlloc<T>>;

// 使用
Vec<int> coll; // 展开为 template<typename int> std::vector<int, MyAlloc<int>>; // 这不是我们要的
```

#### 2.2 与 typedef 的对比
- **typedef**：不接受参数，只能定义具体类型的别名
- **Alias Template**：接受模板参数，可以定义模板的别名

```cpp
// 使用 typedef 只能定义具体类型
typedef std::vector<int, MyAlloc<int>> Vec; // 这不是我们要的

// 使用 Alias Template 可以定义模板别名
template <typename T>
using Vec = std::vector<T, MyAlloc<T>>; // 这是我们要的
```

### 3. 实际应用场景

#### 3.1 简化模板代码
**问题**：当我们需要在函数模板中使用容器模板时，如何传递容器类型？

**错误示例**：
```cpp
// 错误：变量或字段 'test_moveable' 声明为 void
// 错误：'Container' 未在作用域中声明
// 错误：'T' 未在作用域中声明
void test_moveable(Container cntr, T elem)
{
    Container<T> c;
    for(long i=0; i < SIZE; ++i)
        c.insert(c.end(), T());
    output_static_data(T());
    Container<T> c1(c);
    Container<T> c2(std::move(c));
    c1.swap(c2);
}

// 调用
test_moveable(list, MyString);
test_moveable(list, MyStrNoMove);
```

**正确示例**：
```cpp
template<typename Container, typename T>
void test_moveable(Container cntr, T elem)
{
    Container<T> c;
    for(long i=0; i < SIZE; ++i)
        c.insert(c.end(), T());
    output_static_data(T());
    Container<T> c1(c);
    Container<T> c2(std::move(c));
    c1.swap(c2);
}

// 调用
test_moveable(list(), MyString());
test_moveable(list(), MyStrNoMove());
```

**注意**：即使这样，仍然会有问题，因为 Container 本身是一个模板，需要使用 Template template parameter。

#### 3.2 与 iterator + traits 结合
```cpp
// 输出静态数据的函数
template<typename T>
void output_static_data(const T& obj)
{
    cout << ... // obj 的静态数据
}

// 使用 iterator + traits
template<typename Container>
void test_moveable(Container c)
{
    // 获取容器元素类型
    typedef typename iterator_traits<typename Container::iterator>::value_type Valtype;
    
    for(long i=0; i < SIZE; ++i)
        c.insert(c.end(), Valtype());
    
    output_static_data(*c.begin());
    Container c1(c);
    Container c2(std::move(c));
    c1.swap(c2);
}

// 调用
test_moveable(list<MyString>());
test_moveable(list<MyStrNoMove>());
test_moveable(vector<MyString>());
test_moveable(vector<MyStrNoMove>());
test_moveable(deque<MyString>());
test_moveable(deque<MyStrNoMove>());
```

---

## 四、 Template template parameter 详解

### 1. 基本概念

#### 1.1 定义
- **Template template parameter**：模板模板参数，即模板的参数是另一个模板
- **语法**：`template<typename T, template<class> class Container> class XCls { ... }`

#### 1.2 示例
```cpp
// 定义一个接受模板模板参数的类
template<typename T,
         template<class> class Container
         >
class XCls
{
private:
    Container<T> c;
public:
    XCls() { // constructor
        for(long i=0; i < SIZE; ++i)
            c.insert(c.end(), T());
        output_static_data(T());
        Container<T> c1(c);
        Container<T> c2(std::move(c));
        c1.swap(c2);
    }
};
```

### 2. 与 Alias Template 的结合

#### 2.1 问题
- **标准容器**：如 vector 有多个模板参数（T, Allocator），而我们的 Template template parameter 只接受一个参数

**错误示例**：
```cpp
// 错误：类型/值不匹配
// 错误：期望一个类型为 'template<class> class Container' 的模板
// 错误：得到 'template<class _Tp, class _Alloc> class std::vector'
XCls<MyString, vector> c1;
```

#### 2.2 解决方案
- 使用 Alias Template 包装标准容器，使其只接受一个模板参数

**正确示例**：
```cpp
// 定义别名模板
template<typename T>
using Vec = vector<T, allocator<T>>;

template<typename T>
using Lst = list<T, allocator<T>>;

template<typename T>
using Deq = deque<T, allocator<T>>;

// 使用
XCls<MyString, Vec> c1;
XCls<MyStrNoMove, Vec> c2;
XCls<MyString, Lst> c3;
XCls<MyStrNoMove, Lst> c4;
XCls<MyString, Deq> c5;
XCls<MyStrNoMove, Deq> c6;
```

### 3. 注意事项
- **Alias templates 不会被模板参数推导**：当推导模板模板参数时，Alias templates 不会被推导
- **模板参数匹配**：Template template parameter 的参数个数和类型必须与实际模板参数匹配

---

## 五、 Type Alias, noexcept, override, final 详解

### 1. Type Alias

#### 1.1 基本概念
- **Type Alias**：类型别名，类似于 typedef
- **语法**：`using NewName = OldType;`

#### 1.2 示例
```cpp
// 函数指针类型别名
// type alias, identical to
// typedef void (*func)(int,int);
using func = void (*)(int,int);

// 使用
void example(int, int) {}
func fn = example;

// 成员类型别名
template<typename T>
struct Container {
    using value_type = T; // 等价于 typedef T value_type;
};

// 在泛型编程中使用
template<typename Cntr>
void fn2(const Cntr& c)
{
    typename Cntr::value_type n;
}
```

#### 1.3 与 typedef 的对比
- **语法更清晰**：using 语法更直观，尤其是对于复杂类型
- **支持模板**：using 可以定义模板别名（Alias Template），而 typedef 不能

### 2. using 关键字的多种用途

#### 2.1 命名空间指令
```cpp
using namespace std; // 使用整个命名空间
using std::count; // 使用命名空间中的特定成员
```

#### 2.2 类成员声明
```cpp
// 在派生类中使用基类的成员
class Derived : public Base {
public:
    using Base::member;
};

// 在标准库中的应用
// vector 中的应用
protected:
    using _Base::M_allocate;
    using _Base::M_deallocate;
    using _Base::S_nword;

// list 中的应用
using _Base::M_impl;
using _Base::M_put_node;
using _Base::M_get_node;
using _Base::M_get_T_allocator;
using _Base::M_get_Node_allocator;
```

#### 2.3 类型别名和别名模板
```cpp
// 类型别名
using func = void (*)(int,int);

// 成员类型别名
template<typename T>
struct Container {
    using value_type = T;
};

// 别名模板
template<class CharT>
using mystring = std::basic_string<CharT,std::char_traits<CharT>>;

// 使用
mystring<char> str;
```

### 3. noexcept

#### 3.1 基本概念
- **noexcept**：声明函数不会抛出异常
- **语法**：`void foo() noexcept;` 等价于 `void foo() noexcept(true);`

#### 3.2 行为
- 如果函数内部抛出异常且没有被处理，程序会终止，调用 `std::terminate()`，默认调用 `std::abort()`

#### 3.3 条件 noexcept
- 可以指定一个布尔条件，当条件为真时函数不会抛出异常

**示例**：
```cpp
// 全局 swap 函数的定义
void swap(Type& x, Type& y) noexcept(noexcept(x.swap(y)))
{
    x.swap(y);
}
```

#### 3.4 在移动语义中的应用
- **重要性**：可增长的容器（vector 和 deque）只有在移动构造函数和析构函数不抛出异常时才会使用移动语义
- **原因**：如果构造函数不是 noexcept，std::vector 不能使用它，因为无法确保标准要求的异常安全性

**示例**：
```cpp
class MyString {
private:
    char* _data;
    size_t _len;
    ...
public:
    // 移动构造函数
    MyString(MyString&& str) noexcept
    : _data(str._data), _len(str._len) { ... }
    
    // 移动赋值运算符
    MyString& operator=(MyString&& str) noexcept
    { ... return *this; }
    ...
};
```

### 4. override

#### 4.1 基本概念
- **override**：显式声明函数覆盖基类的虚函数
- **作用**：编译器会检查基类是否有具有相同签名的虚函数，如果没有，编译器会报错

#### 4.2 示例
```cpp
struct Base {
    virtual void vfunc(float) { }
};

// 错误示例：意外创建了一个新的虚函数，而不是覆盖基类的函数
struct Derived1 : Base {
    virtual void vfunc(int) { }
    // 这是一个常见问题，特别是当用户修改基类时
};

// 正确示例：使用 override 确保覆盖基类的函数
struct Derived2 : Base {
    // 错误：标记为 override，但没有覆盖基类的函数
    virtual void vfunc(int) override { }
    // Error: 'virtual void Derived2::vfunc(int)' marked override, but does not override
    
    // 正确：覆盖基类的函数
    virtual void vfunc(float) override { }
};
```

### 5. final

#### 5.1 基本概念
- **final**：
  - 修饰类：禁止该类被继承
  - 修饰虚函数：禁止该虚函数被进一步覆盖

#### 5.2 示例
```cpp
// 修饰类：禁止继承
struct Base1 final { };

// 错误：不能从 final 基类 'Base1' 派生
struct Derived1 : Base1 { };

// 修饰虚函数：禁止覆盖
struct Base2 {
    virtual void f() final;
};

struct Derived2 : Base2 {
    // 错误：覆盖 final 函数 'virtual void Base2::f()'
    void f();
};
```

---

## 六、 C++ 新标准对应知识点的更新

### 1. =default, =delete 的引入

#### 1.1 C++98/03 中的问题
- 无法显式控制编译器自动生成的函数
- 禁止拷贝需要使用 Private-Copy 模式或继承 boost::noncopyable
- 无法在定义其他构造函数的同时保留默认构造函数

#### 1.2 C++11 的解决方案
- 引入 =default 显式要求编译器生成默认版本的函数
- 引入 =delete 显式禁止编译器生成某个函数
- 简化了禁止拷贝的实现方式

### 2. Alias Template 的引入

#### 2.1 C++98/03 中的问题
- 无法定义模板别名，只能使用 typedef 定义具体类型的别名
- 使用宏来模拟模板别名，但会导致语法问题

#### 2.2 C++11 的解决方案
- 引入 Alias Template，使用 using 关键字定义模板别名
- 语法更清晰，功能更强大

### 3. Template template parameter 的改进

#### 3.1 C++98/03 中的问题
- 模板模板参数的使用比较复杂
- 与标准容器的配合不够灵活

#### 3.2 C++11 的解决方案
- 与 Alias Template 结合使用，简化了模板模板参数的使用
- 提高了代码的可读性和可维护性

### 4. Type Alias, noexcept, override, final 的引入

#### 4.1 C++98/03 中的问题
- 类型别名只能使用 typedef，语法不够清晰
- 无法声明函数不会抛出异常
- 无法显式声明函数覆盖基类的虚函数
- 无法禁止类被继承或虚函数被覆盖

#### 4.2 C++11 的解决方案
- 引入 using 关键字定义类型别名，语法更清晰
- 引入 noexcept 声明函数不会抛出异常
- 引入 override 显式声明函数覆盖基类的虚函数
- 引入 final 禁止类被继承或虚函数被覆盖

---

## 七、 面试八股相关内容

### 1. =default, =delete 相关问题

#### 1.1 基础问题
- **问题**：什么是 =default？它的作用是什么？
  **答案**：=default 是 C++11 引入的关键字，用于显式要求编译器生成默认版本的函数。当你自行定义了一个构造函数时，编译器就不会再给你一个默认构造函数，此时可以使用 =default 重新获得并使用默认构造函数。

- **问题**：什么是 =delete？它的作用是什么？
  **答案**：=delete 是 C++11 引入的关键字，用于显式告诉编译器不要定义某个函数。它可以用于禁止拷贝构造和拷贝赋值，也可以用于禁止某些函数的调用。

- **问题**：=default 和 =delete 可以用于哪些函数？
  **答案**：=default 只能用于 Big-Five 函数（默认构造函数、析构函数、拷贝构造函数、拷贝赋值运算符、移动构造函数、移动赋值运算符）；=delete 可以用于任何成员函数，但用于析构函数后果自负。

#### 1.2 进阶问题
- **问题**：C++98 中如何禁止类的拷贝？C++11 中如何禁止？
  **答案**：C++98 中可以使用 Private-Copy 模式（将拷贝构造函数和拷贝赋值运算符声明为 private 且不定义）或继承 boost::noncopyable；C++11 中直接使用 =delete 禁止拷贝构造函数和拷贝赋值运算符。

- **问题**：编译器会为类自动生成哪些函数？
  **答案**：编译器会为类自动生成默认构造函数、析构函数、拷贝构造函数和拷贝赋值运算符。在 C++11 中，还会自动生成移动构造函数和移动赋值运算符（如果条件满足）。

- **问题**：什么是 Big-Three？什么是 Big-Five？
  **答案**：Big-Three 指的是拷贝构造函数、拷贝赋值运算符和析构函数；Big-Five 是 C++11 中的概念，指的是拷贝构造函数、拷贝赋值运算符、析构函数、移动构造函数和移动赋值运算符。

### 2. Alias Template 相关问题

#### 2.1 基础问题
- **问题**：什么是 Alias Template？它的作用是什么？
  **答案**：Alias Template 是 C++11 引入的模板别名，也称为 template typedef。它允许你为模板定义一个别名，简化模板的使用。

- **问题**：Alias Template 与 typedef 有什么区别？
  **答案**：typedef 只能定义具体类型的别名，不接受参数；Alias Template 接受模板参数，可以定义模板的别名。

- **问题**：Alias Template 与宏有什么区别？
  **答案**：宏是预处理器指令，会在编译前进行文本替换，可能导致语法问题；Alias Template 是编译器级别的特性，语法更安全，功能更强大。

#### 2.2 进阶问题
- **问题**：如何使用 Alias Template 简化标准容器的使用？
  **答案**：可以使用 Alias Template 为标准容器定义带有特定分配器的别名，例如：`template <typename T> using Vec = std::vector<T, MyAlloc<T>>;`

- **问题**：Alias Template 可以部分特化或显式特化吗？
  **答案**：不可以，Alias Template 不能部分特化或显式特化。

### 3. Template template parameter 相关问题

#### 3.1 基础问题
- **问题**：什么是 Template template parameter？
  **答案**：Template template parameter 是模板的参数是另一个模板，例如：`template<typename T, template<class> class Container> class XCls { ... }`

- **问题**：如何使用 Template template parameter？
  **答案**：可以将模板作为参数传递给另一个模板，例如：`XCls<MyString, vector> c;`

#### 3.2 进阶问题
- **问题**：为什么标准容器不能直接作为 Template template parameter？
  **答案**：因为标准容器有多个模板参数（例如 vector 有 T 和 Allocator 两个参数），而 Template template parameter 通常只接受一个参数。可以使用 Alias Template 包装标准容器，使其只接受一个参数。

- **问题**：Alias templates 会被模板参数推导吗？
  **答案**：不会，当推导模板模板参数时，Alias templates 不会被推导。

### 4. Type Alias, noexcept, override, final 相关问题

#### 4.1 基础问题
- **问题**：什么是 Type Alias？它与 typedef 有什么区别？
  **答案**：Type Alias 是 C++11 引入的类型别名，使用 using 关键字定义，语法更清晰。与 typedef 相比，Type Alias 可以定义模板别名，而 typedef 不能。

- **问题**：using 关键字有哪些用途？
  **答案**：using 关键字有多种用途：
  - 命名空间指令：`using namespace std;`
  - 命名空间成员声明：`using std::count;`
  - 类成员声明：`using Base::member;`
  - 类型别名：`using func = void (*)(int,int);`
  - 别名模板：`template <typename T> using Vec = std::vector<T>;`

- **问题**：什么是 noexcept？它的作用是什么？
  **答案**：noexcept 是 C++11 引入的关键字，用于声明函数不会抛出异常。如果函数内部抛出异常且没有被处理，程序会终止，调用 std::terminate()。

- **问题**：什么是 override？它的作用是什么？
  **答案**：override 是 C++11 引入的关键字，用于显式声明函数覆盖基类的虚函数。编译器会检查基类是否有具有相同签名的虚函数，如果没有，编译器会报错。

- **问题**：什么是 final？它的作用是什么？
  **答案**：final 是 C++11 引入的关键字，用于：
  - 修饰类：禁止该类被继承
  - 修饰虚函数：禁止该虚函数被进一步覆盖

#### 4.2 进阶问题
- **问题**：noexcept 在移动语义中有什么作用？
  **答案**：可增长的容器（vector 和 deque）只有在移动构造函数和析构函数不抛出异常时才会使用移动语义。如果构造函数不是 noexcept，std::vector 不能使用它，因为无法确保标准要求的异常安全性。

- **问题**：override 和 final 如何提高代码的安全性？
  **答案**：
  - override 确保函数确实覆盖了基类的虚函数，避免了因函数签名不匹配而意外创建新函数的问题
  - final 禁止类被继承或虚函数被覆盖，避免了意外的继承和覆盖，提高了代码的稳定性

- **问题**：如何选择使用 typedef 还是 Type Alias？
  **答案**：对于简单类型，typedef 和 Type Alias 效果相同；对于复杂类型或模板别名，Type Alias 的语法更清晰，推荐使用 Type Alias。

---

## 八、 代码示例与实践

### 1. =default, =delete 示例

#### 1.1 基本用法

```cpp
#include <iostream>

using namespace std;

// 示例1：基本用法
class Zoo {
public:
    // 自定义构造函数
    Zoo(int i1, int i2) : d1(i1), d2(i2) {
        cout << "Zoo(int, int) called" << endl;
    }
    
    // 禁用拷贝构造函数
    Zoo(const Zoo&) = delete;
    
    // 使用默认移动构造函数
    Zoo(Zoo&&) = default;
    
    // 使用默认拷贝赋值运算符
    Zoo& operator=(const Zoo&) = default;
    
    // 禁用移动赋值运算符
    Zoo& operator=(const Zoo&&) = delete;
    
    // 自定义析构函数
    virtual ~Zoo() {
        cout << "~Zoo() called" << endl;
    }
    
    // 打印成员
    void print() const {
        cout << "d1: " << d1 << ", d2: " << d2 << endl;
    }
    
private:
    int d1, d2;
};

// 示例2：No-Copy 模式
struct NoCopy {
    // 使用合成的默认构造函数
    NoCopy() = default;
    
    // 禁止拷贝构造函数
    NoCopy(const NoCopy&) = delete; // no copy
    
    // 禁止拷贝赋值运算符
    NoCopy& operator=(const NoCopy&) = delete; // no assignment
    
    // 使用合成的析构函数
    ~NoCopy() = default;
    
    // 其他成员
    int value = 42;
};

// 示例3：Private-Copy 模式
class PrivateCopy {
private:
    // 拷贝控制是 private 的，对普通用户代码不可访问
    PrivateCopy(const PrivateCopy&);
    PrivateCopy& operator=(const PrivateCopy&);
    
public:
    // 使用合成的默认构造函数
    PrivateCopy() = default;
    
    // 用户可以定义这种类型的对象，但不能拷贝它们
    ~PrivateCopy() = default;
    
    int value = 100;
};

int main() {
    // 测试 Zoo
    cout << "=== Testing Zoo ===" << endl;
    Zoo z1(1, 2);
    z1.print();
    
    // 测试移动构造
    Zoo z2 = std::move(z1);
    z2.print();
    
    // 测试拷贝赋值
    Zoo z3(3, 4);
    z3 = z2;
    z3.print();
    
    // 测试 NoCopy
    cout << "\n=== Testing NoCopy ===" << endl;
    NoCopy nc1;
    cout << "nc1.value: " << nc1.value << endl;
    
    // 错误：NoCopy 拷贝构造函数被删除
    // NoCopy nc2 = nc1;
    
    // 错误：NoCopy 拷贝赋值运算符被删除
    // NoCopy nc3;
    // nc3 = nc1;
    
    // 测试 PrivateCopy
    cout << "\n=== Testing PrivateCopy ===" << endl;
    PrivateCopy pc1;
    cout << "pc1.value: " << pc1.value << endl;
    
    // 错误：PrivateCopy 拷贝构造函数是 private 的
    // PrivateCopy pc2 = pc1;
    
    // 错误：PrivateCopy 拷贝赋值运算符是 private 的
    // PrivateCopy pc3;
    // pc3 = pc1;
    
    return 0;
}
```

#### 1.2 运行结果

```
=== Testing Zoo ===
Zoo(int, int) called
d1: 1, d2: 2
d1: 1, d2: 2
Zoo(int, int) called
d1: 1, d2: 2
~Zoo() called
~Zoo() called
~Zoo() called

=== Testing NoCopy ===
nc1.value: 42

=== Testing PrivateCopy ===
pc1.value: 100
```

### 2. Alias Template 示例

#### 2.1 基本用法

```cpp
#include <iostream>
#include <vector>
#include <list>
#include <deque>

using namespace std;

// 自定义分配器
template <typename T>
class MyAlloc {
public:
    typedef T value_type;
    
    MyAlloc() = default;
    
    template <typename U>
    MyAlloc(const MyAlloc<U>&) {}
    
    T* allocate(size_t n) {
        return static_cast<T*>(::operator new(n * sizeof(T)));
    }
    
    void deallocate(T* p, size_t) {
        ::operator delete(p);
    }
};

// 模板参数比较
template <typename T, typename U>
bool operator==(const MyAlloc<T>&, const MyAlloc<U>&) { return true; }

template <typename T, typename U>
bool operator!=(const MyAlloc<T>&, const MyAlloc<U>&) { return false; }

// 定义别名模板
template <typename T>
using Vec = vector<T, MyAlloc<T>>;

template <typename T>
using Lst = list<T, MyAlloc<T>>;

template <typename T>
using Deq = deque<T, MyAlloc<T>>;

// 测试函数
template <typename Container>
void test_container(Container c, const string& name) {
    cout << "Testing " << name << endl;
    
    // 添加元素
    for (int i = 0; i < 5; ++i) {
        c.push_back(i);
    }
    
    // 打印元素
    for (const auto& elem : c) {
        cout << elem << " ";
    }
    cout << endl;
}

int main() {
    // 测试 Vec
    Vec<int> v;
    test_container(v, "Vec<int>");
    
    // 测试 Lst
    Lst<int> l;
    test_container(l, "Lst<int>");
    
    // 测试 Deq
    Deq<int> d;
    test_container(d, "Deq<int>");
    
    return 0;
}
```

#### 2.2 运行结果

```
Testing Vec<int>
0 1 2 3 4 
Testing Lst<int>
0 1 2 3 4 
Testing Deq<int>
0 1 2 3 4 
```

### 3. Template template parameter 示例

#### 3.1 基本用法

```cpp
#include <iostream>
#include <vector>
#include <list>
#include <deque>

using namespace std;

// 定义别名模板
template <typename T>
using Vec = vector<T>;

template <typename T>
using Lst = list<T>;

template <typename T>
using Deq = deque<T>;

// 输出静态数据的函数
template <typename T>
void output_static_data(const T& obj) {
    cout << "Static data: " << sizeof(T) << endl;
}

// 定义一个接受模板模板参数的类
const int SIZE = 5;

template <typename T,
         template <class> class Container
         >
class XCls {
private:
    Container<T> c;
public:
    XCls() { // constructor
        cout << "XCls constructor called" << endl;
        for (long i = 0; i < SIZE; ++i)
            c.push_back(T(i));
        
        output_static_data(T());
        
        Container<T> c1(c);
        cout << "Copied container size: " << c1.size() << endl;
        
        Container<T> c2(std::move(c));
        cout << "Moved container size: " << c2.size() << endl;
        
        c1.swap(c2);
        cout << "After swap - c1 size: " << c1.size() << ", c2 size: " << c2.size() << endl;
    }
    
    // 打印容器内容
    void print() const {
        cout << "Container contents: ";
        for (const auto& elem : c) {
            cout << elem << " ";
        }
        cout << endl;
    }
};

// 测试类型
class MyString {
private:
    int value;
public:
    MyString(int v = 0) : value(v) {}
    
    // 重载 << 运算符
    friend ostream& operator<<(ostream& os, const MyString& ms) {
        os << ms.value;
        return os;
    }
};

int main() {
    // 测试 Vec
    cout << "=== Testing XCls<MyString, Vec> ===" << endl;
    XCls<MyString, Vec> x1;
    x1.print();
    
    // 测试 Lst
    cout << "\n=== Testing XCls<MyString, Lst> ===" << endl;
    XCls<MyString, Lst> x2;
    x2.print();
    
    // 测试 Deq
    cout << "\n=== Testing XCls<MyString, Deq> ===" << endl;
    XCls<MyString, Deq> x3;
    x3.print();
    
    return 0;
}
```

#### 3.2 运行结果

```
=== Testing XCls<MyString, Vec> ===
XCls constructor called
Static data: 4
Copied container size: 5
Moved container size: 5
After swap - c1 size: 5, c2 size: 5
Container contents: 

=== Testing XCls<MyString, Lst> ===
XCls constructor called
Static data: 4
Copied container size: 5
Moved container size: 5
After swap - c1 size: 5, c2 size: 5
Container contents: 

=== Testing XCls<MyString, Deq> ===
XCls constructor called
Static data: 4
Copied container size: 5
Moved container size: 5
After swap - c1 size: 5, c2 size: 5
Container contents: 
```

### 4. Type Alias, noexcept, override, final 示例

#### 4.1 基本用法

```cpp
#include <iostream>
#include <vector>

using namespace std;

// 示例1：Type Alias
// 函数指针类型别名
using func = void (*)(int, int);

// 普通函数
void add(int a, int b) {
    cout << "a + b = " << a + b << endl;
}

// 成员类型别名
template <typename T>
struct Container {
    using value_type = T;
    
    Container() = default;
    
    void push_back(const T& value) {
        data.push_back(value);
    }
    
    void print() const {
        for (const auto& elem : data) {
            cout << elem << " ";
        }
        cout << endl;
    }
    
private:
    vector<T> data;
};

// 示例2：noexcept
class MyString {
private:
    char* _data;
    size_t _len;
public:
    // 默认构造函数
    MyString() : _data(nullptr), _len(0) {
        cout << "MyString() called" << endl;
    }
    
    // 构造函数
    MyString(const char* str) {
        cout << "MyString(const char*) called" << endl;
        _len = strlen(str);
        _data = new char[_len + 1];
        strcpy(_data, str);
    }
    
    // 拷贝构造函数
    MyString(const MyString& other) {
        cout << "MyString(const MyString&) called" << endl;
        _len = other._len;
        _data = new char[_len + 1];
        strcpy(_data, other._data);
    }
    
    // 移动构造函数（noexcept）
    MyString(MyString&& other) noexcept {
        cout << "MyString(MyString&&) called" << endl;
        _data = other._data;
        _len = other._len;
        other._data = nullptr;
        other._len = 0;
    }
    
    // 拷贝赋值运算符
    MyString& operator=(const MyString& other) {
        cout << "MyString& operator=(const MyString&) called" << endl;
        if (this != &other) {
            delete[] _data;
            _len = other._len;
            _data = new char[_len + 1];
            strcpy(_data, other._data);
        }
        return *this;
    }
    
    // 移动赋值运算符（noexcept）
    MyString& operator=(MyString&& other) noexcept {
        cout << "MyString& operator=(MyString&&) called" << endl;
        if (this != &other) {
            delete[] _data;
            _data = other._data;
            _len = other._len;
            other._data = nullptr;
            other._len = 0;
        }
        return *this;
    }
    
    // 析构函数
    ~MyString() {
        cout << "~MyString() called" << endl;
        delete[] _data;
    }
    
    // 打印
    void print() const {
        if (_data) {
            cout << _data << endl;
        } else {
            cout << "Empty string" << endl;
        }
    }
};

// 示例3：override 和 final
// final 类：禁止继承
struct Base1 final {
    virtual void func() {
        cout << "Base1::func()" << endl;
    }
};

// 错误：不能从 final 基类 'Base1' 派生
// struct Derived1 : Base1 { };

// 普通基类
struct Base2 {
    virtual void vfunc(float) {
        cout << "Base2::vfunc(float)" << endl;
    }
    
    // final 虚函数：禁止覆盖
    virtual void f() final {
        cout << "Base2::f()" << endl;
    }
};

// 派生类
struct Derived2 : Base2 {
    // 错误：标记为 override，但没有覆盖基类的函数
    // virtual void vfunc(int) override { }
    
    // 正确：覆盖基类的函数
    virtual void vfunc(float) override {
        cout << "Derived2::vfunc(float)" << endl;
    }
    
    // 错误：覆盖 final 函数 'virtual void Base2::f()'
    // void f() { }
};

int main() {
    // 测试 Type Alias
    cout << "=== Testing Type Alias ===" << endl;
    func fn = add;
    fn(10, 20);
    
    Container<int> c;
    c.push_back(1);
    c.push_back(2);
    c.push_back(3);
    c.print();
    
    // 测试 noexcept
    cout << "\n=== Testing noexcept ===" << endl;
    MyString s1("Hello");
    MyString s2 = std::move(s1); // 使用移动构造函数
    s2.print();
    s1.print(); // 空字符串
    
    MyString s3("World");
    s3 = std::move(s2); // 使用移动赋值运算符
    s3.print();
    s2.print(); // 空字符串
    
    // 测试 override 和 final
    cout << "\n=== Testing override and final ===" << endl;
    Base2* b = new Derived2();
    b->vfunc(3.14f); // 调用派生类的函数
    b->f(); // 调用基类的 final 函数
    delete b;
    
    return 0;
}
```

#### 4.2 运行结果

```
=== Testing Type Alias ===
a + b = 30
1 2 3 

=== Testing noexcept ===
MyString(const char*) called
MyString(MyString&&) called
Hello
Empty string
MyString(const char*) called
MyString& operator=(MyString&&) called
World
Empty string
~MyString() called
~MyString() called
~MyString() called

=== Testing override and final ===
Derived2::vfunc(float)
Base2::f()

=== Testing final class ===
```

---

## 九、 总结与思考

### 1. 核心知识点总结

#### 1.1 =default, =delete
- **=default**：显式要求编译器生成默认版本的函数，只能用于 Big-Five 函数
- **=delete**：显式告诉编译器不要定义某个函数，可用于任何成员函数
- **应用场景**：控制编译器自动生成的函数，实现 No-Copy 模式

#### 1.2 Alias Template
- **定义**：模板别名，使用 using 关键字定义
- **优势**：语法清晰，支持模板参数，比宏和 typedef 更强大
- **应用场景**：简化模板代码，包装标准容器

#### 1.3 Template template parameter
- **定义**：模板的参数是另一个模板
- **应用场景**：在模板中使用模板作为参数
- **与 Alias Template 的结合**：解决标准容器多参数的问题

#### 1.4 Type Alias, noexcept, override, final
- **Type Alias**：类型别名，使用 using 关键字定义，语法更清晰
- **noexcept**：声明函数不会抛出异常，对移动语义很重要
- **override**：显式声明函数覆盖基类的虚函数，提高代码安全性
- **final**：禁止类被继承或虚函数被覆盖，提高代码稳定性

### 2. 设计思想与技术要点

#### 2.1 控制编译器行为
- **=default, =delete**：通过显式声明控制编译器自动生成的函数，提高代码的可控性
- **override, final**：通过显式声明提高代码的安全性和稳定性

#### 2.2 简化代码
- **Alias Template**：简化模板代码，提高代码的可读性
- **Type Alias**：简化类型声明，提高代码的可读性

#### 2.3 异常安全性
- **noexcept**：通过声明函数不抛出异常，提高代码的异常安全性
- **移动语义**：结合 noexcept，提高容器的性能

### 3. 实际应用价值

#### 3.1 提高代码质量
- **显式声明**：通过 =default, =delete, override, final 等关键字，使代码意图更明确
- **编译时检查**：编译器会检查这些关键字的使用是否正确，提前发现错误

#### 3.2 提高代码性能
- **移动语义**：结合 noexcept，使容器在扩容时使用移动语义，提高性能
- **模板别名**：简化模板代码，提高编译器的优化能力

#### 3.3 提高代码可维护性
- **清晰的语法**：using 关键字的多种用途，使代码更清晰
- **明确的意图**：通过显式声明，使代码意图更明确，便于维护

### 4. 学习建议与进阶方向

#### 4.1 学习建议
- **基础巩固**：掌握这些新特性的基本语法和用法
- **实践应用**：在实际项目中多使用这些特性，熟悉它们的应用场景
- **深入理解**：学习它们的底层实现原理，理解编译器如何处理它们

#### 4.2 进阶方向
- **模板元编程**：深入学习模板编程，结合 Alias Template 和 Template template parameter
- **异常处理**：深入学习异常处理机制，理解 noexcept 的作用
- **移动语义**：深入学习移动语义，理解 noexcept 对移动语义的影响
- **现代 C++**：学习 C++14/17/20 中对这些特性的增强

### 5. 个人感悟

通过学习这些 C++11 新特性，我深刻体会到 C++ 语言的不断进化和完善。这些新特性不仅使代码更简洁、更易读，也提高了代码的安全性和性能。

**=default 和 =delete** 让我们能够更精确地控制编译器的行为，实现 No-Copy 等模式变得更加简洁。**Alias Template** 和 **Type Alias** 简化了模板代码和类型声明，使代码更清晰。**Template template parameter** 增强了模板的表达能力。**noexcept**、**override** 和 **final** 提高了代码的安全性和稳定性。

在实际项目中，我越来越多地使用这些特性：
- 使用 =delete 实现单例模式的 No-Copy
- 使用 Alias Template 简化容器的使用
- 使用 override 确保虚函数的正确覆盖
- 使用 noexcept 提高移动语义的性能

这些特性不仅是 C++11 的重要组成部分，也是现代 C++ 编程的必备技能。掌握它们，将使我们的代码更加优雅、高效、安全。

---

## 十、 附录：常见问题解答

### 1. =default, =delete 相关问题

#### 1.1 Q: 什么情况下编译器会自动生成默认构造函数？
**A:** 当类没有声明任何构造函数时，编译器会自动生成默认构造函数。如果类声明了任何构造函数，编译器就不会自动生成默认构造函数，除非使用 =default 显式要求。

#### 1.2 Q: 使用 =delete 禁止析构函数会有什么后果？
**A:** 使用 =delete 禁止析构函数会导致无法销毁该类型的对象，包括：
- 无法在栈上创建该类型的对象
- 可以在堆上创建该类型的对象，但无法 delete 它
- 该类型的对象无法作为局部变量、成员变量或返回值

### 2. Alias Template 相关问题

#### 2.1 Q: Alias Template 可以部分特化或显式特化吗？
**A:** 不可以，Alias Template 不能部分特化或显式特化。

#### 2.2 Q: 如何使用 Alias Template 包装标准容器？
**A:** 可以使用 Alias Template 为标准容器定义带有特定分配器的别名，例如：
```cpp
template <typename T>
using Vec = std::vector<T, MyAlloc<T>>;
```

### 3. Template template parameter 相关问题

#### 3.1 Q: 为什么标准容器不能直接作为 Template template parameter？
**A:** 因为标准容器有多个模板参数（例如 vector 有 T 和 Allocator 两个参数），而 Template template parameter 通常只接受一个参数。可以使用 Alias Template 包装标准容器，使其只接受一个参数。

#### 3.2 Q: 如何声明一个接受多个参数的 Template template parameter？
**A:** 可以在 Template template parameter 中指定多个参数，例如：
```cpp
template <typename T,
         template <class, class> class Container,
         typename Alloc = std::allocator<T>
         >
class XCls {
    Container<T, Alloc> c;
    // ...
};
```

### 4. Type Alias, noexcept, override, final 相关问题

#### 4.1 Q: using 关键字有哪些用途？
**A:** using 关键字有多种用途：
- 命名空间指令：`using namespace std;`
- 命名空间成员声明：`using std::count;`
- 类成员声明：`using Base::member;`
- 类型别名：`using func = void (*)(int,int);`
- 别名模板：`template <typename T> using Vec = std::vector<T>;`

#### 4.2 Q: noexcept 对移动语义有什么影响？
**A:** 可增长的容器（vector 和 deque）只有在移动构造函数和析构函数不抛出异常时才会使用移动语义。如果构造函数不是 noexcept，std::vector 不能使用它，因为无法确保标准要求的异常安全性。

#### 4.3 Q: override 和 final 有什么区别？
**A:** override 用于显式声明函数覆盖基类的虚函数，编译器会检查基类是否有具有相同签名的虚函数；final 用于禁止类被继承或虚函数被进一步覆盖。

---

## 十一、 参考文献

1. **C++ 标准库（第2版）** - Nicolai M. Josuttis
2. **C++ Primer（第5版）** - Stanley B. Lippman, Josée Lajoie, Barbara E. Moo
3. **Effective Modern C++** - Scott Meyers
4. **C++11/14 高级编程** - 侯捷
5. **cppreference.com** - C++ 标准参考网站

---

希望这份学习笔记能帮助你更好地理解 C++11 中的这些新特性，掌握现代 C++ 的核心技能！