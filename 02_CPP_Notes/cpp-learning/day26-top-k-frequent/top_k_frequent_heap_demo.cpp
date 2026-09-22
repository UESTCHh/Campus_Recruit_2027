// 文件职责：
// unordered_map统计frequency
// ↓
// pair保存(frequency,value)
// ↓
// Min Heap维护Top K Frequent
#include <functional>
#include <iostream>
#include <queue>
#include <unordered_map>
#include <utility>
#include <vector>

using namespace std;

int main()
{
    vector<int> nums{
        1, 1, 1, 2, 2, 3
    };

    int k = 2;

    // key   = 原始元素
    // value = 出现频率
    unordered_map<int, int> frequency;

    for (int value : nums)
    {
        ++frequency[value];
    }

    // pair.first  = frequency
    // pair.second = 原始元素
    //
    // greater<pair<int,int>>
    // 让较小的pair位于Heap顶部。
    //
    // pair默认先比较first，
    // 因此主要先按照frequency比较。
    priority_queue<
        pair<int, int>,
        vector<pair<int, int>>,
        greater<pair<int, int>>
    > minHeap;

    for (const auto& entry : frequency)
    {
        int value = entry.first;
        int count = entry.second;

        minHeap.push(
            {count, value}
        );

        // 始终只留下frequency最高的k个候选。
        if (
            static_cast<int>(
                minHeap.size()
            ) > k
        ) {
            minHeap.pop();
        }
    }

    cout << "Top "
         << k
         << " frequent elements:"
         << endl;

    while (!minHeap.empty())
    {
        int count =
            minHeap.top().first;

        int value =
            minHeap.top().second;

        cout << "value = "
             << value
             << ", frequency = "
             << count
             << endl;

        minHeap.pop();
    }

    return 0;
}
