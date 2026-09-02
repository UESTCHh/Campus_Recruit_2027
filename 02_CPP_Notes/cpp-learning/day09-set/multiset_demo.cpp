#include <iostream>
#include <set>

using namespace std;

int main()
{
    multiset<int> numbers;

    numbers.insert(30);
    numbers.insert(20);
    numbers.insert(20);
    numbers.insert(10);
    numbers.insert(20);

    cout << "=== multiset ===" << endl;

    for (int value : numbers)
    {
        cout << value << ' ';
    }

    cout << endl;

    cout << "count(20): "
         << numbers.count(20)
         << endl;

    auto lower = numbers.lower_bound(20);
    auto upper = numbers.upper_bound(20);

    cout << "values equal to 20: ";

    for (auto it = lower; it != upper; ++it)
    {
        cout << *it << ' ';
    }

    cout << endl;

    auto erased = numbers.erase(20);

    cout << "erase(20) count: "
         << erased
         << endl;

    cout << "=== after erase ===" << endl;

    for (int value : numbers)
    {
        cout << value << ' ';
    }

    cout << endl;

    return 0;
}