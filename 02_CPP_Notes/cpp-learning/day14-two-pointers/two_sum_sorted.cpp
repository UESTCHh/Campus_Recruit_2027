// 文件作用：
//     学习：
//     相向双指针
//     有序数组 Two Sum
//     left / right
//     根据sum缩小Search Space
#include <iostream>
#include <vector>

using namespace std;

// twoSumSorted：
//
// 在一个“升序”的 vector<int> 中，
// 寻找两个不同位置的元素，
// 使它们的和等于 target。
//
// 找到：
// 将两个下标保存到 result 中，返回 true。
//
// 没找到：
// 返回 false。
//
// 今天重点不是接口设计，
// 而是理解相向双指针：
//
// left  → 从最左侧开始
// right → 从最右侧开始
//
// 根据当前sum和target的大小关系，
// 安全地排除一个端点。
bool twoSumSorted(
    const vector<int>& numbers,
    int target,
    vector<int>& result
)
{
    // left 指向当前搜索范围最左侧。
    int left = 0;

    // right 指向当前搜索范围最右侧。
    int right =
        static_cast<int>(numbers.size()) - 1;

    // 需要两个不同的元素，
    // 所以只有：
    //
    // left < right
    //
    // 时才还有至少两个候选元素。
    while (left < right)
    {
        int sum =
            numbers[left] + numbers[right];

        cout << "left = "
             << left
             << " ("
             << numbers[left]
             << "), right = "
             << right
             << " ("
             << numbers[right]
             << "), sum = "
             << sum
             << endl;

        // 情况1：
        // 当前两个元素之和正好等于target。
        if (sum == target)
        {
            result.push_back(left);
            result.push_back(right);

            return true;
        }

        // 情况2：
        //
        // sum < target
        //
        // 当前和太小。
        //
        // 因为数组已经升序，
        // right已经是当前搜索范围中的较大元素。
        //
        // numbers[left]连和numbers[right]相加
        // 都仍然太小，
        //
        // 那么numbers[left]和更小的元素组合
        // 更不可能得到target。
        //
        // 所以可以安全排除当前left。
        if (sum < target)
        {
            ++left;
        }
        else
        {
            // 情况3：
            //
            // sum > target
            //
            // 当前和太大。
            //
            // left已经是当前搜索范围中的较小元素。
            //
            // numbers[right]连和numbers[left]相加
            // 都已经太大，
            //
            // 那么numbers[right]与其他更大的元素组合
            // 只会更大。
            //
            // 所以可以安全排除当前right。
            --right;
        }
    }

    // left >= right：
    //
    // 已经没有两个不同位置的元素可供组合，
    // 说明没有找到答案。
    return false;
}

int main()
{
    // 相向双指针能够利用有序性。
    //
    // 今天的数据已经按照升序排列。
    vector<int> numbers = {
        1,
        2,
        4,
        6,
        8,
        10
    };

    int target = 10;

    // 用来保存最终找到的两个下标。
    vector<int> result;

    cout << "=== target = "
         << target
         << " ==="
         << endl;

    bool found = twoSumSorted(
        numbers,
        target,
        result
    );

    cout << endl;

    if (found)
    {
        cout << "found: index "
             << result[0]
             << " and "
             << result[1]
             << endl;

        cout << "values: "
             << numbers[result[0]]
             << " + "
             << numbers[result[1]]
             << " = "
             << target
             << endl;
    }
    else
    {
        cout << "not found"
             << endl;
    }

    return 0;
}