#include<iostream>
#include<memory>

using namespace std;


class Student
{

public:

    int age;


    Student(int a)
    {
        age=a;

        cout<<"constructor"<<endl;
    }


    ~Student()
    {
        cout<<"destructor"<<endl;
    }

};



int main()
{

    auto p1=make_unique<Student>(22);

    auto p = make_shared<Student>(23);


    auto p2=move(p1);

    cout<< p2->age<<endl;

    cout << p -> age << endl;


}