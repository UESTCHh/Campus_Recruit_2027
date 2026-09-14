// 文件作用：
// 把第一阶段的手动三轮操作
// 抽象成while循环

// 实现：
// reverseList()

// 验证：
// 任意长度Linked List都可以使用同一算法
#include <iostream>

using namespace std;

// ListNode 表示单链表中的一个节点。
//
// value：当前节点保存的数据。
// next ：指向逻辑上的下一个节点。
struct ListNode
{
    int value;
    ListNode* next;
};


// printList：
//
// 从head开始沿着next遍历整个链表。
//
// 这也是复习：
// Linked List无法像vector一样通过下标随机访问，
// 只能沿next逐个访问节点。
void printList(ListNode* head)
{
    ListNode* current = head;

    while (current != nullptr)
    {
        cout << current->value << ' ';

        current =
            current->next;
    }

    cout << endl;
}


// reverseList：
//
// 使用 Iterative（三指针）方式反转单链表。
//
// 输入：
//
// head
// → 原链表头节点
//
// 返回：
//
// 反转后链表的新头节点。
//
// --------------------------------------------------
// 三个核心指针：
//
// previous
// → 已经完成反转部分的头节点
//
// current
// → 当前正在处理的节点
//
// nextNode
// → 保存原链表中current后面的剩余链表入口
//
// --------------------------------------------------
//
// 每轮固定四步：
//
// 1. Save
//    nextNode = current->next
//
// 2. Reverse
//    current->next = previous
//
// 3. Move previous
//    previous = current
//
// 4. Move current
//    current = nextNode
//
ListNode* reverseList(ListNode* head)
{
    // 一开始：
    //
    // 已反转部分为空。
    ListNode* previous =
        nullptr;

    // 未反转部分从原head开始。
    ListNode* current =
        head;


    // 只要current不为空，
    // 就说明仍然存在尚未处理的节点。
    while (current != nullptr)
    {
        // ------------------------------------------
        // Step 1：保存剩余链表入口
        // ------------------------------------------
        //
        // 等下会覆盖current->next，
        // 所以必须先保存原来的next。
        ListNode* nextNode =
            current->next;


        // ------------------------------------------
        // Step 2：反转当前节点的指针
        // ------------------------------------------
        //
        // 原来：
        //
        // previous    current
        //     ↓          ↓
        // ...          node -> next
        //
        // 修改后：
        //
        // node -> previous
        current->next =
            previous;


        // ------------------------------------------
        // Step 3：扩大已反转部分
        // ------------------------------------------
        //
        // 当前节点已经完成反转，
        // 所以它现在成为已反转链表的新头节点。
        previous =
            current;


        // ------------------------------------------
        // Step 4：进入下一待处理节点
        // ------------------------------------------
        //
        // nextNode保存着刚才提前留下的
        // 原链表剩余部分入口。
        current =
            nextNode;
    }


    // 循环结束：
    //
    // current == nullptr
    //
    // 说明原链表所有节点都已经完成反转。
    //
    // previous正好指向反转后的新头节点。
    return previous;
}


// deleteList：
//
// 释放链表中的所有动态节点。
//
// 再次复习昨天学习的内存安全：
//
// delete current之前，
// 必须先保存current->next。
void deleteList(ListNode*& head)
{
    ListNode* current =
        head;

    while (current != nullptr)
    {
        // 先保存下一个节点。
        ListNode* nextNode =
            current->next;

        // 再释放当前节点。
        delete current;

        // 最后移动到之前保存好的节点。
        current =
            nextNode;
    }

    // 所有节点释放以后，
    // 不让head继续保存已经无效的旧地址。
    head = nullptr;
}


int main()
{
    // --------------------------------------------------
    // 1. 创建链表
    // --------------------------------------------------
    //
    // head
    //  ↓
    // 10 -> 20 -> 30 -> 40 -> nullptr
    ListNode* node10 =
        new ListNode{10, nullptr};

    ListNode* node20 =
        new ListNode{20, nullptr};

    ListNode* node30 =
        new ListNode{30, nullptr};

    ListNode* node40 =
        new ListNode{40, nullptr};

    node10->next =
        node20;

    node20->next =
        node30;

    node30->next =
        node40;


    ListNode* head =
        node10;


    // --------------------------------------------------
    // 2. 打印反转前链表
    // --------------------------------------------------
    cout << "=== original list ==="
         << endl;

    printList(head);


    // --------------------------------------------------
    // 3. 反转链表
    // --------------------------------------------------
    //
    // reverseList返回新的head，
    // 所以必须重新赋值。
    head =
        reverseList(head);


    // --------------------------------------------------
    // 4. 打印反转后的链表
    // --------------------------------------------------
    cout << endl;

    cout << "=== reversed list ==="
         << endl;

    printList(head);


    // --------------------------------------------------
    // 5. 再反转一次
    // --------------------------------------------------
    //
    // 用于验证：
    //
    // reverseList不是针对某个固定方向写死的，
    // 而是一个通用的链表反转操作。
    head =
        reverseList(head);

    cout << endl;

    cout << "=== reversed again ==="
         << endl;

    printList(head);


    // --------------------------------------------------
    // 6. 释放动态内存
    // --------------------------------------------------
    deleteList(head);

    cout << endl;

    if (head == nullptr)
    {
        cout << "head = nullptr after delete"
            << endl;
    }


    // --------------------------------------------------
    // 7. Empty List Test
    // --------------------------------------------------
    //
    // 空链表：
    //
    // nullptr
    //
    // reverseList内部：
    //
    // previous = nullptr
    // current  = nullptr
    //
    // while不会进入，
    // 最终直接返回previous，也就是nullptr。
    ListNode* emptyHead =
        nullptr;

    emptyHead =
        reverseList(emptyHead);

    cout << endl;

    cout << "=== empty list test ==="
        << endl;

    if (emptyHead == nullptr)
    {
        cout << "empty list remains nullptr"
            << endl;
    }


    // --------------------------------------------------
    // 8. Single Node Test
    // --------------------------------------------------
    //
    // 单节点链表：
    //
    // 100 -> nullptr
    //
    // 反转以后应该仍然：
    //
    // 100 -> nullptr
    ListNode* singleHead =
        new ListNode{100, nullptr};

    cout << endl;

    cout << "=== single node before reverse ==="
        << endl;

    printList(singleHead);

    singleHead =
        reverseList(singleHead);

    cout << endl;

    cout << "=== single node after reverse ==="
        << endl;

    printList(singleHead);


    // 释放单节点。
    deleteList(singleHead);

    return 0;
}