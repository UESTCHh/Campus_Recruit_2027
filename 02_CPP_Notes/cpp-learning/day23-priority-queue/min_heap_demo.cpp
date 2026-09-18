#include <functional>
#include <iostream>
#include <queue>
#include <vector>

using namespace std;

int main()
{
    // 默认Priority Queue：
    // 最大元素具有最高优先级。
    priority_queue<int> maxHeap;

    // Min Heap：
    // 最小元素具有最高优先级。
    //
    // 三个模板参数：
    //
    // int
    // → 元素类型
    //
    // vector<int>
    // → 底层容器
    //
    // greater<int>
    // → 比较规则，使较小元素位于top
    priority_queue<
        int,
        vector<int>,
        greater<int>
    > minHeap;

    vector<int> numbers{
        5, 1, 8, 3, 10
    };

    for (int value : numbers)
    {
        maxHeap.push(value);
        minHeap.push(value);
    }

    cout << "max heap top = "
         << maxHeap.top()
         << endl;

    cout << "min heap top = "
         << minHeap.top()
         << endl;

    cout << endl;

    cout << "max heap pop order: ";

    while (!maxHeap.empty())
    {
        cout << maxHeap.top()
             << ' ';

        maxHeap.pop();
    }

    cout << endl;

    cout << "min heap pop order: ";

    while (!minHeap.empty())
    {
        cout << minHeap.top()
             << ' ';

        minHeap.pop();
    }

    cout << endl;

    return 0;
}