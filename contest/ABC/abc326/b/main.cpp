#include <iostream>
using namespace std;

int main()
{
    int n;
    cin >> n;
    for (int i = n; i <= 919; i++)
    {
        int a = i / 100;
        int b = i / 10 % 10;
        int c = i % 10;
        if (a * b == c)
        {
            cout << i << endl;
            return 0;
        }
    }
    cout << -1 << endl;
}
