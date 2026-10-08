
#include <iostream>
#include <queue>

using namespace std;

int main() {
    queue<int> q;

    // 模拟三个任务按照顺序到达
    q.push(10);
    q.push(20);
    q.push(30);

    // front() 只是查看队头，不会删除元素
    cout << "First task: " << q.front() << '\n';

    // 处理第一个任务，再将其出队
    cout << "Processing: " << q.front() << '\n';
    q.pop();

    // 删除10后，队头变成20
    cout << "New front: " << q.front() << '\n';

    // 继续处理剩余任务
    while (!q.empty()) {
        cout << "Processing: " << q.front() << '\n';
        q.pop();
    }

    cout << "Queue empty: " << boolalpha << q.empty() << '\n';

    return 0;
}
