H = int(input())
day, H_hatsuga = 0, 0

while H_hatsuga <= H:
    H_hatsuga += 2 ** day
    day += 1
print(day)