// 职责：
// 使用：
// push所有新元素
// +
// size > k时pop

// 验证它和门槛判断版本结果一致
// 职责：
// 理解：
// 静态Top K和动态Top K区别

// 理解：
// 为什么Heap必须成为对象长期状态

// 实验：
// 持续add新元素并返回当前第K大
#include <functional>
#include <iostream>
#include <queue>
#include <vector>

using namespace std;

class KthLargestTracker
{
private:
    // 要维护的是“第k大”。
    int k_;

    // Heap中始终只保存：
    // 当前出现过的最大k个元素。
    //
    // 因为我们最关心Top K中最小的那个门槛，
    // 所以使用Min Heap。
    priority_queue<
        int,
        vector<int>,
        greater<int>
    > minHeap_;

public:
    KthLargestTracker(
        int k,
        const vector<int>& nums
    )
        : k_(k)
    {
        // 使用初始数据建立Top K状态。
        for (int value : nums)
        {
            add(value);
        }
    }

    int add(int value)
    {
        minHeap_.push(value);

        if (
            static_cast<int>(
                minHeap_.size()
            ) > k_
        ) {
            minHeap_.pop();
        }

        return minHeap_.top();
    }
};

int main()
{
    vector<int> nums{
        4, 5, 8, 2
    };

    KthLargestTracker tracker(
        3,
        nums
    );

    cout << "add 3  -> "
         << tracker.add(3)
         << endl;

    cout << "add 5  -> "
         << tracker.add(5)
         << endl;

    cout << "add 10 -> "
         << tracker.add(10)
         << endl;

    cout << "add 9  -> "
         << tracker.add(9)
         << endl;

    cout << "add 4  -> "
         << tracker.add(4)
         << endl;

    return 0;
}