#include <iostream>
#include <mutex>

using namespace std;

mutex m;

void doWork(){
    lock_guard<mutex> guard(m);

    cout << "lock acquired" << endl;

    cout << "return from doWork" << endl;

    return;
}

int main(){
    doWork();

    if(m.try_lock()){
        cout << "lock has been released" << endl;

        m.unlock();
    }
    else{
        cout << "lock is still held" << endl;
    }

    return 0;
}