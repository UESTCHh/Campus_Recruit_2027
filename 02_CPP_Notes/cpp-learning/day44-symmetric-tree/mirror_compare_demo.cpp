
#include <iostream>

using namespace std;

struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;

    TreeNode(int x)
        : val(x), left(nullptr), right(nullptr) {}
};

// 只比较一对节点是否相同
// 不递归，因此还不能用来判断整棵树是否对称
bool sameCurrentNode(TreeNode* p, TreeNode* q) {
    if (p == nullptr || q == nullptr) {
        return p == q;
    }

    return p->val == q->val;
}

int main() {
    /*
             1
            / \
           2   2
          / \ / \
         3  4 4  3
    */

    TreeNode n1(1);
    TreeNode n2(2);
    TreeNode n3(2);
    TreeNode n4(3);
    TreeNode n5(4);
    TreeNode n6(4);
    TreeNode n7(3);

    n1.left = &n2;
    n1.right = &n3;

    n2.left = &n4;
    n2.right = &n5;

    n3.left = &n6;
    n3.right = &n7;

    TreeNode* root = &n1;

    cout << boolalpha;

    // 两个直接孩子的值是否相同
    cout << "Root children match: "
         << sameCurrentNode(root->left, root->right)
         << '\n';

    // 外侧镜像位置
    cout << "Outer grandchildren match: "
         << sameCurrentNode(
                root->left->left,
                root->right->right)
         << '\n';

    // 内侧镜像位置
    cout << "Inner grandchildren match: "
         << sameCurrentNode(
                root->left->right,
                root->right->left)
         << '\n';

    // 错误方向：左对左
    cout << "Same-side grandchildren match: "
         << sameCurrentNode(
                root->left->left,
                root->right->left)
         << '\n';

    return 0;
}
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    bool isTwoSym(TreeNode* p, TreeNode* q){
        if(p == nullptr || q == nullptr){
            return p == q;
        }
        return (p -> val == q -> val) && 
            isTwoSym(p -> left, q -> right) &&
            isTwoSym(p -> right, q -> left);
    }
    bool isSymmetric(TreeNode* root) {
        TreeNode* p = root -> left;
        TreeNode* q = root -> right;
        return isTwoSym(p, q);
    }
};
