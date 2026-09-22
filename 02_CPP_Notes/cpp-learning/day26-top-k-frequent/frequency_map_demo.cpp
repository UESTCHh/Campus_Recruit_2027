// 文件职责：
// 1. 使用unordered_map统计频率
// 2. 理解 element → frequency 映射
// 3. 将频率信息组织成 pair
// 4. 为后续Min Heap做准备
#include <iostream>
#include <unordered_map>
#include <utility>
#include <vector>

using namespace std;

int main()
{
    vector<int> nums{
    4, 4, 2, 8, 8, 8, 4, 3, 8, 2
};

    // key：
    // 原始元素
    //
    // value：
    // 这个元素出现的次数
    unordered_map<int, int> frequency;

    for (int value : nums)
    {
        ++frequency[value];
    }

    cout << "frequency table:" << endl;

    for (const auto& entry : frequency)
    {
        cout << "value = "
             << entry.first
             << ", frequency = "
             << entry.second
             << endl;
    }

    cout << endl;

    // 为后续Heap准备：
    //
    // pair.first  = frequency
    // pair.second = 原始元素
    vector<pair<int, int>> candidates;

    for (const auto& entry : frequency)
    {
        int value = entry.first;
        int count = entry.second;

        candidates.push_back(
            {count, value}
        );
    }

    cout << "heap candidates:" << endl;

    for (const auto& candidate : candidates)
    {
        cout << "(frequency = "
             << candidate.first
             << ", value = "
             << candidate.second
             << ")"
             << endl;
    }

    return 0;
}