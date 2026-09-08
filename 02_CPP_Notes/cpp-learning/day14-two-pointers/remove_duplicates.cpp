// 文件作用：
//     学习：
//     快慢双指针
//     原地修改
//     有序数组去重
//     slow维护已处理区域
//     fast扫描未处理区域
#include <iostream>
#include <vector>

using namespace std;

// removeDuplicates：
//
// 对一个“已经升序”的 vector<int>
// 进行原地去重。
//
// 例如：
//
// 原始：
// 1 1 2 2 2 3 4 4
//
// 处理以后：
//
// numbers前4个位置：
// 1 2 3 4
//
// 返回：
// 4
//
// 注意：
//
// 我们不要求把vector本身resize。
// 只规定：
//
// [0, newLength)
//
// 是去重以后真正有效的数据。
//
// 算法使用快慢双指针：
//
// fast
// → 扫描尚未处理的数据
//
// slow
// → 指向已经整理好的区域中
//   最后一个唯一元素。
int removeDuplicates(
    vector<int>& numbers
)
{
    // 空vector没有任何元素。
    //
    // 去重后的有效长度自然也是0。
    if (numbers.empty())
    {
        return 0;
    }

    // 第一个元素天然是第一个唯一元素。
    //
    // 所以：
    //
    // slow = 0
    //
    // 表示当前已经整理好的区域：
    //
    // [0, 0]
    //
    // 里面有一个唯一元素。
    int slow = 0;

    // fast从第二个元素开始扫描。
    //
    // fast负责寻找：
    //
    // 下一个与numbers[slow]不同的新值。
    for (
        int fast = 1;
        fast < static_cast<int>(numbers.size());
        ++fast
    )
    {
        cout << "slow = "
             << slow
             << " ("
             << numbers[slow]
             << "), fast = "
             << fast
             << " ("
             << numbers[fast]
             << ")"
             << endl;

        // 因为数组已经升序，
        // 重复元素一定连续出现。
        //
        // 如果：
        //
        // numbers[fast] == numbers[slow]
        //
        // 说明fast当前看到的还是重复值。
        //
        // 不需要写入，
        // fast继续向后扫描即可。
        if (numbers[fast] == numbers[slow])
        {
            continue;
        }

        // 如果执行到这里：
        //
        // numbers[fast] != numbers[slow]
        //
        // 说明发现了一个新的唯一值。
        //
        // 当前：
        //
        // [0, slow]
        //
        // 已经保存了所有发现过的唯一值。
        //
        // 所以下一个唯一值应该写到：
        //
        // slow + 1
        ++slow;

        // 把fast找到的新唯一值
        // 写到整理区域的末尾。
        numbers[slow] = numbers[fast];
    }

    // slow最终指向：
    //
    // 最后一个唯一元素的下标。
    //
    // 例如：
    //
    // slow == 3
    //
    // 有效下标：
    //
    // 0 1 2 3
    //
    // 所以数量：
    //
    // slow + 1 == 4
    return slow + 1;
}

int main()
{
    vector<int> numbers = {
        1,
        1,
        2,
        2,
        2,
        3,
        4,
        4
    };

    cout << "=== original ==="
         << endl;

    for (int value : numbers)
    {
        cout << value << ' ';
    }

    cout << endl
         << endl;

    int newLength =
        removeDuplicates(numbers);

    cout << endl;

    cout << "new length: "
         << newLength
         << endl;

    cout << "unique values: ";

    // 注意：
    //
    // removeDuplicates并没有承诺：
    //
    // [newLength, numbers.size())
    //
    // 里面是什么内容。
    //
    // 我们只读取：
    //
    // [0, newLength)
    //
    // 这一段有效结果。
    for (int i = 0; i < newLength; ++i)
    {
        cout << numbers[i]
             << ' ';
    }

    cout << endl;

    return 0;
}