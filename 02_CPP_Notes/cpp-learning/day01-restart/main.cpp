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

    p = nullptr;// 避免野指针访问

    vector<int> nums;

    nums.push_back(1);
    nums.push_back(2);
    nums.push_back(3);

    //引用只读，处理大对象时不会复制元素
    for(const auto& x:nums){
        cout << x << " ";
    }

    return 0;
}