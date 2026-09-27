// // 文件职责：
// // 独立完成LeetCode 973
// // K Closest Points to Origin
// class Solution {
// public:
//     vector<vector<int>> kClosest(
//         vector<vector<int>>& points,
//         int k
//     ) {
//         int n = static_cast<int>(points.size());
//         priority_queue<pair<int, pair<int, int>>> pq;
//         for(int i = 0; i < n; i++){
//             int x = points[i][0];
//             int y = points[i][1];
//             int distance2 = x * x + y * y;
//             pq.push({distance2, {x, y}});
//             if(static_cast<int>(pq.size()) > k){
//                 pq.pop();
//             }
//         }
//         vector<vector<int>> answer;
//         while(!pq.empty()){
//             answer.push_back({pq.top().second.first,pq.top().second.second});
//             pq.pop();
//         }
//         return answer;
//     }
// };