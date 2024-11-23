from collections import deque

n = int(input())
isQ = list(map(int, input().split()))
data = list(map(int, input().split()))
m = int(input())
inputs = list(map(int, input().split()))
dq = deque()

result = []

for i in range(n):
    if isQ[i] == 0:
        dq.append(data[i])

for i in inputs:
    dq.appendleft(i)
    result.append(dq.pop())

print(" ".join(map(str, result)))