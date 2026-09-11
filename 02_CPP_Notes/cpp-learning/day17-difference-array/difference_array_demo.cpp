// 文件作用：
// 学习：
//     Difference Array
//     差分构造
//     Range Increment Update
//     difference[left] += value
//     difference[right + 1] -= value
//     Prefix Sum恢复原数组
//     多次区间修改
#include <iostream>
#include <vector>

using namespace std;

// buildDifference：
//
// 根据原始数组 numbers
// 构造 Difference Array。
//
// 今天采用：
//
// difference[0]
// = numbers[0]
//
// difference[i]
// = numbers[i] - numbers[i - 1]
//
// Difference 保存的不是：
//
// “当前位置元素是什么”
//
// 而是：
//
// “当前位置相比前一个位置变化了多少”。
//
// --------------------------------------------------
// 为什么 size = n + 1？
// --------------------------------------------------
//
// 多出来的最后一个位置：
//
// difference[n]
//
// 用来方便处理：
//
// right == n - 1
//
// 时：
//
// difference[right + 1] -= value
//
// 从而避免额外的边界 if 判断。
vector<int> buildDifference(
    const vector<int>& numbers
)
{
    vector<int> difference(
        numbers.size() + 1,
        0
    );

    // 空数组没有需要构造的内容。
    if (numbers.empty())
    {
        return difference;
    }

    // 第一个Difference元素：
    //
    // 没有前一个元素可以相减，
    // 所以直接保存numbers[0]。
    difference[0] =
        numbers[0];

    // 后续：
    //
    // difference[i]
    // =
    // numbers[i] - numbers[i - 1]
    for (
        int i = 1;
        i < static_cast<int>(numbers.size());
        ++i
    )
    {
        difference[i] =
            numbers[i] - numbers[i - 1];
    }

    return difference;
}

// addRange：
//
// 给原数组逻辑上的闭区间：
//
// [left, right]
//
// 中的每个元素都增加：
//
// value
//
// 但是我们不直接修改原数组的每个元素。
//
// 而是在Difference Array中只修改两个位置：
//
// difference[left] += value;
//
// difference[right + 1] -= value;
//
// --------------------------------------------------
// 含义：
//
// left：
// 开始产生 +value 的影响。
//
// right + 1：
// 从这里开始取消 +value 的影响。
//
// 这样对difference重新做Prefix Sum以后，
//
// [left, right]
//
// 内的所有原数组元素都会整体增加value。
void addRange(
    vector<int>& difference,
    int left,
    int right,
    int value
)
{
    // 从left开始打开 +value 的影响。
    difference[left] += value;

    // 在right + 1关闭这段影响。
    //
    // 因为difference比原数组多一个辅助位置，
    // 即使right是原数组最后一个下标，
    // 这里仍然合法。
    difference[right + 1] -= value;
}

// restoreArray：
//
// 对Difference Array做Prefix Sum，
// 还原出区间修改后的最终原数组。
//
// 注意：
//
// difference最后多出来的辅助位置
// 不属于真正的原数组。
//
// 所以这里只恢复：
//
// 前 originalSize 个元素。
vector<int> restoreArray(
    const vector<int>& difference,
    int originalSize
)
{
    vector<int> result(
        originalSize,
        0
    );

    if (originalSize == 0)
    {
        return result;
    }

    // Prefix Sum的累计状态。
    int currentValue = 0;

    for (int i = 0; i < originalSize; ++i)
    {
        // Difference的前缀和
        // 就是当前位置真正的数组值。
        currentValue += difference[i];

        result[i] =
            currentValue;
    }

    return result;
}

// printVector：
//
// 学习辅助函数：
// 打印vector中的所有元素。
void printVector(
    const vector<int>& values
)
{
    for (int value : values)
    {
        cout << value << ' ';
    }

    cout << endl;
}

int main()
{
    vector<int> numbers = {
        10,
        20,
        30,
        40,
        50
    };

    cout << "=== original ==="
         << endl;

    printVector(numbers);

    // --------------------------------------------------
    // 1. 构造Difference Array
    // --------------------------------------------------
    vector<int> difference =
        buildDifference(numbers);

    cout << endl;

    cout << "=== difference before updates ==="
         << endl;

    printVector(difference);

    // --------------------------------------------------
    // 2. 第一次区间修改
    // --------------------------------------------------
    //
    // [1, 3] 中所有元素 +5
    //
    // 原本：
    //
    // 10 20 30 40 50
    //
    // 如果只执行这一次，
    // 应该得到：
    //
    // 10 25 35 45 50
    addRange(
        difference,
        1,
        3,
        5
    );

    // --------------------------------------------------
    // 3. 第二次区间修改
    // --------------------------------------------------
    //
    // [0, 2] 中所有元素 +2
    addRange(
        difference,
        0,
        2,
        2
    );

    // --------------------------------------------------
    // 4. 第三次区间修改
    // --------------------------------------------------
    //
    // [2, 4] 中所有元素 -3
    //
    // 注意：
    //
    // value可以是负数。
    //
    // “增加 -3”
    // 就等价于：
    //
    // 所有元素减3。
    addRange(
        difference,
        2,
        4,
        -3
    );

    cout << endl;

    cout << "=== difference after updates ==="
         << endl;

    printVector(difference);

    // --------------------------------------------------
    // 5. 最后统一恢复数组
    // --------------------------------------------------
    //
    // 多次Range Update期间，
    // 我们始终没有逐个修改原数组。
    //
    // 所有变化都记录在Difference Array中。
    //
    // 最后只做一次Prefix Sum，
    // 得到所有更新后的最终结果。
    vector<int> result =
        restoreArray(
            difference,
            static_cast<int>(numbers.size())
        );

    cout << endl;

    cout << "=== result ==="
         << endl;

    printVector(result);

    return 0;
}