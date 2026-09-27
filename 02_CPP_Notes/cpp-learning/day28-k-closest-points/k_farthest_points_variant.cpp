// // 文件职责：
// // 将K Closest模型反向迁移为K Farthest
// #include <iostream>
// #include <queue>
// #include <utility>
// #include <vector>

// using namespace std;

// int main()
// {
//     vector<pair<int, int>> points{
//         {1, 3},
//         {-2, 2},
//         {5, 8},
//         {0, 1}
//     };

//     int k = 3;

    
//     priority_queue<
//         pair<int, pair<int, int>>,
//         vector<pair<int, pair<int, int>>>,
//         greater<pair<int, pair<int, int>>>
//     > minHeap;

//     for (const auto& point : points)
//     {
//         int x = point.first;
//         int y = point.second;

//         // 比较距离大小时不需要sqrt，
//         // 直接使用距离平方即可。
//         int distanceSquared =
//             x * x + y * y;

//         minHeap.push({
//             distanceSquared,
//             {x, y}
//         });

//         // 只保留距离最大的K个候选。
//         if (
//             static_cast<int>(
//                 minHeap.size()
//             ) > k
//         ) {
//             minHeap.pop();
//         }
//     }

//     cout << "K closest points:"
//          << endl;

//     while (!maxHeap.empty())
//     {
//         int distanceSquared =
//             maxHeap.top().first;

//         int x =
//             maxHeap.top().
//                 second.first;

//         int y =
//             maxHeap.top().
//                 second.second;

//         cout << "("
//              << x
//              << ", "
//              << y
//              << ")"
//              << ", distance^2 = "
//              << distanceSquared
//              << endl;

//         maxHeap.pop();
//     }

//     return 0;
// }