// 文件作用：
//     学习有序vector上的：
//     binary_search
//     lower_bound
//     upper_bound

//     并观察：
//     目标存在
//     目标不存在
//     重复元素
#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

int main()
{
    // 二分查找最重要的前提：
    //
    // 当前操作的范围需要满足对应的有序关系。
    //
    // 今天先使用最简单的：
    // 从小到大排列的 vector<int>。
    vector<int> numbers = {
        1,
        2,
        2,
        2,
        4,
        6,
        8,
        10
    };

    cout << "=== numbers ===" << endl;

    for (int value : numbers)
    {
        cout << value << ' ';
    }

    cout << endl;

    // --------------------------------------------------
    // 1. binary_search
    // --------------------------------------------------
    //
    // binary_search只回答：
    //
    // 目标值是否存在？
    //
    // 返回：
    // bool
    //
    // 它不会返回目标元素的Iterator。
    bool hasTwo = binary_search(
        numbers.begin(),
        numbers.end(),
        2
    );

    bool hasFive = binary_search(
        numbers.begin(),
        numbers.end(),
        5
    );

    // boolalpha让bool输出为：
    //
    // true / false
    //
    // 而不是：
    //
    // 1 / 0
    cout << boolalpha;

    cout << endl;
    cout << "binary_search 2: "
         << hasTwo
         << endl;

    cout << "binary_search 5: "
         << hasFive
         << endl;

    // --------------------------------------------------
    // 2. lower_bound
    // --------------------------------------------------
    //
    // 默认升序情况下：
    //
    // lower_bound(value)
    //
    // 返回第一个：
    //
    // >= value
    //
    // 的位置。
    auto lowerTwo = lower_bound(
        numbers.begin(),
        numbers.end(),
        2
    );

    if (lowerTwo != numbers.end())
    {
        cout << endl;
        cout << "lower_bound(2) value: "
             << *lowerTwo
             << endl;

        // vector的Iterator支持随机访问，
        // 所以可以通过Iterator相减计算下标距离。
        cout << "lower_bound(2) index: "
             << lowerTwo - numbers.begin()
             << endl;
    }

    // --------------------------------------------------
    // 3. upper_bound
    // --------------------------------------------------
    //
    // 默认升序情况下：
    //
    // upper_bound(value)
    //
    // 返回第一个：
    //
    // > value
    //
    // 的位置。
    auto upperTwo = upper_bound(
        numbers.begin(),
        numbers.end(),
        2
    );

    if (upperTwo != numbers.end())
    {
        cout << endl;
        cout << "upper_bound(2) value: "
             << *upperTwo
             << endl;

        cout << "upper_bound(2) index: "
             << upperTwo - numbers.begin()
             << endl;
    }

    // --------------------------------------------------
    // 4. 计算重复元素数量
    // --------------------------------------------------
    //
    // [lower_bound(2), upper_bound(2))
    //
    // 正好覆盖全部等于2的元素。
    //
    // 因此：
    //
    // upper - lower
    //
    // 就是2出现的次数。
    auto occurrences =
        upperTwo - lowerTwo;

    cout << endl;
    cout << "count of 2: "
         << occurrences
         << endl;

    // --------------------------------------------------
    // 5. 目标不存在时的lower_bound
    // --------------------------------------------------
    //
    // 当前数据中没有5：
    //
    // 1 2 2 2 4 6 8 10
    //
    // lower_bound(5)
    //
    // 找的是：
    //
    // 第一个 >= 5
    //
    // 因此应该指向6，
    // 而不是简单返回“没找到”。
    auto lowerFive = lower_bound(
        numbers.begin(),
        numbers.end(),
        5
    );

    if (lowerFive != numbers.end())
    {
        cout << endl;
        cout << "lower_bound(5) value: "
             << *lowerFive
             << endl;

        cout << "lower_bound(5) index: "
             << lowerFive - numbers.begin()
             << endl;
    }

    return 0;
}