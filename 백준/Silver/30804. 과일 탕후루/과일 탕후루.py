n = int(input())
arr = list(map(int, input().split()))

maxLen = 0
left = 0
arr1 = {}
count = 0

for right in range(n):
    if arr[right] in arr1:
        arr1[arr[right]] += 1
    else:
        arr1[arr[right]] = 1
        count += 1
    
    while count > 2:
        arr1[arr[left]] -= 1
        if arr1[arr[left]] == 0:
            del arr1[arr[left]]
            count -= 1
        left += 1
    
    maxLen = max(maxLen, right - left + 1)

print(maxLen)