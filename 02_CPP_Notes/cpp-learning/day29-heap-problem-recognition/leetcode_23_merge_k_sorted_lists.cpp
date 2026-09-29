// /**
//  * Definition for singly-linked list.
//  * struct ListNode {
//  *     int val;
//  *     ListNode *next;
//  *     ListNode() : val(0), next(nullptr) {}
//  *     ListNode(int x) : val(x), next(nullptr) {}
//  *     ListNode(int x, ListNode *next) : val(x), next(next) {}
//  * };
//  */
// class Solution {
// private:
//     struct Compare {
//         bool operator()(
//             ListNode* a,
//             ListNode* b
//         ) const {
//             return a->val > b->val;
//         }
//     };

// public:
//     ListNode* mergeKLists(
//         vector<ListNode*>& lists
//     ) {
//         if(static_cast<int>(lists.size()) == 0) return nullptr;
//         priority_queue<
//             ListNode*,
//             vector<ListNode*>,
//             Compare
//         > minHeap;

//         // 1. 初始化Heap
//         for (ListNode* Node : lists){
//             if(!Node) continue;
//             minHeap.push(Node);
//         }
//         // 2. dummy + tail
//         ListNode dummy;
//         ListNode* tail = &dummy;
//         // 3. while Heap不空
//         while(!minHeap.empty()){
//             ListNode* topNode = minHeap.top(); 
//             minHeap.pop(); 
//             tail->next = topNode; 
//             if(topNode && topNode -> next){
//                 minHeap.push(topNode -> next);
//             }
//             tail = tail->next; 
//         }
//         return dummy.next;
//         //    pop最小节点
//         //    接入结果链表
//         //    如果存在next，push next

//         // 4. return
//     }
// };