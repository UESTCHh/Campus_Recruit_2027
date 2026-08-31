// 文件作用：
//     学习：

//     map::insert
//     重复key插入
//     map::erase
//     iterator invalidation
#include <iostream>
#include <map>
#include <string>

using namespace std;

int main()
{
    map<int, string> students = {
        {1001, "Alice"},
        {2001, "Bob"},
        {3001, "Rose"}
    };

    auto it2001 = students.find(2001);
    string* name3001 = &students.find(3001)->second;

    cout << "=== before insert ===" << endl;

    cout << "iterator: "
         << it2001->first
         << " -> "
         << it2001->second
         << endl;

    cout << "pointer: "
         << *name3001
         << endl;

    cout << endl;

    auto result1 = students.insert(
        {2500, "Jerry"}
    );

    cout << "insert 2500 success: "
         << boolalpha
         << result1.second
         << endl;

    auto result2 = students.insert(
        {2001, "New Bob"}
    );

    cout << "insert duplicate 2001 success: "
         << result2.second
         << endl;

    cout << "current 2001: "
         << students.find(2001)->second
         << endl;

    cout << endl;

    cout << "=== after insert ===" << endl;

    cout << "old iterator: "
         << it2001->first
         << " -> "
         << it2001->second
         << endl;

    cout << "old pointer: "
         << *name3001
         << endl;

    cout << endl;

    students.erase(1001);

    cout << "=== after erase 1001 ===" << endl;

    cout << "old iterator: "
         << it2001->first
         << " -> "
         << it2001->second
         << endl;

    cout << "old pointer: "
         << *name3001
         << endl;

    cout << endl;

    for (const auto& student : students)
    {
        cout << student.first
             << " -> "
             << student.second
             << endl;
    }

    return 0;
}