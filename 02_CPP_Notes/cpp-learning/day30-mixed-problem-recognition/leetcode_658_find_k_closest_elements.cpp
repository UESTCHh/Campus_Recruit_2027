// // class Solution {
// // public:
// //     vector<int> findClosestElements(vector<int>& arr, int k, int x) {
// //         int n = static_cast<int> (arr.size());
// //         vector<int> answer;
// //         int left = 0;
// //         int right = n - 1;
// //         int index = -1;
// //         while(left <= right){
// //             int mid = left + (right - left) / 2;
// //             if(arr[mid] < x){
// //                 left = mid + 1;
// //             }
// //             else if(arr[mid] > x){
// //                 right = mid - 1;
// //             }
// //             else{
// //                 index = mid;
// //                 break;
// //             }
// //         }
// //         if(left == n) index = n - 1;
// //         else if(right == -1) index = 0;
// //         else if(right < left) index = left;
        
// //         left = index - 1;
// //         right = index;
// //         for(int i = 0; i < k; i++){
// //             if(left < 0){
// //                 right++;
// //             }
// //             else if(right > n - 1){
// //                  left--;
// //             }
// //             else if(x - arr[left] <= arr[right] - x){
// //                  left--;
// //             }
// //             else{
// //                  right++;
// //             }
// //         }
// //         for(int i = left + 1; i < right; i++){
// //             answer.push_back(arr[i]);
// //         }
// //         return answer;
// //     }
// // };
// class Solution {
// public:
//     vector<int> findClosestElements(
//         vector<int>& arr,
//         int k,
//         int x
//     ) {
//         int n =
//             static_cast<int>(
//                 arr.size()
//             );

//         // 找第一个 >= x 的位置。
//         int low = 0;
//         int high = n;

//         while (low < high) {
//             int mid =
//                 low + (high - low) / 2;

//             if (arr[mid] < x) {
//                 low = mid + 1;
//             }
//             else {
//                 high = mid;
//             }
//         }

//         int right = low;
//         int left = right - 1;

//         // 从x附近向两边选出k个元素。
//         for (int i = 0; i < k; ++i) {
//             if (left < 0) {
//                 ++right;
//             }
//             else if (right >= n) {
//                 --left;
//             }
//             else if (
//                 x - arr[left]
//                 <=
//                 arr[right] - x
//             ) {
//                 // 距离相同时选择较小元素，
//                 // 即选择左侧。
//                 --left;
//             }
//             else {
//                 ++right;
//             }
//         }

//         vector<int> answer;

//         for (
//             int i = left + 1;
//             i < right;
//             ++i
//         ) {
//             answer.push_back(
//                 arr[i]
//             );
//         }

//         return answer;
//     }
// };