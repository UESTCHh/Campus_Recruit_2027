#include <iostream>
#include <vector>

using namespace std;

int main(){
    vector<int> nums;

    for(int i = 1; i <= 20; i++){
        nums.push_back(i);

        // data得到vector当前连续存储区域首元素的地址
        cout << "push: " << i
            << ", size: " << nums.size()
            << ", capacity: " << nums.capacity()
            << ", data: " << static_cast<const void*>(nums.data())
            << endl;
    }

    return 0;
}