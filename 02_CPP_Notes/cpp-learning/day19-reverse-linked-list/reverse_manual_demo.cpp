// 文件作用：
// 不使用循环
// 手动执行三轮链表反转
// 观察 previous / current / nextNode
// 真正理解三个指针如何移动
#include <iostream>

using namespace std;

// ListNode 表示单链表中的一个节点。
//
// value：
// 当前节点保存的数据。
//
// next：
// 指向逻辑上的下一个节点。
struct ListNode
{
    int value;
    ListNode* next;
};

int main()
{
    // --------------------------------------------------
    // 1. 创建原链表
    // --------------------------------------------------
    //
    // head
    //  ↓
    // 10 -> 20 -> 30 -> nullptr
    //
    // 今天仍然手动创建节点，
    // 这样可以把注意力集中在指针变化上。
    ListNode* node10 =
        new ListNode{10, nullptr};

    ListNode* node20 =
        new ListNode{20, nullptr};

    ListNode* node30 =
        new ListNode{30, nullptr};

    node10->next = node20;
    node20->next = node30;

    ListNode* head = node10;

    cout << "=== original list ==="
         << endl;

    ListNode* printCurrent = head;

    while (printCurrent != nullptr)
    {
        cout << printCurrent->value << ' ';
        printCurrent = printCurrent->next;
    }

    cout << endl;


    // --------------------------------------------------
    // 2. 初始化 Reverse Linked List 的两个核心指针
    // --------------------------------------------------
    //
    // previous：
    //
    // 已经完成反转部分的头节点。
    //
    // 一开始一个节点都还没处理，
    // 所以已反转部分为空：
    //
    // previous -> nullptr
    ListNode* previous = nullptr;


    // current：
    //
    // 当前正在处理的节点。
    //
    // 一开始从原链表头节点开始。
    ListNode* current = head;


    // ==================================================
    // 第一轮
    // ==================================================
    //
    // 当前：
    //
    // previous
    //    ↓
    // nullptr
    //
    // current
    //    ↓
    // 10 -> 20 -> 30 -> nullptr


    // --------------------------------------------------
    // Step 1：保存原来的next
    // --------------------------------------------------
    //
    // current->next 现在保存着：
    //
    // node20
    //
    // 等下我们马上要改 current->next，
    // 因此必须在覆盖它之前保存原值。
    ListNode* nextNode = current->next;


    // --------------------------------------------------
    // Step 2：反转当前节点的next
    // --------------------------------------------------
    //
    // 原来：
    //
    // 10 -> 20
    //
    // 现在：
    //
    // 10 -> nullptr
    //
    // 因为：
    //
    // previous == nullptr
    current->next = previous;


    // --------------------------------------------------
    // Step 3：previous向前移动
    // --------------------------------------------------
    //
    // previous现在应该指向：
    //
    // 已经反转好的：
    //
    // 10 -> nullptr
    previous = current;


    // --------------------------------------------------
    // Step 4：current进入剩余链表
    // --------------------------------------------------
    //
    // nextNode之前保存了node20。
    //
    // 所以现在：
    //
    // current -> 20 -> 30 -> nullptr
    current = nextNode;


    cout << endl;
    cout << "=== after round 1 ==="
         << endl;

    cout << "previous = "
         << previous->value
         << endl;

    cout << "current = "
         << current->value
         << endl;


    // ==================================================
    // 第二轮
    // ==================================================
    //
    // previous
    //    ↓
    // 10 -> nullptr
    //
    // current
    //    ↓
    // 20 -> 30 -> nullptr


    // 保存：
    //
    // node30
    nextNode = current->next;


    // 原来：
    //
    // 20 -> 30
    //
    // 改成：
    //
    // 20 -> 10
    current->next = previous;


    // 已反转链表变成：
    //
    // 20 -> 10 -> nullptr
    previous = current;


    // 未反转链表剩下：
    //
    // 30 -> nullptr
    current = nextNode;


    cout << endl;
    cout << "=== after round 2 ==="
         << endl;

    cout << "previous = "
         << previous->value
         << endl;

    cout << "current = "
         << current->value
         << endl;


    // ==================================================
    // 第三轮
    // ==================================================
    //
    // previous
    //    ↓
    // 20 -> 10 -> nullptr
    //
    // current
    //    ↓
    // 30 -> nullptr


    // node30原来的next就是nullptr，
    // 所以这里：
    //
    // nextNode == nullptr
    nextNode = current->next;


    // 反转：
    //
    // 30 -> 20
    current->next = previous;


    // 已反转部分：
    //
    // 30 -> 20 -> 10 -> nullptr
    previous = current;


    // current：
    //
    // nullptr
    //
    // 说明已经没有未处理节点。
    current = nextNode;


    cout << endl;
    cout << "=== after round 3 ==="
         << endl;

    cout << "previous = "
         << previous->value
         << endl;

    if (current == nullptr)
    {
        cout << "current = nullptr"
             << endl;
    }


    // --------------------------------------------------
    // 3. previous就是新的头节点
    // --------------------------------------------------
    head = previous;


    cout << endl;
    cout << "=== reversed list ==="
         << endl;

    printCurrent = head;

    while (printCurrent != nullptr)
    {
        cout << printCurrent->value << ' ';
        printCurrent = printCurrent->next;
    }

    cout << endl;


    // --------------------------------------------------
    // 4. 释放动态内存
    // --------------------------------------------------
    //
    // 注意：
    //
    // 链表已经反转，
    // 所以现在顺序是：
    //
    // 30 -> 20 -> 10 -> nullptr
    //
    // 释放时依然必须：
    //
    // 先保存next
    // 再delete current
    // 再移动current
    //
    // 这正好复习昨天的内存安全知识。
    current = head;

    while (current != nullptr)
    {
        ListNode* nodeToDelete =
            current;

        ListNode* nextToDelete =
            current->next;

        delete nodeToDelete;

        current = nextToDelete;
    }

    head = nullptr;

    return 0;
}