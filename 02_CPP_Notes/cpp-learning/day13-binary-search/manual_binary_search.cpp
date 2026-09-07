// 文件作用：
//     第一次手写Binary Search
//     理解：
//     left
//     right
//     mid
//     Search Space
//     [left, right]
#include <iostream>
#include <vector>

using namespace std;

// binarySearch：
//
// 在一个“升序”的 vector<int> 中寻找 target。
//
// 找到：
// 返回目标值对应的下标。
//
// 没找到：
// 返回 -1。
//
// 今天这个版本使用：
//
// [left, right]
//
// 也就是左右都包含的“闭区间”作为搜索空间。
int binarySearch(
    const vector<int>& numbers,
    int target
)
{
    // 搜索区间左端点。
    int left = 0;

    // 搜索区间右端点。
    //
    // 因为下标从0开始，
    // 最后一个元素下标是 size() - 1。
    int right =
        static_cast<int>(numbers.size()) - 1;

    // 当前Search Space是：
    //
    // [left, right]
    //
    // 因为是闭区间，
    // 当left == right时，
    // 区间中仍然存在一个需要检查的元素。
    //
    // 所以条件是：
    //
    // left <= right
    while (left <= right)
    {
        // 计算当前搜索区间的中间位置。
        //
        // 也可以写：
        //
        // (left + right) / 2
        //
        // 但：
        //
        // left + (right - left) / 2
        //
        // 可以避免left + right非常大时
        // 可能发生的整数溢出问题。
        int mid =
            left + (right - left) / 2;

        cout << "left = "
             << left
             << ", right = "
             << right
             << ", mid = "
             << mid
             << ", value = "
             << numbers[mid]
             << endl;

        // 情况1：
        //
        // 中间元素正好等于target。
        //
        // 说明已经找到，
        // 直接返回下标mid。
        if (numbers[mid] == target)
        {
            return mid;
        }

        // 情况2：
        //
        // 中间元素小于target。
        //
        // 因为numbers是升序：
        //
        // numbers[mid]以及它左边的元素
        // 都不可能等于target。
        //
        // 所以可以全部排除：
        //
        // [left, mid]
        //
        // 新搜索范围变成：
        //
        // [mid + 1, right]
        if (numbers[mid] < target)
        {
            left = mid + 1;
        }
        else
        {
            // 情况3：
            //
            // numbers[mid] > target
            //
            // 因为数组升序，
            // mid以及它右边的元素
            // 都可以排除。
            //
            // 新搜索范围：
            //
            // [left, mid - 1]
            right = mid - 1;
        }
    }

    // 如果循环结束：
    //
    // left > right
    //
    // 表示搜索区间已经为空，
    // 仍然没有找到target。
    return -1;
}

int main()
{
    // Binary Search的前提：
    // 当前数据已经按照升序排列。
    vector<int> numbers = {
        1,
        3,
        5,
        7,
        9,
        11,
        13
    };

    int target = 8;

    cout << "=== search "
         << target
         << " ==="
         << endl;

    int index =
        binarySearch(
            numbers,
            target
        );

    cout << endl;

    if (index != -1)
    {
        cout << "found at index: "
             << index
             << endl;
    }
    else
    {
        cout << "not found"
             << endl;
    }

    return 0;
}