// 文件作用：
// 学习单链表：
//     头部插入
//     中间插入
//     删除中间节点
//     删除头节点
//     new / delete
//     指针修改顺序
#include <iostream>

using namespace std;

// ListNode 表示单链表中的一个节点。
//
// value：
// 保存当前节点的数据。
//
// next：
// 保存下一个ListNode节点的地址。
// 如果next == nullptr，说明当前节点是链表最后一个节点。
struct ListNode
{
    int value;
    ListNode* next;
};

// printList：
//
// 从head开始沿着next不断向后遍历，
// 直到遇到nullptr。
void printList(ListNode* head)
{
    ListNode* current = head;

    while (current != nullptr)
    {
        cout << current->value << ' ';

        current = current->next;
    }

    cout << endl;
}

int main()
{
    // --------------------------------------------------
    // 1. 创建初始链表
    //
    // head
    //  ↓
    // 10 -> 20 -> 30 -> nullptr
    // --------------------------------------------------

    ListNode* node1 = new ListNode();
    ListNode* node2 = new ListNode();
    ListNode* node3 = new ListNode();

    node1->value = 10;
    node1->next = node2;

    node2->value = 20;
    node2->next = node3;

    node3->value = 30;
    node3->next = nullptr;

    // 普通链表中：
    //
    // head直接指向第一个真实数据节点。
    ListNode* head = node1;

    cout << "=== original ===" << endl;
    printList(head);

    // --------------------------------------------------
    // 2. 头部插入5
    //
    // 原：
    //
    // 10 -> 20 -> 30
    //
    // 目标：
    //
    // 5 -> 10 -> 20 -> 30
    // --------------------------------------------------

    ListNode* newHead = new ListNode();

    newHead->value = 5;

    // 一定先保存原来的head：
    //
    // newHead
    //    ↓
    //    5 -> 原head
    newHead->next = head;

    // 再让head指向新的第一个节点。
    head = newHead;

    cout << endl;
    cout << "=== after insert at head ===" << endl;
    printList(head);

    // --------------------------------------------------
    // 3. 在10和20之间插入15
    //
    // 当前：
    //
    // 5 -> 10 -> 20 -> 30
    //
    // 我们已经知道：
    //
    // node1就是值为10的节点。
    //
    // 所以：
    //
    // prev = node1
    // --------------------------------------------------

    ListNode* prev = node1;

    ListNode* newNode = new ListNode();

    newNode->value = 15;

    // 先让新节点保存原来的后继节点地址。
    //
    // 15 -> 20
    newNode->next = prev->next;

    // 再让10指向15。
    //
    // 10 -> 15
    prev->next = newNode;

    cout << endl;
    cout << "=== after insert 15 ===" << endl;
    printList(head);

    // --------------------------------------------------
    // 4. 删除20
    //
    // 当前：
    //
    // 5 -> 10 -> 15 -> 20 -> 30
    //
    // newNode：
    // 指向15。
    //
    // node2：
    // 指向20。
    //
    // 因此：
    //
    // previous = newNode
    // current  = node2
    // --------------------------------------------------

    ListNode* previous = newNode;
    ListNode* current = node2;

    // 删除节点之前，
    // 先把链表重新连接：
    //
    // 15 -> 30
    previous->next = current->next;

    // 20已经不再属于链表，
    // 现在才能安全释放它。
    delete current;

    // node2原来保存的是已经释放的地址。
    //
    // 为了避免以后误用悬空指针，
    // 学习阶段主动设为nullptr。
    node2 = nullptr;

    cout << endl;
    cout << "=== after delete 20 ===" << endl;
    printList(head);

    // --------------------------------------------------
    // 5. 删除头节点5
    //
    // 当前：
    //
    // head
    //  ↓
    // 5 -> 10 -> 15 -> 30
    // --------------------------------------------------

    // 必须先保存旧head，
    // 否则head移动以后就失去旧节点地址。
    ListNode* oldHead = head;

    // head移动到下一个节点10。
    head = head->next;

    // 原来的5已经离开链表。
    delete oldHead;

    oldHead = nullptr;

    cout << endl;
    cout << "=== after delete head ===" << endl;
    printList(head);

    // --------------------------------------------------
    // 6. 程序结束前释放剩余节点
    //
    // 当前链表：
    //
    // 10 -> 15 -> 30
    //
    // 所有节点都是通过new创建的，
    // 因此需要delete释放。
    // --------------------------------------------------

    current = head;

    while (current != nullptr)
    {
        // 删除current之前，
        // 必须先保存下一个节点地址。
        //
        // 否则delete current之后，
        // current->next就不能再安全访问。
        ListNode* nextNode = current->next;

        delete current;

        current = nextNode;
    }

    // 链表节点全部释放以后，
    // head不能继续保存已经释放的地址。
    head = nullptr;

    return 0;
}