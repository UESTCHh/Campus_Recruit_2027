// 作用：
// 验证map遍历顺序由Comparator决定
#include <functional>
#include <iostream>
#include <map>
#include <string>

using namespace std;

int main()
{
    map<
        int,
        string,
        greater<int>
    > students;

    students[3001] = "Rose";
    students[1001] = "Alice";
    students[2001] = "Bob";
    students[5001] = "David";
    students[2500] = "Jerry";

    for (const auto& student : students)
    {
        cout << student.first
             << " -> "
             << student.second
             << endl;
    }

    return 0;
}