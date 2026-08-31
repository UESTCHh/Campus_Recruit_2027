// 文件作用：
//     学习：

//     lower_bound
//     upper_bound
//     有序范围查询
#include <iostream>
#include <map>
#include <string>

using namespace std;

int main()
{
    map<int, string> students = {
        {1001, "Alice"},
        {2001, "Bob"},
        {3001, "Rose"},
        {4001, "David"},
        {5001, "Tom"}
    };

    cout << "=== all students ===" << endl;

    for (const auto& student : students)
    {
        cout << student.first
             << " -> "
             << student.second
             << endl;
    }

    cout << endl;

    auto lower = students.lower_bound(2500);

    if (lower != students.end())
    {
        cout << "lower_bound(3001): "
             << lower->first
             << " -> "
             << lower->second
             << endl;
    }

    auto upper = students.upper_bound(5001);

    if (upper != students.end())
    {
        cout << "upper_bound(3001): "
             << upper->first
             << " -> "
             << upper->second
             << endl;
    }

    return 0;
}