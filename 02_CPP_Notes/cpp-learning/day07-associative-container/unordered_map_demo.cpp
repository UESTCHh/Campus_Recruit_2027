//文件作用：
//学习：
//	key - value
//	插入
//	查询
//	修改
//	find
//	operator[]
//	遍历

#include <iostream>
#include <string>
#include <unordered_map>

using namespace std;

int main() {
	unordered_map<int, string> students;
	students[1001] = "Alice";
	students[1002] = "Bob";
	students[1003] = "Tom";

	cout << "students 1002: " << students[1002] << endl;

	students[1002] = "Bobby";

	cout << "after update: " << students[1002] << endl;

	auto it = students.find(1003);

	if (it != students.end()) {
		cout << "found key: " << it->first << endl;
		cout << "found value: " << it->second << endl;
	}

	auto missing = students.find(9999);

	if (missing == students.end()) {
		cout << "9999 not found" << endl;
	}

	cout << endl;

	cout << "all students:" << endl;

	for (const auto& item : students) {
		cout << item.first << "->" << item.second << endl;
	}
	return 0;
}