// 文件职责：观察 HashSet 如何记录已出现元素、判断重复以及查询存在性。

#include <iostream>
#include <unordered_set>
#include <vector>

using namespace std;

int main() {
    vector<int> nums{2, 4, 2, 5, 4};

    // HashSet中每个元素最多保存一份
    unordered_set<int> seen;

    for (int num : nums) {
        // find()查找当前元素是否已经存在
        if (seen.find(num) != seen.end()) {
            cout << "Duplicate: " << num << '\n';
        }
        else {
            // 首次遇到该元素，放进HashSet
            seen.insert(num);

            cout << "First seen: " << num << '\n';
        }
    }

    // size()返回不同元素的数量
    cout << "Unique count: " << seen.size() << '\n';

    // count()判断元素是否存在
    // 在unordered_set中，结果是0或1
    cout << boolalpha;

    cout << "Contains 5: "
         << (seen.count(5) > 0)
         << '\n';

    cout << "Contains 9: "
         << (seen.count(9) > 0)
         << '\n';

    return 0;
}
