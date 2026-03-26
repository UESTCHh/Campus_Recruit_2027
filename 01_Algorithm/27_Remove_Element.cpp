// /*
//  * @题目: LeetCode 704. 二分查找
//  * @时间复杂度: O(log n) - 每次砍掉一半的搜索空间
//  * @空间复杂度: O(1)
//  */

// #include <vector>
// using namespace std;

// // ==========================================
// // 写法一：左闭右闭区间 [left, right]
// // ==========================================

// // 🎯 核心思想：循环不变量（Loop Invariant）。寻找目标的有效势力范围是两头都包含的闭区间，每一次缩小边界，都坚决不能把“已知不是目标的值”放进接下来的区间里。

// // 左指针 (Left Pointer)： 像一个不断向右推进的左侧界碑，代表搜索范围的起始点（包含）。
// // 右指针 (Right Pointer)： 像一个不断向左推进的右侧界碑，代表搜索范围的结束点（包含）。
// // 中间指针 (Mid Pointer)： 像一个精准的探雷器，每次都直接降落到当前左右界碑的正中间试探。

// // C++
// class SolutionClosed {
// public:
//     int search(vector<int>& nums, int target) {
//         int left = 0; // 左界碑，立在数组第一个元素
//         int right = nums.size() - 1; // 右界碑，立在数组最后一个元素（可以被取到）

//         // 因为 left == right 在闭区间 [left, right] 中是有意义的（比如区间缩到只剩1个元素），所以用 <=
//         while (left <= right) {
//             int mid = left + (right - left) / 2; // 探雷器降落（防整型溢出写法）

//             if (nums[mid] < target) { 
//                 // 探雷发现中间值比目标小，说明目标在右半边
//                 // 因为 mid 肯定不是目标，左界碑坚决越过 mid，推到 mid + 1
//                 left = mid + 1; 
//             }
//             else if (nums[mid] > target) { 
//                 // 探雷发现中间值比目标大，说明目标在左半边
//                 // 同理，右界碑坚决越过 mid，推到 mid - 1
//                 right = mid - 1; 
//             }
//             else {
//                 // 探雷器正好踩中目标，直接返回所在位置！
//                 return mid; 
//             }
//         }
//         return -1; // 两个界碑错开了都没找到，说明目标不存在
//     }
// };


// // ==========================================
// // 写法二：左闭右开区间 [left, right)
// // ==========================================

// // 🎯 核心思想：右边界是一个“只可远观不可亵玩”的隔离网。寻找目标的有效势力范围包含左边界，但绝对不包含右边界。

// // 左指针 (Left Pointer)： 像一个不断向右推进的左侧界碑（包含）。
// // 右指针 (Right Pointer)： 像一个不可触碰的隔离网（开边界，它指向的位置不属于我们的搜索范围）。

// // C++
// class SolutionOpen {
// public:
//     int search(vector<int>& nums, int target) {
//         int left = 0; // 左界碑，立在数组第一个元素
//         int right = nums.size(); // 🚨 右隔离网，立在数组外的第一个位置（取不到）

//         // 因为 left == right 在 [left, right) 中是无意义的空区间，所以只能用 <
//         while (left < right) {
//             int mid = left + (right - left) / 2; // 探雷器降落

//             if (nums[mid] < target) { 
//                 // 探雷发现中间值比目标小，说明目标在右半边
//                 // 左边界是闭合的，所以左界碑依然坚决推到 mid + 1
//                 left = mid + 1; 
//             }
//             else if (nums[mid] > target) { 
//                 // 探雷发现中间值比目标大，说明目标在左半边
//                 // 🚨 极其关键：因为右侧是开区间（本身就不包含），
//                 // 为了把 mid 排除在下一次搜索之外，右隔离网直接移到 mid 的位置即可。
//                 right = mid; 
//             }
//             else {
//                 return mid; // 踩中目标
//             }
//         }
//         return -1; 
//     }
// };