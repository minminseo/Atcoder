#include <iostream>
#include <queue>
using namespace std;

int main()
{
    int Q;
    cin >> Q;

    queue<int> que;

    for (int i = 0; i < Q; ++i)
    {
        int type;
        cin >> type;

        if (type == 1)
        {
            int X;
            cin >> X;
            que.push(X);
        }
        else if (type == 2)
        {
            cout << que.front() << endl;
            que.pop();
        }
    }
}
