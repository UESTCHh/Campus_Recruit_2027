#include <iostream>
#include <memory>

using namespace std;

class Student{
    public:
        int age;

        Student(int a) : age(a){
                cout << "Student constructor" << endl;
        }

        ~Student(){
            cout << "Student destructor" << endl;
        }
};

int main(){
    auto p1 = make_shared<Student>(22);

    auto p2 = p1;

    weak_ptr<Student> w = p1;

    cout << "count: " << p1.use_count() << endl;

    p1.reset();

    cout << "after, count: " << p2.use_count() << endl;

    cout << "expired:" << w.expired() << endl;

    {
    auto p3 = w.lock();

    if(p3){
        cout << "Student still alive" << endl;
    }
    else{
        cout << "Student already destroyed" << endl;
    }
    }

    cout << "after p3 destroyed, count:" << p2.use_count() << endl;
    
    p2.reset();

    cout << "after p2 reset, expired:" << w.expired() << endl;
    
    auto p4 = w.lock();

    if(p4){
        cout << "Student still alive" << endl;
    }
    else{
        cout << "Student already destroyed" << endl;
    }
    return 0;
}