// 文件作用：
//     学习：
//     Variable-size Sliding Window
//     left / right
//     Window Expansion
//     Window Shrinking
//     while条件收缩
//     最短连续子数组
//     O(n)摊还分析
#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

// minSubarrayLength：
//
// 给定一个“全部由正整数构成”的数组 numbers，
// 寻找：
//
// 元素和 >= target 的
// 最短连续子数组长度。
//
// 例如：
//
// numbers:
// 2 3 1 2 4 3
//
// target:
// 7
//
// 最短满足条件的连续子数组：
//
// 4 3
//
// 所以返回：
//
// 2
//
// --------------------------------------------------
// 今天学习的是：
// Variable-size Sliding Window
// 可变长度滑动窗口
// --------------------------------------------------
//
// right：
// 负责不断扩大窗口。
//
// left：
// 当窗口已经满足条件以后，
// 负责尝试缩小窗口。
//
// Window State：
// 当前窗口元素总和 windowSum。
int minSubarrayLength(
    const vector<int>& numbers,
    int target
)
{
    // left 表示当前窗口左边界。
    int left = 0;

    // 当前 Window State：
    //
    // numbers[left] 到 numbers[right]
    // 所有元素之和。
    int windowSum = 0;

    // minLength 保存当前找到的最短合法窗口长度。
    //
    // 初始使用：
    //
    // numbers.size() + 1
    //
    // 表示：
    // 目前还没有找到任何合法窗口。
    int minLength =
        static_cast<int>(numbers.size()) + 1;

    // right 从左向右扫描。
    //
    // 每次 right 向右移动一格，
    // 都意味着一个新元素进入窗口。
    for (
        int right = 0;
        right < static_cast<int>(numbers.size());
        ++right
    )
    {
        // 扩大 Window：
        //
        // 把新进入的元素加入 Window State。
        windowSum += numbers[right];

        cout << "expand: "
             << "left = "
             << left
             << ", right = "
             << right
             << ", enter = "
             << numbers[right]
             << ", sum = "
             << windowSum
             << endl;

        // 只要当前窗口仍然满足：
        //
        // windowSum >= target
        //
        // 就说明它是一个合法候选答案。
        //
        // 但题目要求“最短”，
        // 所以不能满足一次就停止。
        //
        // 必须不断从左边缩小，
        // 看能不能得到更短的合法窗口。
        while (windowSum >= target)
        {
            // 当前窗口：
            //
            // [left, right]
            //
            // 因此窗口长度：
            //
            // right - left + 1
            int currentLength =
                right - left + 1;

            minLength = min(
                minLength,
                currentLength
            );

            cout << "  valid: ["
                 << left
                 << ", "
                 << right
                 << "]"
                 << ", length = "
                 << currentLength
                 << ", minLength = "
                 << minLength
                 << endl;

            // 尝试缩小Window。
            //
            // 当前最左侧元素离开窗口。
            windowSum -= numbers[left];

            cout << "  shrink: leave = "
                 << numbers[left];

            ++left;

            cout << ", new left = "
                 << left
                 << ", sum = "
                 << windowSum
                 << endl;
        }
    }

    // 如果 minLength 从未被更新，
    // 说明不存在任何满足：
    //
    // sum >= target
    //
    // 的连续子数组。
    if (
        minLength >
        static_cast<int>(numbers.size())
    )
    {
        return 0;
    }

    return minLength;
}

int main()
{
    // 今天这个简单可变窗口算法依赖：
    //
    // 所有元素都是正整数。
    vector<int> numbers = {
        2,
        3,
        1,
        2,
        4,
        3
    };

    int target = 7;

    cout << "=== variable sliding window ==="
         << endl;

    cout << "target = "
         << target
         << endl
         << endl;

    int answer =
        minSubarrayLength(
            numbers,
            target
        );

    cout << endl;

    cout << "minimum length: "
         << answer
         << endl;

    return 0;
}