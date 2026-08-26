// 使用标准库RAII类型作为成员后，
// 普通类可以依赖编译器生成的拷贝、移动和析构行为
#include <iostream>
#include <string>
#include <vector>

using namespace std;

class Student{
    public:
        string name;
        vector<int> scores;

        Student(string n, vector<int> s) : name(n), scores(s) {
            cout << "constructor" << endl;
        }
};

int main(){
    Student s1(
        "UESTCHh",
        {90, 95, 100}
    );

    Student s2 = s1;

    s2.name = "Copy";
    s2.scores[0] = 60;

    cout << s1.name << endl;
    cout << s1.scores[0] << endl;
    cout << s2.name << endl;
    cout << s2.scores[0] << endl;

    return 0;
}