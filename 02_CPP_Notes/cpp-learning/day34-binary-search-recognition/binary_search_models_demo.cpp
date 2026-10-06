#include <iostream>
#include <vector>

using namespace std;

// 1. 精确二分：
// 找一个具体target。
// 一旦相等，可以直接返回。
int exactSearch(
    const vector<int>& nums,
    int target
) {
    int left = 0;
    int right =
        static_cast<int>(nums.size()) - 1;

    while (left <= right) {
        int mid =
            left + (right - left) / 2;

        if (nums[mid] == target) {
            return mid;
        }

        if (nums[mid] < target) {
            left = mid + 1;
        }
        else {
            right = mid - 1;
        }
    }

    return -1;
}

// 2. 边界二分：
// 找第一个 >= target 的位置。
//
// 判断结果可以理解为：
// false false false | true true true
int firstGreaterEqual(
    const vector<int>& nums,
    int target
) {
    int left = 0;
    int right =
        static_cast<int>(nums.size());

    while (left < right) {
        int mid =
            left + (right - left) / 2;

        if (nums[mid] < target) {
            left = mid + 1;
        }
        else {
            right = mid;
        }
    }

    return left;
}

// 3. 模拟“二分答案”的判定函数。
//
// 假设任务需要至少 requirement 的能力值。
// capability越大，越容易满足要求。
bool canFinish(
    int capability,
    int requirement
) {
    return capability >= requirement;
}

// 在候选答案 [1, maxCapability] 中，
// 找第一个能够完成任务的 capability。
int minimumCapability(
    int maxCapability,
    int requirement
) {
    int left = 1;
    int right = maxCapability;

    while (left < right) {
        int mid =
            left + (right - left) / 2;

        if (!canFinish(
                mid,
                requirement
            )) {
            left = mid + 1;
        }
        else {
            right = mid;
        }
    }

    return left;
}

int main()
{
    vector<int> nums{
        1, 3, 5, 5, 7, 9
    };

    cout
        << "exact search index = "
        << exactSearch(nums, 7)
        << '\n';

    cout
        << "first >= 5 index = "
        << firstGreaterEqual(nums, 5)
        << '\n';

    cout
        << "minimum capability = "
        << minimumCapability(10, 6)
        << '\n';

    return 0;
}

 
