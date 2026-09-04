#include <deque>
#include <iostream>

using namespace std;

int main()
{
    deque<int> numbers;

    numbers.push_back(20);
    numbers.push_back(30);

    numbers.push_front(10);
    numbers.push_back(40);
    numbers.push_front(0);

    cout << "=== initial deque ===" << endl;

    for (int value : numbers)
    {
        cout << value << ' ';
    }

    cout << endl;

    cout << "front: "
         << numbers.front()
         << endl;

    cout << "back: "
         << numbers.back()
         << endl;

    cout << "numbers[2]: "
         << numbers[2]
         << endl;

    cout << "size: "
         << numbers.size()
         << endl;

    numbers.pop_front();
    numbers.pop_back();

    cout << endl;
    cout << "=== after pop_front and pop_back ==="
         << endl;

    for (int value : numbers)
    {
        cout << value << ' ';
    }

    cout << endl;

    return 0;
}