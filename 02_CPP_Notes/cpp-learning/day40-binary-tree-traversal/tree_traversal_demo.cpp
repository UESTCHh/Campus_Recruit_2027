#include <iostream>

using namespace std;

struct TreeNode{
    int val;
    TreeNode* left;
    TreeNode* right;

    TreeNode(int x)
        : val(x), left(nullptr), right(nullptr)
    {}
};

void preorder(TreeNode* root){
    if(root == nullptr){
        return;
    }
    cout << root -> val << " ";

    preorder(root -> left);

    preorder(root -> right);
}

void inorder(TreeNode* root){
    if(root == nullptr){
        return;
    }
    inorder(root -> left);

    cout << root -> val << " ";

    inorder(root -> right);
}

void postorder(TreeNode* root){
    if(root == nullptr){
        return;
    }

    postorder(root -> left);

    postorder(root -> right);

    cout << root -> val << " ";
}

int main(){
    /*
          1
         / \
        2   3
       / \
      4   5
    */
   TreeNode* root = new TreeNode(1);
   root -> left = new TreeNode(2);
   root -> right = new TreeNode(3);

   root -> left -> left = new TreeNode(4);
   root -> left -> right = new TreeNode(5);

   cout << "Preorder: ";
   preorder(root);
   cout << endl;

   cout << "Inorder: ";
   inorder(root);
   cout << endl;

   cout << "postorder: ";
   postorder(root);
   cout << endl;

   return 0;
}