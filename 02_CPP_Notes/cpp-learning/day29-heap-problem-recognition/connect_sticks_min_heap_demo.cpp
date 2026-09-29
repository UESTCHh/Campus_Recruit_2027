//文件职责：
//训练：
//不是Top K
//但因为需要反复取当前极值
//仍然应该想到Heap
#include <functional>
#include <iostream>
#include <queue>
#include <vector>

using namespace std;

int main()
{
    vector<int> sticks{
        2,
        4,
        3,
        6
    };

    // 每一轮都需要取得当前最短的两根木棍，
    // 因此希望top()始终是当前最小值。
    priority_queue<
        int,
        vector<int>,
        greater<int>
    > minHeap;

    for (int length : sticks)
    {
        minHeap.push(length);
    }

    int totalCost = 0;

    // 只要还剩至少两根木棍，
    // 就继续连接。
    while (minHeap.size() > 1)
    {
        int first =
            minHeap.top();
        minHeap.pop();

        int second =
            minHeap.top();
        minHeap.pop();

        int merged =
            first + second;

        totalCost += merged;

        cout
            << first
            << " + "
            << second
            << " = "
            << merged
            << ", totalCost = "
            << totalCost
            << '\n';

        // 新木棍会重新参与后续比较，
        // 所以必须重新放回Heap。
        minHeap.push(merged);
    }

    cout
        << "minimum total cost = "
        << totalCost
        << '\n';

    return 0;
}