// 文件职责：
// 验证：
// 有序输入
// ↓
// Binary Search定位
// ↓
// Two Pointers扩展
// ↓
// 连续区间答案
#include <iostream>
#include <vector>

using namespace std;

int main()
{
    vector<int> arr{
        1, 2, 3, 4, 5,
        6, 7, 8, 9
    };

    int k = 4;
    int x = 6;

    int n =
        static_cast<int>(
            arr.size()
        );

    // 手写lower_bound：
    // 找到第一个 >= x 的位置。
    int low = 0;
    int high = n;

    while (low < high)
    {
        int mid =
            low + (high - low) / 2;

        if (arr[mid] < x)
        {
            low = mid + 1;
        }
        else
        {
            high = mid;
        }
    }

    // low就是第一个 >= x 的位置。
    int right = low;
    int left = right - 1;

    // 一共选择K个元素。
    for (int count = 0; count < k; ++count)
    {
        if (left < 0)
        {
            ++right;
        }
        else if (right >= n)
        {
            --left;
        }
        else if (
            x - arr[left]
            <=
            arr[right] - x
        ) {
            // 距离相同也选左边，
            // 因为左边值更小。
            --left;
        }
        else
        {
            ++right;
        }
    }

    cout << "result: ";

    // 被选中的元素最终形成连续区间：
    // [left + 1, right)
    for (
        int i = left + 1;
        i < right;
        ++i
    ) {
        cout
            << arr[i]
            << ' ';
    }

    cout << '\n';

    return 0;
}