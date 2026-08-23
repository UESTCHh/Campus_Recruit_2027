#include <iostream>
#include <vector>

using namespace std;

void change(int& x){
    x = 100;
}

int main(){
    int a = 10;

    int *p = new int(20);

    change(a);
    cout << a << endl;

    cout << *p << endl;

    delete p;

    vector<int> nums;

    nums.push_back(1);
    nums.push_back(2);
    nums.push_back(3);

    for(auto x:nums){
        cout << x << " ";
    }
    
    return 0;
}