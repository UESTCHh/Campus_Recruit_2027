// 职责：
// 理解：
// 普通Queue和Priority Queue的区别

// 学习：
// push
// top
// pop
// empty
// size

// 观察：
// 默认priority_queue为什么总能优先拿到最大元素
#include <iostream>
#include <queue>

using namespace std;

int main()
{
    // std::priority_queue 是一个 Container Adapter。
    //
    // priority_queue<int> 默认：
    //
    // 数值越大，
    // 优先级越高。
    //
    // 因此 top() 始终能够得到当前最大元素。
    priority_queue<int> numbers;

    numbers.push(5);

    cout << "push 5, top = "
         << numbers.top()
         << endl;

    numbers.push(1);

    cout << "push 1, top = "
         << numbers.top()
         << endl;

    numbers.push(8);

    cout << "push 8, top = "
         << numbers.top()
         << endl;

    numbers.push(3);

    cout << "push 3, top = "
         << numbers.top()
         << endl;

    cout << endl;

    cout << "current size = "
         << numbers.size()
         << endl;

    cout << endl;

    // 不断读取最高优先级元素，
    // 然后将它删除。
    //
    // 因为默认是最大元素优先，
    // 所以输出会从大到小。
    cout << "pop order: ";

    while (!numbers.empty())
    {
        cout << numbers.top()
             << ' ';

        numbers.pop();
    }

    cout << endl;

    return 0;
}