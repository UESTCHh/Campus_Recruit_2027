
#include <iostream>

using namespace std;

// 定义二叉树节点
struct TreeNode {
    int val;            // 当前节点存储的值
    TreeNode* left;     // 指向左孩子
    TreeNode* right;    // 指向右孩子

    // 创建节点时，默认没有左右孩子
    TreeNode(int x)
        : val(x), left(nullptr), right(nullptr) {}
};

int main() {
    // 第一步：创建5个节点
    TreeNode n1(1);
    TreeNode n2(2);
    TreeNode n3(3);
    TreeNode n4(4);
    TreeNode n5(5);

    // 第二步：建立节点之间的关系
    n1.left = &n2;
    n1.right = &n3;

    n2.left = &n4;
    n2.right = &n5;

    // 第三步：用root指向整棵树的根节点
    TreeNode* root = &n1;

    // 访问根节点
    cout << "Root: " << root->val << '\n';

    // 访问根节点的左右孩子
    cout << "Root left: "
         << root->left->val << '\n';

    cout << "Root right: "
         << root->right->val << '\n';

    // 访问节点2的左右孩子
    cout << "Node 2 left: "
         << root->left->left->val << '\n';

    cout << "Node 2 right: "
         << root->left->right->val << '\n';

    // 节点3没有左孩子
    cout << boolalpha;

    cout << "Node 3 has left child: "
         << (root->right->left != nullptr)
         << '\n';

    return 0;
}
