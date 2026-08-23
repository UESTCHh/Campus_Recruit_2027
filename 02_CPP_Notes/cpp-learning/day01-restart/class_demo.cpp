#include <iostream>

using namespace std;

class Student{

private:
    string name;
        
    int age;

public:
    Student(string n, int a){
        name = n;

        age = a;

        cout << "constructor" << endl;
    }

    void show(){
        cout << name << " " << age << endl;
    }

    ~Student(){
        cout << "destructor" << endl;
    }

};

int main(){
    Student s("hui", 22);

    Student *p=new Student("h",22);

    s.show();

    delete p;
    
    return 0;
}