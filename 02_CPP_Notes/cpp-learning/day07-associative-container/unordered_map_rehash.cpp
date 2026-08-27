// 文件作用
//     验证：
//     rehash前保存：
//     iterator
//     pointer
//     reference

//     ↓

//     强制rehash

//     ↓

//     旧iterator不能再使用
//     pointer/reference仍然有效
#include <iostream>
#include <string>
#include <unordered_map>

using namespace std;

int main()
{
    unordered_map<int, string> students;

    students[1001] = "Alice";
    students[1002] = "Bobby";
    students[1003] = "Tom";

    auto it = students.find(1002);

    string* p = &it->second;
    string& ref = it->second;

    cout << "before rehash" << endl;

    cout << "bucket count: "
        << students.bucket_count()
        << endl;

    cout << "value: "
        << *p
        << endl;

    cout << "value address: "
        << static_cast<const void*>(p)
        << endl;

    size_t oldBucketCount = students.bucket_count();

    students.rehash(oldBucketCount * 10 + 1);

    cout << endl;
    cout << "after rehash" << endl;

    cout << "bucket count: "
        << students.bucket_count()
        << endl;

    cout << "value through pointer: "
        << *p
        << endl;

    cout << "value through reference: "
        << ref
        << endl;

    cout << "old value address: "
        << static_cast<const void*>(p)
        << endl;

    auto newIt = students.find(1002);

    cout << "new value address: "
        << static_cast<const void*>(&newIt->second)
        << endl;

    return 0;
}