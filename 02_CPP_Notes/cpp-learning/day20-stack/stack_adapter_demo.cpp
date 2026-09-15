// 文件作用：
// 验证默认Stack

// ↓

// 验证deque作为底层

// ↓

// 验证vector作为底层

// ↓

// 观察：
// 底层Container改变
// 但Stack使用方式保持一致
#include <deque>
#include <iostream>
#include <stack>
#include <vector>

using namespace std;

// printAndClear:
//
// 为了观察不同底层Container构造出来的Stack
// 是否仍然保持相同的LIFO行为。
//
// 注意：
//
// std::stack没有提供普通迭代器让我们遍历全部元素。
//
// 因此这里采用：
//
// top()
// ↓
// pop()
//
// 不断查看并删除栈顶。
//
// 这个函数会把传入的Stack真正清空，
// 所以参数采用引用。
template <typename StackType>
void printAndClear(
    StackType& numbers)
{
    while (!numbers.empty())
    {
        cout << numbers.top()
             << ' ';

        numbers.pop();
    }

    cout << endl;
}

int main()
{
    // ==================================================
    // 1. 默认std::stack
    // ==================================================
    //
    // stack<int>
    //
    // 默认底层：
    //
    // deque<int>
    stack<int> defaultStack;

    defaultStack.push(10);
    defaultStack.push(20);
    defaultStack.push(30);

    cout << "=== default stack ==="
         << endl;

    cout << "top = "
         << defaultStack.top()
         << endl;

    cout << "pop order = ";

    printAndClear(defaultStack);


    // ==================================================
    // 2. 显式使用deque作为底层Container
    // ==================================================
    //
    // 下面和默认stack<int>
    // 在底层Container选择上等价。
    stack<int, deque<int>> dequeStack;

    dequeStack.push(10);
    dequeStack.push(20);
    dequeStack.push(30);

    cout << endl;

    cout << "=== deque based stack ==="
         << endl;

    cout << "top = "
         << dequeStack.top()
         << endl;

    cout << "pop order = ";

    printAndClear(dequeStack);


    // ==================================================
    // 3. 使用vector作为底层Container
    // ==================================================
    //
    // vector本身提供：
    //
    // back()
    // push_back()
    // pop_back()
    //
    // 因此也可以被std::stack适配。
    stack<int, vector<int>> vectorStack;

    vectorStack.push(10);
    vectorStack.push(20);
    vectorStack.push(30);

    cout << endl;

    cout << "=== vector based stack ==="
         << endl;

    cout << "top = "
         << vectorStack.top()
         << endl;

    cout << "pop order = ";

    printAndClear(vectorStack);


    // ==================================================
    // 4. 验证Stack API不随底层Container改变
    // ==================================================
    //
    // 无论底层是：
    //
    // deque
    // vector
    //
    // 外层使用者仍然只使用：
    //
    // push
    // pop
    // top
    // empty
    //
    // 这就是Container Adapter带来的统一抽象。
    cout << endl;

    cout << "=== adapter conclusion ==="
         << endl;

    cout << "different underlying containers, "
         << "same LIFO interface"
         << endl;

    return 0;
}