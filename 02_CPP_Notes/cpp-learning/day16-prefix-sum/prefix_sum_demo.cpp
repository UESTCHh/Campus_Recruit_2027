// 文件作用：
//     学习：
//     Prefix Sum
//     prefix[n + 1]
//     prefix[i + 1] = prefix[i] + numbers[i]
//     [left, right] 区间查询
//     prefix[right + 1] - prefix[left]
//     O(n)预处理
//     O(1)查询
#include <iostream>
#include <vector>

using namespace std;

// buildPrefixSum：
//
// 根据原始数组 numbers
// 构造 Prefix Sum 数组。
//
// 今天采用的定义：
//
// prefix[i]
//
// 表示：
//
// numbers 中“前 i 个元素”的总和。
//
// 因此：
//
// prefix[0] = 0
//
// prefix[1]
// = numbers[0]
//
// prefix[2]
// = numbers[0] + numbers[1]
//
// ...
//
// 如果：
//
// numbers.size() == n
//
// 那么：
//
// prefix.size() == n + 1
vector<int> buildPrefixSum(
    const vector<int>& numbers
)
{
    // Prefix 比原数组多一个位置。
    //
    // prefix[0] 保持为 0，
    // 表示“前0个元素之和”。
    vector<int> prefix(
        numbers.size() + 1,
        0
    );

    // 递推关系：
    //
    // prefix[i + 1]
    // =
    // prefix[i] + numbers[i]
    //
    // 也就是：
    //
    // 已经知道前 i 个元素之和，
    // 再加入 numbers[i]，
    // 得到前 i + 1 个元素之和。
    for (
        int i = 0;
        i < static_cast<int>(numbers.size());
        ++i
    )
    {
        prefix[i + 1] =
            prefix[i] + numbers[i];
    }

    return prefix;
}

// rangeSum：
//
// 查询原数组闭区间：
//
// [left, right]
//
// 的元素总和。
//
// 核心公式：
//
// prefix[right + 1]
// -
// prefix[left]
//
// 原理：
//
// prefix[right + 1]
//
// 包含：
//
// numbers[0]
// ...
// numbers[right]
//
// prefix[left]
//
// 包含：
//
// numbers[0]
// ...
// numbers[left - 1]
//
// 两者相减以后，
// 前面的公共部分被抵消，
//
// 正好剩下：
//
// numbers[left]
// ...
// numbers[right]
int rangeSum(
    const vector<int>& prefix,
    int left,
    int right
)
{
    return
        prefix[right + 1]
        -
        prefix[left];
}

int main()
{
    vector<int> numbers = {
        2,
        4,
        1,
        7,
        3,
        6
    };

    cout << "=== numbers ==="
         << endl;

    for (int value : numbers)
    {
        cout << value << ' ';
    }

    cout << endl;

    // --------------------------------------------------
    // 1. 构造 Prefix Sum
    // --------------------------------------------------
    vector<int> prefix =
        buildPrefixSum(numbers);

    cout << endl;

    cout << "=== prefix ==="
         << endl;

    for (int value : prefix)
    {
        cout << value << ' ';
    }

    cout << endl;

    // --------------------------------------------------
    // 2. 查询 [1, 4]
    // --------------------------------------------------
    //
    // numbers[1..4]:
    //
    // 4 1 7 3
    //
    // sum = 15
    int left1 = 1;
    int right1 = 4;

    int sum1 = rangeSum(
        prefix,
        left1,
        right1
    );

    cout << endl;

    cout << "sum ["
         << left1
         << ", "
         << right1
         << "] = "
         << sum1
         << endl;

    // --------------------------------------------------
    // 3. 查询 [0, 2]
    // --------------------------------------------------
    //
    // numbers[0..2]:
    //
    // 2 4 1
    //
    // sum = 7
    //
    // 注意：
    //
    // left == 0
    //
    // 仍然直接使用完全相同的公式：
    //
    // prefix[right + 1] - prefix[left]
    //
    // 因为：
    //
    // prefix[0] == 0
    int left2 = 0;
    int right2 = 2;

    int sum2 = rangeSum(
        prefix,
        left2,
        right2
    );

    cout << "sum ["
         << left2
         << ", "
         << right2
         << "] = "
         << sum2
         << endl;

    // --------------------------------------------------
    // 4. 查询整个数组
    // --------------------------------------------------
    //
    // [0, numbers.size() - 1]
    //
    // 结果实际上就是：
    //
    // prefix[n]
    int left3 = 0;

    int right3 =
        static_cast<int>(numbers.size()) - 1;

    int sum3 = rangeSum(
        prefix,
        left3,
        right3
    );

    cout << "sum ["
         << left3
         << ", "
         << right3
         << "] = "
         << sum3
         << endl;

    return 0;
}