
#include <iostream>
#include <algorithm>
#include <cstdlib>

struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;

    TreeNode(int x)
        : val(x), left(nullptr), right(nullptr) {}
};

// 返回以 root 为根的二叉树高度
int maxDepth(TreeNode* root) {
    if (root == nullptr) {
        return 0;
    }

    int leftHeight = maxDepth(root->left);
    int rightHeight = maxDepth(root->right);

    return 1 + std::max(leftHeight, rightHeight);
}

// 只计算某个节点左右子树的高度差
int heightDifference(TreeNode* root) {
    if (root == nullptr) {
        return 0;
    }

    int leftHeight = maxDepth(root->left);
    int rightHeight = maxDepth(root->right);

    return std::abs(leftHeight - rightHeight);
}

int main() {
    // 示例 A：平衡二叉树
    TreeNode a1(1);
    TreeNode a2(2);
    TreeNode a3(3);
    TreeNode a4(4);
    TreeNode a5(5);

    a1.left = &a2;
    a1.right = &a3;
    a2.left = &a4;
    a2.right = &a5;

    std::cout << "=== Tree A ===\n";
    std::cout << "Tree height: "
              << maxDepth(&a1) << '\n';
    std::cout << "Root height difference: "
              << heightDifference(&a1) << '\n';

    // 示例 B：根节点高度差为 0，但整棵树不平衡
    TreeNode b1(1);
    TreeNode b2(2);
    TreeNode b3(3);
    TreeNode b4(4);
    TreeNode b5(5);
    TreeNode b6(6);
    TreeNode b7(7);

    b1.left = &b2;
    b1.right = &b3;
    b2.left = &b4;
    b4.left = &b5;
    b3.right = &b6;
    b6.right = &b7;

    std::cout << "\n=== Tree B ===\n";
    std::cout << "Tree height: "
              << maxDepth(&b1) << '\n';
    std::cout << "Root height difference: "
              << heightDifference(&b1) << '\n';
    std::cout << "Node 2 height difference: "
              << heightDifference(&b2) << '\n';

    return 0;
}
