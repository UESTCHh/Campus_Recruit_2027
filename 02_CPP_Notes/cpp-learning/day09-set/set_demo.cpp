#include <iostream>
#include <set>

using namespace std;

int main()
{
    set<int> numbers;

    numbers.insert(30);
    numbers.insert(10);
    numbers.insert(20);
    numbers.insert(20);

    cout << "=== numbers ===" << endl;

    for (int x : numbers)
    {
        cout << x << endl;
    }

    auto it = numbers.find(20);

    if (it != numbers.end())
    {
        cout << "found: " << *it << endl;
    }

    cout << "count(20): "
         << numbers.count(20)
         << endl;

    cout << "count(99): "
         << numbers.count(99)
         << endl;

    return 0;
}