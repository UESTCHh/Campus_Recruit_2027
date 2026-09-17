#include <stack>
#include <vector>

using namespace std;

class Solution {
public:
    vector<int> dailyTemperatures(vector<int>& temperatures) {
        int n = static_cast<int>(temperatures.size());
        vector<int> answer(n, 0);
        stack<int> s;

        for (int i = 0; i < n; i++) {
            while (!s.empty() &&
                   temperatures[i] > temperatures[s.top()]) {
                answer[s.top()] = i - s.top();
                s.pop();
            }

            s.push(i);
        }

        return answer;
    }
};