// 🎯 核心思想：循环不变量 (Loop Invariant)
// 二分查找看似简单，大厂面试手撕时却极其容易写出死循环。核心痛点在于区间的定义。
// 区间的定义就是我们在寻找目标值时的“有效势力范围”。在整个 while 循环中，必须始终坚持根据最初定义的区间属性来更新边界，这就是所谓的“循环不变量”原则。

// 写法一：左闭右闭区间 [left, right]
// 这是最符合人类直觉的写法。区间包含 left 和 right，意味着 nums[right] 是一个可以被取到并且需要被判断的合法元素。

// 💻 核心代码
// C++

// class Solution {
// public:
//     int search(vector<int>& nums, int target) {
//         int left = 0;
//         int right = nums.size() - 1; // 定义 target 在左闭右闭的区间里，[left, right]

//         // 因为 left == right 在区间 [left, right] 中是有意义的，所以用 <=
//         while (left <= right) {
//             // 防溢出写法，等同于 (left + right)/2
//             int mid = left + (right - left) / 2; 

//             if (nums[mid] < target) {
//                 // target 大于中间值，说明目标在右半边
//                 // 因为 mid 已经判断过了，下一次搜索范围不应包含 mid，所以是 mid + 1
//                 left = mid + 1;
//             } 
//             else if (nums[mid] > target) {
//                 // target 小于中间值，说明目标在左半边
//                 // 同理，mid 已经判断过了，右边界向左收缩到 mid - 1
//                 right = mid - 1;
//             } 
//             else {
//                 // 找到目标，直接返回索引
//                 return mid;
//             }
//         }
//         return -1; // 整个区间遍历完都没找到
//     }
// };
// 💡 面试核心考点：

// 为什么是 <= ？ 因为当 left == right 时（比如区间只剩最后一个元素 [3, 3]），这个元素依然需要被拉出来和 target 比较，如果不加 =，这最后一个元素就被漏掉了。

// 为什么是 right = mid - 1？ 因为 nums[mid] 已经被明确查过不是 target 了。既然是闭区间，我们就坚决不能把已知错误的值包含进下一次的搜索范围里。

// 防溢出细节： 强烈建议写 left + (right - left) / 2，而不是 (left + right) / 2。如果 left 和 right 都非常大（接近 INT_MAX），直接相加会导致整型溢出，这也是 C++ 面试极其爱扣的底层细节！

// 写法二：左闭右开区间 [left, right)
// 这种写法在 C++ STL 的源码（如 std::lower_bound）中非常常见。right 只是一个边界标兵，它不包含在我们的搜索范围内。

// 💻 核心代码
// C++

// class Solution {
// public:
//     int search(vector<int>& nums, int target) {
//         int left = 0;
//         int right = nums.size(); // 定义 target 在左闭右开的区间里，[left, right)

//         // 因为 left == right 在 [left, right) 中是没有意义的（比如 [1, 1) 是空区间），所以用 <
//         while (left < right) {
//             int mid = left + (right - left) / 2;

//             if (nums[mid] < target) {
//                 // 目标在右半边
//                 // mid 已经查过了，左边界是闭区间，所以可以放心地 +1
//                 left = mid + 1;
//             } 
//             else if (nums[mid] > target) {
//                 // 目标在左半边
//                 // 重点：右边界是开区间，本身就不包含在搜索范围内。
//                 // 既然 mid 查过了，下一次我们搜索到 mid 为止（但不包含 mid），所以直接写 mid
//                 right = mid;
//             } 
//             else {
//                 return mid;
//             }
//         }
//         return -1;
//     }
// };
// 💡 面试核心考点：

// 为什么是 < ？ 如果区间变成了 [3, 3)，这是一个非法的空区间。所以在开区间逻辑下，left 永远不可能等于 right，循环条件只能是 <。

// 为什么是 right = mid？ 这是最容易错的地方！因为我们的右边界是开区间（不包含）。我们要把 mid 排除在下一次搜索之外，只需让下一个区间的开边界等于 mid 即可（即搜索到 mid 的前一个元素停止）。如果写成 right = mid - 1，你就会漏掉 mid - 1 这个无辜的元素！