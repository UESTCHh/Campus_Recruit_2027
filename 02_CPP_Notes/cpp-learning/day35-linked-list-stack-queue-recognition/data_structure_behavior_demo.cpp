// 这个文件的职责是：用相同的 1、2、3，观察链表的头插操作、Stack 和 Queue 的行为差别。
#include <iostream>
#include <stack>
#include <queue>

using namespace std;

// 单向链表节点：
// value 保存数据。
// next 指向下一个节点。
struct ListNode {
    int value;
    ListNode* next;
};

int main() {
    // ==================================
    // 1. Linked List：演示头插法
    // ==================================

    ListNode* head = nullptr;

    // 每次新节点放到链表头部。
    for (int x : {1, 2, 3}) {
        head = new ListNode{x, head};
    }

    cout << "Linked List (push front): ";

    ListNode* cur = head;

    while (cur != nullptr) {
        cout << cur->value << ' ';
        cur = cur->next;
    }

    cout << '\n';

    // 释放new创建的节点，避免内存泄漏。
    while (head != nullptr) {
        ListNode* nextNode = head->next;
        delete head;
        head = nextNode;
    }

    // ==================================
    // 2. Stack：后进先出
    // ==================================

    stack<int> st;

    for (int x : {1, 2, 3}) {
        st.push(x);
    }

    cout << "Stack pop order: ";

    while (!st.empty()) {
        cout << st.top() << ' ';
        st.pop();
    }

    cout << '\n';

    // ==================================
    // 3. Queue：先进先出
    // ==================================

    queue<int> q;

    for (int x : {1, 2, 3}) {
        q.push(x);
    }

    cout << "Queue pop order: ";

    while (!q.empty()) {
        cout << q.front() << ' ';
        q.pop();
    }

    cout << '\n';

    return 0;
}