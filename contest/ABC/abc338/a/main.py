S = input()
if len(S) > 1:
    if S[0].isupper() and S[1:].islower():
        print("Yes")
    else:
        print("No")
else:
    if S.isupper():
        print("Yes")
    else:
        print("No")
        