#include <bits/stdc++.h>
using namespace std;

int main()
{
    string T, U;
    cin >> T >> U;

    if (T.size() < U.size())
    {
        cout << "No" << endl;
        return 0;
    }

    for (int i = 0; i < T.size() - U.size() + 1; i++)
    {
        bool ok = true;
        for (int j = 0; j < U.size(); j++)
        {
            if (T[i + j] != '?' && T[i + j] != U[j])
            {
                ok = false;
                break;
            }
        }
        if (ok)
        {
            cout << "Yes" << endl;
            return 0;
        }
    }
    cout << "No" << endl;
}