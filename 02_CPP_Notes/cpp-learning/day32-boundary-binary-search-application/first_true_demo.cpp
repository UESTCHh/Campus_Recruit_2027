// 职责：
// 不再只把Boundary Binary Search理解为target查找，
// 而是训练：
// false false false | true true
// → 找第一个true
#include <iostream>
#include <vector>

using namespace std;

// 找到第一个满足：
// nums[i] >= requirement
// 的位置。
//
// 今天重点不是背 lower_bound 模板，
// 而是把问题理解为：
//
// false false false | true true true
//                   ↑
//             找第一个 true
int firstSatisfied(
    const vector<int>& nums,
    int requirement
) {
    int left = 0;
    int right =
        static_cast<int>(
            nums.size()
        );

    while (left < right) {
        int mid =
            left + (right - left) / 2;

        // 当前mid不满足要求。
        //
        // 因为数组升序，
        // mid以及mid左边都不可能满足，
        // 所以直接排除。
        if (nums[mid] < requirement) {
            left = mid + 1;
        }
        else {
            // 当前mid已经满足要求，
            // 但左边可能存在更早满足的位置。
            //
            // 因此保留mid，
            // 继续向左寻找边界。
            right = mid;
        }
    }

    return left;
}

int main()
{
    vector<int> nums{
        2, 4, 6, 8, 10, 12
    };

    int requirement = 9;

    int index =
        firstSatisfied(
            nums,
            requirement
        );

    if (
        index ==
        static_cast<int>(
            nums.size()
        )
    ) {
        cout
            << "no value satisfies requirement"
            << '\n';
    }
    else {
        cout
            << "first satisfied index = "
            << index
            << '\n';

        cout
            << "value = "
            << nums[index]
            << '\n';
    }

    return 0;
}