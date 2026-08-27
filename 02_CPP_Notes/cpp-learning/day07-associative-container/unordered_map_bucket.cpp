//文件作用
//	观察：
//	bucket_count
//	bucket(key)
//	bucket_size
//	load_factor
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
    students[2001] = "Jack";
    students[3001] = "Rose";

    //当前哈希表拥有多少个元素
    cout << "size: "
        << students.size()
        << endl;

    //当前哈希表拥有多少个bucket
    cout << "bucket count: "
        << students.bucket_count()
        << endl;

    //元素数量 / bucket数量
    cout << "load factor: "
        << students.load_factor()
        << endl;

    cout << endl;

    for (const auto& item : students)
    {
        cout << "key: "
            << item.first
            << ", value: "
            << item.second
            << ", bucket: "
            << students.bucket(item.first) //当前这个 key 会被映射到哪个 bucket。
            << endl;
    }

    cout << endl;
    cout << "bucket details:" << endl;

    for (size_t i = 0; i < students.bucket_count(); ++i)
    {
        cout << "bucket "
            << i
            << " size: "
            << students.bucket_size(i) //第i号bucket当前包含多少个元素
            << endl;
    }

    return 0;
}