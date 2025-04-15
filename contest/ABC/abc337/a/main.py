N = int(input())
x_point, y_point = 0, 0

for i in range(N):
    X, Y = map(int, input().split())
    x_point += X
    y_point += Y
    
if x_point > y_point:
    print("Takahashi")
elif x_point < y_point:
    print("Aoki")
else:
    print("Draw")