// 核心思想：数组元素无法被真正删除，只能被覆盖。

// 快指针 (Fast Pointer)： 像一个无情的搜寻者，跳过我们要删除的目标值，只抓取我们需要保留的“好数据”。

// 慢指针 (Slow Pointer)： 像一个老实的搬砖工，记录当前“好数据”应该安放的最新位置。

// C++

// class Solution {
// public:
//     int removeElement(vector<int>& nums, int val) {
//         int slowIndex = 0; // 慢指针
//         // 快指针遍历整个数组
//         for (int fastIndex = 0; fastIndex < nums.size(); fastIndex++) {
//             if (nums[fastIndex] != val) { // 发现好数据
//                 nums[slowIndex] = nums[fastIndex]; // 覆盖到慢指针位置
//                 slowIndex++; // 慢指针往前走一步
//             }
//         }
//         return slowIndex; // 慢指针的最终位置就是新数组的长度
//     }
// };