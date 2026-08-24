#include <iostream>

using namespace std;

class Student{
    public:
        int* age;

        Student(int a){
            age = new int(a);
        }

        ~Student(){
            delete age;
        }
};

int main(){

    Student s1(22);

    Student s2 = s1;
    
    cout<<s1.age<<endl;
    cout<<s2.age<<endl;
    cout << *s1.age << endl;
    cout << *s2.age << endl;

    return 0;
}