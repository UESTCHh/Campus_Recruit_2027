// 文件职责：
// 使用Queue模拟任务到达
// ↓
// Queue保存尚未处理任务
// ↓
// front获得最早待处理任务
// ↓
// 处理完成后pop
// ↓
// 理解FIFO为什么适合任务排队
#include <iostream>
#include <queue>
#include <string>

using namespace std;

// Task：
//
// 表示一个简单任务。
//
// id：
// 用于标识任务。
//
// name：
// 描述任务内容。
struct Task
{
    int id;
    string name;
};

int main()
{
    // --------------------------------------------------
    // 1. 创建任务Queue
    // --------------------------------------------------
    //
    // Queue中保存：
    //
    // 已经到达
    // 但尚未处理
    //
    // 的Task。
    queue<Task> tasks;


    // --------------------------------------------------
    // 2. 模拟任务依次到达
    // --------------------------------------------------
    //
    // 到达顺序：
    //
    // 101
    // 102
    // 103
    //
    // Queue遵循FIFO，
    // 所以以后也应该按照这个顺序处理。
    tasks.push(
        Task{101, "compile"}
    );

    tasks.push(
        Task{102, "test"}
    );

    tasks.push(
        Task{103, "deploy"}
    );


    cout << "=== tasks arrived ==="
         << endl;

    cout << "queue size = "
         << tasks.size()
         << endl;

    // front：
    //
    // 最早进入但尚未处理的Task。
    cout << "next task id = "
         << tasks.front().id
         << endl;

    // back：
    //
    // 最近进入Queue的Task。
    cout << "latest task id = "
         << tasks.back().id
         << endl;


    // --------------------------------------------------
    // 3. 再有一个新任务到达
    // --------------------------------------------------
    tasks.push(
        Task{104, "monitor"}
    );

    cout << endl;

    cout << "=== task 104 arrived ==="
         << endl;

    cout << "next task id = "
         << tasks.front().id
         << endl;

    cout << "latest task id = "
         << tasks.back().id
         << endl;


    // --------------------------------------------------
    // 4. 按FIFO顺序处理所有任务
    // --------------------------------------------------
    //
    // 每轮：
    //
    // front()
    // → 当前最早等待的任务
    //
    // 处理任务
    //
    // pop()
    // → 当前任务处理完成，
    //   从“尚未处理任务集合”中删除
    cout << endl;

    cout << "=== process tasks ==="
         << endl;

    while (!tasks.empty())
    {
        // 先读取当前待处理任务。
        Task currentTask =
            tasks.front();

        cout << "processing task "
             << currentTask.id
             << ": "
             << currentTask.name
             << endl;

        // 当前任务处理完成。
        //
        // 从Queue中移除。
        tasks.pop();
    }


    // --------------------------------------------------
    // 5. 最终状态
    // --------------------------------------------------
    cout << endl;

    cout << "=== final state ==="
         << endl;

    cout << boolalpha;

    cout << "tasks.empty() = "
         << tasks.empty()
         << endl;

    cout << "tasks.size() = "
         << tasks.size()
         << endl;


    return 0;
}