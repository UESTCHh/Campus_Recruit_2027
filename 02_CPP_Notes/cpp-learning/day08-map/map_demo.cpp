// 文件作用：
//     学习std::map最基本的：

//     插入
//     修改
//     查找
//     遍历
//     有序性
#include <iostream>
#include <map>
#include <string>

using namespace std;

int main()
{
    map<int, string> students;

    students[3001] = "Rose";
    students[1001] = "Alice";
    students[2001] = "Bob";

    students[5001] = "David";
    students[101] = "Tom";
    students[2500] = "Jerry";

    cout << "=== all students ===" << endl;

    for (const auto& student : students)
    {
        cout << student.first
             << " -> "
             << student.second
             << endl;
    }

    cout << endl;

    students[2001] = "Bob Updated";

    cout << "student 2001: "
         << students[2001]
         << endl;

    cout << endl;

    auto it = students.find(3001);

    if (it != students.end())
    {
        cout << "found: "
             << it->first
             << " -> "
             << it->second
             << endl;
    }
    else
    {
        cout << "not found" << endl;
    }

    return 0;
}