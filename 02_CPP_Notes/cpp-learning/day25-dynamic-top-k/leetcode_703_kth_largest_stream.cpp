// // 职责：
// // 使用对象成员保存Top K状态
// // 解决LeetCode 703
// class KthLargest {
// private:
//     priority_queue<int, vector<int>, greater<int>> minheap;
//     int k_;
// public:
//     KthLargest(int k, vector<int>& nums) {
//         k_ = k;
//         for(int val : nums){
//             add(val);
//         }
//     }
    
//     int add(int val) {
//         if(static_cast<int>(minheap.size()) < k_){
//             minheap.push(val);
//         }
//         else{
//             if(minheap.top() < val){
//                 minheap.pop();
//                 minheap.push(val);
//             }
//         }
//         return minheap.top();
//     }
// };

// /**
//  * Your KthLargest object will be instantiated and called as such:
//  * KthLargest* obj = new KthLargest(k, nums);
//  * int param_1 = obj->add(val);
//  */
