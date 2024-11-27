import sys
import copy
from collections import deque

input = sys.stdin.readline

n = int(input())
arr1 = list(map(int, input().split()))
m = int(input())
arr2 = list(map(int, input().split()))

dp = [[0 for j in range((i+1)*500 +1)] for i in range(n+1)]
ans = []

def sol(i, target):
    if i > n:
        return
    if dp[i][target]:
        return

    dp[i][target] = 1

    sol(i+1, target)
    sol(i+1, target + arr1[i-1])
    sol(i+1, abs(target - arr1[i-1]))
    
sol(0,0)

for i in arr2:
    if i > 500*30:
        ans.append("N")
    elif dp[n][i] == 1:
        ans.append("Y")
    else:
        ans.append("N")

print(" ".join(ans))