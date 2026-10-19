
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

// LC104：返回二叉树高度
int maxDepth(TreeNode* root) {
    if (root == nullptr) {
        return 0;
    }

    int leftHeight = maxDepth(root->left);
    int rightHeight = maxDepth(root->right);

    return 1 + std::max(leftHeight, rightHeight);
}

// LC110：基础递归解法
bool isBalanced(TreeNode* root) {
    if (root == nullptr) {
        return true;
    }

    // 1. 计算左右子树高度
    int leftHeight = maxDepth(root->left);
    int rightHeight = maxDepth(root->right);

    // 2. 判断当前节点是否满足平衡条件
    bool currentBalanced =
        std::abs(leftHeight - rightHeight) <= 1;

    // 3. 判断左右子树自身是否平衡
    bool leftBalanced = isBalanced(root->left);
    bool rightBalanced = isBalanced(root->right);

    // 4. 三个条件必须同时成立
    return currentBalanced
        && leftBalanced
        && rightBalanced;
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

    std::cout << "Tree A balanced: "
              << isBalanced(&a1) << '\n';

    // Tree B：根节点高度差为 0，
    // 但左右子树内部存在不平衡节点
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

    std::cout << "Tree B balanced: "
              << isBalanced(&b1) << '\n';

    // 空树测试
    TreeNode* emptyTree = nullptr;

    std::cout << "Empty tree balanced: "
              << isBalanced(emptyTree) << '\n';

    // 单节点树测试
    TreeNode single(100);

    std::cout << "Single node balanced: "
              << isBalanced(&single) << '\n';

    return 0;
}
