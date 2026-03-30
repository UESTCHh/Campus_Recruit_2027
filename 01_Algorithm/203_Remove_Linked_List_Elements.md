# 算法实战复盘：移除链表元素 (LeetCode 203)

> **打卡日期**：2026-03-30 (Day 4)
> **核心主题**：虚拟头节点 (Dummy Node)、指针操作、C++ 内存泄漏防范。

---

## 📝 一、 题目描述与核心要求

* **题目内容：** 给你一个链表的头节点 `head` 和一个整数 `val` ，请你删除链表中所有满足 `Node.val == val` 的节点，并返回新的头节点。
* **核心痛点：** 链表的节点删除操作是 `cur->next = cur->next->next;`。这要求我们必须站在**目标节点的前一个节点**上才能完成删除。但是，**真实的头节点没有“前一个节点”**！如果不做特殊处理，就必须单独写一段极其丑陋的 `while` 循环来专门处理头节点被删除的情况。

---

## 📊 二、 复杂度深度剖析

### 1. 时间复杂度：$O(n)$
**推导过程：** 假设链表有 $N$ 个节点。我们需要用一个 `cur` 指针从头到尾遍历一次链表，去检查每一个节点的值是否等于 `val`。每个节点最多被访问一次，执行的删除和重新指向操作都是常数时间 $O(1)$ 的底层指针赋值。因此总体时间复杂度为绝对线性的 $O(n)$。

### 2. 空间复杂度：$O(1)$
**推导过程：** 在整个遍历和删除的过程中，我们只在栈内存中开辟了三个额外的指针变量（`dummyHead`、`cur`、`tmp`）。无论链表有 10 个节点还是 10 万个节点，额外消耗的内存永远是固定大小的几个指针，因此空间复杂度为 $O(1)$。

---

## 🎯 三、 核心思想：上帝视角的“虚拟头节点”



为了填平“头节点没有前驱”这个坑，我们人为地 `new` 一个**假节点（Dummy Head）**，并让它的 `next` 指向真实的 `head`。
这样一来，原本的真实头节点，就变成了 Dummy Head 的下一个节点。**整个链表的所有真实节点，统统都有了前驱节点！** 删除逻辑实现了史诗级的统一。

---

## 💻 四、 核心代码与极其严谨的内存释放

```cpp
/**
 * Definition for singly-linked list.
 * struct ListNode {
 * int val;
 * ListNode *next;
 * ListNode() : val(0), next(nullptr) {}
 * ListNode(int x) : val(x), next(nullptr) {}
 * ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* removeElements(ListNode* head, int val) {
        // 1. 设置虚拟头节点，其 next 指向真实的 head
        ListNode* dummyHead = new ListNode(0); 
        dummyHead->next = head;

        // 2. cur 指针作为遍历的游标，从虚拟头节点开始
        ListNode* cur = dummyHead; 

        // 只要 cur 的下一个节点不为空，就继续探测
        while (cur->next != nullptr) {
            // 发现目标，准备执行“越级指向”的删除动作
            if (cur->next->val == val) {
                // ⚠️ 致命细节：C++ 必须手动回收堆内存！
                ListNode* tmp = cur->next;       // 先用 tmp 保存即将被孤立的那个节点
                cur->next = cur->next->next;     // 越级指向，逻辑上从链表中踢出
                delete tmp;                      // 物理上销毁内存，防止内存泄漏！
            } 
            else {
                // 没有命中目标，游标安全向后移动一步
                cur = cur->next;
            }
        }

        // 3. 收尾工作：提取真正的新头节点，并销毁刚才 new 出来的虚拟头节点
        ListNode* retNode = dummyHead->next; 
        delete dummyHead; // 极其严谨的 C++ 工程师修养

        return retNode;
    }
};
```
---

## 💡 五、 面试直击灵魂的防暴击拷问
* 灵魂拷问一：为什么在 C++ 代码里，踢出链表后必须要加一句 delete tmp;？用 Java 或 Python 需要吗？
绝杀回答： Java 和 Python 都有垃圾回收机制（GC），当一个对象没有任何指针指向它时，JVM 会自动回收它。但在 C++ 中，链表的节点通常都是 new 出来的堆内存。当我们执行 cur->next = cur->next->next; 时，那个等于 val 的节点只是“在逻辑上脱离了链表”，它在物理内存中依然存活！ 如果不手动 delete tmp;，这块内存将永远无法被访问和回收，造成极其严重的内存泄漏（Memory Leak）。

* 灵魂拷问二：最后为什么要先把 dummyHead->next 存到 retNode 里，再 delete dummyHead;？直接 return dummyHead->next; 不行吗？
绝杀回答： 如果你打算把 new 出来的 dummyHead 留在堆内存里不管（这会导致几十个字节的内存泄漏），你可以直接 return dummyHead->next;。但作为一个严谨的 C++ 工程师，既然函数开头 new 了一个临时假节点，函数结束前就必须亲手斩断它。如果不存入 retNode 就先 delete dummyHead，那原本挂在它后面的真实链表入口就成了悬空野指针，再也找不到了！
---