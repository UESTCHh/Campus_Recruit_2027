// 文件作用：观察链表插入和删除前后的节点连接关系。

#include <iostream>

using namespace std;

// 单向链表节点
struct ListNode {
    int val;
    ListNode* next;
};

// 从头节点开始遍历链表
void printList(const ListNode* head) {
    const ListNode* cur = head;

    while (cur != nullptr) {
        cout << cur->val;

        if (cur->next != nullptr) {
            cout << " -> ";
        }

        cur = cur->next;
    }

    cout << '\n';
}

int main() {
    // 创建链表：1 -> 2 -> 3
    ListNode* node1 = new ListNode{1, nullptr};
    ListNode* node2 = new ListNode{2, nullptr};
    ListNode* node3 = new ListNode{3, nullptr};

    node1->next = node2;
    node2->next = node3;

    cout << "Original: ";
    printList(node1);

    // 在节点2后插入节点4
    ListNode* node4 = new ListNode{4, nullptr};

    // 先连接节点4和原来的节点3
    node4->next = node2->next;

    // 再让节点2指向节点4
    node2->next = node4;

    cout << "After insert: ";
    printList(node1);

    // 删除节点4
    node2->next = node4->next;
    delete node4;

    cout << "After delete: ";
    printList(node1);

    // 释放剩余链表
    ListNode* head = node1;

    while (head != nullptr) {
        ListNode* nextNode = head->next;
        delete head;
        head = nextNode;
    }

    return 0;
}
