// 文件职责：
// 验证默认Queue
// ↓
// 验证deque底层
// ↓
// 验证list底层
// ↓
// 观察底层Container改变
// ↓
// FIFO Interface保持一致
#include <deque>
#include <iostream>
#include <list>
#include <queue>

using namespace std;

// printAndClear：
//
// 接收不同底层Container构造出的Queue。
//
// 无论底层是：
//
// deque
// list
//
// 对外都只使用：
//
// front()
// pop()
// empty()
//
// 按照FIFO顺序依次输出并清空Queue。
template <typename QueueType>
void printAndClear(
    QueueType& numbers)
{
    while (!numbers.empty())
    {
        cout << numbers.front()
             << ' ';

        numbers.pop();
    }

    cout << endl;
}

int main()
{
    // ==================================================
    // 1. 默认std::queue
    // ==================================================
    //
    // 默认底层：
    //
    // deque<int>
    queue<int> defaultQueue;

    defaultQueue.push(10);
    defaultQueue.push(20);
    defaultQueue.push(30);

    cout << "=== default queue ==="
         << endl;

    cout << "front = "
         << defaultQueue.front()
         << endl;

    cout << "back = "
         << defaultQueue.back()
         << endl;

    cout << "pop order = ";

    printAndClear(defaultQueue);


    // ==================================================
    // 2. 显式使用deque
    // ==================================================
    //
    // 和默认queue<int>在底层选择上相同。
    queue<int, deque<int>> dequeQueue;

    dequeQueue.push(10);
    dequeQueue.push(20);
    dequeQueue.push(30);

    cout << endl;

    cout << "=== deque based queue ==="
         << endl;

    cout << "front = "
         << dequeQueue.front()
         << endl;

    cout << "back = "
         << dequeQueue.back()
         << endl;

    cout << "pop order = ";

    printAndClear(dequeQueue);


    // ==================================================
    // 3. 使用list作为底层Container
    // ==================================================
    //
    // list支持Queue所需要的：
    //
    // front()
    // back()
    // push_back()
    // pop_front()
    //
    // 因此可以作为std::queue的底层Container。
    queue<int, list<int>> listQueue;

    listQueue.push(10);
    listQueue.push(20);
    listQueue.push(30);

    cout << endl;

    cout << "=== list based queue ==="
         << endl;

    cout << "front = "
         << listQueue.front()
         << endl;

    cout << "back = "
         << listQueue.back()
         << endl;

    cout << "pop order = ";

    printAndClear(listQueue);


    // ==================================================
    // 4. 为什么这里不使用vector
    // ==================================================
    //
    // std::queue需要底层Container提供：
    //
    // front()
    // back()
    // push_back()
    // pop_front()
    //
    // std::vector没有pop_front()。
    //
    // 所以vector不满足std::queue完整的底层接口要求。
    //
    // 注意：
    //
    // 这不代表vector无法“手动模拟FIFO”，
    // 而是它不满足std::queue Adapter
    // 对底层Container的接口要求。


    // ==================================================
    // 5. 最终结论
    // ==================================================
    cout << endl;

    cout << "=== adapter conclusion ==="
         << endl;

    cout << "different underlying containers, "
         << "same FIFO interface"
         << endl;

    return 0;
}