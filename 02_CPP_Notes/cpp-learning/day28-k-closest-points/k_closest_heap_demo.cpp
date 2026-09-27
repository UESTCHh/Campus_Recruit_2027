// 文件职责：
// 验证：
// point
// ↓
// distance²
// ↓
// Size-K Max Heap
// ↓
// K Closest Points
#include <iostream>
#include <queue>
#include <utility>
#include <vector>

using namespace std;

int main()
{
    vector<pair<int, int>> points{
        {1, 3},
        {-2, 2},
        {5, 8},
        {0, 1}
    };

    int k = 3;

    // Heap中保存：
    //
    // (
    //     distanceSquared,
    //     (x, y)
    // )
    //
    // 因为要求距离最小的K个点，
    // 所以使用Size-K Max Heap。
    //
    // top()表示：
    // 当前K个最近点中距离最大的那个，
    // 也就是最容易被淘汰的候选。
    priority_queue<
        pair<int, pair<int, int>>
    > maxHeap;

    for (const auto& point : points)
    {
        int x = point.first;
        int y = point.second;

        // 比较距离大小时不需要sqrt，
        // 直接使用距离平方即可。
        int distanceSquared =
            x * x + y * y;

        maxHeap.push({
            distanceSquared,
            {x, y}
        });

        // 只保留距离最小的K个候选。
        if (
            static_cast<int>(
                maxHeap.size()
            ) > k
        ) {
            maxHeap.pop();
        }
    }

    cout << "K closest points:"
         << endl;

    while (!maxHeap.empty())
    {
        int distanceSquared =
            maxHeap.top().first;

        int x =
            maxHeap.top().
                second.first;

        int y =
            maxHeap.top().
                second.second;

        cout << "("
             << x
             << ", "
             << y
             << ")"
             << ", distance^2 = "
             << distanceSquared
             << endl;

        maxHeap.pop();
    }

    return 0;
}