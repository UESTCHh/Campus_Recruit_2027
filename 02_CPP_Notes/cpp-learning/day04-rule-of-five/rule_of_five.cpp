#include<iostream>

using namespace std;


class Student
{

public:

    int* age;


    Student(int a)
    {
        age=new int(a);

        cout<<"constructor"<<endl;
    }

    Student(const Student& other)
    {
        age = new int(*other.age);

        cout<<"copy constructor"<<endl;
    }

    Student(Student&& other){
        age = other.age;

        other.age = nullptr;

        cout << "move constructor" << endl;
    }

    Student& operator=(const Student& other)
    {

        if(this == &other)
        {
            return *this;
        }


        delete age;


        age = new int(*other.age);


        cout<<"copy assignment"<<endl;


        return *this;
    }

    Student& operator=(Student&& other){
        if(this == &other){
            return *this;
        }

        delete age;

        age = other.age;

        other.age = nullptr;

        cout << "move assignment" << endl;

        return *this;
    }

    ~Student()
    {
        delete age;

        cout<<"destructor"<<endl;
    }


};

int main()
{

    Student s1(22);


    Student s2 = std::move(s1);


    cout<<*s2.age<<endl;


}