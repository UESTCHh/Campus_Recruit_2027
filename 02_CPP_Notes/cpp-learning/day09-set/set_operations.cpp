// 作用：
//     练习：
//     insert
//     重复插入
//     erase
//     lower_bound
//     upper_bound
//     iterator invalidation
#include <iostream>
#include <set>

using namespace std;

int main()
{
    set<int> numbers = {
        10,
        20,
        30,
        40,
        50
    };

    auto it30 = numbers.find(30);

    cout << "=== before insert ===" << endl;

    if (it30 != numbers.end())
    {
        cout << "it30: " << *it30 << endl;
    }

    cout << endl;

    auto result1 = numbers.insert(25);

    cout << "insert 25 success: "
         << boolalpha
         << result1.second
         << endl;

    cout << "inserted/existing value: "
         << *result1.first
         << endl;

    cout << endl;

    auto result2 = numbers.insert(30);

    cout << "insert 30 again success: "
         << result2.second
         << endl;

    cout << "result2 iterator points to: "
         << *result2.first
         << endl;

    cout << endl;

    cout << "old it30 after insert: "
         << *it30
         << endl;

    cout << endl;

    auto lower = numbers.lower_bound(26);

    if (lower != numbers.end())
    {
        cout << "lower_bound(26): "
             << *lower
             << endl;
    }

    auto upper = numbers.upper_bound(30);

    if (upper != numbers.end())
    {
        cout << "upper_bound(30): "
             << *upper
             << endl;
    }

    cout << endl;

    numbers.erase(10);

    cout << "=== after erase 10 ===" << endl;

    for (int value : numbers)
    {
        cout << value << ' ';
    }

    cout << endl;

    cout << "old it30 after erase 10: "
         << *it30
         << endl;

    return 0;
}