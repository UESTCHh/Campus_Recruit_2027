#include <iostream>
#include <vector>

using namespace std;

int main(){
    vector<int> nums;
    nums.reserve(3);

    nums.push_back(10);
    nums.push_back(20);
    nums.push_back(30);

    cout << "before reallocation" << endl;
    cout << "size:" << nums.size() << endl;
    cout << "capacity:" << nums.capacity() << endl;
    cout << "data:" << static_cast<const void*>(nums.data()) << endl;

    int* p = &nums[0];

    vector<int>::iterator it = nums.begin();

    cout << "p address:" << static_cast<const void*>(p) << endl;
    cout << "iterator value:" << *it << endl;

    nums.push_back(40);
    cout << "after push_back" << endl;
    cout << "size:" << nums.size() << endl;
    cout << "capacity:" << nums.capacity() << endl;
    cout << "data:" << static_cast<const void*>(nums.data()) << endl;
    cout << "new first element address:" << static_cast<const void*>(&nums[0]) << endl;

    return 0;
}