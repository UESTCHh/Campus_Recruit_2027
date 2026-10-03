#include <iostream>
#include <vector>

using namespace std;

// 找第一个严格大于 target 的位置。
// 如果不存在，返回 nums.size()。
int firstGreater(
    const vector<int>& nums,
    int target
) {
    int left = 0;
    int right =
        static_cast<int>(
            nums.size()
        );

    while (left < right) {
        int mid =
            left + (right - left) / 2;

        if (nums[mid] <= target) {
            // nums[mid]不满足 > target，
            // 并且mid左边也不可能满足。
            left = mid + 1;
        }
        else {
            // nums[mid]已经 > target，
            // 但左边可能还有更早满足的位置。
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
        firstGreater(
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
            << "no element > target"
            << '\n';
    }

    return 0;
}
