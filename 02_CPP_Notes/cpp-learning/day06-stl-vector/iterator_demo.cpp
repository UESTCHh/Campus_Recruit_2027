#include <iostream>
#include <vector>

using namespace std;

int main(){
    vector<int> nums = {10, 20, 30, 40};

    cout << "iterator:" << endl;

    for(vector<int>::iterator it = nums.begin(); it != nums.end(); ++it){
        cout << *it << endl;
    }
    cout << endl;
    cout << "range-for:" << endl;

    for(int value:nums){
        cout << value << endl;
    }

    return 0;
}