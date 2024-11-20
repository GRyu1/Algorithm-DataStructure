import math

sum_value = 0

n = int(input())
size_amount = map(int, input().split())
t,p = map(int, input().split())

for value in size_amount:
    sum_value += math.ceil(value / t)

print(sum_value)
print(f"{n // p} {n % p}")