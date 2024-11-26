import sys
from collections import deque

input = sys.stdin.readline

m, n = map(int, input().split())
route = [list(map(int, input().split())) for i in range(m)]

dp = [[-1]*n for i in range(m)]
dx = [0, 0, 1, -1]
dy = [1, -1, 0, 0]

def count_route(x,y):
    stack = deque()
    stack.append((x,y))
    dp[x][y] = 0

    while stack:
        cx, cy = stack[-1]

        if cx == m-1 and cy == n-1:
            dp[cx][cy] = 1
            stack.pop()
            continue

        allChildProcessed = True
        for i in range(4):
            nx, ny = cx + dx[i], cy + dy[i]
            if 0 <= nx < m and 0 <= ny < n and route[nx][ny] < route[cx][cy]:
                if dp[nx][ny] == -1:
                    stack.append((nx, ny))
                    allChildProcessed = False
                    
        if allChildProcessed:
            dp[cx][cy] = 0
            for i in range(4):
                nx, ny = cx + dx[i], cy + dy[i]
                if 0 <= nx < m and 0 <= ny < n and route[nx][ny] < route[cx][cy]:
                    dp[cx][cy] += dp[nx][ny]
            stack.pop()
    return dp[x][y]

print(count_route(0,0))