// // 文件职责：
// // 使用大小为K的Min Heap
// // 解决LeetCode 215：
// // Kth Largest Element in an Array
// #include <functional>
// #include <iostream>
// #include <queue>
// #include <vector>
// class Solution {
// public:
//     int findKthLargest(vector<int>& nums, int k) {
//         int n = static_cast<int>(nums.size());

//         priority_queue<
//             int,
//             vector<int>,
//             greater<int>
//         > minHeap;

//         for (int i = 0; i < n; ++i) {
//             if (i < k) {
//                 minHeap.push(nums[i]);
//                 continue;
//             }

//             if (nums[i] > minHeap.top()) {
//                 minHeap.pop();
//                 minHeap.push(nums[i]);
//             }
//         }

//         return minHeap.top();
//     }
// };