# 侯捷 C++ STL标准库和泛型编程 - 仿函数（Functors）与可适配条件深度解析
> **打卡日期**：2026-04-12
> **核心主题**：仿函数（Functors）、可适配条件、STL六大部件关系
> **资料出处**：https://www.bilibili.com/video/BV1kBh5zrEYh?spm_id_from=333.788.videopod.sections&vd_source=e01209adaee5cf8495fa73fbfde26dd1

---

### 第31-32集核心知识点总结

#### 第31集：仿函数（Functors）的基本概念
- 仿函数的定义和特性
- 仿函数的分类：算术类、逻辑运算类、相对关系类
- 仿函数与STL算法的配合使用

#### 第32集：存在多种Adapters
- 适配器的种类：容器适配器、迭代器适配器、仿函数适配器
- 容器适配器的实现：stack、queue
- 适配器与其他STL部件的交互关系

---

## 1. 仿函数（Functors）的基本概念

### 1.1 什么是仿函数
仿函数（Functors）是一种**行为类似函数的对象**，它是一个类，通过重载`operator()`运算符，使得该类的实例可以像函数一样被调用。

```cpp
template <class T>
struct less : public binary_function<T, T, bool> {
    bool operator()(const T& x, const T& y) const {
        return x < y;
    }
};
```

### 1.2 仿函数的优势
- **状态保持**：仿函数可以拥有成员变量，从而在多次调用之间保持状态
- **可适配性**：通过继承特定的基类，可以被函数适配器（Function Adapters）修改和组合
- **编译期优化**：编译器可以内联仿函数的调用，提高性能
- **类型安全**：相比函数指针，仿函数提供了更好的类型安全性

### 1.3 仿函数的分类

#### 1.3.1 算术类（Arithmetic）
| 仿函数 | 功能 | 基类 |
|--------|------|------|
| `plus<T>` | 加法 | `binary_function<T, T, T>` |
| `minus<T>` | 减法 | `binary_function<T, T, T>` |
| `multiplies<T>` | 乘法 | `binary_function<T, T, T>` |
| `divides<T>` | 除法 | `binary_function<T, T, T>` |
| `modulus<T>` | 取模 | `binary_function<T, T, T>` |
| `negate<T>` | 取反 | `unary_function<T, T>` |

#### 1.3.2 逻辑运算类（Logical）
| 仿函数 | 功能 | 基类 |
|--------|------|------|
| `logical_and<T>` | 逻辑与 | `binary_function<T, T, bool>` |
| `logical_or<T>` | 逻辑或 | `binary_function<T, T, bool>` |
| `logical_not<T>` | 逻辑非 | `unary_function<T, bool>` |

#### 1.3.3 相对关系类（Relational）
| 仿函数 | 功能 | 基类 |
|--------|------|------|
| `equal_to<T>` | 等于 | `binary_function<T, T, bool>` |
| `not_equal_to<T>` | 不等于 | `binary_function<T, T, bool>` |
| `greater<T>` | 大于 | `binary_function<T, T, bool>` |
| `less<T>` | 小于 | `binary_function<T, T, bool>` |
| `greater_equal<T>` | 大于等于 | `binary_function<T, T, bool>` |
| `less_equal<T>` | 小于等于 | `binary_function<T, T, bool>` |

## 2. STL六大部件的关系

### 2.1 六大部件概述
STL由六大部件组成，它们相互协作，形成一个完整的体系：

| 部件 | 类型 | 作用 | 示例 |
|------|------|------|------|
| **容器（Containers）** | 类模板 | 存储元素 | vector, list, map |
| **算法（Algorithms）** | 函数模板 | 操作元素 | sort, find, copy |
| **迭代器（Iterators）** | 类模板 | 连接容器和算法 | vector::iterator |
| **仿函数（Functors）** | 类模板 | 提供操作逻辑 | less, greater |
| **适配器（Adapters）** | 类模板 | 修改其他组件的接口 | stack, queue, function adapters |
| **分配器（Allocators）** | 类模板 | 管理内存分配 | std::allocator |

### 2.2 适配器的种类与作用
STL中的适配器主要分为三大类：

| 适配器类型 | 作用 | 示例 |
|-----------|------|------|
| **容器适配器（Container Adapters）** | 封装底层容器，提供特定接口 | stack, queue, priority_queue |
| **迭代器适配器（Iterator Adapters）** | 修改迭代器的行为 | reverse_iterator, insert_iterator |
| **仿函数适配器（Functor Adapters）** | 修改仿函数的行为 | bind1st, bind2nd, not1, not2 |

### 2.3 仿函数在STL中的位置
仿函数在STL中扮演着重要的角色，它们：
- 为算法提供自定义的操作逻辑
- 可以被函数适配器修改和组合
- 与迭代器一起，使算法能够适应不同的容器类型
- 提供了一种灵活的方式来定制算法的行为

### 2.4 部件间的交互关系
STL各部件之间通过特定的方式交互：
- **算法与迭代器**：算法通过迭代器操作容器中的元素，算法会向迭代器提问（如迭代器类别、值类型等），迭代器必须回答这些问题
- **算法与仿函数**：算法使用仿函数定义操作逻辑，仿函数通过可适配条件提供必要的类型信息
- **适配器与其他部件**：适配器修改其他部件的接口，使其适应特定的使用场景

这种设计使得STL具有高度的灵活性和可扩展性，各个部件可以独立演化，同时保持良好的协作关系。

## 3. 仿函数的可适配（Adaptable）条件

### 3.1 什么是可适配仿函数
可适配仿函数是指**继承自特定基类**的仿函数，这些基类定义了一些typedef，使得仿函数可以被函数适配器（Function Adapters）操作。

### 3.2 基类定义

#### 3.2.1 `unary_function`
```cpp
template <class Arg, class Result>
struct unary_function {
    typedef Arg argument_type;
    typedef Result result_type;
};
```

#### 3.2.2 `binary_function`
```cpp
template <class Arg1, class Arg2, class Result>
struct binary_function {
    typedef Arg1 first_argument_type;
    typedef Arg2 second_argument_type;
    typedef Result result_type;
};
```

### 3.3 为什么需要可适配条件
函数适配器（如`bind1st`、`bind2nd`、`not1`、`not2`等）需要知道仿函数的参数类型和返回类型，以便正确地修改和组合仿函数。通过继承`unary_function`或`binary_function`，仿函数提供了这些必要的类型信息。

例如，对于`less<int>`：
```cpp
template <class T>
struct less : public binary_function<T, T, bool> {
    bool operator()(const T& x, const T& y) const {
        return x < y;
    }
};
```

当我们使用`less<int>`时，它会继承三个typedef：
- `first_argument_type`：int
- `second_argument_type`：int
- `result_type`：bool

这些typedef使得函数适配器可以正确地操作`less<int>`。

## 4. 仿函数的使用示例

### 4.1 基本使用

```cpp
#include <iostream>
#include <vector>
#include <algorithm>
#include <functional>

int main() {
    std::vector<int> vec = {3, 1, 4, 1, 5, 9, 2, 6};
    
    // 使用默认比较（operator<）
    std::sort(vec.begin(), vec.end());
    
    // 使用函数作为比较器
    bool myfunc(int i, int j) { return i > j; }
    std::sort(vec.begin(), vec.end(), myfunc);
    
    // 使用对象作为比较器
    struct myclass {
        bool operator()(int i, int j) { return i < j; }
    } myobj;
    std::sort(vec.begin(), vec.end(), myobj);
    
    // 显式使用默认比较（operator<）
    std::sort(vec.begin(), vec.end(), std::less<int>());
    
    // 使用另一种比较标准（operator>）
    std::sort(vec.begin(), vec.end(), std::greater<int>());
    
    return 0;
}
```

### 4.2 与算法配合使用

```cpp
#include <iostream>
#include <vector>
#include <algorithm>
#include <functional>

int main() {
    std::vector<int> vec = {1, 2, 3, 4, 5};
    
    // 使用plus仿函数进行累加
    int sum = std::accumulate(vec.begin(), vec.end(), 0, std::plus<int>());
    std::cout << "Sum: " << sum << std::endl;
    
    // 使用equal_to仿函数查找元素
    auto it = std::find_if(vec.begin(), vec.end(), 
                          std::bind2nd(std::equal_to<int>(), 3));
    if (it != vec.end()) {
        std::cout << "Found: " << *it << std::endl;
    }
    
    // 使用greater仿函数筛选元素
    std::vector<int> result;
    std::copy_if(vec.begin(), vec.end(), 
                std::back_inserter(result),
                std::bind2nd(std::greater<int>(), 2));
    
    std::cout << "Elements greater than 2: ";
    for (int num : result) {
        std::cout << num << " ";
    }
    std::cout << std::endl;
    
    return 0;
}
```

## 5. 仿函数的实现细节

### 5.1 自定义仿函数

```cpp
// 自定义仿函数，计算平方
template <class T>
struct square : public unary_function<T, T> {
    T operator()(const T& x) const {
        return x * x;
    }
};

// 自定义二元仿函数，计算两数之差的绝对值
template <class T>
struct abs_difference : public binary_function<T, T, T> {
    T operator()(const T& x, const T& y) const {
        return x > y ? x - y : y - x;
    }
};
```

### 5.2 仿函数的状态

```cpp
// 带状态的仿函数，计数调用次数
template <class T>
struct counting_less : public binary_function<T, T, bool> {
    mutable int count; // 使用mutable允许在const成员函数中修改
    
    counting_less() : count(0) {}
    
    bool operator()(const T& x, const T& y) const {
        ++count;
        return x < y;
    }
};

int main() {
    std::vector<int> vec = {3, 1, 4, 1, 5};
    counting_less<int> comp;
    
    std::sort(vec.begin(), vec.end(), comp);
    std::cout << "Number of comparisons: " << comp.count << std::endl;
    
    return 0;
}
```

## 6. 适配器的详细实现

### 6.1 容器适配器

#### 6.1.1 stack的实现
```cpp
template <class T, class Sequence=deque<T>>
class stack {
public:
    typedef typename Sequence::value_type value_type;
    typedef typename Sequence::size_type size_type;
    typedef typename Sequence::reference reference;
    typedef typename Sequence::const_reference const_reference;
protected:
    Sequence c; // 底层容器
public:
    bool empty() const { return c.empty(); }
    size_type size() const { return c.size(); }
    reference top() { return c.back(); }
    const_reference top() const { return c.back(); }
    void push(const value_type& x) { c.push_back(x); }
    void pop() { c.pop_back(); }
};
```

#### 6.1.2 queue的实现
```cpp
template <class T, class Sequence=deque<T>>
class queue {
public:
    typedef typename Sequence::value_type value_type;
    typedef typename Sequence::size_type size_type;
    typedef typename Sequence::reference reference;
    typedef typename Sequence::const_reference const_reference;
protected:
    Sequence c; // 底层容器
public:
    bool empty() const { return c.empty(); }
    size_type size() const { return c.size(); }
    reference front() { return c.front(); }
    const_reference front() const { return c.front(); }
    reference back() { return c.back(); }
    const_reference back() const { return c.back(); }
    void push(const value_type& x) { c.push_back(x); }
    void pop() { c.pop_front(); }
};
```

### 6.2 函数适配器

| 适配器 | 功能 | 示例 |
|--------|------|------|
| `bind1st` | 将二元仿函数的第一个参数绑定为固定值 | `bind1st(less<int>(), 5)` |
| `bind2nd` | 将二元仿函数的第二个参数绑定为固定值 | `bind2nd(less<int>(), 5)` |
| `not1` | 对一元仿函数取反 | `not1(bind2nd(less<int>(), 5))` |
| `not2` | 对二元仿函数取反 | `not2(less<int>())` |
| `ptr_fun` | 将函数指针转换为仿函数 | `ptr_fun(myfunc)` |
| `mem_fun` | 将成员函数转换为仿函数 | `mem_fun(&MyClass::method)` |
| `mem_fun_ref` | 将成员函数转换为引用仿函数 | `mem_fun_ref(&MyClass::method)` |

### 6.2.1 迭代器适配器

| 适配器 | 功能 | 示例 |
|--------|------|------|
| `reverse_iterator` | 反向迭代器，从尾到头遍历 | `vector<int>::reverse_iterator` |
| `insert_iterator` | 插入迭代器，将元素插入容器 | `insert_iterator<vector<int>>` |
| `back_insert_iterator` | 尾部插入迭代器 | `back_inserter(vec)` |
| `front_insert_iterator` | 头部插入迭代器 | `front_inserter(deq)` |
| `ostream_iterator` | 输出流迭代器 | `ostream_iterator<int>(cout, " ")` |
| `istream_iterator` | 输入流迭代器 | `istream_iterator<int>(cin)` |

### 6.3 适配器使用示例

#### 6.3.1 函数适配器示例

```cpp
#include <iostream>
#include <vector>
#include <algorithm>
#include <functional>

class Person {
public:
    Person(std::string name, int age) : name(name), age(age) {}
    
    std::string name;
    int age;
    
    void print() const {
        std::cout << name << ", " << age << std::endl;
    }
};

int main() {
    std::vector<Person> people = {
        Person("Alice", 30),
        Person("Bob", 25),
        Person("Charlie", 35)
    };
    
    // 使用mem_fun_ref调用成员函数
    std::for_each(people.begin(), people.end(), 
                  std::mem_fun_ref(&Person::print));
    
    // 使用bind2nd和less查找年龄小于30的人
    auto it = std::find_if(people.begin(), people.end(),
                          std::bind2nd(std::less<int>(), 30));
    if (it != people.end()) {
        std::cout << "First person under 30: " << it->name << std::endl;
    }
    
    // 使用not1和bind2nd查找年龄不小于30的人
    auto it2 = std::find_if(people.begin(), people.end(),
                           std::not1(std::bind2nd(std::less<int>(), 30)));
    if (it2 != people.end()) {
        std::cout << "First person 30 or older: " << it2->name << std::endl;
    }
    
    return 0;
}
```

#### 6.3.2 容器适配器示例

```cpp
#include <iostream>
#include <stack>
#include <queue>
#include <vector>

int main() {
    // 使用stack适配器
    std::stack<int> s;
    s.push(1);
    s.push(2);
    s.push(3);
    
    std::cout << "Stack elements: ";
    while (!s.empty()) {
        std::cout << s.top() << " ";
        s.pop();
    }
    std::cout << std::endl;
    
    // 使用queue适配器
    std::queue<int> q;
    q.push(1);
    q.push(2);
    q.push(3);
    
    std::cout << "Queue elements: ";
    while (!q.empty()) {
        std::cout << q.front() << " ";
        q.pop();
    }
    std::cout << std::endl;
    
    // 使用priority_queue适配器
    std::priority_queue<int> pq;
    pq.push(3);
    pq.push(1);
    pq.push(4);
    pq.push(1);
    pq.push(5);
    
    std::cout << "Priority queue elements: ";
    while (!pq.empty()) {
        std::cout << pq.top() << " ";
        pq.pop();
    }
    std::cout << std::endl;
    
    return 0;
}
```

#### 6.3.3 迭代器适配器示例

```cpp
#include <iostream>
#include <vector>
#include <algorithm>
#include <iterator>

int main() {
    std::vector<int> vec = {1, 2, 3, 4, 5};
    
    // 使用反向迭代器
    std::cout << "Reverse order: ";
    for (auto it = vec.rbegin(); it != vec.rend(); ++it) {
        std::cout << *it << " ";
    }
    std::cout << std::endl;
    
    // 使用输出流迭代器
    std::cout << "Using ostream_iterator: ";
    std::copy(vec.begin(), vec.end(), 
              std::ostream_iterator<int>(std::cout, " "));
    std::cout << std::endl;
    
    // 使用back_insert_iterator
    std::vector<int> dest;
    std::copy(vec.begin(), vec.end(), 
              std::back_inserter(dest));
    
    std::cout << "Copied to dest: ";
    for (int num : dest) {
        std::cout << num << " ";
    }
    std::cout << std::endl;
    
    return 0;
}
```

## 7. C++11及以后的仿函数发展

### 7.1 Lambda表达式
C++11引入了Lambda表达式，它提供了一种更简洁的方式来定义匿名函数对象：

```cpp
#include <iostream>
#include <vector>
#include <algorithm>

int main() {
    std::vector<int> vec = {3, 1, 4, 1, 5, 9, 2, 6};
    
    // 使用Lambda表达式作为比较器
    std::sort(vec.begin(), vec.end(), 
              [](int a, int b) { return a > b; });
    
    // 使用Lambda表达式进行筛选
    std::vector<int> result;
    std::copy_if(vec.begin(), vec.end(), 
                std::back_inserter(result),
                [](int x) { return x % 2 == 0; });
    
    return 0;
}
```

### 7.2 `std::function`和`std::bind`
C++11引入了`std::function`和`std::bind`，它们提供了更灵活的函数对象管理方式：

```cpp
#include <iostream>
#include <functional>

int add(int a, int b) {
    return a + b;
}

class Calculator {
public:
    int multiply(int a, int b) {
        return a * b;
    }
};

int main() {
    // 使用std::function存储函数
    std::function<int(int, int)> func = add;
    std::cout << "5 + 3 = " << func(5, 3) << std::endl;
    
    // 使用std::bind绑定参数
    auto add5 = std::bind(add, 5, std::placeholders::_1);
    std::cout << "5 + 7 = " << add5(7) << std::endl;
    
    // 绑定成员函数
    Calculator calc;
    auto multiply_by_2 = std::bind(&Calculator::multiply, &calc, 2, std::placeholders::_1);
    std::cout << "2 * 8 = " << multiply_by_2(8) << std::endl;
    
    return 0;
}
```

## 8. 仿函数的性能考量

### 8.1 内联优化
仿函数的一个重要优势是编译器可以内联它们的调用，因为仿函数的类型在编译时是已知的。相比之下，函数指针的调用通常不能被内联，因为它们的目标在运行时才能确定。

### 8.2 编译期计算
仿函数可以与模板元编程结合，实现编译期计算：

```cpp
// 编译期计算阶乘
template <int N>
struct Factorial {
    static const int value = N * Factorial<N-1>::value;
};

template <>
struct Factorial<0> {
    static const int value = 1;
};

// 使用编译期计算
int main() {
    constexpr int fact5 = Factorial<5>::value; // 编译期计算，结果为120
    std::cout << "5! = " << fact5 << std::endl;
    return 0;
}
```

## 9. 仿函数的实际应用场景

### 9.1 排序和查找
仿函数常用于自定义排序和查找的比较逻辑：

```cpp
// 按字符串长度排序
std::sort(strings.begin(), strings.end(), 
          [](const std::string& a, const std::string& b) {
              return a.size() < b.size();
          });

// 查找大于100的元素
auto it = std::find_if(numbers.begin(), numbers.end(), 
                       std::bind2nd(std::greater<int>(), 100));
```

### 9.2 算法定制
仿函数可以定制算法的行为，例如：

```cpp
// 自定义变换操作
std::transform(numbers.begin(), numbers.end(), 
               std::back_inserter(results),
               [](int x) { return x * x + 1; });

// 自定义累积操作
int product = std::accumulate(numbers.begin(), numbers.end(), 1, 
                              std::multiplies<int>());
```

### 9.3 事件处理
仿函数可以用于事件处理系统，存储回调函数：

```cpp
class EventSystem {
public:
    using EventHandler = std::function<void()>;
    
    void registerHandler(EventHandler handler) {
        handlers.push_back(handler);
    }
    
    void triggerEvent() {
        for (auto& handler : handlers) {
            handler();
        }
    }
    
private:
    std::vector<EventHandler> handlers;
};
```

## 10. 总结与最佳实践

### 10.1 仿函数的优缺点

#### 优点
- **灵活性**：可以根据需要定制操作逻辑
- **可适配性**：可以与函数适配器配合使用
- **性能**：编译器可以内联调用，提高性能
- **状态管理**：可以保持状态，实现更复杂的逻辑

#### 缺点
- **代码复杂性**：相比普通函数，代码可能更复杂
- **可读性**：对于不熟悉STL的开发者，可能难以理解
- **C++11+的替代方案**：Lambda表达式和std::function提供了更简洁的方式

### 10.2 适配器的设计原则

#### 设计理念
- **封装**：适配器封装底层组件，提供更简洁的接口
- **复用**：适配器复用现有的容器、迭代器或仿函数，避免重复实现
- **接口转换**：适配器将一种接口转换为另一种接口，以满足特定需求
- **组合**：适配器通过组合而非继承的方式使用底层组件

#### 最佳实践
- **选择合适的底层容器**：对于容器适配器，选择适合特定场景的底层容器
- **合理使用适配器**：只在需要修改接口时使用适配器，避免过度使用
- **理解适配器的性能影响**：某些适配器可能会引入额外的性能开销
- **考虑C++11+的替代方案**：对于函数适配器，考虑使用`std::bind`和Lambda表达式

### 10.3 仿函数的最佳实践

1. **优先使用标准仿函数**：STL提供了许多常用的仿函数，应优先使用它们
2. **合理使用继承**：为了使仿函数可适配，应继承适当的基类
3. **考虑Lambda表达式**：在C++11及以后的代码中，考虑使用Lambda表达式替代复杂的仿函数
4. **注意性能**：对于性能敏感的场景，仿函数可能比函数指针更优
5. **保持简洁**：仿函数的实现应该简洁明了，避免过度复杂

### 10.4 学习建议
- 理解仿函数的基本概念和工作原理
- 熟悉STL中常用的仿函数
- 掌握函数适配器的使用方法
- 了解C++11及以后的替代方案
- 通过实际项目练习，加深对仿函数的理解

---

## 面试高频考点

### 1. 仿函数与函数指针的区别
**问题**：请比较仿函数与函数指针的异同。

**回答思路**：
- 相同点：都可以作为函数参数传递，都可以被调用
- 不同点：
  - 仿函数是类的实例，函数指针是指针
  - 仿函数可以保持状态，函数指针不能
  - 仿函数可以被内联，函数指针通常不能
  - 仿函数提供更好的类型安全性
  - 仿函数可以被适配和组合

### 2. 可适配仿函数的条件
**问题**：什么是可适配仿函数？它需要满足什么条件？

**回答思路**：
- 可适配仿函数是指继承自`unary_function`或`binary_function`的仿函数
- 这些基类提供了必要的typedef，如`argument_type`、`result_type`等
- 这些typedef使得函数适配器可以正确地操作仿函数
- 例如，`less<T>`继承自`binary_function<T, T, bool>`，因此可以被`bind2nd`等适配器操作

### 3. 仿函数与Lambda表达式的比较
**问题**：在C++11中，Lambda表达式与仿函数相比有什么优势？

**回答思路**：
- Lambda表达式更简洁，不需要定义单独的类
- Lambda表达式可以捕获外部变量，更灵活
- Lambda表达式的类型是匿名的，需要使用`auto`或`std::function`来存储
- 仿函数可以有更复杂的状态管理和成员函数
- 对于简单的操作，Lambda表达式通常更合适；对于复杂的、可重用的操作，仿函数可能更合适

### 4. STL六大部件的关系
**问题**：请简述STL六大部件的关系。

**回答思路**：
- 容器存储元素，算法操作元素，迭代器连接容器和算法
- 仿函数提供操作逻辑，适配器修改其他组件的接口，分配器管理内存
- 六大部件相互协作，形成一个完整的体系
- 例如，算法通过迭代器操作容器中的元素，使用仿函数定义操作逻辑，通过适配器调整仿函数的行为

### 5. 仿函数的性能优势
**问题**：仿函数相比普通函数有什么性能优势？

**回答思路**：
- 编译器可以内联仿函数的调用，因为仿函数的类型在编译时是已知的
- 函数指针的调用通常不能被内联，因为它们的目标在运行时才能确定
- 仿函数可以与模板元编程结合，实现编译期计算
- 仿函数的调用开销通常比函数指针小

### 6. 容器适配器的实现原理
**问题**：stack和queue的底层实现是什么？它们有什么特点？

**回答思路**：
- stack和queue都是容器适配器，默认使用deque作为底层容器
- stack实现了后进先出（LIFO）的数据结构，提供push、pop、top等操作
- queue实现了先进先出（FIFO）的数据结构，提供push、pop、front、back等操作
- 容器适配器通过组合的方式使用底层容器，而非继承
- 可以通过模板参数指定不同的底层容器，只要该容器提供必要的操作

### 7. 迭代器适配器的作用
**问题**：什么是迭代器适配器？请举例说明。

**回答思路**：
- 迭代器适配器是修改现有迭代器行为的组件
- 例如，reverse_iterator可以反向遍历容器
- insert_iterator可以将元素插入容器，而不是覆盖
- ostream_iterator可以将元素输出到流中
- 迭代器适配器扩展了迭代器的功能，使算法更加灵活

---

## 代码优化建议

### 1. 优先使用标准仿函数
```cpp
// 不推荐
struct MyLess {
    bool operator()(int a, int b) {
        return a < b;
    }
};
std::sort(vec.begin(), vec.end(), MyLess());

// 推荐
std::sort(vec.begin(), vec.end(), std::less<int>());
```

### 2. 合理使用Lambda表达式
```cpp
// 对于简单操作，使用Lambda表达式
std::sort(vec.begin(), vec.end(), 
          [](int a, int b) { return a > b; });

// 对于复杂操作，考虑使用命名仿函数
struct ComplexComparator {
    bool operator()(const MyType& a, const MyType& b) {
        // 复杂的比较逻辑
    }
};
std::sort(vec.begin(), vec.end(), ComplexComparator());
```

### 3. 注意仿函数的可适配性
```cpp
// 不推荐：不可适配
struct MyComparator {
    bool operator()(int a, int b) {
        return a < b;
    }
};

// 推荐：可适配
struct MyAdaptableComparator : public std::binary_function<int, int, bool> {
    bool operator()(int a, int b) const {
        return a < b;
    }
};
```

### 4. 避免不必要的状态
```cpp
// 不推荐：不必要的状态
struct CountingComparator {
    int count = 0;
    bool operator()(int a, int b) {
        ++count;
        return a < b;
    }
};

// 推荐：只在需要时使用状态
struct CountingComparator {
    mutable int count = 0; // 使用mutable
    bool operator()(int a, int b) const {
        ++count;
        return a < b;
    }
};
```

---

## 总结

仿函数是STL中的重要组成部分，它提供了一种灵活的方式来定制算法的行为。通过重载`operator()`运算符，仿函数可以像函数一样被调用，同时还可以保持状态和被适配。

在C++11及以后的版本中，Lambda表达式和`std::function`提供了更简洁的替代方案，但仿函数仍然在许多场景中发挥着重要作用，特别是在需要可重用的、有状态的操作逻辑时。

理解仿函数的工作原理和使用方法，对于掌握STL和C++泛型编程至关重要。通过合理使用仿函数，可以编写更灵活、更高效的代码。