// 职责：
// 理解Next Greater Element
// ↓
// 先用最直接暴力搜索
// ↓
// 观察O(n²)
// ↓
// 为Monotonic Stack优化做准备
#include <iostream>
#include <vector>

using namespace std;

int main()
{
    // nums中的每一个元素，
    // 都要寻找：
    //
    // “右边第一个严格大于它的元素”
    //
    // 如果不存在，答案为-1。
    vector<int> nums{
        2, 1, 2, 4, 3
    };

    // answer[i]：
    //
    // nums[i]右边第一个更大的元素。
    //
    // 默认值先全部设置为-1，
    // 表示暂时没有找到。
    vector<int> answer(
        nums.size(),
        -1
    );

    // --------------------------------------------------
    // 暴力方法
    // --------------------------------------------------
    //
    // 对于每一个nums[i]，
    // 从i + 1开始向右扫描。
    for (int i = 0;
         i < static_cast<int>(nums.size());
         ++i)
    {
        for (int j = i + 1;
             j < static_cast<int>(nums.size());
             ++j)
        {
            // 找到右边第一个严格更大的元素。
            if (nums[j] > nums[i])
            {
                answer[i] =
                    nums[j];

                // 注意：
                //
                // 题目要求的是“第一个”更大的元素，
                // 所以找到以后必须立即停止。
                break;
            }
        }
    }

    cout << "nums:   ";

    for (int value : nums)
    {
        cout << value << ' ';
    }

    cout << endl;

    cout << "answer: ";

    for (int value : answer)
    {
        cout << value << ' ';
    }

    cout << endl;

    return 0;
}