from collections import deque

capacity = int(input())

values = list(map(int, input().split()))
deque_data = deque([(value, idx + 1) for idx, value in enumerate(values)])

result = []
for _ in range(capacity - 1):
    displacement = deque_data.popleft()
    result.append(displacement[1])
    
    if displacement[0] > 0:
        for _ in range(displacement[0] - 1):
            deque_data.append(deque_data.popleft())
    else:
        for _ in range(abs(displacement[0])):
            deque_data.appendleft(deque_data.pop())

result.append(deque_data.popleft()[1])

print(" ".join(map(str, result)))