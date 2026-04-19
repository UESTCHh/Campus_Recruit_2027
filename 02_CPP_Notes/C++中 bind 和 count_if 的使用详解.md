# C++中 bind 和 count_if 的使用详解

### 🔍 核心结论先给你

`fn_` **不是直接调用**，而是把**绑定好的函数对象本身**传给了 `count_if`，由 `count_if` 内部去调用它，所以你看不到显式的 `fn_(x)`。

---

### 🧩 逐行拆解代码逻辑

我们先把最后几行代码完整拎出来，再一步步讲：

```C++

vector<int> v {15,37,94,50,73,58,28,98};

// 第1行：统计「不大于50」的元素个数，用not1+bind2nd
int n = count_if(v.cbegin(), v.cend(), not1(bind2nd(less<int>(),50))); 
cout << "n= " << n << endl; // 输出5

// 第2行：用C++11的std::bind实现等价逻辑
auto fn_ = bind(less<int>(), _1, 50); 

// 第3行：把fn_传给count_if
cout << count_if(v.cbegin(), v.cend(), fn_) << endl; 
```

---

### 1. 先搞懂 `bind(less<int>(), _1, 50)` 做了什么

`std::bind` 的作用是**把一个二元函数（** **`less<int>`** **，需要两个参数 ** **`a < b`** **），绑定成一个一元函数**：

- 原函数：`less<int>{}(a, b)` → 等价于 `a < b`

- `bind(less<int>(), _1, 50)`：

    - 用占位符 `_1` 表示「未来传入的第一个参数」

    - 把第二个参数固定为 `50`

    - 最终生成一个**一元函数对象 ** **`fn_`**，它的行为等价于：

        ```C++
        
        bool fn_(int a) {
            return a < 50; // 因为_1对应a，第二个参数固定为50
        }
        ```

---

### 2. 为什么 `count_if` 里直接写 `fn_`，而不是 `fn_(x)`？

`std::count_if` 的函数签名是：

```C++

template< class InputIt, class UnaryPredicate >
typename iterator_traits<InputIt>::difference_type
count_if( InputIt first, InputIt last, UnaryPredicate p );
```

它的第三个参数 `p` 要求是一个**一元谓词（可以接收1个参数、返回bool的可调用对象）**，作用是：

遍历容器的每个元素 `x`，自动调用 `p(x)`，统计返回 `true` 的次数。

所以这里的逻辑是：

- `fn_` 是一个**已经绑定好的一元函数对象**，完全满足 `count_if` 对谓词的要求

- `count_if` 内部会自动对每个元素 `x` 执行 `fn_(x)`，等价于 `x < 50`

- 你不需要手动写 `fn_(x)`，因为 `count_if` 会帮你完成调用

---

### 3. 对比两种写法，彻底搞懂区别

|写法|本质|适用场景|
|---|---|---|
|`count_if(..., fn_)`|把**函数对象本身**传给算法，由算法内部调用|标准库算法（`count_if`/`for_each`/`find_if`等）的标准用法|
|`fn_(x)`|手动调用函数对象，传入参数 `x`|你自己在代码里主动执行这个函数|
举个直观的例子：

```C++

auto fn_ = bind(less<int>(), _1, 50);
// 手动调用
cout << fn_(10) << endl; // true（10<50）
cout << fn_(60) << endl; // false（60<50不成立）

// 传给count_if，由算法自动调用
vector<int> v = {10,60};
count_if(v.begin(), v.end(), fn_); 
// 等价于：
// count = 0;
// if (fn_(10)) count++; → 1
// if (fn_(60)) count++; → 不变
// 最终返回1
```

---

### 4. 补充：占位符 `_1` 到底是什么？

`_1` 是 `std::placeholders::_1` 的简写，它的作用是**标记「未来会被传入的参数位置」**：

- `bind(func, _1, 50)`：`_1` 对应调用时的第1个参数，第2个参数固定为50

- `bind(func, 50, _1)`：第1个参数固定为50，`_1` 对应调用时的第1个参数，等价于 `func(50, x)`

- `bind(func, _2, _1)`：交换两个参数的顺序，等价于 `func(y, x)`

在 `fn_ = bind(less<int>(), _1, 50)` 里，`_1` 就是在说：「等下调用 `fn_(x)` 时，`x` 会填到这个位置」。

---

### 5. 再看前面的例子，验证逻辑

你图里前面的 `bind` 例子也完全符合这个逻辑：

```C++

auto fn_half = bind(my_divide, _1, 2); // 绑定成一元函数：x/2
cout << fn_half(10) << endl; // 手动调用，传入10，输出5
```

这里 `fn_half` 也是一个一元函数对象，手动调用时需要写 `fn_half(10)`，而如果把它传给 `count_if`，就只需要写 `fn_half`，由算法自动调用。

---

### 💡 避坑提示

- 不要把「函数对象本身」和「函数对象的调用」搞混：

    - `fn_` 是对象，`fn_()` 是调用

    - 标准库算法需要的是**可调用对象**，不是调用结果，所以传对象本身

- `_1` 必须在 `bind` 里使用，它只是一个「占位标记」，不是实际参数

- 成员函数绑定的特殊情况：`bind(&MyPair::multiply, _1)` 里的 `_1` 是用来填 `this` 指针的，所以调用时需要传入对象 `ten_two`，本质也是一样的逻辑

---

要不要我帮你把这段代码补全成一个可直接编译运行的完整示例，同时加注释，让你跑一遍验证每一步的输出？
> （注：文档部分内容可能由 AI 生成）