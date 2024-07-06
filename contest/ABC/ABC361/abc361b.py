#自力

a, b, c, d, e, f = map(int, input().split())
g, h, i, j, k, l = map(int, input().split())

def function(m1, n1, m2, n2):
    return  (m1 >= n2 or n1 <= m2)

if function(a, d, g, j) or function(b, e, h, k) or function(c, f, i, l):
    print("No")
else:
    print("Yes")