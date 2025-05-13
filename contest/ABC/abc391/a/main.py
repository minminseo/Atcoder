D = input()
ans = ""

for i in D:
    if i == "N":
        ans += "S"
    elif i == "S":
        ans += "N"
    elif i == "E":
        ans += "W"
    elif i == "W":
        ans += "E"

print(ans)
