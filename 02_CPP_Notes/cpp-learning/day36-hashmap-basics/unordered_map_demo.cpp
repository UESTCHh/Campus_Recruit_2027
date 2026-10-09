// 文件职责：验证 HashMap 的频次统计、键值映射与查找操作。

#include <iostream>
#include <string>
#include <unordered_map>
#include <vector>

using namespace std;

int main() {
    // ================================
    // 1. 频次统计：元素值 -> 出现次数
    // ================================

    vector<int> nums{4, 7, 1, 7, 3};

    unordered_map<int, int> freq;

    for (int num : nums) {
        // 如果num不存在，operator[]会先创建默认值0
        // 然后将对应次数加1
        freq[num]++;
    }

    cout << "7 count = " << freq[7] << '\n';
    cout << "4 count = " << freq[4] << '\n';

    // ================================
    // 2. 键值映射：学生ID -> 姓名
    // ================================

    unordered_map<int, string> students;

    students[1001] = "Alice";
    students[1002] = "Bob";

    // find()用于查找，不会自动插入新元素
    auto it = students.find(1002);

    if (it != students.end()) {
        cout << "Student 1002 = "
             << it->second
             << '\n';
    }

    // ================================
    // 3. 查找不存在的Key
    // ================================

    auto notFound = students.find(9999);

    if (notFound == students.end()) {
        cout << "Student 9999 not found\n";
    }

    return 0;
}
