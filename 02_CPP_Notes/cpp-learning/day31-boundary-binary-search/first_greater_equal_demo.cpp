// 文件职责：
// 训练：
// 精确查找
// vs
// 边界查找

// 并实现：
// first >= target
#include <iostream>
#include <vector>

using namespace std;

// 找到第一个 >= target 的元素下标。
//
// 如果所有元素都 < target，
// 返回 nums.size()。
int firstGreaterEqual(
    const vector<int>& nums,
    int target
) {
    int left = 0;
    int right =
        static_cast<int>(
            nums.size()
        );

    // 搜索区间始终使用：
    // [left, right)
    while (left < right) {
        int mid =
            left + (right - left) / 2;

        if (nums[mid] < target) {
            // nums[mid]以及mid左侧都不可能成为答案。
            left = mid + 1;
        }
        else {
            // nums[mid]已经满足 >= target，
            // 但左边可能还有更早满足条件的位置。
            //
            // 因此mid不能被排除。
            right = mid;
        }
    }

    return left;
}

int main()
{
    vector<int> nums{
        1, 3, 5, 5, 5, 7, 9
    };

    int target = 5;

    int index =
        firstGreaterEqual(
            nums,
            target
        );

    cout
        << "index = "
        << index
        << '\n';

    if (
        index <
        static_cast<int>(
            nums.size()
        )
    ) {
        cout
            << "value = "
            << nums[index]
            << '\n';
    }
    else {
        cout
            << "no element >= target"
            << '\n';
    }

    return 0;
}