// #include <iostream>

// using namespace std;

// struct ListNode{
//     int value;

//     ListNode* next;
// };

// int main(){
//     ListNode* head = new ListNode();
//     ListNode* node1 = new ListNode();
//     ListNode* node2 = new ListNode();
//     ListNode* node3 = new ListNode();

//     node1 -> value = 10;
//     node1 -> next = node2;

//     node2 -> value = 20;
//     node2 -> next = node3;

//     node3 -> value = 30;
//     node3 -> next = nullptr;

//     head -> next = node1;
//     ListNode* cur = head;

//     while(cur -> next != nullptr){
//         cout << cur -> next -> value << " ";
//         cur = cur -> next;
//     }
//     return 0;
// }
#include <iostream>

using namespace std;


struct ListNode
{
    int value;
    ListNode* next;
};


int main()
{
    // 创建三个节点
    ListNode* node1 = new ListNode;
    ListNode* node2 = new ListNode;
    ListNode* node3 = new ListNode;


    // 设置节点数据
    node1->value = 10;
    node2->value = 20;
    node3->value = 30;


    // 连接节点
    node1->next = node2;
    node2->next = node3;
    node3->next = nullptr;


    // head指向第一个节点
    ListNode* head = node1;


    return 0;
}