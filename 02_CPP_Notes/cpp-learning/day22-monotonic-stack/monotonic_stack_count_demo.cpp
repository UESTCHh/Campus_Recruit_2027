// 作用：
// 统计每个元素的push/pop
// ↓
// 亲眼确认
// ↓
// 一个元素最多push一次、pop一次
#include <iostream>
#include <stack>
#include <vector>

using namespace std;

int main()
{
    vector<int> nums{
        2, 1, 2, 4, 3
    };

    int n =
        static_cast<int>(
            nums.size()
        );

    vector<int> answer(
        n,
        -1
    );

    stack<int> st;

    int pushCount = 0;
    int popCount = 0;

    for (int i = 0; i < n; ++i)
    {
        cout << "scan nums["
             << i
             << "] = "
             << nums[i]
             << endl;

        while (
            !st.empty() &&
            nums[i] > nums[st.top()]
        )
        {
            int unresolvedIndex =
                st.top();

            cout << "  solve nums["
                 << unresolvedIndex
                 << "] = "
                 << nums[unresolvedIndex]
                 << " with "
                 << nums[i]
                 << endl;

            answer[unresolvedIndex] =
                nums[i];

            st.pop();
            ++popCount;
        }

        st.push(i);
        ++pushCount;

        cout << "  push index "
             << i
             << endl;

        cout << endl;
    }

    cout << "answer: ";

    for (int value : answer)
    {
        cout << value << ' ';
    }

    cout << endl;

    cout << "push count = "
         << pushCount
         << endl;

    cout << "pop count = "
         << popCount
         << endl;

    cout << "total stack operations = "
         << pushCount + popCount
         << endl;

    return 0;
}