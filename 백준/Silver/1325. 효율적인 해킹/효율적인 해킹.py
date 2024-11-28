from collections import deque, defaultdict

n, m = map(int, input().split())
graph = [[] for _ in range(n + 1)]

for i in range(m):
    a, b = map(int, input().split())
    graph[b].append(a)

dp = [0] * (n + 1)
max_hacks = -1

def solve(start):
    global max_hacks
    is_visited = [False] * (n + 1)
    queue = deque([start])
    is_visited[start] = True
    hack_count = 0

    while queue:
        current = queue.popleft()
        for next_node in graph[current]:
            if not is_visited[next_node]:
                queue.append(next_node)
                is_visited[next_node] = True
                hack_count += 1

    dp[start] = hack_count
    max_hacks = max(max_hacks, hack_count)

for i in range(1, n + 1):
    solve(i)

for i in range(1, n+1):
    if dp[i] == max_hacks:
        print(i, end=' ')
