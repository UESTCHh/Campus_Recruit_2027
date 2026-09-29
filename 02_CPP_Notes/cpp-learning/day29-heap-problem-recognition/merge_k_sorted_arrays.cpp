//文件职责：
//训练未知Heap题识别
//+
//动态候选集合
//+
//Heap Candidate保存后续状态
//+
//K路归并
#include <iostream>
#include <queue>
#include <vector>

using namespace std;

// Heap中的一个候选元素。
// 除了候选值本身，还必须记录它来自哪个数组、哪个位置，
// 这样弹出以后才能找到该数组的下一个候选。
struct Candidate
{
    int value;
    int arrayIndex;
    int elementIndex;
};

// 让value更小的Candidate位于heap.top()。
struct CompareCandidate
{
    bool operator()(
        const Candidate& a,
        const Candidate& b
        ) const
    {
        return a.value > b.value;
    }
};

int main()
{
    vector<vector<int>> arrays{
        {1, 4, 9},
        {2, 6, 8},
        {3, 5, 10}
    };

    priority_queue<
        Candidate,
        vector<Candidate>,
        CompareCandidate
    > minHeap;

    // 初始化时，每个非空数组只把第一个元素放入Heap。
    for (
        int arrayIndex = 0;
        arrayIndex <
        static_cast<int>(arrays.size());
        ++arrayIndex
        ) {
        if (!arrays[arrayIndex].empty())
        {
            minHeap.push({
                arrays[arrayIndex][0],
                arrayIndex,
                0
                });
        }
    }

    vector<int> result;

    while (!minHeap.empty())
    {
        // 当前所有候选中最小的元素。
        Candidate current =
            minHeap.top();

        minHeap.pop();

        result.push_back(
            current.value
        );

        // current所在数组中的下一个位置。
        int nextElementIndex =
            current.elementIndex + 1;

        // 如果该数组后面还有元素，
        // 就让下一个元素成为新的候选。
        if (
            nextElementIndex <
            static_cast<int>(
                arrays[
                    current.arrayIndex
                ].size()
                        )
            ) {
            minHeap.push({
                arrays[
                    current.arrayIndex
                ][nextElementIndex],

                current.arrayIndex,

                nextElementIndex
                });
        }
    }

    cout << "Merged result: ";

    for (int value : result)
    {
        cout
            << value
            << ' ';
    }

    cout << '\n';

    return 0;
}