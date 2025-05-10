R, X = map(int, input().split())

is_rated = False

if X == 1:
    if 1600 <= R <= 2999:
        is_rated = True

elif X == 2:
    if 1200 <= R <= 2399:
        is_rated = True

if is_rated:
    print("Yes")
else:
    print("No")

# 普通に真偽値使わなくてもええか