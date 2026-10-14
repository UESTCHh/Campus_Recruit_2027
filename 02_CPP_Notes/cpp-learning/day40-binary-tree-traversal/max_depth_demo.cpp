#include <iostream>
#include <algorithm>

using namespace std;

struct TreeNode{
    int val;
    TreeNode* left;
    TreeNode* right;

    TreeNode(int x)
        : val(x), left(nullptr), right(nullptr)
    {}
};

int maxDepth(TreeNode* root){
    if(root == nullptr){
        return 0;
    }

    return 1 + max(maxDepth(root -> left), maxDepth(root -> right));
}

int main(){
    /*
            3
           / \
          9   20
              / \
             15  7
    */

    TreeNode* root = new TreeNode(3);
    root -> left = new TreeNode(9);
    root -> right = new TreeNode(20);

    root -> right -> left = new TreeNode(15);
    root -> right -> right = new TreeNode(7);

    cout << "max depth: " << maxDepth(root) << endl;

    return 0;
}