#include <algorithm>
#include <functional>
#include <iostream>
#include <vector>

using namespace std;

// 这是一个普通的Comparator函数。
//
// sort会在排序过程中反复调用这个函数，
// 用它判断两个元素之间应该采用什么顺序。
//
// Comparator最重要的理解：
//
// 如果 descending(a, b) 返回 true，
// 表示：
// a 应该排在 b 前面。
//
// 这里使用：
// a > b
//
// 所以较大的元素应该排在较小元素前面，
// 最终得到降序。
bool descending(
    int a,
    int b
)
{
    return a > b;
}

void printNumbers(
    const vector<int>& numbers
)
{
    // const vector<int>&：
    //
    // &：
    // 使用引用，避免复制整个vector。
    //
    // const：
    // 保证这个打印函数不会修改numbers。
    for (int value : numbers)
    {
        cout << value << ' ';
    }

    cout << endl;
}

int main()
{
    vector<int> numbers = {
        30,
        10,
        50,
        20,
        40
    };

    cout << "=== original ==="
         << endl;

    printNumbers(numbers);

    // 第一种自定义排序方式：
    //
    // 将普通函数 descending
    // 作为Comparator传给sort。
    //
    // 注意：
    // 这里是 descending
    //
    // 而不是：
    // descending()
    //
    // 因为我们传的是函数本身，
    // sort会在内部自己调用它。
    sort(
        numbers.begin(),
        numbers.end(),
        descending
    );

    cout << endl;
    cout << "=== function comparator ==="
         << endl;

    printNumbers(numbers);

    // greater<int>() 是标准库提供的比较器。
    //
    // 它同样可以用于sort实现降序排列。
    sort(
        numbers.begin(),
        numbers.end(),
        greater<int>()
    );

    cout << endl;
    cout << "=== greater<int> ==="
         << endl;

    printNumbers(numbers);

    // 第三种方式：
    // Lambda Comparator。
    //
    // [](int a, int b)
    // {
    //     return a < b;
    // }
    //
    // 表示：
    // 如果a比b小，
    // 那么a应该排在b前面。
    //
    // 所以最终得到升序。
    sort(
        numbers.begin(),
        numbers.end(),
        [](int a, int b)
        {
            return a < b;
        }
    );

    cout << endl;
    cout << "=== lambda ascending ==="
         << endl;

    printNumbers(numbers);

    // 再使用Lambda实现降序。
    //
    // a > b
    // → a比b大时，a应该在b前面
    // → 降序
    sort(
        numbers.begin(),
        numbers.end(),
        [](int a, int b)
        {
            return a > b;
        }
    );

    cout << endl;
    cout << "=== lambda descending ==="
         << endl;

    printNumbers(numbers);

    return 0;
}