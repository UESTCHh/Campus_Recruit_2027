// 文件作用：
// 学习Dummy Node
// 统一删除头节点和普通节点
// 删除第一个value等于target的节点
#include <iostream>

using namespace std;

struct ListNode
{
    int value;
    ListNode* next;
};

// 从head开始遍历链表并输出。
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

// 删除链表中第一个value等于target的节点。
//
// 这里使用Dummy Node的目的：
//
// 让原来的头节点前面也存在一个prev，
// 从而不需要单独判断：
//
// if (head->value == target)
//
// 头节点和中间节点可以使用同一套删除逻辑。
ListNode* removeFirstValue(
    ListNode* head,
    int target
)
{
    // Dummy Node本身不保存真正的业务数据。
    //
    // 它只是放在真正head之前，
    // 为原来的第一个节点人工创建一个前驱节点。
    ListNode* dummy = new ListNode();

    dummy->value = 0;
    dummy->next = head;

    // previous始终指向current前面的节点。
    ListNode* previous = dummy;

    // current从真正的第一个数据节点开始。
    ListNode* current = head;

    while (current != nullptr)
    {
        if (current->value == target)
        {
            // 无论current是不是原来的head，
            // 都可以统一使用：
            //
            // previous->next = current->next
            //
            // 如果删除的是原head：
            //
            // previous == dummy
            //
            // 如果删除的是中间节点：
            //
            // previous == 普通数据节点
            previous->next = current->next;

            delete current;

            // 只删除第一个匹配节点，
            // 所以找到以后直接结束循环。
            break;
        }

        previous = current;
        current = current->next;
    }

    // 删除完成以后，
    // 真正链表的新head就是dummy->next。
    //
    // 必须在delete dummy之前先保存，
    // 否则dummy释放以后不能再访问dummy->next。
    ListNode* newHead = dummy->next;

    // Dummy Node只是辅助节点，
    // 它本身也是通过new创建的，
    // 所以最终需要释放。
    delete dummy;

    return newHead;
}

// 释放整个链表。
void deleteList(ListNode*& head)
{
    ListNode* current = head;

    while (current != nullptr)
    {
        // 一定先保存next，
        // 再释放current。
        ListNode* nextNode = current->next;

        delete current;

        current = nextNode;
    }

    head = nullptr;
}

int main()
{
    // 创建：
    //
    // 10 -> 20 -> 30 -> nullptr

    ListNode* node1 = new ListNode();
    ListNode* node2 = new ListNode();
    ListNode* node3 = new ListNode();

    node1->value = 10;
    node1->next = node2;

    node2->value = 20;
    node2->next = node3;

    node3->value = 30;
    node3->next = nullptr;

    ListNode* head = node1;

    cout << "=== original ===" << endl;
    printList(head);

    // ------------------------------------------
    // 第一次：
    //
    // 删除原来的头节点10。
    //
    // Dummy Node会让10也拥有previous。
    // ------------------------------------------
    head = removeFirstValue(
        head,
        10
    );

    cout << endl;
    cout << "=== after remove 10 ===" << endl;
    printList(head);

    // ------------------------------------------
    // 第二次：
    //
    // 删除中间节点20。
    //
    // 使用的仍然是完全相同的removeFirstValue。
    // ------------------------------------------
    head = removeFirstValue(
        head,
        20
    );

    cout << endl;
    cout << "=== after remove 20 ===" << endl;
    printList(head);

    // ------------------------------------------
    // 第三次：
    //
    // 删除不存在的100。
    //
    // 整个链表遍历结束以后找不到target，
    // 链表保持不变。
    // ------------------------------------------
    head = removeFirstValue(
        head,
        100
    );

    cout << endl;
    cout << "=== after remove 100 ===" << endl;
    printList(head);

    // 释放剩余链表。
    deleteList(head);

    return 0;
}