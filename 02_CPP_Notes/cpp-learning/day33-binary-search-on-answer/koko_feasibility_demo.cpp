// 文件职责：
// 暂时不做完整二分

// 只训练：

// 候选速度k
// ↓
// canFinish(k)
// ↓
// 观察 F F F | T T T
#include <iostream>
#include <vector>

using namespace std;

// 判断：
// 如果Koko每小时最多吃speed根香蕉，
// 能不能在h小时内吃完所有香蕉。
bool canFinish(
    const vector<int>& piles,
    int h,
    int speed
) {
    long long totalHours = 0;

    for (int pile : piles) {
        // ceil(pile / speed)
        //
        // 例如：
        // pile = 11
        // speed = 4
        //
        // 需要3小时。
        //
        // 整数情况下可以写成：
        // (pile + speed - 1) / speed
        int hours =
            (pile + speed - 1)
            / speed;

        totalHours += hours;
    }

    return totalHours <= h;
}

int main()
{
    vector<int> piles{
        3, 6, 7, 11
    };

    int h = 8;

    for (int speed = 1;
         speed <= 6;
         ++speed) {

        cout
            << "speed = "
            << speed
            << ", canFinish = "
            << boolalpha
            << canFinish(
                   piles,
                   h,
                   speed
               )
            << '\n';
    }

    return 0;
}