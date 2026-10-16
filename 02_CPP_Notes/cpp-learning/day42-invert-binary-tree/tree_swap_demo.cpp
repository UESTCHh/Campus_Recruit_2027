
#include <iostream>
#include <utility>

using namespace std;

struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;

    TreeNode(int x)
        : val(x), left(nullptr), right(nullptr) {}
};

int main() {
    /*
             1
            / \
           2   3
          / \ / \
         4  5 6  7
    */

    TreeNode n1(1);
    TreeNode n2(2);
    TreeNode n3(3);
    TreeNode n4(4);
    TreeNode n5(5);
    TreeNode n6(6);
    TreeNode n7(7);

    n1.left = &n2;
    n1.right = &n3;

    n2.left = &n4;
    n2.right = &n5;

    n3.left = &n6;
    n3.right = &n7;

    TreeNode* root = &n1;

    // 交换前
    cout << "Before swap:\n";

    cout << "Root left: "
         << root->left->val << '\n';

    cout << "Root right: "
         << root->right->val << '\n';

    // 只交换根节点的左右孩子指针
    swap(root->left, root->right);

    // 交换后
    cout << "\nAfter swap:\n";

    cout << "Root left: "
         << root->left->val << '\n';

    cout << "Root right: "
         << root->right->val << '\n';

    // 验证子树跟随节点一起交换了位置
    cout << "Root left-left: "
         << root->left->left->val << '\n';

    cout << "Root right-left: "
         << root->right->left->val << '\n';

    return 0;
}
