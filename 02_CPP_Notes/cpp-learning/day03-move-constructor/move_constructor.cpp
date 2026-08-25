#include <iostream>

using namespace std;

class Student{
    public:
        int* age;

        Student(int a){
            age = new int(a);

            cout << "constructor" << endl;

        }

        Student(const Student& other){
            age = new int(*other.age);

            cout << "copy constructor" << endl;
        }

        Student(Student&& other){
            age = other.age;

            other.age = nullptr;

            cout << "move constructor" << endl;
        }

        ~Student(){
            delete age;
        }
};

Student createStudent(){
    Student s(22);

    return s;
}

int main(){
    Student s = createStudent();

    cout << *s.age << endl;
}