#include <iostream>

using namespace std;

class Student{
    public:
        int age;

        Student(int age){
            this -> age = age;
        }

        void show(){
            cout << this << endl;
            cout << this -> age << endl;
        }
};

int main(){
    Student s(22);

    cout << &s << endl;

    s.show();

    return 0;
}