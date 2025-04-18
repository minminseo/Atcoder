#include <iostream>
#include <vector>
#include <string>

int main()
{
    int H, W;
    std::cin >> H >> W;

    std::vector<std::string> table(H);
    for (int i = 0; i < H; ++i)
    {
        std::cin >> table.at(i);
    }

    std::vector<bool> row(H, false);
    std::vector<bool> col(W, false);
    for (int i = 0; i < H; ++i)
    {
        for (int j = 0; j < W; ++j)
        {
            if (table.at(i).at(j) == '#')
            {
                row.at(i) = true;
                col.at(j) = true;
            }
        }
    }

    for (int i = 0; i < H; ++i)
    {
        if (!row.at(i))
            continue;
        std::string output;
        for (int j = 0; j < W; ++j)
        {
            if (col.at(j))
                output += table.at(i).at(j);
        }
        std::cout << output << std::endl;
    }
}
