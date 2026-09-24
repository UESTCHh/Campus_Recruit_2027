// 文件职责：
// 验证：
// 求最大的K个为什么使用Min Heap
#include <functional>
#include <iostream>
#include <queue>
#include <vector>

using namespace std;

int main()
{
    vector<int> nums{
        10, 4, 8, 3, 15, 6
    };

    int k = 2;

    // 求最大的K个元素：
    // 使用大小为K的Min Heap。
    //
    // top()表示：
    // 当前Top K中最小的元素，
    // 也就是进入Top K的门槛。
    priority_queue<
        int,
        vector<int>,
        greater<int>
    > minHeap;

    for (int value : nums)
    {
        minHeap.push(value);

        if (
            static_cast<int>(
                minHeap.size()
            ) > k
        ) {
            minHeap.pop();
        }

        cout << "after value = "
             << value
             << ", threshold = "
             << minHeap.top()
             << endl;
    }

    cout << endl;

    cout << "Top "
         << k
         << " elements:"
         << endl;

    while (!minHeap.empty())
    {
        cout << minHeap.top()
             << " ";

        minHeap.pop();
    }

    cout << endl;

    return 0;
}