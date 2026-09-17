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

    // 默认全部为-1。
    //
    // 如果某个下标最终仍然没有找到
    // 右边第一个更大的元素，
    // 就保持-1。
    vector<int> answer(
        n,
        -1
    );

    // Monotonic Stack保存“下标”。
    //
    // 栈中的下标对应：
    //
    // 已经扫描到，
    // 但还没有找到Next Greater Element的元素。
    stack<int> st;

    for (int i = 0; i < n; ++i)
    {
        // 当前nums[i]是一个新到来的元素。
        //
        // 只要它严格大于栈顶尚未解决元素，
        // 就说明：
        //
        // nums[i]
        //
        // 是这个栈顶元素右边第一个更大的元素。
        while (
            !st.empty() &&
            nums[i] > nums[st.top()]
        )
        {
            int unresolvedIndex =
                st.top();

            answer[unresolvedIndex] =
                nums[i];

            // 这个元素的问题已经解决，
            // 不再需要留在Stack中等待。
            st.pop();
        }

        // 当前元素自己的Next Greater Element
        // 目前还不知道。
        //
        // 所以把当前下标加入Stack等待未来元素。
        st.push(i);
    }

    cout << "nums:   ";

    for (int value : nums)
    {
        cout << value << ' ';
    }

    cout << endl;

    cout << "answer: ";

    for (int value : answer)
    {
        cout << value << ' ';
    }

    cout << endl;

    return 0;
}