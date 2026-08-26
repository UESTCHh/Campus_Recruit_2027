#include <iostream>
#include <algorithm>
#include <vector>

using namespace std;

int main(){
    vector<int> nums = {40, 10, 30, 20, 30};
    cout << "original:" << endl;
    for(const int& value : nums){
        cout << value << " ";
    }
    cout << endl;
    sort(nums.begin(), nums.end());
    cout << "sorted:" << endl;
    for(const int& value : nums){
        cout << value << " ";
    }
    cout << endl;
    auto it = find(nums.begin(), nums.end(), 30);
    if(it != nums.end()){
        cout << "find:" << *it << endl;
    }
    else{
        cout << "not found" << endl;
    }

    int result = count(nums.begin(), nums.end(), 30);

    cout << "count 30: " << result << endl;

    return 0;
}