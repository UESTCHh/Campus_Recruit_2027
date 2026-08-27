//文件作用
//	验证：
//	自定义hash function
//	↓
//	强制制造collision
//	↓
//	不同key全部进入同一bucket
//	↓
//	unordered_map仍能正确查找
#include <iostream>
#include <string>
#include <unordered_map>

using namespace std;

struct BadHash
{
    // 函数调用运算符重载。
    // 它让一个对象可以像函数一样调用。
    size_t operator()(int key) const
    {
        return 0;
    }
};

int main()
{
    unordered_map<int, string, BadHash> students;// 这个 unordered_map 使用什么对象来计算 key 的 hash。

    students[1001] = "Alice";
    students[1002] = "Bobby";
    students[1003] = "Tom";
    students[2001] = "Jack";
    students[3001] = "Rose";

    cout << "size: "
        << students.size()
        << endl;

    cout << "bucket count: "
        << students.bucket_count()
        << endl;

    cout << "load factor: "
        << students.load_factor()
        << endl;

    cout << endl;
    cout << "element buckets:" << endl;

    for (const auto& item : students)
    {
        cout << "key: "
            << item.first
            << ", bucket: "
            << students.bucket(item.first)
            << endl;
    }

    cout << endl;

    size_t targetBucket = students.bucket(1001);

    cout << "target bucket: "
        << targetBucket
        << endl;

    cout << "target bucket size: "
        << students.bucket_size(targetBucket)
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

    return 0;
}