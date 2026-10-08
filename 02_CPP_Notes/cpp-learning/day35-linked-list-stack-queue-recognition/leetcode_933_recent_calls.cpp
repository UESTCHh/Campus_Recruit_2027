// class RecentCounter {
// public:
//     RecentCounter() {
        
//     }

//     int ping(int t) {
//         this->q.push(t);
//         while(q.front() < t - 3000){
//             q.pop();
//         }
//         return q.size();
//     }
// private:
//     queue<int> q;
// };

// /**
//  * Your RecentCounter object will be instantiated and called as such:
//  * RecentCounter* obj = new RecentCounter();
//  * int param_1 = obj->ping(t);
//  */
// Queue 中应该保存什么？是请求的时间 t，还是累计请求次数？
// 1 100 3001 3002
// 输入
// ["RecentCounter","ping","ping","ping","ping","ping"]
// [[],[642],[1849],[4921],[5936],[5957]]
// 输出
// [null,1,2,2,2,3]
// 预期结果
// [null,1,2,1,2,3]

// 当新请求到达时，如何判断队头请求是否已经过期？特别注意区间 [t-3000,t] 的左端点也是有效的。



// 为什么移除所有过期请求以后，可以直接使用 q.size() 作为答案？