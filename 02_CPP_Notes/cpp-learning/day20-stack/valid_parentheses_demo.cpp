// 文件作用：
// 第一次真正使用Stack解决问题

// 验证：
// Stack保存“最近一个未完成对象”

// 学习：
// LIFO为什么适合括号匹配
#include <iostream>
#include <stack>
#include <string>

using namespace std;

// isMatchingPair：
//
// 判断一个左括号 opening
// 和一个右括号 closing
// 是否属于正确的一对。
//
// 例如：
//
// '(' 和 ')'
// '[' 和 ']'
// '{' 和 '}'
//
// 返回true表示匹配。
bool isMatchingPair(
    char opening,
    char closing)
{
    if (opening == '(' &&
        closing == ')')
    {
        return true;
    }

    if (opening == '[' &&
        closing == ']')
    {
        return true;
    }

    if (opening == '{' &&
        closing == '}')
    {
        return true;
    }

    return false;
}

// isValidParentheses：
//
// 判断只包含：
//
// ()
// []
// {}
//
// 的字符串是否是合法括号序列。
//
// --------------------------------------------------
// Stack中保存什么？
//
// 不是保存所有字符。
//
// 而是保存：
//
// “当前还没有匹配完成的左括号”
//
// Stack top始终代表：
//
// “最近出现、但还没有匹配的左括号”
//
// 这正好利用：
//
// Last In First Out
//
bool isValidParentheses(
    const string& text)
{
    stack<char> brackets;

    // 从左向右扫描字符串。
    for (char ch : text)
    {
        // ------------------------------------------
        // 1. 左括号
        // ------------------------------------------
        //
        // 当前左括号还没有找到右括号，
        // 所以先保存到Stack。
        if (ch == '(' ||
            ch == '[' ||
            ch == '{')
        {
            brackets.push(ch);

            continue;
        }


        // ------------------------------------------
        // 2. 当前字符是右括号
        // ------------------------------------------
        //
        // 它应该匹配：
        //
        // 最近一个尚未匹配的左括号。
        //
        // 也就是Stack top。


        // 如果Stack已经为空，
        // 说明不存在任何左括号可以和当前右括号匹配。
        if (brackets.empty())
        {
            return false;
        }


        // 获取最近一个未匹配左括号。
        char opening =
            brackets.top();


        // 检查括号类型是否匹配。
        //
        // 例如：
        //
        // top == '['
        // ch  == ')'
        //
        // 就是非法。
        if (!isMatchingPair(
                opening,
                ch))
        {
            return false;
        }


        // 当前左右括号已经成功匹配。
        //
        // 所以最近这个左括号已经完成任务，
        // 从Stack中删除。
        brackets.pop();
    }


    // ----------------------------------------------
    // 3. 字符串已经遍历结束
    // ----------------------------------------------
    //
    // 如果Stack为空：
    //
    // 所有左括号都成功找到右括号。
    //
    // 如果Stack非空：
    //
    // 还有左括号没有完成匹配。
    //
    // 例如：
    //
    // "(("
    return brackets.empty();
}

int main()
{
    // --------------------------------------------------
    // Test 1
    // --------------------------------------------------
    //
    // 合法：
    //
    // ()
    string test1 =
        "()";


    // --------------------------------------------------
    // Test 2
    // --------------------------------------------------
    //
    // 合法：
    //
    // (
    //   [
    //   ]
    // )
    // {
    // }
    string test2 =
        "([]){}";


    // --------------------------------------------------
    // Test 3
    // --------------------------------------------------
    //
    // 非法：
    //
    // (
    //   [
    // )
    //   ]
    //
    // 当遇到')'时：
    //
    // Stack top是'['
    //
    // 类型不匹配。
    string test3 =
        "([)]";


    // --------------------------------------------------
    // Test 4
    // --------------------------------------------------
    //
    // 非法：
    //
    // 一开始就是')'
    //
    // 此时Stack为空，
    // 没有左括号可以匹配。
    string test4 =
        ")(";


    // --------------------------------------------------
    // Test 5
    // --------------------------------------------------
    //
    // 非法：
    //
    // 字符串结束以后Stack仍然有：
    //
    // (
    // (
    string test5 =
        "((";


    cout << boolalpha;

    cout << "=== valid parentheses ==="
         << endl;

    cout << test1
         << " -> "
         << isValidParentheses(test1)
         << endl;

    cout << test2
         << " -> "
         << isValidParentheses(test2)
         << endl;

    cout << test3
         << " -> "
         << isValidParentheses(test3)
         << endl;

    cout << test4
         << " -> "
         << isValidParentheses(test4)
         << endl;

    cout << test5
         << " -> "
         << isValidParentheses(test5)
         << endl;


    return 0;
}