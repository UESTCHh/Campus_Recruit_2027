// 文件职责：
// 训练：
// lower boundary
// +
// upper boundary
// +
// 两个边界组合解决真实LeetCode
// class Solution {
// public:
//     vector<int> searchRange(vector<int>& nums, int target) {
//         int n = static_cast<int>(nums.size());
//         if(n == 0) return {-1, -1};

//         vector<int> answer;
//         int left1 = 0;
//         int right1 = n;
//         while(left1 < right1){
//             int mid = left1 + (right1 - left1) / 2;
//             if(nums[mid] < target){
//                 left1 = mid + 1;
//             }
//             else{
//                 right1 = mid;
//             }
//         }
//         if(left1 == n || nums[left1] != target) return {-1, -1};
//         int left2 = 0;
//         int right2 = n;
//         while(left2 < right2){
//             int mid = left2 + (right2 - left2) / 2;
//             if(nums[mid] <= target){
//                 left2 = mid + 1;
//             }
//             else{
//                 right2 = mid;
//             }
//         }
//         answer.push_back(left1, left2);
//         return answer;
//     }
// };