// 文件作用：
// 观察Stack的LIFO行为
// ↓
// 学习push
// ↓
// 学习top
// ↓
// 学习pop
// ↓
// 学习empty
// ↓
// 观察完整出栈顺序
#include <iostream>
#include <stack>

using namespace std;

int main()
{
    // --------------------------------------------------
    // 1. 创建一个保存int的Stack
    // --------------------------------------------------
    //
    // Stack遵循：
    //
    // Last In First Out
    // 后进先出
    //
    // numbers一开始没有任何元素。
    stack<int> numbers;


    // --------------------------------------------------
    // 2. 检查空Stack
    // --------------------------------------------------
    //
    // empty()返回bool：
    //
    // true  -> Stack为空
    // false -> Stack非空
    cout << "=== empty stack ==="
         << endl;

    cout << boolalpha;

    cout << "numbers.empty() = "
         << numbers.empty()
         << endl;


    // --------------------------------------------------
    // 3. Push：元素依次入栈
    // --------------------------------------------------
    //
    // 执行：
    //
    // push(10)
    //
    // Stack：
    //
    // top
    //  ↓
    // 10
    numbers.push(10);

    cout << endl;
    cout << "after push 10:"
         << endl;

    cout << "top = "
         << numbers.top()
         << endl;


    // 再push 20：
    //
    // top
    //  ↓
    // 20
    // 10
    numbers.push(20);

    cout << endl;
    cout << "after push 20:"
         << endl;

    cout << "top = "
         << numbers.top()
         << endl;


    // 再push 30：
    //
    // top
    //  ↓
    // 30
    // 20
    // 10
    numbers.push(30);

    cout << endl;
    cout << "after push 30:"
         << endl;

    cout << "top = "
         << numbers.top()
         << endl;


    // --------------------------------------------------
    // 4. size()
    // --------------------------------------------------
    //
    // 当前Stack中：
    //
    // 30
    // 20
    // 10
    //
    // 所以size应该为3。
    cout << "size = "
         << numbers.size()
         << endl;


    // --------------------------------------------------
    // 5. top()只读取，不删除
    // --------------------------------------------------
    //
    // 连续调用两次top：
    //
    // 得到的都应该是30。
    cout << endl;
    cout << "=== top does not remove ==="
         << endl;

    cout << "first top = "
         << numbers.top()
         << endl;

    cout << "second top = "
         << numbers.top()
         << endl;

    cout << "size after top = "
         << numbers.size()
         << endl;


    // --------------------------------------------------
    // 6. pop()删除栈顶
    // --------------------------------------------------
    //
    // 当前：
    //
    // 30 <- top
    // 20
    // 10
    //
    // pop以后：
    //
    // 20 <- top
    // 10
    numbers.pop();

    cout << endl;
    cout << "=== after one pop ==="
         << endl;

    cout << "top = "
         << numbers.top()
         << endl;

    cout << "size = "
         << numbers.size()
         << endl;


    // --------------------------------------------------
    // 7. 完整观察LIFO
    // --------------------------------------------------
    //
    // 当前Stack：
    //
    // 20 <- top
    // 10
    //
    // 我们先继续加入：
    //
    // 40
    // 50
    numbers.push(40);
    numbers.push(50);


    // 此时：
    //
    // top
    //  ↓
    // 50
    // 40
    // 20
    // 10
    //
    // 依次pop应该得到：
    //
    // 50
    // 40
    // 20
    // 10
    //
    // 这就是：
    //
    // Last In First Out
    cout << endl;
    cout << "=== pop all elements ==="
         << endl;

    while (!numbers.empty())
    {
        // pop()本身不返回被删除的值，
        // 所以如果需要看到当前元素：
        //
        // 先top()
        // 再pop()
        cout << "pop value = "
             << numbers.top()
             << endl;

        numbers.pop();
    }


    // --------------------------------------------------
    // 8. 所有元素删除后
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