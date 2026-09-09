// 文件作用：
//     学习固定大小 Sliding Window：

//     - Window
//     - left / right
//     - Window State
//     - 滑动时增量更新
//     - O(n)
//     - O(1) Extra Space
#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

// maxFixedWindowSum：
//
// 给定一个整数数组 numbers
// 和固定窗口大小 k，
//
// 寻找：
//
// 所有长度为 k 的连续子数组中
// 最大的元素和。
//
// 例如：
//
// numbers:
// 1 3 2 6 4 5
//
// k = 3
//
// 所有窗口：
//
// 1 3 2 → 6
// 3 2 6 → 11
// 2 6 4 → 12
// 6 4 5 → 15
//
// 最终返回：
//
// 15
//
// 今天重点：
//
// 不要每移动一次窗口
// 就重新把窗口里的 k 个元素全部求和。
//
// 而是维护：
//
// windowSum
//
// 每次窗口右移时：
//
// 减去离开的元素
// +
// 加入新进入的元素
//
// 从而把每次窗口更新降低为 O(1)。
int maxFixedWindowSum(
    const vector<int>& numbers,
    int k
)
{
    // 当前学习阶段先处理非法窗口大小。
    //
    // 合法条件：
    //
    // 1 <= k <= numbers.size()
    if (
        k <= 0 ||
        k > static_cast<int>(numbers.size())
    )
    {
        return 0;
    }

    // --------------------------------------------------
    // 1. 计算第一个窗口
    // --------------------------------------------------
    //
    // 第一个窗口是：
    //
    // [0, k)
    //
    // 也就是：
    //
    // numbers[0]
    // ...
    // numbers[k - 1]
    //
    // 这里只需要完整计算一次。
    int windowSum = 0;

    for (int i = 0; i < k; ++i)
    {
        windowSum += numbers[i];
    }

    // 当前只看过第一个窗口，
    // 所以它暂时就是最大窗口和。
    int maxSum = windowSum;

    cout << "initial window: "
         << "[0, "
         << k - 1
         << "]"
         << ", sum = "
         << windowSum
         << endl;

    // --------------------------------------------------
    // 2. 滑动窗口
    // --------------------------------------------------
    //
    // right 表示：
    //
    // 当前新进入窗口的元素下标。
    //
    // 第一个窗口已经处理：
    //
    // [0, k - 1]
    //
    // 因此第一个新进入的元素下标就是：
    //
    // k
    for (
        int right = k;
        right < static_cast<int>(numbers.size());
        ++right
    )
    {
        // 当 right 进入窗口时，
        //
        // 原窗口最左侧元素必须离开，
        // 才能继续保持窗口长度为 k。
        //
        // 离开的元素下标：
        //
        // right - k
        int leavingIndex =
            right - k;

        // Window State 更新：
        //
        // 旧窗口和
        // - 离开的元素
        // + 新进入的元素
        //
        // 而不是重新遍历整个窗口。
        windowSum -= numbers[leavingIndex];
        windowSum += numbers[right];

        // 当前窗口：
        //
        // [right - k + 1, right]
        int left =
            right - k + 1;

        cout << "window: ["
             << left
             << ", "
             << right
             << "]"
             << ", leave = "
             << numbers[leavingIndex]
             << ", enter = "
             << numbers[right]
             << ", sum = "
             << windowSum
             << endl;

        // 更新目前见过的最大窗口和。
        maxSum = max(
            maxSum,
            windowSum
        );
    }

    return maxSum;
}

int main()
{
    vector<int> numbers = {
        1,
        3,
        2,
        6,
        4,
        5
    };

    int k = 3;

    cout << "=== fixed sliding window ==="
         << endl;

    cout << "k = "
         << k
         << endl
         << endl;

    int answer =
        maxFixedWindowSum(
            numbers,
            k
        );

    cout << endl;

    cout << "max window sum: "
         << answer
         << endl;

    return 0;
}