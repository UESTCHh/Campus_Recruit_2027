
#include <iostream>
#include <algorithm>

struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;

    TreeNode(int x)
        : val(x), left(nullptr), right(nullptr) {}
};

// 正确计算最大深度：Day40 已学
int maxDepth(TreeNode* root) {
    if (root == nullptr) {
        return 0;
    }

    int leftH = maxDepth(root->left);
    int rightH = maxDepth(root->right);

    return 1 + std::max(leftH, rightH);
}

// 故意展示错误思路，不能作为 LC111 答案
int naiveMinDepth(TreeNode* root) {
    if (root == nullptr) {
        return 0;
    }

    int leftH = maxDepth(root->left);
    int rightH = maxDepth(root->right);

    return 1 + std::min(leftH, rightH);
}

int main() {
    // 只有左孩子的二叉树
    TreeNode n1(1);
    TreeNode n2(2);
    TreeNode n3(3);

    n1.left = &n2;
    n2.left = &n3;

    std::cout << "Left height: "
              << maxDepth(n1.left) << '\n';

    std::cout << "Right height: "
              << maxDepth(n1.right) << '\n';

    std::cout << "Naive min depth: "
              << naiveMinDepth(&n1) << '\n';

    std::cout << "Expected min depth: 3\n";

    return 0;
}
