// 作用：
//     观察：
//     stack LIFO
//     queue FIFO
//     top/front/back
//     push/pop
#include <iostream>
#include <queue>
#include <stack>

using namespace std;

int main()
{
    stack<int> s;

    s.push(10);
    s.push(20);
    s.push(30);

    cout << "=== stack ===" << endl;

    cout << "top: "
         << s.top()
         << endl;

    cout << "size: "
         << s.size()
         << endl;

    while (!s.empty())
    {
        cout << s.top() << ' ';
        s.pop();
    }

    cout << endl;

    queue<int> q;

    q.push(10);
    q.push(20);
    q.push(30);

    cout << endl;
    cout << "=== queue ===" << endl;

    cout << "front: "
         << q.front()
         << endl;

    cout << "back: "
         << q.back()
         << endl;

    cout << "size: "
         << q.size()
         << endl;

    while (!q.empty())
    {
        cout << q.front() << ' ';
        q.pop();
    }

    cout << endl;

    return 0;
}