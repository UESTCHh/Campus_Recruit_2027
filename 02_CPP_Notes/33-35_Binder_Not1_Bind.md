# 侯捷 C++ STL标准库和泛型编程 - 函数适配器（Binder2nd、not1、bind）深度解析
> **打卡日期**：2026-04-13
> **核心主题**：函数适配器、Binder2nd、not1、C++11 bind
> **资料出处**：https://www.bilibili.com/video/BV1kBh5zrEYh?spm_id_from=333.788.videopod.sections&vd_source=e01209adaee5cf8495fa73fbfde26dd1

---

### 第33-35集核心知识点总结

#### 第33集：Binder2nd
- Binder2nd的定义和实现原理
- Binder2nd如何将二元函数转换为一元函数
- Binder2nd的使用场景和示例

#### 第34集：not1
- not1的定义和实现原理
- not1如何对谓词进行逻辑取反
- not1与其他函数适配器的配合使用

#### 第35集：bind（C++11）
- C++11中bind的引入和优势
- bind的使用方法和语法
- bind与传统绑定器的对比
- 占位符的使用和原理

---

## 1. Binder2nd适配器

### 1.1 Binder2nd的定义
Binder2nd是一种函数适配器，它的作用是**将一个二元函数对象转换为一元函数对象**，通过绑定第二个参数为固定值。

```cpp
// 辅助函数，让用户可以方便使用binder2nd<Op>
template <class Operation, class T>
inline binder2nd<Operation> bind2nd(const Operation& op, const T& x) {
    typedef typename Operation::second_argument_type arg2_type;
    return binder2nd<Operation>(op, arg2_type(x));
}

// 以下将某个Adaptable Binary function 转换为 Unary Function
template <class Operation>
class binder2nd 
    : public unary_function<typename Operation::first_argument_type,
                           typename Operation::result_type> {
protected:
    Operation op; // 内部成员，分别用以记录算式和第二实参
    typename Operation::second_argument_type value;
public:
    // constructor
    binder2nd(const Operation& x, 
              const typename Operation::second_argument_type& y)
        : op(x), value(y) { } // 将算式和第二实参记录下来
    
    typename Operation::result_type
    operator()(const typename Operation::first_argument_type& x) const {
        return op(x, value); // 实际呼叫算式并取value为第二实参
    }
};
```

### 1.2 Binder2nd的工作原理
1. **构造阶段**：保存原始的二元函数对象和要绑定的第二个参数值
2. **调用阶段**：当调用转换后的一元函数对象时，它会将传入的参数作为第一个参数，与保存的第二个参数一起传递给原始的二元函数

### 1.3 Binder2nd的使用示例

```cpp
#include <iostream>
#include <vector>
#include <algorithm>
#include <functional>

int main() {
    std::vector<int> vec = {10, 20, 30, 40, 50};
    
    // 统计大于30的元素个数
    int count = std::count_if(vec.begin(), vec.end(), 
                             std::bind2nd(std::greater<int>(), 30));
    
    std::cout << "Numbers greater than 30: " << count << std::endl;
    // 输出: Numbers greater than 30: 2
    
    return 0;
}
```

## 2. not1适配器

### 2.1 not1的定义
not1是一种函数适配器，它的作用是**对一元谓词进行逻辑取反**，将返回true的谓词转换为返回false，反之亦然。

```cpp
// 辅助函式，使user得以方便使用unary_negate<Pred>
template <class Predicate>
inline unary_negate<Predicate> not1(const Predicate& pred) {
    return unary_negate<Predicate>(pred);
}

// 以下取某个Adaptable Predicate的逻辑负值（logical negation）
template <class Predicate>
class unary_negate 
    : public unary_function<typename Predicate::argument_type, bool> {
protected:
    Predicate pred; // 内部成员
public:
    // constructor
    explicit unary_negate(const Predicate& x) : pred(x) {}
    
    bool operator()(const typename Predicate::argument_type& x) const {
        return !pred(x); // 将pred的运算结果"取否"(negate)
    }
};
```

### 2.2 not1的工作原理
1. **构造阶段**：保存原始的一元谓词
2. **调用阶段**：当调用转换后的谓词时，它会先调用原始谓词，然后对结果取反

### 2.3 not1的使用示例

```cpp
#include <iostream>
#include <vector>
#include <algorithm>
#include <functional>

int main() {
    std::vector<int> vec = {10, 20, 30, 40, 50};
    
    // 统计不大于30的元素个数（即小于等于30）
    int count = std::count_if(vec.begin(), vec.end(), 
                             std::not1(std::bind2nd(std::greater<int>(), 30)));
    
    std::cout << "Numbers not greater than 30: " << count << std::endl;
    // 输出: Numbers not greater than 30: 3
    
    return 0;
}
```

## 3. C++11的bind适配器

### 3.1 bind的引入
C++11引入了`std::bind`，它是一种更通用、更灵活的函数适配器，**可以绑定任意函数对象的任意参数**，而不仅限于二元函数的第二个参数。

### 3.2 bind的基本语法

```cpp
#include <functional>

// 绑定函数对象
auto bound_func = std::bind(func, arg1, arg2, ...);

// 使用占位符
using namespace std::placeholders;
auto bound_func = std::bind(func, _1, _2, ...);
```

### 3.3 bind的功能
`std::bind`可以绑定：
1. 普通函数
2. 函数对象
3. 成员函数（需要提供对象地址）
4. 数据成员（需要提供对象地址）

### 3.4 bind的使用示例

```cpp
#include <iostream>
#include <functional>

// 普通函数
double mydivide(double x, double y) {
    return x / y;
}

// 结构体
struct MyPair {
    double a, b;
    double multiply() { return a * b; }
};

int main() {
    using namespace std::placeholders; // 使_1, _2, _3等可见
    
    // 绑定普通函数的所有参数
    auto fn_five = std::bind(mydivide, 10, 2);
    std::cout << fn_five() << std::endl; // 输出: 5
    
    // 绑定部分参数，使用占位符
    auto fn_half = std::bind(mydivide, _1, 2);
    std::cout << fn_half(10) << std::endl; // 输出: 5
    
    // 交换参数顺序
    auto fn_invert = std::bind(mydivide, _2, _1);
    std::cout << fn_invert(10, 2) << std::endl; // 输出: 0.2
    
    // 绑定成员函数
    MyPair ten_two = {10, 2};
    auto bound_memfn = std::bind(&MyPair::multiply, _1);
    std::cout << bound_memfn(ten_two) << std::endl; // 输出: 20
    
    // 绑定数据成员
    auto bound_memdata = std::bind(&MyPair::a, ten_two);
    std::cout << bound_memdata() << std::endl; // 输出: 10
    
    return 0;
}
```

## 4. 传统绑定器与C++11 bind的对比

| 特性 | 传统绑定器（bind1st/bind2nd） | C++11 bind |
|------|-------------------------------|------------|
| 适用范围 | 仅适用于二元函数 | 适用于任意函数 |
| 绑定位置 | 只能绑定第一个或第二个参数 | 可以绑定任意位置的参数 |
| 参数顺序 | 不能改变参数顺序 | 可以通过占位符改变参数顺序 |
| 返回类型 | 固定为unary_function | 可以返回任意可调用对象 |
| 灵活性 | 较低 | 非常高 |
| 可读性 | 较差 | 较好 |

## 5. 深度解析：bind中的占位符机制

### 5.1 占位符的定义和使用
占位符是`std::placeholders`命名空间中的对象，用于表示`std::bind`中未绑定的参数位置。常用的占位符有`_1`、`_2`、`_3`等，分别表示调用时的第一个、第二个、第三个参数。

```cpp
using namespace std::placeholders;

// _1表示调用时的第一个参数
auto fn = std::bind(myfunc, _1, 固定值);
```

### 5.2 占位符的工作原理
当使用`std::bind`创建一个绑定函数时：
1. 它会保存原始函数和所有绑定的参数（包括占位符）
2. 当调用绑定函数时，它会将传入的参数按照占位符的顺序传递给原始函数
3. 占位符`_1`、`_2`等会被替换为调用时的对应位置的参数

### 5.3 问题解答：bind方法中占位符_1的使用

**问题**：在最后测试程序的两行代码中，当使用bind方法时出现的占位符_1，为何在实际调用时没有传递参数，而是直接使用了fn_的原因。

**解答**：

这个问题涉及到`std::bind`的工作机制和函数对象的特性。让我们通过一个具体的例子来分析：

```cpp
// 假设有以下代码
using namespace std::placeholders;
auto fn_ = std::bind(less<int>(), _1, 59);
// 然后直接调用 fn_() 而没有传递参数
```

**分析**：

1. **bind的返回值**：`std::bind`返回一个函数对象，这个函数对象存储了原始函数（这里是`less<int>()`）和绑定的参数（这里是`_1`和`59`）。

2. **函数对象的operator()**：当调用`fn_()`时，实际上是调用了这个函数对象的`operator()`方法。

3. **参数传递**：在这个例子中，`fn_`被定义为接受一个参数（因为使用了`_1`占位符），所以正确的调用应该是`fn_(x)`，其中`x`会被传递给`less<int>()`作为第一个参数。

4. **为什么可以直接调用fn_()**：
   - 如果`fn_`被定义为不需要参数（即没有使用任何占位符），那么可以直接调用`fn_()`
   - 如果`fn_`被定义为需要参数但调用时没有提供，这会导致编译错误

5. **可能的误解**：
   - 可能视频中的例子有上下文，`fn_`实际上是被用作某个算法的谓词参数
   - 例如，在`count_if`算法中：
     ```cpp
     std::count_if(vec.begin(), vec.end(), fn_);
     ```
     这里算法会自动为`fn_`传递容器中的元素作为参数，所以看起来像是没有传递参数直接调用了`fn_`

### 5.4 占位符的高级用法

#### 5.4.1 绑定成员函数

```cpp
struct MyClass {
    void print(int x, int y) {
        std::cout << "x: " << x << ", y: " << y << std::endl;
    }
};

MyClass obj;
auto bound_memfn = std::bind(&MyClass::print, &obj, _1, _2);
bound_memfn(10, 20); // 输出: x: 10, y: 20
```

#### 5.4.2 组合使用多个占位符

```cpp
auto fn = std::bind(myfunc, _2, _1); // 交换参数顺序
fn(10, 20); // 相当于 myfunc(20, 10)
```

#### 5.4.3 绑定部分参数

```cpp
auto fn = std::bind(myfunc, 10, _1, 30); // 绑定第一个和第三个参数
fn(20); // 相当于 myfunc(10, 20, 30)
```

## 6. 函数适配器的实际应用

### 6.1 在算法中的应用

```cpp
#include <iostream>
#include <vector>
#include <algorithm>
#include <functional>

int main() {
    std::vector<int> vec = {15, 37, 94, 50, 73, 58, 28, 98};
    
    // 使用传统绑定器
    int n1 = std::count_if(vec.begin(), vec.end(), 
                          std::not1(std::bind2nd(std::less<int>(), 50)));
    std::cout << "Numbers not less than 50: " << n1 << std::endl;
    
    // 使用C++11 bind
    using namespace std::placeholders;
    int n2 = std::count_if(vec.begin(), vec.end(), 
                          std::bind(std::less<int>(), 59, _1));
    std::cout << "Numbers greater than 59: " << n2 << std::endl;
    
    return 0;
}
```

### 6.2 在事件处理中的应用

```cpp
#include <iostream>
#include <functional>
#include <vector>

class Button {
public:
    using Callback = std::function<void()>;
    
    void onClick(Callback cb) {
        callbacks.push_back(cb);
    }
    
    void trigger() {
        for (auto& cb : callbacks) {
            cb();
        }
    }
    
private:
    std::vector<Callback> callbacks;
};

class Application {
public:
    void onButtonClick(int buttonId) {
        std::cout << "Button " << buttonId << " clicked!" << std::endl;
    }
};

int main() {
    Button button;
    Application app;
    
    // 使用bind绑定成员函数和参数
    using namespace std::placeholders;
    button.onClick(std::bind(&Application::onButtonClick, &app, 1));
    
    // 触发事件
    button.trigger(); // 输出: Button 1 clicked!
    
    return 0;
}
```

## 7. 总结与最佳实践

### 7.1 函数适配器的优缺点

#### 优点
- **灵活性**：可以修改函数的行为和接口
- **代码复用**：可以基于现有函数创建新的函数对象
- **表达能力**：可以简洁地表达复杂的函数组合
- **与STL算法集成**：可以方便地与STL算法配合使用

#### 缺点
- **可读性**：复杂的适配器组合可能降低代码可读性
- **编译错误**：使用不当可能导致难以理解的编译错误
- **性能开销**：适配器可能引入轻微的性能开销

### 7.2 最佳实践

1. **优先使用C++11 bind**：相比传统绑定器，`std::bind`更加灵活和强大
2. **合理使用占位符**：使用占位符可以创建更通用的函数对象
3. **考虑Lambda表达式**：对于简单的场景，Lambda表达式可能比bind更简洁
4. **注意参数传递**：确保绑定的参数类型与函数期望的类型匹配
5. **避免过度使用**：复杂的适配器组合可能使代码难以理解，考虑使用命名函数替代

### 7.3 学习建议
- 理解函数适配器的基本原理和工作机制
- 掌握传统绑定器和C++11 bind的使用方法
- 熟悉占位符的使用规则和技巧
- 通过实际项目练习，加深对函数适配器的理解
- 了解函数适配器在现代C++中的应用场景和替代方案

---

## 面试高频考点

### 1. 函数适配器的作用
**问题**：什么是函数适配器？它在STL中有什么作用？

**回答思路**：
- 函数适配器是修改函数对象行为的组件
- 它可以将一种函数接口转换为另一种接口
- 常见的函数适配器包括bind1st、bind2nd、not1、not2等
- 函数适配器使得STL算法更加灵活，可以适应不同的使用场景

### 2. Binder2nd的实现原理
**问题**：Binder2nd是如何将二元函数转换为一元函数的？

**回答思路**：
- Binder2nd保存一个二元函数对象和一个固定的第二个参数
- 当调用转换后的一元函数时，它会将传入的参数作为第一个参数，与保存的第二个参数一起传递给原始的二元函数
- 它继承自unary_function，提供了标准的函数对象接口

### 3. C++11 bind的优势
**问题**：C++11的bind相比传统的bind1st/bind2nd有什么优势？

**回答思路**：
- bind可以绑定任意函数的任意参数，而不仅限于二元函数的第一个或第二个参数
- bind可以改变参数的顺序
- bind可以绑定成员函数和数据成员
- bind的语法更加灵活和直观
- bind返回的函数对象可以存储和传递

### 4. 占位符的工作原理
**问题**：std::bind中的占位符（如_1、_2）是如何工作的？

**回答思路**：
- 占位符是std::placeholders命名空间中的对象
- 当使用bind创建绑定函数时，占位符表示未绑定的参数位置
- 当调用绑定函数时，传入的参数会按照占位符的顺序替换到对应的位置
- 占位符的编号表示调用时参数的位置

### 5. 函数适配器与Lambda表达式的对比
**问题**：在什么情况下应该使用函数适配器，什么情况下应该使用Lambda表达式？

**回答思路**：
- 对于简单的操作，Lambda表达式通常更简洁、可读性更好
- 对于需要复用的复杂逻辑，函数适配器可能更合适
- 当需要与旧代码兼容时，可能需要使用传统的函数适配器
- 当需要绑定成员函数或改变参数顺序时，bind可能比Lambda更方便
- 对于性能敏感的场景，Lambda表达式可能有更好的内联机会

---

## 代码优化建议

### 1. 优先使用C++11 bind
```cpp
// 不推荐：使用传统绑定器
std::count_if(vec.begin(), vec.end(), 
              std::not1(std::bind2nd(std::less<int>(), 50)));

// 推荐：使用C++11 bind
using namespace std::placeholders;
std::count_if(vec.begin(), vec.end(), 
              std::bind(std::greater_equal<int>(), _1, 50));
```

### 2. 合理使用Lambda表达式
```cpp
// 对于简单操作，使用Lambda表达式更简洁
std::count_if(vec.begin(), vec.end(), 
              [](int x) { return x >= 50; });
```

### 3. 注意bind的参数传递
```cpp
// 绑定成员函数时，确保传递正确的对象指针
auto bound_memfn = std::bind(&MyClass::method, &obj, _1);

// 绑定数据成员时，同样需要对象指针
auto bound_memdata = std::bind(&MyClass::member, &obj);
```

### 4. 避免过度复杂的绑定
```cpp
// 不推荐：过于复杂的绑定
auto complex_bind = std::bind(func, 
                              std::bind(subfunc1, _1, _2), 
                              std::bind(subfunc2, _3, _1));

// 推荐：使用命名函数或Lambda表达式
auto complex_func = [](int a, int b, int c) {
    return func(subfunc1(a, b), subfunc2(c, a));
};
```

---

## 总结

函数适配器是STL中非常重要的组件，它们通过修改函数对象的行为，使得STL算法更加灵活和强大。从传统的bind1st、bind2nd、not1到C++11引入的std::bind，函数适配器的功能不断增强，使用方式也更加灵活。

理解函数适配器的工作原理和使用方法，对于掌握STL和C++泛型编程至关重要。通过合理使用函数适配器，可以编写更简洁、更灵活的代码，提高代码的复用性和可维护性。

在现代C++中，虽然Lambda表达式在很多场景下替代了传统的函数适配器，但函数适配器仍然在一些特定场景下发挥着重要作用，特别是std::bind在绑定成员函数和调整参数顺序方面的能力。

通过学习和实践，我们可以根据具体的使用场景，选择最合适的函数对象和适配器，编写高质量的C++代码。