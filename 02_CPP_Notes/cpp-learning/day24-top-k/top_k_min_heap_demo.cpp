// 文件职责：
// 通过大小固定为K的Min Heap
// 动态维护当前Top K最大元素
#include <functional>
#include <iostream>
#include <queue>
#include <vector>

using namespace std;

int main()
{
    vector<int> nums{
        3, 2, 1, 5, 6, 4
    };

    int k = 3;

    // Min Heap。
    //
    // Heap中始终只保留当前最大的K个元素。
    //
    // top()保存的是：
    // “当前Top K中最小的那个元素”。
    //
    // 它就是当前进入Top K的门槛。
    priority_queue<
        int,
        vector<int>,
        greater<int>
    > minHeap;

    int replacementCount = 0;
    for (int value : nums)
    {
        // Heap还没有K个元素时，
        // 当前元素直接加入。
        if (
            static_cast<int>(
                minHeap.size()
            ) < k
        ) {
            minHeap.push(value);
        }
        else if (
            value > minHeap.top()
        ) {
            // 当前Heap已经有K个元素。
            //
            // 如果新元素比Top K中最小的元素还大，
            // 那么新元素应该进入Top K，
            // 原来的最小元素被淘汰。
            minHeap.pop();
            minHeap.push(value);
            ++replacementCount;
        }

        cout << "read "
             << value
             << ", current threshold = "
             << minHeap.top()
             << endl;
    }

    cout << endl;

    cout << "3rd largest = "
         << minHeap.top()
         << endl;

    cout << "replacement count = "
     << replacementCount
     << endl;
    return 0;
}
// // 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

 

// // 示例 1:

// // 输入: [3,2,1,5,6,4], k = 2
// // 输出: 5
// // 示例 2:

// // 输入: [3,2,3,1,2,4,5,5,6], k = 4
// // 输出: 4
 

// // 提示：

// // 1 <= k <= nums.length <= 105
// // -104 <= nums[i] <= 104
// // 输入
// // nums =
// // [-1,2,0]
// // k =
// // 2

// // 添加到测试用例
// // 输出
// // 2
// // 预期结果
// // 0
// class Solution {
// public:
//     int findKthLargest(vector<int>& nums, int k) {
//         int n = static_cast<int> (nums.size());
//         priority_queue<int, vector<int>, greater<int>> mh;
//         for(int i = 0; i < n; i++){
//             if(i < k){
//                 mh.push(nums[i]);
//                 continue;
//             } 
//             if(mh.top() < nums[i]){
//                 mh.pop();
//                 mh.push(nums[i]);
//             }
//         }
//         return mh.top();
//     }
// };