// 文件作用：
// 理解FIFO
// ↓
// 观察push
// ↓
// 观察front / back
// ↓
// 观察pop
// ↓
// 使用empty安全清空Queue
#include <iostream>
#include <queue>

using namespace std;

int main()
{
    // --------------------------------------------------
    // 1. 创建一个保存int的Queue
    // --------------------------------------------------
    //
    // Queue遵循：
    //
    // First In First Out
    // FIFO
    // 先进先出
    //
    // numbers一开始为空。
    queue<int> numbers;


    // --------------------------------------------------
    // 2. 检查初始状态
    // --------------------------------------------------
    cout << boolalpha;

    cout << "=== empty queue ==="
         << endl;

    cout << "numbers.empty() = "
         << numbers.empty()
         << endl;

    cout << "size = "
         << numbers.size()
         << endl;


    // --------------------------------------------------
    // 3. push(10)
    // --------------------------------------------------
    //
    // 当前只有一个元素：
    //
    // front / back
    //      ↓
    //     10
    numbers.push(10);

    cout << endl;
    cout << "=== after push 10 ==="
         << endl;

    cout << "front = "
         << numbers.front()
         << endl;

    cout << "back = "
         << numbers.back()
         << endl;


    // --------------------------------------------------
    // 4. push(20)
    // --------------------------------------------------
    //
    // Queue：
    //
    // front       back
    //   ↓           ↓
    //  10   →      20
    numbers.push(20);

    cout << endl;
    cout << "=== after push 20 ==="
         << endl;

    cout << "front = "
         << numbers.front()
         << endl;

    cout << "back = "
         << numbers.back()
         << endl;


    // --------------------------------------------------
    // 5. push(30)
    // --------------------------------------------------
    //
    // Queue：
    //
    // front              back
    //   ↓                  ↓
    //  10   →   20   →    30
    numbers.push(30);

    cout << endl;
    cout << "=== after push 30 ==="
         << endl;

    cout << "front = "
         << numbers.front()
         << endl;

    cout << "back = "
         << numbers.back()
         << endl;

    cout << "size = "
         << numbers.size()
         << endl;


    // --------------------------------------------------
    // 6. front() / back()只读取，不删除
    // --------------------------------------------------
    cout << endl;

    cout << "=== front and back do not remove ==="
         << endl;

    cout << "first front = "
         << numbers.front()
         << endl;

    cout << "second front = "
         << numbers.front()
         << endl;

    cout << "back = "
         << numbers.back()
         << endl;

    cout << "size = "
         << numbers.size()
         << endl;


    // --------------------------------------------------
    // 7. pop()删除front
    // --------------------------------------------------
    //
    // 原：
    //
    // 10 → 20 → 30
    //
    // pop以后：
    //
    // 20 → 30
    //
    // 删除的是最早进入的10。
    numbers.pop();

    cout << endl;
    cout << "=== after one pop ==="
         << endl;

    cout << "front = "
         << numbers.front()
         << endl;

    cout << "back = "
         << numbers.back()
         << endl;

    cout << "size = "
         << numbers.size()
         << endl;


    // --------------------------------------------------
    // 8. 再加入40和50
    // --------------------------------------------------
    //
    // 当前：
    //
    // 20 → 30
    //
    // push 40 / 50后：
    //
    // front                         back
    //   ↓                             ↓
    //  20 → 30 → 40 → 50
    numbers.push(40);
    numbers.push(50);


    // --------------------------------------------------
    // 9. 按FIFO顺序全部出队
    // --------------------------------------------------
    //
    // Queue的pop()本身不返回元素，
    // 所以：
    //
    // 先front()
    // 再pop()
    cout << endl;

    cout << "=== pop all elements ==="
         << endl;

    while (!numbers.empty())
    {
        cout << "pop value = "
             << numbers.front()
             << endl;

        numbers.pop();
    }


    // --------------------------------------------------
    // 10. 最终状态
    // --------------------------------------------------
    cout << endl;

    cout << "=== final state ==="
         << endl;

    cout << "numbers.empty() = "
         << numbers.empty()
         << endl;

    cout << "size = "
         << numbers.size()
         << endl;


    return 0;
}