// 文件职责：
// 把Stage 1的canFinish
// 和
// Boundary Binary Search

// 组合成：
// Binary Search on Answer
#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

// 判断给定速度speed是否能够在h小时内完成。
bool canFinish(
    const vector<int>& piles,
    int h,
    int speed
) {
    long long totalHours = 0;

    for (int pile : piles) {
        int hours =
            (pile + speed - 1)
            / speed;

        totalHours += hours;

        // 如果已经超过h，
        // 后面没必要继续累计。
        if (totalHours > h) {
            return false;
        }
    }

    return true;
}

// 对“吃香蕉速度”这个答案空间进行二分。
// 寻找第一个 canFinish(speed) == true 的速度。
int minEatingSpeed(
    const vector<int>& piles,
    int h
) {
    int maxPile =
        *max_element(
            piles.begin(),
            piles.end()
        );

    // 真正答案一定在 [1, maxPile]。
    //
    // left：当前最小的未排除候选。
    // right：当前已知可行的速度上界。
    //
    // 当 left == right 时，
    // 就得到了最小可行速度。
    int left = 1;
    int right = maxPile;
    while (left < right) {
        int mid =
            left + (right - left) / 2;

        if (!canFinish(
                piles,
                h,
                mid
            )) {
            // mid不够快。
            //
            // 根据单调性，
            // <= mid 的速度都不可能成为答案。
            left = mid + 1;
        }
        else {
            // mid已经可以完成，
            // 但题目要求最小速度。
            //
            // mid仍然可能就是答案，
            // 因此不能排除。
            right = mid;
        }
    }

    return left;
}

int main()
{
    vector<int> piles{
        3, 6, 7, 11
    };

    int h = 8;

    cout
        << "minimum eating speed = "
        << minEatingSpeed(
               piles,
               h
           )
        << '\n';

    return 0;
}
