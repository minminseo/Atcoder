A, B, C, D = map(int, input().split())
# 数字全部異なるとこ見落とした
# if A > C:
#     print("Yes")
# elif A == C:
#     if B > D:
#         print("Yes")
#     else:
#         print("No")
# elif A < C:
#     print("No")


if A > C:
    print("Yes")
elif A < C:
    print("No")
elif B > D:
    print("Yes")
else:
    print("No")
