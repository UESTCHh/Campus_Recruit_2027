#include <iostream>
#include <string>

using namespace std;

int main()
{
    string text = "hello";

    cout << "=== basic ===" << endl;

    cout << "text: "
         << text
         << endl;

    cout << "size: "
         << text.size()
         << endl;

    cout << "first char: "
         << text.front()
         << endl;

    cout << "last char: "
         << text.back()
         << endl;

    cout << "text[1]: "
         << text[1]
         << endl;

    text[0] = 'H';

    text.push_back('!');

    cout << "after modify: "
         << text
         << endl;

    text.pop_back();

    cout << "after pop_back: "
         << text
         << endl;

    string sentence = text + " world";

    cout << endl;
    cout << "=== string operations ==="
         << endl;

    cout << "sentence: "
         << sentence
         << endl;

    cout << "substr(0, 5): "
         << sentence.substr(0, 5)
         << endl;

    cout << "substr(6, 5): "
         << sentence.substr(6, 5)
         << endl;

    size_t pos = sentence.find("world");

    if (pos != string::npos)
    {
        cout << "world found at: "
             << pos
             << endl;
    }

    size_t missing = sentence.find("C++");

    if (missing == string::npos)
    {
        cout << "C++ not found"
             << endl;
    }

    return 0;
}