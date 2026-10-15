#include <iostream>
#include <queue>

using namespace std;

//二叉树节点
struct TreeNode{
    int val;
    TreeNode* left;
    TreeNode* right;

    TreeNode(int x)
        : val(x), left(nullptr), right(nullptr) {}
};

//BFS:使用队列实现层序遍历
void bfs(TreeNode* root){
    //空树直接返回
    if(root == nullptr){
        return;
    }

    //队列保存节点指针
    queue<TreeNode*> q;

    //根节点首先进入队列
    q.push(root);

    //队列不为空，就继续遍历
    while(!q.empty()){
        //获取队头节点
        TreeNode* current = q.front();

        //移除队头
        q.pop();

        //处理当前节点
        cout << current -> val << " ";

        //如果左右孩子存在，则加入队列
        if(current -> left != nullptr){
            q.push(current -> left);
        }
        if(current -> right != nullptr){
            q.push(current -> right);
        }
    }
}

int main(){
    /*
            1
           / \
          2   3
         / \
        4   5
    */
   TreeNode n1(1);
   TreeNode n2(2);
   TreeNode n3(3);
   TreeNode n4(4);
   TreeNode n5(5);

   //建立节点连接关系
   n1.left = &n2;
   n1.right = &n3;
   n2.left = &n4;
   n2.right = &n5;

   TreeNode* root = &n1;

   cout << "BFS: ";
   bfs(root);
   cout << "\n";

   return 0;
}
