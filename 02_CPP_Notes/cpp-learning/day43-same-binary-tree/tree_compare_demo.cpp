
#include <iostream>

using namespace std;

struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;

    TreeNode(int x)
        : val(x), left(nullptr), right(nullptr) {}
};

// 只比较当前两个节点
// 注意：这不是完整的二叉树比较算法
bool sameCurrentNode(TreeNode* p, TreeNode* q) {
    // 如果至少一方为空：
    // 只有两者都为空时才相同
    if (p == nullptr || q == nullptr) {
        return p == q;
    }

    // 两者都非空，只比较当前节点值
    return p->val == q->val;
}

int main() {
    /*
        树 A：       树 B：

          1            1
         /            /
        2            3
    */

    TreeNode a1(1);
    TreeNode a2(2);

    TreeNode b1(1);
    TreeNode b2(3);

    a1.left = &a2;
    b1.left = &b2;

    TreeNode* p = &a1;
    TreeNode* q = &b1;

    cout << boolalpha;

    // 两个根节点是否是同一个对象？
    cout << "Same address: "
         << (p == q) << '\n';

    // 两个根节点的值是否相同？
    cout << "Same root value: "
         << sameCurrentNode(p, q) << '\n';

    // 左孩子的值是否相同？
    cout << "Same left child: "
         << sameCurrentNode(p->left, q->left) << '\n';

    // 两棵树的右孩子都为空
    cout << "Same right child: "
         << sameCurrentNode(p->right, q->right) << '\n';

    return 0;
}
