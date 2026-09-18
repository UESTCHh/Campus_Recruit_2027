// 只观察：
// Heap的数组表示
// +
// Parent / Left Child / Right Child
#include <iostream>
#include <vector>

using namespace std;

int main()
{
    // 这是一个合法的Max Heap。
    //
    // 如果画成树：
    //
    //         10
    //        /  \
    //       8    9
    //      / \   /
    //     1   3 5
    //
    // Heap虽然逻辑上可以看成树，
    // 实际上通常可以直接存放在连续数组中。
    vector<int> heap{
        10, 8, 9, 1, 3, 5
    };

    int n =
        static_cast<int>(
            heap.size()
        );

    for (int i = 0; i < n; ++i)
    {
        cout << "index "
             << i
             << ", value = "
             << heap[i]
             << endl;

        // 左孩子下标：
        // 2 * i + 1
        int left =
            2 * i + 1;

        // 右孩子下标：
        // 2 * i + 2
        int right =
            2 * i + 2;

        if (left < n)
        {
            cout << "  left child: index "
                 << left
                 << ", value = "
                 << heap[left]
                 << endl;
        }

        if (right < n)
        {
            cout << "  right child: index "
                 << right
                 << ", value = "
                 << heap[right]
                 << endl;
        }

        // 根节点index 0没有父节点。
        if (i > 0)
        {
            int parent =
                (i - 1) / 2;

            cout << "  parent: index "
                 << parent
                 << ", value = "
                 << heap[parent]
                 << endl;
        }

        cout << endl;
    }

    return 0;
}