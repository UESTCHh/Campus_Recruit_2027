// **************************************************
// 通过先查找再插入，避免重复使用同一个数组下标。
// **************************************************

// class Solution {
// public:
//     vector<int> twoSum(vector<int>& nums, int target) {
//         unordered_map<int, int> map;
//         for(int i = 0; i < static_cast<int>(nums.size()); i++){
//             auto it = map.find(target - nums[i]);
//             if(it != map.end()){
//                 return {i, map[target - nums[i]]};
//             }
//             map[nums[i]] = i;
//         }
//         return {-1, -1};
//     }
// };
// 2 7 11 15 t = 9
// 2 0
// 7 1

// 3 3 t = 6
// 3 0
// 3 1

// 3 0
// 3 1

// class Solution {
// public:
//     vector<int> twoSum(vector<int>& nums, int target) {
//         unordered_map<int, int> map;

//         for (int i = 0; i < static_cast<int>(nums.size()); i++) {

//             // 计算当前元素所需要的另一个数
//             int complement = target - nums[i];

//             // 查找之前是否出现过这个数
//             auto it = map.find(complement);

//             if (it != map.end()) {
//                 // it->second为之前元素的下标
//                 return {it->second, i};
//             }

//             // 没找到时，记录当前元素及下标
//             map[nums[i]] = i;
//         }

//         return {-1, -1};
//     }
// };