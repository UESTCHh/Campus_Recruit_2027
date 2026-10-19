
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

// 返回以 root 为根的子树高度
// 如果发现该子树不平衡，则返回 -1
int getHeight(TreeNode* root) {
    if (root == nullptr) {
        return 0;
    }

    // 先计算左子树
    int leftHeight = getHeight(root->left);

    if (leftHeight == -1) {
        return -1;
    }

    // 再计算右子树
    int rightHeight = getHeight(root->right);

    if (rightHeight == -1) {
        return -1;
    }

    // 左右子树都平衡，再检查当前节点
    if (std::abs(leftHeight - rightHeight) > 1) {
        return -1;
    }

    // 当前整棵子树平衡，返回实际高度
    return 1 + std::max(leftHeight, rightHeight);
}

// 对外提供的判断函数
bool isBalanced(TreeNode* root) {
    return getHeight(root) != -1;
}

int main() {
    std::cout << std::boolalpha;

    // Tree A：平衡二叉树
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
    std::cout << "Height result: "
              << getHeight(&a1) << '\n';
    std::cout << "Is balanced: "
              << isBalanced(&a1) << '\n';

    // Tree B：根节点高度差为 0
    // 但内部子树不平衡
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
    std::cout << "Height result: "
              << getHeight(&b1) << '\n';
    std::cout << "Is balanced: "
              << isBalanced(&b1) << '\n';

    // 空树
    TreeNode* emptyTree = nullptr;

    std::cout << "\n=== Empty Tree ===\n";
    std::cout << "Height result: "
              << getHeight(emptyTree) << '\n';
    std::cout << "Is balanced: "
              << isBalanced(emptyTree) << '\n';

    // 单节点树
    TreeNode single(100);

    std::cout << "\n=== Single Node ===\n";
    std::cout << "Height result: "
              << getHeight(&single) << '\n';
    std::cout << "Is balanced: "
              << isBalanced(&single) << '\n';

    return 0;
}
