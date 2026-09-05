#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

int main()
{
    // vector 是真正存储数据的容器。
    //
    // 今天学习的 std::sort、std::reverse、
    // std::find、std::count 并不属于 vector，
    // 而是 <algorithm> 提供的通用算法。
    vector<int> numbers = {
        30,
        10,
        50,
        20,
        40,
        20
    };

    cout << "=== original ===" << endl;

    for (int value : numbers)
    {
        cout << value << ' ';
    }

    cout << endl;

    // std::sort 接收一段 iterator range。
    //
    // numbers.begin()
    // → 区间起点
    //
    // numbers.end()
    // → 最后一个元素之后的位置
    //
    // 因此操作区间是：
    // [begin, end)
    //
    // 默认按照从小到大的顺序排序。
    //
    // sort 会直接修改 numbers 中元素的顺序，
    // 并不会返回一个新的 vector。
    sort(
        numbers.begin(),
        numbers.end()
    );

    cout << endl;
    cout << "=== after sort ===" << endl;

    for (int value : numbers)
    {
        cout << value << ' ';
    }

    cout << endl;

    // std::find 会在指定 iterator range 中，
    // 从前往后寻找第一个等于目标值的元素。
    //
    // 返回值是 iterator。
    auto it = find(
        numbers.begin(),
        numbers.end(),
        30
    );

    // 如果找到：
    // iterator != end()
    //
    // 如果没找到：
    // iterator == end()
    if (it != numbers.end())
    {
        cout << endl;
        cout << "30 found" << it - numbers.begin()
             << endl;
    }

    // std::count 统计范围中，
    // 与指定值相等的元素一共有多少个。
    int occurrences = count(
        numbers.begin(),
        numbers.end(),
        20
    );

    cout << "20 count: "
         << occurrences
         << endl;

    // reverse 和 sort 不一样。
    //
    // reverse 不关心元素数值大小，
    // 它只是把当前范围中的元素顺序整体翻转。
    reverse(
        numbers.begin(),
        numbers.end()
    );

    cout << endl;
    cout << "=== after reverse ==="
         << endl;

    for (int value : numbers)
    {
        cout << value << ' ';
    }

    cout << endl;

    return 0;
}