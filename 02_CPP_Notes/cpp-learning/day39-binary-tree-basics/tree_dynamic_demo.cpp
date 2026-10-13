
#include <iostream>

using namespace std;

// 二叉树节点定义
struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;

    // 初始化节点值和左右孩子指针
    TreeNode(int x)
        : val(x), left(nullptr), right(nullptr) {}
};

// 动态创建一棵二叉树
// 返回值：根节点的地址
TreeNode* createTree() {
    // 创建根节点1
    TreeNode* root = new TreeNode(1);

    // 创建左右孩子2和3
    root->left = new TreeNode(2);
    root->right = new TreeNode(3);

    // 给节点2创建左孩子4
    root->left->left = new TreeNode(4);

    // 返回根节点指针
    return root;
}

int main() {
    // 调用函数，获取动态创建的树
    TreeNode* root = createTree();

    // 使用根节点指针访问整棵树
    cout << "Root: " << root->val << '\n';
    cout << "Left: " << root->left->val << '\n';
    cout << "Right: " << root->right->val << '\n';
    cout << "Left-left: "
         << root->left->left->val << '\n';

    // 判断节点3是否有左孩子
    cout << boolalpha;
    cout << "Right-left is null: "
         << (root->right->left == nullptr)
         << '\n';

    // 释放动态创建的全部节点
    // 注意：这些语句只适用于当前固定树结构

    // 先释放节点4
    delete root->left->left;

    // 再释放节点2和3
    delete root->left;
    delete root->right;

    // 最后释放根节点1
    delete root;

    // 避免继续保留已经失效的根节点地址
    root = nullptr;

    return 0;
}
