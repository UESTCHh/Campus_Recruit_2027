// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

 

// 示例 1：

// 输入：nums = [1,1,1,2,2,3], k = 2

// 输出：[1,2]

// 示例 2：

// 输入：nums = [1], k = 1

// 输出：[1]

// 示例 3：

// 输入：nums = [1,2,1,2,1,2,3,1,3,2], k = 2

// 输出：[1,2]

 

// 提示：

// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104
// k 的取值范围是 [1, 数组中不相同的元素的个数]
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
 

// 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
// class Solution {
// public:
//     vector<int> topKFrequent(vector<int>& nums, int k) {
//         int n = static_cast<int> (nums.size());
//         unordered_map<int, int> map;
//         for(int val : nums){
//             ++map[val];
//         }
//         priority_queue<pair<int, int>, vector<pair<int, int>>, greater<pair<int, int>>> minheap;
//         for(auto val : map){
//             minheap.push({val.second, val.first});
//             if(minheap.size() > k){
//                 minheap.pop();
//             }
//         }
//         vector<int> answer;
//         while(!minheap.empty()){
//             answer.push_back(minheap.top().second);
//             minheap.pop();
//         }
//         return answer;
//     }
// };
// class Solution {
// public:
//     vector<int> topKFrequent(
//         vector<int>& nums,
//         int k
//     ) {
//         unordered_map<int, int> frequency;

//         // 统计：
//         // value -> frequency
//         for (int value : nums) {
//             ++frequency[value];
//         }

//         // pair.first  = frequency
//         // pair.second = 原始元素
//         //
//         // 使用Min Heap维护frequency最大的K个元素。
//         priority_queue<
//             pair<int, int>,
//             vector<pair<int, int>>,
//             greater<pair<int, int>>
//         > minHeap;

//         for (const auto& entry : frequency) {
//             minHeap.push({
//                 entry.second,
//                 entry.first
//             });

//             if (
//                 static_cast<int>(
//                     minHeap.size()
//                 ) > k
//             ) {
//                 minHeap.pop();
//             }
//         }

//         vector<int> answer;

//         // 题目只要求返回Top K元素，
//         // 不要求按frequency降序排列。
//         while (!minHeap.empty()) {
//             answer.push_back(
//                 minHeap.top().second
//             );

//             minHeap.pop();
//         }

//         return answer;
//     }
// };